package main

import (
	"GoChat/Shared"
	"errors"
	"fmt"
	"strings"
)

type UserInRoom struct {
	minArgs     int
	commandName string
}

func (u *UserInRoom) CommandName() string {
	return u.commandName
}
func (u *UserInRoom) ExecuteCommand(client *Client, args []string) error {
	if client.Room == nil {
		return errors.New("You aren't in any room!")
	}

	names := make([]string, 0)

	for _, user := range client.Room.Clients {
		if user == client {
			names = append(names, "You")
			continue
		}
		names = append(names, user.Name)
	}

	client.Connection.Write([]byte(fmt.Sprintf("Users: %s\n", strings.Join(names, ", "))))

	return nil
}

func (u *UserInRoom) CompareInput(input string) bool {
	if Shared.RemoveSendingCharacters(input) == u.commandName {
		return true
	}
	return false
}
func (u *UserInRoom) MinArguments() int {
	return u.minArgs
}

func CreateUserInRoomCommand() *UserInRoom {
	return &UserInRoom{
		minArgs:     0,
		commandName: "/users",
	}
}
