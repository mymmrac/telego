package telegoutil

import (
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoapi"
)

// ID creates telego.ChatID from user's identifier
func ID(id int64) telego.ChatID {
	return telego.ChatID{ID: id}
}

// Username creates telego.ChatID from username
func Username(username string) telego.ChatID {
	return telego.ChatID{Username: username}
}

// File creates telego.InputFile from telegoapi.NamedReader
func File(file telegoapi.NamedReader) telego.InputFile {
	return telego.InputFile{
		File: file,
	}
}

// FileByURL creates telego.InputFile from URL
func FileByURL(url string) telego.InputFile {
	return telego.InputFile{
		URL: url,
	}
}

// FileByID creates telego.InputFile from file ID
func FileByID(id string) telego.InputFile {
	return telego.InputFile{
		FileID: id,
	}
}

// Keyboard creates telego.ReplyKeyboardMarkup from slice of keyboard buttons
func Keyboard(rows ...[]telego.KeyboardButton) *telego.ReplyKeyboardMarkup {
	return &telego.ReplyKeyboardMarkup{
		Keyboard: rows,
	}
}

// KeyboardRow creates slice of telego.KeyboardButton
func KeyboardRow(buttons ...telego.KeyboardButton) []telego.KeyboardButton {
	return buttons
}

// KeyboardButton creates telego.KeyboardButton with required fields
func KeyboardButton(text string) telego.KeyboardButton {
	return telego.KeyboardButton{
		Text: text,
	}
}

// PollTypeAny creates telego.KeyboardButtonPollType with any type
func PollTypeAny() *telego.KeyboardButtonPollType {
	return &telego.KeyboardButtonPollType{}
}

// PollTypeRegular creates telego.KeyboardButtonPollType with type regular
func PollTypeRegular() *telego.KeyboardButtonPollType {
	return &telego.KeyboardButtonPollType{Type: telego.PollTypeRegular}
}

// PollTypeQuiz creates telego.KeyboardButtonPollType with type quiz
func PollTypeQuiz() *telego.KeyboardButtonPollType {
	return &telego.KeyboardButtonPollType{Type: telego.PollTypeQuiz}
}
