package telegohandler

import (
	"context"
	"sync"

	"github.com/mymmrac/telego"
)

// conditionalHandler represents handler with respectful predicates
type conditionalHandler struct {
	Handler    Handler
	Predicates []Predicate
}

// match matches the current update and handler
func (h conditionalHandler) match(update telego.Update) bool {
	for _, p := range h.Predicates {
		if !p(update.Clone()) {
			return false
		}
	}
	return true
}

// HandlerGroup represents a group of handlers, middlewares and child groups
type HandlerGroup struct {
	parent      *HandlerGroup
	predicates  []Predicate
	groups      []*HandlerGroup
	handlers    []conditionalHandler
	middlewares []Middleware
}

// applyMiddlewares applies all parent middlewares and its own in reverse order
func (h *HandlerGroup) applyMiddlewares(next Handler) Handler {
	for i := len(h.middlewares) - 1; i >= 0; i-- {
		next = h.middlewares[i](next)
	}

	if h.parent != nil {
		next = h.parent.applyMiddlewares(next)
	}

	return next
}

// useHandlers tries to match update to a handler
func (h *HandlerGroup) useHandlers(bot *telego.Bot, update telego.Update, wg *sync.WaitGroup) bool {
	for _, handler := range h.handlers {
		if !handler.match(update) {
			continue
		}

		wg.Add(1)
		go func(ch conditionalHandler) {
			ctx, cancel := context.WithCancel(update.Context())
			h.applyMiddlewares(ch.Handler)(bot, update.WithContext(ctx))
			wg.Done()
			cancel()
		}(handler)

		return true
	}

	return false
}

// match matches the current update and group
func (h *HandlerGroup) match(update telego.Update) bool {
	for _, p := range h.predicates {
		if !p(update.Clone()) {
			return false
		}
	}
	return true
}

// useHandlers tries to match update to a group
func (h *HandlerGroup) useGroups(bot *telego.Bot, update telego.Update, wg *sync.WaitGroup) bool {
	for _, group := range h.groups {
		if !group.match(update) {
			continue
		}

		if ok := group.useGroups(bot, update, wg); ok {
			return true
		}
	}

	if ok := h.useHandlers(bot, update, wg); ok {
		return true
	}

	return false
}

// Handle registers new handler in the group, update will be processed only by first-matched handler,
// order of registration determines the order of matching handlers.
// Important to notice, update's context will be automatically canceled once the handler will finish processing.
// Note: All handlers will process updates in parallel, there is no guaranty on order of processed updates, also keep
// in mind that predicates are checked sequentially.
//
// Warning: Panics if nil handler or predicates passed
func (h *HandlerGroup) Handle(handler Handler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil handlers not allowed")
	}

	for _, p := range predicates {
		if p == nil {
			panic("Telego: nil predicates not allowed")
		}
	}

	h.handlers = append(h.handlers, conditionalHandler{
		Handler:    handler,
		Predicates: predicates,
	})
}

// Group creates a new group of handlers and middlewares from the parent group
// Note: Updates first checked by group and only after that by handler
//
// Warning: Panics if nil predicates passed
func (h *HandlerGroup) Group(predicates ...Predicate) *HandlerGroup {
	for _, p := range predicates {
		if p == nil {
			panic("Telego: nil predicates not allowed")
		}
	}

	group := &HandlerGroup{
		parent:     h,
		predicates: predicates,
	}
	h.groups = append(h.groups, group)

	return group
}

// Use applies middleware to the group
// Note: The Handler chain will be stopped if middleware doesn't call the next func
//
// Warning: Panics if nil middlewares passed
func (h *HandlerGroup) Use(middlewares ...Middleware) {
	for _, m := range middlewares {
		if m == nil {
			panic("Telego: nil middlewares not allowed")
		}
	}

	h.middlewares = append(h.middlewares, middlewares...)
}
