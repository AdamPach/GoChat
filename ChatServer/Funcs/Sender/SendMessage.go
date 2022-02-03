package Sender

import "GoChat/ChatServer/Models/ServerModels"

func SendMessage(c *ServerModels.ChatClient, message string) {
	if c.Name == "" || c.Room == nil {
		SendInfo(c, "You haven't set Nickname or Room yet! Please set it")
		return
	}
	BroadcastMessageToRoom(c, message)
}

func BroadcastMessageToRoom(c *ServerModels.ChatClient, message string) {
	for _, client := range c.Room.Clients {
		if client != c {
			client.Connection.Write([]byte(message))
		}
	}
}
