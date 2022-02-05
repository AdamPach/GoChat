package Admin

import (
	"GoChat/ChatServer/Funcs/Sender"
	"GoChat/ChatServer/Models/ServerModels"
	"GoChat/Shared"
	"bufio"
	"errors"
	"fmt"
)

func CreateKickClientCommand() *ServerModels.ServerCommand {
	return &ServerModels.ServerCommand{
		Name:        "kick",
		ExecCommand: ExecKickCommand,
	}
}

func ExecKickCommand(s *ServerModels.Server, reader *bufio.Reader) error {
	PrintAllChatClients(s)
	name, err := Shared.EnterCommand(*reader, "Enter client name(Case sensitive): ")

	if err != nil {
		return err
	}

	KickedClient := GetChatClientByName(s, name)

	if KickedClient == nil {
		return errors.New("This client don't exist!")
	}

	Sender.SendInfo(KickedClient, "You have been kicked from this server!")

	s.DeadConnections <- KickedClient.Connection

	fmt.Printf("Client %s now been kicked!\n", name)

	return nil
}
