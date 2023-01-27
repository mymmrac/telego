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

	// Create Bot with debug on
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get bot user
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Bot user: %+v\n", botUser)

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler
	bh, _ := th.NewBotHandler(bot, updates)

	// Handle any message
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		// Get chat ID from message
		chatID := tu.ID(message.Chat.ID)

		// Copy sent message back to user
		_, _ = bot.CopyMessage(
			tu.CopyMessage(chatID, chatID, message.MessageID),
		)
	})

	// Stop handling updates on exit
	defer bh.Stop()
	defer bot.StopLongPolling()

	// Start handling updates
	bh.Start()
}
