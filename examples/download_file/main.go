package main

import (
	"fmt"
	"os"

	"github.com/chococola/telego"
	tu "github.com/chococola/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Document parameters
	document := tu.Document(
		tu.ID(1234567),
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

// Helper function to open file or panic
func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
