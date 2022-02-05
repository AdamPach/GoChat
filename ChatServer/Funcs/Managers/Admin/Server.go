package Admin

import (
	"GoChat/ChatServer/Config"
	"GoChat/ChatServer/Config/ConfigModels"
	"GoChat/ChatServer/Models/ServerModels"
	"GoChat/Shared"
	"bufio"
	"fmt"
	"os"
)

func ManageRunningServer(s *ServerModels.Server) {
	input := bufio.NewReader(os.Stdin)

	AdminCommadns := InitServerCommands()

	for s.Running {
		fmt.Print("> ")
		command, err := input.ReadString('\n')
		if err != nil {
			s.Running = false
			CloseAllConections(s)
			os.Exit(1)
		}

		AdminCommand := AdminCommadns[Shared.FormatCommand(command)]

		if AdminCommand == nil {
			fmt.Println("Bad command name")
			continue
		}

		err = AdminCommand.ExecCommand(s, input)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

	}
}

func SaveConfig(s *ServerModels.Server) {
	NewRooms := make([]ConfigModels.Room, 0)
	for rName := range s.Rooms {
		NewRooms = append(NewRooms, ConfigModels.Room{Name: rName})
	}
	s.ServerConfig.Configuration.Rooms = NewRooms
	err := Config.WriteConfig(s.ServerConfig)

	if err != nil {
		fmt.Println("[ERROR]: Cant save config!")
	}
}

func CloseAllConections(s *ServerModels.Server) {
	for c := range s.Connections {
		c.Close()
	}
}

func InitServerCommands() map[string]*ServerModels.ServerCommand {
	AdminCommands := make(map[string]*ServerModels.ServerCommand)

	NewCmd := CreateExitCommand()
	AdminCommands[NewCmd.Name] = NewCmd

	NewCmd = CreateAdminRoomsCommand()
	AdminCommands[NewCmd.Name] = NewCmd

	NewCmd = CreateKickClientCommand()
	AdminCommands[NewCmd.Name] = NewCmd

	return AdminCommands
}
