package main

import (
	"log"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func RegisterHandlers(bh *th.BotHandler) {
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, err := bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Menu").
			WithReplyMarkup(tu.Keyboard(
				tu.KeyboardRow(
					tu.KeyboardButton("Sub menu 1"),
					tu.KeyboardButton("Sub menu 2"),
				),
				tu.KeyboardRow(
					tu.KeyboardButton("Sub menu 3"),
				),
			).WithResizeKeyboard()))
		if err != nil {
			log.Printf("Error on start: %s", err)
		}
	}, th.Union(th.CommandEqual("start"), th.TextEqual("Back")))

	subMenu := bh.Group(th.TextPrefix("Sub menu"))
	subMenu.Use(func(bot *telego.Bot, update telego.Update, next th.Handler) {
		log.Println("Sub menu group")
		next(bot, update)
	})

	subMenu.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, err := bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Sub menu 1 content").
			WithReplyMarkup(tu.Keyboard(tu.KeyboardRow(tu.KeyboardButton("Back"))).WithResizeKeyboard()))
		if err != nil {
			log.Printf("Error on sub menu 1: %s", err)
		}
	}, th.TextSuffix("1"))

	subMenu.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, err := bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Sub menu 2 content").
			WithReplyMarkup(tu.Keyboard(tu.KeyboardRow(tu.KeyboardButton("Back"))).WithResizeKeyboard()))
		if err != nil {
			log.Printf("Error on sub menu 2: %s", err)
		}
	}, th.TextSuffix("2"))

	subMenu.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, err := bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Sub menu 3 content").
			WithReplyMarkup(tu.Keyboard(tu.KeyboardRow(tu.KeyboardButton("Back"))).WithResizeKeyboard()))
		if err != nil {
			log.Printf("Error on sub menu 3: %s", err)
		}
	}, th.TextSuffix("3"))
}
