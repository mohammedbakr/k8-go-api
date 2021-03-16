package router

import "net/http"

type Router interface {
	POST(url string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}
