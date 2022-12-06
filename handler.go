package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// MessageCreate function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	config := InitConfig()
	s.Identify.Intents = discordgo.IntentsMessageContent

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Ignore all messages not starting with the bot's prefix
	if m.Content[0:len(config.BotPrefix)] != config.BotPrefix {
		return
	}

	// if the beginning of the m.Content is "!ada " then send the rest of the message to the AI
	if m.Content[0:len(config.BotPrefix+"genie ")] == config.BotPrefix+"genie " {
		contentRecieved := m.Content

		contentRecieved = contentRecieved[len(config.BotPrefix):]
		contentRecieved = contentRecieved[len("genie"):]
		prompt := GetPrompt()

		AIResponse := InitAI("davinci", prompt)

		// Clean up the response from the AI stop characters at . or ! or ?

		//AIResponse = AIResponse[0:findStop(AIResponse)]

		// if AIResponse is empty, ask the user to rephrase the question
		if len(AIResponse) == 0 {
			AIResponse = "I'm sorry, I didn't quite understand that. Could you rephrase your question?"
		}

		AIResponse = refineGpt3Response(AIResponse)
		_, err := SendMessage(s, m.ChannelID, AIResponse)
		if err != nil {
			fmt.Println("Error sending message: ", err)
		}
	}

	if m.Content == config.BotPrefix+"help" {
		_, err := SendMessage(s, m.ChannelID, "Commands: `!help`, `!genie` + `<your question>`")
		if err != nil {
			fmt.Println("Error sending message: ", err)
		}

	}

}

// findStop finds the last stop character in the string and only returns the string up to that point
func findStop(response string) int {
	for i := 0; i < len(response); i++ {
		if response[i] == '.' || response[i] == '!' || response[i] == '?' {
			return i
		}
	}
	return len(response)
}

// func filterWithStop(response string, stop int) string {
//	return response[0:stop]
// }

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
