package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Configuration struct {
	BotToken      string `json:"bot_token"`
	BotPrefix     string `json:"bot_prefix"`
	PermissionInt string `json:"permission_int"`
}

func InitConfig() Configuration {
	config := Configuration{}

	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error parsing config file: ", err)
	}
	fmt.Println("Config file loaded successfully.")

	return config
}
