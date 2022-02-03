package ClientCommands

import (
	"GoChat/ChatServer/Models/ServerModels"
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
