package Static

import (
	"GoChat/ChatServer/Config/ConfigModels"
	"GoChat/ChatServer/Models/ServerModels"
)

func CreateRoom(RoomName string) *ServerModels.Room {
	return &ServerModels.Room{
		RoomName: RoomName,
		Clients:  make(map[string]*ServerModels.ChatClient),
	}
}

func ReadRoomsFromConfig(r *ConfigModels.RunningConfig) map[string]*ServerModels.Room {
	ReaderRooms := make(map[string]*ServerModels.Room)

	for _, Room := range r.Configuration.Rooms {
		ReaderRooms[Room.Name] = CreateRoom(Room.Name)
	}

	return ReaderRooms
}
