package telego

import (
	"fmt"

	ta "github.com/mymmrac/telego/telegoapi"
)

// GetUpdatesParams - Represents parameters of getUpdates method.
type GetUpdatesParams struct {
	// Offset - Optional. Identifier of the first update to be returned. Must be greater by one than the highest
	// among the identifiers of previously received updates. By default, updates starting with the earliest
	// unconfirmed update are returned. An update is considered confirmed as soon as getUpdates
	// (https://core.telegram.org/bots/api#getupdates) is called with an offset higher than its update_id. The
	// negative offset can be specified to retrieve updates starting from -offset update from the end of the updates
	// queue. All previous updates will be forgotten.
	Offset int `json:"offset,omitempty"`

	// Limit - Optional. Limits the number of updates to be retrieved. Values between 1-100 are accepted.
	// Defaults to 100.
	Limit int `json:"limit,omitempty"`

	// Timeout - Optional. Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should
	// be positive, short polling should be used for testing purposes only.
	Timeout int `json:"timeout,omitempty"`

	// AllowedUpdates - Optional. A JSON-serialized list of the update types you want your bot to receive. For
	// example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types.
	// See Update (https://core.telegram.org/bots/api#update) for a complete list of available update types. Specify
	// an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count
	// (default). If not specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted
	// updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// Update types you want your bot to receive
const (
	MessageUpdates            = "message"
	EditedMessageUpdates      = "edited_message"
	ChannelPostUpdates        = "channel_post"
	EditedChannelPostUpdates  = "edited_channel_post"
	MessageReaction           = "message_reaction"
	MessageReactionCount      = "message_reaction_count"
	InlineQueryUpdates        = "inline_query"
	ChosenInlineResultUpdates = "chosen_inline_result"
	CallbackQueryUpdates      = "callback_query"
	ShippingQueryUpdates      = "shipping_query"
	PreCheckoutQueryUpdates   = "pre_checkout_query"
	PollUpdates               = "poll"
	PollAnswerUpdates         = "poll_answer"
	MyChatMemberUpdates       = "my_chat_member"
	ChatMemberUpdates         = "chat_member"
	ChatJoinRequestUpdates    = "chat_join_request"
)

// GetUpdates - Use this method to receive incoming updates using long polling (wiki
// (https://en.wikipedia.org/wiki/Push_technology#Long_polling)). Returns an Array of Update
// (https://core.telegram.org/bots/api#update) objects.
func (b *Bot) GetUpdates(params *GetUpdatesParams) ([]Update, error) {
	var updates []Update
	err := b.performRequest("getUpdates", params, &updates)
	if err != nil {
		return nil, fmt.Errorf("telego: getUpdates(): %w", err)
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

	// MaxConnections - Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook
	// for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and
	// higher values to increase your bot's throughput.
	MaxConnections int `json:"max_connections,omitempty"`

	// AllowedUpdates - Optional. A JSON-serialized list of the update types you want your bot to receive. For
	// example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types.
	// See Update (https://core.telegram.org/bots/api#update) for a complete list of available update types. Specify
	// an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count
	// (default). If not specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted
	// updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`

	// DropPendingUpdates - Optional. Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`

	// SecretToken - Optional. A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in
	// every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is
	// useful to ensure that the request comes from a webhook set by you.
	SecretToken string `json:"secret_token,omitempty"`
}

func (p *SetWebhookParams) fileParameters() map[string]ta.NamedReader {
	fp := make(map[string]ta.NamedReader)

	if p.Certificate != nil {
		fp["certificate"] = p.Certificate.File
	}

	return fp
}

// SetWebhook - Use this method to specify a URL and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing
// a JSON-serialized Update (https://core.telegram.org/bots/api#update). In case of an unsuccessful request, we
// will give up after a reasonable amount of attempts. Returns True on success.
// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter
// secret_token. If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the
// secret token as content.
func (b *Bot) SetWebhook(params *SetWebhookParams) error {
	err := b.performRequest("setWebhook", params)
	if err != nil {
		return fmt.Errorf("telego: setWebhook(): %w", err)
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
	err := b.performRequest("deleteWebhook", params)
	if err != nil {
		return fmt.Errorf("telego: deleteWebhook(): %w", err)
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
		return nil, fmt.Errorf("telego: getWebhookInfo(): %w", err)
	}

	return webhookInfo, nil
}

// GetMe - A simple method for testing your bot's authentication token. Requires no parameters. Returns basic
// information about the bot in form of a User (https://core.telegram.org/bots/api#user) object.
func (b *Bot) GetMe() (*User, error) {
	var user *User
	err := b.performRequest("getMe", nil, &user)
	if err != nil {
		return nil, fmt.Errorf("telego: getMe(): %w", err)
	}

	return user, nil
}

// LogOut - Use this method to log out from the cloud Bot API server before launching the bot locally. You
// must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive
// updates. After a successful call, you can immediately log in on a local server, but will not be able to log
// in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
func (b *Bot) LogOut() error {
	err := b.performRequest("logOut", nil)
	if err != nil {
		return fmt.Errorf("telego: logOut(): %w", err)
	}

	return nil
}

// Close - Use this method to close the bot instance before moving it from one local server to another. You
// need to delete the webhook before calling this method to ensure that the bot isn't launched again after
// server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns
// True on success. Requires no parameters.
func (b *Bot) Close() error {
	err := b.performRequest("close", nil)
	if err != nil {
		return fmt.Errorf("telego: close(): %w", err)
	}

	return nil
}

// SendMessageParams - Represents parameters of sendMessage method.
type SendMessageParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Text - Text of the message to be sent, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// ParseMode - Optional. Mode for parsing entities in the message text. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Entities - Optional. A JSON-serialized list of special entities that appear in message text, which can be
	// specified instead of parse_mode
	Entities []MessageEntity `json:"entities,omitempty"`

	// LinkPreviewOptions - Optional. Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
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
		return nil, fmt.Errorf("telego: sendMessage(): %w", err)
	}

	return message, nil
}

// ForwardMessageParams - Represents parameters of forwardMessage method.
type ForwardMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// FromChatID - Unique identifier for the chat where the original message was sent (or channel username in
	// the format @channel_username)
	FromChatID ChatID `json:"from_chat_id"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the forwarded message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageID - Message identifier in the chat specified in from_chat_id
	MessageID int `json:"message_id"`
}

// ForwardMessage - Use this method to forward messages of any kind. Service messages and messages with
// protected content can't be forwarded. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) ForwardMessage(params *ForwardMessageParams) (*Message, error) {
	var message *Message
	err := b.performRequest("forwardMessage", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: forwardMessage(): %w", err)
	}

	return message, nil
}

// ForwardMessagesParams - Represents parameters of forwardMessages method.
type ForwardMessagesParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// FromChatID - Unique identifier for the chat where the original messages were sent (or channel username in
	// the format @channel_username)
	FromChatID ChatID `json:"from_chat_id"`

	// MessageIDs - A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to forward.
	// The identifiers must be specified in a strictly increasing order.
	MessageIDs []int `json:"message_ids"`

	// DisableNotification - Optional. Sends the messages silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the forwarded messages from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`
}

// ForwardMessages - Use this method to forward multiple messages of any kind. If some of the specified
// messages can't be found or forwarded, they are skipped. Service messages and messages with protected content
// can't be forwarded. Album grouping is kept for forwarded messages. On success, an array of MessageID
// (https://core.telegram.org/bots/api#messageid) of the sent messages is returned.
func (b *Bot) ForwardMessages(params *ForwardMessagesParams) (*MessageID, error) {
	var messageID *MessageID
	err := b.performRequest("forwardMessages", params, &messageID)
	if err != nil {
		return nil, fmt.Errorf("telego: forwardMessages(): %w", err)
	}

	return messageID, nil
}

// CopyMessageParams - Represents parameters of copyMessage method.
type CopyMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// FromChatID - Unique identifier for the chat where the original message was sent (or channel username in
	// the format @channel_username)
	FromChatID ChatID `json:"from_chat_id"`

	// MessageID - Message identifier in the chat specified in from_chat_id
	MessageID int `json:"message_id"`

	// Caption - Optional. New caption for media, 0-1024 characters after entities parsing. If not specified,
	// the original caption is kept
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the new caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the new caption,
	// which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media.
	// Ignored if a new caption isn't specified.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// CopyMessage - Use this method to copy messages of any kind. Service messages, giveaway messages, giveaway
// winners messages, and invoice messages can't be copied. A quiz poll (https://core.telegram.org/bots/api#poll)
// can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous
// to the method forwardMessage (https://core.telegram.org/bots/api#forwardmessage), but the copied message
// doesn't have a link to the original message. Returns the MessageID
// (https://core.telegram.org/bots/api#messageid) of the sent message on success.
func (b *Bot) CopyMessage(params *CopyMessageParams) (*MessageID, error) {
	var messageID *MessageID
	err := b.performRequest("copyMessage", params, &messageID)
	if err != nil {
		return nil, fmt.Errorf("telego: copyMessage(): %w", err)
	}

	return messageID, nil
}

// CopyMessagesParams - Represents parameters of copyMessages method.
type CopyMessagesParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// FromChatID - Unique identifier for the chat where the original messages were sent (or channel username in
	// the format @channel_username)
	FromChatID ChatID `json:"from_chat_id"`

	// MessageIDs - A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to copy.
	// The identifiers must be specified in a strictly increasing order.
	MessageIDs []int `json:"message_ids"`

	// DisableNotification - Optional. Sends the messages silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent messages from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// RemoveCaption - Optional. Pass True to copy the messages without their captions
	RemoveCaption bool `json:"remove_caption,omitempty"`
}

// CopyMessages - Use this method to copy messages of any kind. If some of the specified messages can't be
// found or copied, they are skipped. Service messages, giveaway messages, giveaway winners messages, and
// invoice messages can't be copied. A quiz poll (https://core.telegram.org/bots/api#poll) can be copied only if
// the value of the field correct_option_id is known to the bot. The method is analogous to the method
// forwardMessages (https://core.telegram.org/bots/api#forwardmessages), but the copied messages don't have a
// link to the original message. Album grouping is kept for copied messages. On success, an array of MessageID
// (https://core.telegram.org/bots/api#messageid) of the sent messages is returned.
func (b *Bot) CopyMessages(params *CopyMessagesParams) (*MessageID, error) {
	var messageID *MessageID
	err := b.performRequest("copyMessages", params, &messageID)
	if err != nil {
		return nil, fmt.Errorf("telego: copyMessages(): %w", err)
	}

	return messageID, nil
}

// SendPhotoParams - Represents parameters of sendPhoto method.
type SendPhotoParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Photo - Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new
	// photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must
	// not exceed 10000 in total. Width and height ratio must be at most 20. More information on Sending Files »
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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// HasSpoiler - Optional. Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendPhotoParams) fileParameters() map[string]ta.NamedReader {
	return map[string]ta.NamedReader{
		"photo": p.Photo.File,
	}
}

// SendPhoto - Use this method to send photos. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendPhoto(params *SendPhotoParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendPhoto", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendPhoto(): %w", err)
	}

	return message, nil
}

// SendAudioParams - Represents parameters of sendAudio method.
type SendAudioParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Audio - Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or
	// upload a new one using multipart/form-data. More information on Sending Files »
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

	// Thumbnail - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumbnail *InputFile `json:"thumbnail,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendAudioParams) fileParameters() map[string]ta.NamedReader {
	fp := make(map[string]ta.NamedReader)

	fp["audio"] = p.Audio.File
	if p.Thumbnail != nil {
		fp["thumbnail"] = p.Thumbnail.File
	}

	return fp
}

// SendAudio - Use this method to send audio files, if you want Telegram clients to display them in the music
// player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned. Bots can currently send audio files of up to 50 MB
// in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice (https://core.telegram.org/bots/api#sendvoice) method instead.
func (b *Bot) SendAudio(params *SendAudioParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendAudio", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendAudio(): %w", err)
	}

	return message, nil
}

// SendDocumentParams - Represents parameters of sendDocument method.
type SendDocumentParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Document - File to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one
	// using multipart/form-data. More information on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	Document InputFile `json:"document"`

	// Thumbnail - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumbnail *InputFile `json:"thumbnail,omitempty"`

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

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendDocumentParams) fileParameters() map[string]ta.NamedReader {
	fp := make(map[string]ta.NamedReader)

	fp["document"] = p.Document.File
	if p.Thumbnail != nil {
		fp["thumbnail"] = p.Thumbnail.File
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
		return nil, fmt.Errorf("telego: sendDocument(): %w", err)
	}

	return message, nil
}

// SendVideoParams - Represents parameters of sendVideo method.
type SendVideoParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Video - Video to send. Pass a file_id as String to send a video that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new
	// video using multipart/form-data. More information on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	Video InputFile `json:"video"`

	// Duration - Optional. Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`

	// Width - Optional. Video width
	Width int `json:"width,omitempty"`

	// Height - Optional. Video height
	Height int `json:"height,omitempty"`

	// Thumbnail - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumbnail *InputFile `json:"thumbnail,omitempty"`

	// Caption - Optional. Video caption (may also be used when resending videos by file_id), 0-1024 characters
	// after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the video caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// HasSpoiler - Optional. Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`

	// SupportsStreaming - Optional. Pass True if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendVideoParams) fileParameters() map[string]ta.NamedReader {
	fp := make(map[string]ta.NamedReader)

	fp["video"] = p.Video.File
	if p.Thumbnail != nil {
		fp["thumbnail"] = p.Thumbnail.File
	}

	return fp
}

// SendVideo - Use this method to send video files, Telegram clients support MPEG4 videos (other formats may
// be sent as Document (https://core.telegram.org/bots/api#document)). On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned. Bots can currently send video files of up to 50 MB
// in size, this limit may be changed in the future.
func (b *Bot) SendVideo(params *SendVideoParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendVideo", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendVideo(): %w", err)
	}

	return message, nil
}

// SendAnimationParams - Represents parameters of sendAnimation method.
type SendAnimationParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Animation - Animation to send. Pass a file_id as String to send an animation that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or
	// upload a new animation using multipart/form-data. More information on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	Animation InputFile `json:"animation"`

	// Duration - Optional. Duration of sent animation in seconds
	Duration int `json:"duration,omitempty"`

	// Width - Optional. Animation width
	Width int `json:"width,omitempty"`

	// Height - Optional. Animation height
	Height int `json:"height,omitempty"`

	// Thumbnail - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumbnail *InputFile `json:"thumbnail,omitempty"`

	// Caption - Optional. Animation caption (may also be used when resending animation by file_id), 0-1024
	// characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the animation caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// HasSpoiler - Optional. Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendAnimationParams) fileParameters() map[string]ta.NamedReader {
	fp := make(map[string]ta.NamedReader)

	fp["animation"] = p.Animation.File
	if p.Thumbnail != nil {
		fp["thumbnail"] = p.Thumbnail.File
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
		return nil, fmt.Errorf("telego: sendAnimation(): %w", err)
	}

	return message, nil
}

// SendVoiceParams - Represents parameters of sendVoice method.
type SendVoiceParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Voice - Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one
	// using multipart/form-data. More information on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
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

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendVoiceParams) fileParameters() map[string]ta.NamedReader {
	return map[string]ta.NamedReader{
		"voice": p.Voice.File,
	}
}

// SendVoice - Use this method to send audio files, if you want Telegram clients to display the file as a
// playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3
// format, or in .M4A format (other formats may be sent as Audio (https://core.telegram.org/bots/api#audio) or
// Document (https://core.telegram.org/bots/api#document)). On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned. Bots can currently send voice messages of up to 50
// MB in size, this limit may be changed in the future.
func (b *Bot) SendVoice(params *SendVoiceParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendVoice", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendVoice(): %w", err)
	}

	return message, nil
}

// SendVideoNoteParams - Represents parameters of sendVideoNote method.
type SendVideoNoteParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// VideoNote - Video note to send. Pass a file_id as String to send a video note that exists on the Telegram
	// servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files). Sending video notes by a URL is currently unsupported
	VideoNote InputFile `json:"video_note"`

	// Duration - Optional. Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`

	// Length - Optional. Video width and height, i.e. diameter of the video message
	Length int `json:"length,omitempty"`

	// Thumbnail - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumbnail *InputFile `json:"thumbnail,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendVideoNoteParams) fileParameters() map[string]ta.NamedReader {
	fp := make(map[string]ta.NamedReader)

	fp["video_note"] = p.VideoNote.File
	if p.Thumbnail != nil {
		fp["thumbnail"] = p.Thumbnail.File
	}

	return fp
}

// SendVideoNote - As of v.4.0 (https://telegram.org/blog/video-messages-and-telescope), Telegram clients
// support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On
// success, the sent Message (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendVideoNote(params *SendVideoNoteParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendVideoNote", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendVideoNote(): %w", err)
	}

	return message, nil
}

// SendMediaGroupParams - Represents parameters of sendMediaGroup method.
type SendMediaGroupParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Media - A JSON-serialized array describing messages to be sent, must include 2-10 items
	Media []InputMedia `json:"media"`

	// DisableNotification - Optional. Sends messages silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent messages from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
}

func (p *SendMediaGroupParams) fileParameters() map[string]ta.NamedReader {
	fp := make(map[string]ta.NamedReader)

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
		return nil, fmt.Errorf("telego: sendMediaGroup(): %w", err)
	}

	return messages, nil
}

// SendLocationParams - Represents parameters of sendLocation method.
type SendLocationParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Latitude - Latitude of the location
	Latitude float64 `json:"latitude"`

	// Longitude - Longitude of the location
	Longitude float64 `json:"longitude"`

	// HorizontalAccuracy - Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// LivePeriod - Optional. Period in seconds during which the location will be updated (see Live Locations
	// (https://telegram.org/blog/live-locations), should be between 60 and 86400, or 0x7FFFFFFF for live locations
	// that can be edited indefinitely.
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

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendLocation - Use this method to send point on the map. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendLocation(params *SendLocationParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendLocation", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendLocation(): %w", err)
	}

	return message, nil
}

// SendVenueParams - Represents parameters of sendVenue method.
type SendVenueParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

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

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendVenue - Use this method to send information about a venue. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendVenue(params *SendVenueParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendVenue", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendVenue(): %w", err)
	}

	return message, nil
}

// SendContactParams - Represents parameters of sendContact method.
type SendContactParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

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

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendContact - Use this method to send phone contacts. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendContact(params *SendContactParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendContact", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendContact(): %w", err)
	}

	return message, nil
}

// SendPollParams - Represents parameters of sendPoll method.
type SendPollParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Question - Poll question, 1-300 characters
	Question string `json:"question"`

	// QuestionParseMode - Optional. Mode for parsing entities in the question. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details. Currently, only custom emoji
	// entities are allowed
	QuestionParseMode string `json:"question_parse_mode,omitempty"`

	// QuestionEntities - Optional. A JSON-serialized list of special entities that appear in the poll question.
	// It can be specified instead of question_parse_mode
	QuestionEntities []MessageEntity `json:"question_entities,omitempty"`

	// Options - A JSON-serialized list of 2-10 answer options
	Options []InputPollOption `json:"options"`

	// IsAnonymous - Optional. True, if the poll needs to be anonymous, defaults to True
	IsAnonymous *bool `json:"is_anonymous,omitempty"`

	// Type - Optional. Poll type, “quiz” or “regular”, defaults to “regular”
	Type string `json:"type,omitempty"`

	// AllowsMultipleAnswers - Optional. True, if the poll allows multiple answers, ignored for polls in quiz
	// mode, defaults to False
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`

	// CorrectOptionID - Optional. 0-based identifier of the correct answer option, required for polls in quiz
	// mode
	CorrectOptionID *int `json:"correct_option_id,omitempty"`

	// Explanation - Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp
	// icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	Explanation string `json:"explanation,omitempty"`

	// ExplanationParseMode - Optional. Mode for parsing entities in the explanation. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ExplanationParseMode string `json:"explanation_parse_mode,omitempty"`

	// ExplanationEntities - Optional. A JSON-serialized list of special entities that appear in the poll
	// explanation. It can be specified instead of explanation_parse_mode
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`

	// OpenPeriod - Optional. Amount of time in seconds the poll will be active after creation, 5-600. Can't be
	// used together with close_date.
	OpenPeriod int `json:"open_period,omitempty"`

	// CloseDate - Optional. Point in time (Unix timestamp) when the poll will be automatically closed. Must be
	// at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
	CloseDate int64 `json:"close_date,omitempty"`

	// IsClosed - Optional. Pass True if the poll needs to be immediately closed. This can be useful for poll
	// preview.
	IsClosed bool `json:"is_closed,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendPoll - Use this method to send a native poll. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendPoll(params *SendPollParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendPoll", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendPoll(): %w", err)
	}

	return message, nil
}

// SendDiceParams - Represents parameters of sendDice method.
type SendDiceParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Emoji - Optional. Emoji on which the dice throw animation is based. Currently, must be one of “🎲”,
	// “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”. Dice can have values 1-6 for “🎲”,
	// “🎯” and “🎳”, values 1-5 for “🏀” and “⚽”, and values 1-64 for “🎰”. Defaults
	// to “🎲”
	Emoji string `json:"emoji,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendDice - Use this method to send an animated emoji that will display a random value. On success, the
// sent Message (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendDice(params *SendDiceParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendDice", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendDice(): %w", err)
	}

	return message, nil
}

// SendChatActionParams - Represents parameters of sendChatAction method.
type SendChatActionParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// action will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread; for supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Action - Type of action to broadcast. Choose one, depending on what the user is about to receive: typing
	// for text messages (https://core.telegram.org/bots/api#sendmessage), upload_photo for photos
	// (https://core.telegram.org/bots/api#sendphoto), record_video or upload_video for videos
	// (https://core.telegram.org/bots/api#sendvideo), record_voice or upload_voice for voice notes
	// (https://core.telegram.org/bots/api#sendvoice), upload_document for general files
	// (https://core.telegram.org/bots/api#senddocument), choose_sticker for stickers
	// (https://core.telegram.org/bots/api#sendsticker), find_location for location data
	// (https://core.telegram.org/bots/api#sendlocation), record_video_note or upload_video_note for video notes
	// (https://core.telegram.org/bots/api#sendvideonote).
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
// Example: The ImageBot (https://t.me/imagebot) needs some time to process a request and upload the image.
// Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use
// sendChatAction (https://core.telegram.org/bots/api#sendchataction) with action = upload_photo. The user will
// see a “sending photo” status for the bot.
// We only recommend using this method when a response from the bot will take a noticeable amount of time to
// arrive.
func (b *Bot) SendChatAction(params *SendChatActionParams) error {
	err := b.performRequest("sendChatAction", params)
	if err != nil {
		return fmt.Errorf("telego: sendChatAction(): %w", err)
	}

	return nil
}

// SetMessageReactionParams - Represents parameters of setMessageReaction method.
type SetMessageReactionParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageID - Identifier of the target message. If the message belongs to a media group, the reaction is
	// set to the first non-deleted message in the group instead.
	MessageID int `json:"message_id"`

	// Reaction - Optional. A JSON-serialized list of reaction types to set on the message. Currently, as
	// non-premium users, bots can set up to one reaction per message. A custom emoji reaction can be used if it is
	// either already present on the message or explicitly allowed by chat administrators.
	Reaction []ReactionType `json:"reaction,omitempty"`

	// IsBig - Optional. Pass True to set the reaction with a big animation
	IsBig bool `json:"is_big,omitempty"`
}

// SetMessageReaction - Use this method to change the chosen reactions on a message. Service messages can't
// be reacted to. Automatically forwarded messages from a channel to its discussion group have the same
// available reactions as messages in the channel. Returns True on success.
func (b *Bot) SetMessageReaction(params *SetMessageReactionParams) error {
	err := b.performRequest("setMessageReaction", params)
	if err != nil {
		return fmt.Errorf("telego: setMessageReaction(): %w", err)
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

	// Limit - Optional. Limits the number of photos to be retrieved. Values between 1-100 are accepted.
	// Defaults to 100.
	Limit int `json:"limit,omitempty"`
}

// GetUserProfilePhotos - Use this method to get a list of profile pictures for a user. Returns a
// UserProfilePhotos (https://core.telegram.org/bots/api#userprofilephotos) object.
func (b *Bot) GetUserProfilePhotos(params *GetUserProfilePhotosParams) (*UserProfilePhotos, error) {
	var userProfilePhotos *UserProfilePhotos
	err := b.performRequest("getUserProfilePhotos", params, &userProfilePhotos)
	if err != nil {
		return nil, fmt.Errorf("telego: getUserProfilePhotos(): %w", err)
	}

	return userProfilePhotos, nil
}

// GetFileParams - Represents parameters of getFile method.
type GetFileParams struct {
	// FileID - File identifier to get information about
	FileID string `json:"file_id"`
}

// GetFile - Use this method to get basic information about a file and prepare it for downloading. For the
// moment, bots can download files of up to 20MB in size. On success, a File
// (https://core.telegram.org/bots/api#file) object is returned. The file can then be downloaded via the link
// https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is
// guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested
// by calling getFile (https://core.telegram.org/bots/api#getfile) again.
func (b *Bot) GetFile(params *GetFileParams) (*File, error) {
	var file *File
	err := b.performRequest("getFile", params, &file)
	if err != nil {
		return nil, fmt.Errorf("telego: getFile(): %w", err)
	}

	return file, nil
}

// BanChatMemberParams - Represents parameters of banChatMember method.
type BanChatMemberParams struct {
	// ChatID - Unique identifier for the target group or username of the target supergroup or channel (in the
	// format @channel_username)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// UntilDate - Optional. Date when the user will be unbanned; Unix time. If user is banned for more than 366
	// days or less than 30 seconds from the current time they are considered to be banned forever. Applied for
	// supergroups and channels only.
	UntilDate int64 `json:"until_date,omitempty"`

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
	err := b.performRequest("banChatMember", params)
	if err != nil {
		return fmt.Errorf("telego: banChatMember(): %w", err)
	}

	return nil
}

// UnbanChatMemberParams - Represents parameters of unbanChatMember method.
type UnbanChatMemberParams struct {
	// ChatID - Unique identifier for the target group or username of the target supergroup or channel (in the
	// format @channel_username)
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
	err := b.performRequest("unbanChatMember", params)
	if err != nil {
		return fmt.Errorf("telego: unbanChatMember(): %w", err)
	}

	return nil
}

// RestrictChatMemberParams - Represents parameters of restrictChatMember method.
type RestrictChatMemberParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// Permissions - A JSON-serialized object for new user permissions
	Permissions ChatPermissions `json:"permissions"`

	// UseIndependentChatPermissions - Optional. Pass True if chat permissions are set independently. Otherwise,
	// the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages,
	// can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and
	// can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
	UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`

	// UntilDate - Optional. Date when restrictions will be lifted for the user; Unix time. If user is
	// restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be
	// restricted forever
	UntilDate int64 `json:"until_date,omitempty"`
}

// RestrictChatMember - Use this method to restrict a user in a supergroup. The bot must be an administrator
// in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all
// permissions to lift restrictions from a user. Returns True on success.
func (b *Bot) RestrictChatMember(params *RestrictChatMemberParams) error {
	err := b.performRequest("restrictChatMember", params)
	if err != nil {
		return fmt.Errorf("telego: restrictChatMember(): %w", err)
	}

	return nil
}

// PromoteChatMemberParams - Represents parameters of promoteChatMember method.
type PromoteChatMemberParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// IsAnonymous - Optional. Pass True if the administrator's presence in the chat is hidden
	IsAnonymous *bool `json:"is_anonymous,omitempty"`

	// CanManageChat - Optional. Pass True if the administrator can access the chat event log, get boost list,
	// see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other
	// administrator privilege.
	CanManageChat *bool `json:"can_manage_chat,omitempty"`

	// CanDeleteMessages - Optional. Pass True if the administrator can delete messages of other users
	CanDeleteMessages *bool `json:"can_delete_messages,omitempty"`

	// CanManageVideoChats - Optional. Pass True if the administrator can manage video chats
	CanManageVideoChats *bool `json:"can_manage_video_chats,omitempty"`

	// CanRestrictMembers - Optional. Pass True if the administrator can restrict, ban or unban chat members, or
	// access supergroup statistics
	CanRestrictMembers *bool `json:"can_restrict_members,omitempty"`

	// CanPromoteMembers - Optional. Pass True if the administrator can add new administrators with a subset of
	// their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by
	// administrators that were appointed by him)
	CanPromoteMembers *bool `json:"can_promote_members,omitempty"`

	// CanChangeInfo - Optional. Pass True if the administrator can change chat title, photo and other settings
	CanChangeInfo *bool `json:"can_change_info,omitempty"`

	// CanInviteUsers - Optional. Pass True if the administrator can invite new users to the chat
	CanInviteUsers *bool `json:"can_invite_users,omitempty"`

	// CanPostStories - Optional. Pass True if the administrator can post stories to the chat
	CanPostStories *bool `json:"can_post_stories,omitempty"`

	// CanEditStories - Optional. Pass True if the administrator can edit stories posted by other users, post
	// stories to the chat page, pin chat stories, and access the chat's story archive
	CanEditStories *bool `json:"can_edit_stories,omitempty"`

	// CanDeleteStories - Optional. Pass True if the administrator can delete stories posted by other users
	CanDeleteStories *bool `json:"can_delete_stories,omitempty"`

	// CanPostMessages - Optional. Pass True if the administrator can post messages in the channel, or access
	// channel statistics; for channels only
	CanPostMessages *bool `json:"can_post_messages,omitempty"`

	// CanEditMessages - Optional. Pass True if the administrator can edit messages of other users and can pin
	// messages; for channels only
	CanEditMessages *bool `json:"can_edit_messages,omitempty"`

	// CanPinMessages - Optional. Pass True if the administrator can pin messages; for supergroups only
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`

	// CanManageTopics - Optional. Pass True if the user is allowed to create, rename, close, and reopen forum
	// topics; for supergroups only
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
}

// PromoteChatMember - Use this method to promote or demote a user in a supergroup or a channel. The bot must
// be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass
// False for all boolean parameters to demote a user. Returns True on success.
func (b *Bot) PromoteChatMember(params *PromoteChatMemberParams) error {
	err := b.performRequest("promoteChatMember", params)
	if err != nil {
		return fmt.Errorf("telego: promoteChatMember(): %w", err)
	}

	return nil
}

// SetChatAdministratorCustomTitleParams - Represents parameters of setChatAdministratorCustomTitle method.
type SetChatAdministratorCustomTitleParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// CustomTitle - New custom title for the administrator; 0-16 characters, emoji are not allowed
	CustomTitle string `json:"custom_title"`
}

// SetChatAdministratorCustomTitle - Use this method to set a custom title for an administrator in a
// supergroup promoted by the bot. Returns True on success.
func (b *Bot) SetChatAdministratorCustomTitle(params *SetChatAdministratorCustomTitleParams) error {
	err := b.performRequest("setChatAdministratorCustomTitle", params)
	if err != nil {
		return fmt.Errorf("telego: setChatAdministratorCustomTitle(): %w", err)
	}

	return nil
}

// BanChatSenderChatParams - Represents parameters of banChatSenderChat method.
type BanChatSenderChatParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// SenderChatID - Unique identifier of the target sender chat
	SenderChatID int64 `json:"sender_chat_id"`
}

// BanChatSenderChat - Use this method to ban a channel chat in a supergroup or a channel. Until the chat is
// unbanned (https://core.telegram.org/bots/api#unbanchatsenderchat), the owner of the banned chat won't be able
// to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or
// channel for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) BanChatSenderChat(params *BanChatSenderChatParams) error {
	err := b.performRequest("banChatSenderChat", params)
	if err != nil {
		return fmt.Errorf("telego: banChatSenderChat(): %w", err)
	}

	return nil
}

// UnbanChatSenderChatParams - Represents parameters of unbanChatSenderChat method.
type UnbanChatSenderChatParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// SenderChatID - Unique identifier of the target sender chat
	SenderChatID int64 `json:"sender_chat_id"`
}

// UnbanChatSenderChat - Use this method to unban a previously banned channel chat in a supergroup or
// channel. The bot must be an administrator for this to work and must have the appropriate administrator
// rights. Returns True on success.
func (b *Bot) UnbanChatSenderChat(params *UnbanChatSenderChatParams) error {
	err := b.performRequest("unbanChatSenderChat", params)
	if err != nil {
		return fmt.Errorf("telego: unbanChatSenderChat(): %w", err)
	}

	return nil
}

// SetChatPermissionsParams - Represents parameters of setChatPermissions method.
type SetChatPermissionsParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// Permissions - A JSON-serialized object for new default chat permissions
	Permissions ChatPermissions `json:"permissions"`

	// UseIndependentChatPermissions - Optional. Pass True if chat permissions are set independently. Otherwise,
	// the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages,
	// can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and
	// can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
	UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`
}

// SetChatPermissions - Use this method to set default chat permissions for all members. The bot must be an
// administrator in the group or a supergroup for this to work and must have the can_restrict_members
// administrator rights. Returns True on success.
func (b *Bot) SetChatPermissions(params *SetChatPermissionsParams) error {
	err := b.performRequest("setChatPermissions", params)
	if err != nil {
		return fmt.Errorf("telego: setChatPermissions(): %w", err)
	}

	return nil
}

// ExportChatInviteLinkParams - Represents parameters of exportChatInviteLink method.
type ExportChatInviteLinkParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`
}

// ExportChatInviteLink - Use this method to generate a new primary invite link for a chat; any previously
// generated primary link is revoked. The bot must be an administrator in the chat for this to work and must
// have the appropriate administrator rights. Returns the new invite link as String on success.
func (b *Bot) ExportChatInviteLink(params *ExportChatInviteLinkParams) (*string, error) {
	var inviteLink *string
	err := b.performRequest("exportChatInviteLink", params, &inviteLink)
	if err != nil {
		return nil, fmt.Errorf("telego: exportChatInviteLink(): %w", err)
	}

	return inviteLink, nil
}

// CreateChatInviteLinkParams - Represents parameters of createChatInviteLink method.
type CreateChatInviteLinkParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// Name - Optional. Invite link name; 0-32 characters
	Name string `json:"name,omitempty"`

	// ExpireDate - Optional. Point in time (Unix timestamp) when the link will expire
	ExpireDate int64 `json:"expire_date,omitempty"`

	// MemberLimit - Optional. The maximum number of users that can be members of the chat simultaneously after
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
		return nil, fmt.Errorf("telego: createChatInviteLink(): %w", err)
	}

	return chatInviteLink, nil
}

// EditChatInviteLinkParams - Represents parameters of editChatInviteLink method.
type EditChatInviteLinkParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// InviteLink - The invite link to edit
	InviteLink string `json:"invite_link"`

	// Name - Optional. Invite link name; 0-32 characters
	Name string `json:"name,omitempty"`

	// ExpireDate - Optional. Point in time (Unix timestamp) when the link will expire
	ExpireDate int64 `json:"expire_date,omitempty"`

	// MemberLimit - Optional. The maximum number of users that can be members of the chat simultaneously after
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
		return nil, fmt.Errorf("telego: editChatInviteLink(): %w", err)
	}

	return chatInviteLink, nil
}

// RevokeChatInviteLinkParams - Represents parameters of revokeChatInviteLink method.
type RevokeChatInviteLinkParams struct {
	// ChatID - Unique identifier of the target chat or username of the target channel (in the format
	// @channel_username)
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
		return nil, fmt.Errorf("telego: revokeChatInviteLink(): %w", err)
	}

	return chatInviteLink, nil
}

// ApproveChatJoinRequestParams - Represents parameters of approveChatJoinRequest method.
type ApproveChatJoinRequestParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// ApproveChatJoinRequest - Use this method to approve a chat join request. The bot must be an administrator
// in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (b *Bot) ApproveChatJoinRequest(params *ApproveChatJoinRequestParams) error {
	err := b.performRequest("approveChatJoinRequest", params)
	if err != nil {
		return fmt.Errorf("telego: approveChatJoinRequest(): %w", err)
	}

	return nil
}

// DeclineChatJoinRequestParams - Represents parameters of declineChatJoinRequest method.
type DeclineChatJoinRequestParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// DeclineChatJoinRequest - Use this method to decline a chat join request. The bot must be an administrator
// in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (b *Bot) DeclineChatJoinRequest(params *DeclineChatJoinRequestParams) error {
	err := b.performRequest("declineChatJoinRequest", params)
	if err != nil {
		return fmt.Errorf("telego: declineChatJoinRequest(): %w", err)
	}

	return nil
}

// SetChatPhotoParams - Represents parameters of setChatPhoto method.
type SetChatPhotoParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// Photo - New chat photo, uploaded using multipart/form-data
	Photo InputFile `json:"photo"`
}

func (p *SetChatPhotoParams) fileParameters() map[string]ta.NamedReader {
	return map[string]ta.NamedReader{
		"photo": p.Photo.File,
	}
}

// SetChatPhoto - Use this method to set a new profile photo for the chat. Photos can't be changed for
// private chats. The bot must be an administrator in the chat for this to work and must have the appropriate
// administrator rights. Returns True on success.
func (b *Bot) SetChatPhoto(params *SetChatPhotoParams) error {
	err := b.performRequest("setChatPhoto", params)
	if err != nil {
		return fmt.Errorf("telego: setChatPhoto(): %w", err)
	}

	return nil
}

// DeleteChatPhotoParams - Represents parameters of deleteChatPhoto method.
type DeleteChatPhotoParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`
}

// DeleteChatPhoto - Use this method to delete a chat photo. Photos can't be changed for private chats. The
// bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
func (b *Bot) DeleteChatPhoto(params *DeleteChatPhotoParams) error {
	err := b.performRequest("deleteChatPhoto", params)
	if err != nil {
		return fmt.Errorf("telego: deleteChatPhoto(): %w", err)
	}

	return nil
}

// SetChatTitleParams - Represents parameters of setChatTitle method.
type SetChatTitleParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// Title - New chat title, 1-128 characters
	Title string `json:"title"`
}

// SetChatTitle - Use this method to change the title of a chat. Titles can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator
// rights. Returns True on success.
func (b *Bot) SetChatTitle(params *SetChatTitleParams) error {
	err := b.performRequest("setChatTitle", params)
	if err != nil {
		return fmt.Errorf("telego: setChatTitle(): %w", err)
	}

	return nil
}

// SetChatDescriptionParams - Represents parameters of setChatDescription method.
type SetChatDescriptionParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// Description - Optional. New chat description, 0-255 characters
	Description string `json:"description,omitempty"`
}

// SetChatDescription - Use this method to change the description of a group, a supergroup or a channel. The
// bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
func (b *Bot) SetChatDescription(params *SetChatDescriptionParams) error {
	err := b.performRequest("setChatDescription", params)
	if err != nil {
		return fmt.Errorf("telego: setChatDescription(): %w", err)
	}

	return nil
}

// PinChatMessageParams - Represents parameters of pinChatMessage method.
type PinChatMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageID - Identifier of a message to pin
	MessageID int `json:"message_id"`

	// DisableNotification - Optional. Pass True if it is not necessary to send a notification to all chat
	// members about the new pinned message. Notifications are always disabled in channels and private chats.
	DisableNotification bool `json:"disable_notification,omitempty"`
}

// PinChatMessage - Use this method to add a message to the list of pinned messages in a chat. If the chat is
// not a private chat, the bot must be an administrator in the chat for this to work and must have the
// 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a
// channel. Returns True on success.
func (b *Bot) PinChatMessage(params *PinChatMessageParams) error {
	err := b.performRequest("pinChatMessage", params)
	if err != nil {
		return fmt.Errorf("telego: pinChatMessage(): %w", err)
	}

	return nil
}

// UnpinChatMessageParams - Represents parameters of unpinChatMessage method.
type UnpinChatMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
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
	err := b.performRequest("unpinChatMessage", params)
	if err != nil {
		return fmt.Errorf("telego: unpinChatMessage(): %w", err)
	}

	return nil
}

// UnpinAllChatMessagesParams - Represents parameters of unpinAllChatMessages method.
type UnpinAllChatMessagesParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`
}

// UnpinAllChatMessages - Use this method to clear the list of pinned messages in a chat. If the chat is not
// a private chat, the bot must be an administrator in the chat for this to work and must have the
// 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a
// channel. Returns True on success.
func (b *Bot) UnpinAllChatMessages(params *UnpinAllChatMessagesParams) error {
	err := b.performRequest("unpinAllChatMessages", params)
	if err != nil {
		return fmt.Errorf("telego: unpinAllChatMessages(): %w", err)
	}

	return nil
}

// LeaveChatParams - Represents parameters of leaveChat method.
type LeaveChatParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channel_username)
	ChatID ChatID `json:"chat_id"`
}

// LeaveChat - Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
func (b *Bot) LeaveChat(params *LeaveChatParams) error {
	err := b.performRequest("leaveChat", params)
	if err != nil {
		return fmt.Errorf("telego: leaveChat(): %w", err)
	}

	return nil
}

// GetChatParams - Represents parameters of getChat method.
type GetChatParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channel_username)
	ChatID ChatID `json:"chat_id"`
}

// GetChat - Use this method to get up-to-date information about the chat. Returns a ChatFullInfo
// (https://core.telegram.org/bots/api#chatfullinfo) object on success.
func (b *Bot) GetChat(params *GetChatParams) (*ChatFullInfo, error) {
	var chatFullInfo *ChatFullInfo
	err := b.performRequest("getChat", params, &chatFullInfo)
	if err != nil {
		return nil, fmt.Errorf("telego: getChat(): %w", err)
	}

	return chatFullInfo, nil
}

// GetChatAdministratorsParams - Represents parameters of getChatAdministrators method.
type GetChatAdministratorsParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channel_username)
	ChatID ChatID `json:"chat_id"`
}

// GetChatAdministrators - Use this method to get a list of administrators in a chat, which aren't bots.
// Returns an Array of ChatMember (https://core.telegram.org/bots/api#chatmember) objects.
func (b *Bot) GetChatAdministrators(params *GetChatAdministratorsParams) ([]ChatMember, error) {
	var chatMembersData []chatMemberData
	err := b.performRequest("getChatAdministrators", params, &chatMembersData)
	if err != nil {
		return nil, fmt.Errorf("telego: getChatAdministrators(): %w", err)
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
	// format @channel_username)
	ChatID ChatID `json:"chat_id"`
}

// GetChatMemberCount - Use this method to get the number of members in a chat. Returns Int on success.
func (b *Bot) GetChatMemberCount(params *GetChatMemberCountParams) (*int, error) {
	var chatMemberCount *int
	err := b.performRequest("getChatMemberCount", params, &chatMemberCount)
	if err != nil {
		return nil, fmt.Errorf("telego: getChatMemberCount(): %w", err)
	}

	return chatMemberCount, nil
}

// GetChatMemberParams - Represents parameters of getChatMember method.
type GetChatMemberParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channel_username)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// GetChatMember - Use this method to get information about a member of a chat. The method is only guaranteed
// to work for other users if the bot is an administrator in the chat. Returns a ChatMember
// (https://core.telegram.org/bots/api#chatmember) object on success.
func (b *Bot) GetChatMember(params *GetChatMemberParams) (ChatMember, error) {
	var memberData chatMemberData
	err := b.performRequest("getChatMember", params, &memberData)
	if err != nil {
		return nil, fmt.Errorf("telego: getChatMember(): %w", err)
	}

	return memberData.Data, nil
}

// SetChatStickerSetParams - Represents parameters of setChatStickerSet method.
type SetChatStickerSetParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// StickerSetName - Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name"`
}

// SetChatStickerSet - Use this method to set a new group sticker set for a supergroup. The bot must be an
// administrator in the chat for this to work and must have the appropriate administrator rights. Use the field
// can_set_sticker_set optionally returned in getChat (https://core.telegram.org/bots/api#getchat) requests to
// check if the bot can use this method. Returns True on success.
func (b *Bot) SetChatStickerSet(params *SetChatStickerSetParams) error {
	err := b.performRequest("setChatStickerSet", params)
	if err != nil {
		return fmt.Errorf("telego: setChatStickerSet(): %w", err)
	}

	return nil
}

// DeleteChatStickerSetParams - Represents parameters of deleteChatStickerSet method.
type DeleteChatStickerSetParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`
}

// DeleteChatStickerSet - Use this method to delete a group sticker set from a supergroup. The bot must be an
// administrator in the chat for this to work and must have the appropriate administrator rights. Use the field
// can_set_sticker_set optionally returned in getChat (https://core.telegram.org/bots/api#getchat) requests to
// check if the bot can use this method. Returns True on success.
func (b *Bot) DeleteChatStickerSet(params *DeleteChatStickerSetParams) error {
	err := b.performRequest("deleteChatStickerSet", params)
	if err != nil {
		return fmt.Errorf("telego: deleteChatStickerSet(): %w", err)
	}

	return nil
}

// GetForumTopicIconStickers - Use this method to get custom emoji stickers, which can be used as a forum
// topic icon by any user. Requires no parameters. Returns an Array of Sticker
// (https://core.telegram.org/bots/api#sticker) objects.
func (b *Bot) GetForumTopicIconStickers() ([]Sticker, error) {
	var stickers []Sticker
	err := b.performRequest("getForumTopicIconStickers", nil, &stickers)
	if err != nil {
		return nil, fmt.Errorf("telego: getForumTopicIconStickers(): %w", err)
	}

	return stickers, nil
}

// CreateForumTopicParams - Represents parameters of createForumTopic method.
type CreateForumTopicParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// Name - Topic name, 1-128 characters
	Name string `json:"name"`

	// IconColor - Optional. Color of the topic icon in RGB format. Currently, must be one of 7322096
	// (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047
	// (0xFB6F5F)
	IconColor int `json:"icon_color,omitempty"`

	// IconCustomEmojiID - Optional. Unique identifier of the custom emoji shown as the topic icon. Use
	// getForumTopicIconStickers (https://core.telegram.org/bots/api#getforumtopiciconstickers) to get all allowed
	// custom emoji identifiers.
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// CreateForumTopic - Use this method to create a topic in a forum supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns
// information about the created topic as a ForumTopic (https://core.telegram.org/bots/api#forumtopic) object.
func (b *Bot) CreateForumTopic(params *CreateForumTopicParams) (*ForumTopic, error) {
	var forumTopic *ForumTopic
	err := b.performRequest("createForumTopic", params, &forumTopic)
	if err != nil {
		return nil, fmt.Errorf("telego: createForumTopic(): %w", err)
	}

	return forumTopic, nil
}

// EditForumTopicParams - Represents parameters of editForumTopic method.
type EditForumTopicParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Unique identifier for the target message thread of the forum topic
	MessageThreadID int `json:"message_thread_id"`

	// Name - Optional. New topic name, 0-128 characters. If not specified or empty, the current name of the
	// topic will be kept
	Name string `json:"name,omitempty"`

	// IconCustomEmojiID - Optional. New unique identifier of the custom emoji shown as the topic icon. Use
	// getForumTopicIconStickers (https://core.telegram.org/bots/api#getforumtopiciconstickers) to get all allowed
	// custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be
	// kept
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
}

// EditForumTopic - Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must
// be an administrator in the chat for this to work and must have can_manage_topics administrator rights, unless
// it is the creator of the topic. Returns True on success.
func (b *Bot) EditForumTopic(params *EditForumTopicParams) error {
	err := b.performRequest("editForumTopic", params)
	if err != nil {
		return fmt.Errorf("telego: editForumTopic(): %w", err)
	}

	return nil
}

// CloseForumTopicParams - Represents parameters of closeForumTopic method.
type CloseForumTopicParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Unique identifier for the target message thread of the forum topic
	MessageThreadID int `json:"message_thread_id"`
}

// CloseForumTopic - Use this method to close an open topic in a forum supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless
// it is the creator of the topic. Returns True on success.
func (b *Bot) CloseForumTopic(params *CloseForumTopicParams) error {
	err := b.performRequest("closeForumTopic", params)
	if err != nil {
		return fmt.Errorf("telego: closeForumTopic(): %w", err)
	}

	return nil
}

// ReopenForumTopicParams - Represents parameters of reopenForumTopic method.
type ReopenForumTopicParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Unique identifier for the target message thread of the forum topic
	MessageThreadID int `json:"message_thread_id"`
}

// ReopenForumTopic - Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless
// it is the creator of the topic. Returns True on success.
func (b *Bot) ReopenForumTopic(params *ReopenForumTopicParams) error {
	err := b.performRequest("reopenForumTopic", params)
	if err != nil {
		return fmt.Errorf("telego: reopenForumTopic(): %w", err)
	}

	return nil
}

// DeleteForumTopicParams - Represents parameters of deleteForumTopic method.
type DeleteForumTopicParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Unique identifier for the target message thread of the forum topic
	MessageThreadID int `json:"message_thread_id"`
}

// DeleteForumTopic - Use this method to delete a forum topic along with all its messages in a forum
// supergroup chat. The bot must be an administrator in the chat for this to work and must have the
// can_delete_messages administrator rights. Returns True on success.
func (b *Bot) DeleteForumTopic(params *DeleteForumTopicParams) error {
	err := b.performRequest("deleteForumTopic", params)
	if err != nil {
		return fmt.Errorf("telego: deleteForumTopic(): %w", err)
	}

	return nil
}

// UnpinAllForumTopicMessagesParams - Represents parameters of unpinAllForumTopicMessages method.
type UnpinAllForumTopicMessagesParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Unique identifier for the target message thread of the forum topic
	MessageThreadID int `json:"message_thread_id"`
}

// UnpinAllForumTopicMessages - Use this method to clear the list of pinned messages in a forum topic. The
// bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator
// right in the supergroup. Returns True on success.
func (b *Bot) UnpinAllForumTopicMessages(params *UnpinAllForumTopicMessagesParams) error {
	err := b.performRequest("unpinAllForumTopicMessages", params)
	if err != nil {
		return fmt.Errorf("telego: unpinAllForumTopicMessages(): %w", err)
	}

	return nil
}

// EditGeneralForumTopicParams - Represents parameters of editGeneralForumTopic method.
type EditGeneralForumTopicParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// Name - New topic name, 1-128 characters
	Name string `json:"name"`
}

// EditGeneralForumTopic - Use this method to edit the name of the 'General' topic in a forum supergroup
// chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics
// administrator rights. Returns True on success.
func (b *Bot) EditGeneralForumTopic(params *EditGeneralForumTopicParams) error {
	err := b.performRequest("editGeneralForumTopic", params)
	if err != nil {
		return fmt.Errorf("telego: editGeneralForumTopic(): %w", err)
	}

	return nil
}

// CloseGeneralForumTopicParams - Represents parameters of closeGeneralForumTopic method.
type CloseGeneralForumTopicParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`
}

// CloseGeneralForumTopic - Use this method to close an open 'General' topic in a forum supergroup chat. The
// bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator
// rights. Returns True on success.
func (b *Bot) CloseGeneralForumTopic(params *CloseGeneralForumTopicParams) error {
	err := b.performRequest("closeGeneralForumTopic", params)
	if err != nil {
		return fmt.Errorf("telego: closeGeneralForumTopic(): %w", err)
	}

	return nil
}

// ReopenGeneralForumTopicParams - Represents parameters of reopenGeneralForumTopic method.
type ReopenGeneralForumTopicParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`
}

// ReopenGeneralForumTopic - Use this method to reopen a closed 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics
// administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
func (b *Bot) ReopenGeneralForumTopic(params *ReopenGeneralForumTopicParams) error {
	err := b.performRequest("reopenGeneralForumTopic", params)
	if err != nil {
		return fmt.Errorf("telego: reopenGeneralForumTopic(): %w", err)
	}

	return nil
}

// HideGeneralForumTopicParams - Represents parameters of hideGeneralForumTopic method.
type HideGeneralForumTopicParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`
}

// HideGeneralForumTopic - Use this method to hide the 'General' topic in a forum supergroup chat. The bot
// must be an administrator in the chat for this to work and must have the can_manage_topics administrator
// rights. The topic will be automatically closed if it was open. Returns True on success.
func (b *Bot) HideGeneralForumTopic(params *HideGeneralForumTopicParams) error {
	err := b.performRequest("hideGeneralForumTopic", params)
	if err != nil {
		return fmt.Errorf("telego: hideGeneralForumTopic(): %w", err)
	}

	return nil
}

// UnhideGeneralForumTopicParams - Represents parameters of unhideGeneralForumTopic method.
type UnhideGeneralForumTopicParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`
}

// UnhideGeneralForumTopic - Use this method to unhide the 'General' topic in a forum supergroup chat. The
// bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator
// rights. Returns True on success.
func (b *Bot) UnhideGeneralForumTopic(params *UnhideGeneralForumTopicParams) error {
	err := b.performRequest("unhideGeneralForumTopic", params)
	if err != nil {
		return fmt.Errorf("telego: unhideGeneralForumTopic(): %w", err)
	}

	return nil
}

// UnpinAllGeneralForumTopicMessagesParams - Represents parameters of unpinAllGeneralForumTopicMessages
// method.
type UnpinAllGeneralForumTopicMessagesParams struct {
	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`
}

// UnpinAllGeneralForumTopicMessages - Use this method to clear the list of pinned messages in a General
// forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages
// administrator right in the supergroup. Returns True on success.
func (b *Bot) UnpinAllGeneralForumTopicMessages(params *UnpinAllGeneralForumTopicMessagesParams) error {
	err := b.performRequest("unpinAllGeneralForumTopicMessages", params)
	if err != nil {
		return fmt.Errorf("telego: unpinAllGeneralForumTopicMessages(): %w", err)
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
	// (https://t.me/botfather), specify the URL that opens your game - note that this will only work if the query
	// comes from a callback_game (https://core.telegram.org/bots/api#inlinekeyboardbutton) button.
	// Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	URL string `json:"url,omitempty"`

	// CacheTime - Optional. The maximum amount of time in seconds that the result of the callback query may be
	// cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime int `json:"cache_time,omitempty"`
}

// AnswerCallbackQuery - Use this method to send answers to callback queries sent from inline keyboards
// (https://core.telegram.org/bots/features#inline-keyboards). The answer will be displayed to the user as a
// notification at the top of the chat screen or as an alert. On success, True is returned.
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first
// create a game for your bot via @BotFather (https://t.me/botfather) and accept the terms. Otherwise, you may
// use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
func (b *Bot) AnswerCallbackQuery(params *AnswerCallbackQueryParams) error {
	err := b.performRequest("answerCallbackQuery", params)
	if err != nil {
		return fmt.Errorf("telego: answerCallbackQuery(): %w", err)
	}

	return nil
}

// GetUserChatBoostsParams - Represents parameters of getUserChatBoosts method.
type GetUserChatBoostsParams struct {
	// ChatID - Unique identifier for the chat or username of the channel (in the format @channel_username)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// GetUserChatBoosts - Use this method to get the list of boosts added to a chat by a user. Requires
// administrator rights in the chat. Returns a UserChatBoosts
// (https://core.telegram.org/bots/api#userchatboosts) object.
func (b *Bot) GetUserChatBoosts(params *GetUserChatBoostsParams) (*UserChatBoosts, error) {
	var userChatBoosts *UserChatBoosts
	err := b.performRequest("getUserChatBoosts", params, &userChatBoosts)
	if err != nil {
		return nil, fmt.Errorf("telego: getUserChatBoosts(): %w", err)
	}

	return userChatBoosts, nil
}

// GetBusinessConnectionParams - Represents parameters of getBusinessConnection method.
type GetBusinessConnectionParams struct {
	// BusinessConnectionID - Unique identifier of the business connection
	BusinessConnectionID string `json:"business_connection_id"`
}

// GetBusinessConnection - Use this method to get information about the connection of the bot with a business
// account. Returns a BusinessConnection (https://core.telegram.org/bots/api#businessconnection) object on
// success.
func (b *Bot) GetBusinessConnection(params *GetBusinessConnectionParams) (*BusinessConnection, error) {
	var businessConnection *BusinessConnection
	err := b.performRequest("getBusinessConnection", params, &businessConnection)
	if err != nil {
		return nil, fmt.Errorf("telego: getBusinessConnection(): %w", err)
	}

	return businessConnection, nil
}

// SetMyCommandsParams - Represents parameters of setMyCommands method.
type SetMyCommandsParams struct {
	// Commands - A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most
	// 100 commands can be specified.
	Commands []BotCommand `json:"commands"`

	// Scope - Optional. A JSON-serialized object, describing scope of users for which the commands are
	// relevant. Defaults to BotCommandScopeDefault (https://core.telegram.org/bots/api#botcommandscopedefault).
	Scope BotCommandScope `json:"scope,omitempty"`

	// LanguageCode - Optional. A two-letter ISO 639-1 language code. If empty, commands will be applied to all
	// users from the given scope, for whose language there are no dedicated commands
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyCommands - Use this method to change the list of the bot's commands. See this manual
// (https://core.telegram.org/bots/features#commands) for more details about bot commands. Returns True on
// success.
func (b *Bot) SetMyCommands(params *SetMyCommandsParams) error {
	err := b.performRequest("setMyCommands", params)
	if err != nil {
		return fmt.Errorf("telego: setMyCommands(): %w", err)
	}

	return nil
}

// DeleteMyCommandsParams - Represents parameters of deleteMyCommands method.
type DeleteMyCommandsParams struct {
	// Scope - Optional. A JSON-serialized object, describing scope of users for which the commands are
	// relevant. Defaults to BotCommandScopeDefault (https://core.telegram.org/bots/api#botcommandscopedefault).
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
	err := b.performRequest("deleteMyCommands", params)
	if err != nil {
		return fmt.Errorf("telego: deleteMyCommands(): %w", err)
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
// language. Returns an Array of BotCommand (https://core.telegram.org/bots/api#botcommand) objects. If commands
// aren't set, an empty list is returned.
func (b *Bot) GetMyCommands(params *GetMyCommandsParams) ([]BotCommand, error) {
	var botCommands []BotCommand
	err := b.performRequest("getMyCommands", params, &botCommands)
	if err != nil {
		return nil, fmt.Errorf("telego: getMyCommands(): %w", err)
	}

	return botCommands, nil
}

// SetMyNameParams - Represents parameters of setMyName method.
type SetMyNameParams struct {
	// Name - Optional. New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the
	// given language.
	Name string `json:"name,omitempty"`

	// LanguageCode - Optional. A two-letter ISO 639-1 language code. If empty, the name will be shown to all
	// users for whose language there is no dedicated name.
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyName - Use this method to change the bot's name. Returns True on success.
func (b *Bot) SetMyName(params *SetMyNameParams) error {
	err := b.performRequest("setMyName", params)
	if err != nil {
		return fmt.Errorf("telego: setMyName(): %w", err)
	}

	return nil
}

// GetMyNameParams - Represents parameters of getMyName method.
type GetMyNameParams struct {
	// LanguageCode - Optional. A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyName - Use this method to get the current bot name for the given user language. Returns BotName
// (https://core.telegram.org/bots/api#botname) on success.
func (b *Bot) GetMyName(params *GetMyNameParams) (*BotName, error) {
	var botName *BotName
	err := b.performRequest("getMyName", params, &botName)
	if err != nil {
		return nil, fmt.Errorf("telego: getMyName(): %w", err)
	}

	return botName, nil
}

// SetMyDescriptionParams - Represents parameters of setMyDescription method.
type SetMyDescriptionParams struct {
	// Description - Optional. New bot description; 0-512 characters. Pass an empty string to remove the
	// dedicated description for the given language.
	Description string `json:"description,omitempty"`

	// LanguageCode - Optional. A two-letter ISO 639-1 language code. If empty, the description will be applied
	// to all users for whose language there is no dedicated description.
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyDescription - Use this method to change the bot's description, which is shown in the chat with the
// bot if the chat is empty. Returns True on success.
func (b *Bot) SetMyDescription(params *SetMyDescriptionParams) error {
	err := b.performRequest("setMyDescription", params)
	if err != nil {
		return fmt.Errorf("telego: setMyDescription(): %w", err)
	}

	return nil
}

// GetMyDescriptionParams - Represents parameters of getMyDescription method.
type GetMyDescriptionParams struct {
	// LanguageCode - Optional. A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyDescription - Use this method to get the current bot description for the given user language. Returns
// BotDescription (https://core.telegram.org/bots/api#botdescription) on success.
func (b *Bot) GetMyDescription(params *GetMyDescriptionParams) (*BotDescription, error) {
	var botDescription *BotDescription
	err := b.performRequest("getMyDescription", params, &botDescription)
	if err != nil {
		return nil, fmt.Errorf("telego: getMyDescription(): %w", err)
	}

	return botDescription, nil
}

// SetMyShortDescriptionParams - Represents parameters of setMyShortDescription method.
type SetMyShortDescriptionParams struct {
	// ShortDescription - Optional. New short description for the bot; 0-120 characters. Pass an empty string to
	// remove the dedicated short description for the given language.
	ShortDescription string `json:"short_description,omitempty"`

	// LanguageCode - Optional. A two-letter ISO 639-1 language code. If empty, the short description will be
	// applied to all users for whose language there is no dedicated short description.
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyShortDescription - Use this method to change the bot's short description, which is shown on the bot's
// profile page and is sent together with the link when users share the bot. Returns True on success.
func (b *Bot) SetMyShortDescription(params *SetMyShortDescriptionParams) error {
	err := b.performRequest("setMyShortDescription", params)
	if err != nil {
		return fmt.Errorf("telego: setMyShortDescription(): %w", err)
	}

	return nil
}

// GetMyShortDescriptionParams - Represents parameters of getMyShortDescription method.
type GetMyShortDescriptionParams struct {
	// LanguageCode - Optional. A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyShortDescription - Use this method to get the current bot short description for the given user
// language. Returns BotShortDescription (https://core.telegram.org/bots/api#botshortdescription) on success.
func (b *Bot) GetMyShortDescription(params *GetMyShortDescriptionParams) (*BotShortDescription, error) {
	var botShortDescription *BotShortDescription
	err := b.performRequest("getMyShortDescription", params, &botShortDescription)
	if err != nil {
		return nil, fmt.Errorf("telego: getMyShortDescription(): %w", err)
	}

	return botShortDescription, nil
}

// SetChatMenuButtonParams - Represents parameters of setChatMenuButton method.
type SetChatMenuButtonParams struct {
	// ChatID - Optional. Unique identifier for the target private chat. If not specified, default bot's menu
	// button will be changed
	ChatID int64 `json:"chat_id,omitempty"`

	// MenuButton - Optional. A JSON-serialized object for the bot's new menu button. Defaults to
	// MenuButtonDefault (https://core.telegram.org/bots/api#menubuttondefault)
	MenuButton MenuButton `json:"menu_button,omitempty"`
}

// SetChatMenuButton - Use this method to change the bot's menu button in a private chat, or the default menu
// button. Returns True on success.
func (b *Bot) SetChatMenuButton(params *SetChatMenuButtonParams) error {
	err := b.performRequest("setChatMenuButton", params)
	if err != nil {
		return fmt.Errorf("telego: setChatMenuButton(): %w", err)
	}

	return nil
}

// GetChatMenuButtonParams - Represents parameters of getChatMenuButton method.
type GetChatMenuButtonParams struct {
	// ChatID - Optional. Unique identifier for the target private chat. If not specified, default bot's menu
	// button will be returned
	ChatID int64 `json:"chat_id,omitempty"`
}

// GetChatMenuButton - Use this method to get the current value of the bot's menu button in a private chat,
// or the default menu button. Returns MenuButton (https://core.telegram.org/bots/api#menubutton) on success.
func (b *Bot) GetChatMenuButton(params *GetChatMenuButtonParams) (MenuButton, error) {
	var menuButton menuButtonData
	err := b.performRequest("getChatMenuButton", params, &menuButton)
	if err != nil {
		return nil, fmt.Errorf("telego: getChatMenuButton(): %w", err)
	}

	return menuButton.Data, nil
}

// SetMyDefaultAdministratorRightsParams - Represents parameters of setMyDefaultAdministratorRights method.
type SetMyDefaultAdministratorRightsParams struct {
	// Rights - Optional. A JSON-serialized object describing new default administrator rights. If not
	// specified, the default administrator rights will be cleared.
	Rights *ChatAdministratorRights `json:"rights,omitempty"`

	// ForChannels - Optional. Pass True to change the default administrator rights of the bot in channels.
	// Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
	ForChannels bool `json:"for_channels,omitempty"`
}

// SetMyDefaultAdministratorRights - Use this method to change the default administrator rights requested by
// the bot when it's added as an administrator to groups or channels. These rights will be suggested to users,
// but they are free to modify the list before adding the bot. Returns True on success.
func (b *Bot) SetMyDefaultAdministratorRights(params *SetMyDefaultAdministratorRightsParams) error {
	err := b.performRequest("setMyDefaultAdministratorRights", params)
	if err != nil {
		return fmt.Errorf("telego: setMyDefaultAdministratorRights(): %w", err)
	}

	return nil
}

// GetMyDefaultAdministratorRightsParams - Represents parameters of getMyDefaultAdministratorRights method.
type GetMyDefaultAdministratorRightsParams struct {
	// ForChannels - Optional. Pass True to get default administrator rights of the bot in channels. Otherwise,
	// default administrator rights of the bot for groups and supergroups will be returned.
	ForChannels bool `json:"for_channels,omitempty"`
}

// GetMyDefaultAdministratorRights - Use this method to get the current default administrator rights of the
// bot. Returns ChatAdministratorRights (https://core.telegram.org/bots/api#chatadministratorrights) on success.
func (b *Bot) GetMyDefaultAdministratorRights(
	params *GetMyDefaultAdministratorRightsParams,
) (*ChatAdministratorRights, error) {
	var chatAdministratorRights *ChatAdministratorRights
	err := b.performRequest("getMyDefaultAdministratorRights", params, &chatAdministratorRights)
	if err != nil {
		return nil, fmt.Errorf("telego: getMyDefaultAdministratorRights(): %w", err)
	}

	return chatAdministratorRights, nil
}

// EditMessageTextParams - Represents parameters of editMessageText method.
type EditMessageTextParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channel_username)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the
	// inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Text - New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// ParseMode - Optional. Mode for parsing entities in the message text. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Entities - Optional. A JSON-serialized list of special entities that appear in message text, which can be
	// specified instead of parse_mode
	Entities []MessageEntity `json:"entities,omitempty"`

	// LinkPreviewOptions - Optional. Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageText - Use this method to edit text and game (https://core.telegram.org/bots/api#games)
// messages. On success, if the edited message is not an inline message, the edited Message
// (https://core.telegram.org/bots/api#message) is returned, otherwise True is returned.
func (b *Bot) EditMessageText(params *EditMessageTextParams) (*Message, error) {
	var message *Message
	var success *bool
	err := b.performRequest("editMessageText", params, &message, &success)
	if err != nil {
		return nil, fmt.Errorf("telego: editMessageText(): %w", err)
	}

	return message, nil
}

// EditMessageCaptionParams - Represents parameters of editMessageCaption method.
type EditMessageCaptionParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channel_username)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the
	// inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Caption - Optional. New caption of the message, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the message caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. A JSON-serialized list of special entities that appear in the caption, which
	// can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media.
	// Supported only for animation, photo and video messages.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageCaption - Use this method to edit captions of messages. On success, if the edited message is
// not an inline message, the edited Message (https://core.telegram.org/bots/api#message) is returned, otherwise
// True is returned.
func (b *Bot) EditMessageCaption(params *EditMessageCaptionParams) (*Message, error) {
	var message *Message
	var success *bool
	err := b.performRequest("editMessageCaption", params, &message, &success)
	if err != nil {
		return nil, fmt.Errorf("telego: editMessageCaption(): %w", err)
	}

	return message, nil
}

// EditMessageMediaParams - Represents parameters of editMessageMedia method.
type EditMessageMediaParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channel_username)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the
	// inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Media - A JSON-serialized object for a new media content of the message
	Media InputMedia `json:"media"`

	// ReplyMarkup - Optional. A JSON-serialized object for a new inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *EditMessageMediaParams) fileParameters() map[string]ta.NamedReader {
	fp := make(map[string]ta.NamedReader)

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
	var success *bool
	err := b.performRequest("editMessageMedia", params, &message, &success)
	if err != nil {
		return nil, fmt.Errorf("telego: editMessageMedia(): %w", err)
	}

	return message, nil
}

// EditMessageLiveLocationParams - Represents parameters of editMessageLiveLocation method.
type EditMessageLiveLocationParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channel_username)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the
	// inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Latitude - Latitude of new location
	Latitude float64 `json:"latitude"`

	// Longitude - Longitude of new location
	Longitude float64 `json:"longitude"`

	// LivePeriod - Optional. New period in seconds during which the location can be updated, starting from the
	// message send date. If 0x7FFFFFFF is specified, then the location can be updated forever. Otherwise, the new
	// value must not exceed the current live_period by more than a day, and the live location expiration date must
	// remain within the next 90 days. If not specified, then live_period remains unchanged
	LivePeriod int `json:"live_period,omitempty"`

	// HorizontalAccuracy - Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// Heading - Optional. Direction in which the user is moving, in degrees. Must be between 1 and 360 if
	// specified.
	Heading int `json:"heading,omitempty"`

	// ProximityAlertRadius - Optional. The maximum distance for proximity alerts about approaching another chat
	// member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for a new inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageLiveLocation - Use this method to edit live location messages. A location can be edited until
// its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation
// (https://core.telegram.org/bots/api#stopmessagelivelocation). On success, if the edited message is not an
// inline message, the edited Message (https://core.telegram.org/bots/api#message) is returned, otherwise True
// is returned.
func (b *Bot) EditMessageLiveLocation(params *EditMessageLiveLocationParams) (*Message, error) {
	var message *Message
	var success *bool
	err := b.performRequest("editMessageLiveLocation", params, &message, &success)
	if err != nil {
		return nil, fmt.Errorf("telego: editMessageLiveLocation(): %w", err)
	}

	return message, nil
}

// StopMessageLiveLocationParams - Represents parameters of stopMessageLiveLocation method.
type StopMessageLiveLocationParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channel_username)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message with live
	// location to stop
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the
	// inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for a new inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// StopMessageLiveLocation - Use this method to stop updating a live location message before live_period
// expires. On success, if the message is not an inline message, the edited Message
// (https://core.telegram.org/bots/api#message) is returned, otherwise True is returned.
func (b *Bot) StopMessageLiveLocation(params *StopMessageLiveLocationParams) (*Message, error) {
	var message *Message
	var success *bool
	err := b.performRequest("stopMessageLiveLocation", params, &message, &success)
	if err != nil {
		return nil, fmt.Errorf("telego: stopMessageLiveLocation(): %w", err)
	}

	return message, nil
}

// EditMessageReplyMarkupParams - Represents parameters of editMessageReplyMarkup method.
type EditMessageReplyMarkupParams struct {
	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	// or username of the target channel (in the format @channel_username)
	ChatID ChatID `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the
	// inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkup - Use this method to edit only the reply markup of messages. On success, if the
// edited message is not an inline message, the edited Message (https://core.telegram.org/bots/api#message) is
// returned, otherwise True is returned.
func (b *Bot) EditMessageReplyMarkup(params *EditMessageReplyMarkupParams) (*Message, error) {
	var message *Message
	var success *bool
	err := b.performRequest("editMessageReplyMarkup", params, &message, &success)
	if err != nil {
		return nil, fmt.Errorf("telego: editMessageReplyMarkup(): %w", err)
	}

	return message, nil
}

// StopPollParams - Represents parameters of stopPoll method.
type StopPollParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageID - Identifier of the original message with the poll
	MessageID int `json:"message_id"`

	// ReplyMarkup - Optional. A JSON-serialized object for a new message inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards).
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// StopPoll - Use this method to stop a poll which was sent by the bot. On success, the stopped Poll
// (https://core.telegram.org/bots/api#poll) is returned.
func (b *Bot) StopPoll(params *StopPollParams) (*Poll, error) {
	var poll *Poll
	err := b.performRequest("stopPoll", params, &poll)
	if err != nil {
		return nil, fmt.Errorf("telego: stopPoll(): %w", err)
	}

	return poll, nil
}

// DeleteMessageParams - Represents parameters of deleteMessage method.
type DeleteMessageParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageID - Identifier of the message to delete
	MessageID int `json:"message_id"`
}

// DeleteMessage - Use this method to delete a message, including service messages, with the following
// limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - Service messages about a supergroup, channel, or forum topic creation can't be deleted.
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message
// there.
// Returns True on success.
func (b *Bot) DeleteMessage(params *DeleteMessageParams) error {
	err := b.performRequest("deleteMessage", params)
	if err != nil {
		return fmt.Errorf("telego: deleteMessage(): %w", err)
	}

	return nil
}

// DeleteMessagesParams - Represents parameters of deleteMessages method.
type DeleteMessagesParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageIDs - A JSON-serialized list of 1-100 identifiers of messages to delete. See deleteMessage
	// (https://core.telegram.org/bots/api#deletemessage) for limitations on which messages can be deleted
	MessageIDs []int `json:"message_ids"`
}

// DeleteMessages - Use this method to delete multiple messages simultaneously. If some of the specified
// messages can't be found, they are skipped. Returns True on success.
func (b *Bot) DeleteMessages(params *DeleteMessagesParams) error {
	err := b.performRequest("deleteMessages", params)
	if err != nil {
		return fmt.Errorf("telego: deleteMessages(): %w", err)
	}

	return nil
}

// SendStickerParams - Represents parameters of sendSticker method.
type SendStickerParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Sticker - Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload
	// a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data. More information on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files). Video and animated stickers can't be sent via an HTTP
	// URL.
	Sticker InputFile `json:"sticker"`

	// Emoji - Optional. Emoji associated with the sticker; only for just uploaded stickers
	Emoji string `json:"emoji,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. Additional interface options. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards), custom reply keyboard
	// (https://core.telegram.org/bots/features#keyboards), instructions to remove a reply keyboard or to force a
	// reply from the user
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p *SendStickerParams) fileParameters() map[string]ta.NamedReader {
	return map[string]ta.NamedReader{
		"sticker": p.Sticker.File,
	}
}

// SendSticker - Use this method to send static .WEBP, animated (https://telegram.org/blog/animated-stickers)
// .TGS, or video (https://telegram.org/blog/video-stickers-better-reactions) .WEBM stickers. On success, the
// sent Message (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendSticker(params *SendStickerParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendSticker", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendSticker(): %w", err)
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
		return nil, fmt.Errorf("telego: getStickerSet(): %w", err)
	}

	return stickerSet, nil
}

// GetCustomEmojiStickersParams - Represents parameters of getCustomEmojiStickers method.
type GetCustomEmojiStickersParams struct {
	// CustomEmojiIDs - A JSON-serialized list of custom emoji identifiers. At most 200 custom emoji identifiers
	// can be specified.
	CustomEmojiIDs []string `json:"custom_emoji_ids"`
}

// GetCustomEmojiStickers - Use this method to get information about custom emoji stickers by their
// identifiers. Returns an Array of Sticker (https://core.telegram.org/bots/api#sticker) objects.
func (b *Bot) GetCustomEmojiStickers(params *GetCustomEmojiStickersParams) ([]Sticker, error) {
	var stickers []Sticker
	err := b.performRequest("getCustomEmojiStickers", params, &stickers)
	if err != nil {
		return nil, fmt.Errorf("telego: getCustomEmojiStickers(): %w", err)
	}

	return stickers, nil
}

// UploadStickerFileParams - Represents parameters of uploadStickerFile method.
type UploadStickerFileParams struct {
	// UserID - User identifier of sticker file owner
	UserID int64 `json:"user_id"`

	// Sticker - A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format. See
	// https://core.telegram.org/stickers (https://core.telegram.org/stickers) for technical requirements. More
	// information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Sticker InputFile `json:"sticker"`

	// StickerFormat - Format of the sticker, must be one of “static”, “animated”, “video”
	StickerFormat string `json:"sticker_format"`
}

// Sticker formats
const (
	StickerFormatStatic   = "static"
	StickerFormatAnimated = "animated"
	StickerFormatVideo    = "video"
)

func (p *UploadStickerFileParams) fileParameters() map[string]ta.NamedReader {
	return map[string]ta.NamedReader{
		"sticker": p.Sticker.File,
	}
}

// UploadStickerFile - Use this method to upload a file with a sticker for later use in the
// createNewStickerSet (https://core.telegram.org/bots/api#createnewstickerset), addStickerToSet
// (https://core.telegram.org/bots/api#addstickertoset), or replaceStickerInSet
// (https://core.telegram.org/bots/api#replacestickerinset) methods (the file can be used multiple times).
// Returns the uploaded File (https://core.telegram.org/bots/api#file) on success.
func (b *Bot) UploadStickerFile(params *UploadStickerFileParams) (*File, error) {
	var file *File
	err := b.performRequest("uploadStickerFile", params, &file)
	if err != nil {
		return nil, fmt.Errorf("telego: uploadStickerFile(): %w", err)
	}

	return file, nil
}

// CreateNewStickerSetParams - Represents parameters of createNewStickerSet method.
type CreateNewStickerSetParams struct {
	// UserID - User identifier of created sticker set owner
	UserID int64 `json:"user_id"`

	// Name - Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only
	// English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and
	// must end in "_by_<bot_username>". <bot_username> is case insensitive. 1-64 characters.
	Name string `json:"name"`

	// Title - Sticker set title, 1-64 characters
	Title string `json:"title"`

	// Stickers - A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
	Stickers []InputSticker `json:"stickers"`

	// StickerType - Optional. Type of stickers in the set, pass “regular”, “mask”, or
	// “custom_emoji”. By default, a regular sticker set is created.
	StickerType string `json:"sticker_type,omitempty"`

	// NeedsRepainting - Optional. Pass True if stickers in the sticker set must be repainted to the color of
	// text when used in messages, the accent color if used as emoji status, white on chat photos, or another
	// appropriate color based on context; for custom emoji sticker sets only
	NeedsRepainting bool `json:"needs_repainting,omitempty"`
}

func (p *CreateNewStickerSetParams) fileParameters() map[string]ta.NamedReader {
	fp := make(map[string]ta.NamedReader)

	for i, s := range p.Stickers {
		p.Stickers[i].Sticker.needAttach = true

		file := s.Sticker.File
		if isNil(file) {
			continue
		}

		fp[file.Name()] = file
	}

	return fp
}

// CreateNewStickerSet - Use this method to create a new sticker set owned by a user. The bot will be able to
// edit the sticker set thus created. Returns True on success.
func (b *Bot) CreateNewStickerSet(params *CreateNewStickerSetParams) error {
	err := b.performRequest("createNewStickerSet", params)
	if err != nil {
		return fmt.Errorf("telego: createNewStickerSet(): %w", err)
	}

	return nil
}

// AddStickerToSetParams - Represents parameters of addStickerToSet method.
type AddStickerToSetParams struct {
	// UserID - User identifier of sticker set owner
	UserID int64 `json:"user_id"`

	// Name - Sticker set name
	Name string `json:"name"`

	// Sticker - A JSON-serialized object with information about the added sticker. If exactly the same sticker
	// had already been added to the set, then the set isn't changed.
	Sticker InputSticker `json:"sticker"`
}

func (p *AddStickerToSetParams) fileParameters() map[string]ta.NamedReader {
	file := p.Sticker.Sticker.File
	if isNil(file) {
		return map[string]ta.NamedReader{}
	}

	p.Sticker.Sticker.needAttach = true
	return map[string]ta.NamedReader{
		file.Name(): file,
	}
}

// AddStickerToSet - Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can
// have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.
func (b *Bot) AddStickerToSet(params *AddStickerToSetParams) error {
	err := b.performRequest("addStickerToSet", params)
	if err != nil {
		return fmt.Errorf("telego: addStickerToSet(): %w", err)
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
	err := b.performRequest("setStickerPositionInSet", params)
	if err != nil {
		return fmt.Errorf("telego: setStickerPositionInSet(): %w", err)
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
	err := b.performRequest("deleteStickerFromSet", params)
	if err != nil {
		return fmt.Errorf("telego: deleteStickerFromSet(): %w", err)
	}

	return nil
}

// ReplaceStickerInSetParams - Represents parameters of replaceStickerInSet method.
type ReplaceStickerInSetParams struct {
	// UserID - User identifier of the sticker set owner
	UserID int64 `json:"user_id"`

	// Name - Sticker set name
	Name string `json:"name"`

	// OldSticker - File identifier of the replaced sticker
	OldSticker string `json:"old_sticker"`

	// Sticker - A JSON-serialized object with information about the added sticker. If exactly the same sticker
	// had already been added to the set, then the set remains unchanged.
	Sticker InputSticker `json:"sticker"`
}

// ReplaceStickerInSet - Use this method to replace an existing sticker in a sticker set with a new one. The
// method is equivalent to calling deleteStickerFromSet
// (https://core.telegram.org/bots/api#deletestickerfromset), then addStickerToSet
// (https://core.telegram.org/bots/api#addstickertoset), then setStickerPositionInSet
// (https://core.telegram.org/bots/api#setstickerpositioninset). Returns True on success.
func (b *Bot) ReplaceStickerInSet(params *ReplaceStickerInSetParams) error {
	err := b.performRequest("replaceStickerInSet", params)
	if err != nil {
		return fmt.Errorf("telego: replaceStickerInSet(): %w", err)
	}

	return nil
}

// SetStickerEmojiListParams - Represents parameters of setStickerEmojiList method.
type SetStickerEmojiListParams struct {
	// Sticker - File identifier of the sticker
	Sticker string `json:"sticker"`

	// EmojiList - A JSON-serialized list of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list"`
}

// SetStickerEmojiList - Use this method to change the list of emoji assigned to a regular or custom emoji
// sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
func (b *Bot) SetStickerEmojiList(params *SetStickerEmojiListParams) error {
	err := b.performRequest("setStickerEmojiList", params)
	if err != nil {
		return fmt.Errorf("telego: setStickerEmojiList(): %w", err)
	}

	return nil
}

// SetStickerKeywordsParams - Represents parameters of setStickerKeywords method.
type SetStickerKeywordsParams struct {
	// Sticker - File identifier of the sticker
	Sticker string `json:"sticker"`

	// Keywords - Optional. A JSON-serialized list of 0-20 search keywords for the sticker with total length of
	// up to 64 characters
	Keywords []string `json:"keywords,omitempty"`
}

// SetStickerKeywords - Use this method to change search keywords assigned to a regular or custom emoji
// sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
func (b *Bot) SetStickerKeywords(params *SetStickerKeywordsParams) error {
	err := b.performRequest("setStickerKeywords", params)
	if err != nil {
		return fmt.Errorf("telego: setStickerKeywords(): %w", err)
	}

	return nil
}

// SetStickerMaskPositionParams - Represents parameters of setStickerMaskPosition method.
type SetStickerMaskPositionParams struct {
	// Sticker - File identifier of the sticker
	Sticker string `json:"sticker"`

	// MaskPosition - Optional. A JSON-serialized object with the position where the mask should be placed on
	// faces. Omit the parameter to remove the mask position.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// SetStickerMaskPosition - Use this method to change the mask position
// (https://core.telegram.org/bots/api#maskposition) of a mask sticker. The sticker must belong to a sticker set
// that was created by the bot. Returns True on success.
func (b *Bot) SetStickerMaskPosition(params *SetStickerMaskPositionParams) error {
	err := b.performRequest("setStickerMaskPosition", params)
	if err != nil {
		return fmt.Errorf("telego: setStickerMaskPosition(): %w", err)
	}

	return nil
}

// SetStickerSetTitleParams - Represents parameters of setStickerSetTitle method.
type SetStickerSetTitleParams struct {
	// Name - Sticker set name
	Name string `json:"name"`

	// Title - Sticker set title, 1-64 characters
	Title string `json:"title"`
}

// SetStickerSetTitle - Use this method to set the title of a created sticker set. Returns True on success.
func (b *Bot) SetStickerSetTitle(params *SetStickerSetTitleParams) error {
	err := b.performRequest("setStickerSetTitle", params)
	if err != nil {
		return fmt.Errorf("telego: setStickerSetTitle(): %w", err)
	}

	return nil
}

// SetStickerSetThumbnailParams - Represents parameters of setStickerSetThumbnail method.
type SetStickerSetThumbnailParams struct {
	// Name - Sticker set name
	Name string `json:"name"`

	// UserID - User identifier of the sticker set owner
	UserID int64 `json:"user_id"`

	// Thumbnail - Optional. A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and
	// have a width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size
	// (see https://core.telegram.org/stickers#animated-sticker-requirements
	// (https://core.telegram.org/stickers#animated-sticker-requirements) for animated sticker technical
	// requirements), or a WEBM video with the thumbnail up to 32 kilobytes in size; see
	// https://core.telegram.org/stickers#video-sticker-requirements
	// (https://core.telegram.org/stickers#video-sticker-requirements) for video sticker technical requirements.
	// Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a
	// String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More
	// information on Sending Files » (https://core.telegram.org/bots/api#sending-files). Animated and video
	// sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is dropped and the
	// first sticker is used as the thumbnail.
	Thumbnail *InputFile `json:"thumbnail,omitempty"`

	// Format - Format of the thumbnail, must be one of “static” for a .WEBP or .PNG image, “animated”
	// for a .TGS animation, or “video” for a WEBM video
	Format string `json:"format"`
}

func (p *SetStickerSetThumbnailParams) fileParameters() map[string]ta.NamedReader {
	fp := make(map[string]ta.NamedReader)

	if p.Thumbnail != nil {
		fp["thumbnail"] = p.Thumbnail.File
	}

	return fp
}

// SetStickerSetThumbnail - Use this method to set the thumbnail of a regular or mask sticker set. The format
// of the thumbnail file must match the format of the stickers in the set. Returns True on success.
func (b *Bot) SetStickerSetThumbnail(params *SetStickerSetThumbnailParams) error {
	err := b.performRequest("setStickerSetThumbnail", params)
	if err != nil {
		return fmt.Errorf("telego: setStickerSetThumbnail(): %w", err)
	}

	return nil
}

// SetCustomEmojiStickerSetThumbnailParams - Represents parameters of setCustomEmojiStickerSetThumbnail
// method.
type SetCustomEmojiStickerSetThumbnailParams struct {
	// Name - Sticker set name
	Name string `json:"name"`

	// CustomEmojiID - Optional. Custom emoji identifier of a sticker from the sticker set; pass an empty string
	// to drop the thumbnail and use the first sticker as the thumbnail.
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// SetCustomEmojiStickerSetThumbnail - Use this method to set the thumbnail of a custom emoji sticker set.
// Returns True on success.
func (b *Bot) SetCustomEmojiStickerSetThumbnail(params *SetCustomEmojiStickerSetThumbnailParams) error {
	err := b.performRequest("setCustomEmojiStickerSetThumbnail", params)
	if err != nil {
		return fmt.Errorf("telego: setCustomEmojiStickerSetThumbnail(): %w", err)
	}

	return nil
}

// DeleteStickerSetParams - Represents parameters of deleteStickerSet method.
type DeleteStickerSetParams struct {
	// Name - Sticker set name
	Name string `json:"name"`
}

// DeleteStickerSet - Use this method to delete a sticker set that was created by the bot. Returns True on
// success.
func (b *Bot) DeleteStickerSet(params *DeleteStickerSetParams) error {
	err := b.performRequest("deleteStickerSet", params)
	if err != nil {
		return fmt.Errorf("telego: deleteStickerSet(): %w", err)
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

	// IsPersonal - Optional. Pass True if results may be cached on the server side only for the user that sent
	// the query. By default, results may be returned to any user who sends the same query.
	IsPersonal bool `json:"is_personal,omitempty"`

	// NextOffset - Optional. Pass the offset that a client should send in the next query with the same text to
	// receive more results. Pass an empty string if there are no more results or if you don't support pagination.
	// Offset length can't exceed 64 bytes.
	NextOffset string `json:"next_offset,omitempty"`

	// Button - Optional. A JSON-serialized object describing a button to be shown above inline query results
	Button *InlineQueryResultsButton `json:"button,omitempty"`
}

// AnswerInlineQuery - Use this method to send answers to an inline query. On success, True is returned.
// No more than 50 results per query are allowed.
func (b *Bot) AnswerInlineQuery(params *AnswerInlineQueryParams) error {
	err := b.performRequest("answerInlineQuery", params)
	if err != nil {
		return fmt.Errorf("telego: answerInlineQuery(): %w", err)
	}

	return nil
}

// AnswerWebAppQueryParams - Represents parameters of answerWebAppQuery method.
type AnswerWebAppQueryParams struct {
	// WebAppQueryID - Unique identifier for the query to be answered
	WebAppQueryID string `json:"web_app_query_id"`

	// Result - A JSON-serialized object describing the message to be sent
	Result InlineQueryResult `json:"result"`
}

// AnswerWebAppQuery - Use this method to set the result of an interaction with a Web App
// (https://core.telegram.org/bots/webapps) and send a corresponding message on behalf of the user to the chat
// from which the query originated. On success, a SentWebAppMessage
// (https://core.telegram.org/bots/api#sentwebappmessage) object is returned.
func (b *Bot) AnswerWebAppQuery(params *AnswerWebAppQueryParams) (*SentWebAppMessage, error) {
	var sentWebAppMessage *SentWebAppMessage
	err := b.performRequest("answerWebAppQuery", params, &sentWebAppMessage)
	if err != nil {
		return nil, fmt.Errorf("telego: answerWebAppQuery(): %w", err)
	}

	return sentWebAppMessage, nil
}

// SendInvoiceParams - Represents parameters of sendInvoice method.
type SendInvoiceParams struct {
	// ChatID - Unique identifier for the target chat or username of the target channel (in the format
	// @channel_username)
	ChatID ChatID `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Title - Product name, 1-32 characters
	Title string `json:"title"`

	// Description - Product description, 1-255 characters
	Description string `json:"description"`

	// Payload - Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your
	// internal processes.
	Payload string `json:"payload"`

	// ProviderToken - Optional. Payment provider token, obtained via @BotFather (https://t.me/botfather). Pass
	// an empty string for payments in Telegram Stars (https://t.me/BotNews/90).
	ProviderToken string `json:"provider_token,omitempty"`

	// Currency - Three-letter ISO 4217 currency code, see more on currencies
	// (https://core.telegram.org/bots/payments#supported-currencies). Pass “XTR” for payments in Telegram Stars
	// (https://t.me/BotNews/90).
	Currency string `json:"currency"`

	// Prices - Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars
	// (https://t.me/BotNews/90).
	Prices []LabeledPrice `json:"prices"`

	// MaxTipAmount - Optional. The maximum accepted amount for tips in the smallest units of the currency
	// (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the
	// exp parameter in currencies.json (https://core.telegram.org/bots/payments/currencies.json), it shows the
	// number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0.
	// Not supported for payments in Telegram Stars (https://t.me/BotNews/90).
	MaxTipAmount int `json:"max_tip_amount,omitempty"`

	// SuggestedTipAmounts - Optional. A JSON-serialized array of suggested amounts of tips in the smallest
	// units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The
	// suggested tip amounts must be positive, passed in a strictly increased order and must not exceed
	// max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`

	// StartParameter - Optional. Unique deep-linking parameter. If left empty, forwarded copies of the sent
	// message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the
	// same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to
	// the bot (instead of a Pay button), with the value used as the start parameter
	StartParameter string `json:"start_parameter,omitempty"`

	// ProviderData - Optional. JSON-serialized data about the invoice, which will be shared with the payment
	// provider. A detailed description of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// PhotoURL - Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing
	// image for a service. People like it better when they see what they are paying for.
	PhotoURL string `json:"photo_url,omitempty"`

	// PhotoSize - Optional. Photo size in bytes
	PhotoSize int `json:"photo_size,omitempty"`

	// PhotoWidth - Optional. Photo width
	PhotoWidth int `json:"photo_width,omitempty"`

	// PhotoHeight - Optional. Photo height
	PhotoHeight int `json:"photo_height,omitempty"`

	// NeedName - Optional. Pass True if you require the user's full name to complete the order. Ignored for
	// payments in Telegram Stars (https://t.me/BotNews/90).
	NeedName bool `json:"need_name,omitempty"`

	// NeedPhoneNumber - Optional. Pass True if you require the user's phone number to complete the order.
	// Ignored for payments in Telegram Stars (https://t.me/BotNews/90).
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// NeedEmail - Optional. Pass True if you require the user's email address to complete the order. Ignored
	// for payments in Telegram Stars (https://t.me/BotNews/90).
	NeedEmail bool `json:"need_email,omitempty"`

	// NeedShippingAddress - Optional. Pass True if you require the user's shipping address to complete the
	// order. Ignored for payments in Telegram Stars (https://t.me/BotNews/90).
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// SendPhoneNumberToProvider - Optional. Pass True if the user's phone number should be sent to the
	// provider. Ignored for payments in Telegram Stars (https://t.me/BotNews/90).
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// SendEmailToProvider - Optional. Pass True if the user's email address should be sent to the provider.
	// Ignored for payments in Telegram Stars (https://t.me/BotNews/90).
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// IsFlexible - Optional. Pass True if the final price depends on the shipping method. Ignored for payments
	// in Telegram Stars (https://t.me/BotNews/90).
	IsFlexible bool `json:"is_flexible,omitempty"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards). If empty, one 'Pay total price' button will be
	// shown. If not empty, the first button must be a Pay button.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// SendInvoice - Use this method to send invoices. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendInvoice(params *SendInvoiceParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendInvoice", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendInvoice(): %w", err)
	}

	return message, nil
}

// CreateInvoiceLinkParams - Represents parameters of createInvoiceLink method.
type CreateInvoiceLinkParams struct {
	// Title - Product name, 1-32 characters
	Title string `json:"title"`

	// Description - Product description, 1-255 characters
	Description string `json:"description"`

	// Payload - Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your
	// internal processes.
	Payload string `json:"payload"`

	// ProviderToken - Optional. Payment provider token, obtained via @BotFather (https://t.me/botfather). Pass
	// an empty string for payments in Telegram Stars (https://t.me/BotNews/90).
	ProviderToken string `json:"provider_token,omitempty"`

	// Currency - Three-letter ISO 4217 currency code, see more on currencies
	// (https://core.telegram.org/bots/payments#supported-currencies). Pass “XTR” for payments in Telegram Stars
	// (https://t.me/BotNews/90).
	Currency string `json:"currency"`

	// Prices - Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars
	// (https://t.me/BotNews/90).
	Prices []LabeledPrice `json:"prices"`

	// MaxTipAmount - Optional. The maximum accepted amount for tips in the smallest units of the currency
	// (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the
	// exp parameter in currencies.json (https://core.telegram.org/bots/payments/currencies.json), it shows the
	// number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0.
	// Not supported for payments in Telegram Stars (https://t.me/BotNews/90).
	MaxTipAmount int `json:"max_tip_amount,omitempty"`

	// SuggestedTipAmounts - Optional. A JSON-serialized array of suggested amounts of tips in the smallest
	// units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The
	// suggested tip amounts must be positive, passed in a strictly increased order and must not exceed
	// max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`

	// ProviderData - Optional. JSON-serialized data about the invoice, which will be shared with the payment
	// provider. A detailed description of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// PhotoURL - Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing
	// image for a service.
	PhotoURL string `json:"photo_url,omitempty"`

	// PhotoSize - Optional. Photo size in bytes
	PhotoSize int `json:"photo_size,omitempty"`

	// PhotoWidth - Optional. Photo width
	PhotoWidth int `json:"photo_width,omitempty"`

	// PhotoHeight - Optional. Photo height
	PhotoHeight int `json:"photo_height,omitempty"`

	// NeedName - Optional. Pass True if you require the user's full name to complete the order. Ignored for
	// payments in Telegram Stars (https://t.me/BotNews/90).
	NeedName bool `json:"need_name,omitempty"`

	// NeedPhoneNumber - Optional. Pass True if you require the user's phone number to complete the order.
	// Ignored for payments in Telegram Stars (https://t.me/BotNews/90).
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// NeedEmail - Optional. Pass True if you require the user's email address to complete the order. Ignored
	// for payments in Telegram Stars (https://t.me/BotNews/90).
	NeedEmail bool `json:"need_email,omitempty"`

	// NeedShippingAddress - Optional. Pass True if you require the user's shipping address to complete the
	// order. Ignored for payments in Telegram Stars (https://t.me/BotNews/90).
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// SendPhoneNumberToProvider - Optional. Pass True if the user's phone number should be sent to the
	// provider. Ignored for payments in Telegram Stars (https://t.me/BotNews/90).
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// SendEmailToProvider - Optional. Pass True if the user's email address should be sent to the provider.
	// Ignored for payments in Telegram Stars (https://t.me/BotNews/90).
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// IsFlexible - Optional. Pass True if the final price depends on the shipping method. Ignored for payments
	// in Telegram Stars (https://t.me/BotNews/90).
	IsFlexible bool `json:"is_flexible,omitempty"`
}

// CreateInvoiceLink - Use this method to create a link for an invoice. Returns the created invoice link as
// String on success.
func (b *Bot) CreateInvoiceLink(params *CreateInvoiceLinkParams) (*string, error) {
	var invoiceLink *string
	err := b.performRequest("createInvoiceLink", params, &invoiceLink)
	if err != nil {
		return nil, fmt.Errorf("telego: createInvoiceLink(): %w", err)
	}

	return invoiceLink, nil
}

// AnswerShippingQueryParams - Represents parameters of answerShippingQuery method.
type AnswerShippingQueryParams struct {
	// ShippingQueryID - Unique identifier for the query to be answered
	ShippingQueryID string `json:"shipping_query_id"`

	// Ok - Pass True if delivery to the specified address is possible and False if there are any problems (for
	// example, if delivery to the specified address is not possible)
	Ok bool `json:"ok"`

	// ShippingOptions - Optional. Required if ok is True. A JSON-serialized array of available shipping
	// options.
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
	err := b.performRequest("answerShippingQuery", params)
	if err != nil {
		return fmt.Errorf("telego: answerShippingQuery(): %w", err)
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
	err := b.performRequest("answerPreCheckoutQuery", params)
	if err != nil {
		return fmt.Errorf("telego: answerPreCheckoutQuery(): %w", err)
	}

	return nil
}

// RefundStarPaymentParams - Represents parameters of refundStarPayment method.
type RefundStarPaymentParams struct {
	// UserID - Identifier of the user whose payment will be refunded
	UserID int64 `json:"user_id"`

	// TelegramPaymentChargeID - Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
}

// RefundStarPayment - Refunds a successful payment in Telegram Stars (https://t.me/BotNews/90). Returns True
// on success.
func (b *Bot) RefundStarPayment(params *RefundStarPaymentParams) error {
	err := b.performRequest("refundStarPayment", params)
	if err != nil {
		return fmt.Errorf("telego: refundStarPayment(): %w", err)
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
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any
// reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence
// of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the
// issues.
func (b *Bot) SetPassportDataErrors(params *SetPassportDataErrorsParams) error {
	err := b.performRequest("setPassportDataErrors", params)
	if err != nil {
		return fmt.Errorf("telego: setPassportDataErrors(): %w", err)
	}

	return nil
}

// SendGameParams - Represents parameters of sendGame method.
type SendGameParams struct {
	// BusinessConnectionID - Optional. Unique identifier of the business connection on behalf of which the
	// message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// ChatID - Unique identifier for the target chat
	// Note: Should be int64 not ChatID according to documentation (https://core.telegram.org/bots/api#sendgame)
	ChatID int64 `json:"chat_id"`

	// MessageThreadID - Optional. Unique identifier for the target message thread (topic) of the forum; for
	// forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// GameShortName - Short name of the game, serves as the unique identifier for the game. Set up your games
	// via @BotFather (https://t.me/botfather).
	GameShortName string `json:"game_short_name"`

	// DisableNotification - Optional. Sends the message silently
	// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent - Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// MessageEffectID - Optional. Unique identifier of the message effect to be added to the message; for
	// private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// ReplyParameters - Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup - Optional. A JSON-serialized object for an inline keyboard
	// (https://core.telegram.org/bots/features#inline-keyboards). If empty, one 'Play game_title' button will be
	// shown. If not empty, the first button must launch the game.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// SendGame - Use this method to send a game. On success, the sent Message
// (https://core.telegram.org/bots/api#message) is returned.
func (b *Bot) SendGame(params *SendGameParams) (*Message, error) {
	var message *Message
	err := b.performRequest("sendGame", params, &message)
	if err != nil {
		return nil, fmt.Errorf("telego: sendGame(): %w", err)
	}

	return message, nil
}

// SetGameScoreParams - Represents parameters of setGameScore method.
type SetGameScoreParams struct {
	// UserID - User identifier
	UserID int64 `json:"user_id"`

	// Score - New score, must be non-negative
	Score int `json:"score"`

	// Force - Optional. Pass True if the high score is allowed to decrease. This can be useful when fixing
	// mistakes or banning cheaters
	Force bool `json:"force,omitempty"`

	// DisableEditMessage - Optional. Pass True if the game message should not be automatically edited to
	// include the current scoreboard
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`

	// ChatID - Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatID int64 `json:"chat_id,omitempty"`

	// MessageID - Optional. Required if inline_message_id is not specified. Identifier of the sent message
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the
	// inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// SetGameScore - Use this method to set the score of the specified user in a game message. On success, if
// the message is not an inline message, the Message (https://core.telegram.org/bots/api#message) is returned,
// otherwise True is returned. Returns an error, if the new score is not greater than the user's current score
// in the chat and force is False.
func (b *Bot) SetGameScore(params *SetGameScoreParams) (*Message, error) {
	var message *Message
	var success *bool
	err := b.performRequest("setGameScore", params, &message, &success)
	if err != nil {
		return nil, fmt.Errorf("telego: setGameScore(): %w", err)
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

	// InlineMessageID - Optional. Required if chat_id and message_id are not specified. Identifier of the
	// inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// GetGameHighScores - Use this method to get data for high score tables. Will return the score of the
// specified user and several of their neighbors in a game. Returns an Array of GameHighScore
// (https://core.telegram.org/bots/api#gamehighscore) objects.
// This method will currently return scores for the target user, plus two of their closest neighbors on each
// side. Will also return the top three users if the user and their neighbors are not among them. Please note
// that this behavior is subject to change.
func (b *Bot) GetGameHighScores(params *GetGameHighScoresParams) ([]GameHighScore, error) {
	var gameHighScores []GameHighScore
	err := b.performRequest("getGameHighScores", params, &gameHighScores)
	if err != nil {
		return nil, fmt.Errorf("telego: getGameHighScores(): %w", err)
	}

	return gameHighScores, nil
}
