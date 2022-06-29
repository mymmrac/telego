package main

import (
	"fmt"
	"os"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := os.Getenv("TOKEN")

	var myLogger telego.Logger // Used just for example (not valid in real use)

	// Create bot
	bot, err := telego.NewBot(botToken,
		// Change bot API server URL (default: https://api.telegram.org)
		telego.WithAPIServer("new bot api server"),

		// Change HTTP client (default: &fasthttp.Client{})
		telego.WithFastHTTPClient(&fasthttp.Client{}),

		// Configuration of default logger, enable printing debug information and errors (default: false, true)
		// Note: Please keep in mind that default logger exposes your bot token, use in development only
		telego.WithDefaultLogger(true, true),

		// Create you custom logger that implements telego.Logger (default: telego has build in default logger)
		telego.WithLogger(myLogger),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe
	botUser, _ := bot.GetMe()
	fmt.Printf("Bot user: %#v\n", botUser)
}
