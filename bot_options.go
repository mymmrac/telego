package telego

import (
	"errors"
	"os"

	"github.com/valyala/fasthttp"
)

// FastHTTPClient - Sets fasthttp client to use. Redefines existing API caller
func FastHTTPClient(client *fasthttp.Client) BotOption {
	return func(bot *Bot) error {
		bot.api = fasthttpAPICaller{Client: client}
		return nil
	}
}

// DefaultLogger - Configure default logger. Redefines existing logger
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

// SetLogger - Set logger. Redefines existing logger
func SetLogger(log Logger) BotOption {
	return func(bot *Bot) error {
		bot.log = log
		return nil
	}
}

// SetAPIServer - Sets bot API server URL
func SetAPIServer(apiURL string) BotOption {
	return func(bot *Bot) error {
		if apiURL == "" {
			return errors.New("empty bot api server url")
		}

		bot.apiURL = apiURL
		return nil
	}
}
