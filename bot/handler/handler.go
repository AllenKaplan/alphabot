package handler

import (
	"context"
	"github.com/AllenKaplan/alphabot/meetup"
	"github.com/AllenKaplan/alphabot/weather"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type Client struct {
	ctx 	context.Context
	session   *discordgo.Session
	message *discordgo.MessageCreate
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := ","
	ctx := context.WithValue(context.Background(), "prefix", prefix)
	client := Client{ ctx, s, m}

	cmd := client.message.Content
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	if client.message.Author.ID == client.session.State.User.ID {
		return
	}

	if !strings.HasPrefix(client.message.Content, prefix){
		return
	}

	log.Print("Prefix hit")
	cmd = strings.TrimPrefix(client.message.Content, prefix)
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	var routes = map[string]interface{}{
		"weather": client.GetWeather,
		"meetup get": client.CreateMeetup,
	}

	if handleFunc, ok := routes[cmd]; ok {
		handleFunc.(func())()
	}
}

func (client *Client) GetWeather() {
	log.Printf("go.bot.handler.GetWeather request recieved")

	cmd := client.ctx.Value("cmd").(string)
	cmd = strings.TrimPrefix(cmd, "weather ")
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	cmd = client.ctx.Value("cmd").(string)

	weatherClient := weather.NewWeatherClient(client.ctx)
	res, err := weatherClient.GetWeatherByLocation(cmd)
	if err != nil {
		log.Printf("error getting weather: %s | %v", cmd, err)
	}

	if msg, err := client.session.ChannelMessageSend(client.message.ChannelID, res); err != nil {
		log.Printf("error sending message | %v | %v", msg, err)
	}
}

func (client *Client) CreateMeetup() {
	log.Printf("go.bot.handler.GetWeather request recieved")

	cmd := client.ctx.Value("cmd").(string)
	cmd = strings.TrimPrefix(cmd, "meetup create")
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	cmd = client.ctx.Value("cmd").(string)

	weatherClient := meetup.NewMeetupClient(client.ctx)
	res, err := weatherClient.CreateMeetup(cmd)
	if err != nil {
		log.Printf("error getting weather: %s | %v", cmd, err)
	}

	if msg, err := client.session.ChannelMessageSend(client.message.ChannelID, res); err != nil {
		log.Printf("error sending message | %v | %v", msg, err)
	}
}

func (client *Client) CreateMeetup() {
	log.Printf("go.bot.handler.GetWeather request recieved")

	cmd := client.ctx.Value("cmd").(string)
	cmd = strings.TrimPrefix(cmd, "meetup get")
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	cmd = client.ctx.Value("cmd").(string)

	weatherClient := meetup.NewMeetupClient(client.ctx)
	res, err := weatherClient.GetMeetup(cmd)
	if err != nil {
		log.Printf("error getting weather: %s | %v", cmd, err)
	}

	if msg, err := client.session.ChannelMessageSend(client.message.ChannelID, res); err != nil {
		log.Printf("error sending message | %v | %v", msg, err)
	}
}

