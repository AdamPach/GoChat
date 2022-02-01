package main

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	name       string
	room       *Room
	connection net.Conn
}

func CreateIncomeClient(conn net.Conn) Client {
	return Client{connection: conn}
}

func (c *Client) HandleIncomingMessages(deadChannel chan net.Conn) {
	reader := bufio.NewReader(c.connection)
	for {
		mesg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		//Handle requset
		if ItsCommand(mesg) {

		} else {
			c.SendMessage(mesg)
		}
	}

	deadChannel <- c.connection
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
