package controllers

import (
	"encoding/json"
	"io/ioutil"
	"k8-go-api/models"
	"k8-go-api/utils"
	"log"
	"net/http"

	"github.com/rs/zerolog/hlog"
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
	if err != nil {
		log.Println("ioutilReadAll", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "file not found")
		return
	}

	if handler.Header.Get("Content-Type") != "application/zip" || http.DetectContentType(buf) != "application/zip" {
		log.Println("mediatype is", handler.Header.Get("Content-Type"))
		//utils.ResponseWithError(w, http.StatusUnsupportedMediaType, "uploaded file should be zip format")
		//return
	}

	hlog.FromRequest(r).Info().
		Str("Filename", handler.Filename).
		Int64("Filesize", int64(handler.Size)).
		Str("Content-Type", handler.Header.Get("Content-Type")).
		Msg("after response")

	//uploaded file log info
	/*
		log.Printf("Filename: %v\n", handler.Filename)
		log.Printf("File size: %v\n", handler.Size)
		log.Printf("Content-Type: %v\n", handler.Header.Get("Content-Type"))
		log.Printf("Content-Type: %v\n", http.DetectContentType(buf))
	*/
	//GW custom header
	utils.AddGWHeader(w, models.Temp)

	_, e := w.Write(buf)
	if e != nil {
		log.Println(e)
		return
	}
}
