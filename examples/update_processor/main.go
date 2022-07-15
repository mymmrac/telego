package main

import (
	"fmt"
	"os"
	"sync/atomic"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Create Bot
	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPulling(nil)

	// Stop reviving updates from updates channel
	defer bot.StopLongPulling()

	fmt.Println("Listening for updates...")

	count := int64(0)

	// Process updates for something (here to count them)
	processedUpdates := tu.UpdateProcessor(updates, 100, func(update telego.Update) telego.Update {
		atomic.AddInt64(&count, 1)

		currentCount := atomic.LoadInt64(&count)
		fmt.Println("Update count:", currentCount)

		// Stop bot when processed 3 updates
		if currentCount >= 3 {
			bot.StopLongPulling()
		}

		return update
	})

	// Log update IDs
	for update := range processedUpdates {
		fmt.Println("Update ID:", update.UpdateID)
	}

	fmt.Println("Bye")
}
