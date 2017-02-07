package objectstore_test

import (
	"encoding/json"
	objectstore_ "github.com/sivaramsk/cloud-backup/backup-server/objectstore"
	"log"
	//golang_os "os"
	"testing"
)

func TestInitialize(t *testing.T) {
	var os objectstore_.ObjectStore
	if os.Initialize("127.0.0.1:9000", "E3WWG9H129KOKWC1O34O", "itg4ijdOh7KZH0qSRu6fPupgXVza1uXg0EHtmn1R", true) != nil {
		t.Error("Error Initializing objectstore")
	}
}

func TestCreateBucket(t *testing.T) {
	var os objectstore_.ObjectStore
	bucketName := "tbtb"
	if os.Initialize("127.0.0.1:9000", "E3WWG9H129KOKWC1O34O", "itg4ijdOh7KZH0qSRu6fPupgXVza1uXg0EHtmn1R", false) != nil {
		t.Error("Error Initializing objectstore")
	}

	os.DeleteBucket(bucketName)

	err := os.CreateBucket(bucketName)
	if err != nil {
		log.Println("Bucket creation failed:", err)
		t.Error("Bucket creation failed")
	} else {
		log.Println("Now that the bucket creation had sucedded, delete the bucket")
		os.DeleteBucket(bucketName)
	}

}

/*

sample.json

{
    name: "backupname",
    target: "<ip of the target machine or FQDN>",
    backupsrc: "<folder to be backup>",
    schedule: "<hourly/daily/weekly>",
    password: "<restic password>"
}

*/

func TestPutGetObject(t *testing.T) {

	// Initialize the test bucket and object name for the test
	// TODO: Have test do a different objectname length and multiple bucket names
	bucketName := "testbucket"
	objectName := "backup/testjob1/config"
	/*
		Configuration: config/testjob1/schedule
		Logs: config/testjob1/logs/<timestamp>/{job_status}
	*/

	type BackupJob struct {
		BackupName     string
		Target         string
		Backupfolder   string
		BackupSchedule string
		BackupPassword string
	}

	testBackup := BackupJob{BackupName: "testjob1", Target: "testserver.local.in", Backupfolder: "/etc", BackupSchedule: "daily", BackupPassword: "testpassword"}

	// Create a test bucke first
	var os objectstore_.ObjectStore
	if os.Initialize("127.0.0.1:9000", "E3WWG9H129KOKWC1O34O", "itg4ijdOh7KZH0qSRu6fPupgXVza1uXg0EHtmn1R", false) != nil {
		t.Error("Error Initializing objectstore")
	}

	// Delete the bucket, if it already exists
	os.DeleteBucket(bucketName)

	err := os.CreateBucket(bucketName)
	if err != nil {
		t.Error("Bucket creation failed", err)
	}

	// Create a json object
	j, err := json.Marshal(testBackup)
	if err != nil {
		t.Error("Error marshalling json object")
	}

	log.Println("Testing PutObject....")
	// Put the object in to the bucket
	err = os.PutObject(bucketName, objectName, j)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Testing GetObject....")
	// Get the object uploaded above
	out, err := os.GetObject(bucketName, objectName)
	if err != nil {
		log.Fatalln(err)
	}

	var testJson BackupJob
	err = json.Unmarshal(out, &testJson)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", testJson)

	log.Println("Delete the test object", objectName)
	// Delete the Created object
	err = os.MinioObject.RemoveObject(bucketName, objectName)
	if err != nil {
		log.Fatalln(err)
	}

}
