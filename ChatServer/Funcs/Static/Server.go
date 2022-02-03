package Static

import (
	"GoChat/ChatServer/Funcs/ClientCommands"
	"GoChat/ChatServer/Models/ServerModels"
	"net"
)

func CreateServer(listener net.Listener) *ServerModels.Server {
	defRooms := make(map[string]*ServerModels.Room)
	defRooms["default"] = CreateRoom("default")

	return &ServerModels.Server{
		Listener:          listener,
		Rooms:             defRooms,
		ClientCommands:    ClientCommands.InitCommands(),
		Running:           true,
		Connections:       make(map[net.Conn]*ServerModels.ChatClient),
		IncomeConnections: make(chan net.Conn),
		DeadConnections:   make(chan net.Conn),
	}
}
