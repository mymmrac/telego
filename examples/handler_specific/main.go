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

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer func() { _ = bh.Stop() }()

	// Register new handler with match on command `/start`
	bh.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		// Send a message with inline keyboard
		_, _ = bot.SendMessage(ctx, tu.Messagef(
			tu.ID(message.Chat.ID),
			"Hello %s!", message.From.FirstName,
		).WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(tu.InlineKeyboardButton("Go!").WithCallbackData("go"))),
		))
		return nil
	}, th.CommandEqual("start"))

	// Register new handler with match on a call back query with data equal to `go` and non-nil message
	bh.HandleCallbackQuery(func(ctx *th.Context, query telego.CallbackQuery) error {
		// Send message
		_, _ = bot.SendMessage(ctx, tu.Message(tu.ID(query.Message.GetChat().ID), "GO"))

		// Answer callback query
		_ = bot.AnswerCallbackQuery(ctx, tu.CallbackQuery(query.ID).WithText("Done"))

		return nil
	}, th.AnyCallbackQueryWithMessage(), th.CallbackDataEqual("go"))

	// Start handling updates
	_ = bh.Start()
}
