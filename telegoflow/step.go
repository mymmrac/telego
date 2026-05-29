package telegoflow

import (
	"context"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

// Handler handles a flow step or lifecycle event.
type Handler[T any] func(ctx *Context[T]) error

// ErrorHandler handles an error returned by a step handler.
type ErrorHandler[T any] func(ctx *Context[T], err error) error

// Middleware wraps a flow handler.
type Middleware[T any] func(next Handler[T]) Handler[T]

type stepRoute[T any] struct {
	predicates []th.Predicate
	handler    Handler[T]
}

func (r stepRoute[T]) match(ctx context.Context, update telego.Update) bool {
	for _, predicate := range r.predicates {
		if !predicate(ctx, update) {
			return false
		}
	}
	return true
}

// Step represents one state in a conversation flow.
type Step[T any] struct {
	id          string
	enter       Handler[T]
	routes      []stepRoute[T]
	fallback    Handler[T]
	middlewares []Middleware[T]
	canGo       map[string]struct{}
	canComplete bool
}

// NewStep creates a step with the provided ID.
func NewStep[T any](id string) *Step[T] {
	return &Step[T]{
		id:    id,
		canGo: make(map[string]struct{}),
	}
}

// ID returns the step ID.
func (s *Step[T]) ID() string {
	if s == nil {
		return ""
	}
	return s.id
}

// Enter sets a handler that is executed when the session enters this step.
func (s *Step[T]) Enter(handler Handler[T]) *Step[T] {
	s.enter = handler
	return s
}

// Handle adds an input route to the step.
//
// When all predicates match, the route handler is executed. Routes are checked
// in the order they were added. If no predicates are provided, the route matches
// any update.
func (s *Step[T]) Handle(handler Handler[T], predicates ...th.Predicate) *Step[T] {
	s.routes = append(s.routes, stepRoute[T]{
		predicates: predicates,
		handler:    handler,
	})
	return s
}

// Fallback sets a handler that is executed when no input route matches.
func (s *Step[T]) Fallback(handler Handler[T]) *Step[T] {
	s.fallback = handler
	return s
}

// Use adds middleware for this step.
func (s *Step[T]) Use(middlewares ...Middleware[T]) *Step[T] {
	s.middlewares = append(s.middlewares, middlewares...)
	return s
}

// CanGo declares controlled transitions allowed from this step.
func (s *Step[T]) CanGo(stepIDs ...string) *Step[T] {
	if s.canGo == nil {
		s.canGo = make(map[string]struct{}, len(stepIDs))
	}
	for _, stepID := range stepIDs {
		s.canGo[stepID] = struct{}{}
	}
	return s
}

// CanComplete declares that this step can complete the flow.
//
// Context.Complete is still allowed from any handler; this declaration is used
// by Flow.Graph to show the intended flow exit explicitly.
func (s *Step[T]) CanComplete() *Step[T] {
	s.canComplete = true
	return s
}

func (s *Step[T]) allows(stepID string) bool {
	_, ok := s.canGo[stepID]
	return ok
}
