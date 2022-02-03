package Admin

import (
	"GoChat/ChatServer/Models/ServerModels"
	"GoChat/Shared"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ManageRunningServer(s *ServerModels.Server) {
	input := bufio.NewReader(os.Stdin)

	for s.Running {
		fmt.Print("> ")
		command, err := input.ReadString('\n')
		if err != nil {
			s.Running = false
			CloseAllConections(s)
			os.Exit(1)
		}

		if strings.ToLower(Shared.RemoveSendingCharacters(command)) == "exit" {
			s.Running = false
			CloseAllConections(s)
			os.Exit(0)
		} else if strings.ToLower(Shared.RemoveSendingCharacters(command)) == "rooms" {
			err = ManageRooms(s, input)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func CloseAllConections(s *ServerModels.Server) {
	for c := range s.Connections {
		c.Close()
	}
}
