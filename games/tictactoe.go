package games

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type TicTacToeGame struct {
	Board       [3][3]string
	Player1     string // ID of player 1
	Player2     string //	ID of player 2
	CurrentTurn string // ID of the current player
	Winner      string //	ID of the winner
}

var activeGames = make(map[int]*TicTacToeGame)
var gameIDCounter = 1

// This function starts a new game of Tic Tac Toe
func StartTicTacToe(sess *discordgo.Session, msg *discordgo.MessageCreate, opponentID string) int {
	// Initialize game state, setting Player1 as the message author and Player2 as the mentioned user
	game := &TicTacToeGame{
		Player1:     msg.Author.ID,
		Player2:     opponentID,
		CurrentTurn: msg.Author.ID,
		Board:       [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}},
	}

	// Store the game state in a map using a unique gameID
	currentID := gameIDCounter
	activeGames[currentID] = game
	gameIDCounter++

	boardRepresentation := drawBoard(game.Board)
	player1Mention := fmt.Sprintf("<@%s>", game.Player1)
	player2Mention := fmt.Sprintf("<@%s>", game.Player2)
	currentTurnMention := fmt.Sprintf("<@%s>", game.CurrentTurn)

	// Create an embed for the game state message
	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Tic Tac Toe - Game ID: %d", currentID),
		Description: fmt.Sprintf("%s vs %s\n\nIt's now %s's turn!", player1Mention, player2Mention, currentTurnMention),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Board",
				Value:  "```" + boardRepresentation + "```",
				Inline: false,
			},
		},
		Color: 0x00ff00,
	}

	// Send the embedded message
	_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
	if err != nil {
		fmt.Printf("Error sending embed: %v", err)
		return -1
	}
	return currentID

}

// This function draws the game board
func drawBoard(board [3][3]string) string {
	var sb strings.Builder
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "" {
				sb.WriteString("   ")
			} else {
				sb.WriteString(fmt.Sprintf(" %s ", board[i][j]))
			}
			if j < 2 {
				sb.WriteString("|")
			}
		}
		if i < 2 {
			sb.WriteString("\n---+---+---\n")
		}
	}
	return sb.String()
}

// This function processes a move in the game
func ProcessMove(sess *discordgo.Session, msg *discordgo.MessageCreate, gameID string, row int, column int) {
	id, _ := strconv.Atoi(gameID)
	game, exists := activeGames[id]
	if !exists {
		sess.ChannelMessageSend(msg.ChannelID, "Game not found. Please ensure you have the correct game ID.")
		return
	}
	//Check if the player is a participant in the game
	if msg.Author.ID != game.Player1 && msg.Author.ID != game.Player2 {
		sess.ChannelMessageSend(msg.ChannelID, "You are not a participant in this game.")
		return
	}
	// Check if it's the player's turn
	if msg.Author.ID != game.CurrentTurn {
		sess.ChannelMessageSend(msg.ChannelID, "It's not your turn.")
		return
	}

	// Validate the move
	if row < 0 || row > 2 || column < 0 || column > 2 {
		sess.ChannelMessageSend(msg.ChannelID, "Invalid move. Row and column numbers must be between 1 and 3.")
		return
	}
	if game.Board[row][column] != " " {
		sess.ChannelMessageSend(msg.ChannelID, "This cell is already taken. Please choose another one.")
		return
	}

	// Make the move
	currentPlayerSymbol := "X"
	if game.Player1 == msg.Author.ID {
		currentPlayerSymbol = "X"
	} else if game.Player2 == msg.Author.ID {
		currentPlayerSymbol = "O"
	}
	game.Board[row][column] = currentPlayerSymbol

	// Check for a win or a tie
	if checkWin(game.Board, currentPlayerSymbol) {
		boardRepresentation := drawBoard(game.Board)
		embed := &discordgo.MessageEmbed{
			Title:       "Game Over",
			Description: fmt.Sprintf("<@%s> wins!", msg.Author.ID),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Final Board",
					Value:  "```" + boardRepresentation + "```",
					Inline: false,
				},
			},
			Color: 0xff0000, //Red
		}
		// Send the embedded message
		_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
		if err != nil {
			fmt.Printf("Error sending embed: %v", err)
		}
		delete(activeGames, id)
		return
	} else if checkTie(game.Board) {
		boardRepresentation := drawBoard(game.Board)
		embed := &discordgo.MessageEmbed{
			Title:       "Game Over",
			Description: "It's a tie!",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Final Board",
					Value:  "```" + boardRepresentation + "```",
					Inline: false,
				},
			},
			Color: 0x00ff00, //Green
		}
		// Send the embedded message
		_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
		if err != nil {
			fmt.Printf("Error sending embed: %v", err)
		}
		delete(activeGames, id)
		return
	}

	// Switch turns
	nextPlayerID := ""
	if game.CurrentTurn == game.Player1 {
		game.CurrentTurn = game.Player2
		nextPlayerID = game.Player2
	} else {
		game.CurrentTurn = game.Player1
		nextPlayerID = game.Player1
	}

	// Send the updated board
	boardRepresentation := drawBoard(game.Board)
	// Prepare an embedded message
	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Tic Tac Toe - Game ID: %d", id),
		Description: fmt.Sprintf("It's now %s's turn!", fmt.Sprintf("<@%s>", nextPlayerID)),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Current Board",
				Value:  "```" + boardRepresentation + "```",
				Inline: false,
			},
		},
		Color: 0x00ff00,
	}

	// Send the embedded message
	_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
	if err != nil {
		log.Println("Error sending embed: ", err)
	}

}

// check for a win
func checkWin(board [3][3]string, playerSymbol string) bool {
	// Check rows, columns, and diagonals for a win
	for i := 0; i < 3; i++ {
		if board[i][0] == playerSymbol && board[i][1] == playerSymbol && board[i][2] == playerSymbol {
			return true
		}
		if board[0][i] == playerSymbol && board[1][i] == playerSymbol && board[2][i] == playerSymbol {
			return true
		}
	}
	if board[0][0] == playerSymbol && board[1][1] == playerSymbol && board[2][2] == playerSymbol {
		return true
	}
	if board[0][2] == playerSymbol && board[1][1] == playerSymbol && board[2][0] == playerSymbol {
		return true
	}
	return false
}

// check for a tie
func checkTie(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false // If any cell is empty, it's not a tie
			}
		}
	}
	return true
}
