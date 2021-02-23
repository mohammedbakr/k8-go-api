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

const (
	contentManagementFlagJson = `{
		"PdfContentManagement": {
		  "Metadata": 1,
		  "InternalHyperlinks": 1,
		  "ExternalHyperlinks": 1,
		  "EmbeddedFiles": 1,
		  "EmbeddedImages": 1,
		  "Javascript": 1,
		  "Acroform": 1,
		  "ActionsAll": 1
		},
		"ExcelContentManagement": {
		  "Metadata": 1,
		  "InternalHyperlinks": 1,
		  "ExternalHyperlinks": 1,
		  "EmbeddedFiles": 1,
		  "EmbeddedImages": 1,
		  "DynamicDataExchange": 1,
		  "Macros": 1,
		  "ReviewComments": 1
		},
		"PowerPointContentManagement": {
		  "Metadata": 1,
		  "InternalHyperlinks": 1,
		  "ExternalHyperlinks": 1,
		  "EmbeddedFiles": 1,
		  "EmbeddedImages": 1,
		  "Macros": 1,
		  "ReviewComments": 1
		},
		"WordContentManagement": {
		  "Metadata": 1,
		  "InternalHyperlinks": 1,
		  "ExternalHyperlinks": 1,
		  "EmbeddedFiles": 1,
		  "EmbeddedImages": 1,
		  "DynamicDataExchange": 1,
		  "Macros": 1,
		  "ReviewComments": 1
		}
	  }`
)

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
		if output, _ := rebuildfileconnect(test.flags, "file.pdf"); output != test.status {
			t.Errorf("Test Failed: {%s} flags, {%d} status value, output: {%d}", test.flags, test.status, output)

		}
	}

}
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
		if output, _ := rebuildfileconnect(test.flags, "file.zip"); output != test.status {
			t.Errorf("Test Failed: {%s} flags, {%d} status value, output: {%d}", test.flags, test.status, output)

		}
	}

}

func rebuildfileconnect(flag, filename string) (int, string) {

	endpoint := http.HandlerFunc(RebuildFile)

	ts := httptest.NewServer(endpoint)
	fpath := fmt.Sprintf("/home/ibrahim/sampledata/%s", filename)
	client := &http.Client{}

	//req, err := http.NewRequest("POST", ts.URL, strings.NewReader("empty message"))

	extraParams := map[string]string{
		"contentManagementFlagJson": contentManagementFlagJson,
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

	status := resp.StatusCode
	statusm := resp.Status

	return status, statusm
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

/*
   var tests = []struct {
       input    int
       expected int
   }{
       {2, 4},
       {-1, 1},
       {0, 2},
       {-5, -3},
       {99999, 100001},
   }

   for _, test := range tests {
       if output := Calculate(test.input); output != test.expected {
           t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
       }
   }
*/
