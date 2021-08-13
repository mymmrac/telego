package main

import (
	"fmt"
	"os"
	"time"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.DefaultLogger(true, true)

	// Set interval of getting updates (default: 0.5s)
	// If you want to get updates as fast as possible set to 0
	bot.SetUpdateInterval(time.Second / 2)

	// Get updates channel
	updates, _ := bot.GetUpdatesChan(&telego.GetUpdatesParams{})

	// Stop reviving updates from updates channel
	defer bot.StopGettingUpdates()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %#v\n", update)
	}
}
