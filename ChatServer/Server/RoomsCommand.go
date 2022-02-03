package Server

import (
	"GoChat/Shared"
	"fmt"
	"strings"
)

type RoomsCommand struct {
	minArgs     int
	commandName string
}

func (r *RoomsCommand) CommandName() string {
	return r.commandName
}
func (r *RoomsCommand) ExecuteCommand(client *Client, args []string) error {
	roomsNames := make([]string, 0)
	for rn := range client.BaseServer.Rooms {
		roomsNames = append(roomsNames, rn)
	}
	client.SendErrorMessage(fmt.Sprintf("Rooms: %s\n", strings.Join(roomsNames, ", ")))

	return nil
}

func (r *RoomsCommand) CompareInput(input string) bool {
	if Shared.RemoveSendingCharacters(input) == r.commandName {
		return true
	}
	return false
}
func (r *RoomsCommand) MinArguments() int {
	return r.minArgs
}

func CreateRoomsCommand() *RoomsCommand {
	return &RoomsCommand{
		minArgs:     0,
		commandName: "/rooms",
	}
}
