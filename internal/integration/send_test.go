package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func TestSendMessage(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		msg, err := bot.SendMessage(&telego.SendMessageParams{
			ChatID: tu.ID(chatID),
			Text:   "SendMessage " + timeNow,
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("complex", func(t *testing.T) {
		keyboard := tu.InlineKeyboard(
			tu.InlineKeyboardRow(tu.InlineKeyboardButton("Test").WithCallbackData("OK")),
		)
		msg, err := bot.SendMessage(
			tu.MessageWithEntities(tu.ID(chatID),
				tu.Entity("SendMessage").Bold(), tu.Entity(" "), tu.Entity(timeNow).Code(),
			).WithReplyMarkup(keyboard),
		)

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})
}

func TestSendPhoto(t *testing.T) {
	msg, err := bot.SendPhoto(&telego.SendPhotoParams{
		ChatID:  tu.ID(chatID),
		Photo:   tu.File(open(img1Jpg)),
		Caption: "SendPhoto " + timeNow,
	})

	require.NoError(t, err)
	assert.NotNil(t, msg)
}

func TestSendAudio(t *testing.T) {
	msg, err := bot.SendAudio(&telego.SendAudioParams{
		ChatID:  tu.ID(chatID),
		Audio:   tu.File(open(kittenMp3)),
		Caption: "SendAudio " + timeNow,
	})

	require.NoError(t, err)
	assert.NotNil(t, msg)
}
