package Sender

import (
	"GoChat/ChatServer/Models/ServerModels"
	"fmt"
)

func SendInfo(c *ServerModels.ChatClient, message string) {
	c.Connection.Write([]byte(fmt.Sprintf("%s\n", message)))
}
