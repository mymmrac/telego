package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger exposes your bot token, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create telego.ChatID from int64
	chatID := tu.ID(123)

	// Create message with only required parameters
	_, _ = bot.SendMessage(tu.Message(chatID, "Hello"))

	// Create telego.ChatID username and send message
	_, _ = bot.SendMessage(tu.Message(tu.Username("@user"), "World"))

	// Create message and change optional parameters
	msg := tu.Message(chatID, "Hello World").
		WithReplyToMessageID(1234).
		WithDisableNotification().
		WithProtectContent()
	_, _ = bot.SendMessage(msg)

	var file *os.File // Used just for example (not valid in real use)

	// Create document using *os.File as telego.InputFile
	_, _ = bot.SendDocument(tu.Document(chatID, tu.File(file)))

	var reader io.Reader // Used just for example (not valid in real use)

	// Create document using io.Reader by "naming" it and send as document
	_, _ = bot.SendDocument(tu.Document(chatID, tu.File(tu.NameReader(reader, "my_file"))))

	// Create document using URL to file as telego.InputFile
	_, _ = bot.SendDocument(tu.Document(chatID, tu.FileByURL("https://example.com/test.txt")))

	// Create contact from phone and first name
	_, _ = bot.SendContact(tu.Contact(chatID, "+123454321", "John"))

	// Small example of parsing commands
	updates, _ := bot.GetUpdates(nil)
	for _, u := range updates {
		if u.Message != nil {
			text := u.Message.Text

			// Parse text into command and its arguments
			command, args := tu.ParseCommand(text)

			// Check if text contains command
			if command != "" {
				fmt.Println("Command:", command)
				fmt.Println("Args:", args)
			} else {
				fmt.Println("Not a command:", text)
			}
		}
	}
}
