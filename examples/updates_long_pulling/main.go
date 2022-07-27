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
	updates, _ := bot.UpdatesViaLongPulling(nil,
		// Set interval of getting updates (default: 0.5s).
		// If you want to get updates as fast as possible set to 0, but webhook method is recommended for this.
		telego.WithLongPullingUpdateInterval(time.Second/2),

		// Set retry timeout that will be used if an error occurs (default 3s)
		telego.WithLongPullingRetryTimeout(time.Second*3),

		// Set chan buffer (default 100)
		telego.WithLongPullingBuffer(100),
	)

	// Stop reviving updates from updates channel
	defer bot.StopLongPulling()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %+v\n", update)
	}
}
