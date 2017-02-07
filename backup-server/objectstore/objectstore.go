package objectstore

import (
	"bytes"
	minio_ "github.com/minio/minio-go"
	"log"
)

type ObjectStore struct {
	Endpoint        string
	accessKeyID     string
	secretAccessKey string
	bucketName      string
	location        string
	useSSL          bool
	configPrefix    string
	storePrefix     string
	MinioObject     *minio_.Client
}

func (obs *ObjectStore) Initialize(endPoint string, accessKeyID string, secretAccessKey string, useSSL bool) error {
	obs.Endpoint = endPoint
	obs.accessKeyID = accessKeyID
	obs.secretAccessKey = secretAccessKey
	obs.useSSL = useSSL

	var err error

	// Initialize minio client object.
	obs.MinioObject, err = minio_.New(obs.Endpoint, obs.accessKeyID, obs.secretAccessKey, obs.useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (obs *ObjectStore) ListBuckets() {
	buckets, err := obs.MinioObject.ListBuckets()
	if err != nil {
		log.Println(err)
		return
	}
	for _, bucket := range buckets {
		log.Println(bucket)
	}

}

func (obs *ObjectStore) DeleteBucket(bucketName string) {

	log.Println("Inside DeleteBucket")
	err := obs.MinioObject.RemoveBucket(bucketName)
	if err != nil {
		log.Println("DeleteBucket failed: ", err)
	}

}

func (obs *ObjectStore) CreateBucket(bucketName string) error {

	log.Println("Inside CreateBucket function in objectstore package: ", bucketName)
	err := obs.MinioObject.MakeBucket(bucketName, "")
	log.Println("MakeBucket return value: ", err)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := obs.MinioObject.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Println("Error testing if bucketexists")
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	return err
}

func (obs *ObjectStore) PutObject(bucketName string, objectName string, data []byte) error {
	// Copy the data to the given object name
	r := bytes.NewReader(data)
	n, err := obs.MinioObject.PutObject(bucketName, objectName, r, "application/octet-stream")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
	return err
}

func (obs *ObjectStore) GetObject(bucketName string, objectName string) ([]byte, error) {

	object, err := obs.MinioObject.GetObject(bucketName, objectName)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert the object to bytes buffer and return bytes[]
	buf := new(bytes.Buffer)
	buf.ReadFrom(object)

	log.Printf("Successfully downloaded %s of size %s\n", objectName, buf)

	return buf.Bytes(), err

}

func (obs *ObjectStore) IsBucketExists(bucketName string) bool {
	found, err := obs.MinioObject.BucketExists(bucketName)
	if err != nil {
		log.Fatalln("The bucket" + bucketName + " not found")
	}

	return found
}
