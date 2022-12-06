package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// GetChannelID to get channel id from channel name
func GetChannelID(s *discordgo.Session, channelName string, guildName string) string {
	channels, err := s.GuildChannels(GetGuildID(s, guildName))
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
func GetChannelName(s *discordgo.Session, channelID string, guildName string) string {
	channels, err := s.GuildChannels(GetGuildID(s, guildName))
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

// GetGuildID to get guild id from guild name (server name)
func GetGuildID(s *discordgo.Session, guildName string) string {
	guilds, err := s.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("Error getting guilds: ", err)
	}
	for _, guild := range guilds {
		if guild.Name == guildName {
			return guild.ID
		}
	}
	return ""
}
