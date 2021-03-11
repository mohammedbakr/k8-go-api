package controllers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

//ignore enpoint tests for the moment until we restructure them to itegration tests
func TestRebuildfile(t *testing.T) {
}
func TestRebuildzip(t *testing.T) {
}

/*
func TestRebuildfile(t *testing.T) {

	var tests = []struct {
		flags  string
		status int
	}{
		{"NORMAL", http.StatusOK},
		{"EMPTY", http.StatusBadRequest},
		{"DIFNAME", http.StatusBadRequest},
	}

	for _, test := range tests {
		if output := rebuildconnect(test.flags, PdfFileName); output != test.status {
			t.Errorf("Test Failed: {%s} flags, {%d} status value, output: {%d}", test.flags, test.status, output)

		}
	}

}
*/
/*
func TestRebuildzip(t *testing.T) {

	var tests = []struct {
		flags  string
		status int
	}{
		{"NORMAL", http.StatusOK},
		{"EMPTY", http.StatusBadRequest},
		{"DIFNAME", http.StatusBadRequest},
	}

	for _, test := range tests {
		if output := rebuildconnect(test.flags, ZipFileName); output != test.status {
			t.Errorf("Test Failed: {%s} flags, {%d} status value, output: {%d}", test.flags, test.status, output)

		}
	}

}
*/
func rebuildconnect(flag, filename string) int {
	var endpoint http.HandlerFunc
	if fileExt(filename) == "zip" {
		endpoint = http.HandlerFunc(Rebuildzip)

	} else {
		endpoint = http.HandlerFunc(RebuildFile)

	}

	ts := httptest.NewServer(endpoint)
	fpath := fmt.Sprintf("%s%s", SampleDataPath, filename)
	client := &http.Client{}

	//req, err := http.NewRequest("POST", ts.URL, strings.NewReader("empty message"))

	extraParams := map[string]string{
		"contentManagementFlagJson": contentManagementFlagJSON,
	}

	var req *http.Request
	var err error
	switch flag {
	case "EMPTY":
		req, err = newfileUploadRequestEmpty(ts.URL, extraParams)
		if err != nil {
			log.Fatal(err)
		}
	case "NORMAL":
		req, err = newfileUploadRequest(ts.URL, extraParams, "file", fpath)
		if err != nil {
			log.Fatal(err)
		}
	case "DIFNAME":
		req, err = newfileUploadRequest(ts.URL, extraParams, "diffnamefile", fpath)
		if err != nil {
			log.Fatal(err)
		}

	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	status := resp.StatusCode

	return status
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))

	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func newfileUploadRequestEmpty(uri string, params map[string]string) (*http.Request, error) {

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func fileExt(s string) string {
	if len(s) > 4 {
		return s[2:]
	}
	return ""
}
