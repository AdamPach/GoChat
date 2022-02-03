package Static

import (
	"GoChat/ChatServer/Models"
	"GoChat/ChatServer/Models/ServerModels"
	"net"
)

func CreateServer(listener net.Listener) *Models.Server {
	defRooms := make(map[string]*ServerModels.Room)
	defRooms["default"] = CreateRoom("default")

	return &Models.Server{
		Listener: listener,
		Rooms:    defRooms,
		//ClientCommands:    InitCommands(),
		Running:           true,
		Connections:       make(map[net.Conn]*ServerModels.ChatClient),
		IncomeConnections: make(chan net.Conn),
		DeadConnections:   make(chan net.Conn),
	}
}
