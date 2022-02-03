package ServerModels

type Room struct {
	RoomName string
	Clients  map[string]*ChatClient
}
