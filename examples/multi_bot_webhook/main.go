package main

import (
	"fmt"
	"os"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego"
)

func main() {
	botToken1 := os.Getenv("TOKEN_1")
	botToken2 := os.Getenv("TOKEN_2")

	// Create multiple bots with different tokens
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot1, _ := telego.NewBot(botToken1, telego.WithDefaultDebugLogger())
	bot2, _ := telego.NewBot(botToken2, telego.WithDefaultDebugLogger())

	// Set up a webhook on Telegram side for each bot with different URLs
	_ = bot1.SetWebhook(&telego.SetWebhookParams{
		URL: "https://example.com/bot" + bot1.Token(),
	})
	_ = bot2.SetWebhook(&telego.SetWebhookParams{
		URL: "https://example.com/bot" + bot2.Token(),
	})

	// Create common server and router for all bots
	srv := &telego.MultiBotWebhookServer{
		Server: telego.FastHTTPWebhookServer{
			Server: &fasthttp.Server{},
			Router: router.New(),
		},
	}

	// Get updates chan from webhook with respect to webhook URL
	// Note: Each bot should use the same server and router
	updates1, _ := bot1.UpdatesViaWebhook(
		"/bot"+bot1.Token(),
		telego.WithWebhookServer(srv),
	)
	updates2, _ := bot2.UpdatesViaWebhook(
		"/bot"+bot2.Token(),
		telego.WithWebhookServer(srv),
	)

	// Start server for receiving requests from the Telegram
	// Note: You still need to start both bot webhook servers
	go func() {
		_ = bot1.StartWebhook("localhost:443")
	}()
	go func() {
		_ = bot2.StartWebhook("localhost:443")
	}()

	// Stop reviving updates from updates chan and shutdown webhook server
	// Note: You still need to stop both bot webhook servers
	defer func() {
		_ = bot1.StopWebhook()
		_ = bot2.StopWebhook()
	}()

	// Loop through all updates when they came
	go func() {
		for update := range updates1 {
			fmt.Printf("Update 1: %+v\n", update)
		}
	}()
	for update := range updates2 {
		fmt.Printf("Update 2: %+v\n", update)
	}
}
