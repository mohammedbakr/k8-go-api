package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k8-proxy/k8-go-api/middleware"
)

type muxRouter struct{}

var mx = mux.NewRouter()

func NewMuxRouter() Router {
	return &muxRouter{}
}
func (*muxRouter) POST(url string, f func(w http.ResponseWriter, r *http.Request)) {
	mx.Use(middleware.LogMiddleware, middleware.AuthMiddleware)
	mx.HandleFunc(url, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Println("Server is ready to handle requests at port", port)
	http.ListenAndServe(port, mx)
}
