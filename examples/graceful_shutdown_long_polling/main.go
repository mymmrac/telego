package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/chococola/telego"
	th "github.com/chococola/telego/telegohandler"
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
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	// Initialize done chan
	done := make(chan struct{}, 1)

	// Get updates
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler with stop timeout
	bh, _ := th.NewBotHandler(bot, updates)

	// Handle updates
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		fmt.Println("Processing update:", update.UpdateID)
		time.Sleep(time.Second * 15) // Simulate long process time
		fmt.Println("Done update:", update.UpdateID)
	})

	// Handle stop signal (Ctrl+C)
	go func() {
		// Wait for stop signal
		<-sigs

		fmt.Println("Stopping...")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		bot.StopLongPolling()
		fmt.Println("Long polling done")

		bh.StopWithContext(ctx)
		fmt.Println("Bot handler done")

		// Notify that stop is done
		done <- struct{}{}
	}()

	// Start handling in goroutine
	go bh.Start()
	fmt.Println("Handling updates...")

	// Wait for the stop process to be completed
	<-done
	fmt.Println("Done")
}
