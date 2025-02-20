package telegohandler

import (
	"time"

	"github.com/mymmrac/telego"
)

// PanicRecovery returns a middleware that will recover handler from panic
// Note: It's not recommend to ignore panics, use [PanicRecoveryHandler] instead to handle them
func PanicRecovery() Handler {
	return PanicRecoveryHandler(nil)
}

// PanicRecoveryHandler returns a middleware that will recover handler from panic and call panic handler.
// Error returned from panic handler will be returned as handler error, if panic handler is nil, panic will be ignored
// (not recommended, try to always handle panics).
func PanicRecoveryHandler(panicHandler func(recovered any) error) Handler {
	return func(ctx *Context, update telego.Update) (err error) {
		defer func() {
			if recovered := recover(); recovered != nil && panicHandler != nil {
				err = panicHandler(recovered)
			}
		}()
		return ctx.Next(update)
	}
}

// Timeout returns a middleware that will add timeout to context
func Timeout(timeout time.Duration) Handler {
	return func(ctx *Context, update telego.Update) error {
		ctx, cancel := ctx.WithTimeout(timeout)
		defer cancel()
		return ctx.Next(update)
	}
}
