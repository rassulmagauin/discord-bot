package games

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// prefix is defined
const tictac = "!tictac"

func TicTac(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	// Check if the message is from the bot itself to avoid self-response
	if msg.Author.ID == sess.State.User.ID {
		return
	}

	split := strings.Split(msg.Content, " ")
	// Check for the 'start' command to create a new game
	if len(split) >= 2 && split[0] == tictac {
		switch split[1] {
		case "start":
			if len(split) < 3 {
				sess.ChannelMessageSend(msg.ChannelID, "Please mention an opponent to start a game. Usage: `!servant start @opponent`")
				return
			}
			opponentID := split[2]

			// Extracting the ID from the mention
			// <@USER_ID> or <@!USER_ID> for nickname mentions
			re := regexp.MustCompile(`<@!?(\d+)>`)
			matches := re.FindStringSubmatch(opponentID)
			if len(matches) < 2 {
				sess.ChannelMessageSend(msg.ChannelID, "Invalid opponent format. Please make sure you're mentioning a user.")
				return
			}
			opponentID = matches[1] // ID of the mentioned user
			if opponentID == msg.Author.ID {
				sess.ChannelMessageSend(msg.ChannelID, "You can't play against yourself.")
				return
			}
			StartTicTacToe(sess, msg, opponentID)
			return
		case "move":
			if len(split) == 5 { // !servant move [game id] [row] [column]
				gameID := split[2]
				row, errRow := strconv.Atoi(split[3])
				column, errCol := strconv.Atoi(split[4])
				if errRow != nil || errCol != nil || row < 1 || row > 3 || column < 1 || column > 3 {
					sess.ChannelMessageSend(msg.ChannelID, "Invalid move. Please use the format: !servant move [game id] [row] [column] where row and column are between 1 and 3.")
					return
				}
				// Adjust row and column to be zero-indexed for internal representation
				ProcessMove(sess, msg, gameID, row-1, column-1)
			} else {
				sess.ChannelMessageSend(msg.ChannelID, "Invalid command format. Please use: !servant move [game id] [row] [column]")
			}
			return
		default:
			sess.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Unknown command: %s. Please use a valid command.", split[1]))
		}
	}
}
