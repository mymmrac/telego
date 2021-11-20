package telegoutil

import (
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoapi"
)

// ChatID creates telego.ChatID from user ID
func ChatID(id int64) telego.ChatID {
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
