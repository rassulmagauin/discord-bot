package translation

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var flagToLang = map[string]string{
	"🇺🇸": "en",
	"🇩🇪": "de",
	"🇫🇷": "fr",
	"🇪🇸": "es",
	"🇮🇹": "it",
	"🇵🇹": "pt",
	"🇷🇺": "ru",
	"🇦🇱": "sq",
	"🇸🇦": "ar",
	"🇧🇦": "bs",
	"🇧🇬": "bg",
	"🇨🇳": "zh-CN",
	"🇭🇷": "hr",
	"🇨🇿": "cs",
	"🇩🇰": "da",
	"🇪🇪": "et",
	"🇫🇮": "fi",
	"🇬🇷": "el",
	"🇭🇺": "hu",
	"🇮🇩": "id",
	"🇮🇳": "hi",
	"🇮🇪": "ga",
	"🇮🇸": "is",
	"🇮🇱": "he",
	"🇯🇵": "ja",
	"🇰🇷": "ko",
	"🇱🇻": "lv",
	"🇱🇹": "lt",
	"🇲🇹": "mt",
	"🇲🇪": "sr",
	"🇳🇱": "nl",
	"🇳🇴": "no",
	"🇵🇰": "ur",
	"🇵🇱": "pl",
	"🇷🇴": "ro",
	"🇷🇸": "sr",
	"🇸🇰": "sk",
	"🇸🇮": "sl",
	"🇸🇬": "sv",
	"🇹🇭": "th",
	"🇹🇷": "tr",
	"🇹🇼": "zh-TW",
	"🇺🇦": "uk",
	"🇻🇦": "la",
}

func ReactionAdd(sess *discordgo.Session, reaction *discordgo.MessageReactionAdd) {

	//Because call for API is expensive, it is better to run it in a  different goroutine
	go func() {
		// Check if the reaction is a flag emoji and get the corresponding language code
		targetLang, ok := flagToLang[reaction.Emoji.Name]
		if !ok {
			// If the emoji is not a flag we're interested in, ignore the reaction
			return
		}

		// Retrieve the original message
		msg, err := sess.ChannelMessage(reaction.ChannelID, reaction.MessageID)
		if err != nil {
			log.Printf("Error retrieving message: %v", err)
			sess.ChannelMessageSend(reaction.ChannelID, "Error: Unable to retrieve the message.")
			return
		}

		// Translate the message
		translatedText, err := translateText(msg.Content, targetLang)
		if err != nil {
			log.Printf("Error translating text: %v", err)
			sess.ChannelMessageSend(reaction.ChannelID, "Error: Unable to translate the message.")
			return
		}

		langName, ok := languageFullName[targetLang]
		if !ok {
			langName = targetLang // Fallback to language code
		}
		userMention := fmt.Sprintf("<@%s>", reaction.UserID)
		// Create a new embed message
		embed := &discordgo.MessageEmbed{
			Title:       "Translation",
			Description: fmt.Sprintf("Requested by %s", userMention),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Original Text",
					Value:  msg.Content,
					Inline: false,
				},
				{
					Name:   fmt.Sprintf("Translated to %s", langName),
					Value:  translatedText,
					Inline: false,
				},
			},
			Color: 0x00ff00, // Set the color of the embed
		}

		// Send the embedded message
		sess.ChannelMessageSendEmbed(reaction.ChannelID, embed)
	}()
}

// It is created to write language names beautifully
var languageFullName = map[string]string{
	"en":    "English",
	"de":    "German",
	"fr":    "French",
	"es":    "Spanish",
	"it":    "Italian",
	"pt":    "Portuguese",
	"ru":    "Russian",
	"sq":    "Albanian",
	"ar":    "Arabic",
	"bs":    "Bosnian",
	"bg":    "Bulgarian",
	"zh-CN": "Chinese (Simplified)",
	"hr":    "Croatian",
	"cs":    "Czech",
	"da":    "Danish",
	"et":    "Estonian",
	"fi":    "Finnish",
	"el":    "Greek",
	"hu":    "Hungarian",
	"id":    "Indonesian",
	"hi":    "Hindi",
	"ga":    "Irish",
	"is":    "Icelandic",
	"he":    "Hebrew",
	"ja":    "Japanese",
	"ko":    "Korean",
	"lv":    "Latvian",
	"lt":    "Lithuanian",
	"mt":    "Maltese",
	"sr":    "Serbian",
	"nl":    "Dutch",
	"no":    "Norwegian",
	"ur":    "Urdu",
	"pl":    "Polish",
	"ro":    "Romanian",
	"sk":    "Slovak",
	"sl":    "Slovenian",
	"sv":    "Swedish",
	"th":    "Thai",
	"tr":    "Turkish",
	"zh-TW": "Chinese (Traditional)",
	"uk":    "Ukrainian",
	"la":    "Latin",
}
