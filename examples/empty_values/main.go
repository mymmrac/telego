package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger(),
		telego.WithEmptyValues(), // Enable empty values
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// In cases, when empty value is a valid parameter value, like here in telego.InlineKeyboardButton we have
	// SwitchInlineQueryCurrentChat that has an empty value as expected argument
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			telego.InlineKeyboardButton{
				Text: "Inline query in current chat",

				// Pass empty value that will be properly passed in request as "switch_inline_query_current_chat": ""
				// Warning: If telego.WithEmptyValues() or telego.WithCustomEmptyValues() bot options are not used,
				// this will do nothing
				SwitchInlineQueryCurrentChat: bot.EmptyValue(),
			},
		),
	)

	_, _ = bot.SendMessage(
		tu.Message(tu.ID(1234567), "Query").WithReplyMarkup(keyboard),
	)
}
