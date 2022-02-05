package Admin

import (
	"GoChat/ChatServer/Models/ServerModels"
	"bufio"
	"os"
)

func CreateExitCommand() *ServerModels.ServerCommand {
	return &ServerModels.ServerCommand{
		Name:        "exit",
		ExecCommand: ExitCommandExec,
	}
}

func ExitCommandExec(s *ServerModels.Server, reader *bufio.Reader) error {
	s.Running = false
	CloseAllConections(s)
	os.Exit(0)
	return nil
}
