package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/rassulmagauin/discord-bot/utils"
)

// WeatherCommand function to get weather data from the API
func WeatherCommand(sess *discordgo.Session, msg *discordgo.MessageCreate, location string) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, location)

	resp, err := http.Get(url)
	if err != nil {
		utils.HandleError(sess, msg, err, "Error getting weather data.")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.HandleError(sess, msg, err, "Error reading weather data.")
		return
	}
	// saving the response to a map
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	currentWeather, ok := result["current"].(map[string]interface{})
	if !ok {
		utils.HandleError(sess, msg, err, "Error parsing weather data.")
		return
	}
	// Extracting weather data
	loc, ok := result["location"].(map[string]interface{})
	if !ok {
		utils.HandleError(sess, msg, err, "Error parsing weather data.")
		return
	}

	// Extracting needed weather data
	locName, _ := loc["name"].(string)
	tempC, _ := currentWeather["temp_c"].(float64)
	condition, _ := currentWeather["condition"].(map[string]interface{})
	weatherCondition, _ := condition["text"].(string)
	humidity, _ := currentWeather["humidity"].(float64)

	userMention := fmt.Sprintf("<@%s>", msg.Author.ID)

	// Creating an embed message
	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Weather in %s", locName),
		Description: fmt.Sprintf("%s, here's the current weather:", userMention),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Temperature",
				Value:  fmt.Sprintf("%.1f°C", tempC),
				Inline: true,
			},
			{
				Name:   "Condition",
				Value:  weatherCondition,
				Inline: true,
			},
			{
				Name:   "Humidity",
				Value:  fmt.Sprintf("%.0f%%", humidity),
				Inline: true,
			},
		},
		Color: 0x00ff00,
	}

	sess.ChannelMessageSendEmbed(msg.ChannelID, embed)

}
