package rabbitmq

import (
	"fmt"
	"net/url"

	"github.com/streadway/amqp"
)

type RMQConnection struct {
	*amqp.Connection
}

type RMQChannel struct {
	*amqp.Channel
}

func NewInstance(rabbitHost string, rabbitPort string, messagebrokeruser string, messagebrokerpassword string) (*amqp.Connection, error) {

	if messagebrokeruser == "" {
		messagebrokeruser = "guest"
	}

	if messagebrokerpassword == "" {
		messagebrokerpassword = "guest"
	}

	amqpUrl := url.URL{
		Scheme: "amqp",
		User:   url.UserPassword(messagebrokeruser, messagebrokerpassword),
		Host:   fmt.Sprintf("%s:%s", rabbitHost, rabbitPort),
		Path:   "/",
	}
	conn, err := amqp.Dial(amqpUrl.String())
	if err != nil {
		return conn, err
	}

	return conn, err

}

func NewQueueConsumer(connection *amqp.Connection, queueName string, exchange string, routingKey string) (<-chan amqp.Delivery, *amqp.Channel, error) {

	ch, err := connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	err = ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		return nil, nil, err
	}

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return nil, nil, err
	}

	err = ch.QueueBind(q.Name, routingKey, exchange, false, nil)
	if err != nil {
		return nil, nil, err
	}

	consumer, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	return consumer, ch, err

}

func NewQueuePublisher(connection *amqp.Connection, exchange string) (*amqp.Channel, error) {

	channel, err := connection.Channel()
	if err != nil {
		return channel, err
	}

	if err := channel.ExchangeDeclare(
		exchange, // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // noWait
		nil,      // arguments
	); err != nil {
		return channel, err
	}

	return channel, nil

}

func PublishMessage(channel *amqp.Channel, exchange string, routingKey string, header amqp.Table, message []byte) error {

	err := channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         header, ////????
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            message,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	)

	return err

}
