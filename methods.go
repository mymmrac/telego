package telego

import (
	"fmt"

	"github.com/mymmrac/telego/telegoapi"
)

// GetUpdatesParams - Represents parameters of getUpdates method.
type GetUpdatesParams struct {
	// Offset - Optional. Identifier of the first update to be returned. Must be greater by one than the highest
	// among the identifiers of previously received updates. By default, updates starting with the earliest
	// unconfirmed update are returned. An update is considered confirmed as soon as getUpdates
	// (https://core.telegram.org/bots/api#getupdates) is called with an offset higher than its update_id. The
	// negative offset can be specified to retrieve updates starting from -offset update from the end of the updates
	// queue. All previous updates will forgotten.
	Offset int `json:"offset,omitempty"`

	// Limit - Optional. Limits the number of updates to be retrieved. Values between 1-100 are accepted.
	// Defaults to 100.
	Limit int `json:"limit,omitempty"`

	// Timeout - Optional. Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should
	// be positive, short polling should be used for testing purposes only.
	Timeout int `json:"timeout,omitempty"`

	// AllowedUpdates - Optional. A JSON-serialized list of the update types you want your bot to receive. For
	// example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of
	// these types. See Update (https://core.telegram.org/bots/api#update) for a complete list of available update
	// types. Specify an empty list to receive all update types except chat_member (default). If not specified, the
	// previous setting will be used.Please note that this parameter doesn't affect updates created before the call
	// to the getUpdates, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// Update types you want your bot to receive
const (
	MessageUpdates            = "message"
	EditedMessageUpdates      = "edited_message"
	ChannelPostUpdates        = "channel_post"
	EditedChannelPostUpdates  = "edited_channel_post"
	InlineQueryUpdates        = "inline_query"
	ChosenInlineResultUpdates = "chosen_inline_result"
	CallbackQueryUpdates      = "callback_query"
	ShippingQueryUpdates      = "shipping_query"
	PreCheckoutQueryUpdates   = "pre_checkout_query"
	PollUpdates               = "poll"
	PollAnswerUpdates         = "poll_answer"
	MyChatMemberUpdates       = "my_chat_member"
	ChatMemberUpdates         = "chat_member"
)

// GetUpdates - Use this method to receive incoming updates using long polling (wiki
// (https://en.wikipedia.org/wiki/Push_technology#Long_polling)). An Array of Update
// (https://core.telegram.org/bots/api#update) objects is returned.
func (b *Bot) GetUpdates(params *GetUpdatesParams) ([]Update, error) {
	var updates []Update
	err := b.performRequest("getUpdates", params, &updates)
	if err != nil {
		return nil, fmt.Errorf("getUpdates(): %w", err)
	}

	return updates, nil
}

// SetWebhookParams - Represents parameters of setWebhook method.
type SetWebhookParams struct {
	// URL - HTTPS URL to send updates to. Use an empty string to remove webhook integration
	URL string `json:"url"`

	// Certificate - Optional. Upload your public key certificate so that the root certificate in use can be
	// checked. See our self-signed guide (https://core.telegram.org/bots/self-signed) for details.
	// Please upload as File, sending a FileID or URL will not work.
	Certificate *InputFile `json:"certificate,omitempty"`

	// IPAddress - Optional. The fixed IP address which will be used to send webhook requests instead of the IP
	// address resolved through DNS
	IPAddress string `json:"ip_address,omitempty"`

	// MaxConnections - Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for
	// update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher
	// values to increase your bot's throughput.
	MaxConnections int `json:"max_connections,omitempty"`

	// AllowedUpdates - Optional. A JSON-serialized list of the update types you want your bot to receive. For
	// example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of
	// these types. See Update (https://core.telegram.org/bots/api#update) for a complete list of available update
	// types. Specify an empty list to receive all update types except chat_member (default). If not specified, the
	// previous setting will be used.Please note that this parameter doesn't affect updates created before the call
	// to the setWebhook, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`

	// DropPendingUpdates - Optional. Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

func (s *SetWebhookParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	if s.Certificate != nil {
		fp["certificate"] = s.Certificate.File
	}

	return fp
}

// SetWebhook - Use this method to specify a URL and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing
// a JSON-serialized Update (https://core.telegram.org/bots/api#update). In case of an unsuccessful request, we
// will give up after a reasonable amount of attempts. Returns True on success.
func (b *Bot) SetWebhook(params *SetWebhookParams) error {
	err := b.performRequest("setWebhook", params, nil)
	if err != nil {
		return fmt.Errorf("setWebhook(): %w", err)
	}

	return nil
}

// DeleteWebhookParams - Represents parameters of deleteWebhook method.
type DeleteWebhookParams struct {
	// DropPendingUpdates - Optional. Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// DeleteWebhook - Use this method to remove webhook integration if you decide to switch back to getUpdates
// (https://core.telegram.org/bots/api#getupdates). Returns True on success.
func (b *Bot) DeleteWebhook(params *DeleteWebhookParams) error {
	err := b.performRequest("deleteWebhook", params, nil)
	if err != nil {
		return fmt.Errorf("deleteWebhook(): %w", err)
	}

	return nil
}

// GetWebhookInfo - Use this method to get current webhook status. Requires no parameters. On success,
// returns a WebhookInfo (https://core.telegram.org/bots/api#webhookinfo) object. If the bot is using getUpdates
// (https://core.telegram.org/bots/api#getupdates), will return an object with the URL field empty.
func (b *Bot) GetWebhookInfo() (*WebhookInfo, error) {
	var webhookInfo *WebhookInfo
	err := b.performRequest("getWebhookInfo", nil, &webhookInfo)
	if err != nil {
		return nil, fmt.Errorf("getWebhookInfo(): %w", err)
	}

	return webhookInfo, nil
}

// GetMe - A simple method for testing your bot's authentication token. Requires no parameters. Returns basic
// information about the bot in form of a User (https://core.telegram.org/bots/api#user) object.
func (b *Bot) GetMe() (*User, error) {
	var user *User
	err := b.performRequest("getMe", nil, &user)
	if err != nil {
		return nil, fmt.Errorf("getMe(): %w", err)
	}

	return user, nil
}

// LogOut - Use this method to log out from the cloud Bot API server before launching the bot locally. You
// must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive
// updates. After a successful call, you can immediately log in on a local server, but will not be able to log
// in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
func (b *Bot) LogOut() error {
	err := b.performRequest("logOut", nil, nil)
	if err != nil {
		return fmt.Errorf("logOut(): %w", err)
	}

	return nil
}

// Close - Use this method to close the bot instance before moving it from one local server to another. You
// need to delete the webhook before calling this method to ensure that the bot isn't launched again after
// server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns
// True on success. Requires no parameters.
func (b *Bot) Close() error {
	err := b.performRequest("close", nil, nil)
	if err != nil {
		return fmt.Errorf("close(): %w", err)
	}

	return nil
}

// SendMessageParams - Represents parameters of sendMessage method.
type SendMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Text - Text of the message to be sent, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// ParseMode - Optional. Mode for parsing entities in the message text. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Entities - Optional. A JSON-serialized list of special entities that appear in message text, which can be
	// specified instead of parse_mode
	Entities []MessageEntity `json:"entities,omitempty"`

	// DisableWebPagePreview - Optional. Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Parse modes
const (
	ModeHTML       = "HTML"
	ModeMarkdown   = "Markdown"
	ModeMarkdownV2 = "MarkdownV2"
)

// SendMessage - Use this method to send text messages. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendMessage(params *SendMessageParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendMessage", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendMessage(): %w", err)
	}

	return message, nil
}

// ForwardMessageParams - Represents parameters of forwardMessage method.
type ForwardMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// FromChatID - Unique identifier for the chat where the original message was sent (or channel username in
	// the format @channelusername)
	FromChatID ChatID `json:"from_chat_id"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// MessageID - Message identifier in the chat specified in from_chat_id
	MessageID int `json:"message_id"`
}

// ForwardMessage - Use this method to forward messages of any kind. Service messages can't be forwarded. On
// success, the sent Message (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) ForwardMessage(params *ForwardMessageParams) (*Message, error) {
	var message *Message
	err := b.performRequest("forwardMessage", params, &message)
	if err != nil {
		return nil, fmt.Errorf("forwardMessage(): %w", err)
	}

	return message, nil
}

// CopyMessageParams - Represents parameters of copyMessage method.
type CopyMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// FromChatID - Unique identifier for the chat where the original message was sent (or channel username in
	// the format @channelusername)
	FromChatID ChatID `json:"from_chat_id"`

	// MessageID - Message identifier in the chat specified in from_chat_id
	MessageID int `json:"message_id"`

	// Caption - Optional. New caption for media, 0-1024 characters after entities parsing. If not specified, the
	// original caption is kept
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the new caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the new caption,
	// which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// CopyMessage - Use this method to copy messages of any kind. Service messages and invoice messages can't be
// copied. The method is analogous to the method forwardMessage
// (https://core.telegram.org/bots/api#forwardmessage), but the copied message doesn't have a link to the
// original message. Returns the MessageID (https://core.telegram.org/bots/api#messageid) of the sent message on
// success.
func (b *Bot) CopyMessage(params *CopyMessageParams) (*MessageID, error) {
	var messageID *MessageID
	err := b.performRequest("copyMessage", params, &messageID)
	if err != nil {
		return nil, fmt.Errorf("copyMessage(): %w", err)
	}

	return messageID, nil
}

// SendPhotoParams - Represents parameters of sendPhoto method.
type SendPhotoParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Photo - Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new
	// photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must
	// not exceed 10000 in total. Width and height ratio must be at most 20. More info on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	Photo InputFile `json:"photo"`

	// Caption - Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters
	// after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the photo caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendPhotoParams) fileParameters() map[string]telegoapi.NamedReader {
	return map[string]telegoapi.NamedReader{
		"photo": p.Photo.File,
	}
}

// SendPhoto - Use this method to send photos. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendPhoto(params *SendPhotoParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendPhoto", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendPhoto(): %w", err)
	}

	return message, nil
}

// SendAudioParams - Represents parameters of sendAudio method.
type SendAudioParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Audio - Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or
	// upload a new one using multipart/form-data. More info on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	Audio InputFile `json:"audio"`

	// Caption - Optional. Audio caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the audio caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Duration - Optional. Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`

	// Performer - Optional. Performer
	Performer string `json:"performer,omitempty"`

	// Title - Optional. Track name
	Title string `json:"title,omitempty"`

	// Thumb - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More info on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumb *InputFile `json:"thumb,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendAudioParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	fp["audio"] = p.Audio.File
	if p.Thumb != nil {
		fp["thumb"] = p.Thumb.File
	}

	return fp
}

// SendAudio - Use this method to send audio files, if you want Telegram clients to display them in the music
// player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned. Bots can currently send audio files of up to 50 MB
// in size, this limit may be changed in the future.
// For sending voice messages, use the SendVoice (https://core.telegram.org/bots/api#sendvoice) method instead.
func (b *Bot) SendAudio(params *SendAudioParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendAudio", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendAudio(): %w", err)
	}

	return message, nil
}

// SendDocumentParams - Represents parameters of sendDocument method.
type SendDocumentParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Document - File to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one
	// using multipart/form-data. More info on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Document InputFile `json:"document"`

	// Thumb - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More info on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumb *InputFile `json:"thumb,omitempty"`

	// Caption - Optional. Document caption (may also be used when resending documents by file_id), 0-1024
	// characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the document caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// DisableContentTypeDetection - Optional. Disables automatic server-side content type detection for files
	// uploaded using multipart/form-data
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendDocumentParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	fp["document"] = p.Document.File
	if p.Thumb != nil {
		fp["thumb"] = p.Thumb.File
	}

	return fp
}

// SendDocument - Use this method to send general files. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned. Bots can currently send files of any type of up to
// 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendDocument(params *SendDocumentParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendDocument", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendDocument(): %w", err)
	}

	return message, nil
}

// SendVideoParams - Represents parameters of sendVideo method.
type SendVideoParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Video - Video to send. Pass a file_id as String to send a video that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new
	// video using multipart/form-data. More info on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	Video InputFile `json:"video"`

	// Duration - Optional. Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`

	// Width - Optional. Video width
	Width int `json:"width,omitempty"`

	// Height - Optional. Video height
	Height int `json:"height,omitempty"`

	// Thumb - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More info on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumb *InputFile `json:"thumb,omitempty"`

	// Caption - Optional. Video caption (may also be used when resending videos by file_id), 0-1024 characters
	// after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the video caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// SupportsStreaming - Optional. Pass True, if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendVideoParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	fp["video"] = p.Video.File
	if p.Thumb != nil {
		fp["thumb"] = p.Thumb.File
	}

	return fp
}

// SendVideo - Use this method to send video files, Telegram clients support mp4 videos (other formats may be
// sent as Document (https://core.telegram.org/bots/api#document)). On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned. Bots can currently send video files of up to 50 MB
// in size, this limit may be changed in the future.
func (b *Bot) SendVideo(params *SendVideoParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendVideo", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendVideo(): %w", err)
	}

	return message, nil
}

// SendAnimationParams - Represents parameters of sendAnimation method.
type SendAnimationParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Animation - Animation to send. Pass a file_id as String to send an animation that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or
	// upload a new animation using multipart/form-data. More info on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	Animation InputFile `json:"animation"`

	// Duration - Optional. Duration of sent animation in seconds
	Duration int `json:"duration,omitempty"`

	// Width - Optional. Animation width
	Width int `json:"width,omitempty"`

	// Height - Optional. Animation height
	Height int `json:"height,omitempty"`

	// Thumb - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More info on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumb *InputFile `json:"thumb,omitempty"`

	// Caption - Optional. Animation caption (may also be used when resending animation by file_id), 0-1024
	// characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the animation caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendAnimationParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	fp["animation"] = p.Animation.File
	if p.Thumb != nil {
		fp["thumb"] = p.Thumb.File
	}

	return fp
}

// SendAnimation - Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On
// success, the sent Message (https://core.telegram.org/bots/api#message) is returned. Bots can currently send
// animation files of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendAnimation(params *SendAnimationParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendAnimation", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendAnimation(): %w", err)
	}

	return message, nil
}

// SendVoiceParams - Represents parameters of sendVoice method.
type SendVoiceParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Voice - Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one
	// using multipart/form-data. More info on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Voice InputFile `json:"voice"`

	// Caption - Optional. Voice message caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the voice message caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Duration - Optional. Duration of the voice message in seconds
	Duration int `json:"duration,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendVoiceParams) fileParameters() map[string]telegoapi.NamedReader {
	return map[string]telegoapi.NamedReader{
		"voice": p.Voice.File,
	}
}

// SendVoice - Use this method to send audio files, if you want Telegram clients to display the file as a
// playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats
// may be sent as Audio (https://core.telegram.org/bots/api#audio) or Document
// (https://core.telegram.org/bots/api#document)). On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned. Bots can currently send voice messages of up to 50
// MB in size, this limit may be changed in the future.
func (b *Bot) SendVoice(params *SendVoiceParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendVoice", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendVoice(): %w", err)
	}

	return message, nil
}

// SendVideoNoteParams - Represents parameters of sendVideoNote method.
type SendVideoNoteParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// VideoNote - Video note to send. Pass a file_id as String to send a video note that exists on the Telegram
	// servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files). Sending video notes by a URL is currently unsupported
	VideoNote InputFile `json:"video_note"`

	// Duration - Optional. Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`

	// Length - Optional. Video width and height, i.e. diameter of the video message
	Length int `json:"length,omitempty"`

	// Thumb - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More info on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumb *InputFile `json:"thumb,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendVideoNoteParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	fp["video_note"] = p.VideoNote.File
	if p.Thumb != nil {
		fp["thumb"] = p.Thumb.File
	}

	return fp
}

// SendVideoNote - As of v.4.0 (https://telegram.org/blog/video-messages-and-telescope), Telegram clients
// support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success,
// the sent Message (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendVideoNote(params *SendVideoNoteParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendVideoNote", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendVideoNote(): %w", err)
	}

	return message, nil
}

// SendMediaGroupParams - Represents parameters of sendMediaGroup method.
type SendMediaGroupParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Media - A JSON-serialized array describing messages to be sent, must include 2-10 items
	Media []InputMedia `json:"media"`

	// DisableNotification - Optional. Sends messages silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the messages are a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
}

func (p *SendMediaGroupParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	for _, m := range p.Media {
		for _, v := range m.fileParameters() {
			if isNil(v) {
				continue
			}
			fp[v.Name()] = v
		}
	}

	return fp
}

// SendMediaGroup - Use this method to send a group of photos, videos, documents or audios as an album.
// Documents and audio files can be only grouped in an album with messages of the same type. On success, an
// array of Messages (https://core.telegram.org/bots/api#message) that were sent is returned.
func (b *Bot) SendMediaGroup(params *SendMediaGroupParams) ([]Message, error) {
	var messages []Message
	err := b.performRequest("sendMediaGroup", params, &messages)
	if err != nil {
		return nil, fmt.Errorf("sendMediaGroup(): %w", err)
	}

	return messages, nil
}

// SendLocationParams - Represents parameters of sendLocation method.
type SendLocationParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Latitude - Latitude of the location
	Latitude float64 `json:"latitude"`

	// Longitude - Longitude of the location
	Longitude float64 `json:"longitude"`

	// HorizontalAccuracy - Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// LivePeriod - Optional. Period in seconds for which the location will be updated (see Live Locations
	// (https://telegram.org/blog/live-locations), should be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`

	// Heading - Optional. For live locations, a direction in which the user is moving, in degrees. Must be
	// between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`

	// ProximityAlertRadius - Optional. For live locations, a maximum distance for proximity alerts about
	// approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendLocation - Use this method to send point on the map. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendLocation(params *SendLocationParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendLocation", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendLocation(): %w", err)
	}

	return message, nil
}

// EditMessageLiveLocationParams - Represents parameters of editMessageLiveLocation method.
type EditMessageLiveLocationParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the inline
	// message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Latitude - Latitude of new location
	Latitude float64 `json:"latitude"`

	// Longitude - Longitude of new location
	Longitude float64 `json:"longitude"`

	// HorizontalAccuracy - Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// Heading - Optional. Direction in which the user is moving, in degrees. Must be between 1 and 360 if
	// specified.
	Heading int `json:"heading,omitempty"`

	// ProximityAlertRadius - Optional. Maximum distance for proximity alerts about approaching another chat
	// member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for a new inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageLiveLocation - Use this method to edit live location messages. A location can be edited until
// its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation
// (https://core.telegram.org/bots/api#stopmessagelivelocation). On success, if the edited message is not an
// inline message, the edited Message (https://core.telegram.org/bots/api#message) is returned, otherwise True
// is returned.
func (b *Bot) EditMessageLiveLocation(params *EditMessageLiveLocationParams) (*Message, error) {
	var message *Message
	err := b.performRequest("editMessageLiveLocation", params, &message)
	if err != nil {
		return nil, fmt.Errorf("editMessageLiveLocation(): %w", err)
	}

	return message, nil
}

// StopMessageLiveLocationParams - Represents parameters of stopMessageLiveLocation method.
type StopMessageLiveLocationParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message with live
	// location to stop
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the inline
	// message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for a new inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// StopMessageLiveLocation - Use this method to stop updating a live location message before live_period
// expires. On success, if the message is not an inline message, the edited Message
// (https://core.telegram.org/bots/api#message) is returned, otherwise True is returned.
func (b *Bot) StopMessageLiveLocation(params *StopMessageLiveLocationParams) (*Message, error) {
	var message *Message
	err := b.performRequest("stopMessageLiveLocation", params, &message)
	if err != nil {
		return nil, fmt.Errorf("stopMessageLiveLocation(): %w", err)
	}

	return message, nil
}

// SendVenueParams - Represents parameters of sendVenue method.
type SendVenueParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Latitude - Latitude of the venue
	Latitude float64 `json:"latitude"`

	// Longitude - Longitude of the venue
	Longitude float64 `json:"longitude"`

	// Title - Name of the venue
	Title string `json:"title"`

	// Address - Address of the venue
	Address string `json:"address"`

	// FoursquareID - Optional. Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`

	// FoursquareType - Optional. Foursquare type of the venue, if known. (For example,
	// “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// GooglePlaceID - Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// GooglePlaceType - Optional. Google Places type of the venue. (See supported types
	// (https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType string `json:"google_place_type,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendVenue - Use this method to send information about a venue. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendVenue(params *SendVenueParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendVenue", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendVenue(): %w", err)
	}

	return message, nil
}

// SendContactParams - Represents parameters of sendContact method.
type SendContactParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// PhoneNumber - Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// FirstName - Contact's first name
	FirstName string `json:"first_name"`

	// LastName - Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Vcard - Optional. Additional data about the contact in the form of a vCard
	// (https://en.wikipedia.org/wiki/VCard), 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove keyboard or to force a reply from the
	// user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendContact - Use this method to send phone contacts. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendContact(params *SendContactParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendContact", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendContact(): %w", err)
	}

	return message, nil
}

// SendPollParams - Represents parameters of sendPoll method.
type SendPollParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Question - Poll question, 1-300 characters
	Question string `json:"question"`

	// Options - A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	Options []string `json:"options"`

	// IsAnonymous - Optional. True, if the poll needs to be anonymous, defaults to True
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// Type - Optional. Poll type, “quiz” or “regular”, defaults to “regular”
	Type string `json:"type,omitempty"`

	// AllowsMultipleAnswers - Optional. True, if the poll allows multiple answers, ignored for polls in quiz
	// mode, defaults to False
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`

	// CorrectOptionID - Optional. 0-based identifier of the correct answer option, required for polls in quiz
	// mode
	CorrectOptionID int `json:"correct_option_id,omitempty"`

	// Explanation - Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp
	// icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	Explanation string `json:"explanation,omitempty"`

	// ExplanationParseMode - Optional. Mode for parsing entities in the explanation. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ExplanationParseMode string `json:"explanation_parse_mode,omitempty"`

	// ExplanationEntities - Optional. A JSON-serialized list of special entities that appear in the poll
	// explanation, which can be specified instead of parse_mode
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`

	// OpenPeriod - Optional. Amount of time in seconds the poll will be active after creation, 5-600. Can't be
	// used together with close_date.
	OpenPeriod int `json:"open_period,omitempty"`

	// CloseDate - Optional. Point in time (Unix timestamp) when the poll will be automatically closed. Must be
	// at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
	CloseDate int `json:"close_date,omitempty"`

	// IsClosed - Optional. Pass True, if the poll needs to be immediately closed. This can be useful for poll
	// preview.
	IsClosed bool `json:"is_closed,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendPoll - Use this method to send a native poll. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendPoll(params *SendPollParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendPoll", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendPoll(): %w", err)
	}

	return message, nil
}

// SendDiceParams - Represents parameters of sendDice method.
type SendDiceParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Emoji - Optional. Emoji on which the dice throw animation is based. Currently, must be one of “🎲”,
	// “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”. Dice can have values 1-6 for “🎲”,
	// “🎯” and “🎳”, values 1-5 for “🏀” and “⚽”, and values 1-64 for “🎰”. Defaults
	// to “🎲”
	Emoji string `json:"emoji,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendDice - Use this method to send an animated emoji that will display a random value. On success, the
// sent Message (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendDice(params *SendDiceParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendDice", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendDice(): %w", err)
	}

	return message, nil
}

// SendChatActionParams - Represents parameters of sendChatAction method.
type SendChatActionParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Action - Type of action to broadcast. Choose one, depending on what the user is about to receive:
	// typing for text messages (https://core.telegram.org/bots/api#sendmessage),
	// upload_photo for photos (https://core.telegram.org/bots/api#sendphoto),
	// record_video or upload_video for videos (https://core.telegram.org/bots/api#sendvideo),
	// record_voice or upload_voice for voice notes (https://core.telegram.org/bots/api#sendvoice),
	// upload_document for general files (https://core.telegram.org/bots/api#senddocument),
	// choose_sticker for stickers (https://core.telegram.org/bots/api#sendsticker),
	// find_location for location data (https://core.telegram.org/bots/api#sendlocation),
	// record_video_note or upload_video_note for video notes (https://core.telegram.org/bots/api#sendvideonote).
	Action string `json:"action"`
}

// Chat actions
const (
	ChatActionTyping          = "typing"
	ChatActionUploadPhoto     = "upload_photo"
	ChatActionRecordVideo     = "record_video"
	ChatActionUploadVideo     = "upload_video"
	ChatActionRecordVoice     = "record_voice"
	ChatActionUploadVoice     = "upload_voice"
	ChatActionUploadDocument  = "upload_document"
	ChatActionChooseSticker   = "choose_sticker"
	ChatActionFindLocation    = "find_location"
	ChatActionRecordVideoNote = "record_video_note"
	ChatActionUploadVideoNote = "upload_video_note"
)

// SendChatAction - Use this method when you need to tell the user that something is happening on the bot's
// side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear
// its typing status). Returns True on success.
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
func (b *Bot) SendChatAction(params *SendChatActionParams) error {
	err := b.performRequest("sendChatAction", params, nil)
	if err != nil {
		return fmt.Errorf("sendChatAction(): %w", err)
	}

	return nil
}

// GetUserProfilePhotosParams - Represents parameters of getUserProfilePhotos method.
type GetUserProfilePhotosParams struct {
	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// Offset - Optional. Sequential number of the first photo to be returned. By default, all photos are
	// returned.
	Offset int `json:"offset,omitempty"`

	// Limit - Optional. Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults
	// to 100.
	Limit int `json:"limit,omitempty"`
}

// GetUserProfilePhotos - Use this method to get a list of profile pictures for a user. Returns a
// UserProfilePhotos (https://core.telegram.org/bots/api#userprofilephotos) object.
func (b *Bot) GetUserProfilePhotos(params *GetUserProfilePhotosParams) (*UserProfilePhotos, error) {
	var userProfilePhotos *UserProfilePhotos
	err := b.performRequest("getUserProfilePhotos", params, &userProfilePhotos)
	if err != nil {
		return nil, fmt.Errorf("getUserProfilePhotos(): %w", err)
	}

	return userProfilePhotos, nil
}

// GetFileParams - Represents parameters of getFile method.
type GetFileParams struct {
	// FileID - File identifier to get info about
	FileID string `json:"file_id"`
}

// GetFile - Use this method to get basic info about a file and prepare it for downloading. For the moment,
// bots can download files of up to 20MB in size. On success, a File (https://core.telegram.org/bots/api#file)
// object is returned. The file can then be downloaded via the link
// https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is
// guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested
// by calling getFile (https://core.telegram.org/bots/api#getfile) again.
func (b *Bot) GetFile(params *GetFileParams) (*File, error) {
	var file *File
	err := b.performRequest("getFile", params, &file)
	if err != nil {
		return nil, fmt.Errorf("getFile(): %w", err)
	}

	return file, nil
}

// BanChatMemberParams - Represents parameters of banChatMember method.
type BanChatMemberParams struct {
	// ChatID - Unique identifier for the target group or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// UntilDate - Optional. Date when the user will be unbanned, unix time. If user is banned for more than 366
	// days or less than 30 seconds from the current time they are considered to be banned forever. Applied for
	// supergroups and channels only.
	UntilDate int `json:"until_date,omitempty"`

	// RevokeMessages - Optional. Pass True to delete all messages from the chat for the user that is being
	// removed. If False, the user will be able to see messages in the group that were sent before the user was
	// removed. Always True for supergroups and channels.
	RevokeMessages bool `json:"revoke_messages,omitempty"`
}

// BanChatMember - Use this method to ban a user in a group, a supergroup or a channel. In the case of
// supergroups and channels, the user will not be able to return to the chat on their own using invite links,
// etc., unless unbanned (https://core.telegram.org/bots/api#unbanchatmember) first. The bot must be an
// administrator in the chat for this to work and must have the appropriate administrator rights. Returns True
// on success.
func (b *Bot) BanChatMember(params *BanChatMemberParams) error {
	err := b.performRequest("banChatMember", params, nil)
	if err != nil {
		return fmt.Errorf("banChatMember(): %w", err)
	}

	return nil
}

// UnbanChatMemberParams - Represents parameters of unbanChatMember method.
type UnbanChatMemberParams struct {
	// ChatID - Unique identifier for the target group or username of the target supergroup or channel (in the
	// format @username)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// OnlyIfBanned - Optional. Do nothing if the user is not banned
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

// UnbanChatMember - Use this method to unban a previously banned user in a supergroup or channel. The user
// will not return to the group or channel automatically, but will be able to join via link, etc. The bot must
// be an administrator for this to work. By default, this method guarantees that after the call the user is not
// a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be
// removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
func (b *Bot) UnbanChatMember(params *UnbanChatMemberParams) error {
	err := b.performRequest("unbanChatMember", params, nil)
	if err != nil {
		return fmt.Errorf("unbanChatMember(): %w", err)
	}

	return nil
}

// RestrictChatMemberParams - Represents parameters of restrictChatMember method.
type RestrictChatMemberParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// Permissions - A JSON-serialized object for new user permissions
	Permissions ChatPermissions `json:"permissions"`

	// UntilDate - Optional. Date when restrictions will be lifted for the user, unix time. If user is restricted
	// for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted
	// forever
	UntilDate int `json:"until_date,omitempty"`
}

// RestrictChatMember - Use this method to restrict a user in a supergroup. The bot must be an administrator
// in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all
// permissions to lift restrictions from a user. Returns True on success.
func (b *Bot) RestrictChatMember(params *RestrictChatMemberParams) error {
	err := b.performRequest("restrictChatMember", params, nil)
	if err != nil {
		return fmt.Errorf("restrictChatMember(): %w", err)
	}

	return nil
}

// PromoteChatMemberParams - Represents parameters of promoteChatMember method.
type PromoteChatMemberParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// IsAnonymous - Optional. Pass True, if the administrator's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// CanManageChat - Optional. Pass True, if the administrator can access the chat event log, chat statistics,
	// message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore
	// slow mode. Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat,omitempty"`

	// CanPostMessages - Optional. Pass True, if the administrator can create channel posts, channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// CanEditMessages - Optional. Pass True, if the administrator can edit messages of other users and can pin
	// messages, channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// CanDeleteMessages - Optional. Pass True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

	// CanManageVoiceChats - Optional. Pass True, if the administrator can manage voice chats
	CanManageVoiceChats bool `json:"can_manage_voice_chats,omitempty"`

	// CanRestrictMembers - Optional. Pass True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

	// CanPromoteMembers - Optional. Pass True, if the administrator can add new administrators with a subset of
	// their own privileges or demote administrators that he has promoted, directly or indirectly (promoted by
	// administrators that were appointed by him)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`

	// CanChangeInfo - Optional. Pass True, if the administrator can change chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// CanInviteUsers - Optional. Pass True, if the administrator can invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// CanPinMessages - Optional. Pass True, if the administrator can pin messages, supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// PromoteChatMember - Use this method to promote or demote a user in a supergroup or a channel. The bot must
// be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass
// False for all boolean parameters to demote a user. Returns True on success.
func (b *Bot) PromoteChatMember(params *PromoteChatMemberParams) error {
	err := b.performRequest("promoteChatMember", params, nil)
	if err != nil {
		return fmt.Errorf("promoteChatMember(): %w", err)
	}

	return nil
}

// SetChatAdministratorCustomTitleParams - Represents parameters of setChatAdministratorCustomTitle method.
type SetChatAdministratorCustomTitleParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// CustomTitle - New custom title for the administrator; 0-16 characters, emoji are not allowed
	CustomTitle string `json:"custom_title"`
}

// SetChatAdministratorCustomTitle - Use this method to set a custom title for an administrator in a
// supergroup promoted by the bot. Returns True on success.
func (b *Bot) SetChatAdministratorCustomTitle(params *SetChatAdministratorCustomTitleParams) error {
	err := b.performRequest("setChatAdministratorCustomTitle", params, nil)
	if err != nil {
		return fmt.Errorf("setChatAdministratorCustomTitle(): %w", err)
	}

	return nil
}

// SetChatPermissionsParams - Represents parameters of setChatPermissions method.
type SetChatPermissionsParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID ChatID `json:"chat_id"`

	// Permissions - A JSON-serialized object for new default chat permissions
	Permissions ChatPermissions `json:"permissions"`
}

// SetChatPermissions - Use this method to set default chat permissions for all members. The bot must be an
// administrator in the group or a supergroup for this to work and must have the can_restrict_members
// administrator rights. Returns True on success.
func (b *Bot) SetChatPermissions(params *SetChatPermissionsParams) error {
	err := b.performRequest("setChatPermissions", params, nil)
	if err != nil {
		return fmt.Errorf("setChatPermissions(): %w", err)
	}

	return nil
}

// ExportChatInviteLinkParams - Represents parameters of exportChatInviteLink method.
type ExportChatInviteLinkParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`
}

// ExportChatInviteLink - Use this method to generate a new primary invite link for a chat; any previously
// generated primary link is revoked. The bot must be an administrator in the chat for this to work and must
// have the appropriate administrator rights. Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other
// administrators. If you want your bot to work with invite links, it will need to generate its own link using
// exportChatInviteLink or by calling the getChat method. If your bot needs to generate a new primary invite link
// replacing its previous one, use exportChatInviteLink again.
func (b *Bot) ExportChatInviteLink(params *ExportChatInviteLinkParams) (*string, error) {
	var inviteLink *string
	err := b.performRequest("exportChatInviteLink", params, &inviteLink)
	if err != nil {
		return nil, fmt.Errorf("exportChatInviteLink(): %w", err)
	}

	return inviteLink, nil
}

// CreateChatInviteLinkParams - Represents parameters of createChatInviteLink method.
type CreateChatInviteLinkParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Name - Optional. Invite link name; 0-32 characters
	Name string `json:"name,omitempty"`

	// ExpireDate - Optional. Point in time (Unix timestamp) when the link will expire
	ExpireDate int `json:"expire_date,omitempty"`

	// MemberLimit - Optional. Maximum number of users that can be members of the chat simultaneously after
	// joining the chat via this invite link; 1-99999
	MemberLimit int `json:"member_limit,omitempty"`

	// CreatesJoinRequest - Optional. True, if users joining the chat via the link need to be approved by chat
	// administrators. If True, member_limit can't be specified
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// CreateChatInviteLink - Use this method to create an additional invite link for a chat. The bot must be an
// administrator in the chat for this to work and must have the appropriate administrator rights. The link can
// be revoked using the method revokeChatInviteLink (https://core.telegram.org/bots/api#revokechatinvitelink).
// Returns the new invite link as ChatInviteLink (https://core.telegram.org/bots/api#chatinvitelink) object.
func (b *Bot) CreateChatInviteLink(params *CreateChatInviteLinkParams) (*ChatInviteLink, error) {
	var chatInviteLink *ChatInviteLink
	err := b.performRequest("createChatInviteLink", params, &chatInviteLink)
	if err != nil {
		return nil, fmt.Errorf("createChatInviteLink(): %w", err)
	}

	return chatInviteLink, nil
}

// EditChatInviteLinkParams - Represents parameters of editChatInviteLink method.
type EditChatInviteLinkParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// InviteLink - The invite link to edit
	InviteLink string `json:"invite_link"`

	// Name - Optional. Invite link name; 0-32 characters
	Name string `json:"name,omitempty"`

	// ExpireDate - Optional. Point in time (Unix timestamp) when the link will expire
	ExpireDate int `json:"expire_date,omitempty"`

	// MemberLimit - Optional. Maximum number of users that can be members of the chat simultaneously after
	// joining the chat via this invite link; 1-99999
	MemberLimit int `json:"member_limit,omitempty"`

	// CreatesJoinRequest - Optional. True, if users joining the chat via the link need to be approved by chat
	// administrators. If True, member_limit can't be specified
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// EditChatInviteLink - Use this method to edit a non-primary invite link created by the bot. The bot must be
// an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the
// edited invite link as a ChatInviteLink (https://core.telegram.org/bots/api#chatinvitelink) object.
func (b *Bot) EditChatInviteLink(params *EditChatInviteLinkParams) (*ChatInviteLink, error) {
	var chatInviteLink *ChatInviteLink
	err := b.performRequest("editChatInviteLink", params, &chatInviteLink)
	if err != nil {
		return nil, fmt.Errorf("editChatInviteLink(): %w", err)
	}

	return chatInviteLink, nil
}

// RevokeChatInviteLinkParams - Represents parameters of revokeChatInviteLink method.
type RevokeChatInviteLinkParams struct {
	// ChatID - Unique identifier of the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// InviteLink - The invite link to revoke
	InviteLink string `json:"invite_link"`
}

// RevokeChatInviteLink - Use this method to revoke an invite link created by the bot. If the primary link is
// revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work
// and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink
// (https://core.telegram.org/bots/api#chatinvitelink) object.
func (b *Bot) RevokeChatInviteLink(params *RevokeChatInviteLinkParams) (*ChatInviteLink, error) {
	var chatInviteLink *ChatInviteLink
	err := b.performRequest("revokeChatInviteLink", params, &chatInviteLink)
	if err != nil {
		return nil, fmt.Errorf("revokeChatInviteLink(): %w", err)
	}

	return chatInviteLink, nil
}

// ApproveChatJoinRequestParams - Represents parameters of approveChatJoinRequest method.
type ApproveChatJoinRequestParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// ApproveChatJoinRequest - Use this method to approve a chat join request. The bot must be an administrator
// in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (b *Bot) ApproveChatJoinRequest(params *ApproveChatJoinRequestParams) error {
	err := b.performRequest("approveChatJoinRequest", params, nil)
	if err != nil {
		return fmt.Errorf("approveChatJoinRequest(): %w", err)
	}

	return nil
}

// DeclineChatJoinRequestParams - Represents parameters of declineChatJoinRequest method.
type DeclineChatJoinRequestParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// DeclineChatJoinRequest - Use this method to decline a chat join request. The bot must be an administrator
// in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (b *Bot) DeclineChatJoinRequest(params *DeclineChatJoinRequestParams) error {
	err := b.performRequest("declineChatJoinRequest", params, nil)
	if err != nil {
		return fmt.Errorf("declineChatJoinRequest(): %w", err)
	}

	return nil
}

// SetChatPhotoParams - Represents parameters of setChatPhoto method.
type SetChatPhotoParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Photo - New chat photo, uploaded using multipart/form-data
	Photo InputFile `json:"photo"`
}

func (p *SetChatPhotoParams) fileParameters() map[string]telegoapi.NamedReader {
	return map[string]telegoapi.NamedReader{
		"photo": p.Photo.File,
	}
}

// SetChatPhoto - Use this method to set a new profile photo for the chat. Photos can't be changed for
// private chats. The bot must be an administrator in the chat for this to work and must have the appropriate
// administrator rights. Returns True on success.
func (b *Bot) SetChatPhoto(params *SetChatPhotoParams) error {
	err := b.performRequest("setChatPhoto", params, nil)
	if err != nil {
		return fmt.Errorf("setChatPhoto(): %w", err)
	}

	return nil
}

// DeleteChatPhotoParams - Represents parameters of deleteChatPhoto method.
type DeleteChatPhotoParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`
}

// DeleteChatPhoto - Use this method to delete a chat photo. Photos can't be changed for private chats. The
// bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
func (b *Bot) DeleteChatPhoto(params *DeleteChatPhotoParams) error {
	err := b.performRequest("deleteChatPhoto", params, nil)
	if err != nil {
		return fmt.Errorf("deleteChatPhoto(): %w", err)
	}

	return nil
}

// SetChatTitleParams - Represents parameters of setChatTitle method.
type SetChatTitleParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Title - New chat title, 1-255 characters
	Title string `json:"title"`
}

// SetChatTitle - Use this method to change the title of a chat. Titles can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator
// rights. Returns True on success.
func (b *Bot) SetChatTitle(params *SetChatTitleParams) error {
	err := b.performRequest("setChatTitle", params, nil)
	if err != nil {
		return fmt.Errorf("setChatTitle(): %w", err)
	}

	return nil
}

// SetChatDescriptionParams - Represents parameters of setChatDescription method.
type SetChatDescriptionParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Description - Optional. New chat description, 0-255 characters
	Description string `json:"description,omitempty"`
}

// SetChatDescription - Use this method to change the description of a group, a supergroup or a channel. The
// bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
func (b *Bot) SetChatDescription(params *SetChatDescriptionParams) error {
	err := b.performRequest("setChatDescription", params, nil)
	if err != nil {
		return fmt.Errorf("setChatDescription(): %w", err)
	}

	return nil
}

// PinChatMessageParams - Represents parameters of pinChatMessage method.
type PinChatMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// MessageID - Identifier of a message to pin
	MessageID int `json:"message_id"`

	// DisableNotification - Optional. Pass True, if it is not necessary to send a notification to all chat
	// members about the new pinned message. Notifications are always disabled in channels and private chats.
	DisableNotification bool `json:"disable_notification,omitempty"`
}

// PinChatMessage - Use this method to add a message to the list of pinned messages in a chat. If the chat is
// not a private chat, the bot must be an administrator in the chat for this to work and must have the
// 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a
// channel. Returns True on success.
func (b *Bot) PinChatMessage(params *PinChatMessageParams) error {
	err := b.performRequest("pinChatMessage", params, nil)
	if err != nil {
		return fmt.Errorf("pinChatMessage(): %w", err)
	}

	return nil
}

// UnpinChatMessageParams - Represents parameters of unpinChatMessage method.
type UnpinChatMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// MessageID - Optional. Identifier of a message to unpin. If not specified, the most recent pinned message
	// (by sending date) will be unpinned.
	MessageID int `json:"message_id,omitempty"`
}

// UnpinChatMessage - Use this method to remove a message from the list of pinned messages in a chat. If the
// chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the
// 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a
// channel. Returns True on success.
func (b *Bot) UnpinChatMessage(params *UnpinChatMessageParams) error {
	err := b.performRequest("unpinChatMessage", params, nil)
	if err != nil {
		return fmt.Errorf("unpinChatMessage(): %w", err)
	}

	return nil
}

// UnpinAllChatMessagesParams - Represents parameters of unpinAllChatMessages method.
type UnpinAllChatMessagesParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`
}

// UnpinAllChatMessages - Use this method to clear the list of pinned messages in a chat. If the chat is not
// a private chat, the bot must be an administrator in the chat for this to work and must have the
// 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a
// channel. Returns True on success.
func (b *Bot) UnpinAllChatMessages(params *UnpinAllChatMessagesParams) error {
	err := b.performRequest("unpinAllChatMessages", params, nil)
	if err != nil {
		return fmt.Errorf("unpinAllChatMessages(): %w", err)
	}

	return nil
}

// LeaveChatParams - Represents parameters of leaveChat method.
type LeaveChatParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID ChatID `json:"chat_id"`
}

// LeaveChat - Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
func (b *Bot) LeaveChat(params *LeaveChatParams) error {
	err := b.performRequest("leaveChat", params, nil)
	if err != nil {
		return fmt.Errorf("leaveChat(): %w", err)
	}

	return nil
}

// GetChatParams - Represents parameters of getChat method.
type GetChatParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID ChatID `json:"chat_id"`
}

// GetChat - Use this method to get up to date information about the chat (current name of the user for
// one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat
// (https://core.telegram.org/bots/api#chat) object on success.
func (b *Bot) GetChat(params *GetChatParams) (*Chat, error) {
	var chat *Chat
	err := b.performRequest("getChat", params, &chat)
	if err != nil {
		return nil, fmt.Errorf("getChat(): %w", err)
	}

	return chat, nil
}

// GetChatAdministratorsParams - Represents parameters of getChatAdministrators method.
type GetChatAdministratorsParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID ChatID `json:"chat_id"`
}

// GetChatAdministrators - Use this method to get a list of administrators in a chat. On success, returns an
// Array of ChatMember (https://core.telegram.org/bots/api#chatmember) objects that contains information about
// all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were
// appointed, only the creator will be returned.
func (b *Bot) GetChatAdministrators(params *GetChatAdministratorsParams) ([]ChatMember, error) {
	var chatMembersData []chatMemberData
	err := b.performRequest("getChatAdministrators", params, &chatMembersData)
	if err != nil {
		return nil, fmt.Errorf("getChatAdministrators(): %w", err)
	}

	chatMembers := make([]ChatMember, len(chatMembersData))
	for i, d := range chatMembersData {
		chatMembers[i] = d.Data
	}
	return chatMembers, nil
}

// GetChatMemberCountParams - Represents parameters of getChatMemberCount method.
type GetChatMemberCountParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID ChatID `json:"chat_id"`
}

// GetChatMemberCount - Use this method to get the number of members in a chat. Returns Int on success.
func (b *Bot) GetChatMemberCount(params *GetChatMemberCountParams) (*int, error) {
	var chatMemberCount *int
	err := b.performRequest("getChatMemberCount", params, &chatMemberCount)
	if err != nil {
		return nil, fmt.Errorf("getChatMemberCount(): %w", err)
	}

	return chatMemberCount, nil
}

// GetChatMemberParams - Represents parameters of getChatMember method.
type GetChatMemberParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// GetChatMember - Use this method to get information about a member of a chat. Returns a ChatMember
// (https://core.telegram.org/bots/api#chatmember) object on success.
func (b *Bot) GetChatMember(params *GetChatMemberParams) (ChatMember, error) {
	var memberData chatMemberData
	err := b.performRequest("getChatMember", params, &memberData)
	if err != nil {
		return nil, fmt.Errorf("getChatMember(): %w", err)
	}

	return memberData.Data, nil
}

// SetChatStickerSetParams - Represents parameters of setChatStickerSet method.
type SetChatStickerSetParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID ChatID `json:"chat_id"`

	// StickerSetName - Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name"`
}

// SetChatStickerSet - Use this method to set a new group sticker set for a supergroup. The bot must be an
// administrator in the chat for this to work and must have the appropriate administrator rights. Use the field
// can_set_sticker_set optionally returned in getChat (https://core.telegram.org/bots/api#getchat) requests to
// check if the bot can use this method. Returns True on success.
func (b *Bot) SetChatStickerSet(params *SetChatStickerSetParams) error {
	err := b.performRequest("setChatStickerSet", params, nil)
	if err != nil {
		return fmt.Errorf("setChatStickerSet(): %w", err)
	}

	return nil
}

// DeleteChatStickerSetParams - Represents parameters of deleteChatStickerSet method.
type DeleteChatStickerSetParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID ChatID `json:"chat_id"`
}

// DeleteChatStickerSet - Use this method to delete a group sticker set from a supergroup. The bot must be an
// administrator in the chat for this to work and must have the appropriate administrator rights. Use the field
// can_set_sticker_set optionally returned in getChat (https://core.telegram.org/bots/api#getchat) requests to
// check if the bot can use this method. Returns True on success.
func (b *Bot) DeleteChatStickerSet(params *DeleteChatStickerSetParams) error {
	err := b.performRequest("deleteChatStickerSet", params, nil)
	if err != nil {
		return fmt.Errorf("deleteChatStickerSet(): %w", err)
	}

	return nil
}

// AnswerCallbackQueryParams - Represents parameters of answerCallbackQuery method.
type AnswerCallbackQueryParams struct {
	// CallbackQueryID - Unique identifier for the query to be answered
	CallbackQueryID string `json:"callback_query_id"`

	// Text - Optional. Text of the notification. If not specified, nothing will be shown to the user, 0-200
	// characters
	Text string `json:"text,omitempty"`

	// ShowAlert - Optional. If True, an alert will be shown by the client instead of a notification at the top
	// of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`

	// URL - Optional. URL that will be opened by the user's client. If you have created a Game
	// (https://core.telegram.org/bots/api#game) and accepted the conditions via @BotFather
	// (https://t.me/botfather), specify the URL that opens your game — note that this will only work if the query
	// comes from a callback_game (https://core.telegram.org/bots/api#inlinekeyboardbutton) button.Otherwise, you
	// may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	URL string `json:"url,omitempty"`

	// CacheTime - Optional. The maximum amount of time in seconds that the result of the callback query may be
	// cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime int `json:"cache_time,omitempty"`
}

// AnswerCallbackQuery - Use this method to send answers to callback queries sent from inline keyboards
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating). The answer will be displayed to
// the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
func (b *Bot) AnswerCallbackQuery(params *AnswerCallbackQueryParams) error {
	err := b.performRequest("answerCallbackQuery", params, nil)
	if err != nil {
		return fmt.Errorf("answerCallbackQuery(): %w", err)
	}

	return nil
}

// SetMyCommandsParams - Represents parameters of setMyCommands method.
type SetMyCommandsParams struct {
	// Commands - A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100
	// commands can be specified.
	Commands []BotCommand `json:"commands"`

	// Scope - Optional. A JSON-serialized object, describing scope of users for which the commands are relevant.
	// Defaults to BotCommandScopeDefault (https://core.telegram.org/bots/api#botcommandscopedefault).
	Scope BotCommandScope `json:"scope,omitempty"`

	// LanguageCode - Optional. A two-letter ISO 639-1 language code. If empty, commands will be applied to all
	// users from the given scope, for whose language there are no dedicated commands
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyCommands - Use this method to change the list of the bot's commands. See
// https://core.telegram.org/bots#commands (https://core.telegram.org/bots#commands) for more details about bot
// commands. Returns True on success.
func (b *Bot) SetMyCommands(params *SetMyCommandsParams) error {
	err := b.performRequest("setMyCommands", params, nil)
	if err != nil {
		return fmt.Errorf("setMyCommands(): %w", err)
	}

	return nil
}

// DeleteMyCommandsParams - Represents parameters of deleteMyCommands method.
type DeleteMyCommandsParams struct {
	// Scope - Optional. A JSON-serialized object, describing scope of users for which the commands are relevant.
	// Defaults to BotCommandScopeDefault (https://core.telegram.org/bots/api#botcommandscopedefault).
	Scope BotCommandScope `json:"scope,omitempty"`

	// LanguageCode - Optional. A two-letter ISO 639-1 language code. If empty, commands will be applied to all
	// users from the given scope, for whose language there are no dedicated commands
	LanguageCode string `json:"language_code,omitempty"`
}

// DeleteMyCommands - Use this method to delete the list of the bot's commands for the given scope and user
// language. After deletion, higher level commands
// (https://core.telegram.org/bots/api#determining-list-of-commands) will be shown to affected users. Returns
// True on success.
func (b *Bot) DeleteMyCommands(params *DeleteMyCommandsParams) error {
	err := b.performRequest("deleteMyCommands", params, nil)
	if err != nil {
		return fmt.Errorf("deleteMyCommands(): %w", err)
	}

	return nil
}

// GetMyCommandsParams - Represents parameters of getMyCommands method.
type GetMyCommandsParams struct {
	// Scope - Optional. A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault
	// (https://core.telegram.org/bots/api#botcommandscopedefault).
	Scope BotCommandScope `json:"scope,omitempty"`

	// LanguageCode - Optional. A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyCommands - Use this method to get the current list of the bot's commands for the given scope and user
// language. Returns Array of BotCommand (https://core.telegram.org/bots/api#botcommand) on success. If commands
// aren't set, an empty list is returned.
func (b *Bot) GetMyCommands(params *GetMyCommandsParams) ([]BotCommand, error) {
	var botCommands []BotCommand
	err := b.performRequest("getMyCommands", params, &botCommands)
	if err != nil {
		return nil, fmt.Errorf("getMyCommands(): %w", err)
	}

	return botCommands, nil
}

// EditMessageTextParams - Represents parameters of editMessageText method.
type EditMessageTextParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the inline
	// message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Text - New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// ParseMode - Optional. Mode for parsing entities in the message text. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Entities - Optional. A JSON-serialized list of special entities that appear in message text, which can be
	// specified instead of parse_mode
	Entities []MessageEntity `json:"entities,omitempty"`

	// DisableWebPagePreview - Optional. Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageText - Use this method to edit text and game (https://core.telegram.org/bots/api#games)
// messages. On success, if the edited message is not an inline message, the edited Message
// (https://core.telegram.org/bots/api#message) is returned, otherwise True is returned.
func (b *Bot) EditMessageText(params *EditMessageTextParams) (*Message, error) {
	var message *Message
	err := b.performRequest("editMessageText", params, &message)
	if err != nil {
		return nil, fmt.Errorf("editMessageText(): %w", err)
	}

	return message, nil
}

// EditMessageCaptionParams - Represents parameters of editMessageCaption method.
type EditMessageCaptionParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the inline
	// message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Caption - Optional. New caption of the message, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the message caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageCaption - Use this method to edit captions of messages. On success, if the edited message is
// not an inline message, the edited Message (https://core.telegram.org/bots/api#message) is returned, otherwise
// True is returned.
func (b *Bot) EditMessageCaption(params *EditMessageCaptionParams) (*Message, error) {
	var message *Message
	err := b.performRequest("editMessageCaption", params, &message)
	if err != nil {
		return nil, fmt.Errorf("editMessageCaption(): %w", err)
	}

	return message, nil
}

// EditMessageMediaParams - Represents parameters of editMessageMedia method.
type EditMessageMediaParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the inline
	// message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Media - A JSON-serialized object for a new media content of the message
	Media InputMedia `json:"media"`

	// ReplyMarkup - Optional. A JSON-serialized object for a new inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *EditMessageMediaParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	for _, v := range p.Media.fileParameters() {
		if isNil(v) {
			continue
		}
		fp[v.Name()] = v
	}

	return fp
}

// EditMessageMedia - Use this method to edit animation, audio, document, photo, or video messages. If a
// message is part of a message album, then it can be edited only to an audio for audio albums, only to a
// document for document albums and to a photo or a video otherwise. When an inline message is edited, a new
// file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the
// edited message is not an inline message, the edited Message (https://core.telegram.org/bots/api#message) is
// returned, otherwise True is returned.
func (b *Bot) EditMessageMedia(params *EditMessageMediaParams) (*Message, error) {
	var message *Message
	err := b.performRequest("editMessageMedia", params, &message)
	if err != nil {
		return nil, fmt.Errorf("editMessageMedia(): %w", err)
	}

	return message, nil
}

// EditMessageReplyMarkupParams - Represents parameters of editMessageReplyMarkup method.
type EditMessageReplyMarkupParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the inline
	// message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkup - Use this method to edit only the reply markup of messages. On success, if the
// edited message is not an inline message, the edited Message (https://core.telegram.org/bots/api#message) is
// returned, otherwise True is returned.
func (b *Bot) EditMessageReplyMarkup(params *EditMessageReplyMarkupParams) (*Message, error) {
	var message *Message
	err := b.performRequest("editMessageReplyMarkup", params, &message)
	if err != nil {
		return nil, fmt.Errorf("editMessageReplyMarkup(): %w", err)
	}

	return message, nil
}

// StopPollParams - Represents parameters of stopPoll method.
type StopPollParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// MessageID - Identifier of the original message with the poll
	MessageID int `json:"message_id"`

	// ReplyMarkup - Optional. A JSON-serialized object for a new message inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// StopPoll - Use this method to stop a poll which was sent by the bot. On success, the stopped Poll
// (https://core.telegram.org/bots/api#poll) is returned.
func (b *Bot) StopPoll(params *StopPollParams) (*Poll, error) {
	var poll *Poll
	err := b.performRequest("stopPoll", params, &poll)
	if err != nil {
		return nil, fmt.Errorf("stopPoll(): %w", err)
	}

	return poll, nil
}

// DeleteMessageParams - Represents parameters of deleteMessage method.
type DeleteMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// MessageID - Identifier of the message to delete
	MessageID int `json:"message_id"`
}

// DeleteMessage - Use this method to delete a message, including service messages, with the following
// limitations:- A message can only be deleted if it was sent less than 48 hours ago.- A dice message in a
// private chat can only be deleted if it was sent more than 24 hours ago.- Bots can delete outgoing messages in
// private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted
// can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a
// group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a
// channel, it can delete any message there.Returns True on success.
func (b *Bot) DeleteMessage(params *DeleteMessageParams) error {
	err := b.performRequest("deleteMessage", params, nil)
	if err != nil {
		return fmt.Errorf("deleteMessage(): %w", err)
	}

	return nil
}

// SendStickerParams - Represents parameters of sendSticker method.
type SendStickerParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Sticker - Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP file from the Internet, or upload a
	// new one using multipart/form-data. More info on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	Sticker InputFile `json:"sticker"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard
	// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from
	// the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendStickerParams) fileParameters() map[string]telegoapi.NamedReader {
	return map[string]telegoapi.NamedReader{
		"sticker": p.Sticker.File,
	}
}

// SendSticker - Use this method to send static .WEBP or animated
// (https://telegram.org/blog/animated-stickers) .TGS stickers. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendSticker(params *SendStickerParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendSticker", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendSticker(): %w", err)
	}

	return message, nil
}

// GetStickerSetParams - Represents parameters of getStickerSet method.
type GetStickerSetParams struct {
	// Name - Name of the sticker set
	Name string `json:"name"`
}

// GetStickerSet - Use this method to get a sticker set. On success, a StickerSet
// (https://core.telegram.org/bots/api#stickerset) object is returned.
func (b *Bot) GetStickerSet(params *GetStickerSetParams) (*StickerSet, error) {
	var stickerSet *StickerSet
	err := b.performRequest("getStickerSet", params, &stickerSet)
	if err != nil {
		return nil, fmt.Errorf("getStickerSet(): %w", err)
	}

	return stickerSet, nil
}

// UploadStickerFileParams - Represents parameters of uploadStickerFile method.
type UploadStickerFileParams struct {
	// UserID - User identifier of sticker file owner
	UserID int64 `json:"user_id"`

	// PngSticker - PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed
	// 512px, and either width or height must be exactly 512px. More info on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	PngSticker InputFile `json:"png_sticker"`
}

func (p *UploadStickerFileParams) fileParameters() map[string]telegoapi.NamedReader {
	return map[string]telegoapi.NamedReader{
		"png_sticker": p.PngSticker.File,
	}
}

// UploadStickerFile - Use this method to upload a .PNG file with a sticker for later use in
// createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File
// (https://core.telegram.org/bots/api#file) on success.
func (b *Bot) UploadStickerFile(params *UploadStickerFileParams) (*File, error) {
	var file *File
	err := b.performRequest("uploadStickerFile", params, &file)
	if err != nil {
		return nil, fmt.Errorf("uploadStickerFile(): %w", err)
	}

	return file, nil
}

// CreateNewStickerSetParams - Represents parameters of createNewStickerSet method.
type CreateNewStickerSetParams struct {
	// UserID - User identifier of created sticker set owner
	UserID int64 `json:"user_id"`

	// Name - Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only
	// english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and
	// must end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
	Name string `json:"name"`

	// Title - Sticker set title, 1-64 characters
	Title string `json:"title"`

	// PngSticker - Optional. PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must
	// not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file
	// that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the
	// Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	PngSticker *InputFile `json:"png_sticker,omitempty"`

	// TgsSticker - Optional. TGS animation with the sticker, uploaded using multipart/form-data. See
	// https://core.telegram.org/animated_stickers#technical-requirements
	// (https://core.telegram.org/animated_stickers#technical-requirements) for technical requirements
	TgsSticker *InputFile `json:"tgs_sticker,omitempty"`

	// Emojis - One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`

	// ContainsMasks - Optional. Pass True, if a set of mask stickers should be created
	ContainsMasks bool `json:"contains_masks,omitempty"`

	// MaskPosition - Optional. A JSON-serialized object for position where the mask should be placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

func (p *CreateNewStickerSetParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	if p.PngSticker != nil {
		fp["png_sticker"] = p.PngSticker.File
	}
	if p.TgsSticker != nil {
		fp["tgs_sticker"] = p.TgsSticker.File
	}

	return fp
}

// CreateNewStickerSet - Use this method to create a new sticker set owned by a user. The bot will be able to
// edit the sticker set thus created. You must use exactly one of the fields png_sticker or tgs_sticker. Returns
// True on success.
func (b *Bot) CreateNewStickerSet(params *CreateNewStickerSetParams) error {
	err := b.performRequest("createNewStickerSet", params, nil)
	if err != nil {
		return fmt.Errorf("createNewStickerSet(): %w", err)
	}

	return nil
}

// AddStickerToSetParams - Represents parameters of addStickerToSet method.
type AddStickerToSetParams struct {
	// UserID - User identifier of sticker set owner
	UserID int64 `json:"user_id"`

	// Name - Sticker set name
	Name string `json:"name"`

	// PngSticker - Optional. PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must
	// not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file
	// that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the
	// Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	PngSticker *InputFile `json:"png_sticker,omitempty"`

	// TgsSticker - Optional. TGS animation with the sticker, uploaded using multipart/form-data. See
	// https://core.telegram.org/animated_stickers#technical-requirements
	// (https://core.telegram.org/animated_stickers#technical-requirements) for technical requirements
	TgsSticker *InputFile `json:"tgs_sticker,omitempty"`

	// Emojis - One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`

	// MaskPosition - Optional. A JSON-serialized object for position where the mask should be placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

func (p *AddStickerToSetParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	if p.PngSticker != nil {
		fp["png_sticker"] = p.PngSticker.File
	}
	if p.TgsSticker != nil {
		fp["tgs_sticker"] = p.TgsSticker.File
	}

	return fp
}

// AddStickerToSet - Use this method to add a new sticker to a set created by the bot. You must use exactly
// one of the fields png_sticker or tgs_sticker. Animated stickers can be added to animated sticker sets and
// only to them. Animated sticker sets can have up to 50 stickers. Static sticker sets can have up to 120
// stickers. Returns True on success.
func (b *Bot) AddStickerToSet(params *AddStickerToSetParams) error {
	err := b.performRequest("addStickerToSet", params, nil)
	if err != nil {
		return fmt.Errorf("addStickerToSet(): %w", err)
	}

	return nil
}

// SetStickerPositionInSetParams - Represents parameters of setStickerPositionInSet method.
type SetStickerPositionInSetParams struct {
	// Sticker - File identifier of the sticker
	Sticker string `json:"sticker"`

	// Position - New sticker position in the set, zero-based
	Position int `json:"position"`
}

// SetStickerPositionInSet - Use this method to move a sticker in a set created by the bot to a specific
// position. Returns True on success.
func (b *Bot) SetStickerPositionInSet(params *SetStickerPositionInSetParams) error {
	err := b.performRequest("setStickerPositionInSet", params, nil)
	if err != nil {
		return fmt.Errorf("setStickerPositionInSet(): %w", err)
	}

	return nil
}

// DeleteStickerFromSetParams - Represents parameters of deleteStickerFromSet method.
type DeleteStickerFromSetParams struct {
	// Sticker - File identifier of the sticker
	Sticker string `json:"sticker"`
}

// DeleteStickerFromSet - Use this method to delete a sticker from a set created by the bot. Returns True on
// success.
func (b *Bot) DeleteStickerFromSet(params *DeleteStickerFromSetParams) error {
	err := b.performRequest("deleteStickerFromSet", params, nil)
	if err != nil {
		return fmt.Errorf("deleteStickerFromSet(): %w", err)
	}

	return nil
}

// SetStickerSetThumbParams - Represents parameters of setStickerSetThumb method.
type SetStickerSetThumbParams struct {
	// Name - Sticker set name
	Name string `json:"name"`

	// UserID - User identifier of the sticker set owner
	UserID int64 `json:"user_id"`

	// Thumb - Optional. A PNG image with the thumbnail, must be up to 128 kilobytes in size and have width and
	// height exactly 100px, or a TGS animation with the thumbnail up to 32 kilobytes in size; see
	// https://core.telegram.org/animated_stickers#technical-requirements
	// (https://core.telegram.org/animated_stickers#technical-requirements) for animated sticker technical
	// requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an
	// HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using
	// multipart/form-data. More info on Sending Files » (https://core.telegram.org/bots/api#sending-files).
	// Animated sticker set thumbnail can't be uploaded via HTTP URL.
	Thumb *InputFile `json:"thumb,omitempty"`
}

func (p *SetStickerSetThumbParams) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	if p.Thumb != nil {
		fp["thumb"] = p.Thumb.File
	}

	return fp
}

// SetStickerSetThumb - Use this method to set the thumbnail of a sticker set. Animated thumbnails can be set
// for animated sticker sets only. Returns True on success.
func (b *Bot) SetStickerSetThumb(params *SetStickerSetThumbParams) error {
	err := b.performRequest("setStickerSetThumb", params, nil)
	if err != nil {
		return fmt.Errorf("setStickerSetThumb(): %w", err)
	}

	return nil
}

// AnswerInlineQueryParams - Represents parameters of answerInlineQuery method.
type AnswerInlineQueryParams struct {
	// InlineQueryID - Unique identifier for the answered query
	InlineQueryID string `json:"inline_query_id"`

	// Results - A JSON-serialized array of results for the inline query
	Results []InlineQueryResult `json:"results"`

	// CacheTime - Optional. The maximum amount of time in seconds that the result of the inline query may be
	// cached on the server. Defaults to 300.
	CacheTime int `json:"cache_time,omitempty"`

	// IsPersonal - Optional. Pass True, if results may be cached on the server side only for the user that sent
	// the query. By default, results may be returned to any user who sends the same query
	IsPersonal bool `json:"is_personal,omitempty"`

	// NextOffset - Optional. Pass the offset that a client should send in the next query with the same text to
	// receive more results. Pass an empty string if there are no more results or if you don't support pagination.
	// Offset length can't exceed 64 bytes.
	NextOffset string `json:"next_offset,omitempty"`

	// SwitchPmText - Optional. If passed, clients will display a button with specified text that switches the
	// user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
	SwitchPmText string `json:"switch_pm_text,omitempty"`

	// SwitchPmParameter - Optional. Deep-linking (https://core.telegram.org/bots#deep-linking) parameter for the
	// /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _
	// and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to
	// their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube
	// account' button above the results, or even before showing any. The user presses the button, switches to a
	// private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an
	// OAuth link. Once done, the bot can offer a switch_inline
	// (https://core.telegram.org/bots/api#inlinekeyboardmarkup) button so that the user can easily return to the
	// chat where they wanted to use the bot's inline capabilities.
	SwitchPmParameter string `json:"switch_pm_parameter,omitempty"`
}

// AnswerInlineQuery - Use this method to send answers to an inline query. On success, True is returned.No
// more than 50 results per query are allowed.
func (b *Bot) AnswerInlineQuery(params *AnswerInlineQueryParams) error {
	err := b.performRequest("answerInlineQuery", params, nil)
	if err != nil {
		return fmt.Errorf("answerInlineQuery(): %w", err)
	}

	return nil
}

// SendInvoiceParams - Represents parameters of sendInvoice method.
type SendInvoiceParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID ChatID `json:"chat_id"`

	// Title - Product name, 1-32 characters
	Title string `json:"title"`

	// Description - Product description, 1-255 characters
	Description string `json:"description"`

	// Payload - Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your
	// internal processes.
	Payload string `json:"payload"`

	// ProviderToken - Payments provider token, obtained via BotFather (https://t.me/botfather)
	ProviderToken string `json:"provider_token"`

	// Currency - Three-letter ISO 4217 currency code, see more on currencies
	// (https://core.telegram.org/bots/payments#supported-currencies)
	Currency string `json:"currency"`

	// Prices - Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.)
	Prices []LabeledPrice `json:"prices"`

	// MaxTipAmount - Optional. The maximum accepted amount for tips in the smallest units of the currency
	// (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the
	// exp parameter in currencies.json (https://core.telegram.org/bots/payments/currencies.json), it shows the
	// number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	MaxTipAmount int `json:"max_tip_amount,omitempty"`

	// SuggestedTipAmounts - Optional. A JSON-serialized array of suggested amounts of tips in the smallest units
	// of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested
	// tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`

	// StartParameter - Optional. Unique deep-linking parameter. If left empty, forwarded copies of the sent
	// message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the
	// same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to
	// the bot (instead of a Pay button), with the value used as the start parameter
	StartParameter string `json:"start_parameter,omitempty"`

	// ProviderData - Optional. A JSON-serialized data about the invoice, which will be shared with the payment
	// provider. A detailed description of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// PhotoURL - Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing
	// image for a service. People like it better when they see what they are paying for.
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

	// NeedShippingAddress - Optional. Pass True, if you require the user's shipping address to complete the
	// order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// SendPhoneNumberToProvider - Optional. Pass True, if user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// SendEmailToProvider - Optional. Pass True, if user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// IsFlexible - Optional. Pass True, if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating). If empty, one 'Pay total price'
	// button will be shown. If not empty, the first button must be a Pay button.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// SendInvoice - Use this method to send invoices. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendInvoice(params *SendInvoiceParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendInvoice", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendInvoice(): %w", err)
	}

	return message, nil
}

// AnswerShippingQueryParams - Represents parameters of answerShippingQuery method.
type AnswerShippingQueryParams struct {
	// ShippingQueryID - Unique identifier for the query to be answered
	ShippingQueryID string `json:"shipping_query_id"`

	// Ok - Specify True if delivery to the specified address is possible and False if there are any problems
	// (for example, if delivery to the specified address is not possible)
	Ok bool `json:"ok"`

	// ShippingOptions - Optional. Required if ok is True. A JSON-serialized array of available shipping options.
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`

	// ErrorMessage - Optional. Required if ok is False. Error message in human readable form that explains why
	// it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable').
	// Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

// AnswerShippingQuery - If you sent an invoice requesting a shipping address and the parameter is_flexible
// was specified, the Bot API will send an Update (https://core.telegram.org/bots/api#update) with a
// shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
func (b *Bot) AnswerShippingQuery(params *AnswerShippingQueryParams) error {
	err := b.performRequest("answerShippingQuery", params, nil)
	if err != nil {
		return fmt.Errorf("answerShippingQuery(): %w", err)
	}

	return nil
}

// AnswerPreCheckoutQueryParams - Represents parameters of answerPreCheckoutQuery method.
type AnswerPreCheckoutQueryParams struct {
	// PreCheckoutQueryID - Unique identifier for the query to be answered
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`

	// Ok - Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed
	// with the order. Use False if there are any problems.
	Ok bool `json:"ok"`

	// ErrorMessage - Optional. Required if ok is False. Error message in human readable form that explains the
	// reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing
	// black T-shirts while you were busy filling out your payment details. Please choose a different color or
	// garment!"). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

// AnswerPreCheckoutQuery - Once the user has confirmed their payment and shipping details, the Bot API sends
// the final confirmation in the form of an Update (https://core.telegram.org/bots/api#update) with the field
// pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned.
// Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (b *Bot) AnswerPreCheckoutQuery(params *AnswerPreCheckoutQueryParams) error {
	err := b.performRequest("answerPreCheckoutQuery", params, nil)
	if err != nil {
		return fmt.Errorf("answerPreCheckoutQuery(): %w", err)
	}

	return nil
}

// SetPassportDataErrorsParams - Represents parameters of setPassportDataErrors method.
type SetPassportDataErrorsParams struct {
	// UserID - User identifier
	UserID int64 `json:"user_id"`

	// Errors - A JSON-serialized array describing the errors
	Errors []PassportElementError `json:"errors"`
}

// SetPassportDataErrors - Informs a user that some of the Telegram Passport elements they provided contains
// errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents
// of the field for which you returned the error must change). Returns True on success.
func (b *Bot) SetPassportDataErrors(params *SetPassportDataErrorsParams) error {
	err := b.performRequest("setPassportDataErrors", params, nil)
	if err != nil {
		return fmt.Errorf("setPassportDataErrors(): %w", err)
	}

	return nil
}

// SendGameParams - Represents parameters of sendGame method.
type SendGameParams struct {
	// ChatID - Unique identifier for the target chat
	// Note: Should be int64 not ChatID according to documentation (https://core.telegram.org/bots/api#sendgame)
	ChatID int64 `json:"chat_id"`

	// GameShortName - Short name of the game, serves as the unique identifier for the game. Set up your games
	// via BotFather (https://t.me/botfather).
	GameShortName string `json:"game_short_name"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID - Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply - Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating). If empty, one 'Play game_title'
	// button will be shown. If not empty, the first button must launch the game.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// SendGame - Use this method to send a game. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendGame(params *SendGameParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendGame", params, &message)
	if err != nil {
		return nil, fmt.Errorf("sendGame(): %w", err)
	}

	return message, nil
}

// SetGameScoreParams - Represents parameters of setGameScore method.
type SetGameScoreParams struct {
	// UserID - User identifier
	UserID int64 `json:"user_id"`

	// Score - New score, must be non-negative
	Score int `json:"score"`

	// Force - Optional. Pass True, if the high score is allowed to decrease. This can be useful when fixing
	// mistakes or banning cheaters
	Force bool `json:"force,omitempty"`

	// DisableEditMessage - Optional. Pass True, if the game message should not be automatically edited to
	// include the current scoreboard
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`

	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatID int64 `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the sent message
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the inline
	// message
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// SetGameScore - Use this method to set the score of the specified user in a game message. On success, if
// the message is not an inline message, the Message (https://core.telegram.org/bots/api#message) is returned,
// otherwise True is returned. Returns an error, if the new score is not greater than the user's current score
// in the chat and force is False.
func (b *Bot) SetGameScore(params *SetGameScoreParams) (*Message, error) {
	var message *Message
	err := b.performRequest("setGameScore", params, &message)
	if err != nil {
		return nil, fmt.Errorf("setGameScore(): %w", err)
	}

	return message, nil
}

// GetGameHighScoresParams - Represents parameters of getGameHighScores method.
type GetGameHighScoresParams struct {
	// UserID - Target user ID
	UserID int64 `json:"user_id"`

	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatID int64 `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the sent message
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the inline
	// message
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// GetGameHighScores - Use this method to get data for high score tables. Will return the score of the
// specified user and several of their neighbors in a game. On success, returns an Array of GameHighScore
// (https://core.telegram.org/bots/api#gamehighscore) objects.
func (b *Bot) GetGameHighScores(params *GetGameHighScoresParams) ([]GameHighScore, error) {
	var gameHighScores []GameHighScore
	err := b.performRequest("getGameHighScores", params, &gameHighScores)
	if err != nil {
		return nil, fmt.Errorf("getGameHighScores(): %w", err)
	}

	return gameHighScores, nil
}
