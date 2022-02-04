package Static

import (
	"GoChat/ChatServer/Config/ConfigModels"
	"GoChat/ChatServer/Funcs/ClientCommands"
	"GoChat/ChatServer/Models/ServerModels"
	"net"
)

func CreateServer(listener net.Listener, config *ConfigModels.RunningConfig) *ServerModels.Server {
	defRooms := make(map[string]*ServerModels.Room)
	defRooms["default"] = CreateRoom("default")

	return &ServerModels.Server{
		Listener:          listener,
		ServerConfig:      config,
		Rooms:             defRooms,
		ClientCommands:    ClientCommands.InitCommands(),
		Running:           true,
		Connections:       make(map[net.Conn]*ServerModels.ChatClient),
		IncomeConnections: make(chan net.Conn),
		DeadConnections:   make(chan net.Conn),
	}
}
