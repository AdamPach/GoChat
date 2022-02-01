package main

import (
	"bufio"
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
	}

	deadChannel <- c.connection
}

func (c *Client) GetRoom() *Room {
	return c.room
}
