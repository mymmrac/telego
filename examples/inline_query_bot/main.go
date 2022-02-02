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

	updates, _ := bot.UpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()

	for update := range updates {
		// Receive inline query
		if update.InlineQuery != nil {
			iq := update.InlineQuery
			name := iq.From.FirstName

			// Answer inline query request
			_ = bot.AnswerInlineQuery(tu.InlineQuery(
				iq.ID,

				tu.ResultArticle(
					"hello",
					"Hello",
					&telego.InputTextMessageContent{ // Hello message with inline query
						ParseMode:   telego.ModeMarkdownV2,
						MessageText: fmt.Sprintf("Hello %s\n\nYour query:\n```%#+v```", name, iq),
					},
				).WithDescription(fmt.Sprintf("Query: %q", iq.Query)),

				tu.ResultArticle(
					"bey",
					"Bye",
					&telego.InputTextMessageContent{ // Bye message with inline query
						ParseMode:   telego.ModeMarkdownV2,
						MessageText: fmt.Sprintf("Bye %s\n\nYour query:\n```%#+v```", name, iq),
					},
				).WithDescription(fmt.Sprintf("Query: %q", iq.Query)),
			).WithIsPersonal())
		}
	}
}
