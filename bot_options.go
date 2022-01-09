package telego

import (
	"errors"
	"os"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/telegoapi"
)

// WithAPICaller sets custom API caller to use
func WithAPICaller(caller telegoapi.Caller) BotOption {
	return func(bot *Bot) error {
		bot.api = caller
		return nil
	}
}

// WithFastHTTPClient sets fasthttp client to use
func WithFastHTTPClient(client *fasthttp.Client) BotOption {
	return func(bot *Bot) error {
		bot.api = telegoapi.FasthttpAPICaller{Client: client}
		return nil
	}
}

// WithRequestConstructor sets custom request constructor to use
func WithRequestConstructor(constructor telegoapi.RequestConstructor) BotOption {
	return func(bot *Bot) error {
		bot.constructor = constructor
		return nil
	}
}

// WithDefaultLogger configures default logger. Redefines existing logger.
// Note: Keep in mind that debug logs will include your bot token, it's only safe to have them enabled in
// testing environment.
func WithDefaultLogger(debugMode, printErrors bool) BotOption {
	return func(bot *Bot) error {
		log := &logger{
			Out:         os.Stderr,
			DebugMode:   debugMode,
			PrintErrors: printErrors,
		}
		bot.log = log
		return nil
	}
}

// WithDiscardLogger configures discard logger. Alias to default logger with disabled logs. Redefines existing logger.
func WithDiscardLogger() BotOption {
	return WithDefaultLogger(false, false)
}

// WithLogger sets logger to use
// Note: Keep in mind that debug logs will include your bot token, it's only safe to have them enabled in
// testing environment.
func WithLogger(log Logger) BotOption {
	return func(bot *Bot) error {
		bot.log = log
		return nil
	}
}

// WithAPIServer sets bot API server URL to use
func WithAPIServer(apiURL string) BotOption {
	return func(bot *Bot) error {
		if apiURL == "" {
			return errors.New("empty bot api server url")
		}

		bot.apiURL = apiURL
		return nil
	}
}

// WithWebhookServer sets bot HTTP server used for listening for webhook
func WithWebhookServer(server *fasthttp.Server) BotOption {
	return func(bot *Bot) error {
		bot.server = server
		return nil
	}
}
