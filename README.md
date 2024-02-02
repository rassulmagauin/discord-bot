
# Discord Bot

## Introduction

This Discord bot has versatile capabilities, including language translation, weather forecasts, interactive polling, and the ability to play Tic Tac Toe. 

## Features

- **Translation**: Instantly translates messages when reacted to with a country's flag emoji, allowing for seamless multilingual communication.
- **Weather Information**: Provides current weather updates for any specified location directly within your Discord server.
- **Polls**: Facilitates decision-making through the creation of customizable polls with multiple choices.
- **Tic Tac Toe**: Engages users in the classic game of Tic Tac Toe, supporting multiplayer gaming within the server.
- **Asynchronous Task Handling**: Performs time-consuming tasks, such as retrieving weather information or processing translations, in separate goroutines to maintain server responsiveness and user experience.

## Setup and Usage

### Initial Setup

1. **Clone the Repository**

   Clone the repository to your local system:

   ```bash
   git clone https://github.com/rassulmagauin/discord-bot.git
   cd discord-bot
   ```

2. **Discord Bot Configuration**
   - Navigate to the https://discord.com/developers/applications and create a new application.
   - Go to the Bot section and add a bot. Here, you will find your `DISCORD_BOT_TOKEN`.
   - Under the Bot permissions, select `Administrator` to give your bot full access to Discord API features. This is necessary for functionalities such as message reading, writing, and more.
   - Enable all intents under the Bot section to ensure your bot can listen to messages, reactions, and other events on Discord.
   - Use the OAuth2 section to generate an invite link for your bot. Make sure to select bot scope and assign it the necessary permissions (or simply `Administrator` for all permissions).
   - Invite the bot to your server using the generated invite link.

3. **Configure Environment**

   In the project's root directory, create a `.env` file containing your Discord bot token and the WeatherAPI.com API key:

   ```plaintext
   DISCORD_BOT_TOKEN=your_discord_bot_token
   WEATHER_API_KEY=your_weatherapi_key
   ```

4. **Install Dependencies**

   Ensure Go is installed on your system and install the project dependencies:

   ```bash
   go get .
   ```

5. **Launch the Bot**

   Execute the bot with:

   ```bash
   go run main.go
   ```

## Commands

- **Translate Messages**: React with a flag emoji to translate the message to the corresponding language.
- **Weather**: Use `!servant weather [location]` to retrieve weather information for the specified location.
- **Create Polls**: Use `!servant poll "Question" "Option1" "Option2"...` to initiate a poll with the provided options.
- **Play Tic Tac Toe**: Use `!servant start @opponent` to begin a game, followed by `!servant move [game id] [row] [column]` for gameplay.
- **Help**: Use `!servant help` to display a list of all available commands and their descriptions for quick reference.


### Presentation

- Demonstration of this bot you can see in this YouTube video: https://youtu.be/cF_W1WXTdCs

