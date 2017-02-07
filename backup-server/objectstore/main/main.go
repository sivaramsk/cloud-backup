package main

import (
	objectstore_ "github.com/sivaramsk/cloud-backup/backup-server/objectstore"
	"log"
)

func main() {
	var os objectstore_.ObjectStore

	log.Println("Initialize the ObjectStore object")
	if os.Initialize("127.0.0.1:9000", "E3WWG9H129KOKWC1O34O", "itg4ijdOh7KZH0qSRu6fPupgXVza1uXg0EHtmn1R", false) != nil {
		log.Println("Error Initializing objectstore")
	}

	log.Println("Create bucket")
	if os.CreateBucket("tbucket") != nil {
		log.Println("Bucket creation failed")
	}

	log.Println("List bucket")
	os.ListBuckets()
}

/*
import (
	"github.com/minio/minio-go"
	"log"
)

func main() {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "E3WWG9H129KOKWC1O34O"
	secretAccessKey := "itg4ijdOh7KZH0qSRu6fPupgXVza1uXg0EHtmn1R"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucked called mymusic.
	bucketName := "mymusic"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	// Upload the zip file
	objectName := "golden-oldies.zip"
	filePath := "/tmp/golden-oldies.zip"
	contentType := "application/zip"

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, contentType)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
*/
