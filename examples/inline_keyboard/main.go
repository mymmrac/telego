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

	bot.DebugMode(true)

	// Inline keyboard parameters
	inlineKeyboard := &telego.InlineKeyboardMarkup{
		InlineKeyboard: [][]telego.InlineKeyboardButton{
			// Row 1
			{
				// Column 1
				telego.InlineKeyboardButton{
					Text:         "Callback data button 1",
					CallbackData: "callback_1",
				},
				// Column 2
				telego.InlineKeyboardButton{
					Text:         "Callback data button 2",
					CallbackData: "callback_2",
				},
			},
			// Row 2
			{
				// Column 1
				telego.InlineKeyboardButton{
					Text: "URL button",
					URL:  "https://example.com",
				},
			},
		},
	}

	// Message parameters
	message := &telego.SendMessageParams{
		ChatID:      telego.ChatID{ID: 1234567},
		Text:        "My message",
		ReplyMarkup: inlineKeyboard,
	}

	// Sending message
	_, _ = bot.SendMessage(message)

	updates, _ := bot.GetUpdatesChan(&telego.GetUpdatesParams{})
	defer bot.StopGettingUpdates()

	// Receiving callback data
	for update := range updates {
		if update.CallbackQuery != nil {
			fmt.Println("Received callback with data:", update.CallbackQuery.Data)
		}
	}
}
