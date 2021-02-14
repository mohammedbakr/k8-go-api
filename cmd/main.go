package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/api/rebuild/file", rebuildFile).Methods("POST")
	r.HandleFunc("/api/rebuild/base64", rebuildBase64).Methods("POST")
	http.ListenAndServe(":8000", r)
}

func main() {
	setupRoutes()
}
