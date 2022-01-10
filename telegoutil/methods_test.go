package telegoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego"
)

var (
	id   = telego.ChatID{ID: 123}
	file = telego.InputFile{FileID: "123"}

	text1 = "test"
	text2 = "test2"
	text3 = "test3"
	text4 = "test4"
	text5 = "test5"

	texts = []string{
		text2, text3,
	}

	float  = 123.0
	float2 = 321.0

	prices = []telego.LabeledPrice{
		{
			Label:  text1,
			Amount: 123,
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
	a := Animation(id, file)
	assert.Equal(t, id, a.ChatID)
	assert.Equal(t, file, a.Animation)
}

func TestAudio(t *testing.T) {
	a := Audio(id, file)
	assert.Equal(t, id, a.ChatID)
	assert.Equal(t, file, a.Audio)
}

func TestChatAction(t *testing.T) {
	c := ChatAction(id, text1)
	assert.Equal(t, id, c.ChatID)
	assert.Equal(t, text1, c.Action)
}

func TestContact(t *testing.T) {
	c := Contact(id, text1, text2)
	assert.Equal(t, id, c.ChatID)
	assert.Equal(t, text1, c.PhoneNumber)
	assert.Equal(t, text2, c.FirstName)
}

func TestDice(t *testing.T) {
	d := Dice(id, text1)
	assert.Equal(t, id, d.ChatID)
	assert.Equal(t, text1, d.Emoji)
}

func TestDocument(t *testing.T) {
	d := Document(id, file)
	assert.Equal(t, id, d.ChatID)
	assert.Equal(t, file, d.Document)
}

func TestGame(t *testing.T) {
	g := Game(id.ID, text1)
	assert.Equal(t, id.ID, g.ChatID)
	assert.Equal(t, text1, g.GameShortName)
}

func TestInvoice(t *testing.T) {
	i := Invoice(id, text1, text2, text3, text4, text5, prices)
	assert.Equal(t, id, i.ChatID)
	assert.Equal(t, text1, i.Title)
	assert.Equal(t, text2, i.Description)
	assert.Equal(t, text3, i.Payload)
	assert.Equal(t, text4, i.ProviderToken)
	assert.Equal(t, text5, i.Currency)
	assert.Equal(t, prices, i.Prices)
}

func TestLocation(t *testing.T) {
	l := Location(id, float, float2)
	assert.Equal(t, id, l.ChatID)
	assert.Equal(t, float, l.Latitude)
	assert.Equal(t, float2, l.Longitude)
}

func TestMediaGroup(t *testing.T) {
	m := MediaGroup(id, mediaGroups)
	assert.Equal(t, id, m.ChatID)
	assert.Equal(t, mediaGroups, m.Media)
}

func TestMessage(t *testing.T) {
	m := Message(id, text1)
	assert.Equal(t, id, m.ChatID)
	assert.Equal(t, text1, m.Text)
}

func TestPhoto(t *testing.T) {
	p := Photo(id, file)
	assert.Equal(t, id, p.ChatID)
	assert.Equal(t, file, p.Photo)
}

func TestPoll(t *testing.T) {
	p := Poll(id, text1, texts)
	assert.Equal(t, id, p.ChatID)
	assert.Equal(t, text1, p.Question)
	assert.Equal(t, texts, p.Options)
}

func TestSticker(t *testing.T) {
	s := Sticker(id, file)
	assert.Equal(t, id, s.ChatID)
	assert.Equal(t, file, s.Sticker)
}

func TestVenue(t *testing.T) {
	v := Venue(id, float, float2, text1, text2)
	assert.Equal(t, id, v.ChatID)
	assert.Equal(t, float, v.Latitude)
	assert.Equal(t, float2, v.Longitude)
	assert.Equal(t, text1, v.Title)
	assert.Equal(t, text2, v.Address)
}

func TestVideo(t *testing.T) {
	v := Video(id, file)
	assert.Equal(t, id, v.ChatID)
	assert.Equal(t, file, v.Video)
}

func TestVideoNote(t *testing.T) {
	v := VideoNote(id, file)
	assert.Equal(t, id, v.ChatID)
	assert.Equal(t, file, v.VideoNote)
}

func TestVoice(t *testing.T) {
	v := Voice(id, file)
	assert.Equal(t, id, v.ChatID)
	assert.Equal(t, file, v.Voice)
}
