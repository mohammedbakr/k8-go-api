package utils

import (
	"encoding/json"
	"net/http"

	"k8-go-api/models"
)

// ResponseWithError to handle errors with JSON
func ResponseWithError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-type", "application/json")
	var error models.Error
	error.Message = message
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}
