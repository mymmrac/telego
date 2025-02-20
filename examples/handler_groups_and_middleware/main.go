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

	// Add global middleware, it will be applied in order of addition
	bh.Use(th.PanicRecovery()) // Will be called first
	bh.Use(
		func(ctx *th.Context, update telego.Update) error {
			fmt.Println("Global middleware") // Will be called second
			return ctx.Next(update)
		},
		func(ctx *th.Context, update telego.Update) error {
			fmt.Println("Global middleware 2") // Will be called third
			return ctx.Next(update)
		},
	)

	// Create any groups with or without predicates
	// Note: Updates first checked by groups and only then by handlers (group -> ... -> group -> handler)
	task := bh.Group(th.TextContains("task"))

	// Add middleware to groups
	task.Use(func(ctx *th.Context, update telego.Update) error {
		fmt.Println("Group-based middleware") // Will be called fourth

		if len(update.Message.Text) < 10 {
			return ctx.Next(update)
		}

		return nil
	})

	// Handle updates on a group
	task.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		fmt.Println("Task...") // Will be called fifth
		return err
	})

	// Start handling updates
	_ = bh.Start()
}
