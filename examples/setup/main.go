package main

import (
	"fmt"
	"net/http"
	"os"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Create bot
	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Change bot token (default: set by telego.NewBot(...))
	_ = bot.SetToken("new bot token")

	// Change bot API server URL (default: https://api.telegram.org)
	_ = bot.SetAPIServer("new bot api server")

	// Change http client (default: http.DefaultClient)
	_ = bot.SetClient(http.DefaultClient)

	// Enable printing debug information (default: false)
	bot.DebugMode(true)

	// Enable printing errors (default: true)
	bot.PrintErrors(true)
}
