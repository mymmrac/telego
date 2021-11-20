package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tg "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Call method getMe
	botUser, _ := bot.GetMe()
	fmt.Printf("Bot User: %#v\n", botUser)

	updates, _ := bot.GetUpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()

	for update := range updates {
		if update.Message != nil {
			// Retrieve chat ID
			chatID := update.Message.Chat.ID

			// Call method sendMessage (https://core.telegram.org/bots/api#sendmessage).
			// Sends message to sender with same text (echo bot).
			sentMessage, _ := bot.SendMessage(tg.Message(tg.ID(chatID), update.Message.Text))

			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}
}
