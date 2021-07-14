package telego

import "encoding/json"

// ChatID - Represents chat ID as int or string
type ChatID struct {
	ID       int64
	Username string
}

func (c ChatID) MarshalJSON() ([]byte, error) {
	if c.Username != "" {
		return json.Marshal(struct {
			ChatID string `json:"chat_id"`
		}{ChatID: c.Username})
	}

	return json.Marshal(struct {
		ChatID int64 `json:"chat_id"`
	}{ChatID: c.ID})
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
