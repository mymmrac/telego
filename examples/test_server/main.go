package main

import (
	"fmt"
	"os"

	"github.com/chococola/telego"
)

// Using bots in the test environment:
// https://core.telegram.org/bots/webapps#using-bots-in-the-test-environment

func main() {
	// Get test Bot token from environment variables
	botToken := os.Getenv("TEST_TOKEN")

	// Create bot and enable debugging info
	// (more on configuration in /examples/configuration/main.go)
	bot, err := telego.NewBot(
		botToken,
		// Use the test server API path instead of the regular API path
		telego.WithTestServerPath(),
		// Set default logger with debug
		// Note: Please keep in mind that default logger may expose sensitive information, use in development only
		telego.WithDefaultDebugLogger(),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe (https://core.telegram.org/bots/api#getme).
	// Note: Bot will call the test server: https://api.telegram.org/bot<token>/test/getMe
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print Bot information
	fmt.Printf("Bot user: %+v\n", botUser)
}
