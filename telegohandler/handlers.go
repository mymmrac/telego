package telegohandler

import (
	"github.com/mymmrac/telego"
)

// MessageHandler handles message that came from bot
type MessageHandler func(ctx *Context, message telego.Message) error

// HandleMessage same as [BotHandler.Handle], but assumes that the update contains a message
func (h *HandlerGroup) HandleMessage(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil message handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.Message)
	}, append([]Predicate{AnyMessage()}, predicates...)...)
}

// HandleMessage same as [BotHandler.Handle], but assumes that the update contains a message
func (h *BotHandler) HandleMessage(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleMessage(handler, predicates...)
}

// HandleEditedMessage same as [BotHandler.Handle], but assumes that the update contains an edited message
func (h *HandlerGroup) HandleEditedMessage(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil edited message handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.EditedMessage)
	}, append([]Predicate{AnyEditedMessage()}, predicates...)...)
}

// HandleEditedMessage same as [BotHandler.Handle], but assumes that the update contains an edited message
func (h *BotHandler) HandleEditedMessage(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleEditedMessage(handler, predicates...)
}

// HandleChannelPost same as [BotHandler.Handle], but assumes that the update contains a channel post
func (h *HandlerGroup) HandleChannelPost(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil channel post-handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ChannelPost)
	}, append([]Predicate{AnyChannelPost()}, predicates...)...)
}

// HandleChannelPost same as [BotHandler.Handle], but assumes that the update contains a channel post
func (h *BotHandler) HandleChannelPost(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleChannelPost(handler, predicates...)
}

// HandleEditedChannelPost same as [BotHandler.Handle], but assumes that the update contains an edited channel post
func (h *HandlerGroup) HandleEditedChannelPost(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil edited channel post-handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.EditedChannelPost)
	}, append([]Predicate{AnyEditedChannelPost()}, predicates...)...)
}

// HandleEditedChannelPost same as [BotHandler.Handle], but assumes that the update contains an edited channel post
func (h *BotHandler) HandleEditedChannelPost(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleEditedChannelPost(handler, predicates...)
}

// BusinessConnectionHandler handles business connection that came from bot
type BusinessConnectionHandler func(ctx *Context, connection telego.BusinessConnection) error

// HandleBusinessConnection same as [BotHandler.Handle], but assumes that the update contains a business connection
func (h *HandlerGroup) HandleBusinessConnection(handler BusinessConnectionHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil business connection handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.BusinessConnection)
	}, append([]Predicate{AnyBusinessConnection()}, predicates...)...)
}

// HandleBusinessConnection same as [BotHandler.Handle], but assumes that the update contains a business connection
func (h *BotHandler) HandleBusinessConnection(handler BusinessConnectionHandler, predicates ...Predicate) {
	h.baseGroup.HandleBusinessConnection(handler, predicates...)
}

// HandleBusinessMessage same as [BotHandler.Handle], but assumes that the update contains a business message
func (h *HandlerGroup) HandleBusinessMessage(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil business message handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.BusinessMessage)
	}, append([]Predicate{AnyBusinessMessage()}, predicates...)...)
}

// HandleBusinessMessage same as [BotHandler.Handle], but assumes that the update contains a business message
func (h *BotHandler) HandleBusinessMessage(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleBusinessMessage(handler, predicates...)
}

// HandleEditedBusinessMessage same as [BotHandler.Handle], but assumes that the update contains an edited
// business message
func (h *HandlerGroup) HandleEditedBusinessMessage(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil edited business message handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.EditedBusinessMessage)
	}, append([]Predicate{AnyEditedBusinessMessage()}, predicates...)...)
}

// HandleEditedBusinessMessage same as [BotHandler.Handle], but assumes that the update contains an edited
// business message
func (h *BotHandler) HandleEditedBusinessMessage(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleEditedBusinessMessage(handler, predicates...)
}

// DeletedBusinessMessagesHandler handles deleted business messages that came from bot
type DeletedBusinessMessagesHandler func(ctx *Context, deletedMessage telego.BusinessMessagesDeleted) error

// HandleDeletedBusinessMessages same as [BotHandler.Handle], but assumes that the update contains a deleted
// business messages
func (h *HandlerGroup) HandleDeletedBusinessMessages(handler DeletedBusinessMessagesHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil deleted business messages handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.DeletedBusinessMessages)
	}, append([]Predicate{AnyDeletedBusinessMessages()}, predicates...)...)
}

// HandleDeletedBusinessMessages same as [BotHandler.Handle], but assumes that the update contains a deleted
// business messages
func (h *BotHandler) HandleDeletedBusinessMessages(handler DeletedBusinessMessagesHandler, predicates ...Predicate) {
	h.baseGroup.HandleDeletedBusinessMessages(handler, predicates...)
}

// MessageReactionHandler handles message reaction that came from bot
type MessageReactionHandler func(ctx *Context, reaction telego.MessageReactionUpdated) error

// HandleMessageReaction same as [BotHandler.Handle], but assumes that the update contains a message reaction
func (h *HandlerGroup) HandleMessageReaction(handler MessageReactionHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil message reaction handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.MessageReaction)
	}, append([]Predicate{AnyMessageReaction()}, predicates...)...)
}

// HandleMessageReaction same as [BotHandler.Handle], but assumes that the update contains a message reaction
func (h *BotHandler) HandleMessageReaction(handler MessageReactionHandler, predicates ...Predicate) {
	h.baseGroup.HandleMessageReaction(handler, predicates...)
}

// MessageReactionCountHandler handles message reaction that came from bot
type MessageReactionCountHandler func(ctx *Context, reaction telego.MessageReactionCountUpdated) error

// HandleMessageReactionCount same as [BotHandler.Handle], but assumes that the update contains a message reaction count
func (h *HandlerGroup) HandleMessageReactionCount(handler MessageReactionCountHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil message reaction count handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.MessageReactionCount)
	}, append([]Predicate{AnyMessageReactionCount()}, predicates...)...)
}

// HandleMessageReactionCount same as [BotHandler.Handle], but assumes that the update contains a message reaction count
func (h *BotHandler) HandleMessageReactionCount(handler MessageReactionCountHandler, predicates ...Predicate) {
	h.baseGroup.HandleMessageReactionCount(handler, predicates...)
}

// InlineQueryHandler handles inline queries that came from bot
type InlineQueryHandler func(ctx *Context, query telego.InlineQuery) error

// HandleInlineQuery same as [BotHandler.Handle], but assumes that the update contains an inline query
func (h *HandlerGroup) HandleInlineQuery(handler InlineQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil inline query handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.InlineQuery)
	}, append([]Predicate{AnyInlineQuery()}, predicates...)...)
}

// HandleInlineQuery same as [BotHandler.Handle], but assumes that the update contains an inline query
func (h *BotHandler) HandleInlineQuery(handler InlineQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandleInlineQuery(handler, predicates...)
}

// ChosenInlineResultHandler handles chosen an inline result that came from bot
type ChosenInlineResultHandler func(ctx *Context, result telego.ChosenInlineResult) error

// HandleChosenInlineResult same as [BotHandler.Handle], but assumes that the update contains a chosen inline result
func (h *HandlerGroup) HandleChosenInlineResult(handler ChosenInlineResultHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chosen inline query handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ChosenInlineResult)
	}, append([]Predicate{AnyChosenInlineResult()}, predicates...)...)
}

// HandleChosenInlineResult same as [BotHandler.Handle], but assumes that the update contains a chosen inline result
func (h *BotHandler) HandleChosenInlineResult(handler ChosenInlineResultHandler, predicates ...Predicate) {
	h.baseGroup.HandleChosenInlineResult(handler, predicates...)
}

// CallbackQueryHandler handles callback queries that came from bot
type CallbackQueryHandler func(ctx *Context, query telego.CallbackQuery) error

// HandleCallbackQuery same as [BotHandler.Handle], but assumes that the update contains a callback query
func (h *HandlerGroup) HandleCallbackQuery(handler CallbackQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil callback query handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.CallbackQuery)
	}, append([]Predicate{AnyCallbackQuery()}, predicates...)...)
}

// HandleCallbackQuery same as [BotHandler.Handle], but assumes that the update contains a callback query
func (h *BotHandler) HandleCallbackQuery(handler CallbackQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandleCallbackQuery(handler, predicates...)
}

// ShippingQueryHandler handles shipping query that came from bot
type ShippingQueryHandler func(ctx *Context, query telego.ShippingQuery) error

// HandleShippingQuery same as [BotHandler.Handle], but assumes that the update contains a shipping query
func (h *HandlerGroup) HandleShippingQuery(handler ShippingQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil shipping query handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ShippingQuery)
	}, append([]Predicate{AnyShippingQuery()}, predicates...)...)
}

// HandleShippingQuery same as [BotHandler.Handle], but assumes that the update contains a shipping query
func (h *BotHandler) HandleShippingQuery(handler ShippingQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandleShippingQuery(handler, predicates...)
}

// PreCheckoutQueryHandler handles pre-checkout query that came from bot
type PreCheckoutQueryHandler func(ctx *Context, query telego.PreCheckoutQuery) error

// HandlePreCheckoutQuery same as [BotHandler.Handle], but assumes that the update contains a pre-checkout query
func (h *HandlerGroup) HandlePreCheckoutQuery(handler PreCheckoutQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil pre-checkout query handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.PreCheckoutQuery)
	}, append([]Predicate{AnyPreCheckoutQuery()}, predicates...)...)
}

// HandlePreCheckoutQuery same as [BotHandler.Handle], but assumes that the update contains a pre-checkout query
func (h *BotHandler) HandlePreCheckoutQuery(handler PreCheckoutQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandlePreCheckoutQuery(handler, predicates...)
}

// PurchasedPaidMediaHandler handles purchased paid media that came from bot
type PurchasedPaidMediaHandler func(ctx *Context, purchase telego.PaidMediaPurchased) error

// HandlePurchasedPaidMedia same as [BotHandler.Handle], but assumes that the update contains a purchased paid media
func (h *HandlerGroup) HandlePurchasedPaidMedia(handler PurchasedPaidMediaHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil purchased paid media handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.PurchasedPaidMedia)
	}, append([]Predicate{AnyPurchasedPaidMedia()}, predicates...)...)
}

// HandlePurchasedPaidMedia same as [BotHandler.Handle], but assumes that the update contains a pre-checkout query
func (h *BotHandler) HandlePurchasedPaidMedia(handler PurchasedPaidMediaHandler, predicates ...Predicate) {
	h.baseGroup.HandlePurchasedPaidMedia(handler, predicates...)
}

// PollHandler handles poll that came from bot
type PollHandler func(ctx *Context, poll telego.Poll) error

// HandlePoll same as [BotHandler.Handle], but assumes that the update contains a poll
func (h *HandlerGroup) HandlePoll(handler PollHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil poll handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.Poll)
	}, append([]Predicate{AnyPoll()}, predicates...)...)
}

// HandlePoll same as [BotHandler.Handle], but assumes that the update contains a poll
func (h *BotHandler) HandlePoll(handler PollHandler, predicates ...Predicate) {
	h.baseGroup.HandlePoll(handler, predicates...)
}

// PollAnswerHandler handles poll answer that came from bot
type PollAnswerHandler func(ctx *Context, answer telego.PollAnswer) error

// HandlePollAnswer same as [BotHandler.Handle], but assumes that the update contains a poll answer
func (h *HandlerGroup) HandlePollAnswer(handler PollAnswerHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil poll answer handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.PollAnswer)
	}, append([]Predicate{AnyPollAnswer()}, predicates...)...)
}

// HandlePollAnswer same as [BotHandler.Handle], but assumes that the update contains a poll answer
func (h *BotHandler) HandlePollAnswer(handler PollAnswerHandler, predicates ...Predicate) {
	h.baseGroup.HandlePollAnswer(handler, predicates...)
}

// ChatMemberUpdatedHandler handles chat member that came from bot
type ChatMemberUpdatedHandler func(ctx *Context, chatMember telego.ChatMemberUpdated) error

// HandleMyChatMemberUpdated same as [BotHandler.Handle], but assumes that the update contains my chat member
func (h *HandlerGroup) HandleMyChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil my chat member update handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.MyChatMember)
	}, append([]Predicate{AnyMyChatMember()}, predicates...)...)
}

// HandleMyChatMemberUpdated same as [BotHandler.Handle], but assumes that the update contains my chat member
func (h *BotHandler) HandleMyChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	h.baseGroup.HandleMyChatMemberUpdated(handler, predicates...)
}

// HandleChatMemberUpdated same as [BotHandler.Handle], but assumes that the update contains chat member
func (h *HandlerGroup) HandleChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chat member update handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ChatMember)
	}, append([]Predicate{AnyChatMember()}, predicates...)...)
}

// HandleChatMemberUpdated same as [BotHandler.Handle], but assumes that the update contains chat member
func (h *BotHandler) HandleChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	h.baseGroup.HandleChatMemberUpdated(handler, predicates...)
}

// ChatJoinRequestHandler handles chat join request that came from bot
type ChatJoinRequestHandler func(ctx *Context, request telego.ChatJoinRequest) error

// HandleChatJoinRequest same as [BotHandler.Handle], but assumes that the update contains chat join request
func (h *HandlerGroup) HandleChatJoinRequest(handler ChatJoinRequestHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chat join request handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ChatJoinRequest)
	}, append([]Predicate{AnyChatJoinRequest()}, predicates...)...)
}

// HandleChatJoinRequest same as [BotHandler.Handle], but assumes that the update contains chat join request
func (h *BotHandler) HandleChatJoinRequest(handler ChatJoinRequestHandler, predicates ...Predicate) {
	h.baseGroup.HandleChatJoinRequest(handler, predicates...)
}

// ChatBoostHandler handles chat boost that came from bot
type ChatBoostHandler func(ctx *Context, boost telego.ChatBoostUpdated) error

// HandleChatBoost same as [BotHandler.Handle], but assumes that the update contains chat boost
func (h *HandlerGroup) HandleChatBoost(handler ChatBoostHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chat boost handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.ChatBoost)
	}, append([]Predicate{AnyChatBoost()}, predicates...)...)
}

// HandleChatBoost same as [BotHandler.Handle], but assumes that the update contains chat boost
func (h *BotHandler) HandleChatBoost(handler ChatBoostHandler, predicates ...Predicate) {
	h.baseGroup.HandleChatBoost(handler, predicates...)
}

// RemovedChatBoostHandler handles removed chat boost that came from bot
type RemovedChatBoostHandler func(ctx *Context, removedBoost telego.ChatBoostRemoved) error

// HandleRemovedChatBoost same as [BotHandler.Handle], but assumes that the update contains removed chat boost
func (h *HandlerGroup) HandleRemovedChatBoost(handler RemovedChatBoostHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil removed chat boost handlers not allowed")
	}

	h.Handle(func(ctx *Context, update telego.Update) error {
		return handler(ctx, *update.RemovedChatBoost)
	}, append([]Predicate{AnyRemovedChatBoost()}, predicates...)...)
}

// HandleRemovedChatBoost same as [BotHandler.Handle], but assumes that the update contains removed chat boost
func (h *BotHandler) HandleRemovedChatBoost(handler RemovedChatBoostHandler, predicates ...Predicate) {
	h.baseGroup.HandleRemovedChatBoost(handler, predicates...)
}
