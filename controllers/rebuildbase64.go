package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/k8-proxy/k8-go-api/models"
	"github.com/k8-proxy/k8-go-api/pkg/message"
	"github.com/k8-proxy/k8-go-api/pkg/store"
	"github.com/k8-proxy/k8-go-api/utils"
)

// RebuildBase64 Rebuilds a file using the Base64 encoded representation
func RebuildBase64(w http.ResponseWriter, r *http.Request) {
	var base64 models.Base64

	err := json.NewDecoder(r.Body).Decode(&base64)
	if err != nil {
		log.Println("json unmarshal", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate Base64
	if base64.Request.Base64 == "" {
		log.Println("base64 empty")

		utils.ResponseWithError(w, http.StatusBadRequest, "Base64 is required")

		return
	}

	// Using Regex
	base64regex := regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)
	match := base64regex.MatchString(base64.Request.Base64)
	if !match {
		log.Println("malformed base64 input")

		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid Base64 format")
		return
	}
	// Retun the content as Base64 encoded
	url, err := store.St([]byte("translate this test file "), "pretranslate")
	if err != nil {
		log.Println(err)
	}

	message.AmqpM("auto", "fr", url)
	//GW custom header
	utils.AddGWHeader(w, models.Temp)

	_, err = w.Write([]byte(base64.Request.Base64))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
