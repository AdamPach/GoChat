package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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
	defRooms["default\r\n"] = CreateRoom("default\r\n")

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
				go newClient.HandleIncomingMessages(s.deadConnections, s)
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
	input := bufio.NewReader(os.Stdin)

	for s.running {
		fmt.Print("> ")
		command, err := input.ReadString('\n')
		if err != nil {

		}

		if strings.Contains(strings.ToLower(command), "exit") {
			s.running = false
			break
		}
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

func (s *Server) AddClientToRoom(roomName string, client *Client) bool {
	wantedRoom := s.rooms[roomName]
	if wantedRoom == nil {
		return false
	}
	client.room = wantedRoom
	wantedRoom.Clients[client.name] = client

	return true
}
