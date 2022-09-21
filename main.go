package main

import (
	"log"

	"github.com/minio/minio-go/v6"
	"hjfu.com/minio-client/api"
)

const (
	endpoint        = "138.2.92.64:9000"
	accessKeyID     = "minio"
	secretAccessKey = "minio123"
	useSSL          = false
)

func main() {

	minioClient, error := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)

	if error != nil {
		log.Fatalln(error)
	}

	// bucketName := api.CreateBucket(minioClient, "xfchen12")
	// api.FileUpload(minioClient, bucketName)

	c := &api.C{
		Client: *minioClient,
	}

	// c.ListBuckets()
	bucketName := c.CreateBucket("xfchen")
	// c.FileUpload(bucketName)
	c.ListBucketObjects(bucketName)
	c.GetObject("")

}
