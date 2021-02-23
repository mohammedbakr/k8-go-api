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

	//uploaded file log info
	hlog.FromRequest(r).Info().
		Str("Filename", handler.Filename).
		Int64("Filesize", int64(handler.Size)).
		Str("Content-Type", handler.Header.Get("Content-Type")).
		Msg("after response")
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
