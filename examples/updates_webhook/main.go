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
	// Note: If tou are keeping this call, you shouldn't use telego.WithWebhookSet option below
	_ = bot.SetWebhook(&telego.SetWebhookParams{
		URL: "https://example.com/bot" + bot.Token(),
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo()
	fmt.Printf("Webhook Info: %+v\n", info)

	// Get an update channel from webhook, also all options are optional.
	// Note: For one bot, only one webhook allowed.
	updates, _ := bot.UpdatesViaWebhook("/bot"+bot.Token(),
		// Set chan buffer (default 128)
		telego.WithWebhookBuffer(128),

		// Set fast http server that will be used to handle webhooks (default telego.FastHTTPWebhookServer)
		// Note: If SecretToken is non-empty, it will be verified on each request
		telego.WithWebhookServer(telego.FastHTTPWebhookServer{
			Logger:      bot.Logger(),
			Server:      &fasthttp.Server{},
			Router:      router.New(),
			SecretToken: "token",
		}),

		// Calls SetWebhook before starting webhook
		// Note: If tou are using this option, you shouldn't call bot.SetWebhook above
		telego.WithWebhookSet(&telego.SetWebhookParams{
			URL: "https://example.com/bot" + bot.Token(),
		}),
	)

	// Start server for receiving requests from the Telegram
	_ = bot.StartWebhook("localhost:443")

	// Stop reviving updates from update channel and shutdown webhook server
	defer func() {
		_ = bot.StopWebhook()
	}()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %+v\n", update)
	}
}
