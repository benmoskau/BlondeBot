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
		s.ChannelMessageSend(m.ChannelID, "Oh fuck right off")
	}

	if m.Content == "Henlo" {
		s.ChannelMessageSend(m.ChannelID, "Oh henlo there")
	}

	if m.Content == "<@!207951723292000259>" {
		s.ChannelMessageSend(m.ChannelID, "Yes?")
	}

	if m.Content == "test" {
		s.ChannelMessageSendTTS(m.ChannelID, "/tts Test 1 2")
	}
}
