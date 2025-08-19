package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatalf("Create bot: %s", err)
	}

	// Initialize signal handling
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Generate secret
	secretBytes := sha256.Sum256([]byte(botToken))
	secret := hex.EncodeToString(secretBytes[:])

	// Create webhook server and listener
	srv := &fasthttp.Server{}
	listener, url := WebhookListener(ctx)

	// Get updates via webhook
	updates, err := bot.UpdatesViaWebhook(ctx,
		telego.WebhookFastHTTP(srv, "/bot", secret),
		telego.WithWebhookSet(ctx, tu.Webhook(url+"/bot").WithSecretToken(secret)),
	)
	if err != nil {
		log.Fatalf("Updates via webhoo: %s", err)
	}

	// Create bot handler
	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		log.Fatalf("Create bot handler: %s", err)
	}

	// Register bot handlers
	registerHandlers(bh)

	// Initialize done chan
	done := make(chan struct{}, 1)

	go func() {
		// Wait for stop signal
		<-ctx.Done()
		log.Println("Stopping...")

		stopCtx, stopCancel := context.WithTimeout(context.Background(), time.Second*10)
		defer stopCancel()

		if err = srv.ShutdownWithContext(stopCtx); err != nil {
			log.Printf("Failed to shutdown webhook server: %s", err)
		}

	loop:
		for len(updates) > 0 {
			select {
			case <-stopCtx.Done():
				break loop
			case <-time.After(time.Microsecond * 100):
				log.Printf("Update handler timeout")
			}
		}

		if err = bh.StopWithContext(stopCtx); err != nil {
			log.Printf("Failed to stop bot handler: %s", err)
		}

		// Notify that stop is done
		done <- struct{}{}
	}()

	// Start handling in goroutine
	go func() {
		if startErr := bh.Start(); startErr != nil {
			log.Fatalf("Failed to start bot handler: %s", startErr)
		}
	}()
	log.Println("Handling updates...")

	// Start server for receiving requests from the Telegram
	go func() {
		err = srv.Serve(listener)
		if err != nil {
			log.Fatalf("Failed to start webhook server: %s", err)
		}
	}()

	// Wait for the stop process to be completed
	<-done
	log.Println("Done")
}
