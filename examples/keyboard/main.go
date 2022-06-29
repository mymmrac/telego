package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger exposes your bot token, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Keyboard parameters
	keyboard := tu.Keyboard(
		tu.KeyboardRow( // Row 1
			tu.KeyboardButton("Button 1"), // Column 1
			tu.KeyboardButton("Button 2"), // Column 2
		),
		tu.KeyboardRow( // Row 2
			tu.KeyboardButton("Contact").WithRequestContact(),   // Column 1
			tu.KeyboardButton("Location").WithRequestLocation(), // Column 2
		),
		tu.KeyboardRow( // Row 3
			tu.KeyboardButton("Poll Any").WithRequestPoll(tu.PollTypeAny()),         // Column 1
			tu.KeyboardButton("Poll Regular").WithRequestPoll(tu.PollTypeRegular()), // Column 2
			tu.KeyboardButton("Poll Quiz").WithRequestPoll(tu.PollTypeQuiz()),       // Column 3
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")

	// Message parameters
	message := tu.Message(
		tu.ID(1234567),
		"My message",
	).WithReplyMarkup(keyboard)

	// Sending message
	_, _ = bot.SendMessage(message)
}
