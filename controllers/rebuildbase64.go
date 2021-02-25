package controllers

import (
	b64 "encoding/base64"
	"encoding/json"
	"k8-go-api/models"
	"k8-go-api/utils"
	"log"
	"net/http"
)

const ()

// RebuildBase64 Rebuilds a file using the Base64 encoded representation
func RebuildBase64(w http.ResponseWriter, r *http.Request) {
	var base64 models.Base64

	err := json.NewDecoder(r.Body).Decode(&base64)
	if err != nil {
		log.Println("json decode: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate Base64
	if base64.Request.Base64 == "" {
		log.Println("empty base64: ")
		utils.ResponseWithError(w, http.StatusBadRequest, "Base64 is required")
		return
	}

	// Retun the content as Base64 encoded
	contentEncoded, err := b64.StdEncoding.DecodeString(base64.Request.Base64)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//GW custom header
	utils.AddGWHeader(w, models.Temp)

	_, err = w.Write(contentEncoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
