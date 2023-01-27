package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer bh.Stop()

	// Stop getting updates
	defer bot.StopLongPolling()

	// Define auth predicate
	auth := func(update telego.Update) bool {
		var userID int64

		// Get user ID from message
		if update.Message != nil && update.Message.From != nil {
			userID = update.Message.From.ID
		}

		// Get user ID from callback query
		if update.CallbackQuery != nil {
			userID = update.CallbackQuery.From.ID
		}

		// Reject if no user
		if userID == 0 {
			return false
		}

		// Accept
		if userID == 1234 {
			return true
		}

		return false
	}

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// DO AUTHORIZED STUFF...
	}, auth) // Check for authorization

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// DO NOT AUTHORIZED STUFF...
	}, th.Not(auth)) // Process unauthorized update

	// Start handling updates
	bh.Start()
}
