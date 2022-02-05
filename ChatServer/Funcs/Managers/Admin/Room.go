package Admin

import (
	"GoChat/ChatServer/Funcs/Sender"
	"GoChat/ChatServer/Funcs/Static"
	"GoChat/ChatServer/Models/ServerModels"
	"GoChat/Shared"
	"bufio"
	"errors"
	"fmt"
)

func ManageRooms(s *ServerModels.Server, reader *bufio.Reader) error {

	fmt.Print("1 - Add Room\n2 - Delete Room\n3 - Show All Rooms\n")
	command, err := Shared.EnterCommand(*reader, "Chose operation: ")
	if err != nil {
		return err
	}

	if Shared.FormatCommand(command) == "add" {
		name, err := Shared.EnterCommand(*reader, "Enter name of new room: ")
		if err != nil {
			return err
		}
		s.Rooms[name] = Static.CreateRoom(name)
		SaveConfig(s)
		return nil
	} else if Shared.FormatCommand(command) == "delete" {
		PrintAllRooms(s)
		rn, err := Shared.EnterCommand(*reader, "Enter room name: ")
		if err != nil {
			return err
		}
		err = DeleteRoom(s, rn)
		if err != nil {
			return err
		}
		return nil
	} else if Shared.FormatCommand(command) == "show" {
		PrintAllRooms(s)
		return nil
	}

	return errors.New("You enter bad choice")
}

func PrintAllRooms(s *ServerModels.Server) {
	fmt.Println("All rooms: ")
	for r := range s.Rooms {
		fmt.Println(r)
	}
}

func DeleteRoom(s *ServerModels.Server, roomName string) error {
	deletedRoom := s.Rooms[roomName]

	if deletedRoom == nil {
		return errors.New("You enter bad room name!")
	}

	for _, client := range deletedRoom.Clients {
		Sender.SendInfo(client, "This room was deleted, please select another")
		client.Room = nil
	}
	delete(s.Rooms, deletedRoom.RoomName)
	deletedRoom = nil

	SaveConfig(s)

	return nil
}
