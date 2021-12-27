package telego

import (
	"errors"
	"os"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/telegoapi"
)

// CustomAPICaller sets custom API caller to use
func CustomAPICaller(caller telegoapi.Caller) BotOption {
	return func(bot *Bot) error {
		bot.api = caller
		return nil
	}
}

// FastHTTPClient sets fasthttp client to use
func FastHTTPClient(client *fasthttp.Client) BotOption {
	return func(bot *Bot) error {
		bot.api = telegoapi.FasthttpAPICaller{Client: client}
		return nil
	}
}

// CustomRequestConstructor sets custom request constructor to use
func CustomRequestConstructor(constructor telegoapi.RequestConstructor) BotOption {
	return func(bot *Bot) error {
		bot.constructor = constructor
		return nil
	}
}

// DefaultLogger configures default logger. Redefines existing logger.
// Note: Keep in mind that debug logs will include your bot token, it's only safe to have them enabled in
// testing environment.
func DefaultLogger(debugMode, printErrors bool) BotOption {
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

// SetLogger sets logger to use
// Note: Keep in mind that debug logs will include your bot token, it's only safe to have them enabled in
// testing environment.
func SetLogger(log Logger) BotOption {
	return func(bot *Bot) error {
		bot.log = log
		return nil
	}
}

// SetAPIServer sets bot API server URL to use
func SetAPIServer(apiURL string) BotOption {
	return func(bot *Bot) error {
		if apiURL == "" {
			return errors.New("empty bot api server url")
		}

		bot.apiURL = apiURL
		return nil
	}
}

// SetWebhookServer sets bot HTTP server used for listening for webhook
func SetWebhookServer(server *fasthttp.Server) BotOption {
	return func(bot *Bot) error {
		bot.server = server
		return nil
	}
}
