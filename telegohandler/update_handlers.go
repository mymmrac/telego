package telegohandler

import "github.com/mymmrac/telego"

// MessageHandler handles message that came from bot
type MessageHandler func(bot *telego.Bot, message telego.Message)

// HandleMessage same as Handle, but assumes that the update contains a message
func (h *BotHandler) HandleMessage(handler MessageHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.Message)
	}, append([]Predicate{AnyMessage()}, predicates...)...)
}

// HandleEditedMessage same as Handle, but assumes that the update contains an edited message
func (h *BotHandler) HandleEditedMessage(handler MessageHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.EditedMessage)
	}, append([]Predicate{AnyEditedMessage()}, predicates...)...)
}

// HandleChannelPost same as Handle, but assumes that the update contains a channel post
func (h *BotHandler) HandleChannelPost(handler MessageHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.ChannelPost)
	}, append([]Predicate{AnyChannelPost()}, predicates...)...)
}

// HandleEditedChannelPost same as Handle, but assumes that the update contains an edited channel post
func (h *BotHandler) HandleEditedChannelPost(handler MessageHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.EditedChannelPost)
	}, append([]Predicate{AnyEditedChannelPost()}, predicates...)...)
}

// InlineQueryHandler handles inline queries that came from bot
type InlineQueryHandler func(bot *telego.Bot, query telego.InlineQuery)

// HandleInlineQuery same as Handle, but assumes that the update contains an inline query
func (h *BotHandler) HandleInlineQuery(handler InlineQueryHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.InlineQuery)
	}, append([]Predicate{AnyInlineQuery()}, predicates...)...)
}

// ChosenInlineResultHandler handles chosen inline result that came from bot
type ChosenInlineResultHandler func(bot *telego.Bot, result telego.ChosenInlineResult)

// HandleChosenInlineResult same as Handle, but assumes that the update contains a chosen inline result
func (h *BotHandler) HandleChosenInlineResult(handler ChosenInlineResultHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.ChosenInlineResult)
	}, append([]Predicate{AnyChosenInlineResult()}, predicates...)...)
}

// CallbackQueryHandler handles callback queries that came from bot
type CallbackQueryHandler func(bot *telego.Bot, query telego.CallbackQuery)

// HandleCallbackQuery same as Handle, but assumes that the update contains a callback query
func (h *BotHandler) HandleCallbackQuery(handler CallbackQueryHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.CallbackQuery)
	}, append([]Predicate{AnyCallbackQuery()}, predicates...)...)
}

// ShippingQueryHandler handles shipping query that came from bot
type ShippingQueryHandler func(bot *telego.Bot, query telego.ShippingQuery)

// HandleShippingQuery same as Handle, but assumes that the update contains a shipping query
func (h *BotHandler) HandleShippingQuery(handler ShippingQueryHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.ShippingQuery)
	}, append([]Predicate{AnyShippingQuery()}, predicates...)...)
}

// PreCheckoutQueryHandler handles pre checkout query that came from bot
type PreCheckoutQueryHandler func(bot *telego.Bot, query telego.PreCheckoutQuery)

// HandlePreCheckoutQuery same as Handle, but assumes that the update contains a pre checkout query
func (h *BotHandler) HandlePreCheckoutQuery(handler PreCheckoutQueryHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.PreCheckoutQuery)
	}, append([]Predicate{AnyPreCheckoutQuery()}, predicates...)...)
}

// PollHandler handles poll that came from bot
type PollHandler func(bot *telego.Bot, poll telego.Poll)

// HandlePoll same as Handle, but assumes that the update contains a poll
func (h *BotHandler) HandlePoll(handler PollHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.Poll)
	}, append([]Predicate{AnyPoll()}, predicates...)...)
}

// PollAnswerHandler handles poll answer that came from bot
type PollAnswerHandler func(bot *telego.Bot, answer telego.PollAnswer)

// HandlePollAnswer same as Handle, but assumes that the update contains a poll answer
func (h *BotHandler) HandlePollAnswer(handler PollAnswerHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.PollAnswer)
	}, append([]Predicate{AnyPollAnswer()}, predicates...)...)
}

// ChatMemberUpdatedHandler handles chat member that came from bot
type ChatMemberUpdatedHandler func(bot *telego.Bot, chatMember telego.ChatMemberUpdated)

// HandleMyChatMemberUpdated same as Handle, but assumes that the update contains my chat member
func (h *BotHandler) HandleMyChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.MyChatMember)
	}, append([]Predicate{AnyMyChatMember()}, predicates...)...)
}

// HandleChatMemberUpdated same as Handle, but assumes that the update contains chat member
func (h *BotHandler) HandleChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.ChatMember)
	}, append([]Predicate{AnyChatMember()}, predicates...)...)
}

// ChatJoinRequestHandler handles chat join request that came from bot
type ChatJoinRequestHandler func(bot *telego.Bot, request telego.ChatJoinRequest)

// HandleChatJoinRequest same as Handle, but assumes that the update contains chat join request
func (h *BotHandler) HandleChatJoinRequest(handler ChatJoinRequestHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.ChatJoinRequest)
	}, append([]Predicate{AnyChatJoinRequest()}, predicates...)...)
}
