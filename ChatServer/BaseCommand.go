package main

type BaseCommand interface {
	CommandName() string
	ExecuteCommand(client *Client, args []string) error
	CompareInput(input string) bool
	MinArguments() int
}

func InitCommands() []BaseCommand {
	bc := make([]BaseCommand, 0)
	bc = append(bc, CreateUserInRoomCommand())
	bc = append(bc, CreateNickCommand())
	bc = append(bc, CreateRoomsCommand())
	bc = append(bc, CreateRoomCommand())
	return bc
}

func SearcheCommand(commands []BaseCommand, commandName string) BaseCommand {
	for i := 0; i < len(commands); i++ {
		if commands[i].CommandName() == commandName {
			return commands[i]
		}
	}

	return nil
}
