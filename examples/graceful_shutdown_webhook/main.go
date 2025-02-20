package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/valyala/fasthttp"
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

	// Create HTTP server
	srv := &fasthttp.Server{}

	// Get updates
	updates, _ := bot.UpdatesViaWebhook(ctx, telego.WebhookFastHTTP(srv, "/bot", bot.Token()))

	// Create bot handler
	bh, _ := th.NewBotHandler(bot, updates)

	// Handle updates
	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		fmt.Println("Processing update:", update.UpdateID)
		time.Sleep(time.Second * 5) // Simulate long process time
		fmt.Println("Done update:", update.UpdateID)
		return nil
	})

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

		_ = bh.StopWithContext(stopCtx)
		fmt.Println("Bot handler done")

		// Notify that stop is done
		done <- struct{}{}
	}()

	// Start handling in goroutine
	go func() { _ = bh.Start() }()
	fmt.Println("Handling updates...")

	// Start server for receiving requests from the Telegram
	go func() { _ = srv.ListenAndServe(":443") }()

	// Wait for the stop process to be completed
	<-done
	fmt.Println("Done")
}
