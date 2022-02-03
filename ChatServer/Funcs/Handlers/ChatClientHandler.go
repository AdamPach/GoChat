package Handlers

import (
	"GoChat/ChatServer/Funcs/Sender"
	"GoChat/ChatServer/Models"
	"GoChat/ChatServer/Models/ServerModels"
	"bufio"
)

func HandleClientMessages(c *ServerModels.ChatClient, s *Models.Server) {
	reader := bufio.NewReader(c.Connection)
	for s.Running {
		mesg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if mesg[0] == '/' {

		} else {
			Sender.SendMessage(c, mesg)
		}
	}
	s.DeadConnections <- c.Connection
}
