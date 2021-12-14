package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tg "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Keyboard parameters
	keyboard := &telego.ReplyKeyboardMarkup{
		Keyboard: [][]telego.KeyboardButton{
			// Row 1
			{
				// Column 1
				{
					Text: "Button 1",
				},
				// Column 2
				{
					Text: "Button 2",
				},
			},
			// Row 2
			{
				// Column 1
				{
					Text:           "Contact",
					RequestContact: true,
				},
				// Column 2
				{
					Text:            "Location",
					RequestLocation: true,
				},
			},
			// Row 3
			{
				// Column 1
				{
					Text:        "Poll",
					RequestPoll: &telego.KeyboardButtonPollType{},
				},
				// Column 2
				{
					Text:        "Poll Regular",
					RequestPoll: &telego.KeyboardButtonPollType{Type: telego.PollTypeRegular},
				},
				// Column 3
				{
					Text:        "Poll Quiz",
					RequestPoll: &telego.KeyboardButtonPollType{Type: telego.PollTypeQuiz},
				},
			},
		},
		ResizeKeyboard:        true,
		InputFieldPlaceholder: "Select something",
	}

	// Message parameters
	message := &telego.SendMessageParams{
		ChatID:      tg.ID(1234567),
		Text:        "My message",
		ReplyMarkup: keyboard,
	}

	// Sending message
	_, _ = bot.SendMessage(message)
}
