package controllers

import (
	b64 "encoding/base64"
	"encoding/json"
	"k8-go-api/models"
	"k8-go-api/utils"
	"net/http"
)

// RebuildBase64 Rebuilds a file using the Base64 encoded representation
func RebuildBase64(w http.ResponseWriter, r *http.Request) {
	var base64 models.Base64
	err := json.NewDecoder(r.Body).Decode(&base64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Validate Base64
	if base64.Request.Base64 == "" {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Base64 is required")
	}

	// Retun the content as Base64 encoded
	contentEncoded, err := b64.StdEncoding.DecodeString(base64.Request.Base64)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	//GW custom header
	utils.AddGWHeader(w, models.Temp)

	_, err = w.Write(contentEncoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
