package telegoflow

import (
	"encoding/json"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

// Context is passed to flow handlers.
//
// It embeds [telegohandler.Context], so handlers can use Bot, deadlines,
// cancellation, and values exactly like regular telegohandler handlers.
type Context[T any] struct {
	*th.Context

	manager *Manager
	flow    *Flow[T]
	update  telego.Update
	session *SessionState
	data    *T
	done    bool
}

// Update returns the update currently handled by the flow.
func (c *Context[T]) Update() telego.Update {
	return c.update
}

// Data returns typed session data.
func (c *Context[T]) Data() *T {
	return c.data
}

// FlowID returns the current flow ID.
func (c *Context[T]) FlowID() string {
	return c.flow.id
}

// StepID returns the current step ID.
func (c *Context[T]) StepID() string {
	return c.session.CurrentStep
}

// SessionKey returns the current session key.
func (c *Context[T]) SessionKey() SessionKey {
	return c.session.Key
}

// UserID returns the user ID part of the current session key.
func (c *Context[T]) UserID() int64 {
	return c.session.Key.UserID
}

// ChatID returns the chat ID part of the current session key.
func (c *Context[T]) ChatID() telego.ChatID {
	return telego.ChatID{ID: c.session.Key.ChatID}
}

// Message returns the message carried by the current update, if any.
func (c *Context[T]) Message() *telego.Message {
	switch {
	case c.update.Message != nil:
		return c.update.Message
	case c.update.EditedMessage != nil:
		return c.update.EditedMessage
	case c.update.BusinessMessage != nil:
		return c.update.BusinessMessage
	case c.update.EditedBusinessMessage != nil:
		return c.update.EditedBusinessMessage
	case c.update.GuestMessage != nil:
		return c.update.GuestMessage
	case c.update.CallbackQuery != nil && c.update.CallbackQuery.Message != nil:
		return c.update.CallbackQuery.Message.Message()
	default:
		return nil
	}
}

// CallbackQuery returns the callback query carried by the current update, if any.
func (c *Context[T]) CallbackQuery() *telego.CallbackQuery {
	return c.update.CallbackQuery
}

// Text returns message text from the current update, if any.
func (c *Context[T]) Text() string {
	msg := c.Message()
	if msg == nil {
		return ""
	}
	return msg.Text
}

// Go transitions to another step and executes its enter handler.
func (c *Context[T]) Go(stepID string) error {
	currentStep, ok := c.flow.steps[c.session.CurrentStep]
	if !ok {
		return StepNotFoundError{FlowID: c.flow.id, StepID: c.session.CurrentStep}
	}
	if !currentStep.allows(stepID) {
		return TransitionNotAllowedError{FlowID: c.flow.id, From: c.session.CurrentStep, To: stepID}
	}
	if _, ok = c.flow.steps[stepID]; !ok {
		return StepNotFoundError{FlowID: c.flow.id, StepID: stepID}
	}

	c.session.CurrentStep = stepID
	c.session.UpdatedAt = c.manager.now()
	if err := c.save(); err != nil {
		return err
	}

	if err := c.flow.enterStep(c.Context, c.update, c.session, c.data); err != nil {
		return err
	}

	_, exists, err := c.manager.storage.LoadSession(c.Context.Context(), c.session.Key)
	if err != nil {
		return err
	}
	if !exists {
		c.done = true
	}
	return nil
}

// Stay persists typed data and keeps the session on the current step.
func (c *Context[T]) Stay() error {
	c.session.UpdatedAt = c.manager.now()
	return c.save()
}

// Complete completes the flow, runs its completion hook, and deletes the session.
func (c *Context[T]) Complete() error {
	c.done = true
	if c.flow.onComplete != nil {
		if err := c.flow.onComplete(c); err != nil {
			return err
		}
	}
	return c.manager.storage.DeleteSession(c.Context.Context(), c.session.Key)
}

// Cancel cancels the flow, runs its cancellation hook, and deletes the session.
func (c *Context[T]) Cancel() error {
	c.done = true
	if c.flow.onCancel != nil {
		if err := c.flow.onCancel(c); err != nil {
			return err
		}
	}
	return c.manager.storage.DeleteSession(c.Context.Context(), c.session.Key)
}

func (c *Context[T]) save() error {
	data, err := json.Marshal(c.data)
	if err != nil {
		return SessionDataError{FlowID: c.flow.id, Err: err}
	}

	c.session.Data = data
	return c.manager.storage.SaveSession(c.Context.Context(), c.session)
}
