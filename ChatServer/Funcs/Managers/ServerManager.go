package Managers

import (
	"GoChat/ChatServer/Models/ServerModels"
	"net"
)

func DeleteDeadConnection(s *ServerModels.Server, DeadConn net.Conn) {
	for item := range s.Connections {
		if item == DeadConn {
			deadClient := s.Connections[DeadConn]
			s.RoomLocker.RLock()
			if room := deadClient.Room; room != nil {
				DeleteUserFromRoom(room, deadClient)
			}
			s.RoomLocker.RUnlock()
			delete(s.Connections, DeadConn)
			break
		}
	}
}
