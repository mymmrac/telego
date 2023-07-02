package telegohandler

import (
	"context"
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
		assert.NotNil(t, gr.handlers[0].handler)
		assert.Nil(t, gr.handlers[0].predicates)

		gr.handlers = make([]conditionalHandler, 0)
	})

	predicate := Predicate(func(update telego.Update) bool { return false })

	t.Run("with_predicates", func(t *testing.T) {
		gr.Handle(handler, predicate)

		require.Equal(t, 1, len(gr.handlers))
		assert.NotNil(t, gr.handlers[0].handler)
		assert.NotNil(t, gr.handlers[0].predicates)

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

	middleware := Middleware(func(bot *telego.Bot, update telego.Update, next Handler) {
		next(bot, update)
	})

	t.Run("success", func(t *testing.T) {
		gr.Use(middleware)

		require.Equal(t, 1, len(gr.middlewares))
		assert.NotNil(t, gr.middlewares[0])
	})
}

//nolint:funlen
func TestHandlerGroup_processUpdate(t *testing.T) {
	var order []int
	lock := sync.Mutex{}
	updOrder := func(i int) {
		lock.Lock()
		order = append(order, i)
		lock.Unlock()
	}

	gr := &HandlerGroup{
		predicates: []Predicate{
			func(update telego.Update) bool {
				t.Log("Predicate")
				updOrder(1)
				return true
			},
		},
		middlewares: []Middleware{
			func(bot *telego.Bot, update telego.Update, next Handler) {
				t.Log("Before next")
				updOrder(9)
				next(bot, update)
				t.Log("After next")
				updOrder(10)
			},
			func(bot *telego.Bot, update telego.Update, next Handler) {
				t.Log("Before nested next")
				updOrder(11)
				next(bot, update)
				t.Log("After nested next")
				updOrder(12)
			},
		},
		groups: []*HandlerGroup{
			{
				handlers: []conditionalHandler{
					{
						predicates: []Predicate{
							func(update telego.Update) bool {
								t.Log("Predicate handler nested in a group")
								updOrder(14)
								return false
							},
						},
					},
				},
			},
			{
				middlewares: []Middleware{
					func(bot *telego.Bot, update telego.Update, next Handler) {
						t.Log("Before nested in a group next")
						updOrder(15)
						ctx, cancel := context.WithCancel(update.Context())
						cancel()
						next(bot, update.WithContext(ctx))
						updOrder(19)
					},
				},
				groups: []*HandlerGroup{
					{
						middlewares: []Middleware{
							func(bot *telego.Bot, update telego.Update, next Handler) {
								assert.Fail(t, "shouldn't be called")
							},
						},
					},
				},
			},
			{
				middlewares: []Middleware{
					func(bot *telego.Bot, update telego.Update, next Handler) {
						t.Log("Before nested in a group next")
						updOrder(18)
						ctx, cancel := context.WithTimeout(update.Context(), smallTimeout)
						next(bot, update.WithContext(ctx))
						cancel()
						t.Log("After nested in a group next")
						// Do not record this, since it can be executed in any time and this is expected
						// order = append(order, 17)
					},
					func(bot *telego.Bot, update telego.Update, next Handler) {
						t.Log("Before nested in a group next")
						updOrder(16)
					},
				},
			},
			{
				predicates: []Predicate{
					func(update telego.Update) bool {
						t.Log("Predicate nested in a group")
						updOrder(13)
						return false
					},
				},
			},
			{
				predicates: []Predicate{
					func(update telego.Update) bool {
						t.Log("Predicate nested in a group")
						updOrder(2)
						return true
					},
				},
				middlewares: []Middleware{
					func(bot *telego.Bot, update telego.Update, next Handler) {
						t.Log("Before nested in a group next")
						updOrder(5)
						next(bot, update)
						t.Log("After nested in a group next")
						updOrder(6)
					},
				},
				handlers: []conditionalHandler{
					{
						handler: func(bot *telego.Bot, update telego.Update) {
							t.Log("Handler in a group")
							updOrder(3)
						},
						predicates: []Predicate{
							func(update telego.Update) bool {
								t.Log("Predicate handler nested in a group")
								updOrder(4)
								return true
							},
						},
					},
				},
			},
		},
		handlers: []conditionalHandler{
			{
				handler: func(bot *telego.Bot, update telego.Update) {
					t.Log("Handler")
					updOrder(7)
				},
				predicates: []Predicate{
					func(update telego.Update) bool {
						t.Log("Predicate handler")
						updOrder(8)
						return true
					},
				},
			},
		},
	}

	gr.processUpdate(nil, telego.Update{})

	lock.Lock()
	t.Log("Order:", order)
	assert.Equal(t, []int{1, 9, 11, 14, 15, 19, 18, 16, 13, 2, 5, 4, 3, 6, 12, 10}, order)
	lock.Unlock()
}
