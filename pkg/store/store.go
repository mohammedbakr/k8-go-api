package store

import (
	"bytes"
	"log"
	"time"

	"github.com/k8-proxy/k8-go-comm/pkg/minio"
	min "github.com/minio/minio-go/v7"
)

var (
	minioEndpoint     = "localhost:9000"
	minioAccessKey    = "minioadmin"
	minioSecretKey    = "minioadmin"
	sourceMinioBucket = "test"
)
var (
	cl *min.Client
)

func init() {
	var err error
	cl, err = minio.NewMinioClient(minioEndpoint, minioAccessKey, minioSecretKey, false)
	if err != nil {
		log.Println(err)
	}

}

func St(file []byte, filename string) (string, error) {
	exist, err := minio.CheckIfBucketExists(cl, sourceMinioBucket)
	if err != nil || !exist {
		log.Println("error checkbucket ", err)
		err = minio.CreateNewBucket(cl, "test")
		if err != nil {
			log.Println(err)
			return "", err
		}

	}
	_, errm := minio.UploadFileToMinio(cl, sourceMinioBucket, filename, bytes.NewReader(file))
	if errm != nil {
		log.Println(errm)
		return "", errm
	}
	expirein := time.Second * 24 * 60 * 60
	urlx, err := minio.GetPresignedURLForObject(cl, sourceMinioBucket, filename, expirein)
	if err != nil {
		log.Println(err)
		return "", err

	}
	return urlx.String(), nil

}
