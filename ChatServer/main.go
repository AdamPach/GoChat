package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":7000")

	if err != nil {
		log.Fatal("You can't listen now!")
		return
	}

	defer listener.Close()

	server := CreateServer(listener)
	go server.Listen()
	go server.ManagedConnections()
	server.ManagingServer()
}
