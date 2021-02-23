package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8-go-api/models"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func exm(flag bool) string {

	cont, err := ioutil.ReadFile("/home/ibrahim/Downloads/sample.pdf")
	if err != nil {
		log.Println("ioutilReadAll", err)

	}

	str := base64.StdEncoding.EncodeToString(cont)

	var js models.Base64

	js.Request.FileName = "filename"

	err = json.Unmarshal([]byte(contentManagementFlagJson), &js.Request.ContentManagementFlags)
	if err != nil {
		log.Println("unlarshall", err)

	}

	if flag {
		js.Request.Base64 = "()$--"
	} else {
		js.Request.Base64 = str
	}

	res, err := json.Marshal(js)
	if err != nil {
		log.Println("marshall", err)

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
		if output, _ := rebuildBase64connect(test.flags); output != test.status {
			t.Errorf("Test Failed: {%s} flags, {%d} status value, output: {%d}", test.flags, test.status, output)

		}
	}
}

func rebuildBase64connect(flag string) (int, string) {

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
		body = exm(false)

	case "MAlFORM":
		body = exm(false)[1:]

	case "MALFORM64":
		body = exm(true)
	}

	req, err = newBase64UploadRequest(ts.URL, body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	status := resp.StatusCode
	statusm := resp.Status

	return status, statusm
}

func newBase64UploadRequest(uri, jsonstr string) (*http.Request, error) {

	body := &bytes.Buffer{}

	body.ReadFrom(strings.NewReader(jsonstr))

	req, err := http.NewRequest("POST", uri, body)

	return req, err
}

func deprec() {
	path, _ := os.Getwd()
	path += "/test.pdf"
	extraParams := map[string]string{
		"title":       "My Document",
		"author":      "Matt Aimonetti",
		"description": "A document with all the Go programming language secrets",
	}
	request, err := newfileUploadRequest("https://google.com/upload", extraParams, "file", "/tmp/doc.pdf")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		fmt.Println(body)
	}
}
