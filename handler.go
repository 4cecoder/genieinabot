package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// MessageCreate function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// if message is "ping" reply with "pong"
	if m.Content == "ping" {
		SendMessage(s, m.ChannelID, "pong")
	}
	// set channel to listen to general channel

}

// SendMessage sends a message to the specified channel.
func SendMessage(dg *discordgo.Session, channelID string, message string) (*discordgo.Message, error) {
	// Send the message to the channel.
	msg, err := dg.ChannelMessageSend(channelID, message)
	if err != nil {
		fmt.Println("Error sending message: ", err)
		return nil, err
	}

	return msg, nil
}
