package telegoflow

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilder_Build(t *testing.T) {
	t.Run("empty_flow", func(t *testing.T) {
		flow, err := New[testData]("empty").Build()

		require.ErrorIs(t, err, ErrEmptyFlow)
		assert.Nil(t, flow)
	})

	t.Run("nil_step", func(t *testing.T) {
		flow, err := New[testData]("flow").Steps(nil).Build()

		var stepErr StepNotFoundError
		require.ErrorAs(t, err, &stepErr)
		assert.Equal(t, "flow", stepErr.FlowID)
		assert.Nil(t, flow)
	})

	t.Run("empty_step_id", func(t *testing.T) {
		flow, err := New[testData]("flow").Steps(NewStep[testData]("")).Build()

		var stepErr StepNotFoundError
		require.ErrorAs(t, err, &stepErr)
		assert.Equal(t, "flow", stepErr.FlowID)
		assert.Nil(t, flow)
	})

	t.Run("duplicate_step", func(t *testing.T) {
		flow, err := New[testData]("flow").
			Steps(NewStep[testData]("same"), NewStep[testData]("same")).
			Build()

		var duplicateErr DuplicateStepError
		require.ErrorAs(t, err, &duplicateErr)
		assert.Equal(t, "same", duplicateErr.StepID)
		assert.Nil(t, flow)
	})

	t.Run("invalid_start_step", func(t *testing.T) {
		flow, err := New[testData]("flow").
			Steps(NewStep[testData]("start")).
			StartWith("missing").
			Build()

		var startErr InvalidStartStepError
		require.ErrorAs(t, err, &startErr)
		assert.Equal(t, "flow", startErr.FlowID)
		assert.Equal(t, "missing", startErr.StepID)
		assert.Nil(t, flow)
	})

	t.Run("invalid_transition", func(t *testing.T) {
		flow, err := New[testData]("flow").
			Steps(NewStep[testData]("start").CanGo("missing")).
			Build()

		var transitionErr InvalidTransitionError
		require.ErrorAs(t, err, &transitionErr)
		assert.Equal(t, "flow", transitionErr.FlowID)
		assert.Equal(t, "start", transitionErr.From)
		assert.Equal(t, "missing", transitionErr.To)
		assert.Nil(t, flow)
	})

	t.Run("success_default_start", func(t *testing.T) {
		flow, err := New[testData]("flow").
			Steps(
				NewStep[testData]("start").CanGo("finish"),
				NewStep[testData]("finish"),
			).
			WithTimeout(time.Minute).
			Build()

		require.NoError(t, err)
		require.NotNil(t, flow)
		assert.Equal(t, "flow", flow.ID())
		assert.Equal(t, "start", flow.startStepID())
		assert.Equal(t, time.Minute, flow.timeoutDuration())
	})

	t.Run("success_explicit_start", func(t *testing.T) {
		flow, err := New[testData]("flow").
			Steps(NewStep[testData]("first"), NewStep[testData]("second")).
			StartWith("second").
			Build()

		require.NoError(t, err)
		require.NotNil(t, flow)
		assert.Equal(t, "second", flow.startStepID())
	})
}

func TestStep_BuilderMethods(t *testing.T) {
	step := NewStep[testData]("step")
	enter := Handler[testData](func(_ *Context[testData]) error { return nil })
	handle := Handler[testData](func(_ *Context[testData]) error { return nil })
	fallback := Handler[testData](func(_ *Context[testData]) error { return nil })
	mw := Middleware[testData](func(next Handler[testData]) Handler[testData] { return next })

	returned := step.Enter(enter).
		Handle(handle).
		Fallback(fallback).
		Use(mw).
		CanGo("next")

	assert.Same(t, step, returned)
	assert.Equal(t, "step", step.ID())
	assert.NotNil(t, step.enter)
	require.Len(t, step.routes, 1)
	assert.NotNil(t, step.routes[0].handler)
	assert.NotNil(t, step.fallback)
	require.Len(t, step.middlewares, 1)
	assert.True(t, step.allows("next"))
	assert.False(t, step.allows("missing"))
	assert.Empty(t, (*Step[testData])(nil).ID())
}

func TestFlow_StartBeforeRegister(t *testing.T) {
	flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
	require.NoError(t, err)

	err = flow.Start(nil, messageUpdate("/start"))

	var unregisteredErr FlowNotRegisteredError
	require.ErrorAs(t, err, &unregisteredErr)
	assert.Equal(t, "flow", unregisteredErr.FlowID)
}

func TestFlow_EncodeInitialDataTypeMismatch(t *testing.T) {
	flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
	require.NoError(t, err)

	err = flow.encodeInitialData("wrong", &SessionState{})

	var dataErr SessionDataError
	require.ErrorAs(t, err, &dataErr)
	assert.Equal(t, "flow", dataErr.FlowID)
	assert.ErrorIs(t, dataErr.Unwrap(), errInvalidInitialData)
}
