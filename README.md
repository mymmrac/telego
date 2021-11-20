# Telego | Go Telegram Bot API

[![Go Reference](https://pkg.go.dev/badge/github.com/mymmrac/telego#section-readme.svg)](https://pkg.go.dev/github.com/mymmrac/telego)
[![CI Status](https://github.com/mymmrac/telego/actions/workflows/ci.yml/badge.svg)](https://github.com/mymmrac/telego/actions/workflows/ci.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_telego&metric=alert_status)](https://sonarcloud.io/dashboard?id=mymmrac_telego)
[![Telegram Bot API Version][TelegramVersionBadge]][TelegramLastVersion]
[![Telegram Chat](https://img.shields.io/static/v1?label=Discussion&message=chat&color=29a1d4&logo=telegram)](https://t.me/telegoLibrary)

[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_telego&metric=coverage)](https://sonarcloud.io/dashboard?id=mymmrac_telego)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_telego&metric=bugs)](https://sonarcloud.io/dashboard?id=mymmrac_telego)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_telego&metric=code_smells)](https://sonarcloud.io/dashboard?id=mymmrac_telego)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_telego&metric=ncloc)](https://sonarcloud.io/dashboard?id=mymmrac_telego)

<p align="center">
  <img src="https://i.ibb.co/sChtxzh/Telego-long.png" alt="Telego logo" width="512" style="border-radius: 12px;">
</p>

Telego is Telegram Bot API library for Golang with full [API][TelegramBotAPI] implementation (one-to-one)

The goal of this library was to create API with same types and methods as actual telegram bot API. Every type and method
have been represented in [`types.go`](types.go) and [`methods.go`](methods.go) files with mostly all documentation from
telegram.

> Note: Telego uses [fasthttp](https://github.com/valyala/fasthttp) instead of `net/http` and [jsoniter](https://github.com/json-iterator/go) instead of `encoding/json`.

### ToDo List & Ideas

<details>
<summary>Click to show • hide</summary>

- [X] Refactor [generator](generator)
- [X] Add constants where possible
- [X] Review generated code & comments
- [X] Unit testing of:
    - [X] Core functionality
    - [X] Helper methods
    - [X] Methods
    - [X] Types
- [X] Utility methods
- [ ] Add more examples
- [ ] Create Wiki page
- [ ] Publish stable version
- [ ] Add library to official [Telegram examples](https://core.telegram.org/bots/samples#go)

</details>

## Getting Started

How to get the library:

```shell
go get -u github.com/mymmrac/telego@latest
```

Make sure you get the latest version to have all new features & fixes.

More examples can be seen here:

- [Configuration](examples/configuration/main.go)
- [Sending files (documents, photos, media groups)](examples/sending_fiels/main.go)
- [Inline keyboard](examples/inline_keyboard/main.go)
- [Keyboard](examples/keyboard/main.go)

> Note: Error handling may be missing in examples, but I strongly recommend handling all errors.

> Note: While library in unstable version (v0.x.x) some parts of examples may not work.

### Basic setup

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
	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		return
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

### Getting updates

In order to receive updates you can use two methods:

- using long polling (`bot.GetUpdatesViaLongPulling`)
- using webhook (`bot.GetUpdatesViaWebhook`)

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

	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Optional. Set interval of getting updates (default: 0.5s).
	// If you want to get updates as fast as possible set to 0,
	// but webhook method is recommended for this.
	bot.SetUpdateInterval(time.Second / 2)

	// Get updates channel
	updates, _ := bot.GetUpdatesViaLongPulling(nil)

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

	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		return
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
	updates, _ := bot.GetUpdatesViaWebhook("/bot" + bot.Token())

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

### Using Telegram methods

All Telegram Bot API methods described in [documentation](https://core.telegram.org/bots/api#available-methods) can be
used by the library. They have same names and same parameters, parameters represented by struct with
name: `<methodName>` + `Params`. If method don't have required parameters `nil` value can be used as a parameter.

> Note: [`types.go`](types.go) and [`methods.go`](methods.go) was automatically [generated](generator) from [documentation][TelegramBotAPI], and it's possible that they have errors or missing parts both in comments and actual code.
> Fell free to report such things.

```go
package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Call method getMe
	botUser, _ := bot.GetMe()
	fmt.Printf("Bot User: %#v\n", botUser)

	updates, _ := bot.GetUpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()

	for update := range updates {
		if update.Message != nil {
			// Retrieve chat ID
			chatID := update.Message.Chat.ID

			// Call method sendMessage (https://core.telegram.org/bots/api#sendmessage).
			// Sends message to sender with same text (echo bot).
			sentMessage, _ := bot.SendMessage(&telego.SendMessageParams{
				ChatID: telego.ChatID{ID: chatID},
				Text:   update.Message.Text,
			})

			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}
}
```

## Contribution

1. Fork repo
2. Clone `git clone https://github.com/mymmrac/telego.git`
3. Create new branch `git checkout -b my-new-feature`
4. Make your changes, then add them `git add .`
5. Commit `git commit -m "New feature added"`
6. Push `git push origin my-new-feature`
7. Create pull request

> Note: Please try to use descriptive names for your changes, not just `fix` or `new stuff`.

## License

Telego is distributed under [MIT licence](LICENSE).

[TelegramBotAPI]: https://core.telegram.org/bots/api

[TelegramVersionBadge]: https://img.shields.io/static/v1?label=Supported%20Telegram%20Bot%20API&color=29a1d4&logo=telegram&message=v5.4

[TelegramLastVersion]: https://core.telegram.org/bots/api#november-5-2021
