package main

import (
	"fmt"
	"os"

	"github.com/chococola/telego"
	tu "github.com/chococola/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Edit message text
	_, _ = bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:    tu.ID(1234567),
		MessageID: 1234, // Message ID can be retried when it's sent or with update
		Text:      "<h1>New text</h1>",
		ParseMode: telego.ModeHTML,
	})

	// Edit message caption & reply markup (inline keyboard)
	_, _ = bot.EditMessageCaption(&telego.EditMessageCaptionParams{
		ChatID:    tu.ID(1234567),
		MessageID: 1234, // Message ID can be retried when it's sent or with update
		Caption:   "New caption",
		ReplyMarkup: tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("New button"),
			),
		),
	})

	// Edit message photo
	_, _ = bot.EditMessageMedia(&telego.EditMessageMediaParams{
		ChatID:    tu.ID(1234567),
		MessageID: 1234, // Message ID can be retried when it's sent or with update
		Media:     tu.MediaPhoto(tu.File(mustOpen("photo.png"))),
	})
}

// Helper function to open file or panic
func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
