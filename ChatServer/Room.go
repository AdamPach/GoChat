package main

type Room struct {
	roomName string
	Clients  map[string]*Client
}

func CreateRoom(roomName string) *Room {
	return &Room{
		roomName: roomName,
		Clients:  make(map[string]*Client),
	}
}

func (r *Room) DeleteClient(client *Client) {
	delete(r.Clients, client.name)
}
