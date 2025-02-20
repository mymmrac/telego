package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	ctx := context.Background()
	botToken := os.Getenv("TOKEN")

	// Create Bot with debug on
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get bot user
	botUser, err := bot.GetMe(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Bot user: %+v\n", botUser)

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	// Create bot handler
	bh, _ := th.NewBotHandler(bot, updates)

	// Handle any message
	bh.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		// Get chat ID from the message
		chatID := tu.ID(message.Chat.ID)

		// Copy sent messages back to the user
		_, _ = bot.CopyMessage(ctx, tu.CopyMessage(chatID, chatID, message.MessageID))

		return nil
	})

	// Stop handling updates on exit
	defer func() { _ = bh.Stop() }()

	// Start handling updates
	_ = bh.Start()
}
