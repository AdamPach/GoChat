package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	name       string
	room       *Room
	connection net.Conn
}

func CreateIncomeClient(conn net.Conn) Client {
	return Client{connection: conn}
}

func (c *Client) HandleIncomingMessages(deadChannel chan net.Conn, server *Server) {
	reader := bufio.NewReader(c.connection)
	for {
		mesg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if ItsCommand(mesg) {
			c.UseCommands(mesg, server)
		} else {
			c.SendMessage(mesg)
		}
	}

	deadChannel <- c.connection
}

func (c *Client) UseCommands(command string, server *Server) {
	cmd := strings.Split(command, " ")
	if cmd[0] == "/rooms" {
		roomsNames := make([]string, 0)
		for rn := range server.rooms {
			roomsNames = append(roomsNames, rn)
		}
		c.SendErrorMessage(fmt.Sprintf("Rooms: %s", strings.Join(roomsNames, ", ")))
		return
	}

	if len(cmd) < 2 {
		c.SendErrorMessage("Please enter command with arguments\n")
		return
	}

	if cmd[0] == "/room" {
		if c.name == "" {
			c.SendErrorMessage("Firstly you must set nick with /nick command!\n")
			return
		}
		res := server.AddClientToRoom(cmd[1], c)
		if !res {
			c.SendErrorMessage("Please send existing room name\n")
			return
		}
	} else if cmd[0] == "/nick" {
		c.name = cmd[1]
		return
	} else {
		c.SendMessage("This command don't exist!")
	}
}

func (c *Client) SendErrorMessage(message string) {
	_, err := c.connection.Write([]byte(message))
	if err != nil {
		return
	}
}

func (c *Client) SendMessage(message string) {
	if c.name == "" || c.room == nil {
		_, err := c.connection.Write([]byte("You haven't set Nickname or Room yet! Please set it"))
		if err != nil {
			return
		}
		return
	}
	message = fmt.Sprintf("%s\t: %s\n", c.name, message)

	c.room.SendMessageToRoom(message, c)
}

func (c *Client) GetRoom() *Room {
	return c.room
}
