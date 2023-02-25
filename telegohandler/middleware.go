package telegohandler

import (
	"github.com/mymmrac/telego"
)

// PanicRecovery is a middleware that will recover handler form panic
func PanicRecovery(next Handler) Handler {
	return func(bot *telego.Bot, update telego.Update) {
		defer func() {
			err := recover()
			if err != nil {
				bot.Logger().Errorf("Panic recovered: %v", err)
			}
		}()

		next(bot, update)
	}
}
