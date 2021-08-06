package main

import (
	"fmt"
	"os"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.DebugMode(true)

	// Document parameters
	document := &telego.SendDocumentParams{
		ChatID:  telego.ChatID{ID: 1234567}, // Chat ID as Integer
		Caption: "My cool file from disk",

		// Send using file from disk
		Document: telego.InputFile{File: mustOpen("my_file.txt")},

		// Send using external URL
		//Document: telego.InputFile{URL: "https://example.com/my_file.txt"},

		// Send using file ID
		//Document: telego.InputFile{FileID: "<file ID of your file>"},
	}

	// Sending document
	msg, err := bot.SendDocument(document)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg.Document)

	// =========================================== //

	// Photo parameters
	photo := &telego.SendPhotoParams{
		ChatID:  telego.ChatID{Username: "@my_cool_channel"}, // Chat ID as String (target username)
		Caption: "My cool photo",

		// Send using file from disk
		Photo: telego.InputFile{File: mustOpen("my_photo.png")},
	}
	// Sending photo
	_, _ = bot.SendPhoto(photo)

	// =========================================== //

	// Media group parameters
	mediaGroup := &telego.SendMediaGroupParams{
		ChatID: telego.ChatID{ID: 1234567},

		// Specify slice of telego.InputMedia with media you want to send as a group
		Media: []telego.InputMedia{
			&telego.InputMediaPhoto{
				Type:  telego.MediaTypePhoto,
				Media: telego.InputFile{URL: "https://example.com/my_photo.png"},
			},
			&telego.InputMediaPhoto{
				Type:  telego.MediaTypePhoto,
				Media: telego.InputFile{File: mustOpen("my_photo.png")},
			},
			&telego.InputMediaPhoto{
				Type:  telego.MediaTypePhoto,
				Media: telego.InputFile{FileID: "<file ID of your photo>"},
			},
		},
	}

	// Sending media group
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
