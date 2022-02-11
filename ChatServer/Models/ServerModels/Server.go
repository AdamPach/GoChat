package ServerModels

import (
	"GoChat/ChatServer/Config/ConfigModels"
	"net"
	"sync"
)

type Server struct {
	Listener          net.Listener
	ServerConfig      *ConfigModels.RunningConfig
	IncomeConnections chan net.Conn
	DeadConnections   chan net.Conn
	Running           bool
	Connections       map[net.Conn]*ChatClient
	Rooms             map[string]*Room
	ClientCommands    map[string]*ClientCommand
	RoomLocker        sync.RWMutex
}
