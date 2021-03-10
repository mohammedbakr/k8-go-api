package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k8-proxy/k8-go-api/controllers"
	"github.com/k8-proxy/k8-go-api/middleware"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	mx := mux.NewRouter()
	mx.Use(middleware.LogMiddleware, middleware.AuthMiddleware)

	mx.HandleFunc("/api/rebuild/file", controllers.RebuildFile).Methods("POST")
	mx.HandleFunc("/api/rebuild/zip", controllers.Rebuildzip).Methods("POST")
	mx.HandleFunc("/api/rebuild/base64", controllers.RebuildBase64).Methods("POST")

	fmt.Println("Server is ready to handle requests at port 8100")
	log.Fatal(http.ListenAndServe(":8100", mx))
}
