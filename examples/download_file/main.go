package main

import (
	"fmt"
	tu "github.com/mymmrac/telego/telegoutil"
	"os"

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
	updates, _ := bot.UpdatesViaLongPolling(nil)
	// Stop reviving updates from update channel
	defer bot.StopLongPolling()

	// Loop through all updates when they came
	for update := range updates {
		// Check if update contains a message
		if update.Message != nil {
			// Document parameters
			document := tu.Document(
				update.Message.Chat.ChatID(),
				tu.File(mustOpen("my_file.txt")),
			)

			// Send a test document
			msg, _ := bot.SendDocument(document)

			// Get file info
			// Note: File ID used to get info is only valid for temporary time
			file, _ := bot.GetFile(&telego.GetFileParams{
				FileID: msg.Document.FileID,
			})

			// Download file from Telegram using FileDownloadURL helper func to get full URL
			fileData, err := tu.DownloadFile(bot.FileDownloadURL(file.FilePath))
			fmt.Println(len(fileData), err)
		}
	}

}

// Helper function to open file or panic
func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
