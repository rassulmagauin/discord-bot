package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Poll structure to store poll data
type Poll struct {
	Question string
	Options  []string
}

// For my bot, I decided to store the active polls in a map, not in a database.
// Even though it is not the best solution, it is enough for my bot.
var ActivePolls = make(map[string]*Poll)

//TODO: Add a way to close the poll

// It creates a poll
func CreatePollCommand(sess *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		sess.ChannelMessageSend(msg.ChannelID, "Usage: !servant poll \"Question\" \"Option1\" \"Option2\" ...")
		return
	}

	poll := &Poll{
		Question: args[0],
		Options:  args[1:],
	}

	ActivePolls[msg.ID] = poll

	embed := &discordgo.MessageEmbed{
		Title:       "Poll",
		Description: fmt.Sprintf("**%s**\n\nOptions:\n", poll.Question),
	}

	for i, option := range poll.Options {
		embed.Description += fmt.Sprintf("%s %s\n", string('🇦'+i), option)
	}

	pollMessage, _ := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)

	// Add initial reactions to the poll message
	for i := range poll.Options {
		sess.MessageReactionAdd(msg.ChannelID, pollMessage.ID, string('🇦'+i))
	}
}

//For simple poll function, the above funcation is enough.
// If needed it is possible to add a function to close the poll and get the results.
