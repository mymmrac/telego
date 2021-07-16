package telego

import (
	"encoding/json"
	"fmt"
	"os"
)

// ChatID - Represents chat ID as int or string
type ChatID struct {
	ID       int64
	Username string
}

func (c ChatID) String() string {
	if c.Username != "" {
		return c.Username
	}

	return fmt.Sprintf("%d", c.ID)
}

func (c ChatID) MarshalJSON() ([]byte, error) {
	if c.Username != "" {
		return json.Marshal(c.Username)
	}

	return json.Marshal(fmt.Sprintf("%d", c.ID))
}

// ReplyMarkup - Represents reply markup (inline keyboard, custom reply keyboard, etc.)
type ReplyMarkup interface {
	ReplyType() string
}

// ReplyType - Returns InlineKeyboardMarkup type
func (i *InlineKeyboardMarkup) ReplyType() string {
	return "InlineKeyboardMarkup"
}

// ReplyType - Returns ReplyKeyboardMarkup type
func (i *ReplyKeyboardMarkup) ReplyType() string {
	return "ReplyKeyboardMarkup"
}

// ReplyType - Returns ReplyKeyboardRemove type
func (i *ReplyKeyboardRemove) ReplyType() string {
	return "ReplyKeyboardRemove"
}

// ReplyType - Returns ForceReply type
func (i *ForceReply) ReplyType() string {
	return "ForceReply"
}

func (i InputFile) MarshalJSON() ([]byte, error) {
	if i.File != nil {
		return json.Marshal("")
	}

	if i.FileID != "" {
		return json.Marshal(i.FileID)
	}

	return json.Marshal(i.URL)
}

type fileCompatible interface {
	fileParameters() map[string]*os.File
}

func (p *SendDocumentParams) fileParameters() map[string]*os.File {
	fp := make(map[string]*os.File)

	fp["document"] = p.Document.File
	if p.Thumb != nil {
		fp["thumb"] = p.Thumb.File
	}

	return fp
}
