package Admin

import (
	"GoChat/ChatServer/Models/ServerModels"
	"bufio"
)

func CreateAdminRoomsCommand() *ServerModels.ServerCommand {
	return &ServerModels.ServerCommand{
		Name:        "rooms",
		ExecCommand: ExecAdminRooms,
	}
}

func ExecAdminRooms(s *ServerModels.Server, reader *bufio.Reader) error {
	err := ManageRooms(s, reader)
	if err != nil {
		return err
	}
	return nil
}
