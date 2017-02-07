package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/backup",
		ListBackups,
	},
	Route{
		"ListBackup",
		"POST",
		"/backup",
		ConfigBackup,
	},
	Route{
		"ListBackupById",
		"GET",
		"/backup/{backupId}",
		ListBackupById,
	},
	Route{
		"DeleteBackup",
		"DELETE",
		"/backup/{backupId}",
		DeleteBackupById,
	},
	Route{
		"ConfigSSHKey",
		"POST",
		"/sshkey",
		ConfigSSHKey,
	},
	Route{
		"GetSSHKey",
		"GET",
		"/sshkey",
		GetSSHKey,
	},
}
