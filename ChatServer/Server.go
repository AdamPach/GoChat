package main

import (
	"GoChat/Shared"
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
			continue
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
			s.running = false
			s.CloseAllConections()
			os.Exit(0)
		}

		if strings.ToLower(Shared.RemoveSendingCharacters(command)) == "exit" {
			s.running = false
			s.CloseAllConections()
			os.Exit(0)
		} else if strings.ToLower(Shared.RemoveSendingCharacters(command)) == "rooms" {
			err = s.ManageRooms(input)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (s *Server) CloseAllConections() {
	for c := range s.connections {
		c.Close()
	}
}

func (s *Server) DeleteDeadConnection(deadConn net.Conn) {
	defer deadConn.Close()
	for item := range s.connections {
		if item == deadConn {
			deadClient := s.connections[item]
			if room := deadClient.GetRoom(); room != nil {
				room.DeleteClient(deadClient)
			}
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
	if client.room != nil {
		client.room.RemoveClientFromRoom(client)
	}
	client.room = wantedRoom
	wantedRoom.Clients[client.name] = client

	return true
}
