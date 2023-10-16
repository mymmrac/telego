package telegohandler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

func TestPanicRecovery(t *testing.T) {
	bot, err := telego.NewBot(token, telego.WithDiscardLogger())
	require.NoError(t, err)

	t.Run("no_panic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			PanicRecovery()(bot, telego.Update{}, func(bot *telego.Bot, update telego.Update) {})
		})
	})

	t.Run("panic_recovered", func(t *testing.T) {
		const panicValue = "test panic"
		assert.NotPanics(t, func() {
			PanicRecoveryHandler(func(recovered any) {
				assert.Equal(t, panicValue, recovered)
			})(bot, telego.Update{}, func(bot *telego.Bot, update telego.Update) {
				panic(panicValue)
			})
		})
	})
}

func TestTimeout(t *testing.T) {
	bot, err := telego.NewBot(token, telego.WithDiscardLogger())
	require.NoError(t, err)

	hasDeadline := false
	Timeout(time.Minute)(bot, telego.Update{}, func(bot *telego.Bot, update telego.Update) {
		_, hasDeadline = update.Context().Deadline()
	})
	assert.True(t, hasDeadline)
}
