package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel, all options are optional
	updates, _ := bot.UpdatesViaLongPolling(
		// Set Telegram parameter to get updates, can be nil
		// Note: If nil then timeout will be set to default 8s
		&telego.GetUpdatesParams{
			// Offset:  0, // Will be automatically updated by UpdatesViaLongPolling
			// Timeout: 8, // Can be set instead of using WithLongPollingUpdateInterval (default, recommended way)
		},

		// Set interval of getting updates (default: 0s).
		// If you want to get updates as fast as possible, set to 0 and explicitly set timeout in get updates
		// parameters, but the webhook method is recommended for this.
		telego.WithLongPollingUpdateInterval(time.Second*0),

		// Set retry timeout that will be used if an error occurs (default 8s)
		telego.WithLongPollingRetryTimeout(time.Second*8),

		// Set chan buffer (default 100)
		telego.WithLongPollingBuffer(100),
	)

	// Stop reviving updates from update channel
	defer bot.StopLongPolling()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %+v\n", update)
	}
}
