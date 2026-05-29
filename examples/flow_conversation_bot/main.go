package main

import (
	"context"
	"fmt"
	"net/mail"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	tf "github.com/mymmrac/telego/telegoflow"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// UserData is typed session data shared between flow steps.
type UserData struct {
	Name  string
	Age   uint
	Email string
}

func main() {
	ctx := context.Background()
	botToken := os.Getenv("TOKEN")

	// Create Bot with debug on
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	// Create bot handler
	bh, _ := th.NewBotHandler(bot, updates)

	// Create flow manager with in-memory storage. In production, use persistent storage to restore sessions
	// after restarts (PostgreSQL, Redis, etc.).
	flows := tf.NewManager(tf.NewMemoryStorage())

	registration, err := newRegistrationFlow()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Commands that should work inside a flow should be registered before flow middleware.
	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		if err := flows.Cancel(ctx, update); err != nil {
			_, sendErr := ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), "No active conversation"))
			return sendErr
		}

		_, err := ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), "Conversation canceled"))
		return err
	}, th.CommandEqual("cancel"))

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		_, err := ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), registration.Graph()))
		return err
	}, th.CommandEqual("graph"))

	// Active sessions are handled by flow middleware. If no session is active, processing continues to next handlers.
	bh.Use(flows.Middleware())

	// Register flow and start it with /start.
	if err = flows.Register(registration); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bh.Handle(registration.Start, th.CommandEqual("start"))

	// Fallback for regular messages when no flow is active.
	bh.HandleMessage(func(ctx *th.Context, msg telego.Message) error {
		_, err := ctx.Bot().SendMessage(ctx, tu.Message(
			msg.Chat.ChatID(),
			"Use /start to begin the flow conversation, /cancel to cancel it, or /graph to print the flow graph.",
		))
		return err
	})

	fmt.Println(registration.Graph())

	// Stop handling updates on exit
	defer func() { _ = bh.Stop() }()

	// Start handling updates
	_ = bh.Start()
}

func newRegistrationFlow() (*tf.Flow[UserData], error) {
	return tf.New[UserData]("registration").
		Steps(
			tf.NewStep[UserData]("name").
				Enter(func(ctx *tf.Context[UserData]) error {
					_, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Hello stranger, what's your name?"))
					return err
				}).
				Handle(func(ctx *tf.Context[UserData]) error {
					ctx.Data().Name = ctx.Text()
					return ctx.Go("age")
				}, th.AnyMessageWithText()).
				CanGo("age"),

			tf.NewStep[UserData]("age").
				Enter(func(ctx *tf.Context[UserData]) error {
					_, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "How old are you?"))
					return err
				}).
				Handle(func(ctx *tf.Context[UserData]) error {
					age, err := strconv.ParseUint(ctx.Text(), 10, 32)
					if err != nil || age == 0 {
						_, err = ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Invalid age, please try again"))
						return err
					}

					ctx.Data().Age = uint(age)
					return ctx.Go("email")
				}, th.AnyMessageWithText()).
				CanGo("email"),

			tf.NewStep[UserData]("email").
				Enter(func(ctx *tf.Context[UserData]) error {
					_, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "What's your email?"))
					return err
				}).
				Handle(func(ctx *tf.Context[UserData]) error {
					addr, err := mail.ParseAddress(ctx.Text())
					if err != nil {
						_, err = ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Invalid email, please try again"))
						return err
					}

					ctx.Data().Email = addr.Address
					return ctx.Go("confirm")
				}, th.AnyMessageWithText()).
				CanGo("confirm"),

			tf.NewStep[UserData]("confirm").
				Enter(func(ctx *tf.Context[UserData]) error {
					data := ctx.Data()
					keyboard := tu.InlineKeyboard(
						tu.InlineKeyboardRow(
							tu.InlineKeyboardButton("Confirm").WithCallbackData("confirm_yes"),
							tu.InlineKeyboardButton("Start again").WithCallbackData("confirm_no"),
						),
					)

					_, err := ctx.Bot().SendMessage(ctx, tu.Messagef(
						ctx.ChatID(),
						"Your name is %s, your age is %d, and your email is %s, all correct?",
						data.Name,
						data.Age,
						data.Email,
					).WithReplyMarkup(keyboard))
					return err
				}).
				Handle(func(ctx *tf.Context[UserData]) error {
					query := ctx.CallbackQuery()
					if query == nil {
						return nil
					}

					_ = ctx.Bot().AnswerCallbackQuery(ctx, tu.CallbackQuery(query.ID))

					switch query.Data {
					case "confirm_yes":
						return ctx.Complete()
					case "confirm_no":
						return ctx.Go("name")
					default:
						return nil
					}
				}, th.AnyCallbackQuery(), th.Or(th.CallbackDataEqual("confirm_yes"), th.CallbackDataEqual("confirm_no"))).
				Fallback(func(ctx *tf.Context[UserData]) error {
					_, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Please use the buttons below"))
					return err
				}).
				CanGo("name").CanComplete(),
		).
		StartWith("name").
		OnComplete(func(ctx *tf.Context[UserData]) error {
			_, err := ctx.Bot().SendMessage(ctx, tu.Message(ctx.ChatID(), "Thanks for your data!"))
			return err
		}).
		Build()
}
