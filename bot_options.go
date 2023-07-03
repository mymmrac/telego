package telego

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

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
		bot.warningAsErrors = true
		return nil
	}
}

// WithRetry configure the number of retry attempts and the delay time telego
// will apply when doing API requests.
// `maxAttempts` is the maximum number of attempts it will do and must be >= 1.
// `delayFactor` is the base of exponential time delay and must be >= 1.
// `startDelay` the initial delay between retries.
// `maxDelay` the maximum delay between retries.
func WithRetry(maxAttempts, delayFactor int, startDelay, maxDelay time.Duration) BotOption {
	return func(bot *Bot) error {
		if maxAttempts < 1 {
			return errors.New("`maxAttempts` must be >= 1")
		}
		if delayFactor < 1 {
			return errors.New("`delayFactor` must be >= 1")
		}
		bot.retryOptions = retryOptions{
			maxAttempts,
			delayFactor,
			startDelay,
			maxDelay,
		}
		return nil
	}
}
