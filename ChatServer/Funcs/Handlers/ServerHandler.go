package Handlers

import (
	"GoChat/ChatServer/Funcs/Managers"
	"GoChat/ChatServer/Funcs/Static"
	"GoChat/ChatServer/Models/ServerModels"
	"log"
)

func Listen(s *ServerModels.Server) {
	for s.Running {
		newConn, err := s.Listener.Accept()
		if err != nil {
			log.Println("You cant accept this connection!")
			continue
		}

		s.IncomeConnections <- newConn
	}
}

func ManageConnections(s *ServerModels.Server) {
	for s.Running {
		select {
		case IncomeConn := <-s.IncomeConnections:
			{
				NewClient := Static.CreateClient(IncomeConn)
				s.Connections[IncomeConn] = NewClient
				go HandleClientMessages(NewClient, s)
			}
		case DeadConn := <-s.DeadConnections:
			{
				Managers.DeleteDeadConnection(s, DeadConn)
				DeadConn.Close()
			}

		}
	}
}
