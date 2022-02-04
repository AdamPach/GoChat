package ConfigModels

type Config struct {
	Port  string
	Rooms []Room
}

type Room struct {
	Name string
}

type RunningConfig struct {
	ConfigFilePath string
	Configuration  *Config
}
