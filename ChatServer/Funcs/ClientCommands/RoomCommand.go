package ClientCommands

import (
	"GoChat/ChatServer/Funcs/Managers"
	"GoChat/ChatServer/Models/ServerModels"
	"GoChat/Shared"
	"errors"
)

func CreateRoomCommand() *ServerModels.ClientCommand {
	return &ServerModels.ClientCommand{
		CommandName: "/room",
		MinArgs:     1,
		ExecFunc:    ExecuteRoomCommand,
	}
}

func ExecuteRoomCommand(client *ServerModels.ChatClient, server *ServerModels.Server, args []string) error {
	if len(args) == 0 {
		return errors.New("This command needs args!")
	}
	if client.Name == "" {
		return errors.New("Firstly you must set nick with /nick command!")
	}
	if Shared.RemoveSendingCharacters(args[0]) == "" {
		return errors.New("Please enter room name")
	}
	err := Managers.AddClientToRoom(server, client, Shared.RemoveSendingCharacters(args[0]))
	if err != nil {
		return errors.New("Please send existing room name")
	}

	return nil
}
