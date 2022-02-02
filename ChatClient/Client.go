package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	connection *net.Conn
}

func InitConnection() *Client {
	input := bufio.NewReader(os.Stdin)

	ipAddr := GetIPAddres(input)
	port := GetPor(input)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s,%s", ipAddr, port))

	if err != nil {
		fmt.Println("You cant connect to this server!")
		os.Exit(1)
	}

	return &Client{
		connection: &conn,
	}
}

func (c *Client) ListenForMessages() {

}

func (c *Client) SendMessages() {

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
