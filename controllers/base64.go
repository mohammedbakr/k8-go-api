package controllers

import (
	b64 "encoding/base64"
	"encoding/json"
	"k8-go-api/models"
	"k8-go-api/utils"
	"net/http"
)

// Base64 for controller
type Base64 struct{}

// RebuildBase64 Rebuilds a file using the Base64 encoded representation
func (b Base64) RebuildBase64() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		contentEncoded := b64.URLEncoding.EncodeToString([]byte(base64.Request.Base64))
		_, err = w.Write([]byte(contentEncoded))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
