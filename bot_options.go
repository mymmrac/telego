package telego

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/goccy/go-json"
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
		bot.api = telegoapi.FastHTTPCaller{Client: client}
		return nil
	}
}

// WithHTTPClient sets http client to use
func WithHTTPClient(client *http.Client) BotOption {
	return func(bot *Bot) error {
		bot.api = telegoapi.HTTPCaller{Client: client}
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

// WithExtendedDefaultLogger configures default logger, replacer can be nil. Redefines existing logger.
// Note: Keep in mind that debug logs will include your bot token, it's only safe to have them enabled in
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

// WithLogger sets logger to use. Redefines existing logger.
// Note: Keep in mind that debug logs will include your bot token, it's only safe to have them enabled in
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
			return errors.New("telego: empty bot api server url")
		}

		bot.apiURL = apiURL
		return nil
	}
}

// WithHealthCheck enables health check using Bot.GetMe() method on start
func WithHealthCheck() BotOption {
	return func(bot *Bot) error {
		bot.healthCheckRequested = true
		return nil
	}
}

// WithWarnings treat Telegram warnings as errors
// Note: Any request that has non-empty error will return both result and error
func WithWarnings() BotOption {
	return func(bot *Bot) error {
		bot.warningAsErrors = true
		return nil
	}
}

// WithEmptyValues sets empty value to default one that will be erased from all requests
// Note: Used with Bot.EmptyValue() to get empty strings as parameters
func WithEmptyValues() BotOption {
	return func(bot *Bot) error {
		bot.replaceToEmpty = defaultBotEmptyValue
		return nil
	}
}

// WithCustomEmptyValues sets empty value to custom value that will be erased from all requests
// Note: Used with Bot.EmptyValue() to get empty strings as parameters values
// Warning: Request data is encoded using JSON, so the value will be escaped in JSON and may not match intended value
func WithCustomEmptyValues(emptyValue string) BotOption {
	return func(bot *Bot) error {
		if emptyValue == "" {
			return fmt.Errorf("empty value can't be zero length")
		}

		data, err := json.Marshal(emptyValue)
		if err != nil {
			return fmt.Errorf("marshal empty value: %w", err)
		}

		if fmt.Sprintf(`"%s"`, emptyValue) != string(data) {
			return fmt.Errorf(`empty value does't match it's JSON encoded varian: "%s" not equal to %s`,
				emptyValue, string(data))
		}

		bot.replaceToEmpty = emptyValue
		return nil
	}
}
