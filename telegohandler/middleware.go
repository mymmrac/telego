package telegohandler

import (
	"context"
	"time"

	"github.com/mymmrac/telego"
)

// PanicRecovery returns a middleware that will recover handler from panic
// Note: It's not recommend to ignore panics, use [PanicRecoveryHandler] instead to handle them
func PanicRecovery() Middleware {
	return PanicRecoveryHandler(nil)
}

// PanicRecoveryHandler returns a middleware that will recover handler from panic and call panic handler
func PanicRecoveryHandler(panicHandler func(recovered any)) Middleware {
	return func(bot *telego.Bot, update telego.Update, next Handler) {
		defer func() {
			if recovered := recover(); recovered != nil && panicHandler != nil {
				panicHandler(recovered)
			}
		}()
		next(bot, update)
	}
}

// Timeout returns a middleware that will add timeout to context
func Timeout(timeout time.Duration) Middleware {
	return func(bot *telego.Bot, update telego.Update, next Handler) {
		ctx, cancel := context.WithTimeout(update.Context(), timeout)
		next(bot, update.WithContext(ctx))
		cancel()
	}
}
