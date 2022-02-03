package Models

import (
	"GoChat/ChatServer/Models/CommandModels"
	"GoChat/ChatServer/Models/ServerModels"
	"net"
)

type Server struct {
	listener          net.Listener
	incomeConnections chan net.Conn
	deadConnections   chan net.Conn
	running           bool
	Connections       map[net.Conn]*ServerModels.ChatClient
	Rooms             map[string]*ServerModels.Room
	ClientCommands    map[string]*CommandModels.BaseCommand
}
