package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/valyala/fasthttp"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"

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

	// Initialize signal handling
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Initialize done chan
	done := make(chan struct{}, 1)

	// Create a new Ngrok tunnel to connect local network with the Internet & have HTTPS domain for bot
	tun, err := ngrok.Listen(context.Background(),
		// Forward connections to localhost:8080
		config.HTTPEndpoint(config.WithForwardsTo(":8080")),
		// Authenticate into Ngrok using NGROK_AUTHTOKEN env (optional)
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Prepare fast HTTP server
	srv := &fasthttp.Server{}

	// Get an update channel from webhook using Ngrok
	updates, _ := bot.UpdatesViaWebhook(ctx,
		// Use FastHTTP webhook server
		telego.WebhookFastHTTP(srv, "/bot", bot.SecretToken()),
		// Calls SetWebhook before starting webhook and provide dynamic Ngrok tunnel URL
		telego.WithWebhookSet(ctx, &telego.SetWebhookParams{
			URL:         tun.URL() + "/bot",
			SecretToken: bot.SecretToken(),
		}),
	)

	// Handle stop signal (Ctrl+C)
	go func() {
		// Wait for stop signal
		<-ctx.Done()
		fmt.Println("Stopping...")

		stopCtx, stopCancel := context.WithTimeout(context.Background(), time.Second*20)
		defer stopCancel()

		_ = srv.ShutdownWithContext(stopCtx)
		fmt.Println("Server done")

		for len(updates) > 0 {
			select {
			case <-stopCtx.Done():
				break
			case <-time.After(time.Microsecond * 100):
				// Continue
			}
		}
		fmt.Println("Webhook done")

		// Notify that stop is done
		done <- struct{}{}
	}()

	// Start server for receiving requests from the Telegram using the Ngrok tunnel
	go func() { _ = srv.Serve(tun) }()

	// Loop through all updates when they came
	go func() {
		for update := range updates {
			fmt.Printf("Update: %+v\n", update)
		}
	}()

	// Wait for the stop process to be completed
	<-done
	fmt.Println("Done")
}
