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

	// Add global middleware, it will be applied in order of addition
	bh.Use(th.PanicRecovery()) // Will be called first
	bh.Use(
		func(bot *telego.Bot, update telego.Update, next th.Handler) {
			fmt.Println("Global middleware") // Will be called second
			next(bot, update)
		},
		func(bot *telego.Bot, update telego.Update, next th.Handler) {
			fmt.Println("Global middleware 2") // Will be called third
			next(bot, update)
		},
	)

	// Create any groups with or without predicates
	// Note: Updates first checked by groups and only then by handlers (group -> ... -> group -> handler)
	task := bh.Group(th.TextContains("task"))

	// Add middleware to groups
	task.Use(func(bot *telego.Bot, update telego.Update, next th.Handler) {
		fmt.Println("Group-based middleware") // Will be called fourth

		if len(update.Message.Text) < 10 {
			next(bot, update)
		}
	})

	// Handle updates on a group
	task.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		fmt.Println("Task...") // Will be called fifth
	})

	// Start handling updates
	bh.Start()
}
