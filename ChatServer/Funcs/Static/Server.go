package Static

import (
	"GoChat/ChatServer/Config/ConfigModels"
	"GoChat/ChatServer/Funcs/ClientCommands"
	"GoChat/ChatServer/Models/ServerModels"
	"net"
	"sync"
)

func CreateServer(listener net.Listener, config *ConfigModels.RunningConfig) *ServerModels.Server {

	return &ServerModels.Server{
		Listener:          listener,
		ServerConfig:      config,
		Rooms:             ReadRoomsFromConfig(config),
		ClientCommands:    ClientCommands.InitCommands(),
		Running:           true,
		Connections:       make(map[net.Conn]*ServerModels.ChatClient),
		IncomeConnections: make(chan net.Conn),
		DeadConnections:   make(chan net.Conn),
		RoomLocker:        sync.RWMutex{},
	}
}
