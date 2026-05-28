package telegoflow

import "time"

// Builder provides a fluent API for constructing a typed flow.
type Builder[T any] struct {
	id          string
	steps       []*Step[T]
	startStep   string
	timeout     time.Duration
	middlewares []Middleware[T]
	onComplete  Handler[T]
	onCancel    Handler[T]
	onTimeout   Handler[T]
	onError     ErrorHandler[T]
}

// New creates a builder for a typed flow.
func New[T any](id string) *Builder[T] {
	return &Builder[T]{id: id}
}

// Steps adds one or more steps to the flow.
func (b *Builder[T]) Steps(steps ...*Step[T]) *Builder[T] {
	b.steps = append(b.steps, steps...)
	return b
}

// StartWith sets the initial step.
//
// If StartWith isn't called, the first added step is used.
func (b *Builder[T]) StartWith(stepID string) *Builder[T] {
	b.startStep = stepID
	return b
}

// WithTimeout sets the session timeout for this flow.
func (b *Builder[T]) WithTimeout(timeout time.Duration) *Builder[T] {
	b.timeout = timeout
	return b
}

// Use adds middleware for all steps in this flow.
func (b *Builder[T]) Use(middlewares ...Middleware[T]) *Builder[T] {
	b.middlewares = append(b.middlewares, middlewares...)
	return b
}

// OnComplete sets a hook called when the flow completes.
func (b *Builder[T]) OnComplete(handler Handler[T]) *Builder[T] {
	b.onComplete = handler
	return b
}

// OnCancel sets a hook called when the flow is canceled.
func (b *Builder[T]) OnCancel(handler Handler[T]) *Builder[T] {
	b.onCancel = handler
	return b
}

// OnTimeout sets a hook called when the flow session expires.
func (b *Builder[T]) OnTimeout(handler Handler[T]) *Builder[T] {
	b.onTimeout = handler
	return b
}

// OnError sets a hook called when a step handler returns an error.
func (b *Builder[T]) OnError(handler ErrorHandler[T]) *Builder[T] {
	b.onError = handler
	return b
}

// Build validates and creates a flow.
func (b *Builder[T]) Build() (*Flow[T], error) {
	if len(b.steps) == 0 {
		return nil, ErrEmptyFlow
	}

	steps := make(map[string]*Step[T], len(b.steps))
	for _, step := range b.steps {
		if step == nil || step.id == "" {
			return nil, StepNotFoundError{FlowID: b.id}
		}
		if _, ok := steps[step.id]; ok {
			return nil, DuplicateStepError{StepID: step.id}
		}
		steps[step.id] = step
	}

	startStep := b.startStep
	if startStep == "" {
		startStep = b.steps[0].id
	}
	if _, ok := steps[startStep]; !ok {
		return nil, InvalidStartStepError{FlowID: b.id, StepID: startStep}
	}

	for _, step := range steps {
		for to := range step.canGo {
			if _, ok := steps[to]; !ok {
				return nil, InvalidTransitionError{FlowID: b.id, From: step.id, To: to}
			}
		}
	}

	return &Flow[T]{
		id:          b.id,
		steps:       steps,
		startStep:   startStep,
		timeout:     b.timeout,
		middlewares: b.middlewares,
		onComplete:  b.onComplete,
		onCancel:    b.onCancel,
		onTimeout:   b.onTimeout,
		onError:     b.onError,
	}, nil
}
