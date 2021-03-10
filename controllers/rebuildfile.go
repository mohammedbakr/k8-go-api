package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	// this experemental  , it connect to a translating service process
	connRabMQ, err := rabbitmq.NewInstance(os.Getenv("ADAPTATION_REQUEST_QUEUE_HOSTNAME"), os.Getenv("ADAPTATION_REQUEST_QUEUE_PORT"), os.Getenv("MESSAGE_BROKER_USER"), os.Getenv("MESSAGE_BROKER_PASSWORD"))
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer connRabMQ.Close()

	cl, err := minio.NewMinioClient(os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), false)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	url, err := store.St(cl, buf, "pretranslate")
	if err != nil {
		log.Println(err)
	}

	miniourl := message.AmqpM(connRabMQ, "auto", "es", url)

	buf2, err := getfile(miniourl)
	if err != nil {
		log.Println(err)
	}
	/////////////////////////
	//GW custom header
	utils.AddGWHeader(w, models.Temp)

	_, e := w.Write(buf2)
	if e != nil {
		log.Println(e)
		return
	}
}

func getfile(url string) ([]byte, error) {

	f := []byte{}
	resp, err := http.Get(url)
	if err != nil {
		return f, err
	}
	defer resp.Body.Close()

	f, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return f, err
	}
	return f, nil

}
