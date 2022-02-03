package ClientCommands

import (
	"GoChat/ChatServer/Funcs/Sender"
	"GoChat/ChatServer/Models/ServerModels"
	"GoChat/Shared"
	"fmt"
	"strings"
)

func InitCommands() map[string]*ServerModels.ClientCommand {
	commands := make(map[string]*ServerModels.ClientCommand)
	newCmd := CreateNickCommand()
	commands[newCmd.CommandName] = newCmd

	newCmd = CreateRoomCommand()
	commands[newCmd.CommandName] = newCmd

	newCmd = CreateRoomsCommand()
	commands[newCmd.CommandName] = newCmd

	newCmd = CreateUsersInRoomCommand()
	commands[newCmd.CommandName] = newCmd

	return commands
}

func ExecuteCommand(s *ServerModels.Server, c *ServerModels.ChatClient, mesg string) {
	cmd := strings.Split(mesg, " ")

	SearchedCommand := s.ClientCommands[Shared.RemoveSendingCharacters(cmd[0])]

	if SearchedCommand == nil {
		Sender.SendInfo(c, "You enter command which don't exist!")
		return
	}

	err := SearchedCommand.ExecFunc(c, s, cmd[1:])

	if err != nil {
		Sender.SendInfo(c, fmt.Sprintf("%s", err.Error()))
		return
	}
}
