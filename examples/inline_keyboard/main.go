package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Inline keyboard parameters
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow( // Row 1
			tu.InlineKeyboardButton("Callback data button 1"). // Column 1
										WithCallbackData("callback_1"),
			tu.InlineKeyboardButton("Callback data button 2"). // Column 2
										WithCallbackData("callback_2"),
		),
		tu.InlineKeyboardRow( // Row 2
			tu.InlineKeyboardButton("URL button").WithURL("https://example.com"), // Column 1
		),
	)

	// Message parameters
	message := tu.Message(
		tu.ID(1234567),
		"My message",
	).WithReplyMarkup(inlineKeyboard)

	// Sending message
	_, _ = bot.SendMessage(message)

	updates, _ := bot.UpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()

	// Receiving callback data
	for update := range updates {
		if update.CallbackQuery != nil {
			fmt.Println("Received callback with data:", update.CallbackQuery.Data)
		}
	}
}
