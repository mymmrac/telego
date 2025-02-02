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

	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer func() { _ = bh.Stop() }()

	// Should not be here, because the order of handlers does meter.
	//
	// bh.HandleMessage(func(ctx *th.Context, message telego.Message) error {
	// 	 fmt.Println("Message:", message.Text)
	//	 return nil
	// })
	//
	// When you are defining handlers only the first matched handler will process update, that means that in this
	// specific example, matching any message will "shadow" any other handler that matches on messages.
	// For any message it will match and handling of start command will never happen.
	//
	// The general idea is that you should define the most specific handlers first (like for specific command or etc.)
	// and only then more generic handlers (like any message).

	// Will match any message with command `/start`
	bh.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		fmt.Println("Start")
		return nil
	}, th.CommandEqual("start"))

	// Will match to any message
	bh.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		fmt.Println("Message:", message.Text)
		return nil
	})

	// Start handling
	_ = bh.Start()
}
