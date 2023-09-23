# Telego • Go Telegram Bot API

[![Go Reference](https://pkg.go.dev/badge/github.com/mymmrac/telego#section-readme.svg)](https://pkg.go.dev/github.com/mymmrac/telego)
[![Telego Docs](https://img.shields.io/static/v1?label=Telego&message=docs&color=8ed6fb&logo=hugo)](https://telego.pixelbox.dev)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mymmrac/telego?logo=go)](go.mod)
[![Telegram Bot API Version][TelegramVersionBadge]][TelegramLastVersion]
<br>
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Discussions](https://img.shields.io/github/discussions/mymmrac/telego?color=58a6ff&label=Discussions&logo=github)](https://github.com/mymmrac/telego/discussions)
[![Telegram Chat](https://img.shields.io/static/v1?label=Discussion&message=chat&color=29a1d4&logo=telegram)](https://t.me/telegoLibrary)

[![CI Status](https://github.com/mymmrac/telego/actions/workflows/ci.yml/badge.svg)](https://github.com/mymmrac/telego/actions/workflows/ci.yml)
[![Race Testing](https://github.com/mymmrac/telego/actions/workflows/race-tests.yml/badge.svg)](https://github.com/mymmrac/telego/actions/workflows/race-tests.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_telego&metric=alert_status)](https://sonarcloud.io/dashboard?id=mymmrac_telego)
[![Go Report](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/mymmrac/telego)
<br>
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_telego&metric=coverage)](https://sonarcloud.io/dashboard?id=mymmrac_telego)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_telego&metric=code_smells)](https://sonarcloud.io/dashboard?id=mymmrac_telego)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_telego&metric=ncloc)](https://sonarcloud.io/dashboard?id=mymmrac_telego)

<p align="center">
  <img src="docs/logo/telego-long.png" alt="Telego logo" width="512px" style="border-radius: 12px;">
</p>

Telego is Telegram Bot API library for Golang with full [API][TelegramBotAPI] implementation (one-to-one)

The goal of this library was to create API with the same types and methods as actual Telegram Bot API.
Every type and method have been represented in [`types.go`](types.go) and [`methods.go`](methods.go) files with mostly
all documentation from Telegram.

:warning: Telego is still in v0.x.x version, so do expect breaking changes! :warning:

For more detailed documentation, see docs at [telego.pixelbox.dev](https://telego.pixelbox.dev).

> Note: Telego uses [fasthttp](https://github.com/valyala/fasthttp) instead of `net/http` by default (can be changed)
> and [go-json](https://github.com/goccy/go-json) instead of `encoding/json`.

### :clipboard: Table Of Content

<details>
<summary>Click to show • hide</summary>

- [:zap: Getting Started](#zap-getting-started)
    - [:jigsaw: Basic setup](#jigsaw-basic-setup)
    - [:envelope_with_arrow: Getting updates](#envelope_with_arrow-getting-updates)
    - [:kite: Using Telegram methods](#kite-using-telegram-methods)
    - [:soap: Utility methods](#soap-utility-methods)
    - [:mechanical_arm: Helper `With...` methods](#mechanical_arm-helper-with-methods)
    - [:sun_behind_large_cloud: Bot handlers](#sun_behind_large_cloud-bot-handlers)
- [:art: Contribution](#art-contribution)
- [:star: Stargazers over time](#star-stargazers-over-time)
- [:closed_lock_with_key: License](#closed_lock_with_key-license)

</details>

## :zap: Getting Started

How to get the library:

```shell
go get -u github.com/mymmrac/telego
```

Make sure you get the latest version to have all new features & fixes.

More examples can be seen here:

<details>
<summary>Click to show • hide</summary>

- [Basic](examples/basic/main.go)
- [Configuration](examples/configuration/main.go)
- [Methods](examples/methods/main.go)
- [Updates (long polling)](examples/updates_long_polling/main.go)
- [Updates (webhook)](examples/updates_webhook/main.go)
- [Ngrok webhook](examples/ngrok/main.go)
- [Echo bot](examples/echo_bot/main.go)
- [Echo bot (with handlers)](examples/echo_bot_with_handlers/main.go)
- [Echo bot (handlers + webhook + graceful shutdown + docker)](https://github.com/mymmrac/echo-bot)
- [Sending files (documents, photos, media groups)](examples/sending_files/main.go)
- [Downloading files](examples/download_file/main.go)
- [Inline keyboard](examples/inline_keyboard/main.go)
- [Keyboard](examples/keyboard/main.go)
- [Edit message](examples/edit_message/main.go)
- [Utility methods](examples/utility_methods/main.go)
- [Inline query bot](examples/inline_query_bot/main.go)
- [Bot handlers](examples/handler/main.go)
- [Bot handles (groups + middleware)](examples/handler_groups_and_middleware/main.go)
- [Update's context](examples/handler_with_context/main.go)
- [Graceful shutdown (no helpers)](examples/graceful_shutdown_no_helpers/main.go)
- [Graceful shutdown (long polling)](examples/graceful_shutdown_long_polling/main.go)
- [Graceful shutdown (webhook)](examples/graceful_shutdown_webhook/main.go)
- [Custom predicates for handlers](examples/handler_custom/main.go)
- [Handler ordering](examples/handler_ordering/main.go)
- [Specific handlers](examples/handler_specific/main.go)
- [Predicate as middleware](examples/middleware_with_predicates/main.go)
- [Update processor](examples/update_processor/main.go)
- [Message entities](examples/message_entity/main.go)
- [Multi bot webhook](examples/multi_bot_webhook/main.go)
- [Retry caller](examples/retry_caller/main.go)
- [Menu bot](examples/menu_bot/main.go)
- [Test server](examples/test_server/main.go)

</details>

> Note: Error handling may be missing in examples, but I strongly recommend handling all errors.

Generally, useful information about Telegram Bots and their features:

<details>
<summary>Click to show • hide</summary>

- [:page_facing_up: Telegram Bot API](https://core.telegram.org/bots/api) • Telegram documentation of Bot API (full
  reference)
- [:jigsaw: Telegram Bot Fundamentals](https://core.telegram.org/bots) • Bots: An introduction for developers
- [:star2: Telegram Bot Features](https://core.telegram.org/bots/features) • Describes individual bot elements and
  features in detail
- [:headphones: Telegram Bot Webhooks](https://core.telegram.org/bots/webhooks) • Marvin's Marvellous Guide to All
  Things Webhook
- [:moneybag: Telegram Bot Payments](https://core.telegram.org/bots/payments) • Describes payment system and payment
  lifecycle
- [:iphone: Telegram Bot WebApps](https://core.telegram.org/bots/webapps) • Describes WebApps and interactions with them
- [:link: Ngrok](https://ngrok.com) • Connect localhost to the Internet
- [:shield: Let's Encrypt](https://letsencrypt.org) • TLS certificates for free

</details>

### :jigsaw: Basic setup

[▲ Go Up ▲](#telego--go-telegram-bot-api)

For start, you need to create an instance of your bot and
specify [token](https://core.telegram.org/bots/api#authorizing-your-bot).

```go
package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
)

func main() {
	// Get Bot token from environment variables
	botToken := os.Getenv("TOKEN")

	// Create bot and enable debugging info
	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	// (more on configuration in examples/configuration/main.go)
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe (https://core.telegram.org/bots/api#getme)
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print Bot information
	fmt.Printf("Bot user: %+v\n", botUser)
}
```

### :envelope_with_arrow: Getting updates

[▲ Go Up ▲](#telego--go-telegram-bot-api)

In order to receive updates, you can use one of two methods:

- using long polling (`bot.UpdatesViaLongPolling`)
- using webhook (`bot.UpdatesViaWebhook`)

Let's start from long polling (easier for local testing):

```go
package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	// (more on configuration in examples/updates_long_polling/main.go)
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Stop reviving updates from update channel
	defer bot.StopLongPolling()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %+v\n", update)
	}
}
```

Webhook example (recommended way):

```go
package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set up a webhook on Telegram side
	_ = bot.SetWebhook(&telego.SetWebhookParams{
		URL: "https://example.com/bot" + bot.Token(),
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo()
	fmt.Printf("Webhook Info: %+v\n", info)

	// Get an update channel from webhook.
	// (more on configuration in examples/updates_webhook/main.go)
	updates, _ := bot.UpdatesViaWebhook("/bot" + bot.Token())

	// Start server for receiving requests from the Telegram
	go func() {
		_ = bot.StartWebhook("localhost:443")
	}()

	// Stop reviving updates from update channel and shutdown webhook server
	defer func() {
		_ = bot.StopWebhook()
	}()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %+v\n", update)
	}
}
```

For running multiple bots from a single server, see [this](examples/multi_bot_webhook/main.go) example.

> Tip: For testing webhooks, you can use [Ngrok](https://ngrok.com) to make a tunnel to your localhost,
> and get a random domain available from the Internet.
> It's as simple as `ngrok http 8080`.
> Or follow [Telego + Ngrok example](examples/ngrok/main.go) using [ngrok/ngrok-go](https://github.com/ngrok/ngrok-go)
> for most convenient bot testing.

> Tip: You may wish to use [Let's Encrypt](https://letsencrypt.org) in order to generate your free TLS certificate.

### :kite: Using Telegram methods

[▲ Go Up ▲](#telego--go-telegram-bot-api)

All Telegram Bot API methods described in [documentation](https://core.telegram.org/bots/api#available-methods) can be
used by the library. They have the same names and the same parameters, parameters represented by struct with
name: `<methodName>` + `Params`. If method doesn't have required parameters `nil` value can be used as a parameter.

> Note: [`types.go`](types.go) and [`methods.go`](methods.go) were automatically [generated](internal/generator)
> from [documentation][TelegramBotAPI], and it's possible that they have errors or missing parts both in comments and
> actual code. Feel free to report such things.

```go
package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe
	botUser, _ := bot.GetMe()
	fmt.Printf("Bot User: %+v\n", botUser)

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			// Retrieve chat ID
			chatID := update.Message.Chat.ID

			// Call method sendMessage.
			// Send a message to sender with the same text (echo bot).
			// (https://core.telegram.org/bots/api#sendmessage)
			sentMessage, _ := bot.SendMessage(
				tu.Message(
					tu.ID(chatID),
					update.Message.Text,
				),
			)

			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}
}
```

### :soap: Utility methods

[▲ Go Up ▲](#telego--go-telegram-bot-api)

In Telego even though you have all [`types`](types.go) and [`methods`](methods.go) available, it's often not so
convenient to use them directly. To solve this issues [`telegoutil`](telegoutil) package was created. It contains
utility-helper function that will make your life a bit easier.

I suggest including it with alias to get cleaner code:

```go
import tu "github.com/mymmrac/telego/telegoutil"
```

The package contains couple methods for creating send parameters with all required parameters like:

- `Message(chatID, text) => SendMessageParams`
- `Photo(chatID, photoFile) => SendPhotoParams`
- `Location(chatID, latitude, longitude) => SendLocationParams`
- ...

Or other useful methods like:

- `ID(intID) => ChatID`
- `File(namedReader) => InputFile`
- ...

Utils related to [`methods`](methods.go) can be found in [`telegoutil/methods`](telegoutil/methods.go), for
[`types`](types.go) in [`telegoutil/types`](telegoutil/types.go), for [`handlers`](telegohandler/bot_handler.go) in
[`telegoutil/handler`](telegoutil/handler.go), for [`api`](telegoapi/api.go) in [`telegoutil/api`](telegoutil/api.go).

> Note: If you think that something can be added to [`telegoutil`](telegoutil) package
> fill free to create an issue or pull request with desired changes.

### :mechanical_arm: Helper `With...` methods

[▲ Go Up ▲](#telego--go-telegram-bot-api)

Creating method parameters is sometimes bulky and not convenient, so you can use `with` methods in combination with
`utility` methods.

Here is a simple example of creating a message with a keyboard that has 4 buttons with different parameters.

```go
package main

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	// ... initializing bot (full example in examples/keyboard/main.go)

	// Creating keyboard
	keyboard := tu.Keyboard(
		tu.KeyboardRow( // Row 1
			// Column 1
			tu.KeyboardButton("Button"),

			// Column 2, `with` method
			tu.KeyboardButton("Poll Regular").
				WithRequestPoll(tu.PollTypeRegular()),
		),
		tu.KeyboardRow( // Row 2
			// Column 1, `with` method 
			tu.KeyboardButton("Contact").WithRequestContact(),

			// Column 2, `with` method 
			tu.KeyboardButton("Location").WithRequestLocation(),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")
	// Multiple `with` methods can be chained

	// Creating message
	msg := tu.Message(
		tu.ID(123),
		"Hello World",
	).WithReplyMarkup(keyboard).WithProtectContent() // Multiple `with` method 

	bot.SendMessage(msg)
}
```

Those methods allow you to modify values without directly accessing them, also as you saw `with` methods can be staked
one to another in order to update multiple values.

### :sun_behind_large_cloud: Bot handlers

[▲ Go Up ▲](#telego--go-telegram-bot-api)

Processing updates just in for loop is not the most pleasing thing to do, so Telego provides `net/http` like handlers,
but instead of the path, you provide predicates.

One update will only match to the first handler whose predicates are satisfied, predicates checked in order of handler
registration (it's useful to first specify the most specific predicates and then more general).

Also, all handlers (but not their predicates) are processed in parallel.

I suggest including it with alias to get cleaner code:

```go
import th "github.com/mymmrac/telego/telegohandler"
```

Here is an example of using handlers with long polling updates.
You can see the full list of available predicates in [`telegohandler/predicates`](telegohandler/predicates.go),
or define your own.

```go
package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer bh.Stop()

	// Stop getting updates
	defer bot.StopLongPolling()

	// Register new handler with match on command `/start`
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
		))
	}, th.CommandEqual("start"))

	// Register new handler with match on any command
	// Handlers will match only once and in order of registration, 
	// so this handler will be called on any command except `/start` command
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Unknown command, use /start",
		))
	}, th.AnyCommand())

	// Start handling updates
	bh.Start()
}
```

Also, just handling updates is useful, but handling specific updates like messages or callback queries in most of the
cases are more straightforward and provides cleaner code.

So Telego provides specific handles for all fields of `telego.Update`. See the list of all available handler types in
[`telegohandler/update_handlers`](telegohandler/update_handlers.go), or define your own.

```go
package main

import (
	"fmt"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	// ... initializing bot and bot handler 
	// (full example in examples/handler_specific/main.go)

	// Register new handler with match on command `/start`
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		// Send a message with inline keyboard
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(message.Chat.ID),
			fmt.Sprintf("Hello %s!", message.From.FirstName),
		).WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Go!").WithCallbackData("go"),
			)),
		))
	}, th.CommandEqual("start"))

	// Register new handler with match on the call back query 
	// with data equal to `go` and non-nil message
	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.Chat.ID), "GO GO GO"))

		// Answer callback query
		_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText("Done"))
	}, th.AnyCallbackQueryWithMessage(), th.CallbackDataEqual("go"))

	// ... start bot handler
}
```

One more important part of handlers are groups and middlewares.
Telego allows creating groups with and without predicates and attaching middleware to groups.

```go
package main

import (
	"fmt"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	// Init ...

	// Add global middleware, it will be applied in order of addition
	bh.Use(func(bot *telego.Bot, update telego.Update, next th.Handler) {
		fmt.Println("Global middleware") // Will be called first
		next(bot, update)
	})

	// Create any groups with or without predicates
	// Note: Updates first checked by groups and only then by handlers 
	// (group -> ... -> group -> handler)
	task := bh.Group(th.TextContains("task"))

	// Add middleware to groups
	task.Use(func(bot *telego.Bot, update telego.Update, next th.Handler) {
		fmt.Println("Group-based middleware") // Will be called second

		if len(update.Message.Text) < 10 {
			next(bot, update)
		}
	})

	// Handle updates on a group
	task.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		fmt.Println("Task...") // Will be called third
	})
}

```

## :art: Contribution

Contribution guidelines listed [here](docs/CONTRIBUTING.md).

## :star: Stargazers over time

[![Stargazers over time](https://starchart.cc/mymmrac/telego.svg)](https://starchart.cc/mymmrac/telego)

> Powered by [caarlos0/starcharts](https://github.com/caarlos0/starcharts)

## :closed_lock_with_key: License

Telego is distributed under [MIT licence](LICENSE).

[TelegramBotAPI]: https://core.telegram.org/bots/api

[TelegramVersionBadge]: https://img.shields.io/static/v1?label=Supported%20Telegram%20Bot%20API&color=29a1d4&logo=telegram&message=v6.9

[TelegramLastVersion]: https://core.telegram.org/bots/api#september-22-2023
