package telegohandler

import (
	"context"

	"github.com/mymmrac/telego"
)

// MessageHandler handles message that came from bot
type MessageHandler func(bot *telego.Bot, message telego.Message)

// MessageHandlerCtx handles message that came from bot with context
type MessageHandlerCtx func(ctx context.Context, bot *telego.Bot, message telego.Message)

// HandleMessage same as Handle, but assumes that the update contains a message
func (h *HandlerGroup) HandleMessage(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil message handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.Message)
	}, append([]Predicate{AnyMessage()}, predicates...)...)
}

// HandleMessageCtx same as Handle, but assumes that the update contains a message
func (h *HandlerGroup) HandleMessageCtx(handler MessageHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil message handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.Message)
	}, append([]Predicate{AnyMessage()}, predicates...)...)
}

// HandleMessage same as Handle, but assumes that the update contains a message
func (h *BotHandler) HandleMessage(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleMessage(handler, predicates...)
}

// HandleMessageCtx same as Handle, but assumes that the update contains a message
func (h *BotHandler) HandleMessageCtx(handler MessageHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleMessageCtx(handler, predicates...)
}

// HandleEditedMessage same as Handle, but assumes that the update contains an edited message
func (h *HandlerGroup) HandleEditedMessage(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil edited message handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.EditedMessage)
	}, append([]Predicate{AnyEditedMessage()}, predicates...)...)
}

// HandleEditedMessageCtx same as Handle, but assumes that the update contains an edited message
func (h *HandlerGroup) HandleEditedMessageCtx(handler MessageHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil edited message handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.EditedMessage)
	}, append([]Predicate{AnyEditedMessage()}, predicates...)...)
}

// HandleEditedMessage same as Handle, but assumes that the update contains an edited message
func (h *BotHandler) HandleEditedMessage(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleEditedMessage(handler, predicates...)
}

// HandleEditedMessageCtx same as Handle, but assumes that the update contains an edited message
func (h *BotHandler) HandleEditedMessageCtx(handler MessageHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleEditedMessageCtx(handler, predicates...)
}

// HandleChannelPost same as Handle, but assumes that the update contains a channel post
func (h *HandlerGroup) HandleChannelPost(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil channel post handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.ChannelPost)
	}, append([]Predicate{AnyChannelPost()}, predicates...)...)
}

// HandleChannelPostCtx same as Handle, but assumes that the update contains a channel post
func (h *HandlerGroup) HandleChannelPostCtx(handler MessageHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil channel post handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.ChannelPost)
	}, append([]Predicate{AnyChannelPost()}, predicates...)...)
}

// HandleChannelPost same as Handle, but assumes that the update contains a channel post
func (h *BotHandler) HandleChannelPost(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleChannelPost(handler, predicates...)
}

// HandleChannelPostCtx same as Handle, but assumes that the update contains a channel post
func (h *BotHandler) HandleChannelPostCtx(handler MessageHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleChannelPostCtx(handler, predicates...)
}

// HandleEditedChannelPost same as Handle, but assumes that the update contains an edited channel post
func (h *HandlerGroup) HandleEditedChannelPost(handler MessageHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil edited channel post handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.EditedChannelPost)
	}, append([]Predicate{AnyEditedChannelPost()}, predicates...)...)
}

// HandleEditedChannelPostCtx same as Handle, but assumes that the update contains an edited channel post
func (h *HandlerGroup) HandleEditedChannelPostCtx(handler MessageHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil edited channel post handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.EditedChannelPost)
	}, append([]Predicate{AnyEditedChannelPost()}, predicates...)...)
}

// HandleEditedChannelPost same as Handle, but assumes that the update contains an edited channel post
func (h *BotHandler) HandleEditedChannelPost(handler MessageHandler, predicates ...Predicate) {
	h.baseGroup.HandleEditedChannelPost(handler, predicates...)
}

// HandleEditedChannelPostCtx same as Handle, but assumes that the update contains an edited channel post
func (h *BotHandler) HandleEditedChannelPostCtx(handler MessageHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleEditedChannelPostCtx(handler, predicates...)
}

// InlineQueryHandler handles inline queries that came from bot
type InlineQueryHandler func(bot *telego.Bot, query telego.InlineQuery)

// InlineQueryHandlerCtx handles inline queries that came from bot with context
type InlineQueryHandlerCtx func(ctx context.Context, bot *telego.Bot, query telego.InlineQuery)

// HandleInlineQuery same as Handle, but assumes that the update contains an inline query
func (h *HandlerGroup) HandleInlineQuery(handler InlineQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil inline query handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.InlineQuery)
	}, append([]Predicate{AnyInlineQuery()}, predicates...)...)
}

// HandleInlineQueryCtx same as Handle, but assumes that the update contains an inline query
func (h *HandlerGroup) HandleInlineQueryCtx(handler InlineQueryHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil inline query handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.InlineQuery)
	}, append([]Predicate{AnyInlineQuery()}, predicates...)...)
}

// HandleInlineQuery same as Handle, but assumes that the update contains an inline query
func (h *BotHandler) HandleInlineQuery(handler InlineQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandleInlineQuery(handler, predicates...)
}

// HandleInlineQueryCtx same as Handle, but assumes that the update contains an inline query
func (h *BotHandler) HandleInlineQueryCtx(handler InlineQueryHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleInlineQueryCtx(handler, predicates...)
}

// ChosenInlineResultHandler handles chosen inline result that came from bot
type ChosenInlineResultHandler func(bot *telego.Bot, result telego.ChosenInlineResult)

// ChosenInlineResultHandlerCtx handles chosen inline result that came from bot with context
type ChosenInlineResultHandlerCtx func(ctx context.Context, bot *telego.Bot, result telego.ChosenInlineResult)

// HandleChosenInlineResult same as Handle, but assumes that the update contains a chosen inline result
func (h *HandlerGroup) HandleChosenInlineResult(handler ChosenInlineResultHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chosen inline query handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.ChosenInlineResult)
	}, append([]Predicate{AnyChosenInlineResult()}, predicates...)...)
}

// HandleChosenInlineResultCtx same as Handle, but assumes that the update contains a chosen inline result
func (h *HandlerGroup) HandleChosenInlineResultCtx(handler ChosenInlineResultHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chosen inline query handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.ChosenInlineResult)
	}, append([]Predicate{AnyChosenInlineResult()}, predicates...)...)
}

// HandleChosenInlineResult same as Handle, but assumes that the update contains a chosen inline result
func (h *BotHandler) HandleChosenInlineResult(handler ChosenInlineResultHandler, predicates ...Predicate) {
	h.baseGroup.HandleChosenInlineResult(handler, predicates...)
}

// HandleChosenInlineResultCtx same as Handle, but assumes that the update contains a chosen inline result
func (h *BotHandler) HandleChosenInlineResultCtx(handler ChosenInlineResultHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleChosenInlineResultCtx(handler, predicates...)
}

// CallbackQueryHandler handles callback queries that came from bot
type CallbackQueryHandler func(bot *telego.Bot, query telego.CallbackQuery)

// CallbackQueryHandlerCtx handles callback queries that came from bot with context
type CallbackQueryHandlerCtx func(ctx context.Context, bot *telego.Bot, query telego.CallbackQuery)

// HandleCallbackQuery same as Handle, but assumes that the update contains a callback query
func (h *HandlerGroup) HandleCallbackQuery(handler CallbackQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil callback query handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.CallbackQuery)
	}, append([]Predicate{AnyCallbackQuery()}, predicates...)...)
}

// HandleCallbackQueryCtx same as Handle, but assumes that the update contains a callback query
func (h *HandlerGroup) HandleCallbackQueryCtx(handler CallbackQueryHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil callback query handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.CallbackQuery)
	}, append([]Predicate{AnyCallbackQuery()}, predicates...)...)
}

// HandleCallbackQuery same as Handle, but assumes that the update contains a callback query
func (h *BotHandler) HandleCallbackQuery(handler CallbackQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandleCallbackQuery(handler, predicates...)
}

// HandleCallbackQueryCtx same as Handle, but assumes that the update contains a callback query
func (h *BotHandler) HandleCallbackQueryCtx(handler CallbackQueryHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleCallbackQueryCtx(handler, predicates...)
}

// ShippingQueryHandler handles shipping query that came from bot
type ShippingQueryHandler func(bot *telego.Bot, query telego.ShippingQuery)

// ShippingQueryHandlerCtx handles shipping query that came from bot with context
type ShippingQueryHandlerCtx func(ctx context.Context, bot *telego.Bot, query telego.ShippingQuery)

// HandleShippingQuery same as Handle, but assumes that the update contains a shipping query
func (h *HandlerGroup) HandleShippingQuery(handler ShippingQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil shipping query handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.ShippingQuery)
	}, append([]Predicate{AnyShippingQuery()}, predicates...)...)
}

// HandleShippingQueryCtx same as Handle, but assumes that the update contains a shipping query
func (h *HandlerGroup) HandleShippingQueryCtx(handler ShippingQueryHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil shipping query handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.ShippingQuery)
	}, append([]Predicate{AnyShippingQuery()}, predicates...)...)
}

// HandleShippingQuery same as Handle, but assumes that the update contains a shipping query
func (h *BotHandler) HandleShippingQuery(handler ShippingQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandleShippingQuery(handler, predicates...)
}

// HandleShippingQueryCtx same as Handle, but assumes that the update contains a shipping query
func (h *BotHandler) HandleShippingQueryCtx(handler ShippingQueryHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleShippingQueryCtx(handler, predicates...)
}

// PreCheckoutQueryHandler handles pre checkout query that came from bot
type PreCheckoutQueryHandler func(bot *telego.Bot, query telego.PreCheckoutQuery)

// PreCheckoutQueryHandlerCtx handles pre checkout query that came from bot with context
type PreCheckoutQueryHandlerCtx func(ctx context.Context, bot *telego.Bot, query telego.PreCheckoutQuery)

// HandlePreCheckoutQuery same as Handle, but assumes that the update contains a pre checkout query
func (h *HandlerGroup) HandlePreCheckoutQuery(handler PreCheckoutQueryHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil pre checkout query handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.PreCheckoutQuery)
	}, append([]Predicate{AnyPreCheckoutQuery()}, predicates...)...)
}

// HandlePreCheckoutQueryCtx same as Handle, but assumes that the update contains a pre checkout query
func (h *HandlerGroup) HandlePreCheckoutQueryCtx(handler PreCheckoutQueryHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil pre checkout query handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.PreCheckoutQuery)
	}, append([]Predicate{AnyPreCheckoutQuery()}, predicates...)...)
}

// HandlePreCheckoutQuery same as Handle, but assumes that the update contains a pre checkout query
func (h *BotHandler) HandlePreCheckoutQuery(handler PreCheckoutQueryHandler, predicates ...Predicate) {
	h.baseGroup.HandlePreCheckoutQuery(handler, predicates...)
}

// HandlePreCheckoutQueryCtx same as Handle, but assumes that the update contains a pre checkout query
func (h *BotHandler) HandlePreCheckoutQueryCtx(handler PreCheckoutQueryHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandlePreCheckoutQueryCtx(handler, predicates...)
}

// PollHandler handles poll that came from bot
type PollHandler func(bot *telego.Bot, poll telego.Poll)

// PollHandlerCtx handles poll that came from bot with context
type PollHandlerCtx func(ctx context.Context, bot *telego.Bot, poll telego.Poll)

// HandlePoll same as Handle, but assumes that the update contains a poll
func (h *HandlerGroup) HandlePoll(handler PollHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil poll handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.Poll)
	}, append([]Predicate{AnyPoll()}, predicates...)...)
}

// HandlePollCtx same as Handle, but assumes that the update contains a poll
func (h *HandlerGroup) HandlePollCtx(handler PollHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil poll handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.Poll)
	}, append([]Predicate{AnyPoll()}, predicates...)...)
}

// HandlePoll same as Handle, but assumes that the update contains a poll
func (h *BotHandler) HandlePoll(handler PollHandler, predicates ...Predicate) {
	h.baseGroup.HandlePoll(handler, predicates...)
}

// HandlePollCtx same as Handle, but assumes that the update contains a poll
func (h *BotHandler) HandlePollCtx(handler PollHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandlePollCtx(handler, predicates...)
}

// PollAnswerHandler handles poll answer that came from bot
type PollAnswerHandler func(bot *telego.Bot, answer telego.PollAnswer)

// PollAnswerHandlerCtx handles poll answer that came from bot with context
type PollAnswerHandlerCtx func(ctx context.Context, bot *telego.Bot, answer telego.PollAnswer)

// HandlePollAnswer same as Handle, but assumes that the update contains a poll answer
func (h *HandlerGroup) HandlePollAnswer(handler PollAnswerHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil poll answer handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.PollAnswer)
	}, append([]Predicate{AnyPollAnswer()}, predicates...)...)
}

// HandlePollAnswerCtx same as Handle, but assumes that the update contains a poll answer
func (h *HandlerGroup) HandlePollAnswerCtx(handler PollAnswerHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil poll answer handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.PollAnswer)
	}, append([]Predicate{AnyPollAnswer()}, predicates...)...)
}

// HandlePollAnswer same as Handle, but assumes that the update contains a poll answer
func (h *BotHandler) HandlePollAnswer(handler PollAnswerHandler, predicates ...Predicate) {
	h.baseGroup.HandlePollAnswer(handler, predicates...)
}

// HandlePollAnswerCtx same as Handle, but assumes that the update contains a poll answer
func (h *BotHandler) HandlePollAnswerCtx(handler PollAnswerHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandlePollAnswerCtx(handler, predicates...)
}

// ChatMemberUpdatedHandler handles chat member that came from bot
type ChatMemberUpdatedHandler func(bot *telego.Bot, chatMember telego.ChatMemberUpdated)

// ChatMemberUpdatedHandlerCtx handles chat member that came from bot with context
type ChatMemberUpdatedHandlerCtx func(ctx context.Context, bot *telego.Bot, chatMember telego.ChatMemberUpdated)

// HandleMyChatMemberUpdated same as Handle, but assumes that the update contains my chat member
func (h *HandlerGroup) HandleMyChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil my chat member update handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.MyChatMember)
	}, append([]Predicate{AnyMyChatMember()}, predicates...)...)
}

// HandleMyChatMemberUpdatedCtx same as Handle, but assumes that the update contains my chat member
func (h *HandlerGroup) HandleMyChatMemberUpdatedCtx(handler ChatMemberUpdatedHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil my chat member update handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.MyChatMember)
	}, append([]Predicate{AnyMyChatMember()}, predicates...)...)
}

// HandleMyChatMemberUpdated same as Handle, but assumes that the update contains my chat member
func (h *BotHandler) HandleMyChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	h.baseGroup.HandleMyChatMemberUpdated(handler, predicates...)
}

// HandleMyChatMemberUpdatedCtx same as Handle, but assumes that the update contains my chat member
func (h *BotHandler) HandleMyChatMemberUpdatedCtx(handler ChatMemberUpdatedHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleMyChatMemberUpdatedCtx(handler, predicates...)
}

// HandleChatMemberUpdated same as Handle, but assumes that the update contains chat member
func (h *HandlerGroup) HandleChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chat member update handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.ChatMember)
	}, append([]Predicate{AnyChatMember()}, predicates...)...)
}

// HandleChatMemberUpdatedCtx same as Handle, but assumes that the update contains chat member
func (h *HandlerGroup) HandleChatMemberUpdatedCtx(handler ChatMemberUpdatedHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chat member update handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.ChatMember)
	}, append([]Predicate{AnyChatMember()}, predicates...)...)
}

// HandleChatMemberUpdated same as Handle, but assumes that the update contains chat member
func (h *BotHandler) HandleChatMemberUpdated(handler ChatMemberUpdatedHandler, predicates ...Predicate) {
	h.baseGroup.HandleChatMemberUpdated(handler, predicates...)
}

// HandleChatMemberUpdatedCtx same as Handle, but assumes that the update contains chat member
func (h *BotHandler) HandleChatMemberUpdatedCtx(handler ChatMemberUpdatedHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleChatMemberUpdatedCtx(handler, predicates...)
}

// ChatJoinRequestHandler handles chat join request that came from bot
type ChatJoinRequestHandler func(bot *telego.Bot, request telego.ChatJoinRequest)

// ChatJoinRequestHandlerCtx handles chat join request that came from bot with context
type ChatJoinRequestHandlerCtx func(ctx context.Context, bot *telego.Bot, request telego.ChatJoinRequest)

// HandleChatJoinRequest same as Handle, but assumes that the update contains chat join request
func (h *HandlerGroup) HandleChatJoinRequest(handler ChatJoinRequestHandler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chat join request handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.ChatJoinRequest)
	}, append([]Predicate{AnyChatJoinRequest()}, predicates...)...)
}

// HandleChatJoinRequestCtx same as Handle, but assumes that the update contains chat join request
func (h *HandlerGroup) HandleChatJoinRequestCtx(handler ChatJoinRequestHandlerCtx, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil chat join request handlers not allowed")
	}

	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(update.Context(), bot, *update.ChatJoinRequest)
	}, append([]Predicate{AnyChatJoinRequest()}, predicates...)...)
}

// HandleChatJoinRequest same as Handle, but assumes that the update contains chat join request
func (h *BotHandler) HandleChatJoinRequest(handler ChatJoinRequestHandler, predicates ...Predicate) {
	h.baseGroup.HandleChatJoinRequest(handler, predicates...)
}

// HandleChatJoinRequestCtx same as Handle, but assumes that the update contains chat join request
func (h *BotHandler) HandleChatJoinRequestCtx(handler ChatJoinRequestHandlerCtx, predicates ...Predicate) {
	h.baseGroup.HandleChatJoinRequestCtx(handler, predicates...)
}
