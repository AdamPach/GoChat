package Server

import (
	"GoChat/Shared"
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Name       string
	Room       *Room
	Connection net.Conn
	BaseServer *Server
}

func CreateIncomeClient(conn net.Conn, ser *Server) Client {
	return Client{Connection: conn, BaseServer: ser}
}

func (c *Client) HandleIncomingMessages(deadChannel chan net.Conn, server *Server) {
	reader := bufio.NewReader(c.Connection)
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

	deadChannel <- c.Connection
}
func (c *Client) UseCommands(command string, server *Server) {
	cmd := strings.Split(command, " ")

	SearchedCommand := SearcheCommand(server.ClientCommands, Shared.RemoveSendingCharacters(cmd[0]))

	if SearchedCommand == nil {
		c.SendErrorMessage("You enter command which don't exist!\n")
		return
	}

	err := SearchedCommand.ExecuteCommand(c, cmd[1:])
	if err != nil {
		c.SendErrorMessage(fmt.Sprintf("%s\n", err.Error()))
		return
	}
}

func (c *Client) SendErrorMessage(message string) {
	_, err := c.Connection.Write([]byte(message))
	if err != nil {
		return
	}
}

func (c *Client) SendMessage(message string) {
	if c.Name == "" || c.Room == nil {
		c.SendErrorMessage("You haven't set Nickname or Room yet! Please set it\n")
		return
	}
	message = fmt.Sprintf("%s : %s", c.Name, message)

	c.Room.SendMessageToRoom(message, c)
}

func (c *Client) GetRoom() *Room {
	return c.Room
}
