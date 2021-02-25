package controllers

import (
	"encoding/json"
	"io/ioutil"
	"k8-go-api/models"
	"k8-go-api/utils"
	"log"
	"net/http"

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

	//GW custom header
	utils.AddGWHeader(w, models.Temp)

	_, e := w.Write(buf)
	if e != nil {
		log.Println(e)
		return
	}
}
