package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	"github.com/valyala/fasthttp"
)

func main() {
	ctx := context.Background()
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set up a webhook on Telegram side
	// Note: If you are keeping this call, you shouldn't use telego.WithWebhookSet option below
	_ = bot.SetWebhook(ctx, &telego.SetWebhookParams{
		URL:         "https://example.com/bot",
		SecretToken: bot.Token(),
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo(ctx)
	fmt.Printf("Webhook Info: %+v\n", info)

	// Create HTTP server
	srv := &fasthttp.Server{}

	// Get an update channel from webhook, also all options are optional.
	// Note: For one bot, only one webhook allowed.
	updates, _ := bot.UpdatesViaWebhook(ctx,
		// Use fasthttp webhook server, any other custom server can be used,
		// you need to provide a way to register handler
		telego.WebhookFastHTTP(srv, "/bot", bot.Token()),

		// Telego provides a simple way to use http.Server as a webhook server
		// telego.WebhookHTTPServer(srv, "/bot", bot.Token()),

		// Telego provides a simple way to use http.ServeMux as a webhook server
		// telego.WebhookHTTPServeMux(mux, "POST /bot", bot.Token()),

		// Set chan buffer (default 128)
		telego.WithWebhookBuffer(128),

		// Calls SetWebhook before starting webhook
		// Note: If you are using this option, you shouldn't call bot.SetWebhook above
		telego.WithWebhookSet(ctx, &telego.SetWebhookParams{
			URL:         "https://example.com/bot",
			SecretToken: bot.Token(),
		}),
	)

	// Start server for receiving requests from the Telegram
	go func() { _ = srv.ListenAndServe(":433") }()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %+v\n", update)
	}
}
