package commands

import (
	"github.com/bwmarrin/discordgo"
)

// HelpCommand sends a message with the list of available commands. It also sends descriptions for each command.
func HelpCommand(sess *discordgo.Session, msg *discordgo.MessageCreate, prefix string) {
	helpMessage := prefix + " Here are the commands you can use:\n"
	helpMessage += "`!servant ping` - Responds with 'pong'.\n"
	helpMessage += "`!servant weather [location]` - Shows current weather information for the specified location (e.g., astana, almaty, new_york).\n"
	helpMessage += "`!servant poll \"Question\" \"Option1\" \"Option2\" ...` - Creates a poll with the specified question and options.\n"
	helpMessage += "`!tictac start @opponent` - Starts a new Tic Tac Toe game with the mentioned opponent.\n"
	helpMessage += "`!tictac move [game id] [row] [column]` - Makes a move in the specified Tic Tac Toe game. Rows and columns are 1-indexed.\n"
	helpMessage += "\nIn addition, you can use emoji reactions to translate messages:\n"
	helpMessage += "React to any message with a flag emoji, and I will translate the message to the corresponding language.\n"
	helpMessage += "\nFor Tic Tac Toe:\n"
	helpMessage += "Start a game by mentioning your opponent. Once the game starts, make your moves by specifying the game ID followed by the row and column where you want to place your mark (X or O). The game board is 3x3, and you can win by aligning three of your marks vertically, horizontally, or diagonally.\n"

	sess.ChannelMessageSend(msg.ChannelID, helpMessage)
}
