package main

import (
	"GoChat/ChatServer/Funcs/Handlers"
	"GoChat/ChatServer/Funcs/Managers/Admin"
	"GoChat/ChatServer/Funcs/Static"
	"fmt"
	"net"
	"os"
)

func main() {
	Listener, err := net.Listen("tcp", ":7000")

	if err != nil {
		fmt.Println("You cant listen now!")
		os.Exit(1)
	}

	defer Listener.Close()

	server := Static.CreateServer(Listener)

	go Handlers.Listen(server)
	go Handlers.ManageConnections(server)
	Admin.ManageRunningServer(server)
}
