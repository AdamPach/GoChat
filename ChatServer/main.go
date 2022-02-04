package main

import (
	"GoChat/ChatServer/Config"
	"GoChat/ChatServer/Funcs/Handlers"
	"GoChat/ChatServer/Funcs/Managers/Admin"
	"GoChat/ChatServer/Funcs/Static"
	"fmt"
	"net"
	"os"
)

func main() {
	config, err := Config.InitConfig("config.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Listener, err := net.Listen("tcp", Config.GetPort(config))

	if err != nil {
		fmt.Println("You cant listen now!")
		os.Exit(1)
	}

	defer Listener.Close()

	server := Static.CreateServer(Listener, config)

	go Handlers.Listen(server)
	go Handlers.ManageConnections(server)
	Admin.ManageRunningServer(server)
}
