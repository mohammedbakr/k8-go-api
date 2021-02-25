package controllers

import (
	b64 "encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/mohammedbakr/k8-go-api/models"
	"github.com/mohammedbakr/k8-go-api/utils"
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

	// Using Regex
	base64regex := regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)
	match := base64regex.MatchString(base64.Request.Base64)
	if !match {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Invalid Base64 format")
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
