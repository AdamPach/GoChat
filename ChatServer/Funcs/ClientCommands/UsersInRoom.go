package ClientCommands

import (
	"GoChat/ChatServer/Models/ServerModels"
	"errors"
	"fmt"
	"strings"
)

func CreateUsersInRoomCommand() *ServerModels.ClientCommand {
	return &ServerModels.ClientCommand{
		CommandName: "/users",
		MinArgs:     0,
		ExecFunc:    ExecuteUserInRoom,
	}
}

func ExecuteUserInRoom(client *ServerModels.ChatClient, server *ServerModels.Server, args []string) error {
	if client.Room == nil {
		return errors.New("You aren't in any room!")
	}

	names := make([]string, 0)

	for _, user := range client.Room.Clients {
		if user == client {
			names = append(names, "You")
			continue
		}
		names = append(names, user.Name)
	}

	client.Connection.Write([]byte(fmt.Sprintf("Users: %s\n", strings.Join(names, ", "))))

	return nil
}
