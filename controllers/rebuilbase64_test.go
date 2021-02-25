package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"k8-go-api/models"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func jsonreqbuild(flag bool) string {

	cont, err := ioutil.ReadFile("/home/ibrahim/Downloads/sample.pdf")
	if err != nil {
		log.Println("ioutilReadAll", err)
		return ""
	}

	str := base64.StdEncoding.EncodeToString(cont)

	var js models.Base64

	js.Request.FileName = "filename"

	js.Request.ContentManagementFlags, err = parsecontentManagementFlagJson([]byte(contentManagementFlagJson))
	if err != nil {
		log.Println("unmarshal", err)
		return ""

	}
	if flag {
		js.Request.Base64 = "()$--"
	} else {
		js.Request.Base64 = str
	}

	res, err := json.Marshal(js)
	if err != nil {
		log.Println("marshall", err)
		return ""
	}

	return string(res)

}

func TestRebuilBase64(t *testing.T) {

	var tests = []struct {
		flags  string
		status int
	}{
		{"NORMAL", http.StatusOK},
		{"EMPTY", http.StatusBadRequest},
		{"MAlFORM", http.StatusBadRequest},
		{"MALFORM64", http.StatusBadRequest},
	}

	for _, test := range tests {
		if output := rebuildBase64connect(test.flags); output != test.status {
			t.Errorf("Test Failed: {%s} flags, {%d} status value, output: {%d}", test.flags, test.status, output)

		}
	}
}

func rebuildBase64connect(flag string) int {

	endpoint := http.HandlerFunc(RebuildBase64)

	ts := httptest.NewServer(endpoint)

	client := &http.Client{}

	//req, err := http.NewRequest("POST", ts.URL, strings.NewReader("empty message"))

	var req *http.Request
	var err error

	var body string
	switch flag {
	case "EMPTY":
		body = ""

	case "NORMAL":
		body = jsonreqbuild(false)

	case "MAlFORM":
		body = jsonreqbuild(false)[1:]

	case "MALFORM64":
		body = jsonreqbuild(true)
	}

	req, err = newBase64UploadRequest(ts.URL, body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	status := resp.StatusCode

	return status
}

func newBase64UploadRequest(uri, jsonstr string) (*http.Request, error) {

	body := &bytes.Buffer{}

	body.ReadFrom(strings.NewReader(jsonstr))

	req, err := http.NewRequest("POST", uri, body)

	return req, err
}
