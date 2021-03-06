package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/k8-proxy/k8-go-api/models"
	"github.com/k8-proxy/k8-go-api/utils"

	"github.com/rs/zerolog"
)

// Rebuildzip processes a zip uploaded by the user, returns a zip file with rebuilt files
func Rebuildzip(w http.ResponseWriter, r *http.Request) {
	//handling json , not implemeted yet
	//log.Println(r.PostFormValue("contentManagementFlagJson"))

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
		utils.ResponseWithError(w, http.StatusBadRequest, "file not found or wrong form field  name")
		return
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)

	bufsign := buf[:511]

	if err != nil {
		log.Println("ioutilReadAll", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "file not found")
		return
	}

	if handler.Header.Get("Content-Type") != "application/zip" || http.DetectContentType(bufsign) != "application/zip" {
		log.Println("mediatype is", handler.Header.Get("Content-Type"))
		utils.ResponseWithError(w, http.StatusUnsupportedMediaType, "uploaded file should be zip format")
		return
	}

	logf := zerolog.Ctx(r.Context())
	logf.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("Filename", handler.Filename).
			Int64("Filesize", handler.Size).
			Str("Content-Type", handler.Header.Get("Content-Type"))

	})

	//GW custom header
	utils.AddGWHeader(w, models.Temp)

	_, e := w.Write(buf)
	if e != nil {
		log.Println(e)
		return
	}
}
