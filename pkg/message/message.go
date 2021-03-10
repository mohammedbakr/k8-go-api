package message

import (
	"log"

	"github.com/streadway/amqp"

	"github.com/k8-proxy/k8-go-comm/pkg/rabbitmq"
)

var (
	exchange   = "transalte-exchange"
	routingKey = "transalte-request"
	queueName  = "transalte-queue"
)

func AmqpM(conn *amqp.Connection, source, target, url string) string {

	publisher, err := rabbitmq.NewQueuePublisher(conn, exchange)
	if err != nil {
		log.Fatalf("%s", err)
	}

	defer publisher.Close()

	// Start a consumer
	msgs, ch, err := rabbitmq.NewQueueConsumer(conn, queueName, exchange, routingKey)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer ch.Close()

	table := amqp.Table{
		"sourcelanguage": source,
		"targetlanguage": target,
	}

	err = rabbitmq.PublishMessage(publisher, exchange, routingKey, table, []byte(url))
	if err != nil {
		log.Println("PublishMessage", err)

		return ""
	}

	var miniourl string

	notforever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			miniourl = string(d.Body)
			notforever <- true
		}
	}()
	<-notforever

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	log.Println(miniourl)
	return miniourl
}
