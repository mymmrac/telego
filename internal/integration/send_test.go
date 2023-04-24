//go:build integration

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

	t.Run("unicode-entities", func(t *testing.T) {
		text, entities := tu.MessageEntities(
			tu.Entity("😅").Italic(),
			tu.Entity(" test ").Bold(),
			tu.Entity("🌗").Italic(),
			tu.Entity(" Україна").Bold(),
			tu.Entity(" "),
			tu.Entity("\U0001FAE5 ").Italic(),
			tu.Entity("世界").Bold(),
		)
		msg, err := bot.SendMessage(tu.Message(tu.ID(chatID), text).WithEntities(entities...))
		require.NoError(t, err)

		/*

			TODO: Method is not fully compliant with API, trailing newlines or whitespaces not counted correctly
			https://core.telegram.org/api/entities

			actual  : []telego.MessageEntity{
				telego.MessageEntity{Type:"italic", Offset:0, Length:2, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""},
				telego.MessageEntity{Type:"bold", Offset:3, Length:5, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""},
				telego.MessageEntity{Type:"italic", Offset:8, Length:2, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""},
				telego.MessageEntity{Type:"bold", Offset:11, Length:7, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""},
				telego.MessageEntity{Type:"italic", Offset:19, Length:3, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""},
				telego.MessageEntity{Type:"bold", Offset:22, Length:2, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""}
			}
			expected: []telego.MessageEntity{
				telego.MessageEntity{Type:"italic", Offset:0, Length:2, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""},
				telego.MessageEntity{Type:"bold", Offset:2, Length:6, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""},
				telego.MessageEntity{Type:"italic", Offset:8, Length:2, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""},
				telego.MessageEntity{Type:"bold", Offset:10, Length:8, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""},
				telego.MessageEntity{Type:"italic", Offset:19, Length:3, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""},
				telego.MessageEntity{Type:"bold", Offset:22, Length:2, URL:"", User:(*telego.User)(nil), Language:"", CustomEmojiID:""}
			}

		*/

		assert.Equal(t, msg.Text, text)
		assert.Equal(t, msg.Entities, entities)
	})

	t.Run("new_line", func(t *testing.T) {
		msg, err := bot.SendMessage(&telego.SendMessageParams{
			ChatID: tu.ID(chatID),
			Text:   "Send\nMessage",
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})
}

func TestSendPhoto(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		msg, err := bot.SendPhoto(&telego.SendPhotoParams{
			ChatID:  tu.ID(chatID),
			Photo:   tu.File(open(img1Jpg)),
			Caption: "SendPhoto " + timeNow,
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("new_line", func(t *testing.T) {
		msg, err := bot.SendPhoto(&telego.SendPhotoParams{
			ChatID:  tu.ID(chatID),
			Photo:   tu.File(open(img1Jpg)),
			Caption: "Send\nPhoto \" >",
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("keyboard_and_markdown", func(t *testing.T) {
		msg, err := bot.SendPhoto(&telego.SendPhotoParams{
			ChatID:    tu.ID(chatID),
			Photo:     tu.File(open(img1Jpg)),
			ParseMode: telego.ModeMarkdownV2,
			Caption:   "Send\n`Photo`",
			ReplyMarkup: tu.InlineKeyboard(tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Test").WithCallbackData("ok"),
			)),
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})
}

func TestSendAudio(t *testing.T) {
	t.Run("audio_file", func(t *testing.T) {
		msg, err := bot.SendAudio(&telego.SendAudioParams{
			ChatID:    tu.ID(chatID),
			Audio:     tu.File(open(kittenMp3)),
			Caption:   "SendAudio " + timeNow,
			Thumbnail: telego.ToPtr(tu.File(open(img1Jpg))),
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("url", func(t *testing.T) {
		msg, err := bot.SendAudio(&telego.SendAudioParams{
			ChatID:    tu.ID(chatID),
			Audio:     tu.FileFromURL("https://file-examples.com/storage/fe0e4ffeec64469f8a2ba23/2017/11/file_example_MP3_700KB.mp3"),
			Caption:   "SendAudio " + timeNow,
			Thumbnail: telego.ToPtr(tu.File(open(img1Jpg))), // Expected to be not displayed
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})
}
