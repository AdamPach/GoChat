package ServerModels

type Room struct {
	roomName string
	Clients  map[string]*ChatClient
}
