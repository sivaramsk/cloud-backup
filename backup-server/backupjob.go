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
	BackupName string `json:"name"`
	TargetNode string `json:"target"`
	BackupSrc  string `json:"backupsrc"`
	Schedule   string `json:"schedule"`
	Password   string `json:"password"`
}

type Backups []BackupJob
