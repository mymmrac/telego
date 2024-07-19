package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chococola/telego"
	th "github.com/chococola/telego/telegohandler"
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

	// Define a context key
	type userID bool
	var userIDKey userID

	// Apply middleware that will retrieve user ID from update
	bh.Use(func(bot *telego.Bot, update telego.Update, next th.Handler) {
		// Get initial context
		ctx := update.Context()

		if update.Message != nil && update.Message.From != nil {
			// Set user ID in context
			ctx = context.WithValue(ctx, userIDKey, update.Message.From.ID)
		}

		// Update context
		update = update.WithContext(ctx)
		next(bot, update)
	})

	// Handle messages
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		ctx := update.Context()

		// Retrieve user ID from context
		fmt.Println("User ID:", ctx.Value(userIDKey))
	}, th.AnyMessage())

	// Start handling updates
	bh.Start()
}
