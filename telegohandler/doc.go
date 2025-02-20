/*
Package telegohandler provides handlers and predicates for Telego.

Bot handlers provide an easy way to make net/http like handlers, but with predicates instead of paths.
In addition to just handlers and predicates, it also provides groups and middlewares.

You can create [BotHandler] register new handlers and start processing updates from the update channel.
All handlers process updates concurrently, but keep in mind that predicates are checked sequentially.
This gives an ability to process one update only with the first matched handler.

# Example

This example shows how you can create [BotHandler] and register new handlers.
Note that order of registration directly impacts the order of checking matched handlers, and only the first matched
handler will process the update, however, middlewares can run even if no handler matched since we are searching for any
handler in any group to process the update.

	package main

	import (
		"fmt"
		"os"

		"github.com/mymmrac/telego"
		th "github.com/mymmrac/telego/telegohandler"
		tu "github.com/mymmrac/telego/telegoutil"
	)

	func main() {
		ctx := context.Background()
		botToken := os.Getenv("TOKEN")

		// Note: Please keep in mind that default logger may expose sensitive information, use in development only
		bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Get updates channel
		updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

		// Create bot handler and specify from where to get updates
		bh, _ := th.NewBotHandler(bot, updates)

		// Stop handling updates
		defer func() { _ = bh.Stop() }()

		// Register new handler with match on command `/start`
		bh.Handle(func(ctx *th.Context, update telego.Update) error {
			// Send message
			_, _ = ctx.Bot().SendMessage(ctx, tu.Messagef(
				tu.ID(update.Message.Chat.ID),
				"Hello %s!", update.Message.From.FirstName,
			))
			return nil
		}, th.CommandEqual("start"))

		// Register new handler with match on any command
		// Handlers will match only once and in order of registration, so this handler will be called on any command
		// except `/start` command
		bh.Handle(func(ctx *th.Context, update telego.Update) error {
			// Send message
			_, _ = ctx.Bot().SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Unknown command, use /start",
			))
			return nil
		}, th.AnyCommand())

		// Start handling updates
		_ = bh.Start()
	}

One more example of handler usage.
It shows how to use specific handlers to process individual fields of [telego.Update].

	package main

	import (
		"fmt"
		"os"

		"github.com/mymmrac/telego"
		th "github.com/mymmrac/telego/telegohandler"
		tu "github.com/mymmrac/telego/telegoutil"
	)

	func main() {
		ctx := context.Background()
		botToken := os.Getenv("TOKEN")

		// Note: Please keep in mind that default logger may expose sensitive information, use in development only
		bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Get updates channel
		updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

		// Create bot handler and specify from where to get updates
		bh, _ := th.NewBotHandler(bot, updates)

		// Stop handling updates
		defer func() { _ = bh.Stop() }()

		// Register new handler with match on command `/start`
		bh.HandleMessage(func(ctx *th.Context, message telego.Message) error {
			// Send a message with inline keyboard
			_, _ = ctx.Bot().SendMessage(tu.Messagef(
				tu.ID(message.Chat.ID),
				"Hello %s!", message.From.FirstName,
			).WithReplyMarkup(tu.InlineKeyboard(
				tu.InlineKeyboardRow(tu.InlineKeyboardButton("Go!").WithCallbackData("go"))),
			))
			return nil
		}, th.CommandEqual("start"))

		// Register new handler with match on a call back query with data equal to `go` and non-nil message
		bh.HandleCallbackQuery(func(ctx *th.Context, query telego.CallbackQuery) error {
			// Send message
			_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.Chat.ID), "GO GO GO"))

			// Answer callback query
			_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText("Done"))

			return nil
		}, th.AnyCallbackQueryWithMessage(), th.CallbackDataEqual("go"))

		// Start handling updates
		_ = bh.Start()
	}

In this example, usage of groups and middleware will be shown.

	package main

	import (
		"fmt"
		"os"

		"github.com/mymmrac/telego"
		th "github.com/mymmrac/telego/telegohandler"
	)

	func main() {
		ctx := context.Background()
		botToken := os.Getenv("TOKEN")

		// Note: Please keep in mind that default logger may expose sensitive information, use in development only
		bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Get updates channel
		updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

		// Create bot handler and specify from where to get updates
		bh, _ := th.NewBotHandler(bot, updates)

		// Stop handling updates
		defer func() { _ = bh.Stop() }()

		// Add global middleware, it will be applied in order of addition
		bh.Use(
			func(ctx *th.Context, update telego.Update) error {
				fmt.Println("Global middleware") // Will be called first
				return ctx.Next(update)
			},
			func(ctx *th.Context, update telego.Update) error {
				fmt.Println("Global middleware 2") // Will be called second
				return ctx.Next(update)
			},
		)

		// Create any groups with or without predicates
		task := bh.Group(th.TextContains("task"))

		// Add middleware to groups
		task.Use(func(ctx *th.Context, update telego.Update) error {
			fmt.Println("Group based middleware") // Will be called third

			if len(update.Message.Text) < 10 {
				return ctx.Next(update)
			}

			return nil
		})

		// Handle updates on a group
		task.HandleMessage(func(ctx *th.Context, message telego.Message) error {
			fmt.Println("Task...") // Will be called fourth
			return nil
		})

		// Start handling updates
		_ = bh.Start()
	}
*/
package telegohandler
