package main

import (
	"log"
	internal "net-cat/internal/server"
	"os"
)

func main() {
	port := internal.DefaultPort
	args := os.Args[1:]
	if len(args) == 1 {
		port = args[0]
	} else if len(args) > 1 {
		log.Println(internal.Usage)
		os.Exit(0)
	}

	server := internal.RunServer("tcp", ":"+port)
	log.Println(internal.PortLocalhost + port)
	server.Start()
}
