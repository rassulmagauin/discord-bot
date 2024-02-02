package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/rassulmagauin/discord-bot/commands"
	"github.com/rassulmagauin/discord-bot/games"
	"github.com/rassulmagauin/discord-bot/translation"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	token := os.Getenv("DISCORD_BOT_TOKEN")
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Add handlers for the bot
	sess.AddHandler(commands.MessageCreate)
	sess.AddHandler(translation.ReactionAdd)
	sess.AddHandler(games.TicTac)

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer sess.Close()
	log.Println("Bot is running...")

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	log.Println("Shutting down...")
}
