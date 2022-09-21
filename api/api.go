package api

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v6"
)

type C struct {
	Client minio.Client
}

func (c C) ListBuckets() {

	buckets, err := c.Client.ListBuckets()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}

}

func (c C) CreateBucket(bucketName string) string {

	location := "us-east-1"

	err := c.Client.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := c.Client.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)
	return bucketName
}

func (c C) FileUpload(bucketName string) {
	// 上传一个zip文件。
	objectName := "Centos-7.repo"
	filePath := "C:\\Users\\admin\\Downloads\\Centos-7.repo"
	// contentType := "application/zip"

	// 使用FPutObject上传一个zip文件。
	n, err := c.Client.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: ""})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}

func (c C) ListBucketObjects(bucketName string) {
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := c.Client.ListObjects(bucketName, "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}
}
