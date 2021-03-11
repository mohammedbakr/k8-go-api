package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/k8-proxy/k8-go-api/models"
	"github.com/k8-proxy/k8-go-api/pkg/store"
	"github.com/k8-proxy/k8-go-api/utils"
	"github.com/k8-proxy/k8-go-comm/pkg/minio"
	"github.com/k8-proxy/k8-go-comm/pkg/rabbitmq"

	"github.com/k8-proxy/k8-go-api/pkg/message"
	"github.com/rs/zerolog"
)

// RebuildFile rebuilds a file using its binary data
func RebuildFile(w http.ResponseWriter, r *http.Request) {

	// max 6 MB file size
	r.ParseMultipartForm(6 << 20)

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

	/////////////////////////////
	// this experemental  , it connect to a translating service process
	connRabMQ, err := rabbitmq.NewInstance(os.Getenv("ADAPTATION_REQUEST_QUEUE_HOSTNAME"), os.Getenv("ADAPTATION_REQUEST_QUEUE_PORT"), os.Getenv("MESSAGE_BROKER_USER"), os.Getenv("MESSAGE_BROKER_PASSWORD"))
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "RabbitMQ Error "+err.Error())
		return
	}
	defer connRabMQ.Close()

	cl, err := minio.NewMinioClient(os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), false)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "MinIO Error "+err.Error())
		return
	}

	timer := time.Now()
	url, err := store.St(cl, buf, "originalpdf")
	if err != nil {
		log.Println(err)
	}
	stsince := time.Since(timer)

	reqid := r.Header.Get("Request-Id")

	timer = time.Now()
	miniourl := message.AmqpM(connRabMQ, reqid, url)
	mqsince := time.Since(timer)

	timer = time.Now()
	buf2, err := store.Getfile(miniourl)
	if err != nil {
		log.Println(err)
	}
	gfsince := time.Since(timer)

	/////////////////////////
	//GW custom header
	utils.AddGWHeader(w, models.Temp)

	_, e := w.Write(buf2)
	if e != nil {
		log.Println(e)
		return
	}

	logf := zerolog.Ctx(r.Context())
	logf.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("Filename", handler.Filename).
			Int64("Filesize", handler.Size).
			Str("Content-Type", handler.Header.Get("Content-Type")).
			Dur("mqduration", mqsince).Dur("minio duration", stsince).Dur("getfil duration", gfsince)

	})
}
