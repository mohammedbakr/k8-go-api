package main

import (
	"fmt"
	"k8-go-api/controllers"
	"k8-go-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mx := mux.NewRouter()
	mx.Use(middleware.AuthMiddleware)

	mx.HandleFunc("/api/rebuild/file", controllers.RebuildFile).Methods("POST")
	mx.HandleFunc("/api/rebuild/zip", controllers.Rebuildzip).Methods("POST")
	mx.HandleFunc("/api/rebuild/base64", controllers.RebuildBase64).Methods("POST")

	fmt.Println("Server is ready to handle requests at port 8000")
	log.Fatal(http.ListenAndServe(":8000", mx))
}
