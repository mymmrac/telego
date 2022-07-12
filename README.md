# Telego • Go Telegram Bot API

[![Go Reference](https://pkg.go.dev/badge/github.com/mymmrac/telego#section-readme.svg)](https://pkg.go.dev/github.com/mymmrac/telego)
[![Telegram Bot API Version][TelegramVersionBadge]][TelegramLastVersion]
<br>
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Dis](https://img.shields.io/github/discussions/mymmrac/telego?color=58a6ff&label=Discussions&logo=github)](https://github.com/mymmrac/telego/discussions)
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
  <img src="docs/Telego-long.png" alt="Telego logo" width="512px" style="border-radius: 12px;">
</p>

Telego is Telegram Bot API library for Golang with full [API][TelegramBotAPI] implementation (one-to-one)

The goal of this library was to create API with same types and methods as actual telegram bot API. Every type and method
have been represented in [`types.go`](types.go) and [`methods.go`](methods.go) files with mostly all documentation from
telegram.

> Note: Telego uses [fasthttp](https://github.com/valyala/fasthttp) instead of `net/http` and
> [go-json](https://github.com/goccy/go-json) instead of `encoding/json`.

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
- [:closed_lock_with_key: License](#closed_lock_with_key-license)

</details>

## :zap: Getting Started

[//]: # (TODO: Create Wiki page with Github Wikis or Github Pages)

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
- [Updates (long pulling)](examples/updates_long_pulling/main.go)
- [Updates (webhook)](examples/updates_webhook/main.go)
- [Utility methods](examples/utility_methods/main.go)
- [Bot handlers](examples/handler/main.go)
- [Bot more handlers (custom)](examples/handler_custom/main.go)
- [Echo bot](examples/echo_bot/main.go)
- [Sending files (documents, photos, media groups)](examples/sending_files/main.go)
- [Inline keyboard](examples/inline_keyboard/main.go)
- [Keyboard](examples/keyboard/main.go)
- [Inline query bot](examples/inline_query_bot/main.go)

</details>

> Note: Error handling may be missing in examples, but I strongly recommend handling all errors.

### :jigsaw: Basic setup

[▲ Go Up ▲](#telego--go-telegram-bot-api)

For start, you need to create instance of your bot and
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
	// (more on configuration at /examples/configuration/main.go)
	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
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
	fmt.Printf("Bot user: %#v\n", botUser)
}
```

### :envelope_with_arrow: Getting updates

[▲ Go Up ▲](#telego--go-telegram-bot-api)

In order to receive updates you can use two methods:

- using long polling (`bot.UpdatesViaLongPulling`)
- using webhook (`bot.UpdatesViaWebhook`)

Let's start from long pulling (easier for local testing):

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Optional. Set interval of getting updates (default: 0.5s).
	// If you want to get updates as fast as possible set to 0,
	// but webhook method is recommended for this.
	bot.SetUpdateInterval(time.Second / 2)

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPulling(nil)

	// Stop reviving updates from updates channel
	defer bot.StopLongPulling()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %#v\n", update)
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
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
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
	fmt.Printf("Webhook Info: %#v\n", info)

	// Get updates channel from webhook.
	// Note: For one bot only one webhook allowed.
	updates, _ := bot.UpdatesViaWebhook("/bot" + bot.Token())

	// Start server for receiving requests from Telegram
	bot.StartListeningForWebhook("localhost:443")

	// Stop reviving updates from updates channel and shutdown webhook server
	defer func() {
		_ = bot.StopWebhook()
	}()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %#v\n", update)
	}
}
```

> Note: You may wish to use [Let's Encrypt](https://letsencrypt.org/) in order to generate your free TLS certificate.

### :kite: Using Telegram methods

[▲ Go Up ▲](#telego--go-telegram-bot-api)

All Telegram Bot API methods described in [documentation](https://core.telegram.org/bots/api#available-methods) can be
used by the library. They have same names and same parameters, parameters represented by struct with
name: `<methodName>` + `Params`. If method don't have required parameters `nil` value can be used as a parameter.

> Note: [`types.go`](types.go) and [`methods.go`](methods.go) was automatically [generated](internal/generator)
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
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe
	botUser, _ := bot.GetMe()
	fmt.Printf("Bot User: %#v\n", botUser)

	updates, _ := bot.UpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()

	for update := range updates {
		if update.Message != nil {
			// Retrieve chat ID
			chatID := update.Message.Chat.ID

			// Call method sendMessage (https://core.telegram.org/bots/api#sendmessage).
			// Sends message to sender with same text (echo bot).
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

Package contains couple methods for creating send parameters with all required parameters like:

- `Message(chatID, text) => SendMessageParams`
- `Photo(chatID, photoFile) => SendPhotoParams`
- `Location(chatID, latitude, longitude) => SendLocationParams`
- ...

Or other useful methods like:

- `ID(intID) => ChatID`
- `File(namedReader) => InputFile`
- ...

Utils related to [`methods`](methods.go) can be found in [`telegoutil/methods`](telegoutil/methods.go), for
[`types`](types.go) in [`telegoutil/types`](telegoutil/types.go), for [`handlers`](telegohandler/handler.go) in
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
	// ... initializing bot

	// Creating keyboard
	keyboard := tu.Keyboard(
		tu.KeyboardRow( // Row 1
			tu.KeyboardButton("Button"), // Column 1
			tu.KeyboardButton("Poll Regular"). // Column 2
				WithRequestPoll(tu.PollTypeRegular()), // <- `with` method
		),
		tu.KeyboardRow( // Row 2
			tu.KeyboardButton("Contact").WithRequestContact(),   // Column 1, <- `with` method 
			tu.KeyboardButton("Location").WithRequestLocation(), // Column 2, <- `with` method 
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Select something") // <- multiple `with` methods 

	// Creating message
	msg := tu.Message(
		tu.ID(123),
		"Hello World",
	).WithReplyMarkup(keyboard).WithProtectContent() // <- multiple `with` method 

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
registration (it's useful to first specify most specific predicates and then more general).

Also, all handlers (but not their predicates) are processed in parallel.

I suggest including it with alias to get cleaner code:

```go
import th "github.com/mymmrac/telego/telegohandler"
```

Here is example of using handlers with long pulling updates. You can see full list of available predicates
in [`telegohandler/pradicates`](telegohandler/pradicates.go), or define your own.

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
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Register new handler with match on command `/start`
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
		))
	}, th.CommandEqual("start"))

	// Register new handler with match on any command
	// Handlers will match only once and in order of registration, so this handler will be called on any command except
	// `/start` command
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Unknown command, use /start",
		))
	}, th.AnyCommand())

	// Start handling updates
	bh.Start()

	// Stop handling updates
	defer bh.Stop()
}
```

## :art: Contribution

Contribution guidelines listed [here](docs/CONTRIBUTING.md).

## :closed_lock_with_key: License

Telego is distributed under [MIT licence](LICENSE).

[TelegramBotAPI]: https://core.telegram.org/bots/api

[TelegramVersionBadge]: https://img.shields.io/static/v1?label=Supported%20Telegram%20Bot%20API&color=29a1d4&logo=telegram&message=v6.1

[TelegramLastVersion]: https://core.telegram.org/bots/api#june-20-2022
