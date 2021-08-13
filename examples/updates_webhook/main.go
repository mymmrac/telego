package main

import (
	"fmt"
	"os"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.DefaultLogger(true, true)

	// Setup a webhook
	_ = bot.SetWebhook(&telego.SetWebhookParams{
		URL:         "https://www.google.com:443/" + bot.Token(),
		Certificate: &telego.InputFile{File: mustOpen("cert.pem")},
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo()
	fmt.Printf("Webhook Info: %#v\n", info)

	// Start server for receiving requests from telegram
	bot.StartListeningForWebhook("0.0.0.0:443/"+bot.Token(), "cert.pem", "key.pem")

	// Get updates channel from webhook. Note for one bot only one webhook allowed
	updates, _ := bot.ListenForWebhook("/" + bot.Token())

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %#v\n", update)
	}
}

// Helper function to open file or panic
func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
