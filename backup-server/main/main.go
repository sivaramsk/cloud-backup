package main

import (
	"flag"
	"log"
)

func main() {
	log.Println("backup-server")
	test := flag.String("test", "", "Run automated test.")
	address := flag.String("address", "", "Bind address for the server.")
	port := flag.String("port", "", "Bind port for the server.")

	flag.Parse()
	log.Println("Address: %s, Port: %s", *address, *port)

}
