package Handlers

import (
	"GoChat/ChatServer/Funcs/ClientCommands"
	"GoChat/ChatServer/Funcs/Sender"
	"GoChat/ChatServer/Models/ServerModels"
	"bufio"
)

func HandleClientMessages(c *ServerModels.ChatClient, s *ServerModels.Server) {
	reader := bufio.NewReader(c.Connection)
	for s.Running {
		mesg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if mesg[0] == '/' {
			ClientCommands.ExecuteCommand(s, c, mesg)
		} else {
			Sender.SendMessage(c, mesg)
		}
	}
	s.DeadConnections <- c.Connection
}
