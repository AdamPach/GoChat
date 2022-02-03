package CommandModels

import "GoChat/ChatServer/Models/ServerModels"

type BaseCommand interface {
	CommandName() string
	ExecuteCommand(client *ServerModels.ChatClient, args []string) error
}
