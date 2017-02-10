package main

/*

Json Structure:

{
    name: "backupname",
    target: "<ip of the target machine or FQDN>",
    backupsrc: "<folder to be backup>",
    schedule: "<hourly/daily/weekly>",
    password: "<restic password>"
}

*/

type BackupJob struct {
	BackupName string `json:"BackupName"`
	TargetNode string `json:"TargetNode"`
	BackupSrc  string `json:"BackupSrc"`
	Schedule   string `json:"Schedule"`
	Password   string `json:"Password"`
}

type Backups []BackupJob

type JobList struct {
	BackupList []string
}
