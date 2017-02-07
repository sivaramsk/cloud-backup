package main

import (
	"encoding/json"
	"fmt"
	"github.com/sivaramsk/cloud-backup/backup-server/objectstore"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	//"github.com/minio/minio-go"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ListBackups)
	router.HandleFunc("/backup", ConfigBackup)
	router.HandleFunc("/backup/{backupId}", ListBackupById)
	router.HandleFunc("/backup/{backupId}", DeleteBackupById).Methods("GET")
	//router.HandleFunc("/sshkey", ConfigSSHKey).Methods("POST")
	//router.HandleFunc("/sshkey", ConfigSSHKey).Methods("GET")

	log.Println(http.ListenAndServe(":8080", router))
}

func ListBackups(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "ListBackup!")

	doneCh := make(chan struct{})

	obs, err := getBackendStoreObject()
	if err != nil {
		log.Fatalln("Error getting objectstore instance")
	}

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := obs.MinioObject.ListObjectsV2("cloudbackup", "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}

}

/*
   Configuration: config/testjob1/schedule
   Logs: config/testjob1/logs/<timestamp>/{job_status}
*/

func ConfigBackup(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "ConfigBackup!")

	// Get the backend object
	obs, err := getBackendStoreObject()
	if err != nil {
		log.Fatalln("Error getting objectStore instance")
	}

	if obs.IsBucketExists("cloudbackup") != true {
		log.Fatalln("The bucket" + "cloudbackup" + "does not exists")
	}

	w.Header().Set("Content-Type", "application/json")

	var m BackupJob
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Error reading json.")
	}

	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Fatalln("Error unmarshaling data")
	}

	// Get the objectName
	objectName := "config/" + m.BackupName + "/schedule"

	// Set the config data against the key objectName
	err = obs.PutObject("backupserver", objectName, b)
	if err != nil {
		log.Fatalln("Error Uploading data to the backend store")
	}

}

func ListBackupById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	backupId := vars["backupId"]
	fmt.Fprintln(w, "ListBackupById:", backupId)
}

func DeleteBackupById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	backupId := vars["backupId"]
	fmt.Fprintln(w, "DeleteBackupById!", backupId)
}

func ConfigSSHKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	backupId := vars["backupId"]
	fmt.Fprintln(w, "DeleteBackupById!", backupId)
}

func GetSSHKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	backupId := vars["backupId"]
	fmt.Fprintln(w, "DeleteBackupById!", backupId)
}

func getBackendStoreObject() (objectstore.ObjectStore, error) {
	endpoint := os.Getenv("OBJECSTORE_ENDPOINT")
	accessKeyId := os.Getenv("ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("SECRET_ACCESS_KEY")
	useSSL := false

	var obs objectstore.ObjectStore
	err := obs.Initialize(endpoint, accessKeyId, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln("Error initializing ObjectStore")
	}

	return obs, nil
}
