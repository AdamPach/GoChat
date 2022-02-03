package Static

import (
	"GoChat/ChatServer/Models/ServerModels"
	"net"
)

func CreateClient(conn net.Conn) *ServerModels.ChatClient {
	return &ServerModels.ChatClient{
		Name:       "",
		Room:       nil,
		Connection: conn,
	}
}
