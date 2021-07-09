package main

import (
	"fmt"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

const testToken = "950209960:AAEXV03s6bW5C1O138ydeW8fnYxeG_CcGl4" //nolint:gosec

func main() {
	bot, err := telego.NewBot(testToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bot.GetMe())
}
