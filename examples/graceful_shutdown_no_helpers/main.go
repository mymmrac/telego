package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/chococola/telego"
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

	// Initialize updates processor chan
	updatesProcessed := make(chan struct{}, 1)

	// Handle stop signal (Ctrl+C)
	go func() {
		// Wait for stop signal
		<-sigs

		fmt.Println("Stopping...")

		bot.StopLongPolling()
		fmt.Println("Long polling done")

		<-updatesProcessed
		fmt.Println("Updates processed")

		// Notify that stop is done
		done <- struct{}{}
	}()

	// Get updates
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Handle updates
	go func() {
		for update := range updates {
			fmt.Println("Processing update:", update.UpdateID)
			time.Sleep(time.Second * 5) // Simulate long process time
			fmt.Println("Done update:", update.UpdateID)
		}
		updatesProcessed <- struct{}{}
	}()

	// Wait for the stop process to be completed
	<-done
	fmt.Println("Done")
}
