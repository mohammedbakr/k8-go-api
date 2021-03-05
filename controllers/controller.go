package controllers

import (
	"encoding/json"
	"log"

	"github.com/k8-proxy/k8-go-api/models"
	"github.com/k8-proxy/k8-go-api/pkg/rabbitmq"
	"github.com/streadway/amqp"
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
func parseContentManagementFlagJSON(c []byte) (models.ContentManagementFlags, error) {

	var d models.ContentManagementFlags
	err := json.Unmarshal(c, &d)
	if err != nil {
		log.Println("unmarshall", err)
		return d, err

	}
	return d, nil
}
