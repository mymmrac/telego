package main

import (
	"fmt"
	"os"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.DefaultLogger(true, true)

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
