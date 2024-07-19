package main

import (
	"fmt"
	"os"
	"time"

	"github.com/chococola/telego"
	ta "github.com/chococola/telego/telegoapi"
)

func main() {
	// Get Bot token from environment variables
	botToken := os.Getenv("TOKEN")

	// Create bot and enable debugging info
	// (more on configuration in /examples/configuration/main.go)
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(
		botToken,
		// Set up retry caller
		telego.WithAPICaller(&ta.RetryCaller{
			// Use caller
			Caller: ta.DefaultFastHTTPCaller,
			// Max number of attempts to make call
			MaxAttempts: 4,
			// Exponent base for delay
			ExponentBase: 2,
			// Starting delay duration
			StartDelay: time.Millisecond * 10,
			// Maximum delay duration
			MaxDelay: time.Second,
		}),
		telego.WithDefaultDebugLogger(),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe (https://core.telegram.org/bots/api#getme).
	// In case if this call will fail, retry caller will retry calling Telegram until request is
	// successful (no network errors) or max attempts reached.
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print Bot information
	fmt.Printf("Bot user: %+v\n", botUser)
}
