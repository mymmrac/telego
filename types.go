package telego

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/goccy/go-json"

	"github.com/mymmrac/telego/telegoapi"
)

// Update - This object (https://core.telegram.org/bots/api#available-types) represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	// UpdateID - The update's unique identifier. Update identifiers start from a certain positive number and
	// increase sequentially. This ID becomes especially handy if you're using webhooks
	// (https://core.telegram.org/bots/api#setwebhook), since it allows you to ignore repeated updates or to restore
	// the correct update sequence, should they get out of order. If there are no new updates for at least a week,
	// then identifier of the next update will be chosen randomly instead of sequentially.
	UpdateID int `json:"update_id"`

	// Message - Optional. New incoming message of any kind - text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`

	// EditedMessage - Optional. New version of a message that is known to the bot and was edited
	EditedMessage *Message `json:"edited_message,omitempty"`

	// ChannelPost - Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`

	// EditedChannelPost - Optional. New version of a channel post that is known to the bot and was edited
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// InlineQuery - Optional. New incoming inline (https://core.telegram.org/bots/api#inline-mode) query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`

	// ChosenInlineResult - Optional. The result of an inline (https://core.telegram.org/bots/api#inline-mode)
	// query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback
	// collecting (https://core.telegram.org/bots/inline#collecting-feedback) for details on how to enable these
	// updates for your bot.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

	// CallbackQuery - Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

	// ShippingQuery - Optional. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	// PreCheckoutQuery - Optional. New incoming pre-checkout query. Contains full information about checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	// Poll - Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent
	// by the bot
	Poll *Poll `json:"poll,omitempty"`

	// PollAnswer - Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only
	// in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`

	// MyChatMember - Optional. The bot's chat member status was updated in a chat. For private chats, this
	// update is received only when the bot is blocked or unblocked by the user.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`

	// ChatMember - Optional. A chat member's status was updated in a chat. The bot must be an administrator in
	// the chat and must explicitly specify “chat_member” in the list of allowed_updates to receive these
	// updates.
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`

	// ChatJoinRequest - Optional. A request to join the chat has been sent. The bot must have the
	// can_invite_users administrator right in the chat to receive these updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`

	// ctx - Internal context value can be retrieved using [Update.Context] and set by [Update.WithContext].
	// Value can't be cloned; thus, after calling [Update.Clone] or [Update.CloneSafe] ctx will be the same as in the
	// original update.
	ctx context.Context
}

// Clone returns a deep copy of Update.
//
// Warning: Types like [ChatMember] and [MenuButton] require to have their mandatory fields (like status or type) to be
// filled properly, else [Update.Clone] method will panic. To safely clone, use [Update.CloneSafe] method.
func (u Update) Clone() Update {
	update, err := u.CloneSafe()
	if err != nil {
		panic(err)
	}

	return update
}

// CloneSafe returns a deep copy of Update or an error.
//
// Note: Update's context is carried to the copy as is, to change it use [Update.WithContext] method.
func (u Update) CloneSafe() (Update, error) {
	var update Update

	data, err := json.Marshal(u)
	if err != nil {
		return Update{}, fmt.Errorf("telego: clone update: marshal: %w", err)
	}

	err = json.Unmarshal(data, &update)
	if err != nil {
		return Update{}, fmt.Errorf("telego: clone update: unmarshal: %w", err)
	}
	update.ctx = u.ctx

	return update, nil
}

// Context returns the update's context. To change the context, use WithContext.
// The returned context is always non-nil; it defaults to the background context.
func (u Update) Context() context.Context {
	if u.ctx != nil {
		return u.ctx
	}

	return context.Background()
}

// WithContext returns a shallow copy of the update with its context changed to ctx.
//
// Warning: Panics if nil context passed.
func (u Update) WithContext(ctx context.Context) Update {
	if ctx == nil {
		panic("Telego: nil context not allowed")
	}

	u.ctx = ctx
	return u
}

// WebhookInfo - Describes the current status of a webhook.
type WebhookInfo struct {
	// URL - Webhook URL, may be empty if webhook is not set up
	URL string `json:"url"`

	// HasCustomCertificate - True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// PendingUpdateCount - Number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`

	// IPAddress - Optional. Currently used webhook IP address
	IPAddress string `json:"ip_address,omitempty"`

	// LastErrorDate - Optional. Unix time for the most recent error that happened when trying to deliver an
	// update via webhook
	LastErrorDate int64 `json:"last_error_date,omitempty"`

	// LastErrorMessage - Optional. Error message in human-readable format for the most recent error that
	// happened when trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message,omitempty"`

	// LastSynchronizationErrorDate - Optional. Unix time of the most recent error that happened when trying to
	// synchronize available updates with Telegram datacenters
	LastSynchronizationErrorDate int64 `json:"last_synchronization_error_date,omitempty"`

	// MaxConnections - Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook
	// for update delivery
	MaxConnections int `json:"max_connections,omitempty"`

	// AllowedUpdates - Optional. A list of update types the bot is subscribed to. Defaults to all update types
	// except chat_member
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// User - This object represents a Telegram user or bot.
type User struct {
	// ID - Unique identifier for this user or bot. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	ID int64 `json:"id"`

	// IsBot - True, if this user is a bot
	IsBot bool `json:"is_bot"`

	// FirstName - User's or bot's first name
	FirstName string `json:"first_name"`

	// LastName - Optional. User's or bot's last name
	LastName string `json:"last_name,omitempty"`

	// Username - Optional. User's or bot's username
	Username string `json:"username,omitempty"`

	// LanguageCode - Optional. IETF language tag (https://en.wikipedia.org/wiki/IETF_language_tag) of the
	// user's language
	LanguageCode string `json:"language_code,omitempty"`

	// IsPremium - Optional. True, if this user is a Telegram Premium user
	IsPremium bool `json:"is_premium,omitempty"`

	// AddedToAttachmentMenu - Optional. True, if this user added the bot to the attachment menu
	AddedToAttachmentMenu bool `json:"added_to_attachment_menu,omitempty"`

	// CanJoinGroups - Optional. True, if the bot can be invited to groups. Returned only in getMe
	// (https://core.telegram.org/bots/api#getme).
	CanJoinGroups bool `json:"can_join_groups,omitempty"`

	// CanReadAllGroupMessages - Optional. True, if privacy mode
	// (https://core.telegram.org/bots/features#privacy-mode) is disabled for the bot. Returned only in getMe
	// (https://core.telegram.org/bots/api#getme).
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`

	// SupportsInlineQueries - Optional. True, if the bot supports inline queries. Returned only in getMe
	// (https://core.telegram.org/bots/api#getme).
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}

// Chat - This object represents a chat.
type Chat struct {
	// ID - Unique identifier for this chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// identifier.
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

	// IsForum - Optional. True, if the supergroup chat is a forum (has topics
	// (https://telegram.org/blog/topics-in-groups-collectible-usernames#topics-in-groups) enabled)
	IsForum bool `json:"is_forum,omitempty"`

	// Photo - Optional. Chat photo. Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	Photo *ChatPhoto `json:"photo,omitempty"`

	// ActiveUsernames - Optional. If non-empty, the list of all active chat usernames
	// (https://telegram.org/blog/topics-in-groups-collectible-usernames#collectible-usernames); for private chats,
	// supergroups and channels. Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	ActiveUsernames []string `json:"active_usernames,omitempty"`

	// EmojiStatusCustomEmojiID - Optional. Custom emoji identifier of emoji status of the other party in a
	// private chat. Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	EmojiStatusCustomEmojiID string `json:"emoji_status_custom_emoji_id,omitempty"`

	// EmojiStatusExpirationDate - Optional. Expiration date of the emoji status of the other party in a private
	// chat in Unix time, if any. Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	EmojiStatusExpirationDate int64 `json:"emoji_status_expiration_date,omitempty"`

	// Bio - Optional. Bio of the other party in a private chat. Returned only in getChat
	// (https://core.telegram.org/bots/api#getchat).
	Bio string `json:"bio,omitempty"`

	// HasPrivateForwards - Optional. True, if privacy settings of the other party in the private chat allows to
	// use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat
	// (https://core.telegram.org/bots/api#getchat).
	HasPrivateForwards bool `json:"has_private_forwards,omitempty"`

	// HasRestrictedVoiceAndVideoMessages - Optional. True, if the privacy settings of the other party restrict
	// sending voice and video note messages in the private chat. Returned only in getChat
	// (https://core.telegram.org/bots/api#getchat).
	HasRestrictedVoiceAndVideoMessages bool `json:"has_restricted_voice_and_video_messages,omitempty"`

	// JoinToSendMessages - Optional. True, if users need to join the supergroup before they can send messages.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	JoinToSendMessages bool `json:"join_to_send_messages,omitempty"`

	// JoinByRequest - Optional. True, if all users directly joining the supergroup need to be approved by
	// supergroup administrators. Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	JoinByRequest bool `json:"join_by_request,omitempty"`

	// Description - Optional. Description, for groups, supergroups and channel chats. Returned only in getChat
	// (https://core.telegram.org/bots/api#getchat).
	Description string `json:"description,omitempty"`

	// InviteLink - Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in
	// getChat (https://core.telegram.org/bots/api#getchat).
	InviteLink string `json:"invite_link,omitempty"`

	// PinnedMessage - Optional. The most recent pinned message (by sending date). Returned only in getChat
	// (https://core.telegram.org/bots/api#getchat).
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Permissions - Optional. Default chat member permissions, for groups and supergroups. Returned only in
	// getChat (https://core.telegram.org/bots/api#getchat).
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// SlowModeDelay - Optional. For supergroups, the minimum allowed delay between consecutive messages sent by
	// each unpriviledged user; in seconds. Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`

	// MessageAutoDeleteTime - Optional. The time after which all messages sent to the chat will be
	// automatically deleted; in seconds. Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`

	// HasAggressiveAntiSpamEnabled - Optional. True, if aggressive anti-spam checks are enabled in the
	// supergroup. The field is only available to chat administrators. Returned only in getChat
	// (https://core.telegram.org/bots/api#getchat).
	HasAggressiveAntiSpamEnabled bool `json:"has_aggressive_anti_spam_enabled,omitempty"`

	// HasHiddenMembers - Optional. True, if non-administrators can only get the list of bots and administrators
	// in the chat. Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	HasHiddenMembers bool `json:"has_hidden_members,omitempty"`

	// HasProtectedContent - Optional. True, if messages from the chat can't be forwarded to other chats.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	HasProtectedContent bool `json:"has_protected_content,omitempty"`

	// StickerSetName - Optional. For supergroups, name of group sticker set. Returned only in getChat
	// (https://core.telegram.org/bots/api#getchat).
	StickerSetName string `json:"sticker_set_name,omitempty"`

	// CanSetStickerSet - Optional. True, if the bot can change the group sticker set. Returned only in getChat
	// (https://core.telegram.org/bots/api#getchat).
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`

	// LinkedChatID - Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for
	// a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and
	// some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52
	// bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	LinkedChatID int64 `json:"linked_chat_id,omitempty"`

	// Location - Optional. For supergroups, the location to which the supergroup is connected. Returned only in
	// getChat (https://core.telegram.org/bots/api#getchat).
	Location *ChatLocation `json:"location,omitempty"`
}

// Chat types
const (
	ChatTypeSender     = "sender"
	ChatTypePrivate    = "private"
	ChatTypeGroup      = "group"
	ChatTypeSupergroup = "supergroup"
	ChatTypeChannel    = "channel"
)

// Message - This object represents a message.
type Message struct {
	// MessageID - Unique message identifier inside this chat
	MessageID int `json:"message_id"`

	// MessageThreadID - Optional. Unique identifier of a message thread to which the message belongs; for
	// supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// From - Optional. Sender of the message; empty for messages sent to channels. For backward compatibility,
	// the field contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
	From *User `json:"from,omitempty"`

	// SenderChat - Optional. Sender of the message, sent on behalf of a chat. For example, the channel itself
	// for channel posts, the supergroup itself for messages from anonymous group administrators, the linked channel
	// for messages automatically forwarded to the discussion group. For backward compatibility, the field from
	// contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// Date - Date the message was sent in Unix time
	Date int64 `json:"date"`

	// Chat - Conversation the message belongs to
	Chat Chat `json:"chat"`

	// ForwardFrom - Optional. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from,omitempty"`

	// ForwardFromChat - Optional. For messages forwarded from channels or from anonymous administrators,
	// information about the original sender chat
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`

	// ForwardFromMessageID - Optional. For messages forwarded from channels, identifier of the original message
	// in the channel
	ForwardFromMessageID int `json:"forward_from_message_id,omitempty"`

	// ForwardSignature - Optional. For forwarded messages that were originally sent in channels or by an
	// anonymous chat administrator, signature of the message sender if present
	ForwardSignature string `json:"forward_signature,omitempty"`

	// ForwardSenderName - Optional. Sender's name for messages forwarded from users who disallow adding a link
	// to their account in forwarded messages
	ForwardSenderName string `json:"forward_sender_name,omitempty"`

	// ForwardDate - Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardDate int64 `json:"forward_date,omitempty"`

	// IsTopicMessage - Optional. True, if the message is sent to a forum topic
	IsTopicMessage bool `json:"is_topic_message,omitempty"`

	// IsAutomaticForward - Optional. True, if the message is a channel post that was automatically forwarded to
	// the connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`

	// ReplyToMessage - Optional. For replies, the original message. Note that the Message object in this field
	// will not contain further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// ViaBot - Optional. Bot through which the message was sent
	ViaBot *User `json:"via_bot,omitempty"`

	// EditDate - Optional. Date the message was last edited in Unix time
	EditDate int64 `json:"edit_date,omitempty"`

	// HasProtectedContent - Optional. True, if the message can't be forwarded
	HasProtectedContent bool `json:"has_protected_content,omitempty"`

	// MediaGroupID - Optional. The unique identifier of a media message group this message belongs to
	MediaGroupID string `json:"media_group_id,omitempty"`

	// AuthorSignature - Optional. Signature of the post author for messages in channels, or the custom title of
	// an anonymous group administrator
	AuthorSignature string `json:"author_signature,omitempty"`

	// Text - Optional. For text messages, the actual UTF-8 text of the message
	Text string `json:"text,omitempty"`

	// Entities - Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that
	// appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`

	// Animation - Optional. Message is an animation, information about the animation. For backward
	// compatibility, when this field is set, the document field will also be set
	Animation *Animation `json:"animation,omitempty"`

	// Audio - Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`

	// Document - Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`

	// Photo - Optional. Message is a photo, available sizes of the photo
	Photo []PhotoSize `json:"photo,omitempty"`

	// Sticker - Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// Story - Optional. Message is a forwarded story
	Story *Story `json:"story,omitempty"`

	// Video - Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`

	// VideoNote - Optional. Message is a video note (https://telegram.org/blog/video-messages-and-telescope),
	// information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Voice - Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`

	// Caption - Optional. Caption for the animation, audio, document, photo, video or voice
	Caption string `json:"caption,omitempty"`

	// CaptionEntities - Optional. For messages with a caption, special entities like usernames, URLs, bot
	// commands, etc. that appear in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// HasMediaSpoiler - Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Contact - Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`

	// Dice - Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`

	// Game - Optional. Message is a game, information about the game. More about games »
	// (https://core.telegram.org/bots/api#games)
	Game *Game `json:"game,omitempty"`

	// Poll - Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`

	// Venue - Optional. Message is a venue, information about the venue. For backward compatibility, when this
	// field is set, the location field will also be set
	Venue *Venue `json:"venue,omitempty"`

	// Location - Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`

	// NewChatMembers - Optional. New members that were added to the group or supergroup and information about
	// them (the bot itself may be one of these members)
	NewChatMembers []User `json:"new_chat_members,omitempty"`

	// LeftChatMember - Optional. A member was removed from the group, information about them (this member may
	// be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`

	// NewChatTitle - Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`

	// NewChatPhoto - Optional. A chat photo was change to this value
	NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`

	// DeleteChatPhoto - Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`

	// GroupChatCreated - Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`

	// SupergroupChatCreated - Optional. Service message: the supergroup has been created. This field can't be
	// received in a message coming through updates, because bot can't be a member of a supergroup when it is
	// created. It can only be found in reply_to_message if someone replies to a very first message in a directly
	// created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`

	// ChannelChatCreated - Optional. Service message: the channel has been created. This field can't be
	// received in a message coming through updates, because bot can't be a member of a channel when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`

	// MessageAutoDeleteTimerChanged - Optional. Service message: auto-delete timer settings changed in the chat
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`

	// MigrateToChatID - Optional. The group has been migrated to a supergroup with the specified identifier.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// MigrateFromChatID - Optional. The supergroup has been migrated from a group with the specified
	// identifier. This number may have more than 32 significant bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`

	// PinnedMessage - Optional. Specified message was pinned. Note that the Message object in this field will
	// not contain further reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Invoice - Optional. Message is an invoice for a payment (https://core.telegram.org/bots/api#payments),
	// information about the invoice. More about payments » (https://core.telegram.org/bots/api#payments)
	Invoice *Invoice `json:"invoice,omitempty"`

	// SuccessfulPayment - Optional. Message is a service message about a successful payment, information about
	// the payment. More about payments » (https://core.telegram.org/bots/api#payments)
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`

	// UserShared - Optional. Service message: a user was shared with the bot
	UserShared *UserShared `json:"user_shared,omitempty"`

	// ChatShared - Optional. Service message: a chat was shared with the bot
	ChatShared *ChatShared `json:"chat_shared,omitempty"`

	// ConnectedWebsite - Optional. The domain name of the website on which the user has logged in. More about
	// Telegram Login » (https://core.telegram.org/widgets/login)
	ConnectedWebsite string `json:"connected_website,omitempty"`

	// WriteAccessAllowed - Optional. Service message: the user allowed the bot to write messages after adding
	// it to the attachment or side menu, launching a Web App from a link, or accepting an explicit request from a
	// Web App sent by the method requestWriteAccess (https://core.telegram.org/bots/webapps#initializing-mini-apps)
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed,omitempty"`

	// PassportData - Optional. Telegram Passport data
	PassportData *PassportData `json:"passport_data,omitempty"`

	// ProximityAlertTriggered - Optional. Service message. A user in the chat triggered another user's
	// proximity alert while sharing Live Location.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`

	// ForumTopicCreated - Optional. Service message: forum topic created
	ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created,omitempty"`

	// ForumTopicEdited - Optional. Service message: forum topic edited
	ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited,omitempty"`

	// ForumTopicClosed - Optional. Service message: forum topic closed
	ForumTopicClosed *ForumTopicClosed `json:"forum_topic_closed,omitempty"`

	// ForumTopicReopened - Optional. Service message: forum topic reopened
	ForumTopicReopened *ForumTopicReopened `json:"forum_topic_reopened,omitempty"`

	// GeneralForumTopicHidden - Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicHidden *GeneralForumTopicHidden `json:"general_forum_topic_hidden,omitempty"`

	// GeneralForumTopicUnhidden - Optional. Service message: the 'General' forum topic unhidden
	GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden,omitempty"`

	// VideoChatScheduled - Optional. Service message: video chat scheduled
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`

	// VideoChatStarted - Optional. Service message: video chat started
	VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`

	// VideoChatEnded - Optional. Service message: video chat ended
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`

	// VideoChatParticipantsInvited - Optional. Service message: new participants invited to a video chat
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`

	// WebAppData - Optional. Service message: data sent by a Web App
	WebAppData *WebAppData `json:"web_app_data,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard attached to the message. login_url buttons are represented as
	// ordinary URL buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// MessageID - This object represents a unique message identifier.
type MessageID struct {
	// MessageID - Unique message identifier
	MessageID int `json:"message_id"`
}

// MessageEntity - This object represents one special entity in a text message. For example, hashtags,
// usernames, URLs, etc.
type MessageEntity struct {
	// Type - Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag),
	// “cashtag” ($USD), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email”
	// (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic”
	// (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler”
	// (spoiler message), “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable
	// text URLs), “text_mention” (for users without usernames (https://telegram.org/blog/edit#new-mentions)),
	// “custom_emoji” (for inline custom emoji stickers)
	Type string `json:"type"`

	// Offset - Offset in UTF-16 code units (https://core.telegram.org/api/entities#entity-length) to the start
	// of the entity
	Offset int `json:"offset"`

	// Length - Length of the entity in UTF-16 code units (https://core.telegram.org/api/entities#entity-length)
	Length int `json:"length"`

	// URL - Optional. For “text_link” only, URL that will be opened after user taps on the text
	URL string `json:"url,omitempty"`

	// User - Optional. For “text_mention” only, the mentioned user
	User *User `json:"user,omitempty"`

	// Language - Optional. For “pre” only, the programming language of the entity text
	Language string `json:"language,omitempty"`

	// CustomEmojiID - Optional. For “custom_emoji” only, unique identifier of the custom emoji. Use
	// getCustomEmojiStickers (https://core.telegram.org/bots/api#getcustomemojistickers) to get full information
	// about the sticker
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// MessageEntity types
const (
	EntityTypeMention       = "mention"
	EntityTypeHashtag       = "hashtag"
	EntityTypeCashtag       = "cashtag"
	EntityTypeBotCommand    = "bot_command"
	EntityTypeURL           = "url"
	EntityTypeEmail         = "email"
	EntityTypePhoneNumber   = "phone_number"
	EntityTypeBold          = "bold"
	EntityTypeItalic        = "italic"
	EntityTypeUnderline     = "underline"
	EntityTypeStrikethrough = "strikethrough"
	EntityTypeSpoiler       = "spoiler"
	EntityTypeCode          = "code"
	EntityTypePre           = "pre"
	EntityTypeTextLink      = "text_link"
	EntityTypeTextMention   = "text_mention"
	EntityTypeCustomEmoji   = "custom_emoji"
)

// PhotoSize - This object represents one size of a photo or a file
// (https://core.telegram.org/bots/api#document) / sticker (https://core.telegram.org/bots/api#sticker)
// thumbnail.
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

	// FileSize - Optional. File size in bytes
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

	// Thumbnail - Optional. Animation thumbnail as defined by sender
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// FileName - Optional. Original animation filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// FileSize - Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`
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

	// FileSize - Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`

	// Thumbnail - Optional. Thumbnail of the album cover to which the music file belongs
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// Document - This object represents a general file (as opposed to photos
// (https://core.telegram.org/bots/api#photosize), voice messages (https://core.telegram.org/bots/api#voice) and
// audio files (https://core.telegram.org/bots/api#audio)).
type Document struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Thumbnail - Optional. Document thumbnail as defined by sender
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// FileName - Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// FileSize - Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`
}

// Story - This object represents a message about a forwarded story in the chat. Currently holds no
// information.
type Story struct{}

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

	// Thumbnail - Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// FileName - Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// FileSize - Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`
}

// VideoNote - This object represents a video message
// (https://telegram.org/blog/video-messages-and-telescope) (available in Telegram apps as of v.4.0
// (https://telegram.org/blog/video-messages-and-telescope)).
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

	// Thumbnail - Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// FileSize - Optional. File size in bytes
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

	// FileSize - Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`
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
	// bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most
	// 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	UserID int64 `json:"user_id,omitempty"`

	// Vcard - Optional. Additional data about the contact in the form of a vCard
	// (https://en.wikipedia.org/wiki/VCard)
	Vcard string `json:"vcard,omitempty"`
}

// Dice - This object represents an animated emoji that displays a random value.
type Dice struct {
	// Emoji - Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`

	// Value - Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀”
	// and “⚽” base emoji, 1-64 for “🎰” base emoji
	Value int `json:"value"`
}

// Dice emojis
const (
	EmojiDice        = "🎲"
	EmojiDarts       = "🎯"
	EmojiBowling     = "🎳"
	EmojiBasketball  = "🏀"
	EmojiSoccer      = "⚽"
	EmojiSlotMachine = "🎰"
)

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

	// VoterChat - Optional. The chat that changed the answer to the poll, if the voter is anonymous
	VoterChat *Chat `json:"voter_chat,omitempty"`

	// User - Optional. The user that changed the answer to the poll, if the voter isn't anonymous
	User *User `json:"user,omitempty"`

	// OptionIDs - 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
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

	// CorrectOptionID - Optional. 0-based identifier of the correct answer option. Available only for polls in
	// the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	CorrectOptionID int `json:"correct_option_id,omitempty"`

	// Explanation - Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp
	// icon in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation,omitempty"`

	// ExplanationEntities - Optional. Special entities like usernames, URLs, bot commands, etc. that appear in
	// the explanation
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`

	// OpenPeriod - Optional. Amount of time in seconds the poll will be active after creation
	OpenPeriod int `json:"open_period,omitempty"`

	// CloseDate - Optional. Point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate int64 `json:"close_date,omitempty"`
}

// Poll types
const (
	PollTypeRegular = "regular"
	PollTypeQuiz    = "quiz"
)

// Location - This object represents a point on the map.
type Location struct {
	// Longitude - Longitude as defined by sender
	Longitude float64 `json:"longitude"`

	// Latitude - Latitude as defined by sender
	Latitude float64 `json:"latitude"`

	// HorizontalAccuracy - Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// LivePeriod - Optional. Time relative to the message sending date, during which the location can be
	// updated; in seconds. For active live locations only.
	LivePeriod int `json:"live_period,omitempty"`

	// Heading - Optional. The direction in which user is moving, in degrees; 1-360. For active live locations
	// only.
	Heading int `json:"heading,omitempty"`

	// ProximityAlertRadius - Optional. The maximum distance for proximity alerts about approaching another chat
	// member, in meters. For sent live locations only.
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

	// GooglePlaceType - Optional. Google Places type of the venue. (See supported types
	// (https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// WebAppData - Describes data sent from a Web App (https://core.telegram.org/bots/webapps) to the bot.
type WebAppData struct {
	// Data - The data. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`

	// ButtonText - Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad
	// client can send arbitrary data in this field.
	ButtonText string `json:"button_text"`
}

// ProximityAlertTriggered - This object represents the content of a service message, sent whenever a user in
// the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	// Traveler - User that triggered the alert
	Traveler User `json:"traveler"`

	// Watcher - User that set the alert
	Watcher User `json:"watcher"`

	// Distance - The distance between the users
	Distance int `json:"distance"`
}

// MessageAutoDeleteTimerChanged - This object represents a service message about a change in auto-delete
// timer settings.
type MessageAutoDeleteTimerChanged struct {
	// MessageAutoDeleteTime - New auto-delete time for messages in the chat; in seconds
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// ForumTopicCreated - This object represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	// Name - Name of the topic
	Name string `json:"name"`

	// IconColor - Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`

	// IconCustomEmojiID - Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicClosed - This object represents a service message about a forum topic closed in the chat.
// Currently holds no information.
type ForumTopicClosed struct{}

// ForumTopicEdited - This object represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	// Name - Optional. New name of the topic, if it was edited
	Name string `json:"name,omitempty"`

	// IconCustomEmojiID - Optional. New identifier of the custom emoji shown as the topic icon, if it was
	// edited; an empty string if the icon was removed
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicReopened - This object represents a service message about a forum topic reopened in the chat.
// Currently holds no information.
type ForumTopicReopened struct{}

// GeneralForumTopicHidden - This object represents a service message about General forum topic hidden in the
// chat. Currently holds no information.
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden - This object represents a service message about General forum topic unhidden in
// the chat. Currently holds no information.
type GeneralForumTopicUnhidden struct{}

// UserShared - This object contains information about the user whose identifier was shared with the bot
// using a KeyboardButtonRequestUser (https://core.telegram.org/bots/api#keyboardbuttonrequestuser) button.
type UserShared struct {
	// RequestID - Identifier of the request
	RequestID int `json:"request_id"`

	// UserID - Identifier of the shared user. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	// The bot may not have access to the user and could be unable to use this identifier, unless the user is
	// already known to the bot by some other means.
	UserID int64 `json:"user_id"`
}

// ChatShared - This object contains information about the chat whose identifier was shared with the bot
// using a KeyboardButtonRequestChat (https://core.telegram.org/bots/api#keyboardbuttonrequestchat) button.
type ChatShared struct {
	// RequestID - Identifier of the request
	RequestID int `json:"request_id"`

	// ChatID - Identifier of the shared chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	// The bot may not have access to the chat and could be unable to use this identifier, unless the chat is
	// already known to the bot by some other means.
	ChatID int64 `json:"chat_id"`
}

// WriteAccessAllowed - This object represents a service message about a user allowing a bot to write
// messages after adding it to the attachment menu, launching a Web App from a link, or accepting an explicit
// request from a Web App sent by the method requestWriteAccess
// (https://core.telegram.org/bots/webapps#initializing-mini-apps).
type WriteAccessAllowed struct {
	// FromRequest - Optional. True, if the access was granted after the user accepted an explicit request from
	// a Web App sent by the method requestWriteAccess
	// (https://core.telegram.org/bots/webapps#initializing-mini-apps)
	FromRequest bool `json:"from_request,omitempty"`

	// WebAppName - Optional. Name of the Web App, if the access was granted when the Web App was launched from
	// a link
	WebAppName string `json:"web_app_name,omitempty"`

	// FromAttachmentMenu - Optional. True, if the access was granted when the bot was added to the attachment
	// or side menu
	FromAttachmentMenu bool `json:"from_attachment_menu,omitempty"`
}

// VideoChatScheduled - This object represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	// StartDate - Point in time (Unix timestamp) when the video chat is supposed to be started by a chat
	// administrator
	StartDate int64 `json:"start_date"`
}

// VideoChatStarted - This object represents a service message about a video chat started in the chat.
// Currently holds no information.
type VideoChatStarted struct{}

// VideoChatEnded - This object represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	// Duration - Video chat duration in seconds
	Duration int `json:"duration"`
}

// VideoChatParticipantsInvited - This object represents a service message about new members invited to a
// video chat.
type VideoChatParticipantsInvited struct {
	// Users - New members that were invited to the video chat
	Users []User `json:"users"`
}

// UserProfilePhotos - This object represent a user's profile pictures.
type UserProfilePhotos struct {
	// TotalCount - Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`

	// Photos - Requested profile pictures (in up to 4 sizes each)
	Photos [][]PhotoSize `json:"photos"`
}

// File - This object represents a file ready to be downloaded. The file can be downloaded via the link
// https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at
// least 1 hour. When the link expires, a new one can be requested by calling getFile
// (https://core.telegram.org/bots/api#getfile).
// The maximum file size to download is 20 MB
type File struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// FileSize - Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`

	// FilePath - Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath string `json:"file_path,omitempty"`
}

// WebAppInfo - Describes a Web App (https://core.telegram.org/bots/webapps).
type WebAppInfo struct {
	// URL - An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
	// (https://core.telegram.org/bots/webapps#initializing-mini-apps)
	URL string `json:"url"`
}

// ReplyMarkup - Represents reply markup (inline keyboard, custom reply keyboard, etc.)
type ReplyMarkup interface {
	// ReplyType - Returns type of reply
	ReplyType() string
}

// ReplyMarkup types
const (
	MarkupTypeReplyKeyboard       = "ReplyKeyboardMarkup"
	MarkupTypeReplyKeyboardRemove = "ReplyKeyboardRemove"
	MarkupTypeInlineKeyboard      = "InlineKeyboardMarkup"
	MarkupTypeForceReply          = "ForceReply"
)

// ReplyKeyboardMarkup - This object represents a custom keyboard
// (https://core.telegram.org/bots/features#keyboards) with reply options (see Introduction to bots
// (https://core.telegram.org/bots/features#keyboards) for details and examples).
type ReplyKeyboardMarkup struct {
	// Keyboard - Array of button rows, each represented by an Array of KeyboardButton
	// (https://core.telegram.org/bots/api#keyboardbutton) objects
	Keyboard [][]KeyboardButton `json:"keyboard"`

	// IsPersistent - Optional. Requests clients to always show the keyboard when the regular keyboard is
	// hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	IsPersistent bool `json:"is_persistent,omitempty"`

	// ResizeKeyboard - Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make
	// the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom
	// keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// OneTimeKeyboard - Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard
	// will still be available, but clients will automatically display the usual letter-keyboard in the chat - the
	// user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// InputFieldPlaceholder - Optional. The placeholder to be shown in the input field when the keyboard is
	// active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective - Optional. Use this parameter if you want to show the keyboard to specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message (https://core.telegram.org/bots/api#message)
	// object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	// Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select
	// the new language. Other users in the group don't see the keyboard.
	Selective bool `json:"selective,omitempty"`
}

// ReplyType - Returns ReplyKeyboardMarkup type
func (r *ReplyKeyboardMarkup) ReplyType() string {
	return MarkupTypeReplyKeyboard
}

// KeyboardButton - This object represents one button of the reply keyboard. For simple text buttons, String
// can be used instead of this object to specify the button text. The optional fields web_app, request_user,
// request_chat, request_contact, request_location, and request_poll are mutually exclusive.
type KeyboardButton struct {
	// Text - Text of the button. If none of the optional fields are used, it will be sent as a message when the
	// button is pressed
	Text string `json:"text"`

	// RequestUser - Optional. If specified, pressing the button will open a list of suitable users. Tapping on
	// any user will send their identifier to the bot in a “user_shared” service message. Available in private
	// chats only.
	RequestUser *KeyboardButtonRequestUser `json:"request_user,omitempty"`

	// RequestChat - Optional. If specified, pressing the button will open a list of suitable chats. Tapping on
	// a chat will send its identifier to the bot in a “chat_shared” service message. Available in private chats
	// only.
	RequestChat *KeyboardButtonRequestChat `json:"request_chat,omitempty"`

	// RequestContact - Optional. If True, the user's phone number will be sent as a contact when the button is
	// pressed. Available in private chats only.
	RequestContact bool `json:"request_contact,omitempty"`

	// RequestLocation - Optional. If True, the user's current location will be sent when the button is pressed.
	// Available in private chats only.
	RequestLocation bool `json:"request_location,omitempty"`

	// RequestPoll - Optional. If specified, the user will be asked to create a poll and send it to the bot when
	// the button is pressed. Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`

	// WebApp - Optional. If specified, the described Web App (https://core.telegram.org/bots/webapps) will be
	// launched when the button is pressed. The Web App will be able to send a “web_app_data” service message.
	// Available in private chats only.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// KeyboardButtonRequestUser - This object defines the criteria used to request a suitable user. The
// identifier of the selected user will be shared with the bot when the corresponding button is pressed. More
// about requesting users » (https://core.telegram.org/bots/features#chat-and-user-selection)
type KeyboardButtonRequestUser struct {
	// RequestID - Signed 32-bit identifier of the request, which will be received back in the UserShared
	// (https://core.telegram.org/bots/api#usershared) object. Must be unique within the message
	RequestID int32 `json:"request_id"`

	// UserIsBot - Optional. Pass True to request a bot, pass False to request a regular user. If not specified,
	// no additional restrictions are applied.
	UserIsBot *bool `json:"user_is_bot,omitempty"`

	// UserIsPremium - Optional. Pass True to request a premium user, pass False to request a non-premium user.
	// If not specified, no additional restrictions are applied.
	UserIsPremium *bool `json:"user_is_premium,omitempty"`
}

// KeyboardButtonRequestChat - This object defines the criteria used to request a suitable chat. The
// identifier of the selected chat will be shared with the bot when the corresponding button is pressed. More
// about requesting chats » (https://core.telegram.org/bots/features#chat-and-user-selection)
type KeyboardButtonRequestChat struct {
	// RequestID - Signed 32-bit identifier of the request, which will be received back in the ChatShared
	// (https://core.telegram.org/bots/api#chatshared) object. Must be unique within the message
	RequestID int32 `json:"request_id"`

	// ChatIsChannel - Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsChannel bool `json:"chat_is_channel"`

	// ChatIsForum - Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat.
	// If not specified, no additional restrictions are applied.
	ChatIsForum *bool `json:"chat_is_forum,omitempty"`

	// ChatHasUsername - Optional. Pass True to request a supergroup or a channel with a username, pass False to
	// request a chat without a username. If not specified, no additional restrictions are applied.
	ChatHasUsername *bool `json:"chat_has_username,omitempty"`

	// ChatIsCreated - Optional. Pass True to request a chat owned by the user. Otherwise, no additional
	// restrictions are applied.
	ChatIsCreated *bool `json:"chat_is_created,omitempty"`

	// UserAdministratorRights - Optional. A JSON-serialized object listing the required administrator rights of
	// the user in the chat. The rights must be a superset of bot_administrator_rights. If not specified, no
	// additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`

	// BotAdministratorRights - Optional. A JSON-serialized object listing the required administrator rights of
	// the bot in the chat. The rights must be a subset of user_administrator_rights. If not specified, no
	// additional restrictions are applied.
	BotAdministratorRights *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`

	// BotIsMember - Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional
	// restrictions are applied.
	BotIsMember *bool `json:"bot_is_member,omitempty"`
}

// KeyboardButtonPollType - This object represents type of a poll, which is allowed to be created and sent
// when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// Type - Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If
	// regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll
	// of any type.
	Type string `json:"type,omitempty"`
}

// ReplyKeyboardRemove - Upon receiving a message with this object, Telegram clients will remove the current
// custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a
// new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after
// the user presses a button (see ReplyKeyboardMarkup (https://core.telegram.org/bots/api#replykeyboardmarkup)).
type ReplyKeyboardRemove struct {
	// RemoveKeyboard - Requests clients to remove the custom keyboard (user will not be able to summon this
	// keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in
	// ReplyKeyboardMarkup (https://core.telegram.org/bots/api#replykeyboardmarkup))
	RemoveKeyboard bool `json:"remove_keyboard"`

	// Selective - Optional. Use this parameter if you want to remove the keyboard for specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message (https://core.telegram.org/bots/api#message)
	// object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	// Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the
	// keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
	Selective bool `json:"selective,omitempty"`
}

// ReplyType - Returns ReplyKeyboardRemove type
func (r *ReplyKeyboardRemove) ReplyType() string {
	return MarkupTypeReplyKeyboardRemove
}

// InlineKeyboardMarkup - This object represents an inline keyboard
// (https://core.telegram.org/bots/features#inline-keyboards) that appears right next to the message it belongs
// to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard - Array of button rows, each represented by an Array of InlineKeyboardButton
	// (https://core.telegram.org/bots/api#inlinekeyboardbutton) objects
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// ReplyType - Returns InlineKeyboardMarkup type
func (i *InlineKeyboardMarkup) ReplyType() string {
	return MarkupTypeInlineKeyboard
}

// InlineKeyboardButton - This object represents one button of an inline keyboard. You must use exactly one
// of the optional fields.
type InlineKeyboardButton struct {
	// Text - Label text on the button
	Text string `json:"text"`

	// URL - Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id>
	// can be used to mention a user by their ID without using a username, if this is allowed by their privacy
	// settings.
	URL string `json:"url,omitempty"`

	// CallbackData - Optional. Data to be sent in a callback query
	// (https://core.telegram.org/bots/api#callbackquery) to the bot when button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`

	// WebApp - Optional. Description of the Web App (https://core.telegram.org/bots/webapps) that will be
	// launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of
	// the user using the method answerWebAppQuery (https://core.telegram.org/bots/api#answerwebappquery). Available
	// only in private chats between a user and the bot.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// LoginURL - Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement
	// for the Telegram Login Widget (https://core.telegram.org/widgets/login).
	LoginURL *LoginURL `json:"login_url,omitempty"`

	// SwitchInlineQuery - Optional. If set, pressing the button will prompt the user to select one of their
	// chats, open that chat and insert the bot's username and the specified inline query in the input field. May be
	// empty, in which case just the bot's username will be inserted.
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"`

	// SwitchInlineQueryCurrentChat - Optional. If set, pressing the button will insert the bot's username and
	// the specified inline query in the current chat's input field. May be empty, in which case only the bot's
	// username will be inserted.
	// This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting
	// something from multiple options.
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"`

	// SwitchInlineQueryChosenChat - Optional. If set, pressing the button will prompt the user to select one of
	// their chats of the specified type, open that chat and insert the bot's username and the specified inline
	// query in the input field
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`

	// CallbackGame - Optional. Description of the game that will be launched when the user presses the button.
	// NOTE: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Pay - Optional. Specify True, to send a Pay button (https://core.telegram.org/bots/api#payments).
	// NOTE: This type of button must always be the first button in the first row and can only be used in invoice
	// messages.
	Pay bool `json:"pay,omitempty"`
}

// LoginURL - This object represents a parameter of the inline keyboard button used to automatically
// authorize a user. Serves as a great replacement for the Telegram Login Widget
// (https://core.telegram.org/widgets/login) when the user is coming from Telegram. All the user needs to do is
// tap/click a button and confirm that they want to log in:
// TITLE (https://core.telegram.org/file/811140015/1734/8VZFkwWXalM.97872/6127fa62d8a0bf2b3c)
// Telegram apps support these buttons as of version 5.7
// (https://telegram.org/blog/privacy-discussions-web-bots#meet-seamless-web-bots).
// Sample bot: @discussbot (https://t.me/discussbot)
type LoginURL struct {
	// URL - An HTTPS URL to be opened with user authorization data added to the query string when the button is
	// pressed. If the user refuses to provide authorization data, the original URL without information about the
	// user will be opened. The data added is the same as described in Receiving authorization data
	// (https://core.telegram.org/widgets/login#receiving-authorization-data).
	// NOTE: You must always check the hash of the received data to verify the authentication and the integrity of
	// the data as described in Checking authorization
	// (https://core.telegram.org/widgets/login#checking-authorization).
	URL string `json:"url"`

	// ForwardText - Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`

	// BotUsername - Optional. Username of a bot, which will be used for user authorization. See Setting up a
	// bot (https://core.telegram.org/widgets/login#setting-up-a-bot) for more details. If not specified, the
	// current bot's username will be assumed. The URL's domain must be the same as the domain linked with the bot.
	// See Linking your domain to the bot (https://core.telegram.org/widgets/login#linking-your-domain-to-the-bot)
	// for more details.
	BotUsername string `json:"bot_username,omitempty"`

	// RequestWriteAccess - Optional. Pass True to request the permission for your bot to send messages to the
	// user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// SwitchInlineQueryChosenChat - This object represents an inline button that switches the current user to
// inline mode in a chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
	// Query - Optional. The default inline query to be inserted in the input field. If left empty, only the
	// bot's username will be inserted
	Query string `json:"query,omitempty"`

	// AllowUserChats - Optional. True, if private chats with users can be chosen
	AllowUserChats bool `json:"allow_user_chats,omitempty"`

	// AllowBotChats - Optional. True, if private chats with bots can be chosen
	AllowBotChats bool `json:"allow_bot_chats,omitempty"`

	// AllowGroupChats - Optional. True, if group and supergroup chats can be chosen
	AllowGroupChats bool `json:"allow_group_chats,omitempty"`

	// AllowChannelChats - Optional. True, if channel chats can be chosen
	AllowChannelChats bool `json:"allow_channel_chats,omitempty"`
}

// CallbackQuery - This object represents an incoming callback query from a callback button in an inline
// keyboard (https://core.telegram.org/bots/features#inline-keyboards). If the button that originated the query
// was attached to a message sent by the bot, the field message will be present. If the button was attached to a
// message sent via the bot (in inline mode (https://core.telegram.org/bots/api#inline-mode)), the field
// inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	// ID - Unique identifier for this query
	ID string `json:"id"`

	// From - Sender
	From User `json:"from"`

	// Message - Optional. Message with the callback button that originated the query. Note that message content
	// and message date will not be available if the message is too old
	Message *Message `json:"message,omitempty"`

	// InlineMessageID - Optional. Identifier of the message sent via the bot in inline mode, that originated
	// the query.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// ChatInstance - Global identifier, uniquely corresponding to the chat to which the message with the
	// callback button was sent. Useful for high scores in games (https://core.telegram.org/bots/api#games).
	ChatInstance string `json:"chat_instance"`

	// Data - Optional. Data associated with the callback button. Be aware that the message originated the query
	// can contain no callback buttons with this data.
	Data string `json:"data,omitempty"`

	// GameShortName - Optional. Short name of a Game (https://core.telegram.org/bots/api#games) to be returned,
	// serves as the unique identifier for the game
	GameShortName string `json:"game_short_name,omitempty"`
}

// ForceReply - Upon receiving a message with this object, Telegram clients will display a reply interface to
// the user (act as if the user has selected the bot's message and tapped 'Reply'). This can be extremely useful
// if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode
// (https://core.telegram.org/bots/features#privacy-mode).
type ForceReply struct {
	// ForceReply - Shows reply interface to the user, as if they manually selected the bot's message and tapped
	// 'Reply'
	ForceReply bool `json:"force_reply"`

	// InputFieldPlaceholder - Optional. The placeholder to be shown in the input field when the reply is
	// active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective - Optional. Use this parameter if you want to force reply from specific users only. Targets: 1)
	// users that are @mentioned in the text of the Message (https://core.telegram.org/bots/api#message) object; 2)
	// if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

// ReplyType - Returns ForceReply type
func (f *ForceReply) ReplyType() string {
	return MarkupTypeForceReply
}

// ChatPhoto - This object represents a chat photo.
type ChatPhoto struct {
	// SmallFileID - File identifier of small (160x160) chat photo. This file_id can be used only for photo
	// download and only for as long as the photo is not changed.
	SmallFileID string `json:"small_file_id"`

	// SmallFileUniqueID - Unique file identifier of small (160x160) chat photo, which is supposed to be the
	// same over time and for different bots. Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id"`

	// BigFileID - File identifier of big (640x640) chat photo. This file_id can be used only for photo download
	// and only for as long as the photo is not changed.
	BigFileID string `json:"big_file_id"`

	// BigFileUniqueID - Unique file identifier of big (640x640) chat photo, which is supposed to be the same
	// over time and for different bots. Can't be used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id"`
}

// ChatInviteLink - Represents an invite link for a chat.
type ChatInviteLink struct {
	// InviteLink - The invite link. If the link was created by another chat administrator, then the second part
	// of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`

	// Creator - Creator of the link
	Creator User `json:"creator"`

	// CreatesJoinRequest - True, if users joining the chat via the link need to be approved by chat
	// administrators
	CreatesJoinRequest bool `json:"creates_join_request"`

	// IsPrimary - True, if the link is primary
	IsPrimary bool `json:"is_primary"`

	// IsRevoked - True, if the link is revoked
	IsRevoked bool `json:"is_revoked"`

	// Name - Optional. Invite link name
	Name string `json:"name,omitempty"`

	// ExpireDate - Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	ExpireDate int64 `json:"expire_date,omitempty"`

	// MemberLimit - Optional. The maximum number of users that can be members of the chat simultaneously after
	// joining the chat via this invite link; 1-99999
	MemberLimit int `json:"member_limit,omitempty"`

	// PendingJoinRequestCount - Optional. Number of pending join requests created using this link
	PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`
}

// ChatAdministratorRights - Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	// IsAnonymous - True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// CanManageChat - True, if the administrator can access the chat event log, chat statistics, boost list in
	// channels, message statistics in channels, see channel members, see anonymous administrators in supergroups
	// and ignore slow mode. Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat"`

	// CanDeleteMessages - True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`

	// CanManageVideoChats - True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// CanRestrictMembers - True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members"`

	// CanPromoteMembers - True, if the administrator can add new administrators with a subset of their own
	// privileges or demote administrators that they have promoted, directly or indirectly (promoted by
	// administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`

	// CanChangeInfo - True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// CanInviteUsers - True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// CanPostMessages - Optional. True, if the administrator can post messages in the channel; channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// CanEditMessages - Optional. True, if the administrator can edit messages of other users and can pin
	// messages; channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// CanPinMessages - Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// CanPostStories - Optional. True, if the administrator can post stories in the channel; channels only
	CanPostStories bool `json:"can_post_stories,omitempty"`

	// CanEditStories - Optional. True, if the administrator can edit stories posted by other users; channels
	// only
	CanEditStories bool `json:"can_edit_stories,omitempty"`

	// CanDeleteStories - Optional. True, if the administrator can delete stories posted by other users;
	// channels only
	CanDeleteStories bool `json:"can_delete_stories,omitempty"`

	// CanManageTopics - Optional. True, if the user is allowed to create, rename, close, and reopen forum
	// topics; supergroups only
	CanManageTopics bool `json:"can_manage_topics,omitempty"`
}

// ChatMember - This object contains information about one member of a chat. Currently, the following 6 types
// of chat members are supported:
// ChatMemberOwner (https://core.telegram.org/bots/api#chatmemberowner)
// ChatMemberAdministrator (https://core.telegram.org/bots/api#chatmemberadministrator)
// ChatMemberMember (https://core.telegram.org/bots/api#chatmembermember)
// ChatMemberRestricted (https://core.telegram.org/bots/api#chatmemberrestricted)
// ChatMemberLeft (https://core.telegram.org/bots/api#chatmemberleft)
// ChatMemberBanned (https://core.telegram.org/bots/api#chatmemberbanned)
type ChatMember interface {
	MemberStatus() string
	MemberUser() User
}

// ChatMember statuses
const (
	MemberStatusCreator       = "creator"
	MemberStatusAdministrator = "administrator"
	MemberStatusMember        = "member"
	MemberStatusRestricted    = "restricted"
	MemberStatusLeft          = "left"
	MemberStatusBanned        = "kicked"
)

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
	case MemberStatusCreator:
		var cm *ChatMemberOwner
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	case MemberStatusAdministrator:
		var cm *ChatMemberAdministrator
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	case MemberStatusMember:
		var cm *ChatMemberMember
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	case MemberStatusRestricted:
		var cm *ChatMemberRestricted
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	case MemberStatusLeft:
		var cm *ChatMemberLeft
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	case MemberStatusBanned:
		var cm *ChatMemberBanned
		err = json.Unmarshal(bytes, &cm)
		c.Data = cm
	default:
		return fmt.Errorf("unknown member status: %q", memberStatus.Status)
	}

	return err
}

// ChatMemberOwner - Represents a chat member (https://core.telegram.org/bots/api#chatmember) that owns the
// chat and has all administrator privileges.
type ChatMemberOwner struct {
	// Status - The member's status in the chat, always “creator”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`

	// IsAnonymous - True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// CustomTitle - Optional. Custom title for this user
	CustomTitle string `json:"custom_title,omitempty"`
}

// MemberStatus returns ChatMember status
func (c *ChatMemberOwner) MemberStatus() string {
	return MemberStatusCreator
}

// MemberUser returns ChatMember User
func (c *ChatMemberOwner) MemberUser() User {
	return c.User
}

// ChatMemberAdministrator - Represents a chat member (https://core.telegram.org/bots/api#chatmember) that
// has some additional privileges.
type ChatMemberAdministrator struct {
	// Status - The member's status in the chat, always “administrator”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`

	// CanBeEdited - True, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited"`

	// IsAnonymous - True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// CanManageChat - True, if the administrator can access the chat event log, chat statistics, boost list in
	// channels, message statistics in channels, see channel members, see anonymous administrators in supergroups
	// and ignore slow mode. Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat"`

	// CanDeleteMessages - True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`

	// CanManageVideoChats - True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// CanRestrictMembers - True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members"`

	// CanPromoteMembers - True, if the administrator can add new administrators with a subset of their own
	// privileges or demote administrators that they have promoted, directly or indirectly (promoted by
	// administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`

	// CanChangeInfo - True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// CanInviteUsers - True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// CanPostMessages - Optional. True, if the administrator can post messages in the channel; channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// CanEditMessages - Optional. True, if the administrator can edit messages of other users and can pin
	// messages; channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// CanPinMessages - Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// CanPostStories - Optional. True, if the administrator can post stories in the channel; channels only
	CanPostStories bool `json:"can_post_stories,omitempty"`

	// CanEditStories - Optional. True, if the administrator can edit stories posted by other users; channels
	// only
	CanEditStories bool `json:"can_edit_stories,omitempty"`

	// CanDeleteStories - Optional. True, if the administrator can delete stories posted by other users;
	// channels only
	CanDeleteStories bool `json:"can_delete_stories,omitempty"`

	// CanManageTopics - Optional. True, if the user is allowed to create, rename, close, and reopen forum
	// topics; supergroups only
	CanManageTopics bool `json:"can_manage_topics,omitempty"`

	// CustomTitle - Optional. Custom title for this user
	CustomTitle string `json:"custom_title,omitempty"`
}

// MemberStatus returns ChatMember status
func (c *ChatMemberAdministrator) MemberStatus() string {
	return MemberStatusAdministrator
}

// MemberUser returns ChatMember User
func (c *ChatMemberAdministrator) MemberUser() User {
	return c.User
}

// ChatMemberMember - Represents a chat member (https://core.telegram.org/bots/api#chatmember) that has no
// additional privileges or restrictions.
type ChatMemberMember struct {
	// Status - The member's status in the chat, always “member”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`
}

// MemberStatus returns ChatMember status
func (c *ChatMemberMember) MemberStatus() string {
	return MemberStatusMember
}

// MemberUser returns ChatMember User
func (c *ChatMemberMember) MemberUser() User {
	return c.User
}

// ChatMemberRestricted - Represents a chat member (https://core.telegram.org/bots/api#chatmember) that is
// under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	// Status - The member's status in the chat, always “restricted”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`

	// IsMember - True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member"`

	// CanSendMessages - True, if the user is allowed to send text messages, contacts, invoices, locations and
	// venues
	CanSendMessages bool `json:"can_send_messages"`

	// CanSendAudios - True, if the user is allowed to send audios
	CanSendAudios bool `json:"can_send_audios"`

	// CanSendDocuments - True, if the user is allowed to send documents
	CanSendDocuments bool `json:"can_send_documents"`

	// CanSendPhotos - True, if the user is allowed to send photos
	CanSendPhotos bool `json:"can_send_photos"`

	// CanSendVideos - True, if the user is allowed to send videos
	CanSendVideos bool `json:"can_send_videos"`

	// CanSendVideoNotes - True, if the user is allowed to send video notes
	CanSendVideoNotes bool `json:"can_send_video_notes"`

	// CanSendVoiceNotes - True, if the user is allowed to send voice notes
	CanSendVoiceNotes bool `json:"can_send_voice_notes"`

	// CanSendPolls - True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls"`

	// CanSendOtherMessages - True, if the user is allowed to send animations, games, stickers and use inline
	// bots
	CanSendOtherMessages bool `json:"can_send_other_messages"`

	// CanAddWebPagePreviews - True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`

	// CanChangeInfo - True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// CanInviteUsers - True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// CanPinMessages - True, if the user is allowed to pin messages
	CanPinMessages bool `json:"can_pin_messages"`

	// CanManageTopics - True, if the user is allowed to create forum topics
	CanManageTopics bool `json:"can_manage_topics"`

	// UntilDate - Date when restrictions will be lifted for this user; Unix time. If 0, then the user is
	// restricted forever
	UntilDate int64 `json:"until_date"`
}

// MemberStatus returns ChatMember status
func (c *ChatMemberRestricted) MemberStatus() string {
	return MemberStatusRestricted
}

// MemberUser returns ChatMember User
func (c *ChatMemberRestricted) MemberUser() User {
	return c.User
}

// ChatMemberLeft - Represents a chat member (https://core.telegram.org/bots/api#chatmember) that isn't
// currently a member of the chat, but may join it themselves.
type ChatMemberLeft struct {
	// Status - The member's status in the chat, always “left”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`
}

// MemberStatus returns ChatMember status
func (c *ChatMemberLeft) MemberStatus() string {
	return MemberStatusLeft
}

// MemberUser returns ChatMember User
func (c *ChatMemberLeft) MemberUser() User {
	return c.User
}

// ChatMemberBanned - Represents a chat member (https://core.telegram.org/bots/api#chatmember) that was
// banned in the chat and can't return to the chat or view chat messages.
type ChatMemberBanned struct {
	// Status - The member's status in the chat, always “kicked”
	Status string `json:"status"`

	// User - Information about the user
	User User `json:"user"`

	// UntilDate - Date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned
	// forever
	UntilDate int64 `json:"until_date"`
}

// MemberStatus returns ChatMember status
func (c *ChatMemberBanned) MemberStatus() string {
	return MemberStatusBanned
}

// MemberUser returns ChatMember User
func (c *ChatMemberBanned) MemberUser() User {
	return c.User
}

// ChatMemberUpdated - This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	// Chat - Chat the user belongs to
	Chat Chat `json:"chat"`

	// From - Performer of the action, which resulted in the change
	From User `json:"from"`

	// Date - Date the change was done in Unix time
	Date int64 `json:"date"`

	// OldChatMember - Previous information about the chat member
	OldChatMember ChatMember `json:"old_chat_member"`

	// NewChatMember - New information about the chat member
	NewChatMember ChatMember `json:"new_chat_member"`

	// InviteLink - Optional. Chat invite link, which was used by the user to join the chat; for joining by
	// invite link events only.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`

	// ViaChatFolderInviteLink - Optional. True, if the user joined the chat via a chat folder invite link
	ViaChatFolderInviteLink bool `json:"via_chat_folder_invite_link,omitempty"`
}

// UnmarshalJSON converts JSON to ChatMemberUpdated
func (c *ChatMemberUpdated) UnmarshalJSON(bytes []byte) error {
	var chatMemberUpdatedData struct {
		Chat          Chat            `json:"chat"`
		From          User            `json:"from"`
		Date          int64           `json:"date"`
		OldChatMember chatMemberData  `json:"old_chat_member"`
		NewChatMember chatMemberData  `json:"new_chat_member"`
		InviteLink    *ChatInviteLink `json:"invite_link,omitempty"`
	}

	err := json.Unmarshal(bytes, &chatMemberUpdatedData)
	if err != nil {
		return err
	}

	c.Chat = chatMemberUpdatedData.Chat
	c.From = chatMemberUpdatedData.From
	c.Date = chatMemberUpdatedData.Date
	c.OldChatMember = chatMemberUpdatedData.OldChatMember.Data
	c.NewChatMember = chatMemberUpdatedData.NewChatMember.Data
	c.InviteLink = chatMemberUpdatedData.InviteLink

	return nil
}

// ChatJoinRequest - Represents a join request sent to a chat.
type ChatJoinRequest struct {
	// Chat - Chat to which the request was sent
	Chat Chat `json:"chat"`

	// From - User that sent the join request
	From User `json:"from"`

	// UserChatID - Identifier of a private chat with the user who sent the join request. This number may have
	// more than 32 significant bits and some programming languages may have difficulty/silent defects in
	// interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type
	// are safe for storing this identifier. The bot can use this identifier for 24 hours to send messages until the
	// join request is processed, assuming no other administrator contacted the user.
	UserChatID int64 `json:"user_chat_id"`

	// Date - Date the request was sent in Unix time
	Date int64 `json:"date"`

	// Bio - Optional. Bio of the user.
	Bio string `json:"bio,omitempty"`

	// InviteLink - Optional. Chat invite link that was used by the user to send the join request
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// ChatPermissions - Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// CanSendMessages - Optional. True, if the user is allowed to send text messages, contacts, invoices,
	// locations and venues
	CanSendMessages *bool `json:"can_send_messages,omitempty"`

	// CanSendAudios - Optional. True, if the user is allowed to send audios
	CanSendAudios *bool `json:"can_send_audios,omitempty"`

	// CanSendDocuments - Optional. True, if the user is allowed to send documents
	CanSendDocuments *bool `json:"can_send_documents,omitempty"`

	// CanSendPhotos - Optional. True, if the user is allowed to send photos
	CanSendPhotos *bool `json:"can_send_photos,omitempty"`

	// CanSendVideos - Optional. True, if the user is allowed to send videos
	CanSendVideos *bool `json:"can_send_videos,omitempty"`

	// CanSendVideoNotes - Optional. True, if the user is allowed to send video notes
	CanSendVideoNotes *bool `json:"can_send_video_notes,omitempty"`

	// CanSendVoiceNotes - Optional. True, if the user is allowed to send voice notes
	CanSendVoiceNotes *bool `json:"can_send_voice_notes,omitempty"`

	// CanSendPolls - Optional. True, if the user is allowed to send polls
	CanSendPolls *bool `json:"can_send_polls,omitempty"`

	// CanSendOtherMessages - Optional. True, if the user is allowed to send animations, games, stickers and use
	// inline bots
	CanSendOtherMessages *bool `json:"can_send_other_messages,omitempty"`

	// CanAddWebPagePreviews - Optional. True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`

	// CanChangeInfo - Optional. True, if the user is allowed to change the chat title, photo and other
	// settings. Ignored in public supergroups
	CanChangeInfo *bool `json:"can_change_info,omitempty"`

	// CanInviteUsers - Optional. True, if the user is allowed to invite new users to the chat
	CanInviteUsers *bool `json:"can_invite_users,omitempty"`

	// CanPinMessages - Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`

	// CanManageTopics - Optional. True, if the user is allowed to create forum topics. If omitted defaults to
	// the value of can_pin_messages
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
}

// ChatLocation - Represents a location to which a chat is connected.
type ChatLocation struct {
	// Location - The location to which the supergroup is connected. Can't be a live location.
	Location Location `json:"location"`

	// Address - Location address; 1-64 characters, as defined by the chat owner
	Address string `json:"address"`
}

// ForumTopic - This object represents a forum topic.
type ForumTopic struct {
	// MessageThreadID - Unique identifier of the forum topic
	MessageThreadID int `json:"message_thread_id"`

	// Name - Name of the topic
	Name string `json:"name"`

	// IconColor - Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`

	// IconCustomEmojiID - Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// BotCommand - This object represents a bot command.
type BotCommand struct {
	// Command - Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and
	// underscores.
	Command string `json:"command"`

	// Description - Description of the command; 1-256 characters.
	Description string `json:"description"`
}

// BotCommandScope - This object represents the scope to which bot commands are applied. Currently, the
// following 7 scopes are supported:
// BotCommandScopeDefault (https://core.telegram.org/bots/api#botcommandscopedefault)
// BotCommandScopeAllPrivateChats (https://core.telegram.org/bots/api#botcommandscopeallprivatechats)
// BotCommandScopeAllGroupChats (https://core.telegram.org/bots/api#botcommandscopeallgroupchats)
// BotCommandScopeAllChatAdministrators
// (https://core.telegram.org/bots/api#botcommandscopeallchatadministrators)
// BotCommandScopeChat (https://core.telegram.org/bots/api#botcommandscopechat)
// BotCommandScopeChatAdministrators (https://core.telegram.org/bots/api#botcommandscopechatadministrators)
// BotCommandScopeChatMember (https://core.telegram.org/bots/api#botcommandscopechatmember)
type BotCommandScope interface {
	ScopeType() string
}

// BotCommandScope types
const (
	ScopeTypeDefault               = "default"
	ScopeTypeAllPrivateChats       = "all_private_chats"
	ScopeTypeAllGroupChats         = "all_group_chats"
	ScopeTypeAllChatAdministrators = "all_chat_administrators"
	ScopeTypeChat                  = "chat"
	ScopeTypeChatAdministrators    = "chat_administrators"
	ScopeTypeChatMember            = "chat_member"
)

// BotCommandScopeDefault - Represents the default scope (https://core.telegram.org/bots/api#botcommandscope)
// of bot commands. Default commands are used if no commands with a narrower scope
// (https://core.telegram.org/bots/api#determining-list-of-commands) are specified for the user.
type BotCommandScopeDefault struct {
	// Type - Scope type, must be default
	Type string `json:"type"`
}

// ScopeType returns BotCommandScope type
func (b *BotCommandScopeDefault) ScopeType() string {
	return ScopeTypeDefault
}

// BotCommandScopeAllPrivateChats - Represents the scope (https://core.telegram.org/bots/api#botcommandscope)
// of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	// Type - Scope type, must be all_private_chats
	Type string `json:"type"`
}

// ScopeType returns BotCommandScope type
func (b *BotCommandScopeAllPrivateChats) ScopeType() string {
	return ScopeTypeAllPrivateChats
}

// BotCommandScopeAllGroupChats - Represents the scope (https://core.telegram.org/bots/api#botcommandscope)
// of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	// Type - Scope type, must be all_group_chats
	Type string `json:"type"`
}

// ScopeType returns BotCommandScope type
func (b *BotCommandScopeAllGroupChats) ScopeType() string {
	return ScopeTypeAllGroupChats
}

// BotCommandScopeAllChatAdministrators - Represents the scope
// (https://core.telegram.org/bots/api#botcommandscope) of bot commands, covering all group and supergroup chat
// administrators.
type BotCommandScopeAllChatAdministrators struct {
	// Type - Scope type, must be all_chat_administrators
	Type string `json:"type"`
}

// ScopeType returns BotCommandScope type
func (b *BotCommandScopeAllChatAdministrators) ScopeType() string {
	return ScopeTypeAllChatAdministrators
}

// ChatID - Represents chat ID as int64 or string
type ChatID struct {
	// ID - Unique identifier for the target chat
	ID int64

	// Username - Channel or group username of the target chat (in the format @chanel_username)
	// Note: User username can't be used here, you have to use integer chat ID
	Username string
}

// String returns string representation of ChatID
func (c ChatID) String() string {
	if c.ID != 0 {
		return strconv.FormatInt(c.ID, 10)
	}

	if c.Username != "" {
		return c.Username
	}

	return ""
}

// MarshalJSON returns JSON representation of ChatID
func (c ChatID) MarshalJSON() ([]byte, error) {
	if c.ID != 0 {
		return json.Marshal(c.ID)
	}

	if c.Username != "" {
		return json.Marshal(c.Username)
	}

	return []byte(`""`), nil
}

// BotCommandScopeChat - Represents the scope (https://core.telegram.org/bots/api#botcommandscope) of bot
// commands, covering a specific chat.
type BotCommandScopeChat struct {
	// Type - Scope type, must be chat
	Type string `json:"type"`

	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`
}

// ScopeType returns BotCommandScope type
func (b *BotCommandScopeChat) ScopeType() string {
	return ScopeTypeChat
}

// BotCommandScopeChatAdministrators - Represents the scope
// (https://core.telegram.org/bots/api#botcommandscope) of bot commands, covering all administrators of a
// specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	// Type - Scope type, must be chat_administrators
	Type string `json:"type"`

	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`
}

// ScopeType returns BotCommandScope type
func (b *BotCommandScopeChatAdministrators) ScopeType() string {
	return ScopeTypeChatAdministrators
}

// BotCommandScopeChatMember - Represents the scope (https://core.telegram.org/bots/api#botcommandscope) of
// bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	// Type - Scope type, must be chat_member
	Type string `json:"type"`

	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username)
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// ScopeType returns BotCommandScope type
func (b *BotCommandScopeChatMember) ScopeType() string {
	return ScopeTypeChatMember
}

// BotName - This object represents the bot's name.
type BotName struct {
	// Name - The bot's name
	Name string `json:"name"`
}

// BotDescription - This object represents the bot's description.
type BotDescription struct {
	// Description - The bot's description
	Description string `json:"description"`
}

// BotShortDescription - This object represents the bot's short description.
type BotShortDescription struct {
	// ShortDescription - The bot's short description
	ShortDescription string `json:"short_description"`
}

// MenuButton - This object describes the bot's menu button in a private chat. It should be one of
// MenuButtonCommands (https://core.telegram.org/bots/api#menubuttoncommands)
// MenuButtonWebApp (https://core.telegram.org/bots/api#menubuttonwebapp)
// MenuButtonDefault (https://core.telegram.org/bots/api#menubuttondefault)
// If a menu button other than MenuButtonDefault (https://core.telegram.org/bots/api#menubuttondefault) is set
// for a private chat, then it is applied in the chat. Otherwise the default menu button is applied. By default,
// the menu button opens the list of bot commands.
type MenuButton interface {
	ButtonType() string
}

// MenuButton types
const (
	ButtonTypeCommands = "commands"
	ButtonTypeWebApp   = "web_app"
	ButtonTypeDefault  = "default"
)

type menuButtonData struct {
	Data MenuButton
}

func (m *menuButtonData) UnmarshalJSON(bytes []byte) error {
	var buttonType struct {
		Type string `json:"type"`
	}

	err := json.Unmarshal(bytes, &buttonType)
	if err != nil {
		return err
	}

	switch buttonType.Type {
	case ButtonTypeCommands:
		var mb *MenuButtonCommands
		err = json.Unmarshal(bytes, &mb)
		m.Data = mb
	case ButtonTypeWebApp:
		var mb *MenuButtonWebApp
		err = json.Unmarshal(bytes, &mb)
		m.Data = mb
	case ButtonTypeDefault:
		var mb *MenuButtonDefault
		err = json.Unmarshal(bytes, &mb)
		m.Data = mb
	default:
		return fmt.Errorf("unknown menu button type: %q", buttonType.Type)
	}

	return err
}

// MenuButtonCommands - Represents a menu button, which opens the bot's list of commands.
type MenuButtonCommands struct {
	// Type - Type of the button, must be commands
	Type string `json:"type"`
}

// ButtonType returns MenuButton type
func (m *MenuButtonCommands) ButtonType() string {
	return ButtonTypeCommands
}

// MenuButtonWebApp - Represents a menu button, which launches a Web App
// (https://core.telegram.org/bots/webapps).
type MenuButtonWebApp struct {
	// Type - Type of the button, must be web_app
	Type string `json:"type"`

	// Text - Text on the button
	Text string `json:"text"`

	// WebApp - Description of the Web App that will be launched when the user presses the button. The Web App
	// will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery
	// (https://core.telegram.org/bots/api#answerwebappquery).
	WebApp WebAppInfo `json:"web_app"`
}

// ButtonType returns MenuButton type
func (m *MenuButtonWebApp) ButtonType() string {
	return ButtonTypeWebApp
}

// MenuButtonDefault - Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {
	// Type - Type of the button, must be default
	Type string `json:"type"`
}

// ButtonType returns MenuButton type
func (m *MenuButtonDefault) ButtonType() string {
	return ButtonTypeDefault
}

// fileCompatible - Represents types that can be sent as files
type fileCompatible interface {
	fileParameters() map[string]telegoapi.NamedReader
}

// InputMedia - This object represents the content of a media message to be sent. It should be one of
// InputMediaAnimation (https://core.telegram.org/bots/api#inputmediaanimation)
// InputMediaDocument (https://core.telegram.org/bots/api#inputmediadocument)
// InputMediaAudio (https://core.telegram.org/bots/api#inputmediaaudio)
// InputMediaPhoto (https://core.telegram.org/bots/api#inputmediaphoto)
// InputMediaVideo (https://core.telegram.org/bots/api#inputmediavideo)
type InputMedia interface {
	MediaType() string
	fileCompatible
}

// InputMedia types
const (
	MediaTypePhoto     = "photo"
	MediaTypeVideo     = "video"
	MediaTypeAnimation = "animation"
	MediaTypeAudio     = "audio"
	MediaTypeDocument  = "document"
)

// InputMediaPhoto - Represents a photo to be sent.
type InputMediaPhoto struct {
	// Type - Type of the result, must be photo
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to
	// upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files
	// » (https://core.telegram.org/bots/api#sending-files)
	Media InputFile `json:"media"`

	// Caption - Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the photo caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// HasSpoiler - Optional. Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// MediaType return InputMedia type
func (i *InputMediaPhoto) MediaType() string {
	return MediaTypePhoto
}

func (i *InputMediaPhoto) fileParameters() map[string]telegoapi.NamedReader {
	i.Media.needAttach = true
	return map[string]telegoapi.NamedReader{
		"media": i.Media.File,
	}
}

// InputMediaVideo - Represents a video to be sent.
type InputMediaVideo struct {
	// Type - Type of the result, must be video
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to
	// upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files
	// » (https://core.telegram.org/bots/api#sending-files)
	Media InputFile `json:"media"`

	// Thumbnail - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumbnail *InputFile `json:"thumbnail,omitempty"`

	// Caption - Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the video caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Width - Optional. Video width
	Width int `json:"width,omitempty"`

	// Height - Optional. Video height
	Height int `json:"height,omitempty"`

	// Duration - Optional. Video duration in seconds
	Duration int `json:"duration,omitempty"`

	// SupportsStreaming - Optional. Pass True if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`

	// HasSpoiler - Optional. Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// MediaType return InputMedia type
func (i *InputMediaVideo) MediaType() string {
	return MediaTypeVideo
}

func (i *InputMediaVideo) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	i.Media.needAttach = true
	fp["media"] = i.Media.File
	if i.Thumbnail != nil {
		i.Thumbnail.needAttach = true
		fp["thumbnail"] = i.Thumbnail.File
	}

	return fp
}

// InputMediaAnimation - Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be
// sent.
type InputMediaAnimation struct {
	// Type - Type of the result, must be animation
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to
	// upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files
	// » (https://core.telegram.org/bots/api#sending-files)
	Media InputFile `json:"media"`

	// Thumbnail - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumbnail *InputFile `json:"thumbnail,omitempty"`

	// Caption - Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the animation caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Width - Optional. Animation width
	Width int `json:"width,omitempty"`

	// Height - Optional. Animation height
	Height int `json:"height,omitempty"`

	// Duration - Optional. Animation duration in seconds
	Duration int `json:"duration,omitempty"`

	// HasSpoiler - Optional. Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// MediaType return InputMedia type
func (i *InputMediaAnimation) MediaType() string {
	return MediaTypeAnimation
}

func (i *InputMediaAnimation) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	i.Media.needAttach = true
	fp["media"] = i.Media.File
	if i.Thumbnail != nil {
		i.Thumbnail.needAttach = true
		fp["thumbnail"] = i.Thumbnail.File
	}

	return fp
}

// InputMediaAudio - Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	// Type - Type of the result, must be audio
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to
	// upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files
	// » (https://core.telegram.org/bots/api#sending-files)
	Media InputFile `json:"media"`

	// Thumbnail - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumbnail *InputFile `json:"thumbnail,omitempty"`

	// Caption - Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the audio caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Duration - Optional. Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`

	// Performer - Optional. Performer of the audio
	Performer string `json:"performer,omitempty"`

	// Title - Optional. Title of the audio
	Title string `json:"title,omitempty"`
}

// MediaType return InputMedia type
func (i *InputMediaAudio) MediaType() string {
	return MediaTypeAudio
}

func (i *InputMediaAudio) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	i.Media.needAttach = true
	fp["media"] = i.Media.File
	if i.Thumbnail != nil {
		i.Thumbnail.needAttach = true
		fp["thumbnail"] = i.Thumbnail.File
	}

	return fp
}

// InputMediaDocument - Represents a general file to be sent.
type InputMediaDocument struct {
	// Type - Type of the result, must be document
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to
	// upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files
	// » (https://core.telegram.org/bots/api#sending-files)
	Media InputFile `json:"media"`

	// Thumbnail - Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Thumbnail *InputFile `json:"thumbnail,omitempty"`

	// Caption - Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the document caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// DisableContentTypeDetection - Optional. Disables automatic server-side content type detection for files
	// uploaded using multipart/form-data. Always True, if the document is sent as part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}

// MediaType return InputMedia type
func (i *InputMediaDocument) MediaType() string {
	return MediaTypeDocument
}

func (i *InputMediaDocument) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	i.Media.needAttach = true
	fp["media"] = i.Media.File
	if i.Thumbnail != nil {
		i.Thumbnail.needAttach = true
		fp["thumbnail"] = i.Thumbnail.File
	}

	return fp
}

// InputFile - This object represents the contents of a file to be uploaded. Must be posted using
// multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile struct {
	// File - Object that can be treated as file (has name and data to read).
	// Implemented by os.File.
	File telegoapi.NamedReader

	// FileID - ID of file stored in Telegram
	FileID string

	// URL - URL to get file from
	URL string

	// needAttach used to specify that file field will be named the same as file name
	needAttach bool
}

// String returns string representation of InputFile
func (i InputFile) String() string {
	if i.FileID != "" {
		return i.FileID
	}

	if i.URL != "" {
		return i.URL
	}

	if i.File != nil {
		return i.File.Name()
	}

	return ""
}

// MarshalJSON return JSON representation of InputFile
func (i InputFile) MarshalJSON() ([]byte, error) {
	if i.FileID != "" {
		return json.Marshal(i.FileID)
	}

	if i.URL != "" {
		return json.Marshal(i.URL)
	}

	if !isNil(i.File) {
		if i.needAttach {
			return json.Marshal(attachFile + i.File.Name())
		}
		return []byte(`""`), nil
	}

	return nil, errors.New("telego: file ID, URL and file are empty")
}

// Sticker - This object represents a sticker.
type Sticker struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Type - Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of
	// the sticker is independent from its format, which is determined by the fields is_animated and is_video.
	Type string `json:"type"`

	// Width - Sticker width
	Width int `json:"width"`

	// Height - Sticker height
	Height int `json:"height"`

	// IsAnimated - True, if the sticker is animated (https://telegram.org/blog/animated-stickers)
	IsAnimated bool `json:"is_animated"`

	// IsVideo - True, if the sticker is a video sticker
	// (https://telegram.org/blog/video-stickers-better-reactions)
	IsVideo bool `json:"is_video"`

	// Thumbnail - Optional. Sticker thumbnail in the .WEBP or .JPG format
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// Emoji - Optional. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`

	// SetName - Optional. Name of the sticker set to which the sticker belongs
	SetName string `json:"set_name,omitempty"`

	// PremiumAnimation - Optional. For premium regular stickers, premium animation for the sticker
	PremiumAnimation *File `json:"premium_animation,omitempty"`

	// MaskPosition - Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// CustomEmojiID - Optional. For custom emoji stickers, unique identifier of the custom emoji
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`

	// NeedsRepainting - Optional. True, if the sticker must be repainted to a text color in messages, the color
	// of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in
	// other places
	NeedsRepainting bool `json:"needs_repainting,omitempty"`

	// FileSize - Optional. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// Sticker types
const (
	StickerTypeRegular     = "regular"
	StickerTypeMask        = "mask"
	StickerTypeCustomEmoji = "custom_emoji"
)

// StickerSet - This object represents a sticker set.
type StickerSet struct {
	// Name - Sticker set name
	Name string `json:"name"`

	// Title - Sticker set title
	Title string `json:"title"`

	// StickerType - Type of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”
	StickerType string `json:"sticker_type"`

	// IsAnimated - True, if the sticker set contains animated stickers
	// (https://telegram.org/blog/animated-stickers)
	IsAnimated bool `json:"is_animated"`

	// IsVideo - True, if the sticker set contains video stickers
	// (https://telegram.org/blog/video-stickers-better-reactions)
	IsVideo bool `json:"is_video"`

	// Stickers - List of all set stickers
	Stickers []Sticker `json:"stickers"`

	// Thumbnail - Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// MaskPosition - This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// Point - The part of the face relative to which the mask should be placed. One of “forehead”,
	// “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`

	// XShift - Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For
	// example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float64 `json:"x_shift"`

	// YShift - Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For
	// example, 1.0 will place the mask just below the default mask position.
	YShift float64 `json:"y_shift"`

	// Scale - Mask scaling coefficient. For example, 2.0 means double size.
	Scale float64 `json:"scale"`
}

// MaskPosition points
const (
	PointForehead = "forehead"
	PointEyes     = "eyes"
	PointMouth    = "mouth"
	PointChin     = "chin"
)

// InputSticker - This object describes a sticker to be added to a sticker set.
type InputSticker struct {
	// Sticker - The added sticker. Pass a file_id as a String to send a file that already exists on the
	// Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, upload a new one
	// using multipart/form-data, or pass “attach://<file_attach_name>” to upload a new one using
	// multipart/form-data under <file_attach_name> name. Animated and video stickers can't be uploaded via HTTP
	// URL. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Sticker InputFile `json:"sticker"`

	// EmojiList - List of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list"`

	// MaskPosition - Optional. Position where the mask should be placed on faces. For “mask” stickers only.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// Keywords - Optional. List of 0-20 search keywords for the sticker with total length of up to 64
	// characters. For “regular” and “custom_emoji” stickers only.
	Keywords []string `json:"keywords,omitempty"`
}

// InlineQuery - This object represents an incoming inline query. When the user sends an empty query, your
// bot could return some default or trending results.
type InlineQuery struct {
	// ID - Unique identifier for this query
	ID string `json:"id"`

	// From - Sender
	From User `json:"from"`

	// Query - Text of the query (up to 256 characters)
	Query string `json:"query"`

	// Offset - Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`

	// ChatType - Optional. Type of the chat from which the inline query was sent. Can be either “sender”
	// for a private chat with the inline query sender, “private”, “group”, “supergroup”, or
	// “channel”. The chat type should be always known for requests sent from official clients and most
	// third-party clients, unless the request was sent from a secret chat
	ChatType string `json:"chat_type,omitempty"`

	// Location - Optional. Sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`
}

// InlineQueryResultsButton - This object represents a button to be shown above inline query results. You
// must use exactly one of the optional fields.
type InlineQueryResultsButton struct {
	// Text - Label text on the button
	Text string `json:"text"`

	// WebApp - Optional. Description of the Web App (https://core.telegram.org/bots/webapps) that will be
	// launched when the user presses the button. The Web App will be able to switch back to the inline mode using
	// the method switchInlineQuery (https://core.telegram.org/bots/webapps#initializing-mini-apps) inside the Web
	// App.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// StartParameter - Optional. Deep-linking (https://core.telegram.org/bots/features#deep-linking) parameter
	// for the /start message sent to the bot when a user presses the button. 1-64 characters, only A-Z, a-z, 0-9, _
	// and - are allowed.
	// Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account
	// to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account' button above
	// the results, or even before showing any. The user presses the button, switches to a private chat with the bot
	// and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done, the bot
	// can offer a switch_inline (https://core.telegram.org/bots/api#inlinekeyboardmarkup) button so that the user
	// can easily return to the chat where they wanted to use the bot's inline capabilities.
	StartParameter string `json:"start_parameter,omitempty"`
}

// InlineQueryResult - This object represents one result of an inline query. Telegram clients currently
// support results of the following 20 types:
// InlineQueryResultCachedAudio (https://core.telegram.org/bots/api#inlinequeryresultcachedaudio)
// InlineQueryResultCachedDocument (https://core.telegram.org/bots/api#inlinequeryresultcacheddocument)
// InlineQueryResultCachedGif (https://core.telegram.org/bots/api#inlinequeryresultcachedgif)
// InlineQueryResultCachedMpeg4Gif (https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif)
// InlineQueryResultCachedPhoto (https://core.telegram.org/bots/api#inlinequeryresultcachedphoto)
// InlineQueryResultCachedSticker (https://core.telegram.org/bots/api#inlinequeryresultcachedsticker)
// InlineQueryResultCachedVideo (https://core.telegram.org/bots/api#inlinequeryresultcachedvideo)
// InlineQueryResultCachedVoice (https://core.telegram.org/bots/api#inlinequeryresultcachedvoice)
// InlineQueryResultArticle (https://core.telegram.org/bots/api#inlinequeryresultarticle)
// InlineQueryResultAudio (https://core.telegram.org/bots/api#inlinequeryresultaudio)
// InlineQueryResultContact (https://core.telegram.org/bots/api#inlinequeryresultcontact)
// InlineQueryResultGame (https://core.telegram.org/bots/api#inlinequeryresultgame)
// InlineQueryResultDocument (https://core.telegram.org/bots/api#inlinequeryresultdocument)
// InlineQueryResultGif (https://core.telegram.org/bots/api#inlinequeryresultgif)
// InlineQueryResultLocation (https://core.telegram.org/bots/api#inlinequeryresultlocation)
// InlineQueryResultMpeg4Gif (https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif)
// InlineQueryResultPhoto (https://core.telegram.org/bots/api#inlinequeryresultphoto)
// InlineQueryResultVenue (https://core.telegram.org/bots/api#inlinequeryresultvenue)
// InlineQueryResultVideo (https://core.telegram.org/bots/api#inlinequeryresultvideo)
// InlineQueryResultVoice (https://core.telegram.org/bots/api#inlinequeryresultvoice)
// Note: All URLs passed in inline query results will be available to end users and therefore must be assumed to
// be public.
type InlineQueryResult interface {
	ResultType() string
}

// InlineQueryResult types
const (
	ResultTypeArticle  = "article"
	ResultTypePhoto    = "photo"
	ResultTypeGif      = "gif"
	ResultTypeMpeg4Gif = "mpeg4_gif"
	ResultTypeVideo    = "video"
	ResultTypeAudio    = "audio"
	ResultTypeVoice    = "voice"
	ResultTypeDocument = "document"
	ResultTypeLocation = "location"
	ResultTypeVenue    = "venue"
	ResultTypeContact  = "contact"
	ResultTypeGame     = "game"
	ResultTypeSticker  = "sticker"
)

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

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// URL - Optional. URL of the result
	URL string `json:"url,omitempty"`

	// HideURL - Optional. Pass True if you don't want the URL to be shown in the message
	HideURL bool `json:"hide_url,omitempty"`

	// Description - Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// ThumbnailURL - Optional. URL of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// ThumbnailWidth - Optional. Thumbnail width
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`

	// ThumbnailHeight - Optional. Thumbnail height
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultArticle) ResultType() string {
	return ResultTypeArticle
}

// InlineQueryResultPhoto - Represents a link to a photo. By default, this photo will be sent by the user
// with optional caption. Alternatively, you can use input_message_content to send a message with the specified
// content instead of the photo.
type InlineQueryResultPhoto struct {
	// Type - Type of the result, must be photo
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// PhotoURL - A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
	PhotoURL string `json:"photo_url"`

	// ThumbnailURL - URL of the thumbnail for the photo
	ThumbnailURL string `json:"thumbnail_url"`

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

	// ParseMode - Optional. Mode for parsing entities in the photo caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the photo
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultPhoto) ResultType() string {
	return ResultTypePhoto
}

// InlineQueryResultGif - Represents a link to an animated GIF file. By default, this animated GIF file will
// be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message
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

	// GifDuration - Optional. Duration of the GIF in seconds
	GifDuration int `json:"gif_duration,omitempty"`

	// ThumbnailURL - URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url"`

	// ThumbnailMimeType - Optional. MIME type of the thumbnail, must be one of “image/jpeg”,
	// “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
	ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`

	// Title - Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Caption - Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultGif) ResultType() string {
	return ResultTypeGif
}

// ThumbMimeType types
const (
	MimeTypeImageJpeg      = "image/jpeg"
	MimeTypeImageGif       = "image/gif"
	MimeTypeVideoMp4       = "video/mp4"
	MimeTypeTextHTML       = "text/html"
	MimeTypeApplicationPDF = "application/pdf"
	MimeTypeApplicationZip = "application/zip"
)

// InlineQueryResultMpeg4Gif - Represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can
// use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	// Type - Type of the result, must be mpeg4_gif
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// Mpeg4URL - A valid URL for the MPEG4 file. File size must not exceed 1MB
	Mpeg4URL string `json:"mpeg4_url"`

	// Mpeg4Width - Optional. Video width
	Mpeg4Width int `json:"mpeg4_width,omitempty"`

	// Mpeg4Height - Optional. Video height
	Mpeg4Height int `json:"mpeg4_height,omitempty"`

	// Mpeg4Duration - Optional. Video duration in seconds
	Mpeg4Duration int `json:"mpeg4_duration,omitempty"`

	// ThumbnailURL - URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url"`

	// ThumbnailMimeType - Optional. MIME type of the thumbnail, must be one of “image/jpeg”,
	// “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
	ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`

	// Title - Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Caption - Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the video animation
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultMpeg4Gif) ResultType() string {
	return ResultTypeMpeg4Gif
}

// InlineQueryResultVideo - Represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the video.
// If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you must replace its content
// using input_message_content.
type InlineQueryResultVideo struct {
	// Type - Type of the result, must be video
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// VideoURL - A valid URL for the embedded video player or video file
	VideoURL string `json:"video_url"`

	// MimeType - MIME type of the content of the video URL, “text/html” or “video/mp4”
	MimeType string `json:"mime_type"`

	// ThumbnailURL - URL of the thumbnail (JPEG only) for the video
	ThumbnailURL string `json:"thumbnail_url"`

	// Title - Title for the result
	Title string `json:"title"`

	// Caption - Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the video caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// VideoWidth - Optional. Video width
	VideoWidth int `json:"video_width,omitempty"`

	// VideoHeight - Optional. Video height
	VideoHeight int `json:"video_height,omitempty"`

	// VideoDuration - Optional. Video duration in seconds
	VideoDuration int `json:"video_duration,omitempty"`

	// Description - Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the video. This field is
	// required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultVideo) ResultType() string {
	return ResultTypeVideo
}

// InlineQueryResultAudio - Represents a link to an MP3 audio file. By default, this audio file will be sent
// by the user. Alternatively, you can use input_message_content to send a message with the specified content
// instead of the audio.
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

	// ParseMode - Optional. Mode for parsing entities in the audio caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Performer - Optional. Performer
	Performer string `json:"performer,omitempty"`

	// AudioDuration - Optional. Audio duration in seconds
	AudioDuration int `json:"audio_duration,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the audio
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultAudio) ResultType() string {
	return ResultTypeAudio
}

// InlineQueryResultVoice - Represents a link to a voice recording in an .OGG container encoded with OPUS. By
// default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to
// send a message with the specified content instead of the the voice message.
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

	// ParseMode - Optional. Mode for parsing entities in the voice message caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// VoiceDuration - Optional. Recording duration in seconds
	VoiceDuration int `json:"voice_duration,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the voice recording
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultVoice) ResultType() string {
	return ResultTypeVoice
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

	// ParseMode - Optional. Mode for parsing entities in the document caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// DocumentURL - A valid URL for the file
	DocumentURL string `json:"document_url"`

	// MimeType - MIME type of the content of the file, either “application/pdf” or “application/zip”
	MimeType string `json:"mime_type"`

	// Description - Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the file
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`

	// ThumbnailURL - Optional. URL of the thumbnail (JPEG only) for the file
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// ThumbnailWidth - Optional. Thumbnail width
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`

	// ThumbnailHeight - Optional. Thumbnail height
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultDocument) ResultType() string {
	return ResultTypeDocument
}

// InlineQueryResultLocation - Represents a location on a map. By default, the location will be sent by the
// user. Alternatively, you can use input_message_content to send a message with the specified content instead
// of the location.
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

	// LivePeriod - Optional. Period in seconds for which the location can be updated, should be between 60 and
	// 86400.
	LivePeriod int `json:"live_period,omitempty"`

	// Heading - Optional. For live locations, a direction in which the user is moving, in degrees. Must be
	// between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`

	// ProximityAlertRadius - Optional. For live locations, a maximum distance for proximity alerts about
	// approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the location
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`

	// ThumbnailURL - Optional. URL of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// ThumbnailWidth - Optional. Thumbnail width
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`

	// ThumbnailHeight - Optional. Thumbnail height
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultLocation) ResultType() string {
	return ResultTypeLocation
}

// InlineQueryResultVenue - Represents a venue. By default, the venue will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the
// venue.
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

	// GooglePlaceType - Optional. Google Places type of the venue. (See supported types
	// (https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType string `json:"google_place_type,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the venue
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`

	// ThumbnailURL - Optional. URL of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// ThumbnailWidth - Optional. Thumbnail width
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`

	// ThumbnailHeight - Optional. Thumbnail height
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultVenue) ResultType() string {
	return ResultTypeVenue
}

// InlineQueryResultContact - Represents a contact with a phone number. By default, this contact will be sent
// by the user. Alternatively, you can use input_message_content to send a message with the specified content
// instead of the contact.
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

	// Vcard - Optional. Additional data about the contact in the form of a vCard
	// (https://en.wikipedia.org/wiki/VCard), 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the contact
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`

	// ThumbnailURL - Optional. URL of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// ThumbnailWidth - Optional. Thumbnail width
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`

	// ThumbnailHeight - Optional. Thumbnail height
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultContact) ResultType() string {
	return ResultTypeContact
}

// InlineQueryResultGame - Represents a Game (https://core.telegram.org/bots/api#games).
type InlineQueryResultGame struct {
	// Type - Type of the result, must be game
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// GameShortName - Short name of the game
	GameShortName string `json:"game_short_name"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultGame) ResultType() string {
	return ResultTypeGame
}

// InlineQueryResultCachedPhoto - Represents a link to a photo stored on the Telegram servers. By default,
// this photo will be sent by the user with an optional caption. Alternatively, you can use
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

	// ParseMode - Optional. Mode for parsing entities in the photo caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the photo
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultCachedPhoto) ResultType() string {
	return ResultTypePhoto
}

// InlineQueryResultCachedGif - Represents a link to an animated GIF file stored on the Telegram servers. By
// default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use
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

	// ParseMode - Optional. Mode for parsing entities in the caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultCachedGif) ResultType() string {
	return ResultTypeGif
}

// InlineQueryResultCachedMpeg4Gif - Represents a link to a video animation (H.264/MPEG-4 AVC video without
// sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an
// optional caption. Alternatively, you can use input_message_content to send a message with the specified
// content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	// Type - Type of the result, must be mpeg4_gif
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// Mpeg4FileID - A valid file identifier for the MPEG4 file
	Mpeg4FileID string `json:"mpeg4_file_id"`

	// Title - Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Caption - Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the video animation
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultCachedMpeg4Gif) ResultType() string {
	return ResultTypeMpeg4Gif
}

// InlineQueryResultCachedSticker - Represents a link to a sticker stored on the Telegram servers. By
// default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	// Type - Type of the result, must be sticker
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// StickerFileID - A valid file identifier of the sticker
	StickerFileID string `json:"sticker_file_id"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the sticker
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultCachedSticker) ResultType() string {
	return ResultTypeSticker
}

// InlineQueryResultCachedDocument - Represents a link to a file stored on the Telegram servers. By default,
// this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content
// to send a message with the specified content instead of the file.
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

	// ParseMode - Optional. Mode for parsing entities in the document caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the file
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultCachedDocument) ResultType() string {
	return ResultTypeDocument
}

// InlineQueryResultCachedVideo - Represents a link to a video file stored on the Telegram servers. By
// default, this video file will be sent by the user with an optional caption. Alternatively, you can use
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

	// ParseMode - Optional. Mode for parsing entities in the video caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the video
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultCachedVideo) ResultType() string {
	return ResultTypeVideo
}

// InlineQueryResultCachedVoice - Represents a link to a voice message stored on the Telegram servers. By
// default, this voice message will be sent by the user. Alternatively, you can use input_message_content to
// send a message with the specified content instead of the voice message.
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

	// ParseMode - Optional. Mode for parsing entities in the voice message caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the voice message
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultCachedVoice) ResultType() string {
	return ResultTypeVoice
}

// InlineQueryResultCachedAudio - Represents a link to an MP3 audio file stored on the Telegram servers. By
// default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	// Type - Type of the result, must be audio
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// AudioFileID - A valid file identifier for the audio file
	AudioFileID string `json:"audio_file_id"`

	// Caption - Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the audio caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// InputMessageContent - Optional. Content of the message to be sent instead of the audio
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ResultType returns InlineQueryResult type
func (i *InlineQueryResultCachedAudio) ResultType() string {
	return ResultTypeAudio
}

// InputMessageContent - This object represents the content of a message to be sent as a result of an inline
// query. Telegram clients currently support the following 5 types:
// InputTextMessageContent (https://core.telegram.org/bots/api#inputtextmessagecontent)
// InputLocationMessageContent (https://core.telegram.org/bots/api#inputlocationmessagecontent)
// InputVenueMessageContent (https://core.telegram.org/bots/api#inputvenuemessagecontent)
// InputContactMessageContent (https://core.telegram.org/bots/api#inputcontactmessagecontent)
// InputInvoiceMessageContent (https://core.telegram.org/bots/api#inputinvoicemessagecontent)
type InputMessageContent interface {
	ContentType() string
}

// InputMessageContent types
const (
	ContentTypeText     = "InputTextMessage"
	ContentTypeLocation = "InputLocationMessage"
	ContentTypeVenue    = "InputVenueMessage"
	ContentTypeContact  = "InputContactMessage"
	ContentTypeInvoice  = "InputInvoiceMessage"
)

// InputTextMessageContent - Represents the content (https://core.telegram.org/bots/api#inputmessagecontent)
// of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	// MessageText - Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`

	// ParseMode - Optional. Mode for parsing entities in the message text. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Entities - Optional. List of special entities that appear in message text, which can be specified instead
	// of parse_mode
	Entities []MessageEntity `json:"entities,omitempty"`

	// DisableWebPagePreview - Optional. Disables link previews for links in the sent message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

// ContentType returns InputMessageContent type
func (i *InputTextMessageContent) ContentType() string {
	return ContentTypeText
}

// InputLocationMessageContent - Represents the content
// (https://core.telegram.org/bots/api#inputmessagecontent) of a location message to be sent as the result of an
// inline query.
type InputLocationMessageContent struct {
	// Latitude - Latitude of the location in degrees
	Latitude float64 `json:"latitude"`

	// Longitude - Longitude of the location in degrees
	Longitude float64 `json:"longitude"`

	// HorizontalAccuracy - Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// LivePeriod - Optional. Period in seconds for which the location can be updated, should be between 60 and
	// 86400.
	LivePeriod int `json:"live_period,omitempty"`

	// Heading - Optional. For live locations, a direction in which the user is moving, in degrees. Must be
	// between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`

	// ProximityAlertRadius - Optional. For live locations, a maximum distance for proximity alerts about
	// approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

// ContentType returns InputMessageContent type
func (i *InputLocationMessageContent) ContentType() string {
	return ContentTypeLocation
}

// InputVenueMessageContent - Represents the content (https://core.telegram.org/bots/api#inputmessagecontent)
// of a venue message to be sent as the result of an inline query.
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

	// FoursquareType - Optional. Foursquare type of the venue, if known. (For example,
	// “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// GooglePlaceID - Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// GooglePlaceType - Optional. Google Places type of the venue. (See supported types
	// (https://developers.google.com/places/web-service/supported_types).)
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// ContentType returns InputMessageContent type
func (i *InputVenueMessageContent) ContentType() string {
	return ContentTypeVenue
}

// InputContactMessageContent - Represents the content
// (https://core.telegram.org/bots/api#inputmessagecontent) of a contact message to be sent as the result of an
// inline query.
type InputContactMessageContent struct {
	// PhoneNumber - Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// FirstName - Contact's first name
	FirstName string `json:"first_name"`

	// LastName - Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Vcard - Optional. Additional data about the contact in the form of a vCard
	// (https://en.wikipedia.org/wiki/VCard), 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`
}

// ContentType returns InputMessageContent type
func (i *InputContactMessageContent) ContentType() string {
	return ContentTypeContact
}

// InputInvoiceMessageContent - Represents the content
// (https://core.telegram.org/bots/api#inputmessagecontent) of an invoice message to be sent as the result of an
// inline query.
type InputInvoiceMessageContent struct {
	// Title - Product name, 1-32 characters
	Title string `json:"title"`

	// Description - Product description, 1-255 characters
	Description string `json:"description"`

	// Payload - Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your
	// internal processes.
	Payload string `json:"payload"`

	// ProviderToken - Payment provider token, obtained via @BotFather (https://t.me/botfather)
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

	// SuggestedTipAmounts - Optional. A JSON-serialized array of suggested amounts of tip in the smallest units
	// of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested
	// tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`

	// ProviderData - Optional. A JSON-serialized object for data about the invoice, which will be shared with
	// the payment provider. A detailed description of the required fields should be provided by the payment
	// provider.
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

	// NeedName - Optional. Pass True if you require the user's full name to complete the order
	NeedName bool `json:"need_name,omitempty"`

	// NeedPhoneNumber - Optional. Pass True if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// NeedEmail - Optional. Pass True if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email,omitempty"`

	// NeedShippingAddress - Optional. Pass True if you require the user's shipping address to complete the
	// order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// SendPhoneNumberToProvider - Optional. Pass True if the user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// SendEmailToProvider - Optional. Pass True if the user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// IsFlexible - Optional. Pass True if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"`
}

// ContentType returns InputMessageContent type
func (i *InputInvoiceMessageContent) ContentType() string {
	return ContentTypeInvoice
}

// ChosenInlineResult - Represents a result (https://core.telegram.org/bots/api#inlinequeryresult) of an
// inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	// ResultID - The unique identifier for the result that was chosen
	ResultID string `json:"result_id"`

	// From - The user that chose the result
	From User `json:"from"`

	// Location - Optional. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`

	// InlineMessageID - Optional. Identifier of the sent inline message. Available only if there is an inline
	// keyboard (https://core.telegram.org/bots/api#inlinekeyboardmarkup) attached to the message. Will be also
	// received in callback queries (https://core.telegram.org/bots/api#callbackquery) and can be used to edit
	// (https://core.telegram.org/bots/api#updating-messages) the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Query - The query that was used to obtain the result
	Query string `json:"query"`
}

// SentWebAppMessage - Describes an inline message sent by a Web App (https://core.telegram.org/bots/webapps)
// on behalf of a user.
type SentWebAppMessage struct {
	// InlineMessageID - Optional. Identifier of the sent inline message. Available only if there is an inline
	// keyboard (https://core.telegram.org/bots/api#inlinekeyboardmarkup) attached to the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// LabeledPrice - This object represents a portion of the price for goods or services.
type LabeledPrice struct {
	// Label - Portion label
	Label string `json:"label"`

	// Amount - Price of the product in the smallest units of the currency
	// (https://core.telegram.org/bots/payments#supported-currencies) (integer, not float/double). For example, for
	// a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal
	// point for each currency (2 for the majority of currencies).
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

	// Currency - Three-letter ISO 4217 currency (https://core.telegram.org/bots/payments#supported-currencies)
	// code
	Currency string `json:"currency"`

	// TotalAmount - Total price in the smallest units of the currency (integer, not float/double). For example,
	// for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal
	// point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}

// ShippingAddress - This object represents a shipping address.
type ShippingAddress struct {
	// CountryCode - Two-letter ISO 3166-1 alpha-2 country code
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
	// Currency - Three-letter ISO 4217 currency (https://core.telegram.org/bots/payments#supported-currencies)
	// code
	Currency string `json:"currency"`

	// TotalAmount - Total price in the smallest units of the currency (integer, not float/double). For example,
	// for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal
	// point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload - Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// ShippingOptionID - Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo - Optional. Order information provided by the user
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

	// Currency - Three-letter ISO 4217 currency (https://core.telegram.org/bots/payments#supported-currencies)
	// code
	Currency string `json:"currency"`

	// TotalAmount - Total price in the smallest units of the currency (integer, not float/double). For example,
	// for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal
	// point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload - Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// ShippingOptionID - Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo - Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

// PassportData - Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Data - Array with information about documents and other Telegram Passport elements that was shared with
	// the bot
	Data []EncryptedPassportElement `json:"data"`

	// Credentials - Encrypted credentials required to decrypt the data
	Credentials EncryptedCredentials `json:"credentials"`
}

// PassportFile - This object represents a file uploaded to Telegram Passport. Currently all Telegram
// Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// FileSize - File size in bytes
	FileSize int `json:"file_size"`

	// FileDate - Unix time when the file was uploaded
	FileDate int64 `json:"file_date"`
}

// EncryptedPassportElement - Describes documents or other Telegram Passport elements shared with the bot by
// the user.
type EncryptedPassportElement struct {
	// Type - Element type. One of “personal_details”, “passport”, “driver_license”,
	// “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”,
	// “email”.
	Type string `json:"type"`

	// Data - Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available
	// for “personal_details”, “passport”, “driver_license”, “identity_card”,
	// “internal_passport” and “address” types. Can be decrypted and verified using the accompanying
	// EncryptedCredentials (https://core.telegram.org/bots/api#encryptedcredentials).
	Data string `json:"data,omitempty"`

	// PhoneNumber - Optional. User's verified phone number, available only for “phone_number” type
	PhoneNumber string `json:"phone_number,omitempty"`

	// Email - Optional. User's verified email address, available only for “email” type
	Email string `json:"email,omitempty"`

	// Files - Optional. Array of encrypted files with documents provided by the user, available for
	// “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and
	// “temporary_registration” types. Files can be decrypted and verified using the accompanying
	// EncryptedCredentials (https://core.telegram.org/bots/api#encryptedcredentials).
	Files []PassportFile `json:"files,omitempty"`

	// FrontSide - Optional. Encrypted file with the front side of the document, provided by the user. Available
	// for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be
	// decrypted and verified using the accompanying EncryptedCredentials
	// (https://core.telegram.org/bots/api#encryptedcredentials).
	FrontSide *PassportFile `json:"front_side,omitempty"`

	// ReverseSide - Optional. Encrypted file with the reverse side of the document, provided by the user.
	// Available for “driver_license” and “identity_card”. The file can be decrypted and verified using the
	// accompanying EncryptedCredentials (https://core.telegram.org/bots/api#encryptedcredentials).
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`

	// Selfie - Optional. Encrypted file with the selfie of the user holding a document, provided by the user;
	// available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file
	// can be decrypted and verified using the accompanying EncryptedCredentials
	// (https://core.telegram.org/bots/api#encryptedcredentials).
	Selfie *PassportFile `json:"selfie,omitempty"`

	// Translation - Optional. Array of encrypted files with translated versions of documents provided by the
	// user. Available if requested for “passport”, “driver_license”, “identity_card”,
	// “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”,
	// “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using
	// the accompanying EncryptedCredentials (https://core.telegram.org/bots/api#encryptedcredentials).
	Translation []PassportFile `json:"translation,omitempty"`

	// Hash - Base64-encoded element hash for using in PassportElementErrorUnspecified
	// (https://core.telegram.org/bots/api#passportelementerrorunspecified)
	Hash string `json:"hash"`
}

// EncryptedPassportElement types
const (
	ElementTypePersonalDetails       = "personal_details"
	ElementTypePassport              = "passport"
	ElementTypeDriverLicense         = "driver_license"
	ElementTypeIdentityCard          = "identity_card"
	ElementTypeInternalPassport      = "internal_passport"
	ElementTypeAddress               = "address"
	ElementTypeUtilityBill           = "utility_bill"
	ElementTypeBankStatement         = "bank_statement"
	ElementTypeRentalAgreement       = "rental_agreement"
	ElementTypePassportRegistration  = "passport_registration"
	ElementTypeTemporaryRegistration = "temporary_registration"
	ElementTypePhoneNumber           = "phone_number"
	ElementTypeEmail                 = "email"
)

// EncryptedCredentials - Describes data required for decrypting and authenticating EncryptedPassportElement
// (https://core.telegram.org/bots/api#encryptedpassportelement). See the Telegram Passport Documentation
// (https://core.telegram.org/passport#receiving-information) for a complete description of the data decryption
// and authentication processes.
type EncryptedCredentials struct {
	// Data - Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets
	// required for EncryptedPassportElement (https://core.telegram.org/bots/api#encryptedpassportelement)
	// decryption and authentication
	Data string `json:"data"`

	// Hash - Base64-encoded data hash for data authentication
	Hash string `json:"hash"`

	// Secret - Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
	Secret string `json:"secret"`
}

// PassportElementError - This object represents an error in the Telegram Passport element which was
// submitted that should be resolved by the user. It should be one of:
// PassportElementErrorDataField (https://core.telegram.org/bots/api#passportelementerrordatafield)
// PassportElementErrorFrontSide (https://core.telegram.org/bots/api#passportelementerrorfrontside)
// PassportElementErrorReverseSide (https://core.telegram.org/bots/api#passportelementerrorreverseside)
// PassportElementErrorSelfie (https://core.telegram.org/bots/api#passportelementerrorselfie)
// PassportElementErrorFile (https://core.telegram.org/bots/api#passportelementerrorfile)
// PassportElementErrorFiles (https://core.telegram.org/bots/api#passportelementerrorfiles)
// PassportElementErrorTranslationFile (https://core.telegram.org/bots/api#passportelementerrortranslationfile)
// PassportElementErrorTranslationFiles
// (https://core.telegram.org/bots/api#passportelementerrortranslationfiles)
// PassportElementErrorUnspecified (https://core.telegram.org/bots/api#passportelementerrorunspecified)
type PassportElementError interface {
	ErrorSource() string
}

// PassportElementError sources
const (
	ErrorSourceDataField        = "data"
	ErrorSourceFrontSide        = "front_side"
	ErrorSourceReverseSide      = "reverse_side"
	ErrorSourceSelfie           = "selfie"
	ErrorSourceFile             = "file"
	ErrorSourceFiles            = "files"
	ErrorSourceTranslationFile  = "translation_file"
	ErrorSourceTranslationFiles = "translation_files"
	ErrorSourceUnspecified      = "unspecified"
)

// PassportElementErrorDataField - Represents an issue in one of the data fields that was provided by the
// user. The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	// Source - Error source, must be data
	Source string `json:"source"`

	// Type - The section of the user's Telegram Passport which has the error, one of “personal_details”,
	// “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
	Type string `json:"type"`

	// FieldName - Name of the data field which has the error
	FieldName string `json:"field_name"`

	// DataHash - Base64-encoded data hash
	DataHash string `json:"data_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// ErrorSource returns PassportElementError source
func (p *PassportElementErrorDataField) ErrorSource() string {
	return ErrorSourceDataField
}

// PassportElementErrorFrontSide - Represents an issue with the front side of a document. The error is
// considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	// Source - Error source, must be front_side
	Source string `json:"source"`

	// Type - The section of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`

	// FileHash - Base64-encoded hash of the file with the front side of the document
	FileHash string `json:"file_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// ErrorSource returns PassportElementError source
func (p *PassportElementErrorFrontSide) ErrorSource() string {
	return ErrorSourceFrontSide
}

// PassportElementErrorReverseSide - Represents an issue with the reverse side of a document. The error is
// considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	// Source - Error source, must be reverse_side
	Source string `json:"source"`

	// Type - The section of the user's Telegram Passport which has the issue, one of “driver_license”,
	// “identity_card”
	Type string `json:"type"`

	// FileHash - Base64-encoded hash of the file with the reverse side of the document
	FileHash string `json:"file_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// ErrorSource returns PassportElementError source
func (p *PassportElementErrorReverseSide) ErrorSource() string {
	return ErrorSourceReverseSide
}

// PassportElementErrorSelfie - Represents an issue with the selfie with a document. The error is considered
// resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	// Source - Error source, must be selfie
	Source string `json:"source"`

	// Type - The section of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`

	// FileHash - Base64-encoded hash of the file with the selfie
	FileHash string `json:"file_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// ErrorSource returns PassportElementError source
func (p *PassportElementErrorSelfie) ErrorSource() string {
	return ErrorSourceSelfie
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

// ErrorSource returns PassportElementError source
func (p *PassportElementErrorFile) ErrorSource() string {
	return ErrorSourceFile
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

// ErrorSource returns PassportElementError source
func (p *PassportElementErrorFiles) ErrorSource() string {
	return ErrorSourceFiles
}

// PassportElementErrorTranslationFile - Represents an issue with one of the files that constitute the
// translation of a document. The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	// Source - Error source, must be translation_file
	Source string `json:"source"`

	// Type - Type of element of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// FileHash - Base64-encoded file hash
	FileHash string `json:"file_hash"`

	// Message - Error message
	Message string `json:"message"`
}

// ErrorSource returns PassportElementError source
func (p *PassportElementErrorTranslationFile) ErrorSource() string {
	return ErrorSourceTranslationFile
}

// PassportElementErrorTranslationFiles - Represents an issue with the translated version of a document. The
// error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	// Source - Error source, must be translation_files
	Source string `json:"source"`

	// Type - Type of element of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// FileHashes - List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`

	// Message - Error message
	Message string `json:"message"`
}

// ErrorSource returns PassportElementError source
func (p *PassportElementErrorTranslationFiles) ErrorSource() string {
	return ErrorSourceTranslationFiles
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

// ErrorSource returns PassportElementError source
func (p *PassportElementErrorUnspecified) ErrorSource() string {
	return ErrorSourceUnspecified
}

// Game - This object represents a game. Use BotFather to create and edit games, their short names will act
// as unique identifiers.
type Game struct {
	// Title - Title of the game
	Title string `json:"title"`

	// Description - Description of the game
	Description string `json:"description"`

	// Photo - Photo that will be displayed in the game message in chats.
	Photo []PhotoSize `json:"photo"`

	// Text - Optional. Brief description of the game or high scores included in the game message. Can be
	// automatically edited to include current high scores for the game when the bot calls setGameScore
	// (https://core.telegram.org/bots/api#setgamescore), or manually edited using editMessageText
	// (https://core.telegram.org/bots/api#editmessagetext). 0-4096 characters.
	Text string `json:"text,omitempty"`

	// TextEntities - Optional. Special entities that appear in text, such as usernames, URLs, bot commands,
	// etc.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`

	// Animation - Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
	// (https://t.me/botfather)
	Animation *Animation `json:"animation,omitempty"`
}

// CallbackGame - A placeholder, currently holds no information. Use BotFather (https://t.me/botfather) to
// set up your game.
type CallbackGame struct{}

// GameHighScore - This object represents one row of the high scores table for a game.
type GameHighScore struct {
	// Position - Position in high score table for the game
	Position int `json:"position"`

	// User - User
	User User `json:"user"`

	// Score - Score
	Score int `json:"score"`
}
