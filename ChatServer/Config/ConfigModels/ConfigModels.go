package ConfigModels

type Config struct {
	Port  int
	Rooms []Room
}

type Room struct {
	Name string
}
