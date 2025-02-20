package main

import (
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func registerHandlers(bh *th.BotHandler) {
	bh.Use(func(ctx *th.Context, update telego.Update) error {
		if err := ctx.Next(update); err != nil {
			log.Printf("Handler error: %s", err)
		}
		return nil
	})

	bh.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		_, err := ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(message.Chat.ID), "Menu").
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
			return fmt.Errorf("start: %w", err)
		}
		return nil
	}, th.Or(th.CommandEqual("start"), th.TextEqual("Back")))

	subMenu := bh.Group(th.TextPrefix("Sub menu"))
	subMenu.Use(func(ctx *th.Context, update telego.Update) error {
		log.Println("Sub menu group")
		return ctx.Next(update)
	})

	subMenu.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		_, err := ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(message.Chat.ID), "Sub menu 1 content").
			WithReplyMarkup(tu.Keyboard(tu.KeyboardRow(tu.KeyboardButton("Back"))).WithResizeKeyboard()))
		if err != nil {
			return fmt.Errorf("sub menu 1: %w", err)
		}
		return nil
	}, th.TextSuffix("1"))

	subMenu.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		_, err := ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(message.Chat.ID), "Sub menu 2 content").
			WithReplyMarkup(tu.Keyboard(tu.KeyboardRow(tu.KeyboardButton("Back"))).WithResizeKeyboard()))
		if err != nil {
			return fmt.Errorf("sub menu 2: %w", err)
		}
		return nil
	}, th.TextSuffix("2"))

	subMenu.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		_, err := ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(message.Chat.ID), "Sub menu 3 content").
			WithReplyMarkup(tu.Keyboard(tu.KeyboardRow(tu.KeyboardButton("Back"))).WithResizeKeyboard()))
		if err != nil {
			return fmt.Errorf("sub menu 3: %w", err)
		}
		return nil
	}, th.TextSuffix("3"))
}
