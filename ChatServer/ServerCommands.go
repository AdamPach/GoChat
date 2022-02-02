package main

import (
	"GoChat/Shared"
	"bufio"
	"errors"
	"fmt"
)

func (s *Server) ManageRooms(reader *bufio.Reader) error {

	fmt.Print("1 - Add Room\n2 - Delete Room\nChose operation: ")
	command, err := reader.ReadString('\n')
	Shared.LogError(&err)

	if Shared.FormatCommand(command) == "add" {
		name, err := reader.ReadString('\n')
		Shared.LogError(&err)
		name = Shared.RemoveSendingCharacters(name)
		s.rooms[name] = CreateRoom(name)
		return nil
	} else if Shared.FormatCommand(command) == "delete" {
		s.PrintAllRooms()
		rn, err := reader.ReadString('\n')
		Shared.LogError(&err)
		err = s.DeleteRoom(Shared.RemoveSendingCharacters(rn))
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("You enter bad choice")
}

func (s *Server) PrintAllRooms() {
	fmt.Println("All rooms: ")
	for r := range s.rooms {
		fmt.Println(r)
	}
}

func (s *Server) DeleteRoom(roomName string) error {
	deletedRoom := s.rooms[roomName]

	if deletedRoom == nil {
		return errors.New("You enter bad room name!")
	}

	for _, client := range deletedRoom.Clients {
		client.room = nil
	}
	delete(s.rooms, deletedRoom.roomName)
	deletedRoom = nil

	return nil
}
