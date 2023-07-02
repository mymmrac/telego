package telegohandler

import (
	"github.com/mymmrac/telego"
)

// PanicRecovery is a middleware that will recover handler from panic
func PanicRecovery(bot *telego.Bot, update telego.Update, next Handler) {
	defer func() {
		if err := recover(); err != nil {
			bot.Logger().Errorf("Panic recovered: %v", err)
		}
	}()

	next(bot, update)
}
