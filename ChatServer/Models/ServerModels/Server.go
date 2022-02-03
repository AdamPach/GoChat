package ServerModels

import (
	"net"
)

type Server struct {
	Listener          net.Listener
	IncomeConnections chan net.Conn
	DeadConnections   chan net.Conn
	Running           bool
	Connections       map[net.Conn]*ChatClient
	Rooms             map[string]*Room
	ClientCommands    map[string]*ClientCommand
}
