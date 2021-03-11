package store

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/k8-proxy/k8-go-comm/pkg/minio"
	min "github.com/minio/minio-go/v7"
)

// St for storing in MinIO Bucket
func St(cl *min.Client, file []byte, filename string) (string, error) {
	sourceMinioBucket := os.Getenv("MINIO_SOURCE_BUCKET")
	exist, err := minio.CheckIfBucketExists(cl, sourceMinioBucket)
	if err != nil || !exist {
		log.Println("error checkbucket ", err)
		err = minio.CreateNewBucket(cl, sourceMinioBucket)
		if err != nil {
			log.Println(err)
			return "", err
		}

	}
	_, err = minio.UploadFileToMinio(cl, sourceMinioBucket, filename, bytes.NewReader(file))
	if err != nil {
		log.Println(err)
		return "", err
	}
	expirein := time.Second * 24 * 60 * 60
	urlx, err := minio.GetPresignedURLForObject(cl, sourceMinioBucket, filename, expirein)
	if err != nil {
		log.Println(err)
		return "", err

	}
	return urlx.String(), nil

}

// Getfile to get the file by URL
func Getfile(url string) ([]byte, error) {

	f := []byte{}
	resp, err := http.Get(url)
	if err != nil {
		return f, err
	}
	defer resp.Body.Close()

	f, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return f, err
	}
	return f, nil

}
