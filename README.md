# Go Telegram Bot API

[![Go Reference](https://pkg.go.dev/badge/github.com/mymmrac/go-telegram-bot-api#section-readme.svg)](https://pkg.go.dev/github.com/mymmrac/go-telegram-bot-api#section-readme)
[![CI Status](https://github.com/mymmrac/go-telegram-bot-api/actions/workflows/ci.yml/badge.svg)](https://github.com/mymmrac/go-telegram-bot-api/actions/workflows/ci.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_go-telegram-bot-api&metric=alert_status)](https://sonarcloud.io/dashboard?id=mymmrac_go-telegram-bot-api)
[![Telegram Bot API Version](https://img.shields.io/static/v1?label=Supported%20Telegram%20Bot%20API&message=v5.3&color=29a1d4&logo=telegram)](https://core.telegram.org/bots/api#june-25-2021)

[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_go-telegram-bot-api&metric=bugs)](https://sonarcloud.io/dashboard?id=mymmrac_go-telegram-bot-api)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_go-telegram-bot-api&metric=code_smells)](https://sonarcloud.io/dashboard?id=mymmrac_go-telegram-bot-api)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=mymmrac_go-telegram-bot-api&metric=ncloc)](https://sonarcloud.io/dashboard?id=mymmrac_go-telegram-bot-api)

Telegram Bot API library for Golang with full [API](https://core.telegram.org/bots/api) implementation (one-to-one)

The goal of this library was to create API with same types and methods as actual telegram bot API. Every type and method
have been represented in [`types.go`](https://github.com/mymmrac/go-telegram-bot-api/blob/main/types.go)
and [`methods.go`](https://github.com/mymmrac/go-telegram-bot-api/blob/main/methods.go) files with mostly all
documentation from telegram.

> Note: `types.go` and `methods.go` was automatically [generated](https://github.com/mymmrac/go-telegram-bot-api/tree/main/generator) from [documentation](https://core.telegram.org/bots/api), and it's possible that they have errors or missing parts both in comments and actual code.
> Fell free to report such things.

> Note: While library in unstable version (v0.x.x) some parts of examples may work only in the latest commit.

## Examples

How to get the library: `go get -u github.com/mymmrac/go-telegram-bot-api`

> Note: All methods that have `(default: ...)` in comment isn't required for working bot, they were used just to show available configuration options.

> Note: Error handling may be missing in examples, but I strongly recommend to handle all errors.

More examples can be seen here:
- [Sending files (documents, photos, media groups)](https://github.com/mymmrac/go-telegram-bot-api/blob/main/examples/sending_fiels/main.go)
- [Inline keyboard](https://github.com/mymmrac/go-telegram-bot-api/blob/main/examples/inline_keyboard/main.go)
- [Keyboard](https://github.com/mymmrac/go-telegram-bot-api/blob/main/examples/keyboard/main.go)

### Basic setup

For start, you need to create instance of your bot and specify [token](https://core.telegram.org/bots/api#authorizing-your-bot).

```go
package main

import (
	"fmt"
	"net/http"
	"os"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Create bot
	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Change bot token (default: set by telego.NewBot(...))
	_ = bot.SetToken("new bot token")

	// Change bot API server URL (default: https://api.telegram.org)
	_ = bot.SetAPIServer("new bot api server")

	// Change http client (default: http.DefaultClient)
	_ = bot.SetClient(http.DefaultClient)

	// Settings of default logger, enable printing debug information and errors (default: false, true)
	bot.DefaultLogger(true, true)

	var myLogger telego.Logger
	// Create you custom logger that implements telego.Logger (default: telego has build in default logger)
	bot.SetLogger(myLogger)

	// Call method getMe
	botUser, _ := bot.GetMe()
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

	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.DefaultLogger(true, true)

	// Set interval of getting updates (default: 0.5s)
	bot.SetUpdateInterval(time.Second / 2)

	// Get updates channel
	updates, _ := bot.GetUpdatesChan(&telego.GetUpdatesParams{})

	// Stop reviving updates from updates channel
	defer bot.StopGettingUpdates()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Println("====")
		fmt.Printf("%#v\n", update)
		fmt.Println("====")
	}
}
```

Webhook example: 

> Note: You may wish to use [Let's Encrypt](https://letsencrypt.org/) in order to generate your free TLS certificate.

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

	// Setup a webhook
	_ = bot.SetWebhook(&telego.SetWebhookParams{
		URL:         "https://www.google.com:443/" + bot.Token(),
		Certificate: &telego.InputFile{File: mustOpen("cert.pem")},
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo()
	fmt.Println()
	fmt.Printf("%#v\n", info)
	fmt.Println()

	// Get updates channel from webhook
	updates, _ := bot.ListenForWebhook("/" + bot.Token())

	// Start server for receiving requests from telegram
	bot.StartListeningForWebhook("0.0.0.0:443/"+bot.Token(), "cert.pem", "key.pem")

	// Loop through all updates when they came
	for update := range updates {
		fmt.Println("====")
		fmt.Printf("%#v\n", update)
		fmt.Println("====")
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

### Using Telegram methods

All Telegram Bot API methods described in [documentation](https://core.telegram.org/bots/api#available-methods) can be used by this library.
They have same names and same parameters, parameters represented by struct with name: `<methodName>` + `Params`. 
If method don't have required parameters `nil` value can be used as a parameter.

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
	fmt.Println()
	fmt.Printf("%#v\n", botUser)
	fmt.Println()

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

			fmt.Println()
			fmt.Printf("%v\n", sentMessage)
			fmt.Println()
		}
	}
}
```
