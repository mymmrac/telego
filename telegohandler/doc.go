/*
Package telegohandler provides handlers & predicates for Telego.

Bot handlers provide an easy way to make net/http like handlers, but with predicates instead of paths.

You can create BotHandler, register new handlers and start processing updates from the update channel which you provide.
All handlers process updates concurrently, but keep in mind that predicates are checked sequentially. This gives an
ability to process one update only with the first matched handler.

Example

This example shows how you can create BotHandler and register new handlers. Note, that order of registration directly
impacts order of checking matched handlers, and only the first matched handler will process the update.

	package main

	import (
		"fmt"
		"os"

		"github.com/mymmrac/telego"
		th "github.com/mymmrac/telego/telegohandler"
		tu "github.com/mymmrac/telego/telegoutil"
	)

	func main() {
		botToken := os.Getenv("TOKEN")

		bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Get updates channel
		updates, _ := bot.UpdatesViaLongPulling(nil)
		defer bot.StopLongPulling()

		// Create bot handler and specify from where to get updates
		bh, _ := th.NewBotHandler(bot, updates)

		// Register new handler with match on command `/start`
		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			// Send message
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
			))
		}, th.CommandEqual("start"))

		// Register new handler with match on any command
		// Handlers will match only once and in order of registration, so this handler will be called on any command
		// except `/start` command
		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			// Send message
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Unknown command, use /start",
			))
		}, th.AnyCommand())

		// Start handling updates
		bh.Start()

		// Stop handling updates
		defer bh.Stop()
	}

*/
package telegohandler
