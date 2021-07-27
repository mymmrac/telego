package telego

import (
	"encoding/json"
	"fmt"
	"os"
)

// Update - This object (#available-types) represents an incoming update.At most one of the optional parameters
// can be present in any given update.
type Update struct {
	// UpdateID - The update's unique identifier. Update identifiers start from a certain positive number
	// and increase sequentially. This ID becomes especially handy if you're using Webhooks (#setwebhook), since
	// it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order.
	// If there are no new updates for at least a week, then identifier of the next update will be chosen randomly
	// instead of sequentially.
	UpdateID int `json:"update_id"`

	// Message - Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`

	// EditedMessage - Optional. New version of a message that is known to the bot and was edited
	EditedMessage *Message `json:"edited_message,omitempty"`

	// ChannelPost - Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`

	// EditedChannelPost - Optional. New version of a channel post that is known to the bot and was edited
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// InlineQuery - Optional. New incoming inline (#inline-mode) query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`

	// ChosenInlineResult - Optional. The result of an inline (#inline-mode) query that was chosen by a user and
	// sent to their chat partner. Please see our documentation on the feedback collecting
	// (/bots/inline#collecting-feedback) for details on how to enable these updates for your bot.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

	// CallbackQuery - Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

	// ShippingQuery - Optional. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	// PreCheckoutQuery - Optional. New incoming pre-checkout query. Contains full information about checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	// Poll - Optional. New poll state. Bots receive only updates about stopped polls and polls, which
	// are sent by the bot
	Poll *Poll `json:"poll,omitempty"`

	// PollAnswer - Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls
	// that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`

	// MyChatMember - Optional. The bot's chat member status was updated in a chat. For private chats, this update is
	// received only when the bot is blocked or unblocked by the user.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`

	// ChatMember - Optional. A chat member's status was updated in a chat. The bot must be an administrator in the
	// chat and must explicitly specify “chat_member” in the list of allowed_updates to receive these updates.
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`
}

// WebhookInfo - Contains information about the current status of a webhook.
type WebhookInfo struct {
	// URL - Webhook URL, may be empty if webhook is not set up
	URL string `json:"url"`

	// HasCustomCertificate - True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// PendingUpdateCount - Number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`

	// IPAddress - Optional. Currently used webhook IP address
	IPAddress string `json:"ip_address,omitempty"`

	// LastErrorDate - Optional. Unix time for the most recent error that happened when trying to deliver
	// an update via webhook
	LastErrorDate int `json:"last_error_date,omitempty"`

	// LastErrorMessage - Optional. Error message in human-readable format for the most recent error that happened
	// when trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message,omitempty"`

	// MaxConnections - Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for
	// update delivery
	MaxConnections int `json:"max_connections,omitempty"`

	// AllowedUpdates - Optional. A list of update types the bot is subscribed to. Defaults to all update types
	// except chat_member
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// User - This object represents a Telegram user or bot.
type User struct {
	// ID - Unique identifier for this user or bot. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant
	// bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	ID int64 `json:"id"`

	// IsBot - True, if this user is a bot
	IsBot bool `json:"is_bot"`

	// FirstName - User's or bot's first name
	FirstName string `json:"first_name"`

	// LastName - Optional. User's or bot's last name
	LastName string `json:"last_name,omitempty"`

	// Username - Optional. User's or bot's username
	Username string `json:"username,omitempty"`

	// LanguageCode - Optional. IETF language tag (https://en.wikipedia.org/wiki/IETF_language_tag)
	// of the user's language
	LanguageCode string `json:"language_code,omitempty"`

	// CanJoinGroups - Optional. True, if the bot can be invited to groups. Returned only in getMe (#getme).
	CanJoinGroups bool `json:"can_join_groups,omitempty"`

	// CanReadAllGroupMessages - Optional. True, if privacy mode (https://core.telegram.org/bots#privacy-mode) is
	// disabled for the bot. Returned only in getMe (#getme).
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`

	// SupportsInlineQueries - Optional. True, if the bot supports inline queries. Returned only in getMe (#getme).
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}

// Chat - This object represents a chat.
type Chat struct {
	// ID - Unique identifier for this chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant
	// bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	ID int64 `json:"id"`

	// Type - Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`

	// Title - Optional. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`

	// Username - Optional. Username, for private chats, supergroups and channels if available
	Username string `json:"username,omitempty"`

	// FirstName - Optional. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`

	// LastName - Optional. Last name of the other party in a private chat
	LastName string `json:"last_name,omitempty"`

	// Photo - Optional. Chat photo. Returned only in getChat (#getchat).
	Photo *ChatPhoto `json:"photo,omitempty"`

	// Bio - Optional. Bio of the other party in a private chat. Returned only in getChat (#getchat).
	Bio string `json:"bio,omitempty"`

	// Description - Optional. Description, for groups, supergroups and channel chats.
	// Returned only in getChat (#getchat).
	Description string `json:"description,omitempty"`

	// InviteLink - Optional. Primary invite link, for groups, supergroups and channel chats.
	// Returned only in getChat (#getchat).
	InviteLink string `json:"invite_link,omitempty"`

	// PinnedMessage - Optional. The most recent pinned message (by sending date). Returned only in getChat (#getchat).
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Permissions - Optional. Default chat member permissions, for groups and supergroups.
	// Returned only in getChat (#getchat).
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// SlowModeDelay - Optional. For supergroups, the minimum allowed delay between consecutive messages sent by
	// each unprivileged user. Returned only in getChat (#getchat).
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`

	// MessageAutoDeleteTime - Optional. The time after which all messages sent to the chat will be
	// automatically deleted; in seconds. Returned only in getChat (#getchat).
	MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`

	// StickerSetName - Optional. For supergroups, name of group sticker set. Returned only in getChat (#getchat).
	StickerSetName string `json:"sticker_set_name,omitempty"`

	// CanSetStickerSet - Optional. True, if the bot can change the group sticker set. Returned only
	// in getChat (#getchat).
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`

	// LinkedChatID - Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for
	// a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and
	// some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits,
	// so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	// Returned only in getChat (#getchat).
	LinkedChatID int64 `json:"linked_chat_id,omitempty"`

	// Location - Optional. For supergroups, the location to which the supergroup is connected.
	// Returned only in getChat (#getchat).
	Location *ChatLocation `json:"location,omitempty"`
}

// Message - This object represents a message.
type Message struct {
	// MessageID - Unique message identifier inside this chat
	MessageID int `json:"message_id"`

	// From - Optional. Sender, empty for messages sent to channels
	From *User `json:"from,omitempty"`

	// SenderChat - Optional. Sender of the message, sent on behalf of a chat. The channel itself for channel messages.
	// The supergroup itself for messages from anonymous group administrators. The linked channel for messages
	// automatically forwarded to the discussion group
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// Date - Date the message was sent in Unix time
	Date int `json:"date"`

	// Chat - Conversation the message belongs to
	Chat Chat `json:"chat"`

	// ForwardFrom - Optional. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from,omitempty"`

	// ForwardFromChat - Optional. For messages forwarded from channels or from anonymous administrators,
	// information about the original sender chat
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`

	// ForwardFromMessageID - Optional. For messages forwarded from channels, identifier of the original
	// message in the channel
	ForwardFromMessageID int `json:"forward_from_message_id,omitempty"`

	// ForwardSignature - Optional. For messages forwarded from channels, signature of the post author if present
	ForwardSignature string `json:"forward_signature,omitempty"`

	// ForwardSenderName - Optional. Sender's name for messages forwarded from users who disallow adding a link to
	// their account in forwarded messages
	ForwardSenderName string `json:"forward_sender_name,omitempty"`

	// ForwardDate - Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardDate int `json:"forward_date,omitempty"`

	// ReplyToMessage - Optional. For replies, the original message. Note that the Message object in this field
	// will not contain further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// ViaBot - Optional. Bot through which the message was sent
	ViaBot *User `json:"via_bot,omitempty"`

	// EditDate - Optional. Date the message was last edited in Unix time
	EditDate int `json:"edit_date,omitempty"`

	// MediaGroupID - Optional. The unique identifier of a media message group this message belongs to
	MediaGroupID string `json:"media_group_id,omitempty"`

	// AuthorSignature - Optional. Signature of the post author for messages in channels, or the custom title
	// of an anonymous group administrator
	AuthorSignature string `json:"author_signature,omitempty"`

	// Text - Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters
	Text string `json:"text,omitempty"`

	// Entities - Optional. For text messages, special entities like usernames, URLs, bot commands, etc.
	// that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`

	// Animation - Optional. Message is an animation, information about the animation. For backward compatibility,
	// when this field is set, the document field will also be set
	Animation *Animation `json:"animation,omitempty"`

	// Audio - Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`

	// Document - Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`

	// Photo - Optional. Message is a photo, available sizes of the photo
	Photo []PhotoSize `json:"photo,omitempty"`

	// Sticker - Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// Video - Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`

	// VideoNote - Optional. Message is a video note (https://telegram.org/blog/video-messages-and-telescope),
	// information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Voice - Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`

	// Caption - Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	Caption string `json:"caption,omitempty"`

	// CaptionEntities - Optional. For messages with a caption, special entities like usernames, URLs,
	// bot commands, etc. that appear in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Contact - Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`

	// Dice - Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`

	// Game - Optional. Message is a game, information about the game. More about games » (#games)
	Game *Game `json:"game,omitempty"`

	// Poll - Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`

	// Venue - Optional. Message is a venue, information about the venue. For backward compatibility,
	// when this field is set, the location field will also be set
	Venue *Venue `json:"venue,omitempty"`

	// Location - Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`

	// NewChatMembers - Optional. New members that were added to the group or supergroup and information about
	// them (the bot itself may be one of these members)
	NewChatMembers []User `json:"new_chat_members,omitempty"`

	// LeftChatMember - Optional. A member was removed from the group, information about
	// them (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`

	// NewChatTitle - Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`

	// NewChatPhoto - Optional. A chat photo was change to this value
	NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`

	// DeleteChatPhoto - Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`

	// GroupChatCreated - Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`

	// SupergroupChatCreated - Optional. Service message: the supergroup has been created.
	// This field can't be received in a message coming through updates, because bot can't be a member of
	// a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very
	// first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`

	// ChannelChatCreated - Optional. Service message: the channel has been created.
	// This field can't be received in a message coming through updates, because bot can't be a member
	// of a channel when it is created. It can only be found in reply_to_message if someone replies to a v
	// ery first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`

	// MessageAutoDeleteTimerChanged - Optional. Service message: auto-delete timer settings changed in the chat
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`

	// MigrateToChatID - Optional. The group has been migrated to a supergroup with the specified identifier.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// MigrateFromChatID - Optional. The supergroup has been migrated from a group with the specified identifier.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`

	// PinnedMessage - Optional. Specified message was pinned. Note that the Message object in this field will not
	// contain further reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Invoice - Optional. Message is an invoice for a payment (#payments), information about the invoice.
	// More about payments » (#payments)
	Invoice *Invoice `json:"invoice,omitempty"`

	// SuccessfulPayment - Optional. Message is a service message about a successful payment,
	// information about the payment. More about payments » (#payments)
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`

	// ConnectedWebsite - Optional. The domain name of the website on which the user has logged in.
	// More about Telegram Login » (/widgets/login)
	ConnectedWebsite string `json:"connected_website,omitempty"`

	// PassportData - Optional. Telegram Passport data
	PassportData *PassportData `json:"passport_data,omitempty"`

	// ProximityAlertTriggered - Optional. Service message. A user in the chat triggered another user's
	// proximity alert while sharing Live Location.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`

	// VoiceChatScheduled - Optional. Service message: voice chat scheduled
	VoiceChatScheduled *VoiceChatScheduled `json:"voice_chat_scheduled,omitempty"`

	// VoiceChatStarted - Optional. Service message: voice chat started
	VoiceChatStarted *VoiceChatStarted `json:"voice_chat_started,omitempty"`

	// VoiceChatEnded - Optional. Service message: voice chat ended
	VoiceChatEnded *VoiceChatEnded `json:"voice_chat_ended,omitempty"`

	// VoiceChatParticipantsInvited - Optional. Service message: new participants invited to a voice chat
	VoiceChatParticipantsInvited *VoiceChatParticipantsInvited `json:"voice_chat_participants_invited,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard attached to the message. login_URL buttons are represented
	// as ordinary URL buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// MessageID - This object represents a unique message identifier.
type MessageID struct {
	// MessageID - Unique message identifier
	MessageID int `json:"message_id"`
}

// MessageEntity - This object represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type - Type of the entity. Can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD),
	// “bot_command” (/start@jobs_bot), “URL” (https://telegram.org), “email” (do-not-reply@telegram.org),
	// “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text),
	// “strikethrough” (strikethrough text), “code” (monowidth string), “pre” (monowidth block),
	// “text_link” (for clickable text URLs), “text_mention” (for users without
	// usernames (https://telegram.org/blog/edit#new-mentions))
	Type string `json:"type"`

	// Offset - Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`

	// Length - Length of the entity in UTF-16 code units
	Length int `json:"length"`

	// URL - Optional. For “text_link” only, URL that will be opened after user taps on the text
	URL string `json:"url,omitempty"`

	// User - Optional. For “text_mention” only, the mentioned user
	User *User `json:"user,omitempty"`

	// Language - Optional. For “pre” only, the programming language of the entity text
	Language string `json:"language,omitempty"`
}

// PhotoSize - This object represents one size of a photo or a file (#document) / sticker (#sticker) thumbnail.
type PhotoSize struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width - Photo width
	Width int `json:"width"`

	// Height - Photo height
	Height int `json:"height"`

	// FileSize - Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Animation - This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width - Video width as defined by sender
	Width int `json:"width"`

	// Height - Video height as defined by sender
	Height int `json:"height"`

	// Duration - Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`

	// Thumb - Optional. Animation thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileName - Optional. Original animation filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// FileSize - Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Audio - This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration - Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`

	// Performer - Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer,omitempty"`

	// Title - Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title,omitempty"`

	// FileName - Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// FileSize - Optional. File size
	FileSize int `json:"file_size,omitempty"`

	// Thumb - Optional. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// Document - This object represents a general file (as opposed to photos (#photosize), voice messages (#voice)
// and audio files (#audio)).
type Document struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Thumb - Optional. Document thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileName - Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// FileSize - Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Video - This object represents a video file.
type Video struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width - Video width as defined by sender
	Width int `json:"width"`

	// Height - Video height as defined by sender
	Height int `json:"height"`

	// Duration - Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`

	// Thumb - Optional. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileName - Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. Mime type of a file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// FileSize - Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// VideoNote - This object represents a video message (https://telegram.org/blog/video-messages-and-telescope)
// (available in Telegram apps as of v.4.0 (https://telegram.org/blog/video-messages-and-telescope)).
type VideoNote struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Length - Video width and height (diameter of the video message) as defined by sender
	Length int `json:"length"`

	// Duration - Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`

	// Thumb - Optional. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileSize - Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Voice - This object represents a voice note.
type Voice struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration - Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`

	// MimeType - Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// FileSize - Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Contact - This object represents a phone contact.
type Contact struct {
	// PhoneNumber - Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// FirstName - Contact's first name
	FirstName string `json:"first_name"`

	// LastName - Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// UserID - Optional. Contact's user identifier in Telegram. This number may have more than 32 significant
	// bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at
	// most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	UserID int64 `json:"user_id,omitempty"`

	// Vcard - Optional. Additional data about the contact in the form of a vCard (https://en.wikipedia.org/wiki/VCard)
	Vcard string `json:"vcard,omitempty"`
}

// Dice - This object represents an animated emoji that displays a random value.
type Dice struct {
	// Emoji - Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`

	// Value - Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀” and “⚽” base emoji,
	// 1-64 for “🎰” base emoji
	Value int `json:"value"`
}

// PollOption - This object contains information about one answer option in a poll.
type PollOption struct {
	// Text - Option text, 1-100 characters
	Text string `json:"text"`

	// VoterCount - Number of users that voted for this option
	VoterCount int `json:"voter_count"`
}

// PollAnswer - This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// PollID - Unique poll identifier
	PollID string `json:"poll_id"`

	// User - The user, who changed the answer to the poll
	User User `json:"user"`

	// OptionIDs - 0-based identifiers of answer options, chosen by the user.
	// May be empty if the user retracted their vote.
	OptionIDs []int `json:"option_ids"`
}

// Poll - This object contains information about a poll.
type Poll struct {
	// ID - Unique poll identifier
	ID string `json:"id"`

	// Question - Poll question, 1-300 characters
	Question string `json:"question"`

	// Options - List of poll options
	Options []PollOption `json:"options"`

	// TotalVoterCount - Total number of users that voted in the poll
	TotalVoterCount int `json:"total_voter_count"`

	// IsClosed - True, if the poll is closed
	IsClosed bool `json:"is_closed"`

	// IsAnonymous - True, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`

	// Type - Poll type, currently can be “regular” or “quiz”
	Type string `json:"type"`

	// AllowsMultipleAnswers - True, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`

	// CorrectOptionID - Optional. 0-based identifier of the correct answer option.
	// Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or
	// to the private chat with the bot.
	CorrectOptionID int `json:"correct_option_id,omitempty"`

	// Explanation - Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon
	// in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation,omitempty"`

	// ExplanationEntities - Optional. Special entities like usernames, URLs, bot commands, etc. that
	// appear in the explanation
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`

	// OpenPeriod - Optional. Amount of time in seconds the poll will be active after creation
	OpenPeriod int `json:"open_period,omitempty"`

	// CloseDate - Optional. Point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate int `json:"close_date,omitempty"`
}

// Location - This object represents a point on the map.
type Location struct {
	// Longitude - Longitude as defined by sender
	Longitude float64 `json:"longitude"`

	// Latitude - Latitude as defined by sender
	Latitude float64 `json:"latitude"`

	// HorizontalAccuracy - Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// LivePeriod - Optional. Time relative to the message sending date, during which the location can be updated,
	// in seconds. For active live locations only.
	LivePeriod int `json:"live_period,omitempty"`

	// Heading - Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	Heading int `json:"heading,omitempty"`

	// ProximityAlertRadius - Optional. Maximum distance for proximity alerts about approaching another chat member,
	// in meters. For sent live locations only.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

// Venue - This object represents a venue.
type Venue struct {
	// Location - Venue location. Can't be a live location
	Location Location `json:"location"`

	// Title - Name of the venue
	Title string `json:"title"`

	// Address - Address of the venue
	Address string `json:"address"`

	// FoursquareID - Optional. Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`

	// FoursquareType - Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// GooglePlaceID - Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// GooglePlaceType - Optional. Google Places type of the venue. (See supported
	// types (https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// ProximityAlertTriggered - This object represents the content of a service message, sent whenever
// a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	// Traveler - User that triggered the alert
	Traveler User `json:"traveler"`

	// Watcher - User that set the alert
	Watcher User `json:"watcher"`

	// Distance - The distance between the users
	Distance int `json:"distance"`
}

// MessageAutoDeleteTimerChanged - This object represents a service message about a change
// in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	// MessageAutoDeleteTime - New auto-delete time for messages in the chat
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// VoiceChatScheduled - This object represents a service message about a voice chat scheduled in the chat.
type VoiceChatScheduled struct {
	// StartDate - Point in time (Unix timestamp) when the voice chat is supposed to be started by a chat administrator
	StartDate int `json:"start_date"`
}

// VoiceChatStarted - This object represents a service message about a voice chat started in the chat.
// Currently holds no information.
type VoiceChatStarted struct{}

// VoiceChatEnded - This object represents a service message about a voice chat ended in the chat.
type VoiceChatEnded struct {
	// Duration - Voice chat duration; in seconds
	Duration int `json:"duration"`
}

// VoiceChatParticipantsInvited - This object represents a service message about new members invited to a voice chat.
type VoiceChatParticipantsInvited struct {
	// Users - Optional. New members that were invited to the voice chat
	Users []User `json:"users,omitempty"`
}

// UserProfilePhotos - This object represent a user's profile pictures.
type UserProfilePhotos struct {
	// TotalCount - Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`

	// Photos - Requested profile pictures (in up to 4 sizes each)
	Photos []PhotoSize `json:"photos"`
}

// File - This object represents a file ready to be downloaded. The file can be downloaded via the
// link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for
// at least 1 hour. When the link expires, a new one can be requested by calling getFile (#getfile).
// Maximum file size to download is 20 MB
type File struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// FileSize - Optional. File size, if known
	FileSize int `json:"file_size,omitempty"`

	// FilePath - Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath string `json:"file_path,omitempty"`
}

// ReplyMarkup - Represents reply markup (inline keyboard, custom reply keyboard, etc.)
type ReplyMarkup interface {
	// ReplyType - Returns type of reply
	ReplyType() string
}

// ReplyKeyboardMarkup - This object represents a custom keyboard (https://core.telegram.org/bots#keyboards) with
// reply options (see Introduction to bots (https://core.telegram.org/bots#keyboards) for details and examples).
type ReplyKeyboardMarkup struct {
	// Keyboard - Array of button rows, each represented by an Array of KeyboardButton (#keyboardbutton) objects
	Keyboard []KeyboardButton `json:"keyboard"`

	// ResizeKeyboard - Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make
	// the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom
	// keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// OneTimeKeyboard - Optional. Requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display the usual letter-keyboard
	// in the chat – the user can press a special button in the input field to see the custom keyboard again.
	// Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// InputFieldPlaceholder - Optional. The placeholder to be shown in the input field when the
	// keyboard is active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective - Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1)
	// users that are @mentioned in the text of the Message (#message) object; 2) if the bot's message is
	// a reply (has reply_to_message_id), sender of the original message.Example: A user requests to change
	// the bot's language, bot replies to the request with a keyboard to select the new language. Other users in
	// the group don't see the keyboard.
	Selective bool `json:"selective,omitempty"`
}

// ReplyType - Returns ReplyKeyboardMarkup type
func (i *ReplyKeyboardMarkup) ReplyType() string {
	return "ReplyKeyboardMarkup"
}

// KeyboardButton - This object represents one button of the reply keyboard. For simple text buttons String can be
// used instead of this object to specify text of the button. Optional fields request_contact, request_location,
// and request_poll are mutually exclusive.
type KeyboardButton struct {
	// Text - Text of the button. If none of the optional fields are used, it will be sent as a message when
	// the button is pressed
	Text string `json:"text"`

	// RequestContact - Optional. If True, the user's phone number will be sent as a contact when
	// the button is pressed. Available in private chats only
	RequestContact bool `json:"request_contact,omitempty"`

	// RequestLocation - Optional. If True, the user's current location will be sent when the button is pressed.
	// Available in private chats only
	RequestLocation bool `json:"request_location,omitempty"`

	// RequestPoll - Optional. If specified, the user will be asked to create a poll and send it to the bot when
	// the button is pressed. Available in private chats only
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

// KeyboardButtonPollType - This object represents type of a poll, which is allowed to be created and sent
// when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// Type - Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode.
	// If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed
	// to create a poll of any type.
	Type string `json:"type,omitempty"`
}

// ReplyKeyboardRemove - Upon receiving a message with this object, Telegram clients will remove the current
// custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until
// a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately
// after the user presses a button (see ReplyKeyboardMarkup (#replykeyboardmarkup)).
type ReplyKeyboardRemove struct {
	// RemoveKeyboard - Requests clients to remove the custom keyboard (user will not be able to summon this keyboard;
	// if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard
	// in ReplyKeyboardMarkup (#replykeyboardmarkup))
	RemoveKeyboard bool `json:"remove_keyboard"`

	// Selective - Optional. Use this parameter if you want to remove the keyboard for specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message (#message) object; 2) if the bot's message
	// is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll,
	// bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still
	// showing the keyboard with poll options to users who haven't voted yet.
	Selective bool `json:"selective,omitempty"`
}

// ReplyType - Returns ReplyKeyboardRemove type
func (i *ReplyKeyboardRemove) ReplyType() string {
	return "ReplyKeyboardRemove"
}

// InlineKeyboardMarkup - This object represents an inline keyboard
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating) that appears right next
// to the message it belongs to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard - Array of button rows, each represented by an Array of InlineKeyboardButton
	// (#inlinekeyboardbutton) objects
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// ReplyType - Returns InlineKeyboardMarkup type
func (i *InlineKeyboardMarkup) ReplyType() string {
	return "InlineKeyboardMarkup"
}

// InlineKeyboardButton - This object represents one button of an inline keyboard. You must use exactly
// one of the optional fields.
type InlineKeyboardButton struct {
	// Text - Label text on the button
	Text string `json:"text"`

	// URL - Optional. HTTP or tg:// URL to be opened when button is pressed
	URL string `json:"url,omitempty"`

	// LoginURL - Optional. An HTTP URL used to automatically authorize the user. Can be used as a replacement
	// for the Telegram Login Widget (https://core.telegram.org/widgets/login).
	LoginURL *LoginURL `json:"login_url,omitempty"`

	// CallbackData - Optional. Data to be sent in a callback query (#callbackquery) to the bot when button
	// is pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`

	// SwitchInlineQuery - Optional. If set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input field. Can be empty,
	// in which case just the bot's username will be inserted.Note: This offers an easy way for users to start using
	// your bot in inline mode (/bots/inline) when they are currently in a private chat with it. Especially useful
	// when combined with switch_pm… (#answerinlinequery) actions – in this case the user will be automatically
	// returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

	// SwitchInlineQueryCurrentChat - Optional. If set, pressing the button will insert the bot's username and
	// the specified inline query in the current chat's input field. Can be empty, in which case only the bot's username
	// will be inserted.This offers a quick way for the user to open your bot in inline mode in the same chat – good
	// for selecting something from multiple options.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

	// CallbackGame - Optional. Description of the game that will be launched when the user presses the button.NOTE:
	// This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Pay - Optional. Specify True, to send a Pay button (#payments).NOTE: This type of button must always be
	// the first button in the first row.
	Pay bool `json:"pay,omitempty"`
}

// LoginURL - This object represents a parameter of the inline keyboard button used to automatically
// authorize a user. Serves as a great replacement for the Telegram Login Widget
// (https://core.telegram.org/widgets/login) when the user is coming from Telegram.
// All the user needs to do is tap/click a button and confirm that they want to log in:
// Image: https://core.telegram.org/file/811140015/1734/8VZFkwWXalM.97872/6127fa62d8a0bf2b3c
// Telegram apps support these buttons as of version 5.7.
// Sample bot: @discussbot
type LoginURL struct {
	// URL - An HTTP URL to be opened with user authorization data added to the query string when the button is
	// pressed. If the user refuses to provide authorization data, the original URL without information about the
	// user will be opened. The data added is the same as described in Receiving authorization
	// data (https://core.telegram.org/widgets/login#receiving-authorization-data).NOTE: You must always check the
	// hash of the received data to verify the authentication and the integrity of the data as described in Checking
	// authorization (https://core.telegram.org/widgets/login#checking-authorization).
	URL string `json:"url"`

	// ForwardText - Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`

	// BotUsername - Optional. Username of a bot, which will be used for user authorization.
	// See Setting up a bot (https://core.telegram.org/widgets/login#setting-up-a-bot) for more details.
	// If not specified, the current bot's username will be assumed. The URL's domain must be the same as the domain
	// linked with the bot. See Linking your domain to the
	// bot (https://core.telegram.org/widgets/login#linking-your-domain-to-the-bot) for more details.
	BotUsername string `json:"bot_username,omitempty"`

	// RequestWriteAccess - Optional. Pass True to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// CallbackQuery - This object represents an incoming callback query from a callback button in an inline
// keyboard (/bots#inline-keyboards-and-on-the-fly-updating). If the button that originated the query was attached to a
// message sent by the bot, the field message will be present. If the button was attached to a message sent via the
// bot (in inline mode (#inline-mode)), the field inline_message_id will be present. Exactly one of the fields
// data or game_short_name will be present.
//
// NOTE: After the user presses a callback button, Telegram clients will display a progress bar until you call
// answerCallbackQuery (#answercallbackquery). It is, therefore, necessary to react by calling
// answerCallbackQuery (#answercallbackquery) even if no notification to the user is
// needed (e.g., without specifying any of the optional parameters).
type CallbackQuery struct {
	// ID - Unique identifier for this query
	ID string `json:"id"`

	// From - Sender
	From User `json:"from"`

	// Message - Optional. Message with the callback button that originated the query. Note that message content
	// and message date will not be available if the message is too old
	Message *Message `json:"message,omitempty"`

	// InlineMessageID - Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// ChatInstance - Global identifier, uniquely corresponding to the chat to which the message with the callback
	// button was sent. Useful for high scores in games (#games).
	ChatInstance string `json:"chat_instance"`

	// Data - Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary
	// data in this field.
	Data string `json:"data,omitempty"`

	// GameShortName - Optional. Short name of a Game (#games) to be returned, serves as the unique identifier for the game
	GameShortName string `json:"game_short_name,omitempty"`
}

// ForceReply - Upon receiving a message with this object, Telegram clients will display a reply interface to the
// user (act as if the user has selected the bot's message and tapped 'Reply'). This can be extremely useful if you
// want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode (/bots#privacy-mode).
// Example: https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	// ForceReply - Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	ForceReply bool `json:"force_reply"`

	// InputFieldPlaceholder - Optional. The placeholder to be shown in the input field when the reply
	// is active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective - Optional. Use this parameter if you want to force reply from specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message (#message) object; 2) if the bot's message
	// is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

// ReplyType - Returns ForceReply type
func (i *ForceReply) ReplyType() string {
	return "ForceReply"
}

// ChatPhoto - This object represents a chat photo.
type ChatPhoto struct {
	// SmallFileID - File identifier of small (160x160) chat photo. This file_id can be used only for photo download
	// and only for as long as the photo is not changed.
	SmallFileID string `json:"small_file_id"`

	// SmallFileUniqueID - Unique file identifier of small (160x160) chat photo, which is supposed to be the same
	// over time and for different bots. Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id"`

	// BigFileID - File identifier of big (640x640) chat photo. This file_id can be used only for photo download and
	// only for as long as the photo is not changed.
	BigFileID string `json:"big_file_id"`

	// BigFileUniqueID - Unique file identifier of big (640x640) chat photo, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id"`
}

// ChatInviteLink - Represents an invite link for a chat.
type ChatInviteLink struct {
	// InviteLink - The invite link. If the link was created by another chat administrator,
	// then the second part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`

	// Creator - Creator of the link
	Creator User `json:"creator"`

	// IsPrimary - True, if the link is primary
	IsPrimary bool `json:"is_primary"`

	// IsRevoked - True, if the link is revoked
	IsRevoked bool `json:"is_revoked"`

	// ExpireDate - Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	ExpireDate int `json:"expire_date,omitempty"`

	// MemberLimit - Optional. Maximum number of users that can be members of the chat simultaneously after
	// joining the chat via this invite link; 1-99999
	MemberLimit int `json:"member_limit,omitempty"`
}

// ChatMember - This object contains information about one member of a chat. Currently, the following 6 types of
// chat members are supported:
// ChatMemberOwner
// ChatMemberAdministrator
// ChatMemberMember
// ChatMemberRestricted
// ChatMemberLeft
// ChatMemberBanned
type ChatMember interface {
	MemberStatus() string
}

type chatMemberData struct {
	Data ChatMember
}

func (c *chatMemberData) UnmarshalJSON(bytes []byte) error {
	var memberStatus struct {
		Status string `json:"status"`
	}

	err := json.Unmarshal(bytes, &memberStatus)
	if err != nil {
		return err
	}

	switch memberStatus.Status {
	case "creator":
		var cm *ChatMemberOwner
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	case "administrator":
		var cm *ChatMemberAdministrator
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	case "member":
		var cm *ChatMemberMember
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	case "restricted":
		var cm *ChatMemberRestricted
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	case "left":
		var cm *ChatMemberLeft
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	case "kicked":
		var cm *ChatMemberBanned
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	default:
		return fmt.Errorf("unknown member member status: %q", memberStatus.Status)
	}

	return err
}

// ChatMemberOwner - Represents a chat member (#chatmember) that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	// Status - The member's status in the chat, always “creator”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`

	// CustomTitle - Custom title for this user
	CustomTitle string `json:"custom_title,omitempty"`

	// IsAnonymous - True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`
}

func (c *ChatMemberOwner) MemberStatus() string {
	return c.Status
}

// ChatMemberAdministrator - Represents a chat member (#chatmember) that has some additional privileges.
type ChatMemberAdministrator struct {
	// Status - The member's status in the chat, always “administrator”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`

	// CanBeEdited - True, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited"`

	// CustomTitle - Custom title for this user
	CustomTitle string `json:"custom_title"`

	// IsAnonymous - True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// CanManageChat - True, if the administrator can access the chat event log, chat statistics, message
	// statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode.
	// Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat"`

	// CanPostMessages - True, if the administrator can post in the channel; channels only
	CanPostMessages bool `json:"can_post_messages"`

	// CanEditMessages - True, if the administrator can edit messages of other users and can pin messages;
	// channels only
	CanEditMessages bool `json:"can_edit_messages"`

	// CanDeleteMessages - True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`

	// CanManageVoiceChats - True, if the administrator can manage voice chats
	CanManageVoiceChats bool `json:"can_manage_voice_chats"`

	// CanRestrictMembers - True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members"`

	// CanPromoteMembers - True, if the administrator can add new administrators with a subset of their own privileges
	// or demote administrators that he has promoted, directly or indirectly (promoted by administrators
	// that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`

	// CanChangeInfo - True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// CanInviteUsers - True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// CanPinMessages - True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages"`
}

func (c *ChatMemberAdministrator) MemberStatus() string {
	return c.Status
}

// ChatMemberMember - Represents a chat member (#chatmember) that has no additional privileges or restrictions.
type ChatMemberMember struct {
	// Status - The member's status in the chat, always “member”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`
}

func (c *ChatMemberMember) MemberStatus() string {
	return c.Status
}

// ChatMemberRestricted - Represents a chat member (#chatmember) that is under certain restrictions in the chat.
// Supergroups only.
type ChatMemberRestricted struct {
	// Status - The member's status in the chat, always “restricted”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`

	// IsMember - True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member"`

	// CanChangeInfo - True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// CanInviteUsers - True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// CanPinMessages - True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages"`

	// CanSendMessages - True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMessages bool `json:"can_send_messages"`

	// CanSendMediaMessages - True, if the user is allowed to send audios, documents, photos, videos, video
	// notes and voice notes
	CanSendMediaMessages bool `json:"can_send_media_messages"`

	// CanSendPolls - True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls"`

	// CanSendOtherMessages - True, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages"`

	// CanAddWebPagePreviews - True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`

	// UntilDate - Date when restrictions will be lifted for this user; unix time
	UntilDate int `json:"until_date"`
}

func (c *ChatMemberRestricted) MemberStatus() string {
	return c.Status
}

// ChatMemberLeft - Represents a chat member (#chatmember) that isn't currently a member of the chat,
// but may join it themselves.
type ChatMemberLeft struct {
	// Status - The member's status in the chat, always “left”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`
}

func (c *ChatMemberLeft) MemberStatus() string {
	return c.Status
}

// ChatMemberBanned - Represents a chat member (#chatmember) that was banned in the chat and can't return to
// the chat or view chat messages.
type ChatMemberBanned struct {
	// Status - The member's status in the chat, always “kicked”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`

	// UntilDate - Date when restrictions will be lifted for this user; unix time
	UntilDate int `json:"until_date"`
}

func (c *ChatMemberBanned) MemberStatus() string {
	return c.Status
}

// ChatMemberUpdated - This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	// Chat - Chat the user belongs to
	Chat Chat `json:"chat"`

	// From - Performer of the action, which resulted in the change
	From User `json:"from"`

	// Date - Date the change was done in Unix time
	Date int `json:"date"`

	// OldChatMember - Previous information about the chat member
	OldChatMember ChatMember `json:"old_chat_member"`

	// NewChatMember - New information about the chat member
	NewChatMember ChatMember `json:"new_chat_member"`

	// InviteLink - Optional. Chat invite link, which was used by the user to join the chat; for joining by invite
	// link events only.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// ChatPermissions - Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// CanSendMessages - Optional. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// CanSendMediaMessages - Optional. True, if the user is allowed to send audios, documents, photos, videos,
	// video notes and voice notes, implies can_send_messages
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

	// CanSendPolls - Optional. True, if the user is allowed to send polls, implies can_send_messages
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// CanSendOtherMessages - Optional. True, if the user is allowed to send animations, games, stickers and use
	// inline bots, implies can_send_media_messages
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// CanAddWebPagePreviews - Optional. True, if the user is allowed to add web page previews to their messages,
	// implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// CanChangeInfo - Optional. True, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// CanInviteUsers - Optional. True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// CanPinMessages - Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// ChatLocation - Represents a location to which a chat is connected.
type ChatLocation struct {
	// Location - The location to which the supergroup is connected. Can't be a live location.
	Location Location `json:"location"`

	// Address - Location address; 1-64 characters, as defined by the chat owner
	Address string `json:"address"`
}

// BotCommand - This object represents a bot command.
type BotCommand struct {
	// Command - Text of the command, 1-32 characters. Can contain only lowercase English letters,
	// digits and underscores.
	Command string `json:"command"`

	// Description - Description of the command, 3-256 characters.
	Description string `json:"description"`
}

// BotCommandScope - This object represents the scope to which bot commands are applied. Currently,
// the following 7 scopes are supported:
// BotCommandScopeDefault
// BotCommandScopeAllPrivateChats
// BotCommandScopeAllGroupChats
// BotCommandScopeAllChatAdministrators
// BotCommandScopeChat
// BotCommandScopeChatAdministrators
// BotCommandScopeChatMember
type BotCommandScope interface {
	ScopeType() string
}

/* TODO: Check if needed
type botCommandScopeData struct {
	Data BotCommandScope
}

func (b *botCommandScopeData) UnmarshalJSON(bytes []byte) error {
	var scopeType struct {
		Type string `json:"type"`
	}

	err := json.Unmarshal(bytes, &scopeType)
	if err != nil {
		return err
	}

	switch scopeType.Type {
	case "default":
		var bcs BotCommandScopeDefault
		err = json.Unmarshal(bytes, &bcs)
		b.Data = bcs
	case "all_private_chats":
		var bcs BotCommandScopeAllPrivateChats
		err = json.Unmarshal(bytes, &bcs)
		b.Data = bcs
	case "all_group_chats":
		var bcs BotCommandScopeAllGroupChats
		err = json.Unmarshal(bytes, &bcs)
		b.Data = bcs
	case "all_chat_administrators":
		var bcs BotCommandScopeAllChatAdministrators
		err = json.Unmarshal(bytes, &bcs)
		b.Data = bcs
	case "chat":
		var bcs BotCommandScopeChat
		err = json.Unmarshal(bytes, &bcs)
		b.Data = bcs
	case "chat_administrators":
		var bcs BotCommandScopeChatAdministrators
		err = json.Unmarshal(bytes, &bcs)
		b.Data = bcs
	case "chat_member":
		var bcs BotCommandScopeChatMember
		err = json.Unmarshal(bytes, &bcs)
		b.Data = bcs
	default:
		return errors.New("unknown member status")
	}

	return err
}
*/

// BotCommandScopeDefault - Represents the default scope (#botcommandscope) of bot commands.
// Default commands are used if no commands with a narrower scope (#determining-list-of-commands)
// are specified for the user.
type BotCommandScopeDefault struct {
	// Type - Scope type, must be default
	Type string `json:"type"`
}

func (b *BotCommandScopeDefault) ScopeType() string {
	return b.Type
}

// BotCommandScopeAllPrivateChats - Represents the scope (#botcommandscope) of bot commands,
// covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	// Type - Scope type, must be all_private_chats
	Type string `json:"type"`
}

func (b *BotCommandScopeAllPrivateChats) ScopeType() string {
	return b.Type
}

// BotCommandScopeAllGroupChats - Represents the scope (#botcommandscope) of bot commands,
// covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	// Type - Scope type, must be all_group_chats
	Type string `json:"type"`
}

func (b *BotCommandScopeAllGroupChats) ScopeType() string {
	return b.Type
}

// BotCommandScopeAllChatAdministrators - Represents the scope (#botcommandscope) of bot commands,
// covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {
	// Type - Scope type, must be all_chat_administrators
	Type string `json:"type"`
}

func (b *BotCommandScopeAllChatAdministrators) ScopeType() string {
	return b.Type
}

// ChatID - Represents chat ID as int64 or string
type ChatID struct {
	ID       int64
	Username string
}

func (c *ChatID) String() string {
	if c.Username != "" {
		return c.Username
	}

	return fmt.Sprintf("%d", c.ID)
}

func (c *ChatID) MarshalJSON() ([]byte, error) {
	if c.Username != "" {
		return json.Marshal(c.Username)
	}

	return json.Marshal(fmt.Sprintf("%d", c.ID))
}

// BotCommandScopeChat - Represents the scope (#botcommandscope) of bot commands, covering a specific chat.
type BotCommandScopeChat struct {
	// Type - Scope type, must be chat
	Type string `json:"type"`

	// ChatID - Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername)
	ChatID ChatID `json:"chat_id"`
}

func (b *BotCommandScopeChat) ScopeType() string {
	return b.Type
}

// BotCommandScopeChatAdministrators - Represents the scope (#botcommandscope) of bot commands,
// covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	// Type - Scope type, must be chat_administrators
	Type string `json:"type"`

	// ChatID - Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername)
	ChatID ChatID `json:"chat_id"`
}

func (b *BotCommandScopeChatAdministrators) ScopeType() string {
	return b.Type
}

// BotCommandScopeChatMember - Represents the scope (#botcommandscope) of bot commands,
// covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	// Type - Scope type, must be chat_member
	Type string `json:"type"`

	// ChatID - Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int `json:"user_id"`
}

func (b *BotCommandScopeChatMember) ScopeType() string {
	return b.Type
}

// ResponseParameters - Contains information about why a request was unsuccessful.
type ResponseParameters struct {
	// MigrateToChatID - Optional. The group has been migrated to a supergroup with the specified identifier.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// RetryAfter - Optional. In case of exceeding flood control, the number of seconds left to wait
	// before the request can be repeated
	RetryAfter int `json:"retry_after,omitempty"`
}

// fileCompatible - Represents types that can be send as files
type fileCompatible interface {
	fileParameters() map[string]*os.File
}

// InputFile - This object represents the contents of a file to be uploaded. Must be posted using
// multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile struct {
	File       *os.File
	FileID     string
	URL        string
	needAttach bool
}

func (i *InputFile) MarshalJSON() ([]byte, error) {
	if i.File != nil {
		if i.needAttach {
			return json.Marshal(attachFile + i.File.Name())
		}
		return json.Marshal("")
	}

	if i.FileID != "" {
		return json.Marshal(i.FileID)
	}

	return json.Marshal(i.URL)
}

func (i *InputFile) String() string {
	return fmt.Sprintf("{File: %v ID: %q URL: %q NeedAttach: %t}", i.File, i.FileID, i.URL, i.needAttach)
}

// InputMedia - This object represents the content of a media message to be sent. It should be one of:
// InputMediaAnimation
// InputMediaDocument
// InputMediaAudio
// InputMediaPhoto
// InputMediaVideo
type InputMedia interface {
	MediaType() string
	fileCompatible
}

// InputMediaPhoto - Represents a photo to be sent.
type InputMediaPhoto struct {
	// Type - Type of the result, must be photo
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>”
	// to upload a new one using multipart/form-data under <file_attach_name> name.
	// More info on Sending Files » (#sending-files)
	Media InputFile `json:"media"`

	// Caption - Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the photo caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
}

func (i *InputMediaPhoto) MediaType() string {
	return "photo"
}

func (i *InputMediaPhoto) fileParameters() map[string]*os.File {
	i.Media.needAttach = true
	return map[string]*os.File{
		"media": i.Media.File,
	}
}

// InputMediaVideo - Represents a video to be sent.
type InputMediaVideo struct {
	// Type - Type of the result, must be video
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to
	// upload a new one using multipart/form-data under <file_attach_name> name.
	// More info on Sending Files » (#sending-files)
	Media InputFile `json:"media"`

	// Thumb - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height
	// should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be
	// reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail
	// was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » (#sending-files)
	Thumb *InputFile `json:"thumb,omitempty"`

	// Caption - Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the video caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Width - Optional. Video width
	Width int `json:"width,omitempty"`

	// Height - Optional. Video height
	Height int `json:"height,omitempty"`

	// Duration - Optional. Video duration
	Duration int `json:"duration,omitempty"`

	// SupportsStreaming - Optional. Pass True, if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
}

func (i *InputMediaVideo) MediaType() string {
	return "video"
}

func (i *InputMediaVideo) fileParameters() map[string]*os.File {
	fp := make(map[string]*os.File)

	i.Media.needAttach = true
	fp["media"] = i.Media.File
	if i.Thumb != nil {
		i.Thumb.needAttach = true
		fp["thumb"] = i.Thumb.File
	}

	return fp
}

// InputMediaAnimation - Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	// Type - Type of the result, must be animation
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to
	// upload a new one using multipart/form-data under <file_attach_name> name.
	// More info on Sending Files » (#sending-files)
	Media InputFile `json:"media"`

	// Thumb - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you
	// can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data
	// under <file_attach_name>. More info on Sending Files » (#sending-files)
	Thumb *InputFile `json:"thumb,omitempty"`

	// Caption - Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the animation caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Width - Optional. Animation width
	Width int `json:"width,omitempty"`

	// Height - Optional. Animation height
	Height int `json:"height,omitempty"`

	// Duration - Optional. Animation duration
	Duration int `json:"duration,omitempty"`
}

func (i *InputMediaAnimation) MediaType() string {
	return "animation"
}

func (i *InputMediaAnimation) fileParameters() map[string]*os.File {
	fp := make(map[string]*os.File)

	i.Media.needAttach = true
	fp["media"] = i.Media.File
	if i.Thumb != nil {
		i.Thumb.needAttach = true
		fp["thumb"] = i.Thumb.File
	}

	return fp
}

// InputMediaAudio - Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	// Type - Type of the result, must be audio
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload
	// a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files » (#sending-files)
	Media InputFile `json:"media"`

	// Thumb - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height
	// should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was
	// uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » (#sending-files)
	Thumb *InputFile `json:"thumb,omitempty"`

	// Caption - Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the audio caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Duration - Optional. Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`

	// Performer - Optional. Performer of the audio
	Performer string `json:"performer,omitempty"`

	// Title - Optional. Title of the audio
	Title string `json:"title,omitempty"`
}

func (i *InputMediaAudio) MediaType() string {
	return "audio"
}

func (i *InputMediaAudio) fileParameters() map[string]*os.File {
	fp := make(map[string]*os.File)

	i.Media.needAttach = true
	fp["media"] = i.Media.File
	if i.Thumb != nil {
		i.Thumb.needAttach = true
		fp["thumb"] = i.Thumb.File
	}

	return fp
}

// InputMediaDocument - Represents a general file to be sent.
type InputMediaDocument struct {
	// Type - Type of the result, must be document
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload
	// a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files » (#sending-files)
	Media InputFile `json:"media"`

	// Thumb - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height
	// should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was
	// uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » (#sending-files)
	Thumb *InputFile `json:"thumb,omitempty"`

	// Caption - Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the document caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// DisableContentTypeDetection - Optional. Disables automatic server-side content type detection for files
	// uploaded using multipart/form-data. Always true, if the document is sent as part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}

func (i *InputMediaDocument) MediaType() string {
	return "document"
}

func (i *InputMediaDocument) fileParameters() map[string]*os.File {
	fp := make(map[string]*os.File)

	i.Media.needAttach = true
	fp["media"] = i.Media.File
	if i.Thumb != nil {
		i.Thumb.needAttach = true
		fp["thumb"] = i.Thumb.File
	}

	return fp
}

// TODO: Continue checking

// Sticker - This object represents a sticker.
type Sticker struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width - Sticker width
	Width int `json:"width"`

	// Height - Sticker height
	Height int `json:"height"`

	// IsAnimated - True, if the sticker is animated (https://telegram.org/blog/animated-stickers)
	IsAnimated bool `json:"is_animated"`

	// Thumb - Optional. Sticker thumbnail in the .WEBP or .JPG format
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Emoji - Optional. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`

	// SetName - Optional. Name of the sticker set to which the sticker belongs
	SetName string `json:"set_name,omitempty"`

	// MaskPosition - Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// FileSize - Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// StickerSet - This object represents a sticker set.
type StickerSet struct {
	// Name - Sticker set name
	Name string `json:"name"`

	// Title - Sticker set title
	Title string `json:"title"`

	// IsAnimated - True, if the sticker set contains animated stickers (https://telegram.org/blog/animated-stickers)
	IsAnimated bool `json:"is_animated"`

	// ContainsMasks - True, if the sticker set contains masks
	ContainsMasks bool `json:"contains_masks"`

	// Stickers - List of all set stickers
	Stickers []Sticker `json:"stickers"`

	// Thumb - Optional. Sticker set thumbnail in the .WEBP or .TGS format
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// MaskPosition - This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// Point - The part of the face relative to which the mask should be placed. One of “forehead”,
	// “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`

	// XShift - Shift by X-axis measured in widths of the mask scaled to the face size, from left to right.
	// For example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float64 `json:"x_shift"`

	// YShift - Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom.
	// For example, 1.0 will place the mask just below the default mask position.
	YShift float64 `json:"y_shift"`

	// Scale - Mask scaling coefficient. For example, 2.0 means double size.
	Scale float64 `json:"scale"`
}

// InlineQuery - This object represents an incoming inline query. When the user sends an empty query,
// your bot could return some default or trending results.
type InlineQuery struct {
	// ID - Unique identifier for this query
	ID string `json:"id"`

	// From - Sender
	From User `json:"from"`

	// Query - Text of the query (up to 256 characters)
	Query string `json:"query"`

	// Offset - Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`

	// ChatType - Optional. Type of the chat, from which the inline query was sent. Can be either “sender” for
	// a private chat with the inline query sender, “private”, “group”, “supergroup”, or “channel”.
	// The chat type should be always known for requests sent from official clients and most third-party clients,
	// unless the request was sent from a secret chat
	ChatType string `json:"chat_type,omitempty"`

	// Location - Optional. Sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`
}

// FIXME

// InlineQueryResult - This object represents one result of an inline query. Telegram clients currently
// support results of the following 20 types:
type InlineQueryResult struct {
}

// InlineQueryResultArticle - Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	// Type - Type of the result, must be article
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// Title - Title of the result
	Title string `json:"title"`

	// InputMessageContent - Content of the message to be sent
	InputMessageContent InputMessageContent `json:"input_message_content"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// URL - Optional. URL of the result
	URL string `json:"url,omitempty"`

	// HideURL - Optional. Pass True, if you don't want the URL to be shown in the message
	HideURL bool `json:"hide_url,omitempty"`

	// Description - Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// ThumbURL - Optional. URL of the thumbnail for the result
	ThumbURL string `json:"thumb_url,omitempty"`

	// ThumbWidth - Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`

	// ThumbHeight - Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultPhoto - Represents a link to a photo. By default, this photo will be sent by the user with
// optional caption. Alternatively, you can use input_message_content to send a message with the specified
// content instead of the photo.
type InlineQueryResultPhoto struct {
	// Type - Type of the result, must be photo
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// PhotoURL - A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
	PhotoURL string `json:"photo_url"`

	// ThumbURL - URL of the thumbnail for the photo
	ThumbURL string `json:"thumb_url"`

	// PhotoWidth - Optional. Width of the photo
	PhotoWidth int `json:"photo_width,omitempty"`

	// PhotoHeight - Optional. Height of the photo
	PhotoHeight int `json:"photo_height,omitempty"`

	// Title - Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Description - Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Caption - Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the photo caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the photo
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultGif - Represents a link to an animated GIF file. By default, this animated GIF file will be
// sent by the user with optional caption. Alternatively, you can use input_message_content to send a message
// with the specified content instead of the animation.
type InlineQueryResultGif struct {
	// Type - Type of the result, must be gif
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// GifURL - A valid URL for the GIF file. File size must not exceed 1MB
	GifURL string `json:"gif_url"`

	// GifWidth - Optional. Width of the GIF
	GifWidth int `json:"gif_width,omitempty"`

	// GifHeight - Optional. Height of the GIF
	GifHeight int `json:"gif_height,omitempty"`

	// GifDuration - Optional. Duration of the GIF
	GifDuration int `json:"gif_duration,omitempty"`

	// ThumbURL - URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbURL string `json:"thumb_url"`

	// ThumbMimeType - Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”,
	// or “video/mp4”. Defaults to “image/jpeg”
	ThumbMimeType string `json:"thumb_mime_type,omitempty"`

	// Title - Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Caption - Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultMpeg4Gif - Represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	// Type - Type of the result, must be mpeg4_gif
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// Mpeg4URL - A valid URL for the MP4 file. File size must not exceed 1MB
	Mpeg4URL string `json:"mpeg4_url"`

	// Mpeg4Width - Optional. Video width
	Mpeg4Width int `json:"mpeg4_width,omitempty"`

	// Mpeg4Height - Optional. Video height
	Mpeg4Height int `json:"mpeg4_height,omitempty"`

	// Mpeg4Duration - Optional. Video duration
	Mpeg4Duration int `json:"mpeg4_duration,omitempty"`

	// ThumbURL - URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbURL string `json:"thumb_url"`

	// ThumbMimeType - Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”,
	// or “video/mp4”. Defaults to “image/jpeg”
	ThumbMimeType string `json:"thumb_mime_type,omitempty"`

	// Title - Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Caption - Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the video animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultVideo - Represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
}

// InlineQueryResultAudio - Represents a link to an MP3 audio file. By default, this audio file will be
// sent by the user. Alternatively, you can use input_message_content to send a message with the specified
// content instead of the audio.
type InlineQueryResultAudio struct {
	// Type - Type of the result, must be audio
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// AudioURL - A valid URL for the audio file
	AudioURL string `json:"audio_url"`

	// Title - Title
	Title string `json:"title"`

	// Caption - Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the audio caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can
	// be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Performer - Optional. Performer
	Performer string `json:"performer,omitempty"`

	// AudioDuration - Optional. Audio duration in seconds
	AudioDuration int `json:"audio_duration,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the audio
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultVoice - Represents a link to a voice recording in an .OGG container encoded with OPUS.
// By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content
// to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	// Type - Type of the result, must be voice
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// VoiceURL - A valid URL for the voice recording
	VoiceURL string `json:"voice_url"`

	// Title - Recording title
	Title string `json:"title"`

	// Caption - Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the voice message caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// VoiceDuration - Optional. Recording duration in seconds
	VoiceDuration int `json:"voice_duration,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the voice recording
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultDocument - Represents a link to a file. By default, this file will be sent by the user
// with an optional caption. Alternatively, you can use input_message_content to send a message with the
// specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	// Type - Type of the result, must be document
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// Title - Title for the result
	Title string `json:"title"`

	// Caption - Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the document caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// DocumentURL - A valid URL for the file
	DocumentURL string `json:"document_url"`

	// MimeType - Mime type of the content of the file, either “application/pdf” or “application/zip”
	MimeType string `json:"mime_type"`

	// Description - Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the file
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`

	// ThumbURL - Optional. URL of the thumbnail (jpeg only) for the file
	ThumbURL string `json:"thumb_url,omitempty"`

	// ThumbWidth - Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`

	// ThumbHeight - Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultLocation - Represents a location on a map. By default, the location will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content
// instead of the location.
type InlineQueryResultLocation struct {
	// Type - Type of the result, must be location
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// Latitude - Location latitude in degrees
	Latitude float64 `json:"latitude"`

	// Longitude - Location longitude in degrees
	Longitude float64 `json:"longitude"`

	// Title - Location title
	Title string `json:"title"`

	// HorizontalAccuracy - Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// LivePeriod - Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`

	// Heading - Optional. For live locations, a direction in which the user is moving, in degrees. Must be
	// between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`

	// ProximityAlertRadius - Optional. For live locations, a maximum distance for proximity alerts about
	// approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the location
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`

	// ThumbURL - Optional. URL of the thumbnail for the result
	ThumbURL string `json:"thumb_url,omitempty"`

	// ThumbWidth - Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`

	// ThumbHeight - Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultVenue - Represents a venue. By default, the venue will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	// Type - Type of the result, must be venue
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// Latitude - Latitude of the venue location in degrees
	Latitude float64 `json:"latitude"`

	// Longitude - Longitude of the venue location in degrees
	Longitude float64 `json:"longitude"`

	// Title - Title of the venue
	Title string `json:"title"`

	// Address - Address of the venue
	Address string `json:"address"`

	// FoursquareID - Optional. Foursquare identifier of the venue if known
	FoursquareID string `json:"foursquare_id,omitempty"`

	// FoursquareType - Optional. Foursquare type of the venue, if known. (For example,
	// “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// GooglePlaceID - Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// GooglePlaceType - Optional. Google Places type of the venue. (See supported
	// types (https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType string `json:"google_place_type,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the venue
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`

	// ThumbURL - Optional. URL of the thumbnail for the result
	ThumbURL string `json:"thumb_url,omitempty"`

	// ThumbWidth - Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`

	// ThumbHeight - Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultContact - Represents a contact with a phone number. By default,
// this contact will be sent by the user. Alternatively, you can use input_message_content to send
// a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	// Type - Type of the result, must be contact
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// PhoneNumber - Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// FirstName - Contact's first name
	FirstName string `json:"first_name"`

	// LastName - Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Vcard - Optional. Additional data about the contact in the form of
	// a vCard (https://en.wikipedia.org/wiki/VCard), 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the contact
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`

	// ThumbURL - Optional. URL of the thumbnail for the result
	ThumbURL string `json:"thumb_url,omitempty"`

	// ThumbWidth - Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`

	// ThumbHeight - Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultGame - Represents a Game (#games).
type InlineQueryResultGame struct {
	// Type - Type of the result, must be game
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// GameShortName - Short name of the game
	GameShortName string `json:"game_short_name"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// InlineQueryResultCachedPhoto - Represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	// Type - Type of the result, must be photo
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// PhotoFileID - A valid file identifier of the photo
	PhotoFileID string `json:"photo_file_id"`

	// Title - Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Description - Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Caption - Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the photo caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the photo
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedGif - Represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	// Type - Type of the result, must be gif
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// GifFileID - A valid file identifier for the GIF file
	GifFileID string `json:"gif_file_id"`

	// Title - Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Caption - Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedMpeg4Gif - Represents a link to a video animation
// (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will
// be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message
// with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	// Type - Type of the result, must be mpeg4_gif
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// Mpeg4FileID - A valid file identifier for the MP4 file
	Mpeg4FileID string `json:"mpeg4_file_id"`

	// Title - Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Caption - Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the video animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedSticker - Represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send
// a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	// Type - Type of the result, must be sticker
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// StickerFileID - A valid file identifier of the sticker
	StickerFileID string `json:"sticker_file_id"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the sticker
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedDocument - Represents a link to a file stored on the Telegram servers.
// By default, this file will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	// Type - Type of the result, must be document
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// Title - Title for the result
	Title string `json:"title"`

	// DocumentFileID - A valid file identifier for the file
	DocumentFileID string `json:"document_file_id"`

	// Description - Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Caption - Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the document caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the file
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVideo - Represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	// Type - Type of the result, must be video
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// VideoFileID - A valid file identifier for the video file
	VideoFileID string `json:"video_file_id"`

	// Title - Title for the result
	Title string `json:"title"`

	// Description - Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Caption - Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the video caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the video
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVoice - Represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send
// a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	// Type - Type of the result, must be voice
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// VoiceFileID - A valid file identifier for the voice message
	VoiceFileID string `json:"voice_file_id"`

	// Title - Voice message title
	Title string `json:"title"`

	// Caption - Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the voice message caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the voice message
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedAudio - Represents a link to an MP3 audio file stored on the Telegram servers.
// By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send
// a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	// Type - Type of the result, must be audio
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// AudioFileID - A valid file identifier for the audio file
	AudioFileID string `json:"audio_file_id"`

	// Caption - Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the audio caption. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can
	// be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (/bots#inline-keyboards-and-on-the-fly-updating) attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the audio
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InputMessageContent - This object represents the content of a message to be sent as a result of an
// inline query. Telegram clients currently support the following 5 types:
type InputMessageContent struct {
}

// InputTextMessageContent - Represents the content (#inputmessagecontent) of a text message to be sent
// as the result of an inline query.
type InputTextMessageContent struct {
	// MessageText - Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`

	// ParseMode - Optional. Mode for parsing entities in the message text. See formatting
	// options (#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Entities - Optional. List of special entities that appear in message text, which can be specified
	// instead of parse_mode
	Entities []MessageEntity `json:"entities,omitempty"`

	// DisableWebPagePreview - Optional. Disables link previews for links in the sent message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

// InputLocationMessageContent - Represents the content (#inputmessagecontent) of a location message to be sent as
// the result of an inline query.
type InputLocationMessageContent struct {
	// Latitude - Latitude of the location in degrees
	Latitude float64 `json:"latitude"`

	// Longitude - Longitude of the location in degrees
	Longitude float64 `json:"longitude"`

	// HorizontalAccuracy - Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// LivePeriod - Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`

	// Heading - Optional. For live locations, a direction in which the user is moving, in degrees.
	// Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`

	// ProximityAlertRadius - Optional. For live locations, a maximum distance for proximity alerts about
	// approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

// InputVenueMessageContent - Represents the content (#inputmessagecontent) of a venue message to be sent as
// the result of an inline query.
type InputVenueMessageContent struct {
	// Latitude - Latitude of the venue in degrees
	Latitude float64 `json:"latitude"`

	// Longitude - Longitude of the venue in degrees
	Longitude float64 `json:"longitude"`

	// Title - Name of the venue
	Title string `json:"title"`

	// Address - Address of the venue
	Address string `json:"address"`

	// FoursquareID - Optional. Foursquare identifier of the venue, if known
	FoursquareID string `json:"foursquare_id,omitempty"`

	// FoursquareType - Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// GooglePlaceID - Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// GooglePlaceType - Optional. Google Places type of the venue.
	// (See supported types (https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// InputContactMessageContent - Represents the content (#inputmessagecontent) of a contact message
// to be sent as the result of an inline query.
type InputContactMessageContent struct {
	// PhoneNumber - Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// FirstName - Contact's first name
	FirstName string `json:"first_name"`

	// LastName - Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Vcard - Optional. Additional data about the contact in the form of a
	// vCard (https://en.wikipedia.org/wiki/VCard), 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`
}

// InputInvoiceMessageContent - Represents the content (#inputmessagecontent) of an invoice message to be sent
// as the result of an inline query.
type InputInvoiceMessageContent struct {
	// Title - Product name, 1-32 characters
	Title string `json:"title"`

	// Description - Product description, 1-255 characters
	Description string `json:"description"`

	// Payload - Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for
	// your internal processes.
	Payload string `json:"payload"`

	// ProviderToken - Payment provider token, obtained via Botfather (https://t.me/botfather)
	ProviderToken string `json:"provider_token"`

	// Currency - Three-letter ISO 4217 currency code, see more on currencies (/bots/payments#supported-currencies)
	Currency string `json:"currency"`

	// Prices - Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.)
	Prices []LabeledPrice `json:"prices"`

	// MaxTipAmount - Optional. The maximum accepted amount for tips in the smallest units of the
	// currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145.
	// See the exp parameter in currencies.json (https://core.telegram.org/bots/payments/currencies.json), it shows
	// the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	MaxTipAmount int `json:"max_tip_amount,omitempty"`

	// SuggestedTipAmounts - Optional. A JSON-serialized array of suggested amounts of tip in the smallest units
	// of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified.
	// The suggested tip amounts must be positive, passed in a strictly increased order and must not
	// exceed max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`

	// ProviderData - Optional. A JSON-serialized object for data about the invoice, which will be
	// shared with the payment provider. A detailed description of the required fields should be
	// provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// PhotoURL - Optional. URL of the product photo for the invoice. Can be a photo of the goods or a
	// marketing image for a service. People like it better when they see what they are paying for.
	PhotoURL string `json:"photo_url,omitempty"`

	// PhotoSize - Optional. Photo size
	PhotoSize int `json:"photo_size,omitempty"`

	// PhotoWidth - Optional. Photo width
	PhotoWidth int `json:"photo_width,omitempty"`

	// PhotoHeight - Optional. Photo height
	PhotoHeight int `json:"photo_height,omitempty"`

	// NeedName - Optional. Pass True, if you require the user's full name to complete the order
	NeedName bool `json:"need_name,omitempty"`

	// NeedPhoneNumber - Optional. Pass True, if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// NeedEmail - Optional. Pass True, if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email,omitempty"`

	// NeedShippingAddress - Optional. Pass True, if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// SendPhoneNumberToProvider - Optional. Pass True, if user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// SendEmailToProvider - Optional. Pass True, if user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// IsFlexible - Optional. Pass True, if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"`
}

// ChosenInlineResult - Represents a result (#inlinequeryresult) of an inline query that was chosen
// by the user and sent to their chat partner.
type ChosenInlineResult struct {
	// ResultID - The unique identifier for the result that was chosen
	ResultID string `json:"result_id"`

	// From - The user that chose the result
	From User `json:"from"`

	// Location - Optional. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`

	// InlineMessageID - Optional. Identifier of the sent inline message. Available only if there is an
	// inline keyboard (#inlinekeyboardmarkup) attached to the message. Will be also received in callback
	// queries (#callbackquery) and can be used to edit (#updating-messages) the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Query - The query that was used to obtain the result
	Query string `json:"query"`
}

// LabeledPrice - This object represents a portion of the price for goods or services.
type LabeledPrice struct {
	// Label - Portion label
	Label string `json:"label"`

	// Amount - Price of the product in the smallest units of the currency (/bots/payments#supported-currencies)
	// (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter
	// in currencies.json (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits
	// past the decimal point for each currency (2 for the majority of currencies).
	Amount int `json:"amount"`
}

// Invoice - This object contains basic information about an invoice.
type Invoice struct {
	// Title - Product name
	Title string `json:"title"`

	// Description - Product description
	Description string `json:"description"`

	// StartParameter - Unique bot deep-linking parameter that can be used to generate this invoice
	StartParameter string `json:"start_parameter"`

	// Currency - Three-letter ISO 4217 currency (/bots/payments#supported-currencies) code
	Currency string `json:"currency"`

	// TotalAmount - Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the
	// decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}

// ShippingAddress - This object represents a shipping address.
type ShippingAddress struct {
	// CountryCode - ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`

	// State - State, if applicable
	State string `json:"state"`

	// City - City
	City string `json:"city"`

	// StreetLine1 - First line for the address
	StreetLine1 string `json:"street_line1"`

	// StreetLine2 - Second line for the address
	StreetLine2 string `json:"street_line2"`

	// PostCode - Address post code
	PostCode string `json:"post_code"`
}

// OrderInfo - This object represents information about an order.
type OrderInfo struct {
	// Name - Optional. User name
	Name string `json:"name,omitempty"`

	// PhoneNumber - Optional. User's phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// Email - Optional. User email
	Email string `json:"email,omitempty"`

	// ShippingAddress - Optional. User shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// ShippingOption - This object represents one shipping option.
type ShippingOption struct {
	// ID - Shipping option identifier
	ID string `json:"id"`

	// Title - Option title
	Title string `json:"title"`

	// Prices - List of price portions
	Prices []LabeledPrice `json:"prices"`
}

// SuccessfulPayment - This object contains basic information about a successful payment.
type SuccessfulPayment struct {
	// Currency - Three-letter ISO 4217 currency (/bots/payments#supported-currencies) code
	Currency string `json:"currency"`

	// TotalAmount - Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the
	// decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload - Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// ShippingOptionID - Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo - Optional. Order info provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`

	// TelegramPaymentChargeID - Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

	// ProviderPaymentChargeID - Provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}

// ShippingQuery - This object contains information about an incoming shipping query.
type ShippingQuery struct {
	// ID - Unique query identifier
	ID string `json:"id"`

	// From - User who sent the query
	From User `json:"from"`

	// InvoicePayload - Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// ShippingAddress - User specified shipping address
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery - This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	// ID - Unique query identifier
	ID string `json:"id"`

	// From - User who sent the query
	From User `json:"from"`

	// Currency - Three-letter ISO 4217 currency (/bots/payments#supported-currencies) code
	Currency string `json:"currency"`

	// TotalAmount - Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the
	// decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload - Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// ShippingOptionID - Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo - Optional. Order info provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

// PassportData - Contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Data - Array with information about documents and other Telegram Passport elements that was shared with the bot
	Data []EncryptedPassportElement `json:"data"`

	// Credentials - Encrypted credentials required to decrypt the data
	Credentials EncryptedCredentials `json:"credentials"`
}

// PassportFile - This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport
// files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// FileSize - File size
	FileSize int `json:"file_size"`

	// FileDate - Unix time when the file was uploaded
	FileDate int `json:"file_date"`
}

// EncryptedPassportElement - Contains information about documents or other Telegram Passport elements shared
// with the bot by the user.
type EncryptedPassportElement struct {
	// Type - Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”,
	// “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”,
	// “temporary_registration”, “phone_number”, “email”.
	Type string `json:"type"`

	// Data - Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for
	// “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types.
	// Can be decrypted and verified using the accompanying EncryptedCredentials (#encryptedcredentials).
	Data string `json:"data,omitempty"`

	// PhoneNumber - Optional. User's verified phone number, available only for “phone_number” type
	PhoneNumber string `json:"phone_number,omitempty"`

	// Email - Optional. User's verified email address, available only for “email” type
	Email string `json:"email,omitempty"`

	// Files - Optional. Array of encrypted files with documents provided by the user, available for “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be
	// decrypted and verified using the accompanying EncryptedCredentials (#encryptedcredentials).
	Files []PassportFile `json:"files,omitempty"`

	// FrontSide - Optional. Encrypted file with the front side of the document, provided by the user. Available for
	// “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified
	// using the accompanying EncryptedCredentials (#encryptedcredentials).
	FrontSide *PassportFile `json:"front_side,omitempty"`

	// ReverseSide - Optional. Encrypted file with the reverse side of the document, provided by the user.
	// Available for “driver_license” and “identity_card”. The file can be decrypted and verified using the
	// accompanying EncryptedCredentials (#encryptedcredentials).
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`

	// Selfie - Optional. Encrypted file with the selfie of the user holding a document, provided by the user;
	// available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted
	// and verified using the accompanying EncryptedCredentials (#encryptedcredentials).
	Selfie *PassportFile `json:"selfie,omitempty"`

	// Translation - Optional. Array of encrypted files with translated versions of documents provided by the user.
	// Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”,
	// “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration”
	// types. Files can be decrypted and verified using the accompanying EncryptedCredentials (#encryptedcredentials).
	Translation []PassportFile `json:"translation,omitempty"`

	// Hash - Base64-encoded element hash for using in
	// PassportElementErrorUnspecified (#passportelementerrorunspecified)
	Hash string `json:"hash"`
}

// EncryptedCredentials - Contains data required for decrypting and authenticating EncryptedPassportElement
// (#encryptedpassportelement). See the Telegram Passport Documentation
// (https://core.telegram.org/passport#receiving-information) for a complete description of
// the data decryption and authentication processes.
type EncryptedCredentials struct {
	// Data - Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and
	// secrets required for EncryptedPassportElement (#encryptedpassportelement) decryption and authentication
	Data string `json:"data"`

	// Hash - Base64-encoded data hash for data authentication
	Hash string `json:"hash"`

	// Secret - Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
	Secret string `json:"secret"`
}

// FIXME

// PassportElementError - This object represents an error in the Telegram Passport element which was submitted
// that should be resolved by the user. It should be one of:
type PassportElementError struct {
}

// PassportElementErrorDataField - Represents an issue in one of the data fields that was provided by the user.
// The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	// Source - Error source, must be data
	Source string `json:"source"`

	// Type - The section of the user's Telegram Passport which has the error, one of “personal_details”, “passport”,
	// “driver_license”, “identity_card”, “internal_passport”, “address”
	Type string `json:"type"`

	// FieldName - Name of the data field which has the error
	FieldName string `json:"field_name"`

	// DataHash - Base64-encoded data hash
	DataHash string `json:"data_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// PassportElementErrorFrontSide - Represents an issue with the front side of a document. The error is considered
// resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	// Source - Error source, must be front_side
	Source string `json:"source"`

	// Type - The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”,
	// “identity_card”, “internal_passport”
	Type string `json:"type"`

	// FileHash - Base64-encoded hash of the file with the front side of the document
	FileHash string `json:"file_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// PassportElementErrorReverseSide - Represents an issue with the reverse side of a document. The error is
// considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	// Source - Error source, must be reverse_side
	Source string `json:"source"`

	// Type - The section of the user's Telegram Passport which has the issue, one of “driver_license”, “identity_card”
	Type string `json:"type"`

	// FileHash - Base64-encoded hash of the file with the reverse side of the document
	FileHash string `json:"file_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// PassportElementErrorSelfie - Represents an issue with the selfie with a document. The error is considered
// resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	// Source - Error source, must be selfie
	Source string `json:"source"`

	// Type - The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”,
	// “identity_card”, “internal_passport”
	Type string `json:"type"`

	// FileHash - Base64-encoded hash of the file with the selfie
	FileHash string `json:"file_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// PassportElementErrorFile - Represents an issue with a document scan. The error is considered resolved when
// the file with the document scan changes.
type PassportElementErrorFile struct {
	// Source - Error source, must be file
	Source string `json:"source"`

	// Type - The section of the user's Telegram Passport which has the issue, one of “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// FileHash - Base64-encoded file hash
	FileHash string `json:"file_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// PassportElementErrorFiles - Represents an issue with a list of scans. The error is considered resolved
// when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	// Source - Error source, must be files
	Source string `json:"source"`

	// Type - The section of the user's Telegram Passport which has the issue, one of “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// FileHashes - List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`

	// Message - Error message
	Message string `json:"message"`
}

// PassportElementErrorTranslationFile - Represents an issue with one of the files that constitute the
// translation of a document. The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	// Source - Error source, must be translation_file
	Source string `json:"source"`

	// Type - Type of element of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”,
	// “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// FileHash - Base64-encoded file hash
	FileHash string `json:"file_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// PassportElementErrorTranslationFiles - Represents an issue with the translated version of a document.
// The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	// Source - Error source, must be translation_files
	Source string `json:"source"`

	// Type - Type of element of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”,
	// “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// FileHashes - List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`

	// Message - Error message
	Message string `json:"message"`
}

// PassportElementErrorUnspecified - Represents an issue in an unspecified place. The error is considered
// resolved when new data is added.
type PassportElementErrorUnspecified struct {
	// Source - Error source, must be unspecified
	Source string `json:"source"`

	// Type - Type of element of the user's Telegram Passport which has the issue
	Type string `json:"type"`

	// ElementHash - Base64-encoded element hash
	ElementHash string `json:"element_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// Game - This object represents a game. Use BotFather to create and edit games, their short names
// will act as unique identifiers.
type Game struct {
	// Title - Title of the game
	Title string `json:"title"`

	// Description - Description of the game
	Description string `json:"description"`

	// Photo - Photo that will be displayed in the game message in chats.
	Photo []PhotoSize `json:"photo"`

	// Text - Optional. Brief description of the game or high scores included in the game message.
	// Can be automatically edited to include current high scores for the game when the bot calls
	// setGameScore (#setgamescore), or manually edited using editMessageText (#editmessagetext). 0-4096 characters.
	Text string `json:"text,omitempty"`

	// TextEntities - Optional. Special entities that appear in text, such as usernames, URLFs, bot commands, etc.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`

	// Animation - Optional. Animation that will be displayed in the game message in chats.
	// Upload via BotFather (https://t.me/botfather)
	Animation *Animation `json:"animation,omitempty"`
}

// CallbackGame - A placeholder, currently holds no information.
// Use BotFather (https://t.me/botfather) to set up your game.
type CallbackGame struct {
}

// GameHighScore - This object represents one row of the high scores table for a game.
type GameHighScore struct {
	// Position - Position in high score table for the game
	Position int `json:"position"`

	// User - User
	User User `json:"user"`

	// Score - Score
	Score int `json:"score"`
}
