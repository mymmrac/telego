package telego

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/mymmrac/telego/internal/json"
	"github.com/mymmrac/telego/telegoapi"
)

// Update - This object (https://core.telegram.org/bots/api#available-types) represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	// UpdateID - The update's unique identifier. Update identifiers start from a certain positive number and
	// increase sequentially. This identifier becomes especially handy if you're using webhooks
	// (https://core.telegram.org/bots/api#setwebhook), since it allows you to ignore repeated updates or to restore
	// the correct update sequence, should they get out of order. If there are no new updates for at least a week,
	// then identifier of the next update will be chosen randomly instead of sequentially.
	UpdateID int `json:"update_id"`

	// Message - Optional. New incoming message of any kind - text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`

	// EditedMessage - Optional. New version of a message that is known to the bot and was edited. This update
	// may at times be triggered by changes to message fields that are either unavailable or not actively used by
	// your bot.
	EditedMessage *Message `json:"edited_message,omitempty"`

	// ChannelPost - Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`

	// EditedChannelPost - Optional. New version of a channel post that is known to the bot and was edited. This
	// update may at times be triggered by changes to message fields that are either unavailable or not actively
	// used by your bot.
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// BusinessConnection - Optional. The bot was connected to or disconnected from a business account, or a
	// user edited an existing connection with the bot
	BusinessConnection *BusinessConnection `json:"business_connection,omitempty"`

	// BusinessMessage - Optional. New message from a connected business account
	BusinessMessage *Message `json:"business_message,omitempty"`

	// EditedBusinessMessage - Optional. New version of a message from a connected business account
	EditedBusinessMessage *Message `json:"edited_business_message,omitempty"`

	// DeletedBusinessMessages - Optional. Messages were deleted from a connected business account
	DeletedBusinessMessages *BusinessMessagesDeleted `json:"deleted_business_messages,omitempty"`

	// MessageReaction - Optional. A reaction to a message was changed by a user. The bot must be an
	// administrator in the chat and must explicitly specify "message_reaction" in the list of allowed_updates to
	// receive these updates. The update isn't received for reactions set by bots.
	MessageReaction *MessageReactionUpdated `json:"message_reaction,omitempty"`

	// MessageReactionCount - Optional. Reactions to a message with anonymous reactions were changed. The bot
	// must be an administrator in the chat and must explicitly specify "message_reaction_count" in the list of
	// allowed_updates to receive these updates. The updates are grouped and can be sent with delay up to a few
	// minutes.
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`

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

	// PurchasedPaidMedia - Optional. A user purchased paid media with a non-empty payload sent by the bot in a
	// non-channel chat
	PurchasedPaidMedia *PaidMediaPurchased `json:"purchased_paid_media,omitempty"`

	// Poll - Optional. New poll state. Bots receive only updates about manually stopped polls and polls, which
	// are sent by the bot
	Poll *Poll `json:"poll,omitempty"`

	// PollAnswer - Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only
	// in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`

	// MyChatMember - Optional. The bot's chat member status was updated in a chat. For private chats, this
	// update is received only when the bot is blocked or unblocked by the user.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`

	// ChatMember - Optional. A chat member's status was updated in a chat. The bot must be an administrator in
	// the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`

	// ChatJoinRequest - Optional. A request to join the chat has been sent. The bot must have the
	// can_invite_users administrator right in the chat to receive these updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`

	// ChatBoost - Optional. A chat boost was added or changed. The bot must be an administrator in the chat to
	// receive these updates.
	ChatBoost *ChatBoostUpdated `json:"chat_boost,omitempty"`

	// RemovedChatBoost - Optional. A boost was removed from a chat. The bot must be an administrator in the
	// chat to receive these updates.
	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost,omitempty"`

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
	data, err := json.Marshal(u)
	if err != nil {
		return Update{}, fmt.Errorf("telego: clone update: marshal: %w", err)
	}

	var update Update
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

	// CanConnectToBusiness - Optional. True, if the bot can be connected to a Telegram Business account to
	// receive its messages. Returned only in getMe (https://core.telegram.org/bots/api#getme).
	CanConnectToBusiness bool `json:"can_connect_to_business,omitempty"`

	// HasMainWebApp - Optional. True, if the bot has a main Web App. Returned only in getMe
	// (https://core.telegram.org/bots/api#getme).
	HasMainWebApp bool `json:"has_main_web_app,omitempty"`

	// HasTopicsEnabled - Optional. True, if the bot has forum topic mode enabled in private chats. Returned
	// only in getMe (https://core.telegram.org/bots/api#getme).
	HasTopicsEnabled bool `json:"has_topics_enabled,omitempty"`

	// AllowsUsersToCreateTopics - Optional. True, if the bot allows users to create and delete topics in
	// private chats. Returned only in getMe (https://core.telegram.org/bots/api#getme).
	AllowsUsersToCreateTopics bool `json:"allows_users_to_create_topics,omitempty"`
}

// Chat - This object represents a chat.
type Chat struct {
	// ID - Unique identifier for this chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// identifier.
	ID int64 `json:"id"`

	// Type - Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
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

	// IsDirectMessages - Optional. True, if the chat is the direct messages chat of a channel
	IsDirectMessages bool `json:"is_direct_messages,omitempty"`
}

// ChatID returns [ChatID] of this chat
func (c Chat) ChatID() ChatID {
	return ChatID{
		ID: c.ID,
	}
}

// Chat types
const (
	ChatTypeSender     = "sender"
	ChatTypePrivate    = "private"
	ChatTypeGroup      = "group"
	ChatTypeSupergroup = "supergroup"
	ChatTypeChannel    = "channel"
)

// ChatFullInfo - This object contains full information about a chat.
type ChatFullInfo struct {
	// ID - Unique identifier for this chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// identifier.
	ID int64 `json:"id"`

	// Type - Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
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

	// IsDirectMessages - Optional. True, if the chat is the direct messages chat of a channel
	IsDirectMessages bool `json:"is_direct_messages,omitempty"`

	// AccentColorID - Identifier of the accent color for the chat name and backgrounds of the chat photo, reply
	// header, and link preview. See accent colors (https://core.telegram.org/bots/api#accent-colors) for more
	// details.
	AccentColorID int `json:"accent_color_id"`

	// MaxReactionCount - The maximum number of reactions that can be set on a message in the chat
	MaxReactionCount int `json:"max_reaction_count"`

	// Photo - Optional. Chat photo
	Photo *ChatPhoto `json:"photo,omitempty"`

	// ActiveUsernames - Optional. If non-empty, the list of all active chat usernames
	// (https://telegram.org/blog/topics-in-groups-collectible-usernames#collectible-usernames); for private chats,
	// supergroups and channels
	ActiveUsernames []string `json:"active_usernames,omitempty"`

	// Birthdate - Optional. For private chats, the date of birth of the user
	Birthdate *Birthdate `json:"birthdate,omitempty"`

	// BusinessIntro - Optional. For private chats with business accounts, the intro of the business
	BusinessIntro *BusinessIntro `json:"business_intro,omitempty"`

	// BusinessLocation - Optional. For private chats with business accounts, the location of the business
	BusinessLocation *BusinessLocation `json:"business_location,omitempty"`

	// BusinessOpeningHours - Optional. For private chats with business accounts, the opening hours of the
	// business
	BusinessOpeningHours *BusinessOpeningHours `json:"business_opening_hours,omitempty"`

	// PersonalChat - Optional. For private chats, the personal channel of the user
	PersonalChat *Chat `json:"personal_chat,omitempty"`

	// ParentChat - Optional. Information about the corresponding channel chat; for direct messages chats only
	ParentChat *Chat `json:"parent_chat,omitempty"`

	// AvailableReactions - Optional. List of available reactions allowed in the chat. If omitted, then all
	// emoji reactions (https://core.telegram.org/bots/api#reactiontypeemoji) are allowed.
	AvailableReactions []ReactionType `json:"available_reactions,omitempty"`

	// BackgroundCustomEmojiID - Optional. Custom emoji identifier of the emoji chosen by the chat for the reply
	// header and link preview background
	BackgroundCustomEmojiID string `json:"background_custom_emoji_id,omitempty"`

	// ProfileAccentColorID - Optional. Identifier of the accent color for the chat's profile background. See
	// profile accent colors (https://core.telegram.org/bots/api#profile-accent-colors) for more details.
	ProfileAccentColorID int `json:"profile_accent_color_id,omitempty"`

	// ProfileBackgroundCustomEmojiID - Optional. Custom emoji identifier of the emoji chosen by the chat for
	// its profile background
	ProfileBackgroundCustomEmojiID string `json:"profile_background_custom_emoji_id,omitempty"`

	// EmojiStatusCustomEmojiID - Optional. Custom emoji identifier of the emoji status of the chat or the other
	// party in a private chat
	EmojiStatusCustomEmojiID string `json:"emoji_status_custom_emoji_id,omitempty"`

	// EmojiStatusExpirationDate - Optional. Expiration date of the emoji status of the chat or the other party
	// in a private chat, in Unix time, if any
	EmojiStatusExpirationDate int64 `json:"emoji_status_expiration_date,omitempty"`

	// Bio - Optional. Bio of the other party in a private chat
	Bio string `json:"bio,omitempty"`

	// HasPrivateForwards - Optional. True, if privacy settings of the other party in the private chat allows to
	// use tg://user?id=<user_id> links only in chats with the user
	HasPrivateForwards bool `json:"has_private_forwards,omitempty"`

	// HasRestrictedVoiceAndVideoMessages - Optional. True, if the privacy settings of the other party restrict
	// sending voice and video note messages in the private chat
	HasRestrictedVoiceAndVideoMessages bool `json:"has_restricted_voice_and_video_messages,omitempty"`

	// JoinToSendMessages - Optional. True, if users need to join the supergroup before they can send messages
	JoinToSendMessages bool `json:"join_to_send_messages,omitempty"`

	// JoinByRequest - Optional. True, if all users directly joining the supergroup without using an invite link
	// need to be approved by supergroup administrators
	JoinByRequest bool `json:"join_by_request,omitempty"`

	// Description - Optional. Description, for groups, supergroups and channel chats
	Description string `json:"description,omitempty"`

	// InviteLink - Optional. Primary invite link, for groups, supergroups and channel chats
	InviteLink string `json:"invite_link,omitempty"`

	// PinnedMessage - Optional. The most recent pinned message (by sending date)
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Permissions - Optional. Default chat member permissions, for groups and supergroups
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// AcceptedGiftTypes - Information about types of gifts that are accepted by the chat or by the
	// corresponding user for private chats
	AcceptedGiftTypes AcceptedGiftTypes `json:"accepted_gift_types"`

	// CanSendPaidMedia - Optional. True, if paid media messages can be sent or forwarded to the channel chat.
	// The field is available only for channel chats.
	CanSendPaidMedia bool `json:"can_send_paid_media,omitempty"`

	// SlowModeDelay - Optional. For supergroups, the minimum allowed delay between consecutive messages sent by
	// each unprivileged user; in seconds
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`

	// UnrestrictBoostCount - Optional. For supergroups, the minimum number of boosts that a non-administrator
	// user needs to add in order to ignore slow mode and chat permissions
	UnrestrictBoostCount int `json:"unrestrict_boost_count,omitempty"`

	// MessageAutoDeleteTime - Optional. The time after which all messages sent to the chat will be
	// automatically deleted; in seconds
	MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`

	// HasAggressiveAntiSpamEnabled - Optional. True, if aggressive anti-spam checks are enabled in the
	// supergroup. The field is only available to chat administrators.
	HasAggressiveAntiSpamEnabled bool `json:"has_aggressive_anti_spam_enabled,omitempty"`

	// HasHiddenMembers - Optional. True, if non-administrators can only get the list of bots and administrators
	// in the chat
	HasHiddenMembers bool `json:"has_hidden_members,omitempty"`

	// HasProtectedContent - Optional. True, if messages from the chat can't be forwarded to other chats
	HasProtectedContent bool `json:"has_protected_content,omitempty"`

	// HasVisibleHistory - Optional. True, if new chat members will have access to old messages; available only
	// to chat administrators
	HasVisibleHistory bool `json:"has_visible_history,omitempty"`

	// StickerSetName - Optional. For supergroups, name of the group sticker set
	StickerSetName string `json:"sticker_set_name,omitempty"`

	// CanSetStickerSet - Optional. True, if the bot can change the group sticker set
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`

	// CustomEmojiStickerSetName - Optional. For supergroups, the name of the group's custom emoji sticker set.
	// Custom emoji from this set can be used by all users and bots in the group.
	CustomEmojiStickerSetName string `json:"custom_emoji_sticker_set_name,omitempty"`

	// LinkedChatID - Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for
	// a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and
	// some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52
	// bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	LinkedChatID int64 `json:"linked_chat_id,omitempty"`

	// Location - Optional. For supergroups, the location to which the supergroup is connected
	Location *ChatLocation `json:"location,omitempty"`

	// Rating - Optional. For private chats, the rating of the user if any
	Rating *UserRating `json:"rating,omitempty"`

	// FirstProfileAudio - Optional. For private chats, the first audio added to the profile of the user
	FirstProfileAudio *Audio `json:"first_profile_audio,omitempty"`

	// UniqueGiftColors - Optional. The color scheme based on a unique gift that must be used for the chat's
	// name, message replies and link previews
	UniqueGiftColors *UniqueGiftColors `json:"unique_gift_colors,omitempty"`

	// PaidMessageStarCount - Optional. The number of Telegram Stars a general user have to pay to send a
	// message to the chat
	PaidMessageStarCount int `json:"paid_message_star_count,omitempty"`
}

// unknownReactionTypeErr is an error for unknown reaction type
const unknownReactionTypeErr = "unknown reaction type: %q"

// UnmarshalJSON converts JSON to Chat
func (c *ChatFullInfo) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	type uChatFullInfo ChatFullInfo
	var uc uChatFullInfo

	if value.Exists("available_reactions") {
		availableReactions := value.GetArray("available_reactions")
		uc.AvailableReactions = make([]ReactionType, len(availableReactions))
		for i, reaction := range availableReactions {
			reactionType := string(reaction.GetStringBytes("type"))
			switch reactionType {
			case ReactionEmoji:
				uc.AvailableReactions[i] = &ReactionTypeEmoji{}
			case ReactionCustomEmoji:
				uc.AvailableReactions[i] = &ReactionTypeCustomEmoji{}
			case ReactionPaid:
				uc.AvailableReactions[i] = &ReactionTypePaid{}
			default:
				return fmt.Errorf(unknownReactionTypeErr, reactionType)
			}
		}
	}

	if err = json.Unmarshal(data, &uc); err != nil {
		return err
	}
	*c = ChatFullInfo(uc)

	return nil
}

// Message - This object represents a message.
type Message struct {
	// MessageID - Unique message identifier inside this chat. In specific instances (e.g., message containing a
	// video sent to a big chat), the server might automatically schedule a message instead of sending it
	// immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is
	// actually sent
	MessageID int `json:"message_id"`

	// MessageThreadID - Optional. Unique identifier of a message thread or forum topic to which the message
	// belongs; for supergroups and private chats only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// DirectMessagesTopic - Optional. Information about the direct messages chat topic that contains the
	// message
	DirectMessagesTopic *DirectMessagesTopic `json:"direct_messages_topic,omitempty"`

	// From - Optional. Sender of the message; may be empty for messages sent to channels. For backward
	// compatibility, if the message was sent on behalf of a chat, the field contains a fake sender user in
	// non-channel chats
	From *User `json:"from,omitempty"`

	// SenderChat - Optional. Sender of the message when sent on behalf of a chat. For example, the supergroup
	// itself for messages sent by its anonymous administrators or a linked channel for messages automatically
	// forwarded to the channel's discussion group. For backward compatibility, if the message was sent on behalf of
	// a chat, the field from contains a fake sender user in non-channel chats.
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// SenderBoostCount - Optional. If the sender of the message boosted the chat, the number of boosts added by
	// the user
	SenderBoostCount int `json:"sender_boost_count,omitempty"`

	// SenderBusinessBot - Optional. The bot that actually sent the message on behalf of the business account.
	// Available only for outgoing messages sent on behalf of the connected business account.
	SenderBusinessBot *User `json:"sender_business_bot,omitempty"`

	// SenderTag - Optional. Tag or custom title of the sender of the message; for supergroups only
	SenderTag string `json:"sender_tag,omitempty"`

	// Date - Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	Date int64 `json:"date"`

	// BusinessConnectionID - Optional. Unique identifier of the business connection from which the message was
	// received. If non-empty, the message belongs to a chat of the corresponding business account that is
	// independent from any potential bot chat which might share the same identifier.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// Chat - Chat the message belongs to
	Chat Chat `json:"chat"`

	// ForwardOrigin - Optional. Information about the original message for forwarded messages
	ForwardOrigin MessageOrigin `json:"forward_origin,omitempty"`

	// IsTopicMessage - Optional. True, if the message is sent to a topic in a forum supergroup or a private
	// chat with the bot
	IsTopicMessage bool `json:"is_topic_message,omitempty"`

	// IsAutomaticForward - Optional. True, if the message is a channel post that was automatically forwarded to
	// the connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`

	// ReplyToMessage - Optional. For replies in the same chat and message thread, the original message. Note
	// that the Message (https://core.telegram.org/bots/api#message) object in this field will not contain further
	// reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// ExternalReply - Optional. Information about the message that is being replied to, which may come from
	// another chat or forum topic
	ExternalReply *ExternalReplyInfo `json:"external_reply,omitempty"`

	// Quote - Optional. For replies that quote part of the original message, the quoted part of the message
	Quote *TextQuote `json:"quote,omitempty"`

	// ReplyToStory - Optional. For replies to a story, the original story
	ReplyToStory *Story `json:"reply_to_story,omitempty"`

	// ReplyToChecklistTaskID - Optional. Identifier of the specific checklist task that is being replied to
	ReplyToChecklistTaskID int `json:"reply_to_checklist_task_id,omitempty"`

	// ViaBot - Optional. Bot through which the message was sent
	ViaBot *User `json:"via_bot,omitempty"`

	// EditDate - Optional. Date the message was last edited in Unix time
	EditDate int64 `json:"edit_date,omitempty"`

	// HasProtectedContent - Optional. True, if the message can't be forwarded
	HasProtectedContent bool `json:"has_protected_content,omitempty"`

	// IsFromOffline - Optional. True, if the message was sent by an implicit action, for example, as an away or
	// a greeting business message, or as a scheduled message
	IsFromOffline bool `json:"is_from_offline,omitempty"`

	// IsPaidPost - Optional. True, if the message is a paid post. Note that such posts must not be deleted for
	// 24 hours to receive the payment and can't be edited.
	IsPaidPost bool `json:"is_paid_post,omitempty"`

	// MediaGroupID - Optional. The unique identifier inside this chat of a media message group this message
	// belongs to
	MediaGroupID string `json:"media_group_id,omitempty"`

	// AuthorSignature - Optional. Signature of the post author for messages in channels, or the custom title of
	// an anonymous group administrator
	AuthorSignature string `json:"author_signature,omitempty"`

	// PaidStarCount - Optional. The number of Telegram Stars that were paid by the sender of the message to
	// send it
	PaidStarCount int `json:"paid_star_count,omitempty"`

	// Text - Optional. For text messages, the actual UTF-8 text of the message
	Text string `json:"text,omitempty"`

	// Entities - Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that
	// appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`

	// LinkPreviewOptions - Optional. Options used for link preview generation for the message, if it is a text
	// message and link preview options were changed
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// SuggestedPostInfo - Optional. Information about suggested post parameters if the message is a suggested
	// post in a channel direct messages chat. If the message is an approved or declined suggested post, then it
	// can't be edited.
	SuggestedPostInfo *SuggestedPostInfo `json:"suggested_post_info,omitempty"`

	// EffectID - Optional. Unique identifier of the message effect added to the message
	EffectID string `json:"effect_id,omitempty"`

	// Animation - Optional. Message is an animation, information about the animation. For backward
	// compatibility, when this field is set, the document field will also be set
	Animation *Animation `json:"animation,omitempty"`

	// Audio - Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`

	// Document - Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`

	// PaidMedia - Optional. Message contains paid media; information about the paid media
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`

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

	// Caption - Optional. Caption for the animation, audio, document, paid media, photo, video or voice
	Caption string `json:"caption,omitempty"`

	// CaptionEntities - Optional. For messages with a caption, special entities like usernames, URLs, bot
	// commands, etc. that appear in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ShowCaptionAboveMedia - Optional. True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// HasMediaSpoiler - Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Checklist - Optional. Message is a checklist
	Checklist *Checklist `json:"checklist,omitempty"`

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

	// ChatOwnerLeft - Optional. Service message: chat owner has left
	ChatOwnerLeft *ChatOwnerLeft `json:"chat_owner_left,omitempty"`

	// ChatOwnerChanged - Optional. Service message: chat owner has changed
	ChatOwnerChanged *ChatOwnerChanged `json:"chat_owner_changed,omitempty"`

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

	// PinnedMessage - Optional. Specified message was pinned. Note that the Message
	// (https://core.telegram.org/bots/api#message) object in this field will not contain further reply_to_message
	// fields even if it itself is a reply.
	PinnedMessage MaybeInaccessibleMessage `json:"pinned_message,omitempty"`

	// Invoice - Optional. Message is an invoice for a payment (https://core.telegram.org/bots/api#payments),
	// information about the invoice. More about payments » (https://core.telegram.org/bots/api#payments)
	Invoice *Invoice `json:"invoice,omitempty"`

	// SuccessfulPayment - Optional. Message is a service message about a successful payment, information about
	// the payment. More about payments » (https://core.telegram.org/bots/api#payments)
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`

	// RefundedPayment - Optional. Message is a service message about a refunded payment, information about the
	// payment. More about payments » (https://core.telegram.org/bots/api#payments)
	RefundedPayment *RefundedPayment `json:"refunded_payment,omitempty"`

	// UsersShared - Optional. Service message: users were shared with the bot
	UsersShared *UsersShared `json:"users_shared,omitempty"`

	// ChatShared - Optional. Service message: a chat was shared with the bot
	ChatShared *ChatShared `json:"chat_shared,omitempty"`

	// Gift - Optional. Service message: a regular gift was sent or received
	Gift *GiftInfo `json:"gift,omitempty"`

	// UniqueGift - Optional. Service message: a unique gift was sent or received
	UniqueGift *UniqueGiftInfo `json:"unique_gift,omitempty"`

	// GiftUpgradeSent - Optional. Service message: upgrade of a gift was purchased after the gift was sent
	GiftUpgradeSent *GiftInfo `json:"gift_upgrade_sent,omitempty"`

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

	// BoostAdded - Optional. Service message: user boosted the chat
	BoostAdded *ChatBoostAdded `json:"boost_added,omitempty"`

	// ChatBackgroundSet - Optional. Service message: chat background set
	ChatBackgroundSet *ChatBackground `json:"chat_background_set,omitempty"`

	// ChecklistTasksDone - Optional. Service message: some tasks in a checklist were marked as done or not done
	ChecklistTasksDone *ChecklistTasksDone `json:"checklist_tasks_done,omitempty"`

	// ChecklistTasksAdded - Optional. Service message: tasks were added to a checklist
	ChecklistTasksAdded *ChecklistTasksAdded `json:"checklist_tasks_added,omitempty"`

	// DirectMessagePriceChanged - Optional. Service message: the price for paid messages in the corresponding
	// direct messages chat of a channel has changed
	DirectMessagePriceChanged *DirectMessagePriceChanged `json:"direct_message_price_changed,omitempty"`

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

	// GiveawayCreated - Optional. Service message: a scheduled giveaway was created
	GiveawayCreated *GiveawayCreated `json:"giveaway_created,omitempty"`

	// Giveaway - Optional. The message is a scheduled giveaway message
	Giveaway *Giveaway `json:"giveaway,omitempty"`

	// GiveawayWinners - Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`

	// GiveawayCompleted - Optional. Service message: a giveaway without public winners was completed
	GiveawayCompleted *GiveawayCompleted `json:"giveaway_completed,omitempty"`

	// PaidMessagePriceChanged - Optional. Service message: the price for paid messages has changed in the chat
	PaidMessagePriceChanged *PaidMessagePriceChanged `json:"paid_message_price_changed,omitempty"`

	// SuggestedPostApproved - Optional. Service message: a suggested post was approved
	SuggestedPostApproved *SuggestedPostApproved `json:"suggested_post_approved,omitempty"`

	// SuggestedPostApprovalFailed - Optional. Service message: approval of a suggested post has failed
	SuggestedPostApprovalFailed *SuggestedPostApprovalFailed `json:"suggested_post_approval_failed,omitempty"`

	// SuggestedPostDeclined - Optional. Service message: a suggested post was declined
	SuggestedPostDeclined *SuggestedPostDeclined `json:"suggested_post_declined,omitempty"`

	// SuggestedPostPaid - Optional. Service message: payment for a suggested post was received
	SuggestedPostPaid *SuggestedPostPaid `json:"suggested_post_paid,omitempty"`

	// SuggestedPostRefunded - Optional. Service message: payment for a suggested post was refunded
	SuggestedPostRefunded *SuggestedPostRefunded `json:"suggested_post_refunded,omitempty"`

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

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message. login_url buttons are represented as ordinary URL buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// UnmarshalJSON converts JSON to Message
func (m *Message) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	type uMessage Message
	var um uMessage

	if value.Exists("pinned_message") {
		if value.GetInt("pinned_message", "date") == 0 {
			um.PinnedMessage = &InaccessibleMessage{}
		} else {
			um.PinnedMessage = &Message{}
		}
	}

	if value.Exists("forward_origin") {
		forwardOriginType := string(value.GetStringBytes("forward_origin", "type"))
		switch forwardOriginType {
		case OriginTypeUser:
			um.ForwardOrigin = &MessageOriginUser{}
		case OriginTypeHiddenUser:
			um.ForwardOrigin = &MessageOriginHiddenUser{}
		case OriginTypeChat:
			um.ForwardOrigin = &MessageOriginChat{}
		case OriginTypeChannel:
			um.ForwardOrigin = &MessageOriginChannel{}
		default:
			return fmt.Errorf("unknown forward message origin: %q", forwardOriginType)
		}
	}

	if err = json.Unmarshal(data, &um); err != nil {
		return err
	}
	*m = Message(um)

	return nil
}

// IsAccessible returns true if message accessible for bot
func (m *Message) IsAccessible() bool {
	return true
}

// Message returns [Message] if message accessible for bot, otherwise returns nil
func (m *Message) Message() *Message {
	return m
}

// InaccessibleMessage returns [InaccessibleMessage] if message not accessible for bot, otherwise returns nil
func (m *Message) InaccessibleMessage() *InaccessibleMessage {
	return nil
}

// GetChat returns message chat
func (m *Message) GetChat() Chat {
	return m.Chat
}

// GetMessageID returns message ID
func (m *Message) GetMessageID() int {
	return m.MessageID
}

// GetDate returns message date
func (m *Message) GetDate() int64 {
	return m.Date
}

func (m *Message) iMaybeInaccessibleMessage() {}

// MessageID - This object represents a unique message identifier.
type MessageID struct {
	// MessageID - Unique message identifier. In specific instances (e.g., message containing a video sent to a
	// big chat), the server might automatically schedule a message instead of sending it immediately. In such
	// cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageID int `json:"message_id"`
}

// InaccessibleMessage - This object describes a message that was deleted or is otherwise inaccessible to the
// bot.
type InaccessibleMessage struct {
	// Chat - Chat the message belonged to
	Chat Chat `json:"chat"`

	// MessageID - Unique message identifier inside the chat
	MessageID int `json:"message_id"`

	// Date - Always 0. The field can be used to differentiate regular and inaccessible messages.
	Date int64 `json:"date"`
}

// IsAccessible returns true if message accessible for bot
func (m *InaccessibleMessage) IsAccessible() bool {
	return false
}

// Message returns [Message] if message accessible for bot, otherwise returns nil
func (m *InaccessibleMessage) Message() *Message {
	return nil
}

// InaccessibleMessage returns [InaccessibleMessage] if message not accessible for bot, otherwise returns nil
func (m *InaccessibleMessage) InaccessibleMessage() *InaccessibleMessage {
	return m
}

// GetChat returns message chat
func (m *InaccessibleMessage) GetChat() Chat {
	return m.Chat
}

// GetMessageID returns message ID
func (m *InaccessibleMessage) GetMessageID() int {
	return m.MessageID
}

// GetDate returns message date
func (m *InaccessibleMessage) GetDate() int64 {
	return m.Date
}

func (m *InaccessibleMessage) iMaybeInaccessibleMessage() {}

// MaybeInaccessibleMessage - This object describes a message that can be inaccessible to the bot. It can be
// one of
// Message (https://core.telegram.org/bots/api#message)
// InaccessibleMessage (https://core.telegram.org/bots/api#inaccessiblemessage)
type MaybeInaccessibleMessage interface {
	// IsAccessible returns true if message accessible for bot
	IsAccessible() bool
	// Message returns [Message] if message accessible for bot, otherwise returns nil
	Message() *Message
	// InaccessibleMessage returns [InaccessibleMessage] if message not accessible for bot, otherwise returns nil
	InaccessibleMessage() *InaccessibleMessage

	// GetChat returns message chat
	GetChat() Chat
	// GetMessageID returns message ID
	GetMessageID() int
	// GetDate returns message date
	GetDate() int64

	// Disallow external implementations
	iMaybeInaccessibleMessage()
}

// MessageEntity - This object represents one special entity in a text message. For example, hashtags,
// usernames, URLs, etc.
type MessageEntity struct {
	// Type - Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag or
	// #hashtag@chatusername), “cashtag” ($USD or $USD@chatusername), “bot_command” (/start@jobs_bot),
	// “url” (https://telegram.org), “email” (do-not-reply@telegram.org), “phone_number”
	// (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text),
	// “strikethrough” (strikethrough text), “spoiler” (spoiler message), “blockquote” (block
	// quotation), “expandable_blockquote” (collapsed-by-default block quotation), “code” (monowidth
	// string), “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention” (for
	// users without usernames (https://telegram.org/blog/edit#new-mentions)), “custom_emoji” (for inline custom
	// emoji stickers), or “date_time” (for formatted date and time)
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

	// UnixTime - Optional. For “date_time” only, the Unix time associated with the entity
	UnixTime int64 `json:"unix_time,omitempty"`

	// DateTimeFormat - Optional. For “date_time” only, the string that defines the formatting of the date
	// and time. See date-time entity formatting (https://core.telegram.org/bots/api#date-time-entity-formatting)
	// for more details.
	DateTimeFormat string `json:"date_time_format,omitempty"`
}

// MessageEntity types
const (
	EntityTypeMention              = "mention"
	EntityTypeHashtag              = "hashtag"
	EntityTypeCashtag              = "cashtag"
	EntityTypeBotCommand           = "bot_command"
	EntityTypeURL                  = "url"
	EntityTypeEmail                = "email"
	EntityTypePhoneNumber          = "phone_number"
	EntityTypeBold                 = "bold"
	EntityTypeItalic               = "italic"
	EntityTypeUnderline            = "underline"
	EntityTypeStrikethrough        = "strikethrough"
	EntityTypeSpoiler              = "spoiler"
	EntityTypeBlockquote           = "blockquote"
	EntityTypeExpandableBlockquote = "expandable_blockquote"
	EntityTypeCode                 = "code"
	EntityTypePre                  = "pre"
	EntityTypeTextLink             = "text_link"
	EntityTypeTextMention          = "text_mention"
	EntityTypeCustomEmoji          = "custom_emoji"
	EntityTypeDateTime             = "date_time"
)

// TextQuote - This object contains information about the quoted part of a message that is replied to by the
// given message.
type TextQuote struct {
	// Text - Text of the quoted part of a message that is replied to by the given message
	Text string `json:"text"`

	// Entities - Optional. Special entities that appear in the quote. Currently, only bold, italic, underline,
	// strikethrough, spoiler, and custom_emoji entities are kept in quotes.
	Entities []MessageEntity `json:"entities,omitempty"`

	// Position - Approximate quote position in the original message in UTF-16 code units as specified by the
	// sender
	Position int `json:"position"`

	// IsManual - Optional. True, if the quote was chosen manually by the message sender. Otherwise, the quote
	// was added automatically by the server.
	IsManual bool `json:"is_manual,omitempty"`
}

// ExternalReplyInfo - This object contains information about a message that is being replied to, which may
// come from another chat or forum topic.
type ExternalReplyInfo struct {
	// Origin - Origin of the message replied to by the given message
	Origin MessageOrigin `json:"origin"`

	// Chat - Optional. Chat the original message belongs to. Available only if the chat is a supergroup or a
	// channel.
	Chat *Chat `json:"chat,omitempty"`

	// MessageID - Optional. Unique message identifier inside the original chat. Available only if the original
	// chat is a supergroup or a channel.
	MessageID int `json:"message_id,omitempty"`

	// LinkPreviewOptions - Optional. Options used for link preview generation for the original message, if it
	// is a text message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// Animation - Optional. Message is an animation, information about the animation
	Animation *Animation `json:"animation,omitempty"`

	// Audio - Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`

	// Document - Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`

	// PaidMedia - Optional. Message contains paid media; information about the paid media
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`

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

	// HasMediaSpoiler - Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Checklist - Optional. Message is a checklist
	Checklist *Checklist `json:"checklist,omitempty"`

	// Contact - Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`

	// Dice - Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`

	// Game - Optional. Message is a game, information about the game. More about games »
	// (https://core.telegram.org/bots/api#games)
	Game *Game `json:"game,omitempty"`

	// Giveaway - Optional. Message is a scheduled giveaway, information about the giveaway
	Giveaway *Giveaway `json:"giveaway,omitempty"`

	// GiveawayWinners - Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`

	// Invoice - Optional. Message is an invoice for a payment (https://core.telegram.org/bots/api#payments),
	// information about the invoice. More about payments » (https://core.telegram.org/bots/api#payments)
	Invoice *Invoice `json:"invoice,omitempty"`

	// Location - Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`

	// Poll - Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`

	// Venue - Optional. Message is a venue, information about the venue
	Venue *Venue `json:"venue,omitempty"`
}

// UnmarshalJSON converts JSON to ExternalReplyInfo
func (e *ExternalReplyInfo) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	if !value.Exists("origin") {
		return errors.New("no origin")
	}

	type uExternalReplyInfo ExternalReplyInfo
	var ue uExternalReplyInfo

	originType := string(value.GetStringBytes("origin", "type"))
	switch originType {
	case OriginTypeUser:
		ue.Origin = &MessageOriginUser{}
	case OriginTypeHiddenUser:
		ue.Origin = &MessageOriginHiddenUser{}
	case OriginTypeChat:
		ue.Origin = &MessageOriginChat{}
	case OriginTypeChannel:
		ue.Origin = &MessageOriginChannel{}
	default:
		return fmt.Errorf("unknown origin: %q", originType)
	}

	if err = json.Unmarshal(data, &ue); err != nil {
		return err
	}
	*e = ExternalReplyInfo(ue)

	return nil
}

// ReplyParameters - Describes reply parameters for the message that is being sent.
type ReplyParameters struct {
	// MessageID - Identifier of the message that will be replied to in the current chat, or in the chat chat_id
	// if it is specified
	MessageID int `json:"message_id"`

	// ChatID - Optional. If the message to be replied to is from a different chat, unique identifier for the
	// chat or username of the channel (in the format @channel_username). Not supported for messages sent on behalf
	// of a business account and messages from channel direct messages chats.
	ChatID ChatID `json:"chat_id,omitzero"`

	// AllowSendingWithoutReply - Optional. Pass True if the message should be sent even if the specified
	// message to be replied to is not found. Always False for replies in another chat or forum topic. Always True
	// for messages sent on behalf of a business account.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// Quote - Optional. Quoted part of the message to be replied to; 0-1024 characters after entities parsing.
	// The quote must be an exact substring of the message to be replied to, including bold, italic, underline,
	// strikethrough, spoiler, and custom_emoji entities. The message will fail to send if the quote isn't found in
	// the original message.
	Quote string `json:"quote,omitempty"`

	// QuoteParseMode - Optional. Mode for parsing entities in the quote. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	QuoteParseMode string `json:"quote_parse_mode,omitempty"`

	// QuoteEntities - Optional. A JSON-serialized list of special entities that appear in the quote. It can be
	// specified instead of quote_parse_mode.
	QuoteEntities []MessageEntity `json:"quote_entities,omitempty"`

	// QuotePosition - Optional. Position of the quote in the original message in UTF-16 code units
	QuotePosition int `json:"quote_position,omitempty"`

	// ChecklistTaskID - Optional. Identifier of the specific checklist task to be replied to
	ChecklistTaskID int `json:"checklist_task_id,omitempty"`
}

// MessageOrigin - This object describes the origin of a message. It can be one of
// MessageOriginUser (https://core.telegram.org/bots/api#messageoriginuser)
// MessageOriginHiddenUser (https://core.telegram.org/bots/api#messageoriginhiddenuser)
// MessageOriginChat (https://core.telegram.org/bots/api#messageoriginchat)
// MessageOriginChannel (https://core.telegram.org/bots/api#messageoriginchannel)
type MessageOrigin interface {
	// OriginType returns original message type
	OriginType() string
	// OriginalDate returns original message date
	OriginalDate() int64
	// Disallow external implementations
	iMessageOrigin()
}

// Message origin types
const (
	OriginTypeUser       = "user"
	OriginTypeHiddenUser = "hidden_user"
	OriginTypeChat       = "chat"
	OriginTypeChannel    = "channel"
)

// MessageOriginUser - The message was originally sent by a known user.
type MessageOriginUser struct {
	// Type - Type of the message origin, always “user”
	Type string `json:"type"`

	// Date - Date the message was sent originally in Unix time
	Date int64 `json:"date"`

	// SenderUser - User that sent the message originally
	SenderUser User `json:"sender_user"`
}

// OriginType returns original message type
func (m *MessageOriginUser) OriginType() string {
	return OriginTypeUser
}

// OriginalDate returns original message date
func (m *MessageOriginUser) OriginalDate() int64 {
	return m.Date
}

func (m *MessageOriginUser) iMessageOrigin() {}

// MessageOriginHiddenUser - The message was originally sent by an unknown user.
type MessageOriginHiddenUser struct {
	// Type - Type of the message origin, always “hidden_user”
	Type string `json:"type"`

	// Date - Date the message was sent originally in Unix time
	Date int64 `json:"date"`

	// SenderUserName - Name of the user that sent the message originally
	SenderUserName string `json:"sender_user_name"`
}

// OriginType returns original message type
func (m *MessageOriginHiddenUser) OriginType() string {
	return OriginTypeHiddenUser
}

// OriginalDate returns original message date
func (m *MessageOriginHiddenUser) OriginalDate() int64 {
	return m.Date
}

func (m *MessageOriginHiddenUser) iMessageOrigin() {}

// MessageOriginChat - The message was originally sent on behalf of a chat to a group chat.
type MessageOriginChat struct {
	// Type - Type of the message origin, always “chat”
	Type string `json:"type"`

	// Date - Date the message was sent originally in Unix time
	Date int64 `json:"date"`

	// SenderChat - Chat that sent the message originally
	SenderChat Chat `json:"sender_chat"`

	// AuthorSignature - Optional. For messages originally sent by an anonymous chat administrator, original
	// message author signature
	AuthorSignature string `json:"author_signature,omitempty"`
}

// OriginType returns original message type
func (m *MessageOriginChat) OriginType() string {
	return OriginTypeChat
}

// OriginalDate returns original message date
func (m *MessageOriginChat) OriginalDate() int64 {
	return m.Date
}

func (m *MessageOriginChat) iMessageOrigin() {}

// MessageOriginChannel - The message was originally sent to a channel chat.
type MessageOriginChannel struct {
	// Type - Type of the message origin, always “channel”
	Type string `json:"type"`

	// Date - Date the message was sent originally in Unix time
	Date int64 `json:"date"`

	// Chat - Channel chat to which the message was originally sent
	Chat Chat `json:"chat"`

	// MessageID - Unique message identifier inside the chat
	MessageID int `json:"message_id"`

	// AuthorSignature - Optional. Signature of the original post author
	AuthorSignature string `json:"author_signature,omitempty"`
}

// OriginType returns original message type
func (m *MessageOriginChannel) OriginType() string {
	return OriginTypeChannel
}

// OriginalDate returns original message date
func (m *MessageOriginChannel) OriginalDate() int64 {
	return m.Date
}

func (m *MessageOriginChannel) iMessageOrigin() {}

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

	// Width - Video width as defined by the sender
	Width int `json:"width"`

	// Height - Video height as defined by the sender
	Height int `json:"height"`

	// Duration - Duration of the video in seconds as defined by the sender
	Duration int `json:"duration"`

	// Thumbnail - Optional. Animation thumbnail as defined by the sender
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// FileName - Optional. Original animation filename as defined by the sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. MIME type of the file as defined by the sender
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

	// Duration - Duration of the audio in seconds as defined by the sender
	Duration int `json:"duration"`

	// Performer - Optional. Performer of the audio as defined by the sender or by audio tags
	Performer string `json:"performer,omitempty"`

	// Title - Optional. Title of the audio as defined by the sender or by audio tags
	Title string `json:"title,omitempty"`

	// FileName - Optional. Original filename as defined by the sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. MIME type of the file as defined by the sender
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

	// Thumbnail - Optional. Document thumbnail as defined by the sender
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// FileName - Optional. Original filename as defined by the sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. MIME type of the file as defined by the sender
	MimeType string `json:"mime_type,omitempty"`

	// FileSize - Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`
}

// Story - This object represents a story.
type Story struct {
	// Chat - Chat that posted the story
	Chat Chat `json:"chat"`

	// ID - Unique identifier for the story in the chat
	ID int `json:"id"`
}

// VideoQuality - This object represents a video file of a specific quality.
type VideoQuality struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width - Video width
	Width int `json:"width"`

	// Height - Video height
	Height int `json:"height"`

	// Codec - Codec that was used to encode the video, for example, “h264”, “h265”, or “av01”
	Codec string `json:"codec"`

	// FileSize - Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`
}

// Video - This object represents a video file.
type Video struct {
	// FileID - Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// FileUniqueID - Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width - Video width as defined by the sender
	Width int `json:"width"`

	// Height - Video height as defined by the sender
	Height int `json:"height"`

	// Duration - Duration of the video in seconds as defined by the sender
	Duration int `json:"duration"`

	// Thumbnail - Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// Cover - Optional. Available sizes of the cover of the video in the message
	Cover []PhotoSize `json:"cover,omitempty"`

	// StartTimestamp - Optional. Timestamp in seconds from which the video will play in the message
	StartTimestamp int `json:"start_timestamp,omitempty"`

	// Qualities - Optional. List of available qualities of the video
	Qualities []VideoQuality `json:"qualities,omitempty"`

	// FileName - Optional. Original filename as defined by the sender
	FileName string `json:"file_name,omitempty"`

	// MimeType - Optional. MIME type of the file as defined by the sender
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

	// Length - Video width and height (diameter of the video message) as defined by the sender
	Length int `json:"length"`

	// Duration - Duration of the video in seconds as defined by the sender
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

	// Duration - Duration of the audio in seconds as defined by the sender
	Duration int `json:"duration"`

	// MimeType - Optional. MIME type of the file as defined by the sender
	MimeType string `json:"mime_type,omitempty"`

	// FileSize - Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`
}

// PaidMediaInfo - Describes the paid media added to a message.
type PaidMediaInfo struct {
	// StarCount - The number of Telegram Stars that must be paid to buy access to the media
	StarCount int `json:"star_count"`

	// PaidMedia - Information about the paid media
	PaidMedia []PaidMedia `json:"paid_media"`
}

// UnmarshalJSON converts JSON to PaidMediaInfo
func (m *PaidMediaInfo) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	type uPaidMediaInfo PaidMediaInfo
	var um uPaidMediaInfo

	if value.Exists("paid_media") {
		paidMedia := value.GetArray("paid_media")
		um.PaidMedia = make([]PaidMedia, len(paidMedia))
		for i, media := range paidMedia {
			mediaType := string(media.GetStringBytes("type"))
			switch mediaType {
			case PaidMediaTypePreview:
				um.PaidMedia[i] = &PaidMediaPreview{}
			case PaidMediaTypePhoto:
				um.PaidMedia[i] = &PaidMediaPhoto{}
			case PaidMediaTypeVideo:
				um.PaidMedia[i] = &PaidMediaVideo{}
			case paidMediaTypeOther:
				um.PaidMedia[i] = &paidMediaOther{}
			default:
				return fmt.Errorf("unknown paid media type: %q", mediaType)
			}
		}
	}

	if err = json.Unmarshal(data, &um); err != nil {
		return err
	}
	*m = PaidMediaInfo(um)

	return nil
}

// PaidMedia - This object describes paid media. Currently, it can be one of
// PaidMediaPreview (https://core.telegram.org/bots/api#paidmediapreview)
// PaidMediaPhoto (https://core.telegram.org/bots/api#paidmediaphoto)
// PaidMediaVideo (https://core.telegram.org/bots/api#paidmediavideo)
//
// WARNING: Telegram introduced undocumented type "other", it will be correctly parsed by Telego, but will not be
// exposed to users directly as it's not documented anywhere, but still used by Telegram
type PaidMedia interface {
	// MediaType returns PaidMedia type
	MediaType() string
	// Disallow external implementations
	iPaidMedia()
}

// Paid media types
const (
	PaidMediaTypePreview = "preview"
	PaidMediaTypePhoto   = "photo"
	PaidMediaTypeVideo   = "video"
	paidMediaTypeOther   = "other"
)

// PaidMediaPreview - The paid media isn't available before the payment.
type PaidMediaPreview struct {
	// Type - Type of the paid media, always “preview”
	Type string `json:"type"`

	// Width - Optional. Media width as defined by the sender
	Width int `json:"width,omitempty"`

	// Height - Optional. Media height as defined by the sender
	Height int `json:"height,omitempty"`

	// Duration - Optional. Duration of the media in seconds as defined by the sender
	Duration int `json:"duration,omitempty"`
}

// MediaType returns PaidMedia type
func (m *PaidMediaPreview) MediaType() string {
	return PaidMediaTypePreview
}

func (m *PaidMediaPreview) iPaidMedia() {}

// PaidMediaPhoto - The paid media is a photo.
type PaidMediaPhoto struct {
	// Type - Type of the paid media, always “photo”
	Type string `json:"type"`

	// Photo - The photo
	Photo []PhotoSize `json:"photo"`
}

// MediaType returns PaidMedia type
func (m *PaidMediaPhoto) MediaType() string {
	return PaidMediaTypePhoto
}

func (m *PaidMediaPhoto) iPaidMedia() {}

// PaidMediaVideo - The paid media is a video.
type PaidMediaVideo struct {
	// Type - Type of the paid media, always “video”
	Type string `json:"type"`

	// Video - The video
	Video Video `json:"video"`
}

// MediaType returns PaidMedia type
func (m *PaidMediaVideo) MediaType() string {
	return PaidMediaTypeVideo
}

func (m *PaidMediaVideo) iPaidMedia() {}

// paidMediaOther - The paid media is a other.
//
// WARNING: This is undocumented type that was created because we saw it being used by Telegram, users of Telego are
// not expected to use this type as it is not documented anywhere
type paidMediaOther struct {
	// Type - Type of the paid media, always “other”
	Type string `json:"type"`
}

// MediaType returns PaidMedia type
func (m *paidMediaOther) MediaType() string {
	return paidMediaTypeOther
}

func (m *paidMediaOther) iPaidMedia() {}

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

	// TextEntities - Optional. Special entities that appear in the option text. Currently, only custom emoji
	// entities are allowed in poll option texts
	TextEntities []MessageEntity `json:"text_entities,omitempty"`

	// VoterCount - Number of users that voted for this option
	VoterCount int `json:"voter_count"`
}

// InputPollOption - This object contains information about one answer option in a poll to be sent.
type InputPollOption struct {
	// Text - Option text, 1-100 characters
	Text string `json:"text"`

	// TextParseMode - Optional. Mode for parsing entities in the text. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details. Currently, only custom emoji
	// entities are allowed
	TextParseMode string `json:"text_parse_mode,omitempty"`

	// TextEntities - Optional. A JSON-serialized list of special entities that appear in the poll option text.
	// It can be specified instead of text_parse_mode
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
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

	// QuestionEntities - Optional. Special entities that appear in the question. Currently, only custom emoji
	// entities are allowed in poll questions
	QuestionEntities []MessageEntity `json:"question_entities,omitempty"`

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

// ChecklistTask - Describes a task in a checklist.
type ChecklistTask struct {
	// ID - Unique identifier of the task
	ID int `json:"id"`

	// Text - Text of the task
	Text string `json:"text"`

	// TextEntities - Optional. Special entities that appear in the task text
	TextEntities []MessageEntity `json:"text_entities,omitempty"`

	// CompletedByUser - Optional. User that completed the task; omitted if the task wasn't completed by a user
	CompletedByUser *User `json:"completed_by_user,omitempty"`

	// CompletedByChat - Optional. Chat that completed the task; omitted if the task wasn't completed by a chat
	CompletedByChat *Chat `json:"completed_by_chat,omitempty"`

	// CompletionDate - Optional. Point in time (Unix timestamp) when the task was completed; 0 if the task
	// wasn't completed
	CompletionDate int64 `json:"completion_date,omitempty"`
}

// Checklist - Describes a checklist.
type Checklist struct {
	// Title - Title of the checklist
	Title string `json:"title"`

	// TitleEntities - Optional. Special entities that appear in the checklist title
	TitleEntities []MessageEntity `json:"title_entities,omitempty"`

	// Tasks - List of tasks in the checklist
	Tasks []ChecklistTask `json:"tasks"`

	// OthersCanAddTasks - Optional. True, if users other than the creator of the list can add tasks to the list
	OthersCanAddTasks bool `json:"others_can_add_tasks,omitempty"`

	// OthersCanMarkTasksAsDone - Optional. True, if users other than the creator of the list can mark tasks as
	// done or not done
	OthersCanMarkTasksAsDone bool `json:"others_can_mark_tasks_as_done,omitempty"`
}

// InputChecklistTask - Describes a task to add to a checklist.
type InputChecklistTask struct {
	// ID - Unique identifier of the task; must be positive and unique among all task identifiers currently
	// present in the checklist
	ID int `json:"id"`

	// Text - Text of the task; 1-100 characters after entities parsing
	Text string `json:"text"`

	// ParseMode - Optional. Mode for parsing entities in the text. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// TextEntities - Optional. List of special entities that appear in the text, which can be specified instead
	// of parse_mode. Currently, only bold, italic, underline, strikethrough, spoiler, and custom_emoji entities are
	// allowed.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

// InputChecklist - Describes a checklist to create.
type InputChecklist struct {
	// Title - Title of the checklist; 1-255 characters after entities parsing
	Title string `json:"title"`

	// ParseMode - Optional. Mode for parsing entities in the title. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// TitleEntities - Optional. List of special entities that appear in the title, which can be specified
	// instead of parse_mode. Currently, only bold, italic, underline, strikethrough, spoiler, and custom_emoji
	// entities are allowed.
	TitleEntities []MessageEntity `json:"title_entities,omitempty"`

	// Tasks - List of 1-30 tasks in the checklist
	Tasks []InputChecklistTask `json:"tasks"`

	// OthersCanAddTasks - Optional. Pass True if other users can add tasks to the checklist
	OthersCanAddTasks bool `json:"others_can_add_tasks,omitempty"`

	// OthersCanMarkTasksAsDone - Optional. Pass True if other users can mark tasks as done or not done in the
	// checklist
	OthersCanMarkTasksAsDone bool `json:"others_can_mark_tasks_as_done,omitempty"`
}

// ChecklistTasksDone - Describes a service message about checklist tasks marked as done or not done.
type ChecklistTasksDone struct {
	// ChecklistMessage - Optional. Message containing the checklist whose tasks were marked as done or not
	// done. Note that the Message (https://core.telegram.org/bots/api#message) object in this field will not
	// contain the reply_to_message field even if it itself is a reply.
	ChecklistMessage *Message `json:"checklist_message,omitempty"`

	// MarkedAsDoneTaskIDs - Optional. Identifiers of the tasks that were marked as done
	MarkedAsDoneTaskIDs []int `json:"marked_as_done_task_ids,omitempty"`

	// MarkedAsNotDoneTaskIDs - Optional. Identifiers of the tasks that were marked as not done
	MarkedAsNotDoneTaskIDs []int `json:"marked_as_not_done_task_ids,omitempty"`
}

// ChecklistTasksAdded - Describes a service message about tasks added to a checklist.
type ChecklistTasksAdded struct {
	// ChecklistMessage - Optional. Message containing the checklist to which the tasks were added. Note that
	// the Message (https://core.telegram.org/bots/api#message) object in this field will not contain the
	// reply_to_message field even if it itself is a reply.
	ChecklistMessage *Message `json:"checklist_message,omitempty"`

	// Tasks - List of tasks added to the checklist
	Tasks []ChecklistTask `json:"tasks"`
}

// Location - This object represents a point on the map.
type Location struct {
	// Latitude - Latitude as defined by the sender
	Latitude float64 `json:"latitude"`

	// Longitude - Longitude as defined by the sender
	Longitude float64 `json:"longitude"`

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

// ChatBoostAdded - This object represents a service message about a user boosting a chat.
type ChatBoostAdded struct {
	// BoostCount - Number of boosts added by the user
	BoostCount int `json:"boost_count"`
}

// BackgroundFill - This object describes the way a background is filled based on the selected colors.
// Currently, it can be one of
// BackgroundFillSolid (https://core.telegram.org/bots/api#backgroundfillsolid)
// BackgroundFillGradient (https://core.telegram.org/bots/api#backgroundfillgradient)
// BackgroundFillFreeformGradient (https://core.telegram.org/bots/api#backgroundfillfreeformgradient)
type BackgroundFill interface {
	// BackgroundFilled returns BackgroundFill type
	BackgroundFilled() string
	// Disallow external implementations
	iBackgroundFill()
}

// Background fill types
const (
	BackgroundFilledSolid            = "solid"
	BackgroundFilledGradient         = "gradient"
	BackgroundFilledFreeformGradient = "freeform_gradient"
)

// BackgroundFillSolid - The background is filled using the selected color.
type BackgroundFillSolid struct {
	// Type - Type of the background fill, always “solid”
	Type string `json:"type"`

	// Color - The color of the background fill in the RGB24 format
	Color int `json:"color"`
}

// BackgroundFilled returns BackgroundFill type
func (b *BackgroundFillSolid) BackgroundFilled() string {
	return BackgroundFilledSolid
}

func (b *BackgroundFillSolid) iBackgroundFill() {}

// BackgroundFillGradient - The background is a gradient fill.
type BackgroundFillGradient struct {
	// Type - Type of the background fill, always “gradient”
	Type string `json:"type"`

	// TopColor - Top color of the gradient in the RGB24 format
	TopColor int `json:"top_color"`

	// BottomColor - Bottom color of the gradient in the RGB24 format
	BottomColor int `json:"bottom_color"`

	// RotationAngle - Clockwise rotation angle of the background fill in degrees; 0-359
	RotationAngle int `json:"rotation_angle"`
}

// BackgroundFilled returns BackgroundFill type
func (b *BackgroundFillGradient) BackgroundFilled() string {
	return BackgroundFilledGradient
}

func (b *BackgroundFillGradient) iBackgroundFill() {}

// BackgroundFillFreeformGradient - The background is a freeform gradient that rotates after every message in
// the chat.
type BackgroundFillFreeformGradient struct {
	// Type - Type of the background fill, always “freeform_gradient”
	Type string `json:"type"`

	// Colors - A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24
	// format
	Colors []int `json:"colors"`
}

// BackgroundFilled returns BackgroundFill type
func (b *BackgroundFillFreeformGradient) BackgroundFilled() string {
	return BackgroundFilledFreeformGradient
}

func (b *BackgroundFillFreeformGradient) iBackgroundFill() {}

// BackgroundType - This object describes the type of a background. Currently, it can be one of
// BackgroundTypeFill (https://core.telegram.org/bots/api#backgroundtypefill)
// BackgroundTypeWallpaper (https://core.telegram.org/bots/api#backgroundtypewallpaper)
// BackgroundTypePattern (https://core.telegram.org/bots/api#backgroundtypepattern)
// BackgroundTypeChatTheme (https://core.telegram.org/bots/api#backgroundtypechattheme)
type BackgroundType interface {
	// BackgroundType returns BackgroundType type
	BackgroundType() string
	// Disallow external implementations
	iBackgroundType()
}

// Background type names
const (
	BackgroundTypeNameFill      = "fill"
	BackgroundTypeNameWallpaper = "wallpaper"
	BackgroundTypeNamePattern   = "pattern"
	BackgroundTypeNameChatTheme = "chat_theme"
)

// BackgroundTypeFill - The background is automatically filled based on the selected colors.
type BackgroundTypeFill struct {
	// Type - Type of the background, always “fill”
	Type string `json:"type"`

	// Fill - The background fill
	Fill BackgroundFill `json:"fill"`

	// DarkThemeDimming - Dimming of the background in dark themes, as a percentage; 0-100
	DarkThemeDimming int `json:"dark_theme_dimming"`
}

// BackgroundType returns BackgroundType type
func (b *BackgroundTypeFill) BackgroundType() string {
	return BackgroundTypeNameFill
}

func (b *BackgroundTypeFill) iBackgroundType() {}

// UnmarshalJSON converts JSON to BackgroundTypeFill
func (b *BackgroundTypeFill) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	if !value.Exists("fill") {
		return errors.New("no fill")
	}

	type uBackgroundTypeFill BackgroundTypeFill
	var ub uBackgroundTypeFill

	fillType := string(value.GetStringBytes("fill", "type"))
	switch fillType {
	case BackgroundFilledSolid:
		ub.Fill = &BackgroundFillSolid{}
	case BackgroundFilledGradient:
		ub.Fill = &BackgroundFillGradient{}
	case BackgroundFilledFreeformGradient:
		ub.Fill = &BackgroundFillFreeformGradient{}
	default:
		return fmt.Errorf("unknown chat background fill type: %q", fillType)
	}

	if err = json.Unmarshal(data, &ub); err != nil {
		return err
	}
	*b = BackgroundTypeFill(ub)

	return nil
}

// BackgroundTypeWallpaper - The background is a wallpaper in the JPEG format.
type BackgroundTypeWallpaper struct {
	// Type - Type of the background, always “wallpaper”
	Type string `json:"type"`

	// Document - Document with the wallpaper
	Document Document `json:"document"`

	// DarkThemeDimming - Dimming of the background in dark themes, as a percentage; 0-100
	DarkThemeDimming int `json:"dark_theme_dimming"`

	// IsBlurred - Optional. True, if the wallpaper is downscaled to fit in a 450x450 square and then
	// box-blurred with radius 12
	IsBlurred bool `json:"is_blurred,omitempty"`

	// IsMoving - Optional. True, if the background moves slightly when the device is tilted
	IsMoving bool `json:"is_moving,omitempty"`
}

// BackgroundType returns BackgroundType type
func (b *BackgroundTypeWallpaper) BackgroundType() string {
	return BackgroundTypeNameWallpaper
}

func (b *BackgroundTypeWallpaper) iBackgroundType() {}

// BackgroundTypePattern - The background is a .PNG or .TGV (gzipped subset of SVG with MIME type
// “application/x-tgwallpattern”) pattern to be combined with the background fill chosen by the user.
type BackgroundTypePattern struct {
	// Type - Type of the background, always “pattern”
	Type string `json:"type"`

	// Document - Document with the pattern
	Document Document `json:"document"`

	// Fill - The background fill that is combined with the pattern
	Fill BackgroundFill `json:"fill"`

	// Intensity - Intensity of the pattern when it is shown above the filled background; 0-100
	Intensity int `json:"intensity"`

	// IsInverted - Optional. True, if the background fill must be applied only to the pattern itself. All other
	// pixels are black in this case. For dark themes only
	IsInverted bool `json:"is_inverted,omitempty"`

	// IsMoving - Optional. True, if the background moves slightly when the device is tilted
	IsMoving bool `json:"is_moving,omitempty"`
}

// BackgroundType returns BackgroundType type
func (b *BackgroundTypePattern) BackgroundType() string {
	return BackgroundTypeNamePattern
}

func (b *BackgroundTypePattern) iBackgroundType() {}

// BackgroundTypeChatTheme - The background is taken directly from a built-in chat theme.
type BackgroundTypeChatTheme struct {
	// Type - Type of the background, always “chat_theme”
	Type string `json:"type"`

	// ThemeName - Name of the chat theme, which is usually an emoji
	ThemeName string `json:"theme_name"`
}

// BackgroundType returns BackgroundType type
func (b *BackgroundTypeChatTheme) BackgroundType() string {
	return BackgroundTypeNameChatTheme
}

func (b *BackgroundTypeChatTheme) iBackgroundType() {}

// ChatBackground - This object represents a chat background.
type ChatBackground struct {
	// Type - Type of the background
	Type BackgroundType `json:"type"`
}

// noTypeErr error
const noTypeErr = "no type"

// UnmarshalJSON converts JSON to ChatBackground
func (c *ChatBackground) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	if !value.Exists("type") {
		return errors.New(noTypeErr)
	}

	type uChatBackground ChatBackground
	var uc uChatBackground

	backgroundType := string(value.GetStringBytes("type", "type"))
	switch backgroundType {
	case BackgroundTypeNameFill:
		uc.Type = &BackgroundTypeFill{}
	case BackgroundTypeNameWallpaper:
		uc.Type = &BackgroundTypeWallpaper{}
	case BackgroundTypeNamePattern:
		uc.Type = &BackgroundTypePattern{}
	case BackgroundTypeNameChatTheme:
		uc.Type = &BackgroundTypeChatTheme{}
	default:
		return fmt.Errorf("unknown chat background type: %q", backgroundType)
	}

	if err = json.Unmarshal(data, &uc); err != nil {
		return err
	}
	*c = ChatBackground(uc)

	return nil
}

// ForumTopicCreated - This object represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	// Name - Name of the topic
	Name string `json:"name"`

	// IconColor - Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`

	// IconCustomEmojiID - Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`

	// IsNameImplicit - Optional. True, if the name of the topic wasn't specified explicitly by its creator and
	// likely needs to be changed by the bot
	IsNameImplicit bool `json:"is_name_implicit,omitempty"`
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

// SharedUser - This object contains information about a user that was shared with the bot using a
// KeyboardButtonRequestUsers (https://core.telegram.org/bots/api#keyboardbuttonrequestusers) button.
type SharedUser struct {
	// UserID - Identifier of the shared user. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so 64-bit integers or double-precision float types are safe for storing these identifiers.
	// The bot may not have access to the user and could be unable to use this identifier, unless the user is
	// already known to the bot by some other means.
	UserID int64 `json:"user_id"`

	// FirstName - Optional. First name of the user, if the name was requested by the bot
	FirstName string `json:"first_name,omitempty"`

	// LastName - Optional. Last name of the user, if the name was requested by the bot
	LastName string `json:"last_name,omitempty"`

	// Username - Optional. Username of the user, if the username was requested by the bot
	Username string `json:"username,omitempty"`

	// Photo - Optional. Available sizes of the chat photo, if the photo was requested by the bot
	Photo []PhotoSize `json:"photo,omitempty"`
}

// UsersShared - This object contains information about the users whose identifiers were shared with the bot
// using a KeyboardButtonRequestUsers (https://core.telegram.org/bots/api#keyboardbuttonrequestusers) button.
type UsersShared struct {
	// RequestID - Identifier of the request
	RequestID int `json:"request_id"`

	// Users - Information about users shared with the bot.
	Users []SharedUser `json:"users"`
}

// ChatShared - This object contains information about a chat that was shared with the bot using a
// KeyboardButtonRequestChat (https://core.telegram.org/bots/api#keyboardbuttonrequestchat) button.
type ChatShared struct {
	// RequestID - Identifier of the request
	RequestID int `json:"request_id"`

	// ChatID - Identifier of the shared chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	// The bot may not have access to the chat and could be unable to use this identifier, unless the chat is
	// already known to the bot by some other means.
	ChatID int64 `json:"chat_id"`

	// Title - Optional. Title of the chat, if the title was requested by the bot.
	Title string `json:"title,omitempty"`

	// Username - Optional. Username of the chat, if the username was requested by the bot and available.
	Username string `json:"username,omitempty"`

	// Photo - Optional. Available sizes of the chat photo, if the photo was requested by the bot
	Photo []PhotoSize `json:"photo,omitempty"`
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

// PaidMessagePriceChanged - Describes a service message about a change in the price of paid messages within
// a chat.
type PaidMessagePriceChanged struct {
	// PaidMessageStarCount - The new number of Telegram Stars that must be paid by non-administrator users of
	// the supergroup chat for each sent message
	PaidMessageStarCount int `json:"paid_message_star_count"`
}

// DirectMessagePriceChanged - Describes a service message about a change in the price of direct messages
// sent to a channel chat.
type DirectMessagePriceChanged struct {
	// AreDirectMessagesEnabled - True, if direct messages are enabled for the channel chat; false otherwise
	AreDirectMessagesEnabled bool `json:"are_direct_messages_enabled"`

	// DirectMessageStarCount - Optional. The new number of Telegram Stars that must be paid by users for each
	// direct message sent to the channel. Does not apply to users who have been exempted by administrators.
	// Defaults to 0.
	DirectMessageStarCount int `json:"direct_message_star_count,omitempty"`
}

// SuggestedPostApproved - Describes a service message about the approval of a suggested post.
type SuggestedPostApproved struct {
	// SuggestedPostMessage - Optional. Message containing the suggested post. Note that the Message
	// (https://core.telegram.org/bots/api#message) object in this field will not contain the reply_to_message field
	// even if it itself is a reply.
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`

	// Price - Optional. Amount paid for the post
	Price *SuggestedPostPrice `json:"price,omitempty"`

	// SendDate - Date when the post will be published
	SendDate int64 `json:"send_date"`
}

// SuggestedPostApprovalFailed - Describes a service message about the failed approval of a suggested post.
// Currently, only caused by insufficient user funds at the time of approval.
type SuggestedPostApprovalFailed struct {
	// SuggestedPostMessage - Optional. Message containing the suggested post whose approval has failed. Note
	// that the Message (https://core.telegram.org/bots/api#message) object in this field will not contain the
	// reply_to_message field even if it itself is a reply.
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`

	// Price - Expected price of the post
	Price SuggestedPostPrice `json:"price"`
}

// SuggestedPostDeclined - Describes a service message about the rejection of a suggested post.
type SuggestedPostDeclined struct {
	// SuggestedPostMessage - Optional. Message containing the suggested post. Note that the Message
	// (https://core.telegram.org/bots/api#message) object in this field will not contain the reply_to_message field
	// even if it itself is a reply.
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`

	// Comment - Optional. Comment with which the post was declined
	Comment string `json:"comment,omitempty"`
}

// SuggestedPostPaid - Describes a service message about a successful payment for a suggested post.
type SuggestedPostPaid struct {
	// SuggestedPostMessage - Optional. Message containing the suggested post. Note that the Message
	// (https://core.telegram.org/bots/api#message) object in this field will not contain the reply_to_message field
	// even if it itself is a reply.
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`

	// Currency - Currency in which the payment was made. Currently, one of “XTR” for Telegram Stars or
	// “TON” for toncoins
	Currency string `json:"currency"`

	// Amount - Optional. The amount of the currency that was received by the channel in nanotoncoins; for
	// payments in toncoins only
	Amount int `json:"amount,omitempty"`

	// StarAmount - Optional. The amount of Telegram Stars that was received by the channel; for payments in
	// Telegram Stars only
	StarAmount *StarAmount `json:"star_amount,omitempty"`
}

// SuggestedPostRefunded - Describes a service message about a payment refund for a suggested post.
type SuggestedPostRefunded struct {
	// SuggestedPostMessage - Optional. Message containing the suggested post. Note that the Message
	// (https://core.telegram.org/bots/api#message) object in this field will not contain the reply_to_message field
	// even if it itself is a reply.
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`

	// Reason - Reason for the refund. Currently, one of “post_deleted” if the post was deleted within 24
	// hours of being posted or removed from scheduled messages without being posted, or “payment_refunded” if
	// the payer refunded their payment.
	Reason string `json:"reason"`
}

// GiveawayCreated - This object represents a service message about the creation of a scheduled giveaway.
type GiveawayCreated struct {
	// PrizeStarCount - Optional. The number of Telegram Stars to be split between giveaway winners; for
	// Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count,omitempty"`
}

// Giveaway - This object represents a message about a scheduled giveaway.
type Giveaway struct {
	// Chats - The list of chats which the user must join to participate in the giveaway
	Chats []Chat `json:"chats"`

	// WinnersSelectionDate - Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnersSelectionDate int64 `json:"winners_selection_date"`

	// WinnerCount - The number of users which are supposed to be selected as winners of the giveaway
	WinnerCount int `json:"winner_count"`

	// OnlyNewMembers - Optional. True, if only users who join the chats after the giveaway started should be
	// eligible to win
	OnlyNewMembers bool `json:"only_new_members,omitempty"`

	// HasPublicWinners - Optional. True, if the list of giveaway winners will be visible to everyone
	HasPublicWinners bool `json:"has_public_winners,omitempty"`

	// PrizeDescription - Optional. Description of additional giveaway prize
	PrizeDescription string `json:"prize_description,omitempty"`

	// CountryCodes - Optional. A list of two-letter ISO 3166-1 alpha-2
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes indicating the countries from which eligible
	// users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a
	// phone number that was bought on Fragment can always participate in giveaways.
	CountryCodes []string `json:"country_codes,omitempty"`

	// PrizeStarCount - Optional. The number of Telegram Stars to be split between giveaway winners; for
	// Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count,omitempty"`

	// PremiumSubscriptionMonthCount - Optional. The number of months the Telegram Premium subscription won from
	// the giveaway will be active for; for Telegram Premium giveaways only
	PremiumSubscriptionMonthCount int `json:"premium_subscription_month_count,omitempty"`
}

// GiveawayWinners - This object represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	// Chat - The chat that created the giveaway
	Chat Chat `json:"chat"`

	// GiveawayMessageID - Identifier of the message with the giveaway in the chat
	GiveawayMessageID int `json:"giveaway_message_id"`

	// WinnersSelectionDate - Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnersSelectionDate int64 `json:"winners_selection_date"`

	// WinnerCount - Total number of winners in the giveaway
	WinnerCount int `json:"winner_count"`

	// Winners - List of up to 100 winners of the giveaway
	Winners []User `json:"winners"`

	// AdditionalChatCount - Optional. The number of other chats the user had to join in order to be eligible
	// for the giveaway
	AdditionalChatCount int `json:"additional_chat_count,omitempty"`

	// PrizeStarCount - Optional. The number of Telegram Stars that were split between giveaway winners; for
	// Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count,omitempty"`

	// PremiumSubscriptionMonthCount - Optional. The number of months the Telegram Premium subscription won from
	// the giveaway will be active for; for Telegram Premium giveaways only
	PremiumSubscriptionMonthCount int `json:"premium_subscription_month_count,omitempty"`

	// UnclaimedPrizeCount - Optional. Number of undistributed prizes
	UnclaimedPrizeCount int `json:"unclaimed_prize_count,omitempty"`

	// OnlyNewMembers - Optional. True, if only users who had joined the chats after the giveaway started were
	// eligible to win
	OnlyNewMembers bool `json:"only_new_members,omitempty"`

	// WasRefunded - Optional. True, if the giveaway was canceled because the payment for it was refunded
	WasRefunded bool `json:"was_refunded,omitempty"`

	// PrizeDescription - Optional. Description of additional giveaway prize
	PrizeDescription string `json:"prize_description,omitempty"`
}

// GiveawayCompleted - This object represents a service message about the completion of a giveaway without
// public winners.
type GiveawayCompleted struct {
	// WinnerCount - Number of winners in the giveaway
	WinnerCount int `json:"winner_count"`

	// UnclaimedPrizeCount - Optional. Number of undistributed prizes
	UnclaimedPrizeCount int `json:"unclaimed_prize_count,omitempty"`

	// GiveawayMessage - Optional. Message with the giveaway that was completed, if it wasn't deleted
	GiveawayMessage *Message `json:"giveaway_message,omitempty"`

	// IsStarGiveaway - Optional. True, if the giveaway is a Telegram Star giveaway. Otherwise, currently, the
	// giveaway is a Telegram Premium giveaway.
	IsStarGiveaway bool `json:"is_star_giveaway,omitempty"`
}

// LinkPreviewOptions - Describes the options used for link preview generation.
type LinkPreviewOptions struct {
	// IsDisabled - Optional. True, if the link preview is disabled
	IsDisabled bool `json:"is_disabled,omitempty"`

	// URL - Optional. URL to use for the link preview. If empty, then the first URL found in the message text
	// will be used
	URL string `json:"url,omitempty"`

	// PreferSmallMedia - Optional. True, if the media in the link preview is supposed to be shrunk; ignored if
	// the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferSmallMedia bool `json:"prefer_small_media,omitempty"`

	// PreferLargeMedia - Optional. True, if the media in the link preview is supposed to be enlarged; ignored
	// if the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferLargeMedia bool `json:"prefer_large_media,omitempty"`

	// ShowAboveText - Optional. True, if the link preview must be shown above the message text; otherwise, the
	// link preview will be shown below the message text
	ShowAboveText bool `json:"show_above_text,omitempty"`
}

// SuggestedPostPrice - Describes the price of a suggested post.
type SuggestedPostPrice struct {
	// Currency - Currency in which the post will be paid. Currently, must be one of “XTR” for Telegram
	// Stars or “TON” for toncoins
	Currency string `json:"currency"`

	// Amount - The amount of the currency that will be paid for the post in the smallest units of the currency,
	// i.e. Telegram Stars or nanotoncoins. Currently, price in Telegram Stars must be between 5 and 100000, and
	// price in nanotoncoins must be between 10000000 and 10000000000000.
	Amount int `json:"amount"`
}

// SuggestedPostInfo - Contains information about a suggested post.
type SuggestedPostInfo struct {
	// State - State of the suggested post. Currently, it can be one of “pending”, “approved”,
	// “declined”.
	State string `json:"state"`

	// Price - Optional. Proposed price of the post. If the field is omitted, then the post is unpaid.
	Price *SuggestedPostPrice `json:"price,omitempty"`

	// SendDate - Optional. Proposed send date of the post. If the field is omitted, then the post can be
	// published at any time within 30 days at the sole discretion of the user or administrator who approves it.
	SendDate int64 `json:"send_date,omitempty"`
}

// SuggestedPostParameters - Contains parameters of a post that is being suggested by the bot.
type SuggestedPostParameters struct {
	// Price - Optional. Proposed price for the post. If the field is omitted, then the post is unpaid.
	Price *SuggestedPostPrice `json:"price,omitempty"`

	// SendDate - Optional. Proposed send date of the post. If specified, then the date must be between 300
	// second and 2678400 seconds (30 days) in the future. If the field is omitted, then the post can be published
	// at any time within 30 days at the sole discretion of the user who approves it.
	SendDate int64 `json:"send_date,omitempty"`
}

// DirectMessagesTopic - Describes a topic of a direct messages chat.
type DirectMessagesTopic struct {
	// TopicID - Unique identifier of the topic. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	TopicID int64 `json:"topic_id"`

	// User - Optional. Information about the user that created the topic. Currently, it is always present
	User *User `json:"user,omitempty"`
}

// UserProfilePhotos - This object represent a user's profile pictures.
type UserProfilePhotos struct {
	// TotalCount - Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`

	// Photos - Requested profile pictures (in up to 4 sizes each)
	Photos [][]PhotoSize `json:"photos"`
}

// UserProfileAudios - This object represents the audios displayed on a user's profile.
type UserProfileAudios struct {
	// TotalCount - Total number of profile audios for the target user
	TotalCount int `json:"total_count"`

	// Audios - Requested profile audios
	Audios []Audio `json:"audios"`
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
	// ReplyType returns type of reply
	ReplyType() string
	// Disallow external implementations
	iReplyMarkup()
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
// (https://core.telegram.org/bots/features#keyboards) for details and examples). Not supported in channels and
// for messages sent on behalf of a Telegram Business account.
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
	// object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the
	// original message.
	// Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select
	// the new language. Other users in the group don't see the keyboard.
	Selective bool `json:"selective,omitempty"`
}

// ReplyType returns ReplyKeyboardMarkup type
func (r *ReplyKeyboardMarkup) ReplyType() string {
	return MarkupTypeReplyKeyboard
}

func (r *ReplyKeyboardMarkup) iReplyMarkup() {}

// KeyboardButton - This object represents one button of the reply keyboard. At most one of the fields other
// than text, icon_custom_emoji_id, and style must be used to specify the type of the button. For simple text
// buttons, String can be used instead of this object to specify the button text.
type KeyboardButton struct {
	// Text - Text of the button. If none of the fields other than text, icon_custom_emoji_id, and style are
	// used, it will be sent as a message when the button is pressed
	Text string `json:"text"`

	// IconCustomEmojiID - Optional. Unique identifier of the custom emoji shown before the text of the button.
	// Can only be used by bots that purchased additional usernames on Fragment (https://fragment.com) or in the
	// messages directly sent by the bot to private, group and supergroup chats if the owner of the bot has a
	// Telegram Premium subscription.
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`

	// Style - Optional. Style of the button. Must be one of “danger” (red), “success” (green) or
	// “primary” (blue). If omitted, then an app-specific style is used.
	Style string `json:"style,omitempty"`

	// RequestUsers - Optional. If specified, pressing the button will open a list of suitable users.
	// Identifiers of selected users will be sent to the bot in a “users_shared” service message. Available in
	// private chats only.
	RequestUsers *KeyboardButtonRequestUsers `json:"request_users,omitempty"`

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

// KeyboardButtonRequestUsers - This object defines the criteria used to request suitable users. Information
// about the selected users will be shared with the bot when the corresponding button is pressed. More about
// requesting users » (https://core.telegram.org/bots/features#chat-and-user-selection)
type KeyboardButtonRequestUsers struct {
	// RequestID - Signed 32-bit identifier of the request that will be received back in the UsersShared
	// (https://core.telegram.org/bots/api#usersshared) object. Must be unique within the message
	RequestID int32 `json:"request_id"`

	// UserIsBot - Optional. Pass True to request bots, pass False to request regular users. If not specified,
	// no additional restrictions are applied.
	UserIsBot *bool `json:"user_is_bot,omitempty"`

	// UserIsPremium - Optional. Pass True to request premium users, pass False to request non-premium users. If
	// not specified, no additional restrictions are applied.
	UserIsPremium *bool `json:"user_is_premium,omitempty"`

	// MaxQuantity - Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	MaxQuantity int `json:"max_quantity,omitempty"`

	// RequestName - Optional. Pass True to request the users' first and last names
	RequestName *bool `json:"request_name,omitempty"`

	// RequestUsername - Optional. Pass True to request the users' usernames
	RequestUsername *bool `json:"request_username,omitempty"`

	// RequestPhoto - Optional. Pass True to request the users' photos
	RequestPhoto *bool `json:"request_photo,omitempty"`
}

// KeyboardButtonRequestChat - This object defines the criteria used to request a suitable chat. Information
// about the selected chat will be shared with the bot when the corresponding button is pressed. The bot will be
// granted requested rights in the chat if appropriate. More about requesting chats »
// (https://core.telegram.org/bots/features#chat-and-user-selection).
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

	// RequestTitle - Optional. Pass True to request the chat's title
	RequestTitle *bool `json:"request_title,omitempty"`

	// RequestUsername - Optional. Pass True to request the chat's username
	RequestUsername *bool `json:"request_username,omitempty"`

	// RequestPhoto - Optional. Pass True to request the chat's photo
	RequestPhoto *bool `json:"request_photo,omitempty"`
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
// Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ReplyKeyboardRemove struct {
	// RemoveKeyboard - Requests clients to remove the custom keyboard (user will not be able to summon this
	// keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in
	// ReplyKeyboardMarkup (https://core.telegram.org/bots/api#replykeyboardmarkup))
	RemoveKeyboard bool `json:"remove_keyboard"`

	// Selective - Optional. Use this parameter if you want to remove the keyboard for specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message (https://core.telegram.org/bots/api#message)
	// object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the
	// original message.
	// Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the
	// keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
	Selective bool `json:"selective,omitempty"`
}

// ReplyType returns ReplyKeyboardRemove type
func (r *ReplyKeyboardRemove) ReplyType() string {
	return MarkupTypeReplyKeyboardRemove
}

func (r *ReplyKeyboardRemove) iReplyMarkup() {}

// InlineKeyboardMarkup - This object represents an inline keyboard
// (https://core.telegram.org/bots/features#inline-keyboards) that appears right next to the message it belongs
// to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard - Array of button rows, each represented by an Array of InlineKeyboardButton
	// (https://core.telegram.org/bots/api#inlinekeyboardbutton) objects
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// ReplyType returns InlineKeyboardMarkup type
func (i *InlineKeyboardMarkup) ReplyType() string {
	return MarkupTypeInlineKeyboard
}

func (i *InlineKeyboardMarkup) iReplyMarkup() {}

// InlineKeyboardButton - This object represents one button of an inline keyboard. Exactly one of the fields
// other than text, icon_custom_emoji_id, and style must be used to specify the type of the button.
type InlineKeyboardButton struct {
	// Text - Label text on the button
	Text string `json:"text"`

	// IconCustomEmojiID - Optional. Unique identifier of the custom emoji shown before the text of the button.
	// Can only be used by bots that purchased additional usernames on Fragment (https://fragment.com) or in the
	// messages directly sent by the bot to private, group and supergroup chats if the owner of the bot has a
	// Telegram Premium subscription.
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`

	// Style - Optional. Style of the button. Must be one of “danger” (red), “success” (green) or
	// “primary” (blue). If omitted, then an app-specific style is used.
	Style string `json:"style,omitempty"`

	// URL - Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id>
	// can be used to mention a user by their identifier without using a username, if this is allowed by their
	// privacy settings.
	URL string `json:"url,omitempty"`

	// CallbackData - Optional. Data to be sent in a callback query
	// (https://core.telegram.org/bots/api#callbackquery) to the bot when the button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`

	// WebApp - Optional. Description of the Web App (https://core.telegram.org/bots/webapps) that will be
	// launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of
	// the user using the method answerWebAppQuery (https://core.telegram.org/bots/api#answerwebappquery). Available
	// only in private chats between a user and the bot. Not supported for messages sent on behalf of a Telegram
	// Business account.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// LoginURL - Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement
	// for the Telegram Login Widget (https://core.telegram.org/widgets/login).
	LoginURL *LoginURL `json:"login_url,omitempty"`

	// SwitchInlineQuery - Optional. If set, pressing the button will prompt the user to select one of their
	// chats, open that chat and insert the bot's username and the specified inline query in the input field. May be
	// empty, in which case just the bot's username will be inserted. Not supported for messages sent in channel
	// direct messages chats and on behalf of a Telegram Business account.
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"`

	// SwitchInlineQueryCurrentChat - Optional. If set, pressing the button will insert the bot's username and
	// the specified inline query in the current chat's input field. May be empty, in which case only the bot's
	// username will be inserted.
	// This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting
	// something from multiple options. Not supported in channels and for messages sent in channel direct messages
	// chats and on behalf of a Telegram Business account.
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"`

	// SwitchInlineQueryChosenChat - Optional. If set, pressing the button will prompt the user to select one of
	// their chats of the specified type, open that chat and insert the bot's username and the specified inline
	// query in the input field. Not supported for messages sent in channel direct messages chats and on behalf of a
	// Telegram Business account.
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`

	// CopyText - Optional. Description of the button that copies the specified text to the clipboard.
	CopyText *CopyTextButton `json:"copy_text,omitempty"`

	// CallbackGame - Optional. Description of the game that will be launched when the user presses the button.
	// NOTE: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Pay - Optional. Specify True, to send a Pay button (https://core.telegram.org/bots/api#payments).
	// Substrings “⭐” and “XTR” in the buttons's text will be replaced with a Telegram Star icon.
	// NOTE: This type of button must always be the first button in the first row and can only be used in invoice
	// messages.
	Pay bool `json:"pay,omitempty"`
}

// Keyboard button styles
const (
	ButtonStyleDanger  = "danger"
	ButtonStyleSuccess = "success"
	ButtonStylePrimary = "primary"
)

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

// CopyTextButton - This object represents an inline keyboard button that copies specified text to the
// clipboard.
type CopyTextButton struct {
	// Text - The text to be copied to the clipboard; 1-256 characters
	Text string `json:"text"`
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

	// Message - Optional. Message sent by the bot with the callback button that originated the query
	Message MaybeInaccessibleMessage `json:"message,omitempty"`

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

// UnmarshalJSON converts JSON to CallbackQuery
func (q *CallbackQuery) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	type uCallbackQuery CallbackQuery
	var uq uCallbackQuery

	if value.Exists("message") {
		if value.GetInt("message", "date") == 0 {
			uq.Message = &InaccessibleMessage{}
		} else {
			uq.Message = &Message{}
		}
	}

	if err = json.Unmarshal(data, &uq); err != nil {
		return err
	}
	*q = CallbackQuery(uq)

	return nil
}

// ForceReply - Upon receiving a message with this object, Telegram clients will display a reply interface to
// the user (act as if the user has selected the bot's message and tapped 'Reply'). This can be extremely useful
// if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode
// (https://core.telegram.org/bots/features#privacy-mode). Not supported in channels and for messages sent on
// behalf of a Telegram Business account.
type ForceReply struct {
	// ForceReply - Shows reply interface to the user, as if they manually selected the bot's message and tapped
	// 'Reply'
	ForceReply bool `json:"force_reply"`

	// InputFieldPlaceholder - Optional. The placeholder to be shown in the input field when the reply is
	// active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective - Optional. Use this parameter if you want to force reply from specific users only. Targets: 1)
	// users that are @mentioned in the text of the Message (https://core.telegram.org/bots/api#message) object; 2)
	// if the bot's message is a reply to a message in the same chat and forum topic, sender of the original
	// message.
	Selective bool `json:"selective,omitempty"`
}

// ReplyType returns ForceReply type
func (f *ForceReply) ReplyType() string {
	return MarkupTypeForceReply
}

func (f *ForceReply) iReplyMarkup() {}

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

	// SubscriptionPeriod - Optional. The number of seconds the subscription will be active for before the next
	// payment
	SubscriptionPeriod int64 `json:"subscription_period,omitempty"`

	// SubscriptionPrice - Optional. The amount of Telegram Stars a user must pay initially and after each
	// subsequent subscription period to be a member of the chat using the link
	SubscriptionPrice int `json:"subscription_price,omitempty"`
}

// ChatAdministratorRights - Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	// IsAnonymous - True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// CanManageChat - True, if the administrator can access the chat event log, get boost list, see hidden
	// supergroup and channel members, report spam messages, ignore slow mode, and send messages to the chat without
	// paying Telegram Stars. Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat"`

	// CanDeleteMessages - True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`

	// CanManageVideoChats - True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// CanRestrictMembers - True, if the administrator can restrict, ban or unban chat members, or access
	// supergroup statistics
	CanRestrictMembers bool `json:"can_restrict_members"`

	// CanPromoteMembers - True, if the administrator can add new administrators with a subset of their own
	// privileges or demote administrators that they have promoted, directly or indirectly (promoted by
	// administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`

	// CanChangeInfo - True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// CanInviteUsers - True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// CanPostStories - True, if the administrator can post stories to the chat
	CanPostStories bool `json:"can_post_stories"`

	// CanEditStories - True, if the administrator can edit stories posted by other users, post stories to the
	// chat page, pin chat stories, and access the chat's story archive
	CanEditStories bool `json:"can_edit_stories"`

	// CanDeleteStories - True, if the administrator can delete stories posted by other users
	CanDeleteStories bool `json:"can_delete_stories"`

	// CanPostMessages - Optional. True, if the administrator can post messages in the channel, approve
	// suggested posts, or access channel statistics; for channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// CanEditMessages - Optional. True, if the administrator can edit messages of other users and can pin
	// messages; for channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// CanPinMessages - Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// CanManageTopics - Optional. True, if the user is allowed to create, rename, close, and reopen forum
	// topics; for supergroups only
	CanManageTopics bool `json:"can_manage_topics,omitempty"`

	// CanManageDirectMessages - Optional. True, if the administrator can manage direct messages of the channel
	// and decline suggested posts; for channels only
	CanManageDirectMessages bool `json:"can_manage_direct_messages,omitempty"`

	// CanManageTags - Optional. True, if the administrator can edit the tags of regular members; for groups and
	// supergroups only. If omitted defaults to the value of can_pin_messages.
	CanManageTags bool `json:"can_manage_tags,omitempty"`
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

	// ViaJoinRequest - Optional. True, if the user joined the chat after sending a direct join request without
	// using an invite link and being approved by an administrator
	ViaJoinRequest bool `json:"via_join_request,omitempty"`

	// ViaChatFolderInviteLink - Optional. True, if the user joined the chat via a chat folder invite link
	ViaChatFolderInviteLink bool `json:"via_chat_folder_invite_link,omitempty"`
}

// UnmarshalJSON converts JSON to ChatMemberUpdated
func (c *ChatMemberUpdated) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	if !value.Exists("old_chat_member") {
		return errors.New("no old chat member")
	}

	if !value.Exists("new_chat_member") {
		return errors.New("no new chat member")
	}

	type uChatMemberUpdated ChatMemberUpdated
	var uc uChatMemberUpdated

	oldMemberStatus := string(value.GetStringBytes("old_chat_member", "status"))
	switch oldMemberStatus {
	case MemberStatusCreator:
		uc.OldChatMember = &ChatMemberOwner{}
	case MemberStatusAdministrator:
		uc.OldChatMember = &ChatMemberAdministrator{}
	case MemberStatusMember:
		uc.OldChatMember = &ChatMemberMember{}
	case MemberStatusRestricted:
		uc.OldChatMember = &ChatMemberRestricted{}
	case MemberStatusLeft:
		uc.OldChatMember = &ChatMemberLeft{}
	case MemberStatusBanned:
		uc.OldChatMember = &ChatMemberBanned{}
	default:
		return fmt.Errorf("unknown chat member status: %q", oldMemberStatus)
	}

	newMemberStatus := string(value.GetStringBytes("new_chat_member", "status"))
	switch newMemberStatus {
	case MemberStatusCreator:
		uc.NewChatMember = &ChatMemberOwner{}
	case MemberStatusAdministrator:
		uc.NewChatMember = &ChatMemberAdministrator{}
	case MemberStatusMember:
		uc.NewChatMember = &ChatMemberMember{}
	case MemberStatusRestricted:
		uc.NewChatMember = &ChatMemberRestricted{}
	case MemberStatusLeft:
		uc.NewChatMember = &ChatMemberLeft{}
	case MemberStatusBanned:
		uc.NewChatMember = &ChatMemberBanned{}
	default:
		return fmt.Errorf("unknown chat member status: %q", newMemberStatus)
	}

	if err = json.Unmarshal(data, &uc); err != nil {
		return err
	}
	*c = ChatMemberUpdated(uc)

	return nil
}

type chatMemberData struct {
	Data ChatMember
}

// UnmarshalJSON converts JSON to chatMemberData
func (c *chatMemberData) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	memberStatus := string(value.GetStringBytes("status"))
	switch memberStatus {
	case MemberStatusCreator:
		var cm *ChatMemberOwner
		err = json.Unmarshal(data, &cm)
		c.Data = cm
	case MemberStatusAdministrator:
		var cm *ChatMemberAdministrator
		err = json.Unmarshal(data, &cm)
		c.Data = cm
	case MemberStatusMember:
		var cm *ChatMemberMember
		err = json.Unmarshal(data, &cm)
		c.Data = cm
	case MemberStatusRestricted:
		var cm *ChatMemberRestricted
		err = json.Unmarshal(data, &cm)
		c.Data = cm
	case MemberStatusLeft:
		var cm *ChatMemberLeft
		err = json.Unmarshal(data, &cm)
		c.Data = cm
	case MemberStatusBanned:
		var cm *ChatMemberBanned
		err = json.Unmarshal(data, &cm)
		c.Data = cm
	default:
		return fmt.Errorf("unknown member status: %q", memberStatus)
	}

	return err
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
	// MemberStatus returns ChatMember status
	MemberStatus() string
	// MemberUser returns ChatMember User
	MemberUser() User
	// MemberIsMember returns true if ChatMember is member
	MemberIsMember() bool
	// Disallow external implementations
	iChatMember()
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

// MemberIsMember returns true if ChatMember is member
func (c *ChatMemberOwner) MemberIsMember() bool {
	return true
}

func (c *ChatMemberOwner) iChatMember() {}

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

	// CanManageChat - True, if the administrator can access the chat event log, get boost list, see hidden
	// supergroup and channel members, report spam messages, ignore slow mode, and send messages to the chat without
	// paying Telegram Stars. Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat"`

	// CanDeleteMessages - True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`

	// CanManageVideoChats - True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// CanRestrictMembers - True, if the administrator can restrict, ban or unban chat members, or access
	// supergroup statistics
	CanRestrictMembers bool `json:"can_restrict_members"`

	// CanPromoteMembers - True, if the administrator can add new administrators with a subset of their own
	// privileges or demote administrators that they have promoted, directly or indirectly (promoted by
	// administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`

	// CanChangeInfo - True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// CanInviteUsers - True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// CanPostStories - True, if the administrator can post stories to the chat
	CanPostStories bool `json:"can_post_stories"`

	// CanEditStories - True, if the administrator can edit stories posted by other users, post stories to the
	// chat page, pin chat stories, and access the chat's story archive
	CanEditStories bool `json:"can_edit_stories"`

	// CanDeleteStories - True, if the administrator can delete stories posted by other users
	CanDeleteStories bool `json:"can_delete_stories"`

	// CanPostMessages - Optional. True, if the administrator can post messages in the channel, approve
	// suggested posts, or access channel statistics; for channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// CanEditMessages - Optional. True, if the administrator can edit messages of other users and can pin
	// messages; for channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// CanPinMessages - Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// CanManageTopics - Optional. True, if the user is allowed to create, rename, close, and reopen forum
	// topics; for supergroups only
	CanManageTopics bool `json:"can_manage_topics,omitempty"`

	// CanManageDirectMessages - Optional. True, if the administrator can manage direct messages of the channel
	// and decline suggested posts; for channels only
	CanManageDirectMessages bool `json:"can_manage_direct_messages,omitempty"`

	// CanManageTags - Optional. True, if the administrator can edit the tags of regular members; for groups and
	// supergroups only. If omitted defaults to the value of can_pin_messages.
	CanManageTags bool `json:"can_manage_tags,omitempty"`

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

// MemberIsMember returns true if ChatMember is member
func (c *ChatMemberAdministrator) MemberIsMember() bool {
	return true
}

func (c *ChatMemberAdministrator) iChatMember() {}

// ChatMemberMember - Represents a chat member (https://core.telegram.org/bots/api#chatmember) that has no
// additional privileges or restrictions.
type ChatMemberMember struct {
	// Status - The member's status in the chat, always “member”
	Status string `json:"status"`

	// Tag - Optional. Tag of the member
	Tag string `json:"tag,omitempty"`

	// User - Information about the user
	User User `json:"user"`

	// UntilDate - Optional. Date when the user's subscription will expire; Unix time
	UntilDate int64 `json:"until_date,omitempty"`
}

// MemberStatus returns ChatMember status
func (c *ChatMemberMember) MemberStatus() string {
	return MemberStatusMember
}

// MemberUser returns ChatMember User
func (c *ChatMemberMember) MemberUser() User {
	return c.User
}

// MemberIsMember returns true if ChatMember is member
func (c *ChatMemberMember) MemberIsMember() bool {
	return true
}

func (c *ChatMemberMember) iChatMember() {}

// ChatMemberRestricted - Represents a chat member (https://core.telegram.org/bots/api#chatmember) that is
// under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	// Status - The member's status in the chat, always “restricted”
	Status string `json:"status"`

	// Tag - Optional. Tag of the member
	Tag string `json:"tag,omitempty"`

	// User - Information about the user
	User User `json:"user"`

	// IsMember - True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member"`

	// CanSendMessages - True, if the user is allowed to send text messages, contacts, giveaways, giveaway
	// winners, invoices, locations and venues
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

	// CanSendPolls - True, if the user is allowed to send polls and checklists
	CanSendPolls bool `json:"can_send_polls"`

	// CanSendOtherMessages - True, if the user is allowed to send animations, games, stickers and use inline
	// bots
	CanSendOtherMessages bool `json:"can_send_other_messages"`

	// CanAddWebPagePreviews - True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`

	// CanEditTag - True, if the user is allowed to edit their own tag
	CanEditTag bool `json:"can_edit_tag"`

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

// MemberIsMember returns true if ChatMember is member
func (c *ChatMemberRestricted) MemberIsMember() bool {
	return c.IsMember
}

func (c *ChatMemberRestricted) iChatMember() {}

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

// MemberIsMember returns true if ChatMember is member
func (c *ChatMemberLeft) MemberIsMember() bool {
	return false
}

func (c *ChatMemberLeft) iChatMember() {}

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

// MemberIsMember returns true if ChatMember is member
func (c *ChatMemberBanned) MemberIsMember() bool {
	return false
}

func (c *ChatMemberBanned) iChatMember() {}

// ChatJoinRequest - Represents a join request sent to a chat.
type ChatJoinRequest struct {
	// Chat - Chat to which the request was sent
	Chat Chat `json:"chat"`

	// From - User that sent the join request
	From User `json:"from"`

	// UserChatID - Identifier of a private chat with the user who sent the join request. This number may have
	// more than 32 significant bits and some programming languages may have difficulty/silent defects in
	// interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type
	// are safe for storing this identifier. The bot can use this identifier for 5 minutes to send messages until
	// the join request is processed, assuming no other administrator contacted the user.
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
	// CanSendMessages - Optional. True, if the user is allowed to send text messages, contacts, giveaways,
	// giveaway winners, invoices, locations and venues
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

	// CanSendPolls - Optional. True, if the user is allowed to send polls and checklists
	CanSendPolls *bool `json:"can_send_polls,omitempty"`

	// CanSendOtherMessages - Optional. True, if the user is allowed to send animations, games, stickers and use
	// inline bots
	CanSendOtherMessages *bool `json:"can_send_other_messages,omitempty"`

	// CanAddWebPagePreviews - Optional. True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`

	// CanEditTag - Optional. True, if the user is allowed to edit their own tag
	CanEditTag *bool `json:"can_edit_tag,omitempty"`

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

// Birthdate - Describes the birthdate of a user.
type Birthdate struct {
	// Day - Day of the user's birth; 1-31
	Day int `json:"day"`

	// Month - Month of the user's birth; 1-12
	Month int `json:"month"`

	// Year - Optional. Year of the user's birth
	Year int `json:"year,omitempty"`
}

// BusinessIntro - Contains information about the start page settings of a Telegram Business account.
type BusinessIntro struct {
	// Title - Optional. Title text of the business intro
	Title string `json:"title,omitempty"`

	// Message - Optional. Message text of the business intro
	Message string `json:"message,omitempty"`

	// Sticker - Optional. Sticker of the business intro
	Sticker *Sticker `json:"sticker,omitempty"`
}

// BusinessLocation - Contains information about the location of a Telegram Business account.
type BusinessLocation struct {
	// Address - Address of the business
	Address string `json:"address"`

	// Location - Optional. Location of the business
	Location *Location `json:"location,omitempty"`
}

// BusinessOpeningHoursInterval - Describes an interval of time during which a business is open.
type BusinessOpeningHoursInterval struct {
	// OpeningMinute - The minute's sequence number in a week, starting on Monday, marking the start of the time
	// interval during which the business is open; 0 - 7 * 24 * 60
	OpeningMinute int `json:"opening_minute"`

	// ClosingMinute - The minute's sequence number in a week, starting on Monday, marking the end of the time
	// interval during which the business is open; 0 - 8 * 24 * 60
	ClosingMinute int `json:"closing_minute"`
}

// BusinessOpeningHours - Describes the opening hours of a business.
type BusinessOpeningHours struct {
	// TimeZoneName - Unique name of the time zone for which the opening hours are defined
	TimeZoneName string `json:"time_zone_name"`

	// OpeningHours - List of time intervals describing business opening hours
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

// UserRating - This object describes the rating of a user based on their Telegram Star spendings.
type UserRating struct {
	// Level - Current level of the user, indicating their reliability when purchasing digital goods and
	// services. A higher level suggests a more trustworthy customer; a negative level is likely reason for concern.
	Level int `json:"level"`

	// Rating - Numerical value of the user's rating; the higher the rating, the better
	Rating int `json:"rating"`

	// CurrentLevelRating - The rating value required to get the current level
	CurrentLevelRating int `json:"current_level_rating"`

	// NextLevelRating - Optional. The rating value required to get to the next level; omitted if the maximum
	// level was reached
	NextLevelRating int `json:"next_level_rating,omitempty"`
}

// StoryAreaPosition - Describes the position of a clickable area within a story.
type StoryAreaPosition struct {
	// XPercentage - The abscissa of the area's center, as a percentage of the media width
	XPercentage float64 `json:"x_percentage"`

	// YPercentage - The ordinate of the area's center, as a percentage of the media height
	YPercentage float64 `json:"y_percentage"`

	// WidthPercentage - The width of the area's rectangle, as a percentage of the media width
	WidthPercentage float64 `json:"width_percentage"`

	// HeightPercentage - The height of the area's rectangle, as a percentage of the media height
	HeightPercentage float64 `json:"height_percentage"`

	// RotationAngle - The clockwise rotation angle of the rectangle, in degrees; 0-360
	RotationAngle float64 `json:"rotation_angle"`

	// CornerRadiusPercentage - The radius of the rectangle corner rounding, as a percentage of the media width
	CornerRadiusPercentage float64 `json:"corner_radius_percentage"`
}

// LocationAddress - Describes the physical address of a location.
type LocationAddress struct {
	// CountryCode - The two-letter ISO 3166-1 alpha-2 country code of the country where the location is located
	CountryCode string `json:"country_code"`

	// State - Optional. State of the location
	State string `json:"state,omitempty"`

	// City - Optional. City of the location
	City string `json:"city,omitempty"`

	// Street - Optional. Street address of the location
	Street string `json:"street,omitempty"`
}

// StoryAreaType - Describes the type of a clickable area on a story. Currently, it can be one of
// StoryAreaTypeLocation (https://core.telegram.org/bots/api#storyareatypelocation)
// StoryAreaTypeSuggestedReaction (https://core.telegram.org/bots/api#storyareatypesuggestedreaction)
// StoryAreaTypeLink (https://core.telegram.org/bots/api#storyareatypelink)
// StoryAreaTypeWeather (https://core.telegram.org/bots/api#storyareatypeweather)
// StoryAreaTypeUniqueGift (https://core.telegram.org/bots/api#storyareatypeuniquegift)
type StoryAreaType interface {
	// StoryAreaType return StoryAreaType type
	StoryAreaType() string
	// Disallow external implementations
	iStoryAreaType()
}

// Story area types
const (
	StoryAreaLocation          = "location"
	StoryAreaSuggestedReaction = "suggested_reaction"
	StoryAreaLink              = "link"
	StoryAreaWeather           = "weather"
	StoryAreaUniqueGift        = "unique_gift"
)

// StoryAreaTypeLocation - Describes a story area pointing to a location. Currently, a story can have up to
// 10 location areas.
type StoryAreaTypeLocation struct {
	// Type - Type of the area, always “location”
	Type string `json:"type"`

	// Latitude - Location latitude in degrees
	Latitude float64 `json:"latitude"`

	// Longitude - Location longitude in degrees
	Longitude float64 `json:"longitude"`

	// Address - Optional. Address of the location
	Address *LocationAddress `json:"address,omitempty"`
}

// StoryAreaType returns StoryAreaType type
func (s *StoryAreaTypeLocation) StoryAreaType() string {
	return StoryAreaLocation
}

func (s *StoryAreaTypeLocation) iStoryAreaType() {}

// StoryAreaTypeSuggestedReaction - Describes a story area pointing to a suggested reaction. Currently, a
// story can have up to 5 suggested reaction areas.
type StoryAreaTypeSuggestedReaction struct {
	// Type - Type of the area, always “suggested_reaction”
	Type string `json:"type"`

	// ReactionType - Type of the reaction
	ReactionType ReactionType `json:"reaction_type"`

	// IsDark - Optional. Pass True if the reaction area has a dark background
	IsDark bool `json:"is_dark,omitempty"`

	// IsFlipped - Optional. Pass True if reaction area corner is flipped
	IsFlipped bool `json:"is_flipped,omitempty"`
}

// StoryAreaType returns StoryAreaType type
func (s *StoryAreaTypeSuggestedReaction) StoryAreaType() string {
	return StoryAreaSuggestedReaction
}

func (s *StoryAreaTypeSuggestedReaction) iStoryAreaType() {}

// StoryAreaTypeLink - Describes a story area pointing to an HTTP or tg:// link. Currently, a story can have
// up to 3 link areas.
type StoryAreaTypeLink struct {
	// Type - Type of the area, always “link”
	Type string `json:"type"`

	// URL - HTTP or tg:// URL to be opened when the area is clicked
	URL string `json:"url"`
}

// StoryAreaType returns StoryAreaType type
func (s *StoryAreaTypeLink) StoryAreaType() string {
	return StoryAreaLink
}

func (s *StoryAreaTypeLink) iStoryAreaType() {}

// StoryAreaTypeWeather - Describes a story area containing weather information. Currently, a story can have
// up to 3 weather areas.
type StoryAreaTypeWeather struct {
	// Type - Type of the area, always “weather”
	Type string `json:"type"`

	// Temperature - Temperature, in degree Celsius
	Temperature float64 `json:"temperature"`

	// Emoji - Emoji representing the weather
	Emoji string `json:"emoji"`

	// BackgroundColor - A color of the area background in the ARGB format
	BackgroundColor int `json:"background_color"`
}

// StoryAreaType returns StoryAreaType type
func (s *StoryAreaTypeWeather) StoryAreaType() string {
	return StoryAreaWeather
}

func (s *StoryAreaTypeWeather) iStoryAreaType() {}

// StoryAreaTypeUniqueGift - Describes a story area pointing to a unique gift. Currently, a story can have at
// most 1 unique gift area.
type StoryAreaTypeUniqueGift struct {
	// Type - Type of the area, always “unique_gift”
	Type string `json:"type"`

	// Name - Unique name of the gift
	Name string `json:"name"`
}

// StoryAreaType returns StoryAreaType type
func (s *StoryAreaTypeUniqueGift) StoryAreaType() string {
	return StoryAreaUniqueGift
}

func (s *StoryAreaTypeUniqueGift) iStoryAreaType() {}

// StoryArea - Describes a clickable area on a story media.
type StoryArea struct {
	// Position - Position of the area
	Position StoryAreaPosition `json:"position"`

	// Type - Type of the area
	Type StoryAreaType `json:"type"`
}

// ChatLocation - Represents a location to which a chat is connected.
type ChatLocation struct {
	// Location - The location to which the supergroup is connected. Can't be a live location.
	Location Location `json:"location"`

	// Address - Location address; 1-64 characters, as defined by the chat owner
	Address string `json:"address"`
}

// ReactionType - This object describes the type of a reaction. Currently, it can be one of
// ReactionTypeEmoji (https://core.telegram.org/bots/api#reactiontypeemoji)
// ReactionTypeCustomEmoji (https://core.telegram.org/bots/api#reactiontypecustomemoji)
// ReactionTypePaid (https://core.telegram.org/bots/api#reactiontypepaid)
type ReactionType interface {
	// ReactionType returns reaction type
	ReactionType() string
	// Disallow external implementations
	iReactionType()
}

// Reaction types
const (
	ReactionEmoji       = "emoji"
	ReactionCustomEmoji = "custom_emoji"
	ReactionPaid        = "paid"
)

// ReactionTypeEmoji - The reaction is based on an emoji.
type ReactionTypeEmoji struct {
	// Type - Type of the reaction, always “emoji”
	Type string `json:"type"`

	// Emoji - Reaction emoji. Currently, it can be one of "❤", "👍", "👎", "🔥", "🥰", "👏",
	// "😁", "🤔", "🤯", "😱", "🤬", "😢", "🎉", "🤩", "🤮", "💩", "🙏", "👌", "🕊",
	// "🤡", "🥱", "🥴", "😍", "🐳", "❤‍🔥", "🌚", "🌭", "💯", "🤣", "⚡", "🍌", "🏆",
	// "💔", "🤨", "😐", "🍓", "🍾", "💋", "🖕", "😈", "😴", "😭", "🤓", "👻",
	// "👨‍💻", "👀", "🎃", "🙈", "😇", "😨", "🤝", "✍", "🤗", "🫡", "🎅", "🎄", "☃",
	// "💅", "🤪", "🗿", "🆒", "💘", "🙉", "🦄", "😘", "💊", "🙊", "😎", "👾", "🤷‍♂",
	// "🤷", "🤷‍♀", "😡"
	Emoji string `json:"emoji"`
}

// ReactionType returns reaction type
func (r *ReactionTypeEmoji) ReactionType() string {
	return ReactionEmoji
}

func (r *ReactionTypeEmoji) iReactionType() {}

// ReactionTypeCustomEmoji - The reaction is based on a custom emoji.
type ReactionTypeCustomEmoji struct {
	// Type - Type of the reaction, always “custom_emoji”
	Type string `json:"type"`

	// CustomEmojiID - Custom emoji identifier
	CustomEmojiID string `json:"custom_emoji_id"`
}

// ReactionType returns reaction type
func (r *ReactionTypeCustomEmoji) ReactionType() string {
	return ReactionCustomEmoji
}

func (r *ReactionTypeCustomEmoji) iReactionType() {}

// ReactionTypePaid - The reaction is paid.
type ReactionTypePaid struct {
	// Type - Type of the reaction, always “paid”
	Type string `json:"type"`
}

// ReactionType returns reaction type
func (r *ReactionTypePaid) ReactionType() string {
	return ReactionPaid
}

func (r *ReactionTypePaid) iReactionType() {}

// ReactionCount - Represents a reaction added to a message along with the number of times it was added.
type ReactionCount struct {
	// Type - Type of the reaction
	Type ReactionType `json:"type"`

	// TotalCount - Number of times the reaction was added
	TotalCount int `json:"total_count"`
}

// UnmarshalJSON converts JSON to ReactionCount
func (c *ReactionCount) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	if !value.Exists("type") {
		return errors.New(noTypeErr)
	}

	type uReactionCount ReactionCount
	var uc uReactionCount

	reactionType := string(value.GetStringBytes("type", "type"))
	switch reactionType {
	case ReactionEmoji:
		uc.Type = &ReactionTypeEmoji{}
	case ReactionCustomEmoji:
		uc.Type = &ReactionTypeCustomEmoji{}
	case ReactionPaid:
		uc.Type = &ReactionTypePaid{}
	default:
		return fmt.Errorf(unknownReactionTypeErr, reactionType)
	}

	if err = json.Unmarshal(data, &uc); err != nil {
		return err
	}
	*c = ReactionCount(uc)

	return nil
}

// MessageReactionUpdated - This object represents a change of a reaction on a message performed by a user.
type MessageReactionUpdated struct {
	// Chat - The chat containing the message the user reacted to
	Chat Chat `json:"chat"`

	// MessageID - Unique identifier of the message inside the chat
	MessageID int `json:"message_id"`

	// User - Optional. The user that changed the reaction, if the user isn't anonymous
	User *User `json:"user,omitempty"`

	// ActorChat - Optional. The chat on behalf of which the reaction was changed, if the user is anonymous
	ActorChat *Chat `json:"actor_chat,omitempty"`

	// Date - Date of the change in Unix time
	Date int64 `json:"date"`

	// OldReaction - Previous list of reaction types that were set by the user
	OldReaction []ReactionType `json:"old_reaction"`

	// NewReaction - New list of reaction types that have been set by the user
	NewReaction []ReactionType `json:"new_reaction"`
}

// UnmarshalJSON converts JSON to MessageReactionUpdated
func (u *MessageReactionUpdated) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	if !value.Exists("old_reaction") {
		return errors.New("no old reactions")
	}

	if !value.Exists("new_reaction") {
		return errors.New("no new reactions")
	}

	type uMessageReactionUpdated MessageReactionUpdated
	var uu uMessageReactionUpdated

	oldReactions := value.GetArray("old_reaction")
	uu.OldReaction = make([]ReactionType, len(oldReactions))
	for i, reaction := range oldReactions {
		reactionType := string(reaction.GetStringBytes("type"))
		switch reactionType {
		case ReactionEmoji:
			uu.OldReaction[i] = &ReactionTypeEmoji{}
		case ReactionCustomEmoji:
			uu.OldReaction[i] = &ReactionTypeCustomEmoji{}
		case ReactionPaid:
			uu.OldReaction[i] = &ReactionTypePaid{}
		default:
			return fmt.Errorf(unknownReactionTypeErr, reactionType)
		}
	}

	newReactions := value.GetArray("new_reaction")
	uu.NewReaction = make([]ReactionType, len(newReactions))
	for i, reaction := range newReactions {
		reactionType := string(reaction.GetStringBytes("type"))
		switch reactionType {
		case ReactionEmoji:
			uu.NewReaction[i] = &ReactionTypeEmoji{}
		case ReactionCustomEmoji:
			uu.NewReaction[i] = &ReactionTypeCustomEmoji{}
		case ReactionPaid:
			uu.NewReaction[i] = &ReactionTypePaid{}
		default:
			return fmt.Errorf(unknownReactionTypeErr, reactionType)
		}
	}

	if err = json.Unmarshal(data, &uu); err != nil {
		return err
	}
	*u = MessageReactionUpdated(uu)

	return nil
}

// MessageReactionCountUpdated - This object represents reaction changes on a message with anonymous
// reactions.
type MessageReactionCountUpdated struct {
	// Chat - The chat containing the message
	Chat Chat `json:"chat"`

	// MessageID - Unique message identifier inside the chat
	MessageID int `json:"message_id"`

	// Date - Date of the change in Unix time
	Date int64 `json:"date"`

	// Reactions - List of reactions that are present on the message
	Reactions []ReactionCount `json:"reactions"`
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

	// IsNameImplicit - Optional. True, if the name of the topic wasn't specified explicitly by its creator and
	// likely needs to be changed by the bot
	IsNameImplicit bool `json:"is_name_implicit,omitempty"`
}

// GiftBackground - This object describes the background of a gift.
type GiftBackground struct {
	// CenterColor - Center color of the background in RGB format
	CenterColor int `json:"center_color"`

	// EdgeColor - Edge color of the background in RGB format
	EdgeColor int `json:"edge_color"`

	// TextColor - Text color of the background in RGB format
	TextColor int `json:"text_color"`
}

// Gift - This object represents a gift that can be sent by the bot.
type Gift struct {
	// ID - Unique identifier of the gift
	ID string `json:"id"`

	// Sticker - The sticker that represents the gift
	Sticker Sticker `json:"sticker"`

	// StarCount - The number of Telegram Stars that must be paid to send the sticker
	StarCount int `json:"star_count"`

	// UpgradeStarCount - Optional. The number of Telegram Stars that must be paid to upgrade the gift to a
	// unique one
	UpgradeStarCount int `json:"upgrade_star_count,omitempty"`

	// IsPremium - Optional. True, if the gift can only be purchased by Telegram Premium subscribers
	IsPremium bool `json:"is_premium,omitempty"`

	// HasColors - Optional. True, if the gift can be used (after being upgraded) to customize a user's
	// appearance
	HasColors bool `json:"has_colors,omitempty"`

	// TotalCount - Optional. The total number of gifts of this type that can be sent by all users; for limited
	// gifts only
	TotalCount int `json:"total_count,omitempty"`

	// RemainingCount - Optional. The number of remaining gifts of this type that can be sent by all users; for
	// limited gifts only
	RemainingCount int `json:"remaining_count,omitempty"`

	// PersonalTotalCount - Optional. The total number of gifts of this type that can be sent by the bot; for
	// limited gifts only
	PersonalTotalCount int `json:"personal_total_count,omitempty"`

	// PersonalRemainingCount - Optional. The number of remaining gifts of this type that can be sent by the
	// bot; for limited gifts only
	PersonalRemainingCount int `json:"personal_remaining_count,omitempty"`

	// Background - Optional. Background of the gift
	Background *GiftBackground `json:"background,omitempty"`

	// UniqueGiftVariantCount - Optional. The total number of different unique gifts that can be obtained by
	// upgrading the gift
	UniqueGiftVariantCount int `json:"unique_gift_variant_count,omitempty"`

	// PublisherChat - Optional. Information about the chat that published the gift
	PublisherChat *Chat `json:"publisher_chat,omitempty"`
}

// Gifts - This object represent a list of gifts.
type Gifts struct {
	// Gifts - The list of gifts
	Gifts []Gift `json:"gifts"`
}

// UniqueGiftModel - This object describes the model of a unique gift.
type UniqueGiftModel struct {
	// Name - Name of the model
	Name string `json:"name"`

	// Sticker - The sticker that represents the unique gift
	Sticker Sticker `json:"sticker"`

	// RarityPerMille - The number of unique gifts that receive this model for every 1000 gift upgrades. Always
	// 0 for crafted gifts.
	RarityPerMille int `json:"rarity_per_mille"`

	// Rarity - Optional. Rarity of the model if it is a crafted model. Currently, can be “uncommon”,
	// “rare”, “epic”, or “legendary”.
	Rarity string `json:"rarity,omitempty"`
}

// Girt rarities
const (
	GiftRarityUncommon  = "uncommon"
	GiftRarityRare      = "rare"
	GiftRarityEpic      = "epic"
	GiftRarityLegendary = "legendary"
)

// UniqueGiftSymbol - This object describes the symbol shown on the pattern of a unique gift.
type UniqueGiftSymbol struct {
	// Name - Name of the symbol
	Name string `json:"name"`

	// Sticker - The sticker that represents the unique gift
	Sticker Sticker `json:"sticker"`

	// RarityPerMille - The number of unique gifts that receive this model for every 1000 gifts upgraded
	RarityPerMille int `json:"rarity_per_mille"`
}

// UniqueGiftBackdropColors - This object describes the colors of the backdrop of a unique gift.
type UniqueGiftBackdropColors struct {
	// CenterColor - The color in the center of the backdrop in RGB format
	CenterColor int `json:"center_color"`

	// EdgeColor - The color on the edges of the backdrop in RGB format
	EdgeColor int `json:"edge_color"`

	// SymbolColor - The color to be applied to the symbol in RGB format
	SymbolColor int `json:"symbol_color"`

	// TextColor - The color for the text on the backdrop in RGB format
	TextColor int `json:"text_color"`
}

// UniqueGiftBackdrop - This object describes the backdrop of a unique gift.
type UniqueGiftBackdrop struct {
	// Name - Name of the backdrop
	Name string `json:"name"`

	// Colors - Colors of the backdrop
	Colors UniqueGiftBackdropColors `json:"colors"`

	// RarityPerMille - The number of unique gifts that receive this backdrop for every 1000 gifts upgraded
	RarityPerMille int `json:"rarity_per_mille"`
}

// UniqueGiftColors - This object contains information about the color scheme for a user's name, message
// replies and link previews based on a unique gift.
type UniqueGiftColors struct {
	// ModelCustomEmojiID - Custom emoji identifier of the unique gift's model
	ModelCustomEmojiID string `json:"model_custom_emoji_id"`

	// SymbolCustomEmojiID - Custom emoji identifier of the unique gift's symbol
	SymbolCustomEmojiID string `json:"symbol_custom_emoji_id"`

	// LightThemeMainColor - Main color used in light themes; RGB format
	LightThemeMainColor int `json:"light_theme_main_color"`

	// LightThemeOtherColors - List of 1-3 additional colors used in light themes; RGB format
	LightThemeOtherColors []int `json:"light_theme_other_colors"`

	// DarkThemeMainColor - Main color used in dark themes; RGB format
	DarkThemeMainColor int `json:"dark_theme_main_color"`

	// DarkThemeOtherColors - List of 1-3 additional colors used in dark themes; RGB format
	DarkThemeOtherColors []int `json:"dark_theme_other_colors"`
}

// UniqueGift - This object describes a unique gift that was upgraded from a regular gift.
type UniqueGift struct {
	// GiftID - Identifier of the regular gift from which the gift was upgraded
	GiftID string `json:"gift_id"`

	// BaseName - Human-readable name of the regular gift from which this unique gift was upgraded
	BaseName string `json:"base_name"`

	// Name - Unique name of the gift. This name can be used in https://t.me/nft/... links and story areas
	Name string `json:"name"`

	// Number - Unique number of the upgraded gift among gifts upgraded from the same regular gift
	Number int `json:"number"`

	// Model - Model of the gift
	Model UniqueGiftModel `json:"model"`

	// Symbol - Symbol of the gift
	Symbol UniqueGiftSymbol `json:"symbol"`

	// Backdrop - Backdrop of the gift
	Backdrop UniqueGiftBackdrop `json:"backdrop"`

	// IsPremium - Optional. True, if the original regular gift was exclusively purchaseable by Telegram Premium
	// subscribers
	IsPremium bool `json:"is_premium,omitempty"`

	// IsBurned - Optional. True, if the gift was used to craft another gift and isn't available anymore
	IsBurned bool `json:"is_burned,omitempty"`

	// IsFromBlockchain - Optional. True, if the gift is assigned from the TON blockchain and can't be resold or
	// transferred in Telegram
	IsFromBlockchain bool `json:"is_from_blockchain,omitempty"`

	// Colors - Optional. The color scheme that can be used by the gift's owner for the chat's name, replies to
	// messages and link previews; for business account gifts and gifts that are currently on sale only
	Colors *UniqueGiftColors `json:"colors,omitempty"`

	// PublisherChat - Optional. Information about the chat that published the gift
	PublisherChat *Chat `json:"publisher_chat,omitempty"`
}

// GiftInfo - Describes a service message about a regular gift that was sent or received.
type GiftInfo struct {
	// Gift - Information about the gift
	Gift Gift `json:"gift"`

	// OwnedGiftID - Optional. Unique identifier of the received gift for the bot; only present for gifts
	// received on behalf of business accounts
	OwnedGiftID string `json:"owned_gift_id,omitempty"`

	// ConvertStarCount - Optional. Number of Telegram Stars that can be claimed by the receiver by converting
	// the gift; omitted if conversion to Telegram Stars is impossible
	ConvertStarCount int `json:"convert_star_count,omitempty"`

	// PrepaidUpgradeStarCount - Optional. Number of Telegram Stars that were prepaid for the ability to upgrade
	// the gift
	PrepaidUpgradeStarCount int `json:"prepaid_upgrade_star_count,omitempty"`

	// IsUpgradeSeparate - Optional. True, if the gift's upgrade was purchased after the gift was sent
	IsUpgradeSeparate bool `json:"is_upgrade_separate,omitempty"`

	// CanBeUpgraded - Optional. True, if the gift can be upgraded to a unique gift
	CanBeUpgraded bool `json:"can_be_upgraded,omitempty"`

	// Text - Optional. Text of the message that was added to the gift
	Text string `json:"text,omitempty"`

	// Entities - Optional. Special entities that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`

	// IsPrivate - Optional. True, if the sender and gift text are shown only to the gift receiver; otherwise,
	// everyone will be able to see them
	IsPrivate bool `json:"is_private,omitempty"`

	// UniqueGiftNumber - Optional. Unique number reserved for this gift when upgraded. See the number field in
	// UniqueGift (https://core.telegram.org/bots/api#uniquegift)
	UniqueGiftNumber int `json:"unique_gift_number,omitempty"`
}

// UniqueGiftInfo - Describes a service message about a unique gift that was sent or received.
type UniqueGiftInfo struct {
	// Gift - Information about the gift
	Gift UniqueGift `json:"gift"`

	// Origin - Origin of the gift. Currently, either “upgrade” for gifts upgraded from regular gifts,
	// “transfer” for gifts transferred from other users or channels, “resale” for gifts bought from other
	// users, “gifted_upgrade” for upgrades purchased after the gift was sent, or “offer” for gifts bought
	// or sold through gift purchase offers
	Origin string `json:"origin"`

	// LastResaleCurrency - Optional. For gifts bought from other users, the currency in which the payment for
	// the gift was done. Currently, one of “XTR” for Telegram Stars or “TON” for toncoins.
	LastResaleCurrency string `json:"last_resale_currency,omitempty"`

	// LastResaleAmount - Optional. For gifts bought from other users, the price paid for the gift in either
	// Telegram Stars or nanotoncoins
	LastResaleAmount int `json:"last_resale_amount,omitempty"`

	// OwnedGiftID - Optional. Unique identifier of the received gift for the bot; only present for gifts
	// received on behalf of business accounts
	OwnedGiftID string `json:"owned_gift_id,omitempty"`

	// TransferStarCount - Optional. Number of Telegram Stars that must be paid to transfer the gift; omitted if
	// the bot cannot transfer the gift
	TransferStarCount int `json:"transfer_star_count,omitempty"`

	// NextTransferDate - Optional. Point in time (Unix timestamp) when the gift can be transferred. If it is in
	// the past, then the gift can be transferred now
	NextTransferDate int64 `json:"next_transfer_date,omitempty"`
}

// Gift origins
const (
	GiftOriginUpgrade       = "upgrade"
	GiftOriginTransfer      = "transfer"
	GiftOriginResale        = "resale"
	GiftOriginGiftedUpgrade = "gifted_upgrade"
	GiftOriginOffer         = "offer"
)

// OwnedGift - This object describes a gift received and owned by a user or a chat. Currently, it can be one
// of
// OwnedGiftRegular (https://core.telegram.org/bots/api#ownedgiftregular)
// OwnedGiftUnique (https://core.telegram.org/bots/api#ownedgiftunique)
type OwnedGift interface {
	// GiftType returns OwnedGift type
	GiftType() string
	// Disallow external implementations
	iOwnedGift()
}

// Gift types
const (
	GiftTypeRegular = "regular"
	GiftTypeUnique  = "unique"
)

// OwnedGiftRegular - Describes a regular gift owned by a user or a chat.
type OwnedGiftRegular struct {
	// Type - Type of the gift, always “regular”
	Type string `json:"type"`

	// Gift - Information about the regular gift
	Gift Gift `json:"gift"`

	// OwnedGiftID - Optional. Unique identifier of the gift for the bot; for gifts received on behalf of
	// business accounts only
	OwnedGiftID string `json:"owned_gift_id,omitempty"`

	// SenderUser - Optional. Sender of the gift if it is a known user
	SenderUser *User `json:"sender_user,omitempty"`

	// SendDate - Date the gift was sent in Unix time
	SendDate int64 `json:"send_date"`

	// Text - Optional. Text of the message that was added to the gift
	Text string `json:"text,omitempty"`

	// Entities - Optional. Special entities that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`

	// IsPrivate - Optional. True, if the sender and gift text are shown only to the gift receiver; otherwise,
	// everyone will be able to see them
	IsPrivate bool `json:"is_private,omitempty"`

	// IsSaved - Optional. True, if the gift is displayed on the account's profile page; for gifts received on
	// behalf of business accounts only
	IsSaved bool `json:"is_saved,omitempty"`

	// CanBeUpgraded - Optional. True, if the gift can be upgraded to a unique gift; for gifts received on
	// behalf of business accounts only
	CanBeUpgraded bool `json:"can_be_upgraded,omitempty"`

	// WasRefunded - Optional. True, if the gift was refunded and isn't available anymore
	WasRefunded bool `json:"was_refunded,omitempty"`

	// ConvertStarCount - Optional. Number of Telegram Stars that can be claimed by the receiver instead of the
	// gift; omitted if the gift cannot be converted to Telegram Stars; for gifts received on behalf of business
	// accounts only
	ConvertStarCount int `json:"convert_star_count,omitempty"`

	// PrepaidUpgradeStarCount - Optional. Number of Telegram Stars that were paid for the ability to upgrade
	// the gift
	PrepaidUpgradeStarCount int `json:"prepaid_upgrade_star_count,omitempty"`

	// IsUpgradeSeparate - Optional. True, if the gift's upgrade was purchased after the gift was sent; for
	// gifts received on behalf of business accounts only
	IsUpgradeSeparate bool `json:"is_upgrade_separate,omitempty"`

	// UniqueGiftNumber - Optional. Unique number reserved for this gift when upgraded. See the number field in
	// UniqueGift (https://core.telegram.org/bots/api#uniquegift)
	UniqueGiftNumber int `json:"unique_gift_number,omitempty"`
}

// GiftType returns OwnedGift type
func (g *OwnedGiftRegular) GiftType() string {
	return GiftTypeRegular
}

func (g *OwnedGiftRegular) iOwnedGift() {}

// OwnedGiftUnique - Describes a unique gift received and owned by a user or a chat.
type OwnedGiftUnique struct {
	// Type - Type of the gift, always “unique”
	Type string `json:"type"`

	// Gift - Information about the unique gift
	Gift UniqueGift `json:"gift"`

	// OwnedGiftID - Optional. Unique identifier of the received gift for the bot; for gifts received on behalf
	// of business accounts only
	OwnedGiftID string `json:"owned_gift_id,omitempty"`

	// SenderUser - Optional. Sender of the gift if it is a known user
	SenderUser *User `json:"sender_user,omitempty"`

	// SendDate - Date the gift was sent in Unix time
	SendDate int64 `json:"send_date"`

	// IsSaved - Optional. True, if the gift is displayed on the account's profile page; for gifts received on
	// behalf of business accounts only
	IsSaved bool `json:"is_saved,omitempty"`

	// CanBeTransferred - Optional. True, if the gift can be transferred to another owner; for gifts received on
	// behalf of business accounts only
	CanBeTransferred bool `json:"can_be_transferred,omitempty"`

	// TransferStarCount - Optional. Number of Telegram Stars that must be paid to transfer the gift; omitted if
	// the bot cannot transfer the gift
	TransferStarCount int `json:"transfer_star_count,omitempty"`

	// NextTransferDate - Optional. Point in time (Unix timestamp) when the gift can be transferred. If it is in
	// the past, then the gift can be transferred now
	NextTransferDate int64 `json:"next_transfer_date,omitempty"`
}

// GiftType returns OwnedGift type
func (g *OwnedGiftUnique) GiftType() string {
	return GiftTypeUnique
}

func (g *OwnedGiftUnique) iOwnedGift() {}

// OwnedGifts - Contains the list of gifts received and owned by a user or a chat.
type OwnedGifts struct {
	// TotalCount - The total number of gifts owned by the user or the chat
	TotalCount int `json:"total_count"`

	// Gifts - The list of gifts
	Gifts []OwnedGift `json:"gifts"`

	// NextOffset - Optional. Offset for the next request. If empty, then there are no more results
	NextOffset string `json:"next_offset,omitempty"`
}

// UnmarshalJSON converts JSON to OwnedGifts
func (g *OwnedGifts) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	type uOwnedGifts OwnedGifts
	var ug uOwnedGifts

	if value.Exists("gifts") {
		gifts := value.GetArray("gifts")
		ug.Gifts = make([]OwnedGift, len(gifts))
		for i, gift := range gifts {
			giftType := string(gift.GetStringBytes("type"))
			switch giftType {
			case GiftTypeRegular:
				ug.Gifts[i] = &OwnedGiftRegular{}
			case GiftTypeUnique:
				ug.Gifts[i] = &OwnedGiftUnique{}
			default:
				return fmt.Errorf("unknown owned gift type: %q", giftType)
			}
		}
	}

	if err = json.Unmarshal(data, &ug); err != nil {
		return err
	}
	*g = OwnedGifts(ug)

	return nil
}

// AcceptedGiftTypes - This object describes the types of gifts that can be gifted to a user or a chat.
type AcceptedGiftTypes struct {
	// UnlimitedGifts - True, if unlimited regular gifts are accepted
	UnlimitedGifts bool `json:"unlimited_gifts"`

	// LimitedGifts - True, if limited regular gifts are accepted
	LimitedGifts bool `json:"limited_gifts"`

	// UniqueGifts - True, if unique gifts or gifts that can be upgraded to unique for free are accepted
	UniqueGifts bool `json:"unique_gifts"`

	// PremiumSubscription - True, if a Telegram Premium subscription is accepted
	PremiumSubscription bool `json:"premium_subscription"`

	// GiftsFromChannels - True, if transfers of unique gifts from channels are accepted
	GiftsFromChannels bool `json:"gifts_from_channels"`
}

// StarAmount - Describes an amount of Telegram Stars.
type StarAmount struct {
	// Amount - Integer amount of Telegram Stars, rounded to 0; can be negative
	Amount int `json:"amount"`

	// NanostarAmount - Optional. The number of 1/1000000000 shares of Telegram Stars; from -999999999 to
	// 999999999; can be negative if and only if amount is non-positive
	NanostarAmount int `json:"nanostar_amount,omitempty"`
}

// BotCommand - This object represents a bot command.
type BotCommand struct {
	// Command - Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and
	// underscores.
	Command string `json:"command"`

	// Description - Description of the command; 1-256 characters.
	Description string `json:"description"`
}

// ChatID - Represents chat ID as int64 or string
type ChatID struct { //nolint:recvcheck
	// ID - Unique identifier for the target chat
	ID int64

	// Username - Channel or group username of the target chat (in the format @channel_username)
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

// UnmarshalJSON parses JSON to ChatID
func (c *ChatID) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.ID); err == nil {
		return nil
	}
	return json.Unmarshal(data, &c.Username)
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
	// ScopeType returns BotCommandScope type
	ScopeType() string
	// Disallow external implementations
	iBotCommandScope()
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

func (b *BotCommandScopeDefault) iBotCommandScope() {}

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

func (b *BotCommandScopeAllPrivateChats) iBotCommandScope() {}

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

func (b *BotCommandScopeAllGroupChats) iBotCommandScope() {}

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

func (b *BotCommandScopeAllChatAdministrators) iBotCommandScope() {}

// BotCommandScopeChat - Represents the scope (https://core.telegram.org/bots/api#botcommandscope) of bot
// commands, covering a specific chat.
type BotCommandScopeChat struct {
	// Type - Scope type, must be chat
	Type string `json:"type"`

	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username). Channel direct messages chats and channel chats aren't supported.
	ChatID ChatID `json:"chat_id"`
}

// ScopeType returns BotCommandScope type
func (b *BotCommandScopeChat) ScopeType() string {
	return ScopeTypeChat
}

func (b *BotCommandScopeChat) iBotCommandScope() {}

// BotCommandScopeChatAdministrators - Represents the scope
// (https://core.telegram.org/bots/api#botcommandscope) of bot commands, covering all administrators of a
// specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	// Type - Scope type, must be chat_administrators
	Type string `json:"type"`

	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username). Channel direct messages chats and channel chats aren't supported.
	ChatID ChatID `json:"chat_id"`
}

// ScopeType returns BotCommandScope type
func (b *BotCommandScopeChatAdministrators) ScopeType() string {
	return ScopeTypeChatAdministrators
}

func (b *BotCommandScopeChatAdministrators) iBotCommandScope() {}

// BotCommandScopeChatMember - Represents the scope (https://core.telegram.org/bots/api#botcommandscope) of
// bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	// Type - Scope type, must be chat_member
	Type string `json:"type"`

	// ChatID - Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroup_username). Channel direct messages chats and channel chats aren't supported.
	ChatID ChatID `json:"chat_id"`

	// UserID - Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// ScopeType returns BotCommandScope type
func (b *BotCommandScopeChatMember) ScopeType() string {
	return ScopeTypeChatMember
}

func (b *BotCommandScopeChatMember) iBotCommandScope() {}

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
	// ButtonType returns MenuButton type
	ButtonType() string
	// Disallow external implementations
	iMenuButton()
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

// UnmarshalJSON converts JSON to menuButtonData
func (m *menuButtonData) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	if !value.Exists("type") {
		return errors.New(noTypeErr)
	}

	buttonType := string(value.GetStringBytes("type"))
	switch buttonType {
	case ButtonTypeCommands:
		var mb *MenuButtonCommands
		err = json.Unmarshal(data, &mb)
		m.Data = mb
	case ButtonTypeWebApp:
		var mb *MenuButtonWebApp
		err = json.Unmarshal(data, &mb)
		m.Data = mb
	case ButtonTypeDefault:
		var mb *MenuButtonDefault
		err = json.Unmarshal(data, &mb)
		m.Data = mb
	default:
		return fmt.Errorf("unknown menu button type: %q", buttonType)
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

func (m *MenuButtonCommands) iMenuButton() {}

// MenuButtonWebApp - Represents a menu button, which launches a Web App
// (https://core.telegram.org/bots/webapps).
type MenuButtonWebApp struct {
	// Type - Type of the button, must be web_app
	Type string `json:"type"`

	// Text - Text on the button
	Text string `json:"text"`

	// WebApp - Description of the Web App that will be launched when the user presses the button. The Web App
	// will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery
	// (https://core.telegram.org/bots/api#answerwebappquery). Alternatively, a t.me link to a Web App of the bot
	// can be specified in the object instead of the Web App's URL, in which case the Web App will be opened as if
	// the user pressed the link.
	WebApp WebAppInfo `json:"web_app"`
}

// ButtonType returns MenuButton type
func (m *MenuButtonWebApp) ButtonType() string {
	return ButtonTypeWebApp
}

func (m *MenuButtonWebApp) iMenuButton() {}

// MenuButtonDefault - Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {
	// Type - Type of the button, must be default
	Type string `json:"type"`
}

// ButtonType returns MenuButton type
func (m *MenuButtonDefault) ButtonType() string {
	return ButtonTypeDefault
}

func (m *MenuButtonDefault) iMenuButton() {}

// ChatBoostSource - This object describes the source of a chat boost. It can be one of
// ChatBoostSourcePremium (https://core.telegram.org/bots/api#chatboostsourcepremium)
// ChatBoostSourceGiftCode (https://core.telegram.org/bots/api#chatboostsourcegiftcode)
// ChatBoostSourceGiveaway (https://core.telegram.org/bots/api#chatboostsourcegiveaway)
type ChatBoostSource interface {
	// BoostSource returns boost source
	BoostSource() string
	// Disallow external implementations
	iChatBoostSource()
}

// Boost sources
const (
	BoostSourcePremium  = "premium"
	BoostSourceGiftCode = "gift_code"
	BoostSourceGiveaway = "giveaway"
)

// ChatBoostSourcePremium - The boost was obtained by subscribing to Telegram Premium or by gifting a
// Telegram Premium subscription to another user.
type ChatBoostSourcePremium struct {
	// Source - Source of the boost, always “premium”
	Source string `json:"source"`

	// User - User that boosted the chat
	User User `json:"user"`
}

// BoostSource returns boost source
func (b *ChatBoostSourcePremium) BoostSource() string {
	return BoostSourcePremium
}

func (b *ChatBoostSourcePremium) iChatBoostSource() {}

// ChatBoostSourceGiftCode - The boost was obtained by the creation of Telegram Premium gift codes to boost a
// chat. Each such code boosts the chat 4 times for the duration of the corresponding Telegram Premium
// subscription.
type ChatBoostSourceGiftCode struct {
	// Source - Source of the boost, always “gift_code”
	Source string `json:"source"`

	// User - User for which the gift code was created
	User User `json:"user"`
}

// BoostSource returns boost source
func (b *ChatBoostSourceGiftCode) BoostSource() string {
	return BoostSourceGiftCode
}

func (b *ChatBoostSourceGiftCode) iChatBoostSource() {}

// ChatBoostSourceGiveaway - The boost was obtained by the creation of a Telegram Premium or a Telegram Star
// giveaway. This boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription
// for Telegram Premium giveaways and prize_star_count / 500 times for one year for Telegram Star giveaways.
type ChatBoostSourceGiveaway struct {
	// Source - Source of the boost, always “giveaway”
	Source string `json:"source"`

	// GiveawayMessageID - Identifier of a message in the chat with the giveaway; the message could have been
	// deleted already. May be 0 if the message isn't sent yet.
	GiveawayMessageID int `json:"giveaway_message_id"`

	// User - Optional. User that won the prize in the giveaway if any; for Telegram Premium giveaways only
	User *User `json:"user,omitempty"`

	// PrizeStarCount - Optional. The number of Telegram Stars to be split between giveaway winners; for
	// Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count,omitempty"`

	// IsUnclaimed - Optional. True, if the giveaway was completed, but there was no user to win the prize
	IsUnclaimed bool `json:"is_unclaimed,omitempty"`
}

// BoostSource returns boost source
func (b *ChatBoostSourceGiveaway) BoostSource() string {
	return BoostSourceGiveaway
}

func (b *ChatBoostSourceGiveaway) iChatBoostSource() {}

// ChatBoost - This object contains information about a chat boost.
type ChatBoost struct {
	// BoostID - Unique identifier of the boost
	BoostID string `json:"boost_id"`

	// AddDate - Point in time (Unix timestamp) when the chat was boosted
	AddDate int64 `json:"add_date"`

	// ExpirationDate - Point in time (Unix timestamp) when the boost will automatically expire, unless the
	// booster's Telegram Premium subscription is prolonged
	ExpirationDate int64 `json:"expiration_date"`

	// Source - Source of the added boost
	Source ChatBoostSource `json:"source"`
}

// UnmarshalJSON converts JSON to ChatBoost
func (b *ChatBoost) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	if !value.Exists("source") {
		return errors.New("no source")
	}

	type uChatBoost ChatBoost
	var ub uChatBoost

	source := string(value.GetStringBytes("source", "source"))
	switch source {
	case BoostSourcePremium:
		ub.Source = &ChatBoostSourcePremium{}
	case BoostSourceGiftCode:
		ub.Source = &ChatBoostSourceGiftCode{}
	case BoostSourceGiveaway:
		ub.Source = &ChatBoostSourceGiveaway{}
	default:
		return fmt.Errorf("unknown chat boost source: %q", source)
	}

	if err = json.Unmarshal(data, &ub); err != nil {
		return err
	}
	*b = ChatBoost(ub)

	return nil
}

// ChatBoostUpdated - This object represents a boost added to a chat or changed.
type ChatBoostUpdated struct {
	// Chat - Chat which was boosted
	Chat Chat `json:"chat"`

	// Boost - Information about the chat boost
	Boost ChatBoost `json:"boost"`
}

// ChatBoostRemoved - This object represents a boost removed from a chat.
type ChatBoostRemoved struct {
	// Chat - Chat which was boosted
	Chat Chat `json:"chat"`

	// BoostID - Unique identifier of the boost
	BoostID string `json:"boost_id"`

	// RemoveDate - Point in time (Unix timestamp) when the boost was removed
	RemoveDate int64 `json:"remove_date"`

	// Source - Source of the removed boost
	Source ChatBoostSource `json:"source"`
}

// UnmarshalJSON converts JSON to ChatBoostRemoved
func (b *ChatBoostRemoved) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	if !value.Exists("source") {
		return errors.New("no source")
	}

	type uChatBoostRemoved ChatBoostRemoved
	var ub uChatBoostRemoved

	source := string(value.GetStringBytes("source", "source"))
	switch source {
	case BoostSourcePremium:
		ub.Source = &ChatBoostSourcePremium{}
	case BoostSourceGiftCode:
		ub.Source = &ChatBoostSourceGiftCode{}
	case BoostSourceGiveaway:
		ub.Source = &ChatBoostSourceGiveaway{}
	default:
		return fmt.Errorf("unknown chat boost source: %q", source)
	}

	if err = json.Unmarshal(data, &ub); err != nil {
		return err
	}
	*b = ChatBoostRemoved(ub)

	return nil
}

// ChatOwnerLeft - Describes a service message about the chat owner leaving the chat.
type ChatOwnerLeft struct {
	// NewOwner - Optional. The user which will be the new owner of the chat if the previous owner does not
	// return to the chat
	NewOwner *User `json:"new_owner,omitempty"`
}

// ChatOwnerChanged - Describes a service message about an ownership change in the chat.
type ChatOwnerChanged struct {
	// NewOwner - The new owner of the chat
	NewOwner User `json:"new_owner"`
}

// UserChatBoosts - This object represents a list of boosts added to a chat by a user.
type UserChatBoosts struct {
	// Boosts - The list of boosts added to the chat by the user
	Boosts []ChatBoost `json:"boosts"`
}

// BusinessBotRights - Represents the rights of a business bot.
type BusinessBotRights struct {
	// CanReply - Optional. True, if the bot can send and edit messages in the private chats that had incoming
	// messages in the last 24 hours
	CanReply bool `json:"can_reply,omitempty"`

	// CanReadMessages - Optional. True, if the bot can mark incoming private messages as read
	CanReadMessages bool `json:"can_read_messages,omitempty"`

	// CanDeleteSentMessages - Optional. True, if the bot can delete messages sent by the bot
	CanDeleteSentMessages bool `json:"can_delete_sent_messages,omitempty"`

	// CanDeleteAllMessages - Optional. True, if the bot can delete all private messages in managed chats
	CanDeleteAllMessages bool `json:"can_delete_all_messages,omitempty"`

	// CanEditName - Optional. True, if the bot can edit the first and last name of the business account
	CanEditName bool `json:"can_edit_name,omitempty"`

	// CanEditBio - Optional. True, if the bot can edit the bio of the business account
	CanEditBio bool `json:"can_edit_bio,omitempty"`

	// CanEditProfilePhoto - Optional. True, if the bot can edit the profile photo of the business account
	CanEditProfilePhoto bool `json:"can_edit_profile_photo,omitempty"`

	// CanEditUsername - Optional. True, if the bot can edit the username of the business account
	CanEditUsername bool `json:"can_edit_username,omitempty"`

	// CanChangeGiftSettings - Optional. True, if the bot can change the privacy settings pertaining to gifts
	// for the business account
	CanChangeGiftSettings bool `json:"can_change_gift_settings,omitempty"`

	// CanViewGiftsAndStars - Optional. True, if the bot can view gifts and the amount of Telegram Stars owned
	// by the business account
	CanViewGiftsAndStars bool `json:"can_view_gifts_and_stars,omitempty"`

	// CanConvertGiftsToStars - Optional. True, if the bot can convert regular gifts owned by the business
	// account to Telegram Stars
	CanConvertGiftsToStars bool `json:"can_convert_gifts_to_stars,omitempty"`

	// CanTransferAndUpgradeGifts - Optional. True, if the bot can transfer and upgrade gifts owned by the
	// business account
	CanTransferAndUpgradeGifts bool `json:"can_transfer_and_upgrade_gifts,omitempty"`

	// CanTransferStars - Optional. True, if the bot can transfer Telegram Stars received by the business
	// account to its own account, or use them to upgrade and transfer gifts
	CanTransferStars bool `json:"can_transfer_stars,omitempty"`

	// CanManageStories - Optional. True, if the bot can post, edit and delete stories on behalf of the business
	// account
	CanManageStories bool `json:"can_manage_stories,omitempty"`
}

// BusinessConnection - Describes the connection of the bot with a business account.
type BusinessConnection struct {
	// ID - Unique identifier of the business connection
	ID string `json:"id"`

	// User - Business account user that created the business connection
	User User `json:"user"`

	// UserChatID - Identifier of a private chat with the user who created the business connection. This number
	// may have more than 32 significant bits and some programming languages may have difficulty/silent defects in
	// interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type
	// are safe for storing this identifier.
	UserChatID int64 `json:"user_chat_id"`

	// Date - Date the connection was established in Unix time
	Date int64 `json:"date"`

	// Rights - Optional. Rights of the business bot
	Rights *BusinessBotRights `json:"rights,omitempty"`

	// IsEnabled - True, if the connection is active
	IsEnabled bool `json:"is_enabled"`
}

// BusinessMessagesDeleted - This object is received when messages are deleted from a connected business
// account.
type BusinessMessagesDeleted struct {
	// BusinessConnectionID - Unique identifier of the business connection
	BusinessConnectionID string `json:"business_connection_id"`

	// Chat - Information about a chat in the business account. The bot may not have access to the chat or the
	// corresponding user.
	Chat Chat `json:"chat"`

	// MessageIDs - The list of identifiers of deleted messages in the chat of the business account
	MessageIDs []int `json:"message_ids"`
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
	// MediaType return InputMedia type
	MediaType() string
	// Disallow external implementations
	iInputMedia()
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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// HasSpoiler - Optional. Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// MediaType return InputMedia type
func (i *InputMediaPhoto) MediaType() string {
	return MediaTypePhoto
}

func (i *InputMediaPhoto) iInputMedia() {}

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

	// Cover - Optional. Cover for the video in the message. Pass a file_id to send a file that exists on the
	// Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name>
	// name. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Cover *InputFile `json:"cover,omitempty"`

	// StartTimestamp - Optional. Start timestamp for the video in the message
	StartTimestamp int `json:"start_timestamp,omitempty"`

	// Caption - Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// ParseMode - Optional. Mode for parsing entities in the video caption. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// CaptionEntities - Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

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

func (i *InputMediaVideo) iInputMedia() {}

func (i *InputMediaVideo) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	i.Media.needAttach = true
	fp["media"] = i.Media.File
	if i.Thumbnail != nil {
		i.Thumbnail.needAttach = true
		fp["thumbnail"] = i.Thumbnail.File
	}
	if i.Cover != nil {
		i.Cover.needAttach = true
		fp["cover"] = i.Cover.File
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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

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

func (i *InputMediaAnimation) iInputMedia() {}

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

func (i *InputMediaAudio) iInputMedia() {}

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

func (i *InputMediaDocument) iInputMedia() {}

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

// InputPaidMedia - This object describes the paid media to be sent. Currently, it can be one of
// InputPaidMediaPhoto (https://core.telegram.org/bots/api#inputpaidmediaphoto)
// InputPaidMediaVideo (https://core.telegram.org/bots/api#inputpaidmediavideo)
type InputPaidMedia interface {
	// MediaType returns InputPaidMedia type
	MediaType() string
	// MediaFile returns InputPaidMedia file
	MediaFile() InputFile
	// Disallow external implementations
	iInputPaidMedia()
	fileCompatible
}

// InputPaidMediaPhoto - The paid media to send is a photo.
type InputPaidMediaPhoto struct {
	// Type - Type of the media, must be photo
	Type string `json:"type"`

	// Media - File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to
	// upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files
	// » (https://core.telegram.org/bots/api#sending-files)
	Media InputFile `json:"media"`
}

// MediaType returns InputPaidMedia type
func (i *InputPaidMediaPhoto) MediaType() string {
	return PaidMediaTypePhoto
}

// MediaFile returns InputPaidMedia file
func (i *InputPaidMediaPhoto) MediaFile() InputFile {
	return i.Media
}

func (i *InputPaidMediaPhoto) iInputPaidMedia() {}

func (i *InputPaidMediaPhoto) fileParameters() map[string]telegoapi.NamedReader {
	i.Media.needAttach = true
	return map[string]telegoapi.NamedReader{
		"media": i.Media.File,
	}
}

// InputPaidMediaVideo - The paid media to send is a video.
type InputPaidMediaVideo struct {
	// Type - Type of the media, must be video
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

	// Cover - Optional. Cover for the video in the message. Pass a file_id to send a file that exists on the
	// Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name>
	// name. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Cover *InputFile `json:"cover,omitempty"`

	// StartTimestamp - Optional. Start timestamp for the video in the message
	StartTimestamp int `json:"start_timestamp,omitempty"`

	// Width - Optional. Video width
	Width int `json:"width,omitempty"`

	// Height - Optional. Video height
	Height int `json:"height,omitempty"`

	// Duration - Optional. Video duration in seconds
	Duration int `json:"duration,omitempty"`

	// SupportsStreaming - Optional. Pass True if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
}

// MediaType returns InputPaidMedia type
func (i *InputPaidMediaVideo) MediaType() string {
	return PaidMediaTypeVideo
}

// MediaFile returns InputPaidMedia file
func (i *InputPaidMediaVideo) MediaFile() InputFile {
	return i.Media
}

func (i *InputPaidMediaVideo) iInputPaidMedia() {}

func (i *InputPaidMediaVideo) fileParameters() map[string]telegoapi.NamedReader {
	fp := make(map[string]telegoapi.NamedReader)

	i.Media.needAttach = true
	fp["media"] = i.Media.File
	if i.Thumbnail != nil {
		i.Thumbnail.needAttach = true
		fp["thumbnail"] = i.Thumbnail.File
	}
	if i.Cover != nil {
		i.Cover.needAttach = true
		fp["cover"] = i.Cover.File
	}

	return fp
}

// InputProfilePhoto - This object describes a profile photo to set. Currently, it can be one of
// InputProfilePhotoStatic (https://core.telegram.org/bots/api#inputprofilephotostatic)
// InputProfilePhotoAnimated (https://core.telegram.org/bots/api#inputprofilephotoanimated)
type InputProfilePhoto interface {
	// ProfilePhotoType return InputProfilePhoto type
	ProfilePhotoType() string
	// Disallow external implementations
	iInputProfilePhoto()
}

// Input profile photo types
const (
	PhotoTypeStatic   = "static"
	PhotoTypeAnimated = "animated"
)

// InputProfilePhotoStatic - A static profile photo in the .JPG format.
type InputProfilePhotoStatic struct {
	// Type - Type of the profile photo, must be static
	Type string `json:"type"`

	// Photo - The static profile photo. Profile photos can't be reused and can only be uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the photo was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Photo InputFile `json:"photo"`
}

// ProfilePhotoType return InputProfilePhoto type
func (i *InputProfilePhotoStatic) ProfilePhotoType() string {
	return PhotoTypeStatic
}

func (i *InputProfilePhotoStatic) iInputProfilePhoto() {}

// InputProfilePhotoAnimated - An animated profile photo in the MPEG4 format.
type InputProfilePhotoAnimated struct {
	// Type - Type of the profile photo, must be animated
	Type string `json:"type"`

	// Animation - The animated profile photo. Profile photos can't be reused and can only be uploaded as a new
	// file, so you can pass “attach://<file_attach_name>” if the photo was uploaded using multipart/form-data
	// under <file_attach_name>. More information on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	Animation InputFile `json:"animation"`

	// MainFrameTimestamp - Optional. Timestamp in seconds of the frame that will be used as the static profile
	// photo. Defaults to 0.0.
	MainFrameTimestamp float64 `json:"main_frame_timestamp,omitempty"`
}

// ProfilePhotoType return InputProfilePhoto type
func (i *InputProfilePhotoAnimated) ProfilePhotoType() string {
	return PhotoTypeAnimated
}

func (i *InputProfilePhotoAnimated) iInputProfilePhoto() {}

// InputStoryContent - This object describes the content of a story to post. Currently, it can be one of
// InputStoryContentPhoto (https://core.telegram.org/bots/api#inputstorycontentphoto)
// InputStoryContentVideo (https://core.telegram.org/bots/api#inputstorycontentvideo)
type InputStoryContent interface {
	// StoryType return InputStoryContent type
	StoryType() string
	// Disallow external implementations
	iInputStoryContent()
}

// Story types
const (
	StoryTypePhoto = "photo"
	StoryTypeVideo = "video"
)

// InputStoryContentPhoto - Describes a photo to post as a story.
type InputStoryContentPhoto struct {
	// Type - Type of the content, must be photo
	Type string `json:"type"`

	// Photo - The photo to post as a story. The photo must be of the size 1080x1920 and must not exceed 10 MB.
	// The photo can't be reused and can only be uploaded as a new file, so you can pass
	// “attach://<file_attach_name>” if the photo was uploaded using multipart/form-data under
	// <file_attach_name>. More information on Sending Files » (https://core.telegram.org/bots/api#sending-files)
	Photo InputFile `json:"photo"`
}

// StoryType return InputStoryContent type
func (i *InputStoryContentPhoto) StoryType() string {
	return StoryTypePhoto
}

func (i *InputStoryContentPhoto) iInputStoryContent() {}

// InputStoryContentVideo - Describes a video to post as a story.
type InputStoryContentVideo struct {
	// Type - Type of the content, must be video
	Type string `json:"type"`

	// Video - The video to post as a story. The video must be of the size 720x1280, streamable, encoded with
	// H.265 codec, with key frames added each second in the MPEG4 format, and must not exceed 30 MB. The video
	// can't be reused and can only be uploaded as a new file, so you can pass “attach://<file_attach_name>” if
	// the video was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files
	// » (https://core.telegram.org/bots/api#sending-files)
	Video InputFile `json:"video"`

	// Duration - Optional. Precise duration of the video in seconds; 0-60
	Duration float64 `json:"duration,omitempty"`

	// CoverFrameTimestamp - Optional. Timestamp in seconds of the frame that will be used as the static cover
	// for the story. Defaults to 0.0.
	CoverFrameTimestamp float64 `json:"cover_frame_timestamp,omitempty"`

	// IsAnimation - Optional. Pass True if the video has no sound
	IsAnimation bool `json:"is_animation,omitempty"`
}

// StoryType return InputStoryContent type
func (i *InputStoryContentVideo) StoryType() string {
	return StoryTypeVideo
}

func (i *InputStoryContentVideo) iInputStoryContent() {}

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
	// Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new file using multipart/form-data under <file_attach_name>
	// name. Animated and video stickers can't be uploaded via HTTP URL. More information on Sending Files »
	// (https://core.telegram.org/bots/api#sending-files)
	Sticker InputFile `json:"sticker"`

	// Format - Format of the added sticker, must be one of “static” for a .WEBP or .PNG image,
	// “animated” for a .TGS animation, “video” for a .WEBM video
	Format string `json:"format"`

	// EmojiList - List of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list"`

	// MaskPosition - Optional. Position where the mask should be placed on faces. For “mask” stickers only.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// Keywords - Optional. List of 0-20 search keywords for the sticker with total length of up to 64
	// characters. For “regular” and “custom_emoji” stickers only.
	Keywords []string `json:"keywords,omitempty"`
}

// Sticker formats
const (
	StickerStatic   = "static"
	StickerAnimated = "animated"
	StickerVideo    = "video"
)

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
	// ResultType returns InlineQueryResult type
	ResultType() string
	// Disallow external implementations
	iInlineQueryResult()
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

func (i *InlineQueryResultArticle) iInlineQueryResult() {}

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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

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

func (i *InlineQueryResultPhoto) iInlineQueryResult() {}

// InlineQueryResultGif - Represents a link to an animated GIF file. By default, this animated GIF file will
// be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message
// with the specified content instead of the animation.
type InlineQueryResultGif struct {
	// Type - Type of the result, must be gif
	Type string `json:"type"`

	// ID - Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// GifURL - A valid URL for the GIF file
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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

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

func (i *InlineQueryResultGif) iInlineQueryResult() {}

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

	// Mpeg4URL - A valid URL for the MPEG4 file
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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

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

func (i *InlineQueryResultMpeg4Gif) iInlineQueryResult() {}

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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

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

func (i *InlineQueryResultVideo) iInlineQueryResult() {}

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

func (i *InlineQueryResultAudio) iInlineQueryResult() {}

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

func (i *InlineQueryResultVoice) iInlineQueryResult() {}

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

	// ReplyMarkup - Optional. Inline keyboard (https://core.telegram.org/bots/features#inline-keyboards)
	// attached to the message
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

func (i *InlineQueryResultDocument) iInlineQueryResult() {}

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

	// LivePeriod - Optional. Period in seconds during which the location can be updated, should be between 60
	// and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
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

func (i *InlineQueryResultLocation) iInlineQueryResult() {}

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

func (i *InlineQueryResultVenue) iInlineQueryResult() {}

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

func (i *InlineQueryResultContact) iInlineQueryResult() {}

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

func (i *InlineQueryResultGame) iInlineQueryResult() {}

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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

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

func (i *InlineQueryResultCachedPhoto) iInlineQueryResult() {}

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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

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

func (i *InlineQueryResultCachedGif) iInlineQueryResult() {}

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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

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

func (i *InlineQueryResultCachedMpeg4Gif) iInlineQueryResult() {}

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

func (i *InlineQueryResultCachedSticker) iInlineQueryResult() {}

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

func (i *InlineQueryResultCachedDocument) iInlineQueryResult() {}

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

	// ShowCaptionAboveMedia - Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

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

func (i *InlineQueryResultCachedVideo) iInlineQueryResult() {}

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

func (i *InlineQueryResultCachedVoice) iInlineQueryResult() {}

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

func (i *InlineQueryResultCachedAudio) iInlineQueryResult() {}

// InputMessageContent - This object represents the content of a message to be sent as a result of an inline
// query. Telegram clients currently support the following 5 types:
// InputTextMessageContent (https://core.telegram.org/bots/api#inputtextmessagecontent)
// InputLocationMessageContent (https://core.telegram.org/bots/api#inputlocationmessagecontent)
// InputVenueMessageContent (https://core.telegram.org/bots/api#inputvenuemessagecontent)
// InputContactMessageContent (https://core.telegram.org/bots/api#inputcontactmessagecontent)
// InputInvoiceMessageContent (https://core.telegram.org/bots/api#inputinvoicemessagecontent)
type InputMessageContent interface {
	// ContentType returns InputMessageContent type
	ContentType() string
	// Disallow external implementations
	iInputMessageContent()
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

	// LinkPreviewOptions - Optional. Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

// ContentType returns InputMessageContent type
func (i *InputTextMessageContent) ContentType() string {
	return ContentTypeText
}

func (i *InputTextMessageContent) iInputMessageContent() {}

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

	// LivePeriod - Optional. Period in seconds during which the location can be updated, should be between 60
	// and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
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

func (i *InputLocationMessageContent) iInputMessageContent() {}

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

func (i *InputVenueMessageContent) iInputMessageContent() {}

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

func (i *InputContactMessageContent) iInputMessageContent() {}

// InputInvoiceMessageContent - Represents the content
// (https://core.telegram.org/bots/api#inputmessagecontent) of an invoice message to be sent as the result of an
// inline query.
type InputInvoiceMessageContent struct {
	// Title - Product name, 1-32 characters
	Title string `json:"title"`

	// Description - Product description, 1-255 characters
	Description string `json:"description"`

	// Payload - Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for
	// your internal processes.
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

// ContentType returns InputMessageContent type
func (i *InputInvoiceMessageContent) ContentType() string {
	return ContentTypeInvoice
}

func (i *InputInvoiceMessageContent) iInputMessageContent() {}

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

// PreparedInlineMessage - Describes an inline message to be sent by a user of a Mini App.
type PreparedInlineMessage struct {
	// ID - Unique identifier of the prepared message
	ID string `json:"id"`

	// ExpirationDate - Expiration date of the prepared message, in Unix time. Expired prepared messages can no
	// longer be used
	ExpirationDate int64 `json:"expiration_date"`
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
	// code, or “XTR” for payments in Telegram Stars (https://t.me/BotNews/90)
	Currency string `json:"currency"`

	// TotalAmount - Total price in the smallest units of the currency (integer, not float/double). For example,
	// for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal
	// point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}

// ShippingAddress - This object represents a shipping address.
type ShippingAddress struct {
	// CountryCode - Two-letter ISO 3166-1 alpha-2 (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country
	// code
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

// SuccessfulPayment - This object contains basic information about a successful payment. Note that if the
// buyer initiates a chargeback with the relevant payment provider following this transaction, the funds may be
// debited from your balance. This is outside of Telegram's control.
type SuccessfulPayment struct {
	// Currency - Three-letter ISO 4217 currency (https://core.telegram.org/bots/payments#supported-currencies)
	// code, or “XTR” for payments in Telegram Stars (https://t.me/BotNews/90)
	Currency string `json:"currency"`

	// TotalAmount - Total price in the smallest units of the currency (integer, not float/double). For example,
	// for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal
	// point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload - Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// SubscriptionExpirationDate - Optional. Expiration date of the subscription, in Unix time; for recurring
	// payments only
	SubscriptionExpirationDate int64 `json:"subscription_expiration_date,omitempty"`

	// IsRecurring - Optional. True, if the payment is a recurring payment for a subscription
	IsRecurring bool `json:"is_recurring,omitempty"`

	// IsFirstRecurring - Optional. True, if the payment is the first payment for a subscription
	IsFirstRecurring bool `json:"is_first_recurring,omitempty"`

	// ShippingOptionID - Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo - Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`

	// TelegramPaymentChargeID - Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

	// ProviderPaymentChargeID - Provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}

// RefundedPayment - This object contains basic information about a refunded payment.
type RefundedPayment struct {
	// Currency - Three-letter ISO 4217 currency (https://core.telegram.org/bots/payments#supported-currencies)
	// code, or “XTR” for payments in Telegram Stars (https://t.me/BotNews/90). Currently, always “XTR”
	Currency string `json:"currency"`

	// TotalAmount - Total refunded price in the smallest units of the currency (integer, not float/double). For
	// example, for a price of US$ 1.45, total_amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal
	// point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload - Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// TelegramPaymentChargeID - Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

	// ProviderPaymentChargeID - Optional. Provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`
}

// ShippingQuery - This object contains information about an incoming shipping query.
type ShippingQuery struct {
	// ID - Unique query identifier
	ID string `json:"id"`

	// From - User who sent the query
	From User `json:"from"`

	// InvoicePayload - Bot-specified invoice payload
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
	// code, or “XTR” for payments in Telegram Stars (https://t.me/BotNews/90)
	Currency string `json:"currency"`

	// TotalAmount - Total price in the smallest units of the currency (integer, not float/double). For example,
	// for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal
	// point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload - Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// ShippingOptionID - Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo - Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

// PaidMediaPurchased - This object contains information about a paid media purchase.
type PaidMediaPurchased struct {
	// From - User who purchased the media
	From User `json:"from"`

	// PaidMediaPayload - Bot-specified paid media payload
	PaidMediaPayload string `json:"paid_media_payload"`
}

// RevenueWithdrawalState - This object describes the state of a revenue withdrawal operation. Currently, it
// can be one of
// RevenueWithdrawalStatePending (https://core.telegram.org/bots/api#revenuewithdrawalstatepending)
// RevenueWithdrawalStateSucceeded (https://core.telegram.org/bots/api#revenuewithdrawalstatesucceeded)
// RevenueWithdrawalStateFailed (https://core.telegram.org/bots/api#revenuewithdrawalstatefailed)
type RevenueWithdrawalState interface {
	// WithdrawalState returns RevenueWithdrawalState type
	WithdrawalState() string
	// Disallow external implementations
	iRevenueWithdrawalState()
}

// Revenue withdrawal state types
const (
	WithdrawalStatePending   = "pending"
	WithdrawalStateSucceeded = "succeeded"
	WithdrawalStateFailed    = "failed"
)

// RevenueWithdrawalStatePending - The withdrawal is in progress.
type RevenueWithdrawalStatePending struct {
	// Type - Type of the state, always “pending”
	Type string `json:"type"`
}

// WithdrawalState returns RevenueWithdrawalState type
func (r *RevenueWithdrawalStatePending) WithdrawalState() string {
	return WithdrawalStatePending
}

func (r *RevenueWithdrawalStatePending) iRevenueWithdrawalState() {}

// RevenueWithdrawalStateSucceeded - The withdrawal succeeded.
type RevenueWithdrawalStateSucceeded struct {
	// Type - Type of the state, always “succeeded”
	Type string `json:"type"`

	// Date - Date the withdrawal was completed in Unix time
	Date int64 `json:"date"`

	// URL - An HTTPS URL that can be used to see transaction details
	URL string `json:"url"`
}

// WithdrawalState returns RevenueWithdrawalState type
func (r *RevenueWithdrawalStateSucceeded) WithdrawalState() string {
	return WithdrawalStateSucceeded
}

func (r *RevenueWithdrawalStateSucceeded) iRevenueWithdrawalState() {}

// RevenueWithdrawalStateFailed - The withdrawal failed and the transaction was refunded.
type RevenueWithdrawalStateFailed struct {
	// Type - Type of the state, always “failed”
	Type string `json:"type"`
}

// WithdrawalState returns RevenueWithdrawalState type
func (r *RevenueWithdrawalStateFailed) WithdrawalState() string {
	return WithdrawalStateFailed
}

func (r *RevenueWithdrawalStateFailed) iRevenueWithdrawalState() {}

// AffiliateInfo - Contains information about the affiliate that received a commission via this transaction.
type AffiliateInfo struct {
	// AffiliateUser - Optional. The bot or the user that received an affiliate commission if it was received by
	// a bot or a user
	AffiliateUser *User `json:"affiliate_user,omitempty"`

	// AffiliateChat - Optional. The chat that received an affiliate commission if it was received by a chat
	AffiliateChat *Chat `json:"affiliate_chat,omitempty"`

	// CommissionPerMille - The number of Telegram Stars received by the affiliate for each 1000 Telegram Stars
	// received by the bot from referred users
	CommissionPerMille int `json:"commission_per_mille"`

	// Amount - Integer amount of Telegram Stars received by the affiliate from the transaction, rounded to 0;
	// can be negative for refunds
	Amount int `json:"amount"`

	// NanostarAmount - Optional. The number of 1/1000000000 shares of Telegram Stars received by the affiliate;
	// from -999999999 to 999999999; can be negative for refunds
	NanostarAmount int `json:"nanostar_amount,omitempty"`
}

// TransactionPartner - This object describes the source of a transaction, or its recipient for outgoing
// transactions. Currently, it can be one of
// TransactionPartnerUser (https://core.telegram.org/bots/api#transactionpartneruser)
// TransactionPartnerChat (https://core.telegram.org/bots/api#transactionpartnerchat)
// TransactionPartnerAffiliateProgram (https://core.telegram.org/bots/api#transactionpartneraffiliateprogram)
// TransactionPartnerFragment (https://core.telegram.org/bots/api#transactionpartnerfragment)
// TransactionPartnerTelegramAds (https://core.telegram.org/bots/api#transactionpartnertelegramads)
// TransactionPartnerTelegramApi (https://core.telegram.org/bots/api#transactionpartnertelegramapi)
// TransactionPartnerOther (https://core.telegram.org/bots/api#transactionpartnerother)
type TransactionPartner interface {
	// PartnerType returns TransactionPartner type
	PartnerType() string
	// Disallow external implementations
	iTransactionPartner()
}

// Transaction partner types
const (
	PartnerTypeUser             = "user"
	PartnerTypeChat             = "chat"
	PartnerTypeAffiliateProgram = "affiliate_program"
	PartnerTypeFragment         = "fragment"
	PartnerTypeTelegramAds      = "telegram_ads"
	PartnerTypeTelegramApi      = "telegram_api" //nolint:revive
	PartnerTypeOther            = "other"
)

// TransactionPartnerUser - Describes a transaction with a user.
type TransactionPartnerUser struct {
	// Type - Type of the transaction partner, always “user”
	Type string `json:"type"`

	// TransactionType - Type of the transaction, currently one of “invoice_payment” for payments via
	// invoices, “paid_media_payment” for payments for paid media, “gift_purchase” for gifts sent by the
	// bot, “premium_purchase” for Telegram Premium subscriptions gifted by the bot,
	// “business_account_transfer” for direct transfers from managed business accounts
	TransactionType string `json:"transaction_type"`

	// User - Information about the user
	User User `json:"user"`

	// Affiliate - Optional. Information about the affiliate that received a commission via this transaction.
	// Can be available only for “invoice_payment” and “paid_media_payment” transactions.
	Affiliate *AffiliateInfo `json:"affiliate,omitempty"`

	// InvoicePayload - Optional. Bot-specified invoice payload. Can be available only for “invoice_payment”
	// transactions.
	InvoicePayload string `json:"invoice_payload,omitempty"`

	// SubscriptionPeriod - Optional. The duration of the paid subscription. Can be available only for
	// “invoice_payment” transactions.
	SubscriptionPeriod int `json:"subscription_period,omitempty"`

	// PaidMedia - Optional. Information about the paid media bought by the user; for “paid_media_payment”
	// transactions only
	PaidMedia []PaidMedia `json:"paid_media,omitempty"`

	// PaidMediaPayload - Optional. Bot-specified paid media payload. Can be available only for
	// “paid_media_payment” transactions.
	PaidMediaPayload string `json:"paid_media_payload,omitempty"`

	// Gift - Optional. The gift sent to the user by the bot; for “gift_purchase” transactions only
	Gift *Gift `json:"gift,omitempty"`

	// PremiumSubscriptionDuration - Optional. Number of months the gifted Telegram Premium subscription will be
	// active for; for “premium_purchase” transactions only
	PremiumSubscriptionDuration int `json:"premium_subscription_duration,omitempty"`
}

// Transaction types
const (
	TransactionTypeInvoicePayment          = "invoice_payment"
	TransactionTypePaidMediaPayment        = "paid_media_payment"
	TransactionTypeGiftPurchase            = "gift_purchase"
	TransactionTypePremiumPurchase         = "premium_purchase"
	TransactionTypeBusinessAccountTransfer = "business_account_transfer"
)

// PartnerType returns TransactionPartner type
func (p *TransactionPartnerUser) PartnerType() string {
	return PartnerTypeUser
}

func (p *TransactionPartnerUser) iTransactionPartner() {}

// UnmarshalJSON converts JSON to PaidMediaInfo
func (p *TransactionPartnerUser) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	type uTransactionPartnerUser TransactionPartnerUser
	var up uTransactionPartnerUser

	if value.Exists("paid_media") {
		paidMedia := value.GetArray("paid_media")
		up.PaidMedia = make([]PaidMedia, len(paidMedia))
		for i, media := range paidMedia {
			mediaType := string(media.GetStringBytes("type"))
			switch mediaType {
			case PaidMediaTypePreview:
				up.PaidMedia[i] = &PaidMediaPreview{}
			case PaidMediaTypePhoto:
				up.PaidMedia[i] = &PaidMediaPhoto{}
			case PaidMediaTypeVideo:
				up.PaidMedia[i] = &PaidMediaVideo{}
			case paidMediaTypeOther:
				up.PaidMedia[i] = &paidMediaOther{}
			default:
				return fmt.Errorf("unknown paid media type: %q", mediaType)
			}
		}
	}

	if err = json.Unmarshal(data, &up); err != nil {
		return err
	}
	*p = TransactionPartnerUser(up)

	return nil
}

// TransactionPartnerChat - Describes a transaction with a chat.
type TransactionPartnerChat struct {
	// Type - Type of the transaction partner, always “chat”
	Type string `json:"type"`

	// Chat - Information about the chat
	Chat Chat `json:"chat"`

	// Gift - Optional. The gift sent to the chat by the bot
	Gift *Gift `json:"gift,omitempty"`
}

// PartnerType returns TransactionPartner type
func (p *TransactionPartnerChat) PartnerType() string {
	return PartnerTypeChat
}

func (p *TransactionPartnerChat) iTransactionPartner() {}

// TransactionPartnerAffiliateProgram - Describes the affiliate program that issued the affiliate commission
// received via this transaction.
type TransactionPartnerAffiliateProgram struct {
	// Type - Type of the transaction partner, always “affiliate_program”
	Type string `json:"type"`

	// SponsorUser - Optional. Information about the bot that sponsored the affiliate program
	SponsorUser *User `json:"sponsor_user,omitempty"`

	// CommissionPerMille - The number of Telegram Stars received by the bot for each 1000 Telegram Stars
	// received by the affiliate program sponsor from referred users
	CommissionPerMille int `json:"commission_per_mille"`
}

// PartnerType returns TransactionPartner type
func (p *TransactionPartnerAffiliateProgram) PartnerType() string {
	return PartnerTypeAffiliateProgram
}

func (p *TransactionPartnerAffiliateProgram) iTransactionPartner() {}

// TransactionPartnerFragment - Describes a withdrawal transaction with Fragment.
type TransactionPartnerFragment struct {
	// Type - Type of the transaction partner, always “fragment”
	Type string `json:"type"`

	// WithdrawalState - Optional. State of the transaction if the transaction is outgoing
	WithdrawalState RevenueWithdrawalState `json:"withdrawal_state,omitempty"`
}

// PartnerType returns TransactionPartner type
func (p *TransactionPartnerFragment) PartnerType() string {
	return PartnerTypeFragment
}

func (p *TransactionPartnerFragment) iTransactionPartner() {}

// TransactionPartnerTelegramAds - Describes a withdrawal transaction to the Telegram Ads platform.
type TransactionPartnerTelegramAds struct {
	// Type - Type of the transaction partner, always “telegram_ads”
	Type string `json:"type"`
}

// PartnerType returns TransactionPartner type
func (p *TransactionPartnerTelegramAds) PartnerType() string {
	return PartnerTypeTelegramAds
}

func (p *TransactionPartnerTelegramAds) iTransactionPartner() {}

// TransactionPartnerTelegramApi - Describes a transaction with payment for paid broadcasting
// (https://core.telegram.org/bots/api#paid-broadcasts).
type TransactionPartnerTelegramApi struct { //nolint:revive
	// Type - Type of the transaction partner, always “telegram_api”
	Type string `json:"type"`

	// RequestCount - The number of successful requests that exceeded regular limits and were therefore billed
	RequestCount int `json:"request_count"`
}

// PartnerType returns TransactionPartner type
func (p *TransactionPartnerTelegramApi) PartnerType() string {
	return PartnerTypeTelegramApi
}

func (p *TransactionPartnerTelegramApi) iTransactionPartner() {}

// TransactionPartnerOther - Describes a transaction with an unknown source or recipient.
type TransactionPartnerOther struct {
	// Type - Type of the transaction partner, always “other”
	Type string `json:"type"`
}

// PartnerType returns TransactionPartner type
func (p *TransactionPartnerOther) PartnerType() string {
	return PartnerTypeOther
}

func (p *TransactionPartnerOther) iTransactionPartner() {}

// StarTransaction - Describes a Telegram Star transaction. Note that if the buyer initiates a chargeback
// with the payment provider from whom they acquired Stars (e.g., Apple, Google) following this transaction, the
// refunded Stars will be deducted from the bot's balance. This is outside of Telegram's control.
type StarTransaction struct {
	// ID - Unique identifier of the transaction. Coincides with the identifier of the original transaction for
	// refund transactions. Coincides with SuccessfulPayment.telegram_payment_charge_id for successful incoming
	// payments from users.
	ID string `json:"id"`

	// Amount - Integer amount of Telegram Stars transferred by the transaction
	Amount int `json:"amount"`

	// NanostarAmount - Optional. The number of 1/1000000000 shares of Telegram Stars transferred by the
	// transaction; from 0 to 999999999
	NanostarAmount int `json:"nanostar_amount,omitempty"`

	// Date - Date the transaction was created in Unix time
	Date int64 `json:"date"`

	// Source - Optional. Source of an incoming transaction (e.g., a user purchasing goods or services, Fragment
	// refunding a failed withdrawal). Only for incoming transactions
	Source TransactionPartner `json:"source,omitempty"`

	// Receiver - Optional. Receiver of an outgoing transaction (e.g., a user for a purchase refund, Fragment
	// for a withdrawal). Only for outgoing transactions
	Receiver TransactionPartner `json:"receiver,omitempty"`
}

// UnmarshalJSON converts JSON to Chat
func (t *StarTransaction) UnmarshalJSON(data []byte) error {
	parser := json.ParserPoll.Get()
	defer json.ParserPoll.Put(parser)

	value, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}

	type uStarTransaction StarTransaction
	var ut uStarTransaction

	if value.Exists("source") {
		partnerType := string(value.GetStringBytes("source", "type"))
		switch partnerType {
		case PartnerTypeUser:
			ut.Source = &TransactionPartnerUser{}
		case PartnerTypeChat:
			ut.Source = &TransactionPartnerChat{}
		case PartnerTypeAffiliateProgram:
			ut.Source = &TransactionPartnerAffiliateProgram{}
		case PartnerTypeFragment:
			ut.Source = &TransactionPartnerFragment{}
		case PartnerTypeTelegramAds:
			ut.Source = &TransactionPartnerTelegramAds{}
		case PartnerTypeTelegramApi:
			ut.Source = &TransactionPartnerTelegramApi{}
		case PartnerTypeOther:
			ut.Source = &TransactionPartnerOther{}
		default:
			return fmt.Errorf("unknown source partner type: %q", partnerType)
		}
	}

	if value.Exists("receiver") {
		partnerType := string(value.GetStringBytes("receiver", "type"))
		switch partnerType {
		case PartnerTypeUser:
			ut.Receiver = &TransactionPartnerUser{}
		case PartnerTypeChat:
			ut.Receiver = &TransactionPartnerChat{}
		case PartnerTypeAffiliateProgram:
			ut.Receiver = &TransactionPartnerAffiliateProgram{}
		case PartnerTypeFragment:
			ut.Receiver = &TransactionPartnerFragment{}
		case PartnerTypeTelegramAds:
			ut.Receiver = &TransactionPartnerTelegramAds{}
		case PartnerTypeTelegramApi:
			ut.Receiver = &TransactionPartnerTelegramApi{}
		case PartnerTypeOther:
			ut.Receiver = &TransactionPartnerOther{}
		default:
			return fmt.Errorf("unknown receiver partner type: %q", partnerType)
		}
	}

	if err = json.Unmarshal(data, &ut); err != nil {
		return err
	}
	*t = StarTransaction(ut)

	return nil
}

// StarTransactions - Contains a list of Telegram Star transactions.
type StarTransactions struct {
	// Transactions - The list of transactions
	Transactions []StarTransaction `json:"transactions"`
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

	// Data - Optional. Base64-encoded encrypted Telegram Passport element data provided by the user; available
	// only for “personal_details”, “passport”, “driver_license”, “identity_card”,
	// “internal_passport” and “address” types. Can be decrypted and verified using the accompanying
	// EncryptedCredentials (https://core.telegram.org/bots/api#encryptedcredentials).
	Data string `json:"data,omitempty"`

	// PhoneNumber - Optional. User's verified phone number; available only for “phone_number” type
	PhoneNumber string `json:"phone_number,omitempty"`

	// Email - Optional. User's verified email address; available only for “email” type
	Email string `json:"email,omitempty"`

	// Files - Optional. Array of encrypted files with documents provided by the user; available only for
	// “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and
	// “temporary_registration” types. Files can be decrypted and verified using the accompanying
	// EncryptedCredentials (https://core.telegram.org/bots/api#encryptedcredentials).
	Files []PassportFile `json:"files,omitempty"`

	// FrontSide - Optional. Encrypted file with the front side of the document, provided by the user; available
	// only for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can
	// be decrypted and verified using the accompanying EncryptedCredentials
	// (https://core.telegram.org/bots/api#encryptedcredentials).
	FrontSide *PassportFile `json:"front_side,omitempty"`

	// ReverseSide - Optional. Encrypted file with the reverse side of the document, provided by the user;
	// available only for “driver_license” and “identity_card”. The file can be decrypted and verified using
	// the accompanying EncryptedCredentials (https://core.telegram.org/bots/api#encryptedcredentials).
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`

	// Selfie - Optional. Encrypted file with the selfie of the user holding a document, provided by the user;
	// available if requested for “passport”, “driver_license”, “identity_card” and
	// “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials
	// (https://core.telegram.org/bots/api#encryptedcredentials).
	Selfie *PassportFile `json:"selfie,omitempty"`

	// Translation - Optional. Array of encrypted files with translated versions of documents provided by the
	// user; available if requested for “passport”, “driver_license”, “identity_card”,
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
	Secret string `json:"secret"` //nolint:gosec
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
	// ErrorSource returns PassportElementError source
	ErrorSource() string
	// Disallow external implementations
	iPassportElementError()
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

func (p *PassportElementErrorDataField) iPassportElementError() {}

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

func (p *PassportElementErrorFrontSide) iPassportElementError() {}

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

func (p *PassportElementErrorReverseSide) iPassportElementError() {}

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

func (p *PassportElementErrorSelfie) iPassportElementError() {}

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

func (p *PassportElementErrorFile) iPassportElementError() {}

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

func (p *PassportElementErrorFiles) iPassportElementError() {}

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

func (p *PassportElementErrorTranslationFile) iPassportElementError() {}

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

func (p *PassportElementErrorTranslationFiles) iPassportElementError() {}

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

func (p *PassportElementErrorUnspecified) iPassportElementError() {}

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
