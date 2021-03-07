package message

import (
	"log"

	"github.com/streadway/amqp"

	"github.com/k8-proxy/k8-go-api/pkg/rabbitmq"
)

var (
	exchange   = "adaptation-exchange"
	routingKey = "adaptation-request"
	queueName  = "adaptation-request-queue"

	processing_exchange   = "processing-exchange"
	processing_routingKey = "processing-request"
	processing_queueName  = "processing-queue"
)
var (
	conn *amqp.Connection
)

func Conn() *amqp.Connection {
	return conn
}
func init() {
	adaptationRequestQueueHostname := "localhost"
	adaptationRequestQueuePort := "5672"

	var err error
	conn, err = rabbitmq.NewInstance(adaptationRequestQueueHostname, adaptationRequestQueuePort, "", "")
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func AmqpM() {

	publisher, err := rabbitmq.NewQueuePublisher(conn, exchange)
	if err != nil {
		log.Fatalf("%s", err)
	}

	defer publisher.Close()

	// Start a consumer
	msgs, ch, err := rabbitmq.NewQueueConsumer(conn, processing_queueName, processing_exchange, processing_routingKey)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer ch.Close()

	table := amqp.Table{
		"file-id":               "myfileid",
		"source-file-location":  "/home/ibrahim/my_work/k8-go-api/sampledata/file.pdf",
		"rebuilt-file-location": "sampledata/file.zip",
	}

	err = rabbitmq.PublishMessage(publisher, exchange, routingKey, table, []byte("ibrahim"))
	if err != nil {
		log.Println("PublishMessage", err)

		return
	}

	var miniourl string

	notforever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			miniourl = d.Headers["source-presigned-url"].(string)
			notforever <- true
		}
	}()
	<-notforever

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	log.Println(miniourl)
}
