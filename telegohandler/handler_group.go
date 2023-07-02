package telegohandler

import (
	"sync"

	"github.com/mymmrac/telego"
)

// conditionalHandler represents handler with respectful predicates
type conditionalHandler struct {
	handler    Handler
	predicates []Predicate
}

// match matches the current update and handler
func (h conditionalHandler) match(update telego.Update) bool {
	update = update.Clone()
	for _, p := range h.predicates {
		if !p(update) {
			return false
		}
	}
	return true
}

// HandlerGroup represents a group of handlers, middlewares and child groups
type HandlerGroup struct {
	predicates  []Predicate
	middlewares []Middleware
	groups      []*HandlerGroup
	handlers    []conditionalHandler
}

// match matches the current update and group
func (h *HandlerGroup) match(update telego.Update) bool {
	update = update.Clone()
	for _, p := range h.predicates {
		if !p(update) {
			return false
		}
	}
	return true
}

// processUpdate checks all group predicates, runs middlewares, checks handler predicates,
// tries to process update in first matched handler
func (h *HandlerGroup) processUpdate(bot *telego.Bot, update telego.Update) {
	_ = h.processUpdateWithMiddlewares(bot, update, h.middlewares)
}

func (h *HandlerGroup) processUpdateWithMiddlewares(
	bot *telego.Bot, update telego.Update, middlewares []Middleware,
) bool {
	ctx := update.Context()
	select {
	case <-ctx.Done():
		return false
	default:
		// Continue
	}

	// Check group predicates once
	if len(middlewares) == len(h.middlewares) && !h.match(update) {
		return false
	}

	// Process all middlewares
	if len(middlewares) != 0 {
		once := sync.Once{}
		done := make(chan bool, 1)
		middlewares[0](bot, update, func(bot *telego.Bot, update telego.Update) {
			once.Do(func() {
				done <- h.processUpdateWithMiddlewares(bot, update, middlewares[1:])
			})
		})

		select {
		case <-ctx.Done():
			return false
		case matched := <-done:
			return matched
		}
	}

	// Process all groups
	for _, group := range h.groups {
		if group.processUpdateWithMiddlewares(bot, update, group.middlewares) {
			return true
		}
	}

	// Process all handlers
	for _, handler := range h.handlers {
		if handler.match(update) {
			handler.handler(bot, update)
			return true
		}
	}

	return false
}

// Handle registers new handler in the group, update will be processed only by first-matched handler,
// order of registration determines the order of matching handlers.
// Important to notice, update's context will be automatically canceled once the handler will finish processing or
// the bot handler stopped.
// Note: All handlers will process updates in parallel, there is no guaranty on order of processed updates, also keep
// in mind that middlewares and predicates are checked sequentially.
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
		handler:    handler,
		predicates: predicates,
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
		predicates: predicates,
	}
	h.groups = append(h.groups, group)

	return group
}

// Use applies middleware to the group
// Note: The chain will be stopped if middleware doesn't call the next func,
// if there is no context timeout then update will be stuck,
// if there is time out then the group will be skipped since not all middlewares were called
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
