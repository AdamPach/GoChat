package main

import (
	"log"
	"net"
)

type Server struct {
	listener          net.Listener
	rooms             map[string]*Room
	running           bool
	connections       []Client
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
		connections:       make([]Client, 0),
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
				s.connections = append(s.connections, newClient)

			}

		}
	}
}
