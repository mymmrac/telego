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

	// Register a handler with union predicate and not predicate
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		fmt.Println("Update with message text `Hmm?` or any other, but without message.")
	}, th.Or(
		th.Not(th.AnyMessage()), // Matches to any not message update
		th.TextEqual("Hmm?"),    // Matches to message update with a text `Hmm?`
	))

	// Register handler with message predicate and custom predicate
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		fmt.Println("Update with the message which text is longer than 7 chars.")
	},
		th.AnyMessage(), // Matches to any message update
		func(update telego.Update) bool { // Matches to message update with text longer then 7
			return len(update.Message.Text) > 7
		},
	)

	// Register handler with commands and specific args
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		fmt.Println("Update with command `start` without args or `help` with any args")
	}, th.TextContains("one"), th.TextPrefix("two"), th.TextSuffix("three"))

	// Start handling updates
	bh.Start()
}
