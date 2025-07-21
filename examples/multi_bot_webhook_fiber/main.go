package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"

	"github.com/mymmrac/telego"
)

func main() {
	ctx := context.Background()
	botToken1 := os.Getenv("TOKEN_1")
	botToken2 := os.Getenv("TOKEN_2")

	// Create multiple bots with different tokens
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot1, _ := telego.NewBot(botToken1, telego.WithDefaultDebugLogger())
	bot2, _ := telego.NewBot(botToken2, telego.WithDefaultDebugLogger())

	// Set up a webhook on Telegram side for each bot with different URLs
	_ = bot1.SetWebhook(ctx, &telego.SetWebhookParams{
		URL:         "https://example.com/bot1",
		SecretToken: bot1.SecretToken(),
	})
	_ = bot2.SetWebhook(ctx, &telego.SetWebhookParams{
		URL:         "https://example.com/bot2",
		SecretToken: bot2.SecretToken(),
	})

	// Create a common webhook serve mux (or another custom server) for all bots
	app := fiber.New()

	// Get updates chan from webhook with respect to webhook URL
	// Note: Each bot should use the same webhook serve mux (or another custom server)
	updates1, _ := bot1.UpdatesViaWebhook(ctx, WebhookFiber(app, "/bot1", bot1.SecretToken()))
	updates2, _ := bot2.UpdatesViaWebhook(ctx, WebhookFiber(app, "/bot2", bot2.SecretToken()))

	// Start server for receiving requests from the Telegram
	go func() { _ = app.Listen(":443") }()

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

// WebhookFiber returns a fiber handler for update webhook
func WebhookFiber(router fiber.Router, path string, secretToken string) func(handler telego.WebhookHandler) error {
	return func(handler telego.WebhookHandler) error {
		router.Post(path, func(fCtx fiber.Ctx) error {
			if secretToken != string(fCtx.Request().Header.Peek(telego.WebhookSecretTokenHeader)) {
				return fCtx.SendStatus(fiber.StatusUnauthorized)
			}

			if err := handler(fCtx, fCtx.Body()); err != nil {
				return fCtx.SendStatus(fiber.StatusInternalServerError)
			}

			return fCtx.SendStatus(fiber.StatusOK)
		})
		return nil
	}
}
