package main

import (
	"GoChat/Shared"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	connection    net.Conn
	runningClient bool
}

func InitConnection() *Client {
	input := bufio.NewReader(os.Stdin)

	ipAddr := Shared.RemoveSendingCharacters(GetIPAddres(input))
	port := Shared.RemoveSendingCharacters(GetPor(input))

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ipAddr, port))

	if err != nil {
		fmt.Println("You cant connect to this server!")
		os.Exit(1)
	}

	return &Client{
		connection:    conn,
		runningClient: true,
	}
}

func (c *Client) ListenForMessages() {
	incomingMessagesReader := bufio.NewReader(c.connection)
	for c.runningClient {
		mesg, err := incomingMessagesReader.ReadString('\n')
		if err != nil {
			fmt.Println("Cant read messages")
			c.runningClient = false
			break
		}
		fmt.Println(mesg)
	}
}

func (c *Client) SendMessages() {
	input := bufio.NewReader(os.Stdin)
	for c.runningClient {
		fmt.Print("> ")
		mesg, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Cant send messages")
			c.runningClient = false
			break
		}

		_, err = c.connection.Write([]byte(mesg))

		if err != nil {
			fmt.Println("Cant send messages")
			c.runningClient = false
			break
		}
	}
}

func GetIPAddres(input *bufio.Reader) string {

	badAddr := true
	var addr string = ""

	for badAddr {
		fmt.Print("Enter valid ip address: ")
		add, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("You cant input ip address now")
			os.Exit(1)
		}

		if len(strings.Split(add, ".")) != 4 {
			fmt.Println("This is not a ip address! ")
		} else {
			addr = add
			badAddr = false
		}

	}
	return addr
}

func GetPor(input *bufio.Reader) string {
	fmt.Print("Enter valid port number: ")
	add, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("You cant input ip address now")
		os.Exit(1)
	}

	return add
}
