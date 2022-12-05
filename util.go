package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// GetChannelID to get channel id from channel name
func GetChannelID(s *discordgo.Session, channelName string) string {
	channels, err := s.GuildChannels("1021491552897474590")
	if err != nil {
		fmt.Println("Error getting channels: ", err)
	}
	for _, channel := range channels {
		if channel.Name == channelName {
			return channel.ID
		}
	}
	return ""
}

// GetChannelName to get channel name from channel id
func GetChannelName(s *discordgo.Session, channelID string) string {
	channels, err := s.GuildChannels("1021491552897474590")
	if err != nil {
		fmt.Println("Error getting channels: ", err)
	}
	for _, channel := range channels {
		if channel.ID == channelID {
			return channel.Name
		}
	}
	return ""
}
