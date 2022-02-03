package Models

import (
	"GoChat/ChatServer/Models/CommandModels"
	"GoChat/ChatServer/Models/ServerModels"
	"net"
)

type Server struct {
	Listener          net.Listener
	IncomeConnections chan net.Conn
	DeadConnections   chan net.Conn
	Running           bool
	Connections       map[net.Conn]*ServerModels.ChatClient
	Rooms             map[string]*ServerModels.Room
	ClientCommands    map[string]*CommandModels.BaseCommand
}
