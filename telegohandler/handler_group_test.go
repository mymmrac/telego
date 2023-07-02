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

	middleware := Middleware(func(bot *telego.Bot, update telego.Update, next Handler) {
		next(bot, update)
	})

	t.Run("success", func(t *testing.T) {
		gr.Use(middleware)

		require.Equal(t, 1, len(gr.middlewares))
		assert.NotNil(t, gr.middlewares[0])
	})
}

func TestHandlerGroup_processUpdate(t *testing.T) {
	var order []int
	gr := &HandlerGroup{
		predicates: []Predicate{
			func(update telego.Update) bool {
				t.Log("Predicate")
				order = append(order, 1)
				return true
			},
		},
		middlewares: []Middleware{
			func(bot *telego.Bot, update telego.Update, next Handler) {
				t.Log("Before next")
				order = append(order, 9)
				next(bot, update)
				t.Log("After next")
				order = append(order, 10)
			},
			func(bot *telego.Bot, update telego.Update, next Handler) {
				t.Log("Before nested next")
				order = append(order, 11)
				next(bot, update)
				t.Log("After nested next")
				order = append(order, 12)
			},
		},
		groups: []*HandlerGroup{
			{
				handlers: []conditionalHandler{
					{
						Predicates: []Predicate{
							func(update telego.Update) bool {
								t.Log("Predicate handler nested in a group")
								order = append(order, 14)
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
						order = append(order, 15)
						ctx, cancel := context.WithCancel(update.Context())
						cancel()
						next(bot, update.WithContext(ctx))
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
						order = append(order, 18)
						ctx, cancel := context.WithTimeout(update.Context(), smallTimeout)
						next(bot, update.WithContext(ctx))
						cancel()
						t.Log("After nested in a group next")
						order = append(order, 17)
					},
					func(bot *telego.Bot, update telego.Update, next Handler) {
						t.Log("Before nested in a group next")
						order = append(order, 16)
					},
				},
			},
			{
				predicates: []Predicate{
					func(update telego.Update) bool {
						t.Log("Predicate nested in a group")
						order = append(order, 13)
						return false
					},
				},
			},
			{
				predicates: []Predicate{
					func(update telego.Update) bool {
						t.Log("Predicate nested in a group")
						order = append(order, 2)
						return true
					},
				},
				middlewares: []Middleware{
					func(bot *telego.Bot, update telego.Update, next Handler) {
						t.Log("Before nested in a group next")
						order = append(order, 5)
						next(bot, update)
						t.Log("After nested in a group next")
						order = append(order, 6)
					},
				},
				handlers: []conditionalHandler{
					{
						Handler: func(bot *telego.Bot, update telego.Update) {
							t.Log("Handler in a group")
							order = append(order, 3)
						},
						Predicates: []Predicate{
							func(update telego.Update) bool {
								t.Log("Predicate handler nested in a group")
								order = append(order, 4)
								return true
							},
						},
					},
				},
			},
		},
		handlers: []conditionalHandler{
			{
				Handler: func(bot *telego.Bot, update telego.Update) {
					t.Log("Handler")
					order = append(order, 7)
				},
				Predicates: []Predicate{
					func(update telego.Update) bool {
						t.Log("Predicate handler")
						order = append(order, 8)
						return true
					},
				},
			},
		},
	}

	gr.processUpdate(nil, telego.Update{})
	t.Log("Order:", order)
	assert.Equal(t, []int{1, 9, 11, 14, 15, 18, 16, 17, 13, 2, 5, 4, 3, 6, 12, 10}, order)
}
