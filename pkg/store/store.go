package store

import (
	"log"

	"github.com/k8-proxy/k8-go-api/pkg/minio"
	min "github.com/minio/minio-go/v7"
)

var (
	minioEndpoint     = "localhost:9000"
	minioAccessKey    = "minioadmin"
	minioSecretKey    = "minioadmin"
	sourceMinioBucket = "test"
)
var (
	cl = initclient()
)

func initclient() *min.Client {
	client := minio.NewMinioClient(minioEndpoint, minioAccessKey, minioSecretKey, false)
	return client

}

func St() {
	exist, err := minio.CheckIfBucketExists(cl, sourceMinioBucket)
	if err != nil || !exist {
		log.Println("error checkbucket ", err)
		err = minio.CreateNewBucket(cl, "test")
		if err != nil {
			log.Println(err)
		}

	}
}
