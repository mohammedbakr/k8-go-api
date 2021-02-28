package controllers

import (
	b64 "encoding/base64"
	"log"
	"regexp"
	"testing"

	"github.com/k8-proxy/k8-go-api/models"
)

func TestRebuildBase64(t *testing.T) {
	var base64 models.Base64
	// base64 encoded
	base64.Request.Base64 = "bG9yZW0NCg=="
	// base64.Request.Base64 = ""
	// base64.Request.Base64 = "Hello"

	// Validate Base64
	if base64.Request.Base64 == "" {
		log.Fatal("Base64 is required")
	}

	// Using Regex
	base64regex := regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)
	match := base64regex.MatchString(base64.Request.Base64)
	if !match {
		log.Fatal("Invalid Base64 format")
	}

	// Retun the content as Base64 decoded
	_, err := b64.StdEncoding.DecodeString(base64.Request.Base64)
	if err != nil {
		log.Fatal(err)
	}
}
