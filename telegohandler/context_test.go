package telegohandler

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

func TestContext_Deadline(t *testing.T) {
	ctx := &Context{
		ctx: context.Background(),
	}
	_, ok := ctx.Deadline()
	assert.False(t, ok)
}

func TestContext_Done(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	bCtx := &Context{
		ctx: ctx,
	}
	cancel()

	select {
	case <-bCtx.Done():
	case <-time.After(timeout):
		t.Fatal("Timeout")
	}
}

func TestContext_Err(t *testing.T) {
	ctx := &Context{
		ctx: context.Background(),
	}
	assert.NoError(t, ctx.Err())
}

func TestContext_Value(t *testing.T) {
	ctx := &Context{
		ctx: context.WithValue(context.Background(), "key", "value"), //nolint:staticcheck
	}
	assert.Equal(t, "value", ctx.Value("key"))
}

func TestContext_WithContext(t *testing.T) {
	bCtx := &Context{}
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		bCtx.WithContext(ctx)
		assert.Equal(t, ctx, bCtx.ctx)
	})

	t.Run("nil", func(t *testing.T) {
		assert.Panics(t, func() {
			bCtx.WithContext(nil) //nolint:staticcheck
		})
	})
}

func TestContext_WithValue(t *testing.T) {
	ctx := &Context{
		ctx: context.Background(),
	}
	ctx.WithValue("key", "value")
	assert.Equal(t, "value", ctx.Value("key"))
}

func TestContext_WithTimeout(t *testing.T) {
	ctx := &Context{
		ctx: context.Background(),
	}
	ctx, cancel := ctx.WithTimeout(time.Minute)
	assert.NotNil(t, cancel)
	_, ok := ctx.Deadline()
	assert.True(t, ok)
}

func TestContext_WithoutCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	bCtx := &Context{
		ctx: ctx,
	}

	bCtx = bCtx.WithoutCancel()
	assert.NotNil(t, cancel)
	_, ok := bCtx.Deadline()
	assert.False(t, ok)
}

func TestContext_Bot(t *testing.T) {
	ctx := &Context{
		bot: &telego.Bot{},
	}
	assert.NotNil(t, ctx.Bot())
}

func TestContext_UpdateID(t *testing.T) {
	ctx := &Context{
		updateID: 1,
	}
	assert.Equal(t, 1, ctx.UpdateID())
}

func TestContext_Next(t *testing.T) {
	run := false
	ctx := &Context{
		group: &HandlerGroup{
			routes: []route{
				{
					handler: func(ctx *Context, update telego.Update) error {
						update.UpdateID = 1
						return ctx.Next(update)
					},
				},
				{
					predicates: []Predicate{
						func(_ context.Context, _ telego.Update) bool {
							return true
						},
					},
					group: &HandlerGroup{
						routes: []route{
							{
								handler: func(_ *Context, update telego.Update) error {
									assert.Equal(t, 1, update.UpdateID)
									run = true
									return nil
								},
							},
						},
					},
				},
			},
		},
		stack: []int{-1},
	}

	err := ctx.Next(telego.Update{})
	require.NoError(t, err)
	assert.True(t, run)
}
