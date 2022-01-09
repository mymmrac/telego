// This bot will send same messages as you sent to him.

package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Create Bot with debug on
	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPulling(nil)

	// Stop reviving updates from updates channel
	defer bot.StopLongPulling()

	// Loop through all updates when they came
	for update := range updates {
		// Check if update contains message
		if update.Message != nil {
			// Get chat ID from message
			chatID := tu.ID(update.Message.Chat.ID)

			// Copy sent message back to user
			_, _ = bot.CopyMessage(&telego.CopyMessageParams{
				ChatID:     chatID,
				FromChatID: chatID,
				MessageID:  update.Message.MessageID,
			})
		}
	}
}
