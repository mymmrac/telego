package main

import (
	"fmt"
	th "github.com/mymmrac/telego/telegohandler"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer bh.Stop()

	// Stop getting updates
	defer bot.StopLongPolling()

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
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		message := tu.Message(
			tu.ID(update.Message.Chat.ID),
			"My message",
		).WithReplyMarkup(inlineKeyboard)

		// Sending message
		_, _ = bot.SendMessage(message)
	}, th.AnyMessage())
	// Register new handler with match on a call back query with data equal to `go` and non-nil message
	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "GO"))

		// Answer callback query
		_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText("Done"))
	}, th.AnyCallbackQueryWithMessage())

	// Start handling updates
	bh.Start()

}
