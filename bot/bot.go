package bot

import (
	"fmt"

	"github.com/SteakBarbare/Gobot2025/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("Sessions creation failed")
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "cc" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "sv ?")
	}

	if m.Author.ID == "199323571846119425" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ta gueule")
	}

	if m.Author.ID == "1158746428559073302" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Poufiasse")
	}

	if m.Author.ID == "396248969904259074" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "t ki ?")
	}

	if m.Author.ID == "964306682085773332" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pepe Chicken mec")
	}
}
