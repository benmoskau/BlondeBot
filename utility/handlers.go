package utility

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "SAD" {
		s.ChannelVoiceJoin("@!288502624360333312", "Oh fuck right off", true, true)
	}

	if m.Content == "<@&452330468826808321>" {
		s.ChannelMessageSend(m.ChannelID, "<@&452330468826808321>")
	}

	if m.Content == "Henlo" {
		s.ChannelMessageSend(m.ChannelID, "Oh henlo there")
	}

	if m.Content == "<@!207951723292000259>" {
		s.ChannelMessageSend(m.ChannelID, "<a:lol:519666784370950144>")
	}

	if m.Content == "test" {
		s.ChannelMessageSendTTS(m.ChannelID, "Test 1 2")
	}
}
