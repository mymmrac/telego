package telegohandler

import (
	"context"
	"slices"

	"github.com/mymmrac/telego"
)

// route represents handler, middleware or group with respectful predicates
type route struct {
	predicates []Predicate

	group   *HandlerGroup
	handler Handler
}

// match matches the current update by predicates
func (r route) match(ctx context.Context, update telego.Update) bool {
	if len(r.predicates) == 0 {
		return true
	}

	update = update.Clone()
	for _, p := range r.predicates {
		if !p(ctx, update) {
			return false
		}
	}

	return true
}

// HandlerGroup represents a group of middlewares and routes (handlers and subgroups)
type HandlerGroup struct {
	parent *HandlerGroup
	routes []route
}

// depth returns the depth of the group's routes
func (h *HandlerGroup) depth(depth int) int {
	localDepth := depth
	for _, r := range h.routes {
		if r.group != nil {
			depth = max(depth, r.group.depth(localDepth+1))
		}
	}
	return depth
}

// Handle registers new handler in the group, update will be processed only by first-matched route,
// order of registration determines the order of matching routes.
// Important to notice handler's context will be automatically canceled once the handler will finish processing or
// the bot handler stopped.
// Note: All handlers will process updates in parallel, there is no guaranty on order of processed updates, also keep
// in mind that middlewares and predicates are run sequentially.
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

	h.routes = append(h.routes, route{
		predicates: predicates,
		handler:    handler,
	})
}

// Group creates a new group of handlers and middlewares from the parent group, update will be processed only by
// first-matched route, order of registration determines the order of matching routes
//
// Warning: Panics if nil predicates passed
func (h *HandlerGroup) Group(predicates ...Predicate) *HandlerGroup {
	for _, p := range predicates {
		if p == nil {
			panic("Telego: nil predicates not allowed")
		}
	}

	group := &HandlerGroup{
		parent: h,
	}

	h.routes = append(h.routes, route{
		predicates: predicates,
		group:      group,
	})

	return group
}

// Use applies middleware to the group, update will be processed only by first-matched route,
// order of registration determines the order of matching routes.
// Note: The chain will be stopped if middleware doesn't call the [Context.Next]
//
// Warning: Panics if nil middlewares passed
func (h *HandlerGroup) Use(middlewares ...Handler) {
	for _, m := range middlewares {
		if m == nil {
			panic("Telego: nil middlewares not allowed")
		}
	}

	h.routes = slices.Grow(h.routes, len(middlewares))
	for _, middleware := range middlewares {
		h.routes = append(h.routes, route{
			handler: middleware,
		})
	}
}
