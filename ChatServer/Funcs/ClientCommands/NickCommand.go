package ClientCommands

import (
	"GoChat/ChatServer/Models/ServerModels"
	"GoChat/Shared"
	"errors"
)

func CreateNickCommand() *ServerModels.ClientCommand {
	return &ServerModels.ClientCommand{
		CommandName: "/nick",
		MinArgs:     1,
		ExecFunc:    ExectuceNickCommand,
	}
}

func ExectuceNickCommand(c *ServerModels.ChatClient, server *ServerModels.Server, args []string) error {
	if len(args) == 0 {
		return errors.New("This command needs args!")
	}
	if Shared.RemoveSendingCharacters(args[0]) == "" {
		return errors.New("You don't enter your name!")
	}
	c.Name = Shared.RemoveSendingCharacters(args[0])
	return nil
}
