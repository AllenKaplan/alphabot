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
	ctx     context.Context
	session *discordgo.Session
	message *discordgo.MessageCreate
}

type Router struct {
	routes map[string]interface{}
}

func (r *Router) handleRoute(c *Client) {
	cmd := c.ctx.Value("cmd").(string)

	for routeName, routeFunc := range r.routes {
		if strings.HasPrefix(cmd, routeName) {
			log.Printf("Route found based on command | %s -> %s", cmd, routeName)
			routeFunc.(func())()
		}
	}
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := ","
	ctx := context.WithValue(context.Background(), "prefix", prefix)
	client := Client{ctx, s, m}

	cmd := client.message.Content
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	if client.message.Author.ID == client.session.State.User.ID {
		return
	}

	if !strings.HasPrefix(client.message.Content, prefix) {
		return
	}

	log.Print("Prefix hit")
	cmd = strings.TrimPrefix(client.message.Content, prefix)
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	routes := map[string]interface{}{
		"weather": client.Weather,
		"meetup":  client.Meetup,
	}

	router := &Router{routes: routes}

	router.handleRoute(&client)
}

func (client *Client) Weather() {
	log.Printf("go.bot.handler.Weather request recieved")

	cmd := client.ctx.Value("cmd").(string)
	cmd = strings.TrimPrefix(cmd, "weather ")
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	routes := map[string]interface{}{
		"get": client.GetWeather,
	}

	weatherRouter := &Router{routes: routes}

	weatherRouter.handleRoute(client)
}

func (client *Client) Meetup() {
	log.Printf("go.bot.handler.Meetup request recieved")

	cmd := client.ctx.Value("cmd").(string)
	cmd = strings.TrimPrefix(cmd, "meetup ")
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	routes := map[string]interface{}{
		"get":    client.GetMeetup,
		"create": client.CreateMeetup,
	}

	meetupRouter := &Router{routes: routes}

	meetupRouter.handleRoute(client)
}

func (client *Client) GetWeather() {
	log.Printf("go.bot.handler.Weather.Get request recieved")

	cmd := client.ctx.Value("cmd").(string)
	cmd = strings.TrimPrefix(cmd, "get ")
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

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
	log.Printf("go.bot.handler.CreateMeetup request recieved")

	cmd := client.ctx.Value("cmd").(string)
	cmd = strings.TrimPrefix(cmd, "create ")
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	weatherClient := meetup.NewMeetupClient(client.ctx)
	res, err := weatherClient.CreateMeetup(cmd)
	if err != nil {
		log.Printf("go.bot.handler.CreateMeetup : %s | %v", cmd, err)
	}

	if msg, err := client.session.ChannelMessageSend(client.message.ChannelID, res); err != nil {
		log.Printf("go.bot.handler.CreateMeetup  | %v | %v", msg, err)
	}
}

func (client *Client) GetMeetup() {
	log.Printf("go.bot.handler.GetMeetup request recieved")

	cmd := client.ctx.Value("cmd").(string)
	cmd = strings.TrimPrefix(cmd, "get ")
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	res := "Could not get meetup"
	meetupClient := meetup.NewMeetupClient(client.ctx)
	meetupRes, err := meetupClient.GetMeetup(cmd)
	if err != nil {
		log.Printf("go.bot.handler.GetMeetup getting meetup: %s | %v", cmd, err)
	} else {
		res = meetupRes.String()
	}

	if msg, err := client.session.ChannelMessageSend(client.message.ChannelID, res); err != nil {
		log.Printf("go.bot.handler.GetMeetup sending message | %v | %v", msg, err)
	}
}
