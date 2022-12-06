package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	BotPrefix     string `json:"bot_prefix"`
	PermissionInt string `json:"permission_int"`
	GuildName     string `json:"guild_name"`
	ChannelName   string `json:"channel_name"`
	ModelName     string `json:"model_name"`
	MaxTokens     string `json:"max_tokens"`
	Prompt        string `json:"prompt"`
}

func InitConfig() Configuration {
	config := Configuration{}

	file, err := os.ReadFile("config.json")
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

func GetMax() string {
	config := InitConfig()
	return config.MaxTokens
}

// GetAIKeyEnv to get OpenAI API key from environment variable
func GetAIKeyEnv() string {
	return os.Getenv("GPT_KEY")
}

// GetBotKeyEnv to get OpenAI API key from environment variable
func GetBotKeyEnv() string {
	return os.Getenv("BOT_KEY")
}

// GetModel to get OpenAI model name from config file
func GetModel() string {
	config := InitConfig()
	return config.ModelName
}

// GetPrompt to get OpenAI prompt from config file
func GetPrompt() string {
	config := InitConfig()
	return config.Prompt
}
