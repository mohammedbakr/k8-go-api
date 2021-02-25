package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/mohammedbakr/k8-proxy/k8-go-api/utils"
)

// AuthMiddleware to check authorization
func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errauth := "you don't have  valid authoriaztion token"
		erremptyauth := "you didn't provide authoriaztion token"

		//log about request
		//there will be logging middleware soon
		log.Printf("method: %v\n", r.Method)
		log.Printf("URL: %v\n", r.URL)
		log.Printf("RemoteAddr: %v\n", r.RemoteAddr)
		log.Printf("Host: %v\n", r.Host)
		log.Printf("Content-Type: %v\n", r.Header.Get("Content-Type"))
		log.Printf("RequestURI: %v\n", r.RequestURI)

		//Authorization: Bearer
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ResponseWithError(w, http.StatusUnauthorized, erremptyauth)
			return
		}

		authHeaderParts := strings.Fields(authHeader)
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			utils.ResponseWithError(w, http.StatusUnauthorized, errauth)
			return
		}

		if authHeaderParts[1] != "mysecrettoken" {
			utils.ResponseWithError(w, http.StatusUnauthorized, errauth)
			return
		}

		next.ServeHTTP(w, r)
	})
}
