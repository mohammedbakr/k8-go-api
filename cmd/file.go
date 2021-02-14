package main

import (
	"io/ioutil"
	"k8-go-api/utils"
	"net/http"
)

// MaxSize => maximmum upload of 6 MB files.
const MaxSize = 6 << 20

func rebuildFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	// Parse our multipart form, MaxSize
	err = r.ParseMultipartForm(MaxSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Read the content of the file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
