package telegoflow

import (
	"errors"
	"fmt"
)

// ErrEmptyFlow is returned when a flow is built without steps.
var ErrEmptyFlow = errors.New("telego: flow must have at least one step")

// NoSessionKeyError is returned when a session key can't be extracted from an update.
type NoSessionKeyError struct{}

func (e NoSessionKeyError) Error() string {
	return "telego: flow session key not found in update"
}

// FlowNotFoundError is returned when a session references an unknown flow.
type FlowNotFoundError struct {
	FlowID string
}

func (e FlowNotFoundError) Error() string {
	return fmt.Sprintf("telego: flow %q not found", e.FlowID)
}

// DuplicateFlowError is returned when a manager already has a flow with the same ID.
type DuplicateFlowError struct {
	FlowID string
}

func (e DuplicateFlowError) Error() string {
	return fmt.Sprintf("telego: duplicate flow %q", e.FlowID)
}

// FlowNotRegisteredError is returned when a flow is started before it was registered in a manager.
type FlowNotRegisteredError struct {
	FlowID string
}

func (e FlowNotRegisteredError) Error() string {
	return fmt.Sprintf("telego: flow %q is not registered in a manager", e.FlowID)
}

// StepNotFoundError is returned when a step is missing in the flow.
type StepNotFoundError struct {
	FlowID string
	StepID string
}

func (e StepNotFoundError) Error() string {
	return fmt.Sprintf("telego: step %q not found in flow %q", e.StepID, e.FlowID)
}

// DuplicateStepError is returned when a flow contains multiple steps with the same ID.
type DuplicateStepError struct {
	StepID string
}

func (e DuplicateStepError) Error() string {
	return fmt.Sprintf("telego: duplicate step %q", e.StepID)
}

// InvalidStartStepError is returned when the configured start step doesn't exist.
type InvalidStartStepError struct {
	FlowID string
	StepID string
}

func (e InvalidStartStepError) Error() string {
	return fmt.Sprintf("telego: start step %q not found in flow %q", e.StepID, e.FlowID)
}

// InvalidTransitionError is returned when a step declares a transition to an unknown step.
type InvalidTransitionError struct {
	FlowID string
	From   string
	To     string
}

func (e InvalidTransitionError) Error() string {
	return fmt.Sprintf("telego: transition from step %q to step %q in flow %q is invalid", e.From, e.To, e.FlowID)
}

// TransitionNotAllowedError is returned when a handler tries to jump to a step not declared by CanGo.
type TransitionNotAllowedError struct {
	FlowID string
	From   string
	To     string
}

func (e TransitionNotAllowedError) Error() string {
	return fmt.Sprintf("telego: transition from step %q to step %q in flow %q is not allowed", e.From, e.To, e.FlowID)
}

// SessionDataError is returned when typed session data can't be encoded or decoded.
type SessionDataError struct {
	FlowID string
	Err    error
}

func (e SessionDataError) Error() string {
	return fmt.Sprintf("telego: session data error in flow %q: %v", e.FlowID, e.Err)
}

func (e SessionDataError) Unwrap() error {
	return e.Err
}
