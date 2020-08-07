package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/J-Rivard/pso2-filter/internal/config"
	"github.com/J-Rivard/pso2-filter/internal/logging"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func main() {

	logger, err := logging.New(zerolog.ConsoleWriter{Out: os.Stderr}, logging.Debug)
	if err != nil {
		panic(err)
	}

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	fmt.Println(logger)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + cfg.BotToken)
	if err != nil {
		fmt.Println("error creating Discord session:", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection:", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if msg.Author.ID == session.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	// if msg.Content == "ping" {
	// 	session.ChannelMessageSend(msg.ChannelID, "Pong!")
	// }

	// // If the message is "pong" reply with "Ping!"
	// if msg.Content == "pong" {
	// 	session.ChannelMessageSend(msg.ChannelID, "Ping!")
	// }
	fmt.Println(msg.Content, msg.Attachments)
}
