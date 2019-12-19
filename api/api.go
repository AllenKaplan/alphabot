package api

import (
	"context"
	"github.com/AllenKaplan/alphabot/weather"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type AlphaBot struct {
	ctx 	context.Context
	session   *discordgo.Session
	message *discordgo.MessageCreate
}

type Endpoints interface {
	messageCreate(m *discordgo.MessageCreate)
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := ","
	ctx := context.WithValue(context.Background(), "prefix", prefix)
	bot := AlphaBot{ctx, s, m}

	cmd := bot.message.Content
	bot.ctx = context.WithValue(bot.ctx, "cmd", cmd)

	if bot.message.Author.ID == bot.session.State.User.ID {
		return
	}

	if !strings.HasPrefix(bot.message.Content, prefix){
		return
	}

	log.Print("Prefix hit")
	cmd = strings.TrimPrefix(bot.message.Content, prefix)
	bot.ctx = context.WithValue(bot.ctx, "cmd", cmd)

	var commands = map[string]interface{}{
		"weather": bot.Weather,
	}

	for command, handleFunc := range commands {
		if strings.HasPrefix(cmd, command) {
			handleFunc.(func())()
		}
	}
}

func (bot AlphaBot) Weather() {
	log.Printf("Weather hit")
	cmd := bot.ctx.Value("cmd").(string)
	cmd = strings.TrimPrefix(cmd, "weather ")
	bot.ctx = context.WithValue(bot.ctx, "cmd", cmd)

	cmd = bot.ctx.Value("cmd").(string)

	w := weather.NewWeatherService()
	feeling, err := w.GetWeatherByLocation(cmd)
	if err != nil {
		log.Printf("error getting weather: %s | %v", cmd, err)
	}

	if msg, err := bot.session.ChannelMessageSend(bot.message.ChannelID, feeling); err != nil {
		log.Printf("error sending message | %v | %v", msg, err)
	}

}

