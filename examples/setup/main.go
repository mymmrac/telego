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

	// Settings of default logger, enable printing debug information and errors (default: false, true)
	bot.DefaultLogger(true, true)

	var myLogger telego.Logger
	// Create you custom logger that implements telego.Logger (default: telego has build in default logger)
	bot.SetLogger(myLogger)

	// Call method getMe
	botUser, _ := bot.GetMe()
	fmt.Printf("Bot user: %#v\n", botUser)
}
