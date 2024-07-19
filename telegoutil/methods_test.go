package telegoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chococola/telego"
)

var (
	id1  = telego.ChatID{ID: 123}
	id2  = telego.ChatID{ID: 321}
	file = telego.InputFile{FileID: "123"}

	text1 = "test"
	text2 = "test2"
	text3 = "test3"
	text4 = "test4"
	text5 = "test5"

	textNonASCII = "Hello, 世界"

	format1 = "%s %s"

	latitude  = 123.0
	longitude = 321.0

	number1       = 123
	number2 int64 = 1234

	prices = []telego.LabeledPrice{
		{
			Label:  text1,
			Amount: number1,
		},
		{
			Label:  text2,
			Amount: number1,
		},
	}

	mediaGroups = []telego.InputMedia{
		&telego.InputMediaDocument{
			Type:  telego.MediaTypeDocument,
			Media: file,
		},
	}
)

func TestAnimation(t *testing.T) {
	a := Animation(id1, file)
	assert.Equal(t, id1, a.ChatID)
	assert.Equal(t, file, a.Animation)
}

func TestAudio(t *testing.T) {
	a := Audio(id1, file)
	assert.Equal(t, id1, a.ChatID)
	assert.Equal(t, file, a.Audio)
}

func TestChatAction(t *testing.T) {
	c := ChatAction(id1, text1)
	assert.Equal(t, id1, c.ChatID)
	assert.Equal(t, text1, c.Action)
}

func TestContact(t *testing.T) {
	c := Contact(id1, text1, text2)
	assert.Equal(t, id1, c.ChatID)
	assert.Equal(t, text1, c.PhoneNumber)
	assert.Equal(t, text2, c.FirstName)
}

func TestDice(t *testing.T) {
	d := Dice(id1, text1)
	assert.Equal(t, id1, d.ChatID)
	assert.Equal(t, text1, d.Emoji)
}

func TestDocument(t *testing.T) {
	d := Document(id1, file)
	assert.Equal(t, id1, d.ChatID)
	assert.Equal(t, file, d.Document)
}

func TestGame(t *testing.T) {
	g := Game(id1.ID, text1)
	assert.Equal(t, id1.ID, g.ChatID)
	assert.Equal(t, text1, g.GameShortName)
}

func TestInvoice(t *testing.T) {
	i := Invoice(id1, text1, text2, text3, text4, text5, prices...)
	assert.Equal(t, id1, i.ChatID)
	assert.Equal(t, text1, i.Title)
	assert.Equal(t, text2, i.Description)
	assert.Equal(t, text3, i.Payload)
	assert.Equal(t, text4, i.ProviderToken)
	assert.Equal(t, text5, i.Currency)
	assert.Equal(t, prices, i.Prices)
}

func TestLocation(t *testing.T) {
	l := Location(id1, latitude, longitude)
	assert.Equal(t, id1, l.ChatID)
	assert.InEpsilon(t, latitude, l.Latitude, epsilon)
	assert.InEpsilon(t, longitude, l.Longitude, epsilon)
}

func TestMediaGroup(t *testing.T) {
	m := MediaGroup(id1, mediaGroups...)
	assert.Equal(t, id1, m.ChatID)
	assert.Equal(t, mediaGroups, m.Media)
}

func TestMessage(t *testing.T) {
	m := Message(id1, text1)
	assert.Equal(t, id1, m.ChatID)
	assert.Equal(t, text1, m.Text)
}

func TestMessagef(t *testing.T) {
	m := Messagef(id1, format1, text1, text2)
	assert.Equal(t, id1, m.ChatID)
	assert.Equal(t, text1+" "+text2, m.Text)
}

func TestMessageWithEntities(t *testing.T) {
	m := MessageWithEntities(id1, Entity(text1).Italic())
	assert.Equal(t, id1, m.ChatID)
	assert.Equal(t, text1, m.Text)
	assert.Equal(t, []telego.MessageEntity{
		{Type: "italic", Offset: 0, Length: 4, URL: "", User: nil, Language: ""},
	}, m.Entities)
}

func TestPhoto(t *testing.T) {
	p := Photo(id1, file)
	assert.Equal(t, id1, p.ChatID)
	assert.Equal(t, file, p.Photo)
}

func TestPoll(t *testing.T) {
	ops := []telego.InputPollOption{
		{Text: text2},
		{Text: text3},
	}
	p := Poll(id1, text1, ops...)
	assert.Equal(t, id1, p.ChatID)
	assert.Equal(t, text1, p.Question)
	assert.Equal(t, ops, p.Options)
}

func TestPollOption(t *testing.T) {
	p := PollOption(text1)
	assert.Equal(t, text1, p.Text)
}

func TestSticker(t *testing.T) {
	s := Sticker(id1, file)
	assert.Equal(t, id1, s.ChatID)
	assert.Equal(t, file, s.Sticker)
}

func TestVenue(t *testing.T) {
	v := Venue(id1, latitude, longitude, text1, text2)
	assert.Equal(t, id1, v.ChatID)
	assert.InEpsilon(t, latitude, v.Latitude, epsilon)
	assert.InEpsilon(t, longitude, v.Longitude, epsilon)
	assert.Equal(t, text1, v.Title)
	assert.Equal(t, text2, v.Address)
}

func TestVideo(t *testing.T) {
	v := Video(id1, file)
	assert.Equal(t, id1, v.ChatID)
	assert.Equal(t, file, v.Video)
}

func TestVideoNote(t *testing.T) {
	v := VideoNote(id1, file)
	assert.Equal(t, id1, v.ChatID)
	assert.Equal(t, file, v.VideoNote)
}

func TestVoice(t *testing.T) {
	v := Voice(id1, file)
	assert.Equal(t, id1, v.ChatID)
	assert.Equal(t, file, v.Voice)
}

func TestCopyMessage(t *testing.T) {
	c := CopyMessage(id1, id2, number1)
	assert.Equal(t, id1, c.ChatID)
	assert.Equal(t, id2, c.FromChatID)
	assert.Equal(t, number1, c.MessageID)
}

func TestCallbackQuery(t *testing.T) {
	c := CallbackQuery(text1)
	assert.Equal(t, text1, c.CallbackQueryID)
}

func TestInlineQuery(t *testing.T) {
	i := InlineQuery(text1, &telego.InlineQueryResultArticle{Title: text2},
		&telego.InlineQueryResultArticle{Title: text3})
	assert.Equal(t, text1, i.InlineQueryID)
	assert.Equal(t, []telego.InlineQueryResult{
		&telego.InlineQueryResultArticle{Title: text2},
		&telego.InlineQueryResultArticle{Title: text3},
	}, i.Results)
}

func TestShippingQuery(t *testing.T) {
	c := ShippingQuery(text1, true, telego.ShippingOption{ID: text2}, telego.ShippingOption{ID: text3})
	assert.Equal(t, text1, c.ShippingQueryID)
	assert.True(t, c.Ok)
	assert.Equal(t, []telego.ShippingOption{{ID: text2}, {ID: text3}}, c.ShippingOptions)
}

func TestPreCheckoutQuery(t *testing.T) {
	c := PreCheckoutQuery(text1, true)
	assert.Equal(t, text1, c.PreCheckoutQueryID)
	assert.True(t, c.Ok)
}

func TestWebAppQuery(t *testing.T) {
	i := WebAppQuery(text1, &telego.InlineQueryResultArticle{Title: text2})
	assert.Equal(t, text1, i.WebAppQueryID)
	assert.Equal(t, &telego.InlineQueryResultArticle{Title: text2}, i.Result)
}

func TestWebhook(t *testing.T) {
	w := Webhook(text1)
	assert.Equal(t, text1, w.URL)
}

func TestDelete(t *testing.T) {
	d := Delete(id1, number1)
	assert.Equal(t, id1, d.ChatID)
	assert.Equal(t, number1, d.MessageID)
}
