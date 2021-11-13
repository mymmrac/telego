package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	"github.com/valyala/fasthttp"
)

func main() {
	botToken := os.Getenv("TOKEN")

	var myLogger telego.Logger // Used just for example (not valid in real use)

	// Create bot
	bot, err := telego.NewBot(botToken,
		// Change bot API server URL (default: https://api.telegram.org)
		telego.SetAPIServer("new bot api server"),

		// Change HTTP client (default: &fasthttp.Client{})
		telego.FastHTTPClient(&fasthttp.Client{}),

		// Configuration of default logger, enable printing debug information and errors (default: false, true)
		telego.DefaultLogger(true, true),

		// Create you custom logger that implements telego.Logger (default: telego has build in default logger)
		telego.SetLogger(myLogger),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Call method getMe
	botUser, _ := bot.GetMe()
	fmt.Printf("Bot user: %#v\n", botUser)
}
