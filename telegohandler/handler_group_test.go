package telegohandler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

func TestHandlerGroup_Handle(t *testing.T) {
	gr := &HandlerGroup{}

	t.Run("panic_nil_handler", func(t *testing.T) {
		assert.Panics(t, func() {
			gr.Handle(nil)
		})
	})

	handler := Handler(func(_ *Context, _ telego.Update) error { return nil })

	t.Run("panic_nil_predicate", func(t *testing.T) {
		assert.Panics(t, func() {
			gr.Handle(handler, nil)
		})
	})

	t.Run("without_predicates", func(t *testing.T) {
		gr.Handle(handler)

		require.Len(t, gr.routes, 1)
		assert.NotNil(t, gr.routes[0].handler)
		assert.Nil(t, gr.routes[0].predicates)

		gr.routes = nil
	})

	predicate := Predicate(func(_ context.Context, _ telego.Update) bool { return false })

	t.Run("with_predicates", func(t *testing.T) {
		gr.Handle(handler, predicate)

		require.Len(t, gr.routes, 1)
		assert.NotNil(t, gr.routes[0].handler)
		assert.NotNil(t, gr.routes[0].predicates)

		gr.routes = nil
	})
}

func TestHandlerGroup_Group(t *testing.T) {
	gr := &HandlerGroup{}

	t.Run("panic_nil_predicate", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = gr.Group(nil)
		})
	})

	t.Run("without_predicates", func(t *testing.T) {
		newGr := gr.Group()

		require.Len(t, gr.routes, 1)
		assert.Equal(t, newGr, gr.routes[0].group)

		gr.routes = nil
	})

	predicate := Predicate(func(_ context.Context, _ telego.Update) bool { return false })

	t.Run("with_predicates", func(t *testing.T) {
		newGr := gr.Group(predicate)

		require.Len(t, gr.routes, 1)
		assert.Equal(t, newGr, gr.routes[0].group)
		assert.NotEmpty(t, gr.routes[0].predicates)

		gr.routes = nil
	})
}

func TestHandlerGroup_Use(t *testing.T) {
	gr := &HandlerGroup{}

	t.Run("panic_nil_middleware", func(t *testing.T) {
		assert.Panics(t, func() {
			gr.Use(nil)
		})
	})

	middleware := Handler(func(ctx *Context, update telego.Update) error {
		return ctx.Next(update)
	})

	t.Run("success", func(t *testing.T) {
		gr.Use(middleware)

		require.Len(t, gr.routes, 1)
		assert.NotNil(t, gr.routes[0].handler)
	})
}

func TestHandlerGroup_depth(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		gr := &HandlerGroup{}
		assert.Equal(t, 1, gr.depth(1))
	})

	t.Run("2", func(t *testing.T) {
		gr := &HandlerGroup{
			routes: []route{
				{},
				{group: &HandlerGroup{}},
			},
		}
		assert.Equal(t, 2, gr.depth(1))
	})

	t.Run("3", func(t *testing.T) {
		gr := &HandlerGroup{
			routes: []route{
				{},
				{group: &HandlerGroup{}},
				{},
				{group: &HandlerGroup{
					routes: []route{
						{},
						{group: &HandlerGroup{}},
					},
				}},
				{group: &HandlerGroup{}},
				{},
			},
		}
		assert.Equal(t, 3, gr.depth(1))
	})
}
