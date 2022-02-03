package main

import (
	"GoChat/Shared"
	"errors"
)

type RoomCommand struct {
	minArgs     int
	commandName string
}

func (r *RoomCommand) CommandName() string {
	return r.commandName
}
func (r *RoomCommand) ExecuteCommand(client *Client, args []string) error {
	if len(args) == 0 {
		return errors.New("This command needs args!")
	}
	if client.Name == "" {
		return errors.New("Firstly you must set nick with /nick command!")
	}
	if Shared.RemoveSendingCharacters(args[0]) == "" {
		return errors.New("Please enter room name")
	}
	res := client.BaseServer.AddClientToRoom(Shared.RemoveSendingCharacters(args[0]), client)
	if !res {
		return errors.New("Please send existing room name")
	}

	return nil
}
func (r *RoomCommand) CompareInput(input string) bool {
	if Shared.RemoveSendingCharacters(input) == r.commandName {
		return true
	}
	return false
}
func (r *RoomCommand) MinArguments() int {
	return r.minArgs
}

func CreateRoomCommand() *RoomCommand {
	return &RoomCommand{
		minArgs:     1,
		commandName: "/room",
	}
}
