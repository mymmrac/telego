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

	// Set up a webhook on Telegram side
	_ = bot.SetWebhook(&telego.SetWebhookParams{
		URL: "https://example.com/bot" + bot.Token(),
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo()
	fmt.Printf("Webhook Info: %#v\n", info)

	// Get updates channel from webhook.
	// Note: For one bot only one webhook allowed.
	updates, _ := bot.GetUpdatesViaWebhook("/bot" + bot.Token())

	// Start server for receiving requests from Telegram
	bot.StartListeningForWebhook("localhost:443")

	// Stop reviving updates from updates channel and shutdown webhook server
	defer func() {
		_ = bot.StopWebhook()
	}()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %#v\n", update)
	}
}
