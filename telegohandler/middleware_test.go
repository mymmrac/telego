package telegohandler

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

func TestPanicRecovery(t *testing.T) {
	t.Run("no_panic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			err := PanicRecovery()(nil, telego.Update{})
			assert.NoError(t, err)
		})
	})

	t.Run("panic_recovered_no_error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			err := PanicRecoveryHandler(func(recovered any) error {
				assert.NotNil(t, recovered)
				return nil
			})(nil, telego.Update{})
			assert.NoError(t, err)
		})
	})

	t.Run("panic_recovered_error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			err := PanicRecoveryHandler(func(recovered any) error {
				assert.NotNil(t, recovered)
				return errTest
			})(nil, telego.Update{})
			assert.ErrorIs(t, err, errTest)
		})
	})
}

func TestTimeout(t *testing.T) {
	ctx := &Context{
		ctx: context.Background(),
		group: &HandlerGroup{
			middlewares: nil,
			routes:      nil,
		},
		middlewareIndex: -1,
	}

	err := Timeout(time.Minute)(ctx, telego.Update{})
	require.NoError(t, err)

	_, hasDeadline := ctx.Deadline()
	assert.True(t, hasDeadline)
}
