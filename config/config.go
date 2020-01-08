package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type configuration struct {
	Port     string
	Database struct {
		Path string
		Port string
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
