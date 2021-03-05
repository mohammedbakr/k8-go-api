package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/k8-proxy/k8-go-api/models"
	"github.com/k8-proxy/k8-go-api/utils"
	"github.com/streadway/amqp"

	"github.com/rs/zerolog"

	"github.com/k8-proxy/k8-go-api/pkg/minio"
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

// RebuildFile rebuilds a file using its binary data
func RebuildFile(w http.ResponseWriter, r *http.Request) {

	// max 6 MB file size
	r.ParseMultipartForm(6 << 20)

	// log.Printf("json payload : %v\n", r.PostFormValue("contentManagementFlagJson"))
	cont := r.PostFormValue("contentManagementFlagJson")
	var mp map[string]json.RawMessage
	err := json.Unmarshal([]byte(cont), &mp)
	if err != nil {
		log.Println("unmarshal json:", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "malformed json format")
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("formfile", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "File is required")
		return
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("ioutilReadAll", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "file not found")
		return
	}
	if handler.Filename == "" {

	}

	logf := zerolog.Ctx(r.Context())
	logf.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("Filename", handler.Filename).
			Int64("Filesize", handler.Size).
			Str("Content-Type", handler.Header.Get("Content-Type"))

	})

	/////////////////////////////
	minioEndpoint := "localhost:9000"
	minioAccessKey := "minioadmin"
	minioSecretKey := "minioadmin"
	sourceMinioBucket := "test"

	cl := minio.NewMinioClient(minioEndpoint, minioAccessKey, minioSecretKey, false)
	exist, err := minio.CheckIfBucketExists(cl, sourceMinioBucket)
	if err != nil || !exist {
		log.Println("error checkbucket ", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, "bucket not exist")
		return
	}

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

	/////////////////////////
	//GW custom header
	utils.AddGWHeader(w, models.Temp)

	_, e := w.Write(buf)
	if e != nil {
		log.Println(e)
		return
	}
}
