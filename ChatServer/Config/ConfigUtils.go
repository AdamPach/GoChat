package Config

import "GoChat/ChatServer/Config/ConfigModels"

func GetPort(r *ConfigModels.RunningConfig) string {
	return r.Configuration.Port
}
