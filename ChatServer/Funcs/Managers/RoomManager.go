package Managers

import (
	"GoChat/ChatServer/Models/ServerModels"
	"errors"
)

func DeleteUserFromRoom(r *ServerModels.Room, c *ServerModels.ChatClient) {
	delete(r.Clients, c.Name)
}

func AddClientToRoom(s *ServerModels.Server, c *ServerModels.ChatClient, roomName string) error {
	wantedRoom := s.Rooms[roomName]
	if wantedRoom == nil {
		return errors.New("You enter invalid room name")
	}
	if c.Room != nil {
		DeleteUserFromRoom(wantedRoom, c)
	}

	c.Room = wantedRoom
	wantedRoom.Clients[c.Name] = c
	return nil
}