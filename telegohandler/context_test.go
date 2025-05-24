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
		ctx: t.Context(),
	}
	_, ok := ctx.Deadline()
	assert.False(t, ok)
}

func TestContext_Done(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
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
		ctx: t.Context(),
	}
	assert.NoError(t, ctx.Err())
}

func TestContext_Value(t *testing.T) {
	ctx := &Context{
		ctx: context.WithValue(t.Context(), "key", "value"), //nolint:staticcheck
	}
	assert.Equal(t, "value", ctx.Value("key"))
}

func TestContext_Context(t *testing.T) {
	ctx := t.Context()
	bCtx := &Context{
		ctx: ctx,
	}

	assert.Equal(t, ctx, bCtx.Context())
}

func TestContext_WithContext(t *testing.T) {
	bCtx := &Context{}
	ctx := t.Context()

	t.Run("success", func(t *testing.T) {
		newCtx := bCtx.WithContext(ctx)
		assert.Equal(t, ctx, newCtx.ctx)
	})

	t.Run("nil", func(t *testing.T) {
		assert.Panics(t, func() {
			bCtx.WithContext(nil) //nolint:staticcheck
		})
	})

	t.Run("recursion", func(t *testing.T) {
		bCtx.ctx = ctx
		bCtx = bCtx.WithValue("key", "value")
		assert.Equal(t, "value", bCtx.Value("key"))
		newCtx := bCtx.WithContext(context.WithoutCancel(bCtx)) //nolint:contextcheck
		assert.Equal(t, "value", newCtx.Value("key"))
	})
}

func TestContext_WithValue(t *testing.T) {
	ctx := &Context{
		ctx: t.Context(),
	}
	newCtx := ctx.WithValue("key", "value")
	assert.Equal(t, "value", newCtx.Value("key"))
}

func TestContext_WithTimeout(t *testing.T) {
	ctx := &Context{
		ctx: t.Context(),
	}
	ctx, cancel := ctx.WithTimeout(time.Minute)
	assert.NotNil(t, cancel)
	_, ok := ctx.Deadline()
	assert.True(t, ok)
}

func TestContext_WithCancel(t *testing.T) {
	ctx := &Context{
		ctx: t.Context(),
	}
	ctx, cancel := ctx.WithCancel()
	assert.NotNil(t, cancel)
	cancel()
	assert.ErrorIs(t, ctx.Err(), context.Canceled)
}

func TestContext_WithoutCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), time.Hour)
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
		ctxBase: &ctxBase{
			bot: &telego.Bot{},
		},
	}
	assert.NotNil(t, ctx.Bot())
}

func TestContext_UpdateID(t *testing.T) {
	ctx := &Context{
		ctxBase: &ctxBase{
			updateID: 1,
		},
	}
	assert.Equal(t, 1, ctx.UpdateID())
}

func TestContext_Next(t *testing.T) {
	run := false

	gr := &HandlerGroup{}

	gr.Use(func(ctx *Context, update telego.Update) error {
		update.UpdateID = 1
		ctx = ctx.WithContext(context.WithValue(ctx, "key", "value")) //nolint:staticcheck
		return ctx.Next(update)
	})

	gr1 := gr.Group()
	gr1.Handle(func(ctx *Context, update telego.Update) error {
		t.Fatalf("Should not be called")
		return nil
	}, None())

	gr2 := gr.Group()
	gr2.Handle(func(ctx *Context, update telego.Update) error {
		assert.Equal(t, 1, update.UpdateID)
		assert.Equal(t, "value", ctx.Value("key"))
		run = true
		return nil
	})

	ctx := &Context{
		ctxBase: &ctxBase{
			group: gr,
			stack: []int{-1},
		},
	}

	err := ctx.Next(telego.Update{})
	require.NoError(t, err)
	assert.True(t, run)
}
