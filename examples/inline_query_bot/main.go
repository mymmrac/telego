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
			_ = bot.AnswerInlineQuery(&telego.AnswerInlineQueryParams{
				InlineQueryID: iq.ID, // Answer with query ID that we received
				IsPersonal:    true,
				Results: []telego.InlineQueryResult{ // Provide list of results
					&telego.InlineQueryResultArticle{
						Type:        telego.ResultTypeArticle,
						ID:          "hello", // Unique result ID
						Title:       "Hello",
						Description: fmt.Sprintf("Query: %q", iq.Query),
						InputMessageContent: &telego.InputTextMessageContent{ // Hello message with inline query
							ParseMode:   telego.ModeMarkdownV2,
							MessageText: fmt.Sprintf("Hello %s\n\nYour query:\n```%#+v```", name, iq),
						},
					},
					&telego.InlineQueryResultArticle{
						Type:        telego.ResultTypeArticle,
						ID:          "bey", // Unique result ID
						Title:       "Bye",
						Description: fmt.Sprintf("Query: %q", iq.Query),
						InputMessageContent: &telego.InputTextMessageContent{ // Bye message with inline query
							ParseMode:   telego.ModeMarkdownV2,
							MessageText: fmt.Sprintf("Bye %s\n\nYour query:\n```%#+v```", name, iq),
						},
					},
				},
			})
		}
	}
}
