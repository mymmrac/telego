package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
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

	updates, _ := bot.GetUpdatesChan(&telego.GetUpdatesParams{})
	defer bot.StopGettingUpdates()

	for update := range updates {
		if update.Message != nil {
			// Retrieve chat ID
			chatID := update.Message.Chat.ID

			// Call method sendMessage. Sends message to sender with same text (echo bot)
			sentMessage, _ := bot.SendMessage(&telego.SendMessageParams{
				ChatID: telego.ChatID{ID: chatID},
				Text:   update.Message.Text,
			})

			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}
}
