package Managers

import (
	"GoChat/ChatServer/Models/ServerModels"
	"net"
)

func DeleteDeadConnection(s *ServerModels.Server, DeadConn net.Conn) {
	for item := range s.Connections {
		if item == DeadConn {
			deadClient := s.Connections[DeadConn]
			if room := deadClient.Room; room != nil {
				DeleteUserFromRoom(room, deadClient)
			}
			delete(s.Connections, DeadConn)
			break
		}
	}
}
