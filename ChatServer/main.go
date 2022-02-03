package main

import (
	"GoChat/ChatServer/Server"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":7000")

	if err != nil {
		log.Fatal("You can't listen now!")
		os.Exit(1)
	}

	defer listener.Close()

	server := Server.CreateServer(listener)
	go server.Listen()
	go server.ManagedConnections()
	server.ManagingServer()
}
