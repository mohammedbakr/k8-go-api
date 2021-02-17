package main

import (
	"k8-go-api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() {

	// Init Controllers
	fileC := controllers.File{}
	base64C := controllers.Base64{}

	// Init Routes
	r := mux.NewRouter()

	r.HandleFunc("/api/rebuild/file", fileC.RebuildFile()).Methods("POST")
	r.HandleFunc("/api/rebuild/base64", base64C.RebuildBase64()).Methods("POST")
	http.ListenAndServe(":8000", r)
}

func main() {
	setupRoutes()
}
