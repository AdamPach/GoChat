package main

import (
	"log"
	"net"
)

type Server struct {
	listener          net.Listener
	rooms             map[string]*Room
	running           bool
	connections       map[net.Conn]*Client
	incomeConnections chan net.Conn
	deadConnections   chan net.Conn
}

func CreateServer(listener net.Listener) *Server {
	defRooms := make(map[string]*Room)
	defRooms["default"] = CreateRoom("default")

	return &Server{
		listener:          listener,
		rooms:             defRooms,
		running:           true,
		connections:       make(map[net.Conn]*Client),
		incomeConnections: make(chan net.Conn),
		deadConnections:   make(chan net.Conn),
	}
}

func (s *Server) Listen() {
	for s.running {
		conn, err := s.listener.Accept()

		if err != nil {
			log.Println("You cant accept this connection!")
		}

		s.incomeConnections <- conn
	}
}

func (s *Server) ManagedConnections() {
	for s.running {
		select {
		case conn := <-s.incomeConnections:
			{
				newClient := CreateIncomeClient(conn)
				s.connections[conn] = &newClient
				go newClient.HandleIncomingMessages(s.deadConnections)
			}
		case deadCon := <-s.deadConnections:
			{
				s.DeleteDeadConnection(deadCon)
				deadCon.Close()
			}

		}
	}
}

func (s *Server) ManagingServer() {
	for s.running {

	}
}

func (s *Server) DeleteDeadConnection(deadConn net.Conn) {
	for item := range s.connections {
		if item == deadConn {
			deadClient := s.connections[item]
			deadClient.GetRoom().DeleteClient(deadClient)
			delete(s.connections, deadConn)
			break
		}
	}
}
