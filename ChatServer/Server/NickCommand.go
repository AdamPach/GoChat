package Server

import (
	"GoChat/Shared"
	"errors"
)

type NickCommand struct {
	minArgs     int
	commandName string
}

func (n *NickCommand) CommandName() string {
	return n.commandName
}
func (n *NickCommand) ExecuteCommand(client *Client, args []string) error {
	if len(args) == 0 {
		return errors.New("This command needs args!")
	}
	if Shared.RemoveSendingCharacters(args[0]) == "" {
		return errors.New("You don't enter your name!")
	}
	client.Name = Shared.RemoveSendingCharacters(args[0])
	return nil
}
func (n *NickCommand) CompareInput(input string) bool {
	if Shared.RemoveSendingCharacters(input) == n.commandName {
		return true
	}
	return false
}
func (n *NickCommand) MinArguments() int {
	return n.minArgs
}

func CreateNickCommand() *NickCommand {
	return &NickCommand{
		minArgs:     1,
		commandName: "/nick",
	}
}
