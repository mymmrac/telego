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
	run := false
	ctx := &Context{
		ctx: context.Background(),
		ctxBase: &ctxBase{
			group: &HandlerGroup{
				routes: []route{
					{
						handler: func(ctx *Context, update telego.Update) error {
							_, hasDeadline := ctx.Deadline()
							assert.True(t, hasDeadline)
							run = true
							return nil
						},
					},
				},
			},
			stack: []int{-1},
		},
	}

	err := Timeout(time.Minute)(ctx, telego.Update{})
	require.NoError(t, err)
	assert.True(t, run)
}
