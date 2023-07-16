package telegohandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

var _ Middleware = PanicRecovery

func TestPanicRecovery(t *testing.T) {
	bot, err := telego.NewBot(token, telego.WithDiscardLogger())
	require.NoError(t, err)

	t.Run("no_panic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			PanicRecovery(bot, telego.Update{}, func(bot *telego.Bot, update telego.Update) {})
		})
	})

	t.Run("panic_recovered", func(t *testing.T) {
		assert.NotPanics(t, func() {
			PanicRecovery(bot, telego.Update{}, func(bot *telego.Bot, update telego.Update) {
				panic("test panic")
			})
		})
	})
}
