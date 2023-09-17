package telego

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/valyala/fasthttp"

	ta "github.com/mymmrac/telego/telegoapi"
)

// WithAPICaller sets custom API caller to use
func WithAPICaller(caller ta.Caller) BotOption {
	return func(bot *Bot) error {
		bot.api = caller
		return nil
	}
}

// WithFastHTTPClient sets fasthttp client to use
func WithFastHTTPClient(client *fasthttp.Client) BotOption {
	return func(bot *Bot) error {
		bot.api = ta.FastHTTPCaller{Client: client}
		return nil
	}
}

// WithHTTPClient sets http client to use
func WithHTTPClient(client *http.Client) BotOption {
	return func(bot *Bot) error {
		bot.api = ta.HTTPCaller{Client: client}
		return nil
	}
}

// WithRequestConstructor sets custom request constructor to use
func WithRequestConstructor(constructor ta.RequestConstructor) BotOption {
	return func(bot *Bot) error {
		bot.constructor = constructor
		return nil
	}
}

// WithDefaultLogger configures default logger. Redefines existing logger.
// Note: Default logger will hide your bot token, but it still may log sensitive information, it's only safe to use
// default logger in testing environment.
func WithDefaultLogger(debugMode, printErrors bool) BotOption {
	return func(bot *Bot) error {
		log := &logger{
			Out:         os.Stderr,
			DebugMode:   debugMode,
			PrintErrors: printErrors,
			Replacer:    defaultReplacer(bot.Token()),
		}
		bot.log = log
		return nil
	}
}

// WithExtendedDefaultLogger configures default logger, replacer can be nil. Redefines existing loggers.
// Note: Keep in mind that debug logs will include your bot token. It's only safe to have them enabled in
// testing environment, or hide sensitive information (like bot token) yourself.
func WithExtendedDefaultLogger(debugMode, printErrors bool, replacer *strings.Replacer) BotOption {
	return func(bot *Bot) error {
		log := &logger{
			Out:         os.Stderr,
			DebugMode:   debugMode,
			PrintErrors: printErrors,
			Replacer:    replacer,
		}
		bot.log = log
		return nil
	}
}

// WithDefaultDebugLogger configures default debug logger. Alias to default logger with debug and error logs.
// Redefines existing logger.
func WithDefaultDebugLogger() BotOption {
	return WithDefaultLogger(true, true)
}

// WithDiscardLogger configures discard logger. Alias to default logger with disabled logs. Redefines existing logger.
func WithDiscardLogger() BotOption {
	return WithDefaultLogger(false, false)
}

// WithLogger sets logger to use. Redefines existing loggers.
// Note: Keep in mind that debug logs will include your bot token. It's only safe to have them enabled in
// testing environment, or hide sensitive information (like bot token) yourself.
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

// WithTestServerPath use the test server API path instead of regular API path
//
// Regular API: https://<api-server>/bot<token>/<method-name>
// Test API:    https://<api-server>/bot<token>/test/<method-name>
func WithTestServerPath() BotOption {
	return func(bot *Bot) error {
		bot.useTestServerPath = true
		return nil
	}
}

// WithHealthCheck enables health check using [Bot.GetMe] method on start
func WithHealthCheck() BotOption {
	return func(bot *Bot) error {
		bot.healthCheckRequested = true
		return nil
	}
}

// WithWarnings treat Telegram warnings as an error
// Note: Any request that has a non-empty error will return both result and error
func WithWarnings() BotOption {
	return func(bot *Bot) error {
		bot.reportWarningAsErrors = true
		return nil
	}
}
