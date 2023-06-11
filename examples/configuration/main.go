package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := os.Getenv("TOKEN")

	var myLogger telego.Logger // Used, just for example (not valid in real use)

	// Create bot, all options are optional
	bot, err := telego.NewBot(botToken,
		// Change bot API server URL (default: https://api.telegram.org)
		telego.WithAPIServer("new bot api server"),

		// Change caller to the Fast HTTP client (default: &fasthttp.Client{})
		telego.WithFastHTTPClient(&fasthttp.Client{}),

		// Change caller to the HTTP client (default: &fasthttp.Client{})
		telego.WithHTTPClient(&http.Client{}),

		// Enables basic health check that will call getMe method before returning bot instance (default: false)
		telego.WithHealthCheck(),

		// Make all warnings an errors for all requests (default: false)
		// Note: Things like `deleteWebhook` may return a result as true, but also error description with warning
		telego.WithWarnings(),

		// Configuration of default logger, enable printing debug information and errors (default: false, true)
		// Note: Please keep in mind that default logger may expose sensitive information, use in development only
		telego.WithDefaultLogger(false, true),

		// Extended configuration of default logger, enable printing debug information, errors and set replacer
		// (default: false, true, default replacer of bot token)
		// Note: Please keep in mind that default logger may expose sensitive information, use in development only
		telego.WithExtendedDefaultLogger(true, true, strings.NewReplacer("old", "new")),

		// Use default logger with enabled debug logs, same as telego.WithDefaultLogger(true, true)
		// Note: Please keep in mind that default logger may expose sensitive information, use in development only
		telego.WithDefaultDebugLogger(),

		// Use default logger with disabled debug and errors, same as telego.WithDefaultLogger(false, false)
		telego.WithDiscardLogger(),

		// Create you custom logger that implements telego.Logger (default: telego has built in default logger)
		// Note: Please keep in mind that logger may expose sensitive information, use in development only or configure
		// it not to leak unwanted content
		telego.WithLogger(myLogger),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe
	botUser, _ := bot.GetMe()
	fmt.Printf("Bot user: %+v\n", botUser)
}
