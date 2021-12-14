package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tg "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Document parameters
	document := &telego.SendDocumentParams{
		ChatID:  tg.ID(1234567), // Chat ID as Integer
		Caption: "My cool file from disk",

		// Send using file from disk
		Document: tg.File(mustOpen("my_file.txt")),

		// Send using external URL
		//Document: tg.FileByURL("https://example.com/my_file.txt"),

		// Send using file ID
		//Document: tg.FileByID("<file ID of your file>"),
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
		ChatID:  tg.Username("@my_cool_channel"), // Chat ID as String (target username)
		Caption: "My cool photo",

		// Send using file from disk
		Photo: tg.File(mustOpen("my_photo.png")),
	}
	// Sending photo
	_, _ = bot.SendPhoto(photo)

	// =========================================== //

	// Media group parameters
	mediaGroup := &telego.SendMediaGroupParams{
		ChatID: tg.ID(1234567),

		// Specify slice of telego.InputMedia with media you want to send as a group
		Media: []telego.InputMedia{
			&telego.InputMediaPhoto{
				Type:  telego.MediaTypePhoto,
				Media: tg.FileByURL("https://example.com/my_photo.png"),
			},
			&telego.InputMediaPhoto{
				Type:  telego.MediaTypePhoto,
				Media: tg.File(mustOpen("my_photo.png")),
			},
			&telego.InputMediaPhoto{
				Type:  telego.MediaTypePhoto,
				Media: tg.FileByID("<file ID of your photo>"),
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
