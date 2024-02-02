```markdown
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

2. **Configure Environment**

   In the project's root directory, create a `.env` file containing your Discord bot token and the WeatherAPI.com API key:

   ```plaintext
   DISCORD_BOT_TOKEN=your_discord_bot_token
   WEATHER_API_KEY=your_weatherapi_key
   ```

3. **Install Dependencies**

   Ensure Go is installed on your system and install the project dependencies:

   ```bash
   go get .
   ```

4. **Launch the Bot**

   Execute the bot with:

   ```bash
   go run main.go
   ```

### Commands

- **Translate Messages**: React with a flag emoji to translate the message to the corresponding language.
- **Weather**: `!servant weather [location]` - Retrieves weather information for the specified location.
- **Create Polls**: `!servant poll "Question" "Option1" "Option2"...` - Initiates a poll with provided options.
- **Play Tic Tac Toe**: `!servant start @opponent` to begin a game, followed by `!servant move [game id] [row] [column]` for gameplay.
- **Asynchronous Operations**: The bot efficiently handles API requests and translations asynchronously, ensuring seamless user interactions.

## Contributing

We welcome contributions of all forms. If you're interested in enhancing the bot's functionality or suggesting improvements, please fork the repository, commit your changes, and open a pull request.

## License

This project is made available under the MIT License. For more details, see the LICENSE file.
```
Ensure to replace `https://example.com/discord-bot.git` with your actual repository URL and adjust any details as necessary to match your bot's features and setup process.