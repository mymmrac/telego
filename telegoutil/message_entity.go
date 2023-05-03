package telegoutil

import (
	"fmt"
	"strings"

	"github.com/mymmrac/telego"
)

// MessageEntityCollection represents text and slice of telego.MessageEntity associated with it
type MessageEntityCollection struct {
	text     string
	entities []telego.MessageEntity
}

// Entity creates new MessageEntityCollection with provided text and no entities
func Entity(text string) MessageEntityCollection {
	return MessageEntityCollection{
		text: text,
	}
}

// Entityf creates new MessageEntityCollection with the provided format and args and no entities
func Entityf(format string, args ...any) MessageEntityCollection {
	return MessageEntityCollection{
		text: fmt.Sprintf(format, args...),
	}
}

// Text returns text associated with a collection
func (c MessageEntityCollection) Text() string {
	return c.text
}

// Entities return message entities associated with a collection
func (c MessageEntityCollection) Entities() []telego.MessageEntity {
	return c.entities
}

// SetOffset sets offset for all entities
func (c MessageEntityCollection) SetOffset(offset int) {
	for i := range c.entities {
		c.entities[i].Offset = offset
	}
}

// Mention assigns mention entity and returns new collection
func (c MessageEntityCollection) Mention() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeMention,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// Hashtag assigns hashtag entity and returns a new collection
func (c MessageEntityCollection) Hashtag() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeHashtag,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// Cashtag assigns cashtag entity and returns a new collection
func (c MessageEntityCollection) Cashtag() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeCashtag,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// BotCommand assigns bot command entity and returns a new collection
func (c MessageEntityCollection) BotCommand() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeBotCommand,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// URL assigns url entity and returns a new collection
func (c MessageEntityCollection) URL() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeURL,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// Email assigns email entity and returns a new collection
func (c MessageEntityCollection) Email() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeEmail,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// PhoneNumber assigns phone number entity and returns a new collection
func (c MessageEntityCollection) PhoneNumber() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypePhoneNumber,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// Bold assigns bold entity and returns a new collection
func (c MessageEntityCollection) Bold() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeBold,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// Italic assigns italic entity and returns a new collection
func (c MessageEntityCollection) Italic() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeItalic,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// Underline assigns underline entity and returns new collection
func (c MessageEntityCollection) Underline() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeUnderline,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// Strikethrough assigns strikethrough entity and returns a new collection
func (c MessageEntityCollection) Strikethrough() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeStrikethrough,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// Spoiler assigns spoiler entity and returns new collection
func (c MessageEntityCollection) Spoiler() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeSpoiler,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// Code assigns code entity and returns new collection
func (c MessageEntityCollection) Code() MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeCode,
		Length: UTF16TextLen(c.text),
	})
	return c
}

// Pre assigns pre entity with language and returns a new collection
func (c MessageEntityCollection) Pre(language string) MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:     telego.EntityTypePre,
		Length:   UTF16TextLen(c.text),
		Language: language,
	})
	return c
}

// TextLink assigns text link entity with URL and returns a new collection
func (c MessageEntityCollection) TextLink(url string) MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeTextLink,
		Length: UTF16TextLen(c.text),
		URL:    url,
	})
	return c
}

// TextMention assigns text mention entity with user and returns new collection
func (c MessageEntityCollection) TextMention(user *telego.User) MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeTextMention,
		Length: UTF16TextLen(c.text),
		User:   user,
	})
	return c
}

// TextMentionWithID assigns text mention entity with just user ID and returns a new collection
func (c MessageEntityCollection) TextMentionWithID(userID int64) MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:   telego.EntityTypeTextMention,
		Length: UTF16TextLen(c.text),
		User:   &telego.User{ID: userID},
	})
	return c
}

// CustomEmoji assigns custom emoji entity and returns a new collection
func (c MessageEntityCollection) CustomEmoji(emojiID string) MessageEntityCollection {
	c.entities = append(c.entities, telego.MessageEntity{
		Type:          telego.EntityTypeCustomEmoji,
		Length:        UTF16TextLen(c.text),
		CustomEmojiID: emojiID,
	})
	return c
}

// MessageEntities coverts entity collections into the text and slice of [telego.MessageEntity] associated with text
// Note: Entity length is not trimmed as described in docs on purpose, Telegram still handles all entities perfectly
// fine, but trimming their length actually limits what can be sent
func MessageEntities(entityCollections ...MessageEntityCollection) (string, []telego.MessageEntity) {
	text := strings.Builder{}
	var entities []telego.MessageEntity

	for _, collection := range entityCollections {
		collection.SetOffset(UTF16TextLen(text.String()))
		entities = append(entities, collection.Entities()...)

		_, _ = text.WriteString(collection.Text())
	}

	return text.String(), entities
}

// UTF16TextLen returns length of a UTF-16 text
// Credit: https://core.telegram.org/api/entities#computing-entity-length
//
//nolint:gomnd
func UTF16TextLen(text string) int {
	length := 0
	for _, b := range []byte(text) {
		if (b & 0xc0) != 0x80 {
			length++
			if b >= 0xf0 {
				length++
			}
		}
	}
	return length
}
