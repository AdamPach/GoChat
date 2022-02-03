package Static

import "GoChat/ChatServer/Models/ServerModels"

func CreateRoom(RoomName string) *ServerModels.Room {
	return &ServerModels.Room{
		RoomName: RoomName,
		Clients:  make(map[string]*ServerModels.ChatClient),
	}
}
