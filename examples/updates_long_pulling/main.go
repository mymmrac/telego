package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set interval of getting updates (default: 0.5s)
	// If you want to get updates as fast as possible set to 0
	bot.SetUpdateInterval(time.Second / 2)

	// Get updates channel
	updates, _ := bot.GetUpdatesViaLongPulling(&telego.GetUpdatesParams{})

	// Stop reviving updates from updates channel
	defer bot.StopGettingUpdates()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %#v\n", update)
	}
}
