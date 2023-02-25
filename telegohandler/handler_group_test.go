package telegohandler

import (
	"sync"
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

	handler := Handler(func(bot *telego.Bot, update telego.Update) {})

	t.Run("panic_nil_predicate", func(t *testing.T) {
		assert.Panics(t, func() {
			gr.Handle(handler, nil)
		})
	})

	t.Run("without_predicates", func(t *testing.T) {
		gr.Handle(handler)

		require.Equal(t, 1, len(gr.handlers))
		assert.NotNil(t, gr.handlers[0].Handler)
		assert.Nil(t, gr.handlers[0].Predicates)

		gr.handlers = make([]conditionalHandler, 0)
	})

	predicate := Predicate(func(update telego.Update) bool { return false })

	t.Run("with_predicates", func(t *testing.T) {
		gr.Handle(handler, predicate)

		require.Equal(t, 1, len(gr.handlers))
		assert.NotNil(t, gr.handlers[0].Handler)
		assert.NotNil(t, gr.handlers[0].Predicates)

		gr.handlers = make([]conditionalHandler, 0)
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

		require.Equal(t, 1, len(gr.groups))
		assert.Equal(t, newGr, gr.groups[0])

		gr.groups = nil
	})

	predicate := Predicate(func(update telego.Update) bool { return false })

	t.Run("with_predicates", func(t *testing.T) {
		newGr := gr.Group(predicate)

		require.Equal(t, 1, len(gr.groups))
		assert.Equal(t, newGr, gr.groups[0])
		assert.NotEmpty(t, gr.groups[0].predicates)

		gr.groups = nil
	})
}

func TestHandlerGroup_Use(t *testing.T) {
	gr := &HandlerGroup{}

	t.Run("panic_nil_middleware", func(t *testing.T) {
		assert.Panics(t, func() {
			gr.Use(nil)
		})
	})

	middleware := Middleware(func(next Handler) Handler {
		return func(bot *telego.Bot, update telego.Update) {
			next(bot, update)
		}
	})

	t.Run("success", func(t *testing.T) {
		gr.Use(middleware)

		require.Equal(t, 1, len(gr.middlewares))
		assert.NotNil(t, gr.middlewares[0])
	})
}

func TestHandlerGroup_useGroups(t *testing.T) {
	gr := &HandlerGroup{}

	middlewareOk1 := false
	middlewareOk2 := false
	gr.Use(func(next Handler) Handler {
		middlewareOk1 = true
		return func(bot *telego.Bot, update telego.Update) {
			middlewareOk2 = true
			next(bot, update)
		}
	})

	notUseOk := true
	notUseGr := gr.Group(func(update telego.Update) bool {
		return false
	})
	notUseGr.Handle(func(bot *telego.Bot, update telego.Update) {
		notUseOk = false
	})

	predicateOk := false
	newGr := gr.Group(func(update telego.Update) bool {
		predicateOk = true
		return true
	})

	handlerOk := false
	newGr.Handle(func(bot *telego.Bot, update telego.Update) {
		handlerOk = true
	})

	wg := &sync.WaitGroup{}
	ok := gr.useGroups(nil, telego.Update{}, wg)

	assert.True(t, ok)
	wg.Wait()

	assert.True(t, handlerOk)
	assert.True(t, predicateOk)
	assert.True(t, middlewareOk1)
	assert.True(t, middlewareOk2)
	assert.True(t, notUseOk)

	ok = gr.Group(func(update telego.Update) bool {
		return false
	}).useGroups(nil, telego.Update{}, wg)
	assert.False(t, ok)
}
