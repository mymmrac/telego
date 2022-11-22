package main

import (
	"fmt"
	"os"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
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
	fmt.Printf("Webhook Info: %+v\n", info)

	// Get updates channel from webhook, all options are optional.
	// Note: For one bot, only one webhook allowed.
	updates, _ := bot.UpdatesViaWebhook("/bot"+bot.Token(),
		// Set chan buffer (default 100)
		telego.WithWebhookBuffer(100),

		// Set fast http server that will be used to handle webhooks (default &fasthttp.Server{})
		telego.WithWebhookServer(&fasthttp.Server{}),

		// Set router to use, you can define your own routes (default router.New())
		telego.WithWebhookRouter(router.New()),

		// Enable default health API on `/health` (default disabled)
		// Note: Should be used only after telego.WithWebhookRouter() if any
		telego.WithWebhookHealthAPI(),
	)

	// Start server for receiving requests from the Telegram
	_ = bot.StartListeningForWebhook("localhost:443")

	// Stop reviving updates from updates channel and shutdown webhook server
	defer func() {
		_ = bot.StopWebhook()
	}()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %+v\n", update)
	}
}
