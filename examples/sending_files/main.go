package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
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
		// Chat ID as Integer
		tu.ID(1234567),

		// Send using file from disk
		tu.File(mustOpen("my_file.txt")),

		// Send using external URL
		// tu.FileByURL("https://example.com/my_file.txt"),

		// Send using file ID
		// tu.FileByID("<file ID of your file>"),
	).WithCaption("My cool file from disk")

	// Sending document
	msg, err := bot.SendDocument(document)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg.Document)

	// =========================================== //

	// Photo parameters
	photo := tu.Photo(
		// Chat ID as String (target username)
		tu.Username("@my_cool_channel"),

		// Send using file from disk
		tu.File(mustOpen("my_photo.png")),
	).WithCaption("My cool photo")

	// Sending photo
	_, _ = bot.SendPhoto(photo)

	// =========================================== //

	// Media group parameters
	mediaGroup := tu.MediaGroup(
		tu.ID(1234567),

		// Specify slice of telego.InputMedia with media you want to send as a group
		tu.MediaPhoto(tu.File(mustOpen("my_photo.png"))),

		tu.MediaPhoto(tu.FileByID("<file ID of your photo>")),

		tu.MediaPhoto(tu.FileByURL("https://example.com/my_photo.png")),
	)

	// Sending a media group
	_, _ = bot.SendMediaGroup(mediaGroup)
}

// Helper function to open file or panic
func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
