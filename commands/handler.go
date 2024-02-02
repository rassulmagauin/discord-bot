package commands

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix = "!servant"

// MessageCreate is a function that handles all the commands
func MessageCreate(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	// Check if the message is from the bot itself
	if msg.Author.ID == sess.State.User.ID {
		return
	}

	split := strings.Split(msg.Content, " ")
	if len(split) == 0 || split[0] != prefix {
		return
	}

	userMention := fmt.Sprintf("<@%s>", msg.Author.ID)

	if len(split) == 1 {
		response := fmt.Sprintf("%s No command provided. For a list of commands, type `!servant help`.", userMention)
		sess.ChannelMessageSend(msg.ChannelID, response)
		return
	}

	switch split[1] {
	case "ping":
		response := fmt.Sprintf("%s pong", userMention)
		sess.ChannelMessageSend(msg.ChannelID, response)
	case "help", "h":
		response := fmt.Sprintf("%s ", userMention)
		HelpCommand(sess, msg, response)
	case "weather":
		if len(split) < 3 {
			sess.ChannelMessageSend(msg.ChannelID, "Please specify a location.")
			return
		}
		location := strings.Join(split[2:], "_")
		// running the weather command in new goroutine, because it can take a while to get the response
		go WeatherCommand(sess, msg, location)
	case "poll":
		re := regexp.MustCompile(`"[^"]+"`)
		matches := re.FindAllString(msg.Content, -1)

		if len(matches) < 2 {
			sess.ChannelMessageSend(msg.ChannelID, "Usage: !servant poll \"Question\" \"Option1\" \"Option2\" ...")
			return
		}

		// Remove quotes from matches
		for i, match := range matches {
			matches[i] = strings.Trim(match, "\"")
		}

		go CreatePollCommand(sess, msg, matches) // Handle poll command in a new goroutine
	default:
		response := fmt.Sprintf("%s Unknown command", userMention)
		sess.ChannelMessageSend(msg.ChannelID, response)
	}

}
