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
	ctx := t.Context()

	t.Run("simple", func(t *testing.T) {
		msg, err := bot.SendMessage(ctx, &telego.SendMessageParams{
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
		msg, err := bot.SendMessage(ctx,
			tu.MessageWithEntities(tu.ID(chatID),
				tu.Entity("SendMessage").Bold(), tu.Entity(" "), tu.Entity(timeNow).Code(),
			).WithReplyMarkup(keyboard),
		)

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("new_line", func(t *testing.T) {
		msg, err := bot.SendMessage(ctx, &telego.SendMessageParams{
			ChatID: tu.ID(chatID),
			Text:   "Send\nMessage",
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("code_space", func(t *testing.T) {
		text, entities := tu.MessageEntities(
			tu.Entity("    Code").Code(),
			tu.Entity("\n"),
			tu.Entity("  Pre\nPre").Pre(""),
		)
		msg, err := bot.SendMessage(ctx, tu.Message(tu.ID(chatID), text).WithEntities(entities...))
		require.NoError(t, err)

		assert.Equal(t, msg.Text, text)
		assert.Equal(t, msg.Entities, entities)
	})

	t.Run("markdown_and_entities", func(t *testing.T) {
		text, entities := tu.MessageEntities(
			tu.Entity("😅").Italic(),
			tu.Entity(" test ").Bold(),
			tu.Entity("🌗").Italic(),
			tu.Entity(" Україна").Bold(),
			tu.Entity(" "),
			tu.Entity("\U0001FAE5 ").Italic(),
			tu.Entity("世界").Bold(),
		)

		msg, err := bot.SendMessage(ctx, tu.Message(tu.ID(chatID), "_😅_* test *_🌗_* Україна* _\U0001FAE5 _*世界*").
			WithParseMode(telego.ModeMarkdownV2))
		require.NoError(t, err)

		assert.Equal(t, msg.Text, text)
		assert.Equal(t, len(msg.Entities), len(entities))

		for i := 0; i < len(entities); i++ {
			assert.Equal(t, msg.Entities[i].Type, entities[i].Type)
		}
	})

	t.Run("entities_check", func(t *testing.T) {
		msg, err := bot.SendMessage(ctx, tu.MessageWithEntities(tu.ID(chatID),
			tu.Entity("Lo").Strikethrough(), tu.Entity("rem").Underline(), tu.Entity(" ipsum "),
			tu.Entity("dolor").Strikethrough().Underline(), tu.Entity(" sit amet, consectetur adipiscing elit."),
			tu.Entity("\n"),
			tu.Entity("Praesent "), tu.Entity("sed mi blandit").Code(),
			tu.Entity(", tristique urna"), tu.Entity(" sit").TextLink("https://example.org"),
			tu.Entity(" amet,"), tu.Entity(" interdum ").Spoiler(), tu.Entity("justo."),
			tu.Entity("\n"),
			tu.Entity("\tMauris eget lobortis elit.").Pre(""),
			tu.Entity("\n"),
			tu.Entity("    Sed posuere pharetra\n justo ac commodo.").Code(),
			tu.Entity("\n"),
			tu.Entity("a  ").Code(), tu.Entity("a"),
			tu.Entity("\n"),
			tu.Entity("a  "), tu.Entity("a"),
		))

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})
}

func TestSendPhoto(t *testing.T) {
	ctx := t.Context()

	t.Run("regular", func(t *testing.T) {
		msg, err := bot.SendPhoto(ctx, &telego.SendPhotoParams{
			ChatID:  tu.ID(chatID),
			Photo:   tu.File(open(img1Jpg)),
			Caption: "SendPhoto " + timeNow,
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("new_line", func(t *testing.T) {
		msg, err := bot.SendPhoto(ctx, &telego.SendPhotoParams{
			ChatID:  tu.ID(chatID),
			Photo:   tu.File(open(img1Jpg)),
			Caption: "Send\nPhoto \" >",
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("keyboard_and_markdown", func(t *testing.T) {
		msg, err := bot.SendPhoto(ctx, &telego.SendPhotoParams{
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

	t.Run("edit_media", func(t *testing.T) {
		msg, err := bot.SendMessage(ctx, &telego.SendMessageParams{
			ChatID: tu.ID(chatID),
			Text:   "SendPhoto " + timeNow,
		})
		require.NoError(t, err)

		msg, err = bot.EditMessageMedia(ctx, tu.EditMessageMedia(
			tu.ID(chatID), msg.MessageID, tu.MediaPhoto(tu.File(open(img1Jpg))),
		))

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})
}

func TestSendAudio(t *testing.T) {
	ctx := t.Context()

	t.Run("audio_file", func(t *testing.T) {
		msg, err := bot.SendAudio(ctx, &telego.SendAudioParams{
			ChatID:    tu.ID(chatID),
			Audio:     tu.File(open(kittenMp3)),
			Caption:   "SendAudio " + timeNow,
			Thumbnail: telego.ToPtr(tu.File(open(img1Jpg))),
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("url", func(t *testing.T) {
		msg, err := bot.SendAudio(ctx, &telego.SendAudioParams{
			ChatID:    tu.ID(chatID),
			Audio:     tu.FileFromURL(exampleMp3),
			Caption:   "SendAudio " + timeNow,
			Thumbnail: telego.ToPtr(tu.File(open(img1Jpg))), // Expected to be not displayed
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})
}

func TestSendPoll(t *testing.T) {
	ctx := t.Context()

	t.Run("anonymous", func(t *testing.T) {
		msg, err := bot.SendPoll(ctx, &telego.SendPollParams{
			ChatID:      tu.ID(chatID),
			Question:    "Test",
			Options:     []telego.InputPollOption{tu.PollOption("Option 1"), tu.PollOption("Option 2")},
			IsAnonymous: nil,
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("not_anonymous", func(t *testing.T) {
		msg, err := bot.SendPoll(ctx, &telego.SendPollParams{
			ChatID:      tu.ID(chatID),
			Question:    "Test",
			Options:     []telego.InputPollOption{tu.PollOption("Option 1"), tu.PollOption("Option 2")},
			IsAnonymous: telego.ToPtr(false),
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("correct_option_id", func(t *testing.T) {
		msg, err := bot.SendPoll(ctx, &telego.SendPollParams{
			ChatID:           tu.ID(chatID),
			Question:         "Test",
			Options:          []telego.InputPollOption{tu.PollOption("Option 1"), tu.PollOption("Option 2")},
			IsAnonymous:      telego.ToPtr(false),
			Type:             telego.PollTypeQuiz,
			CorrectOptionIDs: []int{0},
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})
}

func TestSendRichMessage(t *testing.T) {
	ctx := t.Context()

	t.Run("blocks", func(t *testing.T) {
		msg, err := bot.SendRichMessage(ctx, &telego.SendRichMessageParams{
			ChatID: tu.ID(chatID),
			RichMessage: telego.InputRichMessage{
				Blocks: []telego.InputRichBlock{
					&telego.InputRichBlockSectionHeading{
						Type: telego.BlockTypeSectionHeading,
						Text: telego.ToPtr(telego.RichTextPlain("Hello")),
						Size: 1,
					},
					&telego.InputRichBlockParagraph{
						Type: telego.BlockTypeParagraph,
						Text: telego.ToPtr(telego.RichTextPlain("World")),
					},
					&telego.InputRichBlockPhoto{
						Type:  telego.BlockTypePhoto,
						Photo: *tu.MediaPhoto(tu.File(open(img1Jpg))),
						Caption: &telego.RichBlockCaption{
							Text:   telego.ToPtr(telego.RichTextPlain("Image")),
							Credit: telego.ToPtr(telego.RichTextPlain("Internet")),
						},
					},
					&telego.InputRichBlockCollage{
						Type: telego.BlockTypeCollage,
						Blocks: []telego.InputRichBlock{
							&telego.InputRichBlockPhoto{
								Type:  telego.BlockTypePhoto,
								Photo: *tu.MediaPhoto(tu.File(open(img1Jpg))),
								Caption: &telego.RichBlockCaption{
									Text: telego.ToPtr(telego.RichTextPlain("Image 1")),
								},
							},
							&telego.InputRichBlockPhoto{
								Type:  telego.BlockTypePhoto,
								Photo: *tu.MediaPhoto(tu.File(open(img2Jpg))),
								Caption: &telego.RichBlockCaption{
									Text: telego.ToPtr(telego.RichTextPlain("Image 2")),
								},
							},
						},
						Caption: &telego.RichBlockCaption{
							Text: telego.ToPtr(telego.RichTextPlain("Collage")),
						},
					},
				},
			},
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("html", func(t *testing.T) {
		msg, err := bot.SendRichMessage(ctx, &telego.SendRichMessageParams{
			ChatID: tu.ID(chatID),
			RichMessage: telego.InputRichMessage{
				HTML: `
<a name="chapter-0"></a>
<b>bold text</b>, <strong>bold text</strong>
<i>italic text</i>, <em>italic text</em>
<u>underlined text</u>, <ins>underlined text</ins>
<s>strikethrough text</s>, <strike>strikethrough text</strike>, <del>strikethrough text</del>
<code>inline fixed-width code</code>
<mark>marked text</mark>
<sub>subscript text</sub>
<sup>superscript text</sup>
<tg-spoiler>spoiler</tg-spoiler>
`,
			},
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("markdown", func(t *testing.T) {
		msg, err := bot.SendRichMessage(ctx, &telego.SendRichMessageParams{
			ChatID: tu.ID(chatID),
			RichMessage: telego.InputRichMessage{
				Markdown: `
**bold text**
__bold text__
*italic text*
_italic text_
~~strikethrough text~~
` + "`inline fixed-width code`" + `
==marked text==
||spoiler||
`,
			},
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("with_media", func(t *testing.T) {
		msg, err := bot.SendRichMessage(ctx, &telego.SendRichMessageParams{
			ChatID: tu.ID(chatID),
			RichMessage: telego.InputRichMessage{
				Markdown: `
Image 1:
![](https://telegram.org/example/photo.jpg)
Image 2:
![](tg://photo?id=ph1)
Image 3:
![](tg://photo?id=ph2)
`,
				Media: []telego.InputRichMessageMedia{
					{
						ID:    "ph1",
						Media: tu.MediaPhoto(tu.FileFromURL("https://telegram.org/example/animation.gif")),
					},
					{
						ID:    "ph2",
						Media: tu.MediaPhoto(tu.File(open(img1Jpg))),
					},
				},
			},
		})

		require.NoError(t, err)
		assert.NotNil(t, msg)
	})
}
