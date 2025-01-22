package telegohandler

import (
	"github.com/mymmrac/telego"
)

// TODO: Implement missing handlers

// MessageHandler handles message that came from bot
type MessageHandler func(ctx *Context, message telego.Message) error

// HandleMessage same as Handle, but assumes that the update contains a message
func (h *HandlerGroup) HandleMessage(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil message handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.Message)
	}, append([]Predicate{AnyMessage()}, predicates...)...)
}

// HandleMessage same as Handle, but assumes that the update contains a message
func (h *BotHandler) HandleMessage(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleMessage(handler, predicates...)
}

// HandleEditedMessage same as Handle, but assumes that the update contains an edited message
func (h *HandlerGroup) HandleEditedMessage(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil edited message handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.EditedMessage)
	}, append([]Predicate{AnyEditedMessage()}, predicates...)...)
}

// HandleEditedMessage same as Handle, but assumes that the update contains an edited message
func (h *BotHandler) HandleEditedMessage(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleEditedMessage(handler, predicates...)
}

// HandleChannelPost same as Handle, but assumes that the update contains a channel post
func (h *HandlerGroup) HandleChannelPost(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil channel post handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ChannelPost)
	}, append([]Predicate{AnyChannelPost()}, predicates...)...)
}

// HandleChannelPost same as Handle, but assumes that the update contains a channel post
func (h *BotHandler) HandleChannelPost(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleChannelPost(handler, predicates...)
}

// HandleEditedChannelPost same as Handle, but assumes that the update contains an edited channel post
func (h *HandlerGroup) HandleEditedChannelPost(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil edited channel post handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.EditedChannelPost)
	}, append([]Predicate{AnyEditedChannelPost()}, predicates...)...)
}

// HandleEditedChannelPost same as Handle, but assumes that the update contains an edited channel post
func (h *BotHandler) HandleEditedChannelPost(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleEditedChannelPost(handler, predicates...)
}

// InlineQueryHandler handles inline queries that came from bot
type InlineQueryHandler func(ctx *Context, query telego.InlineQuery) error

// HandleInlineQuery same as Handle, but assumes that the update contains an inline query
func (h *HandlerGroup) HandleInlineQuery(handler InlineQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil inline query handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.InlineQuery)
	}, append([]Predicate{AnyInlineQuery()}, predicates...)...)
}

// HandleInlineQuery same as Handle, but assumes that the update contains an inline query
func (h *BotHandler) HandleInlineQuery(handler InlineQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandleInlineQuery(handler, predicates...)
}

// ChosenInlineResultHandler handles chosen inline result that came from bot
type ChosenInlineResultHandler func(ctx *Context, result telego.ChosenInlineResult) error

// HandleChosenInlineResult same as Handle, but assumes that the update contains a chosen inline result
func (h *HandlerGroup) HandleChosenInlineResult(handler ChosenInlineResultHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chosen inline query handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ChosenInlineResult)
	}, append([]Predicate{AnyChosenInlineResult()}, predicates...)...)
}

// HandleChosenInlineResult same as Handle, but assumes that the update contains a chosen inline result
func (h *BotHandler) HandleChosenInlineResult(handler ChosenInlineResultHandler, predicates ...Predicate) {
	h.baseGroup.HandleChosenInlineResult(handler, predicates...)
}

// CallbackQueryHandler handles callback queries that came from bot
type CallbackQueryHandler func(ctx *Context, query telego.CallbackQuery) error

// HandleCallbackQuery same as Handle, but assumes that the update contains a callback query
func (h *HandlerGroup) HandleCallbackQuery(handler CallbackQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil callback query handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.CallbackQuery)
	}, append([]Predicate{AnyCallbackQuery()}, predicates...)...)
}

// HandleCallbackQuery same as Handle, but assumes that the update contains a callback query
func (h *BotHandler) HandleCallbackQuery(handler CallbackQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandleCallbackQuery(handler, predicates...)
}

// ShippingQueryHandler handles shipping query that came from bot
type ShippingQueryHandler func(ctx *Context, query telego.ShippingQuery) error

// HandleShippingQuery same as Handle, but assumes that the update contains a shipping query
func (h *HandlerGroup) HandleShippingQuery(handler ShippingQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil shipping query handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ShippingQuery)
	}, append([]Predicate{AnyShippingQuery()}, predicates...)...)
}

// HandleShippingQuery same as Handle, but assumes that the update contains a shipping query
func (h *BotHandler) HandleShippingQuery(handler ShippingQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandleShippingQuery(handler, predicates...)
}

// PreCheckoutQueryHandler handles pre checkout query that came from bot
type PreCheckoutQueryHandler func(ctx *Context, query telego.PreCheckoutQuery) error

// HandlePreCheckoutQuery same as Handle, but assumes that the update contains a pre checkout query
func (h *HandlerGroup) HandlePreCheckoutQuery(handler PreCheckoutQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil pre checkout query handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.PreCheckoutQuery)
	}, append([]Predicate{AnyPreCheckoutQuery()}, predicates...)...)
}

// HandlePreCheckoutQuery same as Handle, but assumes that the update contains a pre checkout query
func (h *BotHandler) HandlePreCheckoutQuery(handler PreCheckoutQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandlePreCheckoutQuery(handler, predicates...)
}

// PollHandler handles poll that came from bot
type PollHandler func(ctx *Context, poll telego.Poll) error

// HandlePoll same as Handle, but assumes that the update contains a poll
func (h *HandlerGroup) HandlePoll(handler PollHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil poll handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.Poll)
	}, append([]Predicate{AnyPoll()}, predicates...)...)
}

// HandlePoll same as Handle, but assumes that the update contains a poll
func (h *BotHandler) HandlePoll(handler PollHandler, predicates ...Predicate) {
	h.baseGroup.HandlePoll(handler, predicates...)
}

// PollAnswerHandler handles poll answer that came from bot
type PollAnswerHandler func(ctx *Context, answer telego.PollAnswer) error

// HandlePollAnswer same as Handle, but assumes that the update contains a poll answer
func (h *HandlerGroup) HandlePollAnswer(handler PollAnswerHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil poll answer handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.PollAnswer)
	}, append([]Predicate{AnyPollAnswer()}, predicates...)...)
}

// HandlePollAnswer same as Handle, but assumes that the update contains a poll answer
func (h *BotHandler) HandlePollAnswer(handler PollAnswerHandler, predicates ...Predicate) {
	h.baseGroup.HandlePollAnswer(handler, predicates...)
}

// ChatMemberUpdatedHandler handles chat member that came from bot
type ChatMemberUpdatedHandler func(ctx *Context, chatMember telego.ChatMemberUpdated) error

// HandleMyChatMemberUpdated same as Handle, but assumes that the update contains my chat member
func (h *HandlerGroup) HandleMyChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil my chat member update handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.MyChatMember)
	}, append([]Predicate{AnyMyChatMember()}, predicates...)...)
}

// HandleMyChatMemberUpdated same as Handle, but assumes that the update contains my chat member
func (h *BotHandler) HandleMyChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	h.baseGroup.HandleMyChatMemberUpdated(handler, predicates...)
}

// HandleChatMemberUpdated same as Handle, but assumes that the update contains chat member
func (h *HandlerGroup) HandleChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chat member update handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ChatMember)
	}, append([]Predicate{AnyChatMember()}, predicates...)...)
}

// HandleChatMemberUpdated same as Handle, but assumes that the update contains chat member
func (h *BotHandler) HandleChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	h.baseGroup.HandleChatMemberUpdated(handler, predicates...)
}

// ChatJoinRequestHandler handles chat join request that came from bot
type ChatJoinRequestHandler func(ctx *Context, request telego.ChatJoinRequest) error

// HandleChatJoinRequest same as Handle, but assumes that the update contains chat join request
func (h *HandlerGroup) HandleChatJoinRequest(handler ChatJoinRequestHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chat join request handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ChatJoinRequest)
	}, append([]Predicate{AnyChatJoinRequest()}, predicates...)...)
}

// HandleChatJoinRequest same as Handle, but assumes that the update contains chat join request
func (h *BotHandler) HandleChatJoinRequest(handler ChatJoinRequestHandler, predicates ...Predicate) {
	h.baseGroup.HandleChatJoinRequest(handler, predicates...)
}
