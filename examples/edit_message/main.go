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
	updates, _ := bot.UpdatesViaLongPolling(nil)
	// Stop reviving updates from update channel
	defer bot.StopLongPolling()
	// Loop through all updates when they came
	for update := range updates {
		// Check if update contains a message
		if update.Message != nil {
			message, err := bot.SendMessage(&telego.SendMessageParams{
				ChatID: update.Message.Chat.ChatID(),
				Text:   "hello",
			})
			if err != nil {
				return
			}

			// Edit message text
			_, _ = bot.EditMessageText(&telego.EditMessageTextParams{
				ChatID:    message.Chat.ChatID(),
				MessageID: message.GetMessageID(), // Message ID can be retried when it's sent or with update
				Text:      "<b>Bold</b>",
				ParseMode: telego.ModeHTML,
			})
			// Send a photo with a caption
			mediaMessage, _ := bot.SendPhoto(&telego.SendPhotoParams{
				ChatID:    update.Message.Chat.ChatID(),
				Photo:     tu.FileFromURL("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSmRtCXZnOIjV-MQkeAjjoPd2_SHPmTzJur2A&s"), // Replace with your photo URL or file path
				Caption:   "This is a caption",                                                                                            // Initial caption
				ParseMode: telego.ModeHTML,
			})
			// Edit message caption & reply markup (inline keyboard)
			_, _ = bot.EditMessageCaption(&telego.EditMessageCaptionParams{
				ChatID:    mediaMessage.Chat.ChatID(),
				MessageID: mediaMessage.GetMessageID(), // Message ID can be retried when it's sent or with update
				ReplyMarkup: tu.InlineKeyboard(
					tu.InlineKeyboardRow(
						tu.InlineKeyboardButton("New button").WithCallbackData("test"),
					),
				),
			})
			// Edit message photo
			_, _ = bot.EditMessageMedia(&telego.EditMessageMediaParams{
				ChatID:    mediaMessage.Chat.ChatID(),
				MessageID: mediaMessage.GetMessageID(), // Message ID can be retried when it's sent or with update
				Media:     tu.MediaPhoto(tu.File(mustOpen("photo.png"))),
			})

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
