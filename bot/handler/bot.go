package handler

import (
	"context"
	"fmt"
	"github.com/AllenKaplan/alphabot/meetup"
	"github.com/AllenKaplan/alphabot/weather"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"strings"
	"time"
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

	embeddedMsg := &discordgo.MessageEmbed{Author: &discordgo.MessageEmbedAuthor{},
		Title:       "Weather in " + res.Name + " | " + res.Weather[0].Main,
		Timestamp:   time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Color:       0x0000ff,                        // Blue
		Description: "`,weather get [location]`",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name: "Temperature",
				Value: fmt.Sprintf("%.1f	°C", res.Main.Temp),
			},
			&discordgo.MessageEmbedField{
				Name:  "Full Weather Desc",
				Value: strings.Title(res.Weather[0].Description),
			},
			&discordgo.MessageEmbedField{
				Name:  "Humidity",
				Value: strconv.Itoa(res.Main.Humidity),
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "http://openweathermap.org/img/w/" + res.Weather[0].Icon + ".png",
		},
	}

	if msg, err := client.session.ChannelMessageSendEmbed(client.message.ChannelID, embeddedMsg); err != nil {
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
	cmd = strings.TrimPrefix(cmd, "get")
	cmd = strings.TrimPrefix(cmd, " ")
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	if cmd == "all" {
		client.GetAllMeetups()
		return
	}

	meetupClient := meetup.NewMeetupClient(client.ctx)
	meetupRes, err := meetupClient.GetMeetup(cmd)
	if err != nil {
		log.Printf("go.bot.handler.GetMeetup getting meetup: %s | %v", cmd, err)
		return
	}

	embeddedMsg := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Title:       "Meetup Info",
		Timestamp:   meetupRes.Time.Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Color:       0xff00ff,                            // Blue
		Description: "Meetup Info `,meetup get [name]`",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Name",
				Value: meetupRes.Name,
			},
			&discordgo.MessageEmbedField{
				Name:  "Location",
				Value: meetupRes.Location,
			},
			&discordgo.MessageEmbedField{
				Name:  "Time",
				Value: meetupRes.Time.Format(time.RFC3339),
			},
		},
	}

	if msg, err := client.session.ChannelMessageSendEmbed(client.message.ChannelID, embeddedMsg); err != nil {
		log.Printf("error sending message | %v | %v", msg, err)
	}
}

func (client *Client) GetAllMeetups() {
	log.Printf("go.bot.handler.GetMeetup request recieved")

	cmd := client.ctx.Value("cmd").(string)
	cmd = strings.TrimPrefix(cmd, "get")
	cmd = strings.TrimPrefix(cmd, " ")
	client.ctx = context.WithValue(client.ctx, "cmd", cmd)

	meetupClient := meetup.NewMeetupClient(client.ctx)
	meetupRes, err := meetupClient.GetAllMeetups(cmd)
	if err != nil {
		log.Printf("go.bot.handler.GetMeetup getting meetups: %s | %v", cmd, err)
		return
	}

	var meetups string

	for _, curMeetup := range meetupRes {
		meetups = meetups + "\n" + curMeetup.String()
	}

	if msg, err := client.session.ChannelMessageSend(client.message.ChannelID, meetups); err != nil {
		log.Printf("error sending message | %v | %v", msg, err)
	}
}
