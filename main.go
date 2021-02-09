package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"k8-go-api/utils"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `file`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	// Create directory
	// Check existence of the directory
	_, err = os.Stat("files")
	if os.IsNotExist(err) {
		errDir := os.Mkdir("files", 0755)
		if errDir != nil {
			utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	dst, err := os.Create(filepath.Join("files", filepath.Base(handler.Filename))) // dir is directory where you want to save file.
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// read all of the contents of our uploaded file into a
	// byte slice
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// write this byte array to our new file
	dst.Write(fileBytes)
	// return that we have successfully uploaded our file!
	utils.ResponseJSON(w, "Successfully Uploaded File.")
}
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8000", nil)
}

func main() {
	setupRoutes()
}
