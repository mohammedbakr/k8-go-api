package middleware

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rs/zerolog/hlog"
)

func TestAuthMiddleware(t *testing.T) {

	var tests = []struct {
		hk     string
		hv     string
		status int
	}{
		{"Authorization", "Bearer mysecrettoken", http.StatusOK},
		{"malformed", "Bearer mysecrettoken", http.StatusUnauthorized},
		{"Authorization", "earer mysecrettoken", http.StatusUnauthorized},
		{"Authorization", "Bearer notmysecrettoken", http.StatusUnauthorized},
		{"", "", http.StatusUnauthorized},
	}

	for _, test := range tests {
		if output := authmiddlewareconnect(test.hk, test.hv); output != test.status {
			t.Errorf("Test Failed: {%s} header, {%s} header value, status: {%d},expected: {%d}", test.hk, test.hv, test.status, output)
		}
	}
}

func authmiddlewareconnect(hk, hv string) int {

	endpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "autorization succeful")
	})

	ts := httptest.NewServer(AuthMiddleware(endpoint))

	client := &http.Client{}

	req, err := http.NewRequest("POST", ts.URL, strings.NewReader("empty message"))
	if hk != "" || hv != "" {
		req.Header.Set(hk, hv)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	status := resp.StatusCode

	return status
}

func TestLogMiddleware(t *testing.T) {

	if output := logmiddlewareconnect(); output != http.StatusOK {
		t.Errorf("Test Failed:  status: {%d},expected: {%d}", http.StatusOK, output)
	}

}

func logmiddlewareconnect() int {

	endpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hlog.FromRequest(r).Info().
			Str("status", "ok").
			Msg("Something happened")
	})

	ts := httptest.NewServer(LogMiddleware(endpoint))

	client := &http.Client{}

	req, err := http.NewRequest("POST", ts.URL, strings.NewReader("empty message"))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	status := resp.StatusCode

	return status
}
