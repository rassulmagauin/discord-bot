package utils

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// HandleError is a function that handles errors. It sends a message to the user and logs the error
func HandleError(sess *discordgo.Session, msg *discordgo.MessageCreate, err error, userMessage string) {
	log.Printf("Error: %v\n", err)

	userMention := fmt.Sprintf("<@%s>", msg.Author.ID)

	if userMessage == "" {
		userMessage = fmt.Sprintf("%s An error occurred. Please try again later.", userMention)
	} else {
		userMessage = fmt.Sprintf("%s %s", userMention, userMessage)
	}

	sess.ChannelMessageSend(msg.ChannelID, userMessage)
}
