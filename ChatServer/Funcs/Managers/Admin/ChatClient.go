package Admin

import (
	"GoChat/ChatServer/Models/ServerModels"
	"fmt"
)

func PrintAllChatClients(s *ServerModels.Server) {
	for _, Client := range s.Connections {
		fmt.Println("Client: ", Client.Name)
	}
}

func GetChatClientByName(s *ServerModels.Server, name string) *ServerModels.ChatClient {
	for _, Client := range s.Connections {
		if Client.Name == name {
			return Client
		}
	}
	return nil
}
