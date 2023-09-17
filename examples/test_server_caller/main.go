package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	ta "github.com/mymmrac/telego/telegoapi"
)

func main() {
	// Get Bot token from environment variables
	botToken := os.Getenv("TOKEN")

	// Create bot and enable debugging info
	// (more on configuration in /examples/configuration/main.go)
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(
		botToken,
		// Set up test server caller
		telego.WithAPICaller(&ta.TestServerCaller{
			// Use caller
			Caller: ta.DefaultFastHTTPCaller,
		}),
		telego.WithDefaultDebugLogger(),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe (https://core.telegram.org/bots/api#getme).
	// Bot will call the test server: "https://api.telegram.org/bot<token>/test/getMe"
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print Bot information
	fmt.Printf("Bot user: %+v\n", botUser)
}
