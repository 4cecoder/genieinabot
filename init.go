package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// InitAI initializes the AI for the Discord bot.
func InitAI(modelName string, prompt string) []string {
	completion, err := OpenAICompletion(modelName, prompt)
	if err != nil {
		return nil
	}
	return completion
}

// InitDiscord initializes a new Discord session using the provided bot token.
func InitDiscord(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return nil, err
	}

	return dg, nil
}

// InitHandlers initializes the event handlers for the Discord session.
func InitHandlers(dg *discordgo.Session) {
	dg.AddHandler(MessageCreate)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
}

// InitBot initializes the Discord bot.
func InitBot() (*discordgo.Session, error) {
	dg, err := InitDiscord(GetBotKeyEnv())
	if err != nil {
		return nil, err
	}

	InitHandlers(dg)
	return dg, nil
}

// RunBot runs the Discord bot until it receives a termination signal.
func RunBot(dg *discordgo.Session) {
	// Open a websocket connection to Discord and begin listening.
	err := dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("Bot is now exiting.")
	// Cleanly close down the Discord session.
	dg.Close()
}

// StartProgram initializes and runs the Discord bot.
func StartProgram() {
	dg, err := InitBot()
	if err != nil {
		return
	}

	RunBot(dg)
}
