package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	GwMetricFormFileRead = "gw-metric-formfileread"
	GwMetricFileSize     = "gw-metric-filesize"
	GwVersion            = "gw-version"
	GwMetricDetect       = "gw-metric-detect"
	GwMetricRebuild      = "gw-metric-rebuild"
)

var (
	temp = gwcustomheader{"0.01", "5 mb", "1.39", "0.02", "0.03"}
)

type gwcustomheader struct {
	metricFormFileread string
	metricFileSize     string
	version            string
	metricDetect       string
	metricRebuild      string
}

func addgwheader(w http.ResponseWriter, v gwcustomheader) {
	w.Header().Set(GwMetricFormFileRead, v.metricFormFileread)
	w.Header().Set(GwMetricFileSize, v.metricFileSize)
	w.Header().Set(GwVersion, v.version)
	w.Header().Set(GwMetricDetect, v.metricDetect)
	w.Header().Set(GwMetricRebuild, v.metricRebuild)

}

//there will middleware chain here
func customMiddleware(next http.Handler) http.Handler {
	errauth := "you d'ont have valid authoriaztion token"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Authorization: Bearer
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, errauth, http.StatusUnauthorized)

			return
		}

		authHeaderParts := strings.Fields(authHeader)

		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			http.Error(w, errauth, http.StatusUnauthorized)

			return
		}

		if authHeaderParts[1] != "mysecrettoken" {

			http.Error(w, errauth, http.StatusUnauthorized)

			return
		}

		// some logs
		log.Println(r.RequestURI)
		log.Println("inside middleware")

		next.ServeHTTP(w, r)
	})
}

func main() {
	mx := mux.NewRouter()
	mx.Use(customMiddleware)

	mx.HandleFunc("/api/rebuild/file", rebuildfile).Methods("POST")
	mx.HandleFunc("/api/rebuild/zip", rebuildzip).Methods("POST")
	mx.HandleFunc("/api/rebuild/base64", rebuildbase64).Methods("POST")

	fmt.Println("Server is ready to handle requests at port 8100")
	log.Fatal(http.ListenAndServe(":8100", mx))
}

//github.com/k8-proxy/k8-go-api.git
//"github.com/gorilla/mux"
