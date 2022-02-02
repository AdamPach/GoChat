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

		return nil
	}

	return errors.New("You enter bad choice")
}
