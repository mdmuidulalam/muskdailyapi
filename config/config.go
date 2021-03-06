package config

import (
	"encoding/json"

	"os"
	"sync"

	enableServiceEnum "muskdaily.com/enums/enableService"
	environmentEnum "muskdaily.com/enums/environment"
)

type Configuration struct {
	Environment environmentEnum.Environment
	Port        string
	Database    struct {
		Path string
		Port string
	}
	Smtp struct {
		EnableService enableServiceEnum.EnableService
		Host          string
		Port          string
		UserName      string
		Password      string
	}
	RandomSeedOffset int64
}

var instance *Configuration
var once sync.Once

func GetConfiguration() *Configuration {
	once.Do(func() {
		file, _ := os.Open("./config/config.json")
		defer file.Close()
		decoder := json.NewDecoder(file)
		instance = &Configuration{}
		err := decoder.Decode(&instance)
		if err != nil {
			panic("Configuration parsing error!!!")
		}
	})

	return instance
}
