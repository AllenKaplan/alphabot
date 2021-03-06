package bot

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Variables used for command line parameters
var (
	token string
)

func init() {
	flag.StringVar(&token, "t", "", "Client Token")
	flag.Parse()
}

func main() {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
	}

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		log.Fatalf("error opening connection: %v", err)
	}

	// Cleanly close down the Discord session.
	defer discord.Close()

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(bot.NewBotHandler)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Client is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
