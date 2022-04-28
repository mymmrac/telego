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

// CallbackQueryHandler handles callback queries that came from bot
type CallbackQueryHandler func(bot *telego.Bot, message telego.CallbackQuery)

// HandleCallbackQuery same as Handle, but assumes that the update contains a callback query
func (h *BotHandler) HandleCallbackQuery(handler CallbackQueryHandler, predicates ...Predicate) {
	h.Handle(func(bot *telego.Bot, update telego.Update) {
		handler(bot, *update.CallbackQuery)
	}, append([]Predicate{AnyCallbackQuery()}, predicates...)...)
}

// TODO: Add more handlers
