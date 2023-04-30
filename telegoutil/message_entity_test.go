package telegoutil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

func TestEntity(t *testing.T) {
	e := Entity(text1)
	assert.Equal(t, text1, e.text)
	assert.Equal(t, text1, e.Text())

	assert.Nil(t, e.entities)
	assert.Nil(t, e.Entities())

	testEntities := []telego.MessageEntity{
		{Type: text1},
		{Type: text2},
	}
	e.entities = testEntities
	require.Equal(t, testEntities, e.Entities())

	assert.Equal(t, 0, e.Entities()[0].Offset)
	assert.Equal(t, 0, e.Entities()[1].Offset)
	e.SetOffset(1)
	assert.Equal(t, 1, e.Entities()[0].Offset)
	assert.Equal(t, 1, e.Entities()[1].Offset)
}

func TestEntityf(t *testing.T) {
	e := Entityf("%s|%s", text1, text2)
	assert.Equal(t, text1+"|"+text2, e.text)
}

func TestMessageEntities(t *testing.T) {
	text, entities := MessageEntities(
		Entity(text1),

		Entity(text2).Italic(),
		Entity(text3).Bold(),
		Entity(text4).Strikethrough(),
		Entity(text1).Underline(),
		Entity(text2).Spoiler(),
		Entity(text3).Code(),
		Entity(text4).Pre("go"),
		Entity(text1).Hashtag(),
		Entity(text2).Cashtag(),
		Entity(text3).URL(),
		Entity(text4).BotCommand(),
		Entity(text1).Email(),
		Entity(text2).PhoneNumber(),
		Entity(text3).Mention(),

		Entity(text4).TextLink("https://example.com"),

		Entity(text1).TextMention(&telego.User{}),
		Entity(text2).TextMentionWithID(1234567),

		Entity(text3).Italic().Bold().Spoiler(),
		Entity(text4).URL().Bold(),
		Entity(text1).Spoiler().Email(),

		Entity(text2).CustomEmoji(text1),
	)

	assert.Equal(t, strings.Repeat(text1+text2+text3+text4, 5)+text1+text2, text)
	assert.Equal(t, []telego.MessageEntity{
		{Type: "italic", Offset: 4, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "bold", Offset: 9, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "strikethrough", Offset: 14, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "underline", Offset: 19, Length: 4, URL: "", User: nil, Language: ""},
		{Type: "spoiler", Offset: 23, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "code", Offset: 28, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "pre", Offset: 33, Length: 5, URL: "", User: nil, Language: "go"},
		{Type: "hashtag", Offset: 38, Length: 4, URL: "", User: nil, Language: ""},
		{Type: "cashtag", Offset: 42, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "url", Offset: 47, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "bot_command", Offset: 52, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "email", Offset: 57, Length: 4, URL: "", User: nil, Language: ""},
		{Type: "phone_number", Offset: 61, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "mention", Offset: 66, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "text_link", Offset: 71, Length: 5, URL: "https://example.com", User: nil, Language: ""},
		{Type: "text_mention", Offset: 76, Length: 4, URL: "", User: &telego.User{}, Language: ""},
		{Type: "text_mention", Offset: 80, Length: 5, URL: "", User: &telego.User{ID: 1234567}, Language: ""},
		{Type: "italic", Offset: 85, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "bold", Offset: 85, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "spoiler", Offset: 85, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "url", Offset: 90, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "bold", Offset: 90, Length: 5, URL: "", User: nil, Language: ""},
		{Type: "spoiler", Offset: 95, Length: 4, URL: "", User: nil, Language: ""},
		{Type: "email", Offset: 95, Length: 4, URL: "", User: nil, Language: ""},
		{Type: "custom_emoji", Offset: 99, Length: 5, URL: "", User: nil, Language: "", CustomEmojiID: text1},
	}, entities)

	t.Run("non_ascii", func(t *testing.T) {
		text, entities = MessageEntities(Entity(text1), Entity(textNonASCII).Bold(), Entity(text2).Bold())
		assert.Equal(t, text1+textNonASCII+text2, text)
		assert.Equal(t, []telego.MessageEntity{
			{
				Type:   "bold",
				Offset: len(text1),
				Length: 9,
			},
			{
				Type:   "bold",
				Offset: len(text1) + 9,
				Length: len(text2),
			},
		}, entities)
	})
}

func TestEntityLength(t *testing.T) {
	text, entities := MessageEntities(
		Entity("😅").Italic(),
		Entity(" test ").Bold(),
		Entity("🌗").Italic(),
		Entity(" Україна").Bold(),
		Entity(" "),
		Entity("\U0001FAE5 ").Italic(),
		Entity("世界").Bold(),
	)

	assert.Equal(t, "😅 test 🌗 Україна \U0001FAE5 世界", text)
	require.Equal(t, 6, len(entities))

	assert.Equal(t, telego.MessageEntity{
		Type:   telego.EntityTypeItalic,
		Offset: 0,
		Length: 2,
	}, entities[0]) // "😅"
	assert.Equal(t, telego.MessageEntity{
		Type:   telego.EntityTypeBold,
		Offset: 3,
		Length: 4,
	}, entities[1]) // " test "
	assert.Equal(t, telego.MessageEntity{
		Type:   telego.EntityTypeItalic,
		Offset: 8,
		Length: 2,
	}, entities[2]) // "🌗"
	assert.Equal(t, telego.MessageEntity{
		Type:   telego.EntityTypeBold,
		Offset: 11,
		Length: 7,
	}, entities[3]) // " Україна"
	assert.Equal(t, telego.MessageEntity{
		Type:   telego.EntityTypeItalic,
		Offset: 19,
		Length: 2,
	}, entities[4]) // "\U0001FAE5 "
	assert.Equal(t, telego.MessageEntity{
		Type:   telego.EntityTypeBold,
		Offset: 22,
		Length: 2,
	}, entities[5]) // "世界"
}
