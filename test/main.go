package main

import (
	"fmt"
	telego "github.com/mymmrac/go-telegram-bot-api"
	"os"
)

const testToken = "950209960:AAEXV03s6bW5C1O138ydeW8fnYxeG_CcGl4" //nolint:gosec

func main() {
	bot, err := telego.NewBot(testToken)
	if err != nil {
		panic(err)
	}

	fmt.Println(bot.GetMe())

	file, err := os.Open("doc.txt")
	if err != nil {
		panic(err)
	}

	dp := &telego.SendDocumentParams{
		ChatID:   telego.ChatID{ID: 331849104},
		Document: telego.InputFile{File: file},
		Caption:  "Hello world",
	}
	msg, err := bot.SendDocument(dp)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg.Document)
}
