package ClientCommands

import (
	"GoChat/ChatServer/Funcs/Sender"
	"GoChat/ChatServer/Models/ServerModels"
	"fmt"
	"strings"
)

func CreateRoomsCommand() *ServerModels.ClientCommand {
	return &ServerModels.ClientCommand{
		CommandName: "/rooms",
		MinArgs:     0,
		ExecFunc:    ExecruteRoomsCommand,
	}
}

func ExecruteRoomsCommand(client *ServerModels.ChatClient, server *ServerModels.Server, args []string) error {
	roomsName := make([]string, 0)

	for rn := range server.Rooms {
		roomsName = append(roomsName, rn)
	}

	Sender.SendInfo(client, fmt.Sprintf("Rooms: %s", strings.Join(roomsName, ", ")))

	return nil
}
