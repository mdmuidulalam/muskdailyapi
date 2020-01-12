package config

import (
	"encoding/json"
	"fmt"

	environmentEnum "muskdaily.com/enums"
	"os"
	"sync"
)

type configuration struct {
	Environment environmentEnum.Environment
	Port        string
	Database    struct {
		Path string
		Port string
	}
	Smtp struct {
		Host     string
		Port     string
		UserName string
		Password string
	}
}

var instance *configuration
var once sync.Once

func GetConfiguration() *configuration {
	once.Do(func() {
		file, _ := os.Open("./config/config.json")
		defer file.Close()
		decoder := json.NewDecoder(file)
		instance = &configuration{}
		err := decoder.Decode(&instance)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println("Called")
	})

	return instance
}
