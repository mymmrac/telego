package telegoflow

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

// FlowRunner is implemented by typed flows and accepted by [Manager.Register].
//
// This interface is sealed intentionally; create values with [New] and [Builder.Build]
// instead of implementing it manually.
type FlowRunner interface {
	ID() string
	Start(ctx *th.Context, update telego.Update) error

	register(manager *Manager) error
	startStepID() string
	timeoutDuration() time.Duration
	encodeInitialData(data any, session *SessionState) error
	enterSession(ctx *th.Context, update telego.Update, session *SessionState) error
	handleUpdate(ctx *th.Context, update telego.Update, session *SessionState) error
	timeoutSession(ctx *th.Context, update telego.Update, session *SessionState) error
	cancelSession(ctx *th.Context, update telego.Update, session *SessionState) error
}

var errInvalidInitialData = errors.New("invalid initial flow data type")

// Flow is a typed conversation graph.
type Flow[T any] struct {
	manager *Manager

	id          string
	steps       map[string]*Step[T]
	startStep   string
	timeout     time.Duration
	middlewares []Middleware[T]
	onComplete  Handler[T]
	onCancel    Handler[T]
	onTimeout   Handler[T]
	onError     ErrorHandler[T]
}

// ID returns the flow ID.
func (f *Flow[T]) ID() string {
	return f.id
}

// Start starts this flow with a zero value of T as session data.
func (f *Flow[T]) Start(ctx *th.Context, update telego.Update) error {
	var data T
	return f.startSession(ctx, update, data)
}

// StartWithData returns a telegohandler handler that starts this flow with the provided initial data.
func (f *Flow[T]) StartWithData(data T) th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		return f.startSession(ctx, update, data)
	}
}

func (f *Flow[T]) startSession(ctx *th.Context, update telego.Update, data T) error {
	if f.manager == nil {
		return FlowNotRegisteredError{FlowID: f.id}
	}
	return f.manager.startFlow(ctx, update, f, data)
}

func (f *Flow[T]) register(manager *Manager) error {
	f.manager = manager
	return nil
}

func (f *Flow[T]) startStepID() string {
	return f.startStep
}

func (f *Flow[T]) timeoutDuration() time.Duration {
	return f.timeout
}

func (f *Flow[T]) encodeInitialData(data any, session *SessionState) error {
	typed, ok := data.(T)
	if !ok {
		return SessionDataError{FlowID: f.id, Err: errInvalidInitialData}
	}

	encoded, err := json.Marshal(typed)
	if err != nil {
		return SessionDataError{FlowID: f.id, Err: err}
	}
	session.Data = encoded
	return nil
}

func (f *Flow[T]) enterSession(ctx *th.Context, update telego.Update, session *SessionState) error {
	data, err := f.decodeData(session)
	if err != nil {
		return err
	}
	return f.enterStep(ctx, update, session, data)
}

func (f *Flow[T]) handleUpdate(ctx *th.Context, update telego.Update, session *SessionState) error {
	data, err := f.decodeData(session)
	if err != nil {
		return err
	}

	flowCtx := f.newContext(ctx, update, session, data)
	step, ok := f.steps[session.CurrentStep]
	if !ok {
		return StepNotFoundError{FlowID: f.id, StepID: session.CurrentStep}
	}

	for _, route := range step.routes {
		if route.match(ctx.Context(), update) {
			return f.runHandler(flowCtx, step, route.handler)
		}
	}

	if step.fallback != nil {
		return f.runHandler(flowCtx, step, step.fallback)
	}

	return flowCtx.Stay()
}

func (f *Flow[T]) timeoutSession(ctx *th.Context, update telego.Update, session *SessionState) error {
	data, err := f.decodeData(session)
	if err != nil {
		return err
	}

	flowCtx := f.newContext(ctx, update, session, data)
	flowCtx.done = true
	if f.onTimeout != nil {
		if err = f.onTimeout(flowCtx); err != nil {
			return err
		}
	}
	return f.manager.storage.DeleteSession(ctx.Context(), session.Key)
}

func (f *Flow[T]) cancelSession(ctx *th.Context, update telego.Update, session *SessionState) error {
	data, err := f.decodeData(session)
	if err != nil {
		return err
	}

	flowCtx := f.newContext(ctx, update, session, data)
	return flowCtx.Cancel()
}

func (f *Flow[T]) enterStep(ctx *th.Context, update telego.Update, session *SessionState, data *T) error {
	flowCtx := f.newContext(ctx, update, session, data)
	step, ok := f.steps[session.CurrentStep]
	if !ok {
		return StepNotFoundError{FlowID: f.id, StepID: session.CurrentStep}
	}
	if step.enter == nil {
		return flowCtx.Stay()
	}
	return f.runHandler(flowCtx, step, step.enter)
}

func (f *Flow[T]) runHandler(ctx *Context[T], step *Step[T], handler Handler[T]) error {
	wrapped := handler
	for i := len(step.middlewares) - 1; i >= 0; i-- {
		wrapped = step.middlewares[i](wrapped)
	}
	for i := len(f.middlewares) - 1; i >= 0; i-- {
		wrapped = f.middlewares[i](wrapped)
	}

	err := wrapped(ctx)
	if err != nil {
		if f.onError != nil {
			return f.onError(ctx, err)
		}
		return err
	}
	if ctx.done {
		return nil
	}
	return ctx.Stay()
}

func (f *Flow[T]) newContext(ctx *th.Context, update telego.Update, session *SessionState, data *T) *Context[T] {
	return &Context[T]{
		Context: ctx,
		manager: f.manager,
		flow:    f,
		update:  update,
		session: session,
		data:    data,
	}
}

func (f *Flow[T]) decodeData(session *SessionState) (*T, error) {
	data := new(T)
	if len(session.Data) == 0 {
		return data, nil
	}
	if err := json.Unmarshal(session.Data, data); err != nil {
		return nil, SessionDataError{FlowID: f.id, Err: err}
	}
	return data, nil
}
