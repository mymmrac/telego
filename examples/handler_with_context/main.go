package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	ctx := context.Background()
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer func() { _ = bh.Stop() }()

	// Define a context key
	type userID bool
	var userIDKey userID

	// Apply middleware that will retrieve user ID from update
	bh.Use(func(ctx *th.Context, update telego.Update) error {
		if update.Message != nil && update.Message.From != nil {
			// Set user ID in context
			ctx = ctx.WithValue(userIDKey, update.Message.From.ID)
		}

		// Update context
		update = update.WithContext(ctx)
		return ctx.Next(update)
	})

	// Handle messages
	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		// Retrieve user ID from context
		fmt.Println("User ID:", ctx.Value(userIDKey))
		return nil
	}, th.AnyMessage())

	// Start handling updates
	_ = bh.Start()
}
