package Config

import (
	"GoChat/ChatServer/Config/ConfigModels"
	"GoChat/Shared"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func InitConfig(configPath string) (*ConfigModels.RunningConfig, error) {
	if !CheckIfConfigExist(configPath) {
		return CreateConfigFile(configPath)
	}

	config := ConfigModels.RunningConfig{ConfigFilePath: configPath}

	data, err := ioutil.ReadFile(config.ConfigFilePath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &config.Configuration)

	if err != nil {
		return nil, err
	}

	return &config, err
}

func CheckIfConfigExist(configPath string) bool {
	_, err := os.Stat(configPath)

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}

	return true
}

func CreateConfigFile(configPath string) (*ConfigModels.RunningConfig, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("You have to create config!")
	fmt.Print("Enter port: ")
	port, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	config := ConfigModels.RunningConfig{
		ConfigFilePath: configPath,
		Configuration: &ConfigModels.Config{
			Port: fmt.Sprintf(":%s", Shared.RemoveSendingCharacters(port)),
			Rooms: []ConfigModels.Room{
				{Name: "default"},
			},
		},
	}

	f, err := os.Create(config.ConfigFilePath)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := json.Marshal(config.Configuration)

	if err != nil {
		return nil, err
	}

	_, err = f.WriteString(string(data))

	if err != nil {
		return nil, err
	}

	err = f.Sync()

	if err != nil {
		return nil, err
	}

	return &config, nil
}

func WriteConfig(r *ConfigModels.RunningConfig) error {
	f, err := os.OpenFile(r.ConfigFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return err
	}

	defer f.Close()

	data, err := json.Marshal(r.Configuration)

	if err != nil {
		return err
	}

	_, err = f.WriteString(string(data))

	if err != nil {
		return err
	}

	err = f.Sync()

	if err != nil {
		return err
	}

	return nil
}
