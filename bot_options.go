package telego

import (
	"errors"
	"os"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/go-telegram-bot-api/api"
)

// CustomAPICaller sets custom API caller to use
func CustomAPICaller(caller api.Caller) BotOption {
	return func(bot *Bot) error {
		bot.api = caller
		return nil
	}
}

// FastHTTPClient sets fasthttp client to use
func FastHTTPClient(client *fasthttp.Client) BotOption {
	return func(bot *Bot) error {
		bot.api = api.FasthttpAPICaller{Client: client}
		return nil
	}
}

// CustomRequestConstructor sets custom request constructor to use
func CustomRequestConstructor(constructor api.RequestConstructor) BotOption {
	return func(bot *Bot) error {
		bot.constructor = constructor
		return nil
	}
}

// DefaultLogger configures default logger. Redefines existing logger
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
