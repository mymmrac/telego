# Telego | Go Telegram Bot API

[![Go Reference](https://pkg.go.dev/badge/github.com/mymmrac/go-telegram-bot-api#section-readme.svg)](https://pkg.go.dev/github.com/mymmrac/go-telegram-bot-api#section-readme)
[![CI Status](https://github.com/mymmrac/go-telegram-bot-api/actions/workflows/ci.yml/badge.svg)](https://github.com/mymmrac/go-telegram-bot-api/actions/workflows/ci.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_go-telegram-bot-api&metric=alert_status)](https://sonarcloud.io/dashboard?id=mymmrac_go-telegram-bot-api)
[![Telegram Bot API Version][TelegramVersionBadge]][TelegramLastVersion]
[![Telegram Chat](https://img.shields.io/static/v1?label=Discussion&message=chat&color=29a1d4&logo=telegram)](https://t.me/telegoLibrary)

[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_go-telegram-bot-api&metric=bugs)](https://sonarcloud.io/dashboard?id=mymmrac_go-telegram-bot-api)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_go-telegram-bot-api&metric=code_smells)](https://sonarcloud.io/dashboard?id=mymmrac_go-telegram-bot-api)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_go-telegram-bot-api&metric=ncloc)](https://sonarcloud.io/dashboard?id=mymmrac_go-telegram-bot-api)

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
- [ ] Unit testing of:
    - [ ] Core functionality
    - [ ] Helper methods
    - [ ] Methods & types
- [ ] Add more examples
- [ ] Create Wiki page
- [ ] Publish stable version
- [ ] Add library to official [Telegram examples](https://core.telegram.org/bots/samples#go)

</details>

## Getting Started

How to get the library:

```shell
go get -u github.com/mymmrac/go-telegram-bot-api@latest
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

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	// Get Bot token from environment variables
	botToken := os.Getenv("TOKEN")

	// Create bot and enable debugging info (more on configuration at /examples/configuration/main.go)
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

- using long polling (`bot.GetUpdatesChan`)
- using webhook (`bot.ListenForWebhook`)

Let's start from long pulling (easier for local testing):

```go
package main

import (
	"fmt"
	"os"
	"time"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set interval of getting updates (default: 0.5s)
	// If you want to get updates as fast as possible set to 0
	bot.SetUpdateInterval(time.Second / 2)

	// Get updates channel
	updates, _ := bot.GetUpdatesChan(&telego.GetUpdatesParams{})

	// Stop reviving updates from updates channel
	defer bot.StopGettingUpdates()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %#v\n", update)
	}
}
```

Webhook example:

```go
package main

import (
	"fmt"
	"os"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set up a webhook
	_ = bot.SetWebhook(&telego.SetWebhookParams{
		URL:         "https://www.google.com:443/" + bot.Token(),
		Certificate: &telego.InputFile{File: mustOpen("cert.pem")},
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo()
	fmt.Printf("Webhook Info: %#v\n", info)

	// Start server for receiving requests from telegram
	bot.StartListeningForWebhookTLS("0.0.0.0:443/"+bot.Token(), "cert.pem", "key.pem")

	// Get updates channel from webhook. Note for one bot only one webhook allowed
	updates, _ := bot.ListenForWebhook("/" + bot.Token())

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %#v\n", update)
	}
}

// Helper function to open file or panic
func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
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

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.DefaultLogger(true, true)

	// Call method getMe
	botUser, _ := bot.GetMe()
	fmt.Printf("Bot User: %#v\n", botUser)

	updates, _ := bot.GetUpdatesChan(&telego.GetUpdatesParams{})
	defer bot.StopGettingUpdates()

	for update := range updates {
		if update.Message != nil {
			// Retrieve chat ID
			chatID := update.Message.Chat.ID

			// Call method sendMessage. Sends message to sender with same text (echo bot)
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
2. Clone `git clone https://github.com/mymmrac/go-telegram-bot-api.git`
3. Create new branch `git checkout -b my-new-feature`
4. Make your changes, then add them `git add .`
5. Commit `git commit -m "New feature added"`
6. Push `git push origin my-new-feature`
7. Create pull request

> Note: Please try to use descriptive names for your changes, not just `fix` or `new stuff`.

## License

Telego is distributed under [MIT licence](LICENSE).

[TelegramBotAPI]: https://core.telegram.org/bots/api

[TelegramVersionBadge]: https://img.shields.io/static/v1?label=Supported%20Telegram%20Bot%20API&color=29a1d4&logo=telegram&message=v5.3

[TelegramLastVersion]: https://core.telegram.org/bots/api#june-25-2021
