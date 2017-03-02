package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/minio/minio-go"
)

func main() {

	log.Println("Starting backup server framework on port 8080...")
	http.ListenAndServe("0.0.0.0:8080", Handlers())

}

func Handlers() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ListBackups)
	router.HandleFunc("/backup", ConfigBackup).Methods("POST")
	router.HandleFunc("/backup/{backupId}", ListBackupById).Methods("GET")
	router.HandleFunc("/backup/{backupId}", DeleteBackupById).Methods("DELETE")
	router.HandleFunc("/sshkey", ConfigSSHKey).Methods("POST")
	router.HandleFunc("/sshkey", ConfigSSHKey).Methods("GET")
	router.HandleFunc("/help", Help)

	return router
}

func ListBackups(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "ListBackup!")

	doneCh := make(chan struct{})

	obs, err := GetBackendStoreObject()
	if err != nil {
		log.Fatalln("Error getting objectstore instance")
	}

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := obs.MinioObject.ListObjectsV2("cloudbackup", "", isRecursive, doneCh)
	jsonList := make([]string, 0)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		jsonList = append(jsonList, object.Key)
	}

	b, err := json.Marshal(jsonList)
	if err != nil {
		log.Fatalln("Error marshaling data", err)
	}

	// Send the HTTP response back to the client
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintln(w, "ListBackups:", string(b))

}

/*
   Configuration: config/testjob1/schedule
   Logs: config/testjob1/logs/<timestamp>/{job_status}
*/

func ConfigBackup(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "ConfigBackup!")

	// Get the backend object
	obs, err := GetBackendStoreObject()
	if err != nil {
		log.Fatalln("Error getting objectStore instance")
	}

	if obs.IsBucketExists("cloudbackup") != true {
		log.Fatalln("The bucket" + " cloudbackup" + " does not exists")
	}

	b, err := ioutil.ReadAll(r.Body)
	log.Println(b)
	if err != nil {
		log.Fatalln("Error reading json.")
	}

	var m BackupJob
	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Fatalln("Error marshaling data", err)
	}

	log.Println("JsonData: ", m)

	// Get the objectName
	objectName := "config/" + m.BackupName + "/schedule"

	// Set the config data against the key objectName
	err = obs.PutObject("cloudbackup", objectName, b)
	if err != nil {
		log.Fatalln("Error Uploading data to the backend store")
	}

}

func ListBackupById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	backupId := vars["backupId"]

	obs, err := GetBackendStoreObject()
	if err != nil {
		log.Fatalln("Error getting objectstore instance")
	}

	backupId = "config/" + backupId + "/schedule"
	object, err := obs.GetObject("cloudbackup", backupId)
	if err != nil {
		log.Println("Error Reading object", err)
	}

	var testJson BackupJob
	err = json.Unmarshal(object, &testJson)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", testJson)

	// Send the HTTP response back to the client
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(object)

	fmt.Fprintln(w, "ListBackupById:", backupId)
}

func DeleteBackupById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	backupId := vars["backupId"]

	obs, err := GetBackendStoreObject()
	if err != nil {
		log.Fatalln("Error getting objectstore instance")
	}

	backupId = "config/" + backupId + "/schedule"
	err = obs.MinioObject.RemoveObject("cloudbackup", backupId)
	if err != nil {
		log.Println("Error Reading object", err)
	}

	w.WriteHeader(http.StatusOK)

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

func Help(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "/help - This help listing")
	fmt.Fprintln(w, "/backup - POST - Configure backups")
	fmt.Fprintln(w, "/backup/{backupId} - GET - List backups")
	fmt.Fprintln(w, "/backup/{backupId} - DELETE - Delete backups")
	fmt.Fprintln(w, "/sshkey - POST - Post the SSH key")
	fmt.Fprintln(w, "/sshkey - GET - Get the SSH key")
}
