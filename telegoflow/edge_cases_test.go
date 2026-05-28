package telegoflow

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type badData struct {
	Ch chan int `json:"ch"`
}

type badSaveData struct {
	Value any `json:"value,omitempty"`
}

type fakeFlowRunner struct {
	id  string
	err error
}

func (f fakeFlowRunner) ID() string                                     { return f.id }
func (f fakeFlowRunner) Start(_ *th.Context, _ telego.Update) error     { return nil }
func (f fakeFlowRunner) register(_ *Manager) error                      { return f.err }
func (f fakeFlowRunner) startStepID() string                            { return "" }
func (f fakeFlowRunner) timeoutDuration() time.Duration                 { return 0 }
func (f fakeFlowRunner) encodeInitialData(_ any, _ *SessionState) error { return nil }
func (f fakeFlowRunner) enterSession(_ *th.Context, _ telego.Update, _ *SessionState) error {
	return nil
}

func (f fakeFlowRunner) handleUpdate(_ *th.Context, _ telego.Update, _ *SessionState) error {
	return nil
}

func (f fakeFlowRunner) timeoutSession(_ *th.Context, _ telego.Update, _ *SessionState) error {
	return nil
}

func (f fakeFlowRunner) cancelSession(_ *th.Context, _ telego.Update, _ *SessionState) error {
	return nil
}

func TestFlow_DataEncodingErrors(t *testing.T) {
	t.Run("initial_data_marshal_error", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		flow, err := New[badData]("flow").Steps(NewStep[badData]("start")).Build()
		require.NoError(t, err)
		require.NoError(t, manager.Register(flow))

		err = flow.StartWithData(badData{Ch: make(chan int)})(&th.Context{}, messageUpdate("/start"))

		var dataErr SessionDataError
		require.ErrorAs(t, err, &dataErr)
		assert.Equal(t, "flow", dataErr.FlowID)
	})

	t.Run("save_marshal_error", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		flow, err := New[badSaveData]("flow").
			Steps(NewStep[badSaveData]("start").Handle(func(ctx *Context[badSaveData]) error {
				ctx.Data().Value = make(chan int)
				return nil
			})).
			Build()
		require.NoError(t, err)
		require.NoError(t, manager.Register(flow))
		require.NoError(t, flow.Start(&th.Context{}, messageUpdate("/start")))

		err = handleUpdate(t, newTestGroup(manager), messageUpdate("boom"))

		var dataErr SessionDataError
		require.ErrorAs(t, err, &dataErr)
		assert.Equal(t, "flow", dataErr.FlowID)
	})
}

func TestFlow_DecodeAndStepErrors(t *testing.T) {
	t.Run("enter_session_bad_json", func(t *testing.T) {
		flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
		require.NoError(t, err)

		err = flow.enterSession(&th.Context{}, messageUpdate("/start"), &SessionState{Data: json.RawMessage(`{bad`)})

		var dataErr SessionDataError
		require.ErrorAs(t, err, &dataErr)
	})

	t.Run("handle_update_missing_step", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
		require.NoError(t, err)
		require.NoError(t, flow.register(manager))

		err = flow.handleUpdate(&th.Context{}, messageUpdate("hello"), &SessionState{
			FlowID:      "flow",
			CurrentStep: "missing",
			Data:        json.RawMessage(`{}`),
		})

		var stepErr StepNotFoundError
		require.ErrorAs(t, err, &stepErr)
		assert.Equal(t, "missing", stepErr.StepID)
	})

	t.Run("enter_step_missing_step", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
		require.NoError(t, err)
		require.NoError(t, flow.register(manager))

		err = flow.enterStep(&th.Context{}, messageUpdate("hello"), &SessionState{CurrentStep: "missing"}, &testData{})

		var stepErr StepNotFoundError
		require.ErrorAs(t, err, &stepErr)
		assert.Equal(t, "missing", stepErr.StepID)
	})
}

func TestFlow_LifecycleHookErrors(t *testing.T) {
	t.Run("on_complete_error", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		flow, err := New[testData]("flow").
			Steps(NewStep[testData]("start").Enter(func(ctx *Context[testData]) error { return ctx.Complete() })).
			OnComplete(func(_ *Context[testData]) error { return errTest }).
			Build()
		require.NoError(t, err)
		require.NoError(t, manager.Register(flow))

		err = flow.Start(&th.Context{}, messageUpdate("/start"))

		require.ErrorIs(t, err, errTest)
	})

	t.Run("on_cancel_error", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		flow, err := New[testData]("flow").
			Steps(NewStep[testData]("start").Enter(func(ctx *Context[testData]) error { return ctx.Cancel() })).
			OnCancel(func(_ *Context[testData]) error { return errTest }).
			Build()
		require.NoError(t, err)
		require.NoError(t, manager.Register(flow))

		err = flow.Start(&th.Context{}, messageUpdate("/start"))

		require.ErrorIs(t, err, errTest)
	})

	t.Run("on_timeout_error", func(t *testing.T) {
		storage := NewMemoryStorage()
		manager := NewManager(storage)
		manager.nowFunc = fixedNow
		flow, err := New[testData]("flow").
			Steps(NewStep[testData]("start")).
			OnTimeout(func(_ *Context[testData]) error { return errTest }).
			Build()
		require.NoError(t, err)
		require.NoError(t, manager.Register(flow))
		expiresAt := fixedNow().Add(-time.Second)
		require.NoError(t, storage.SaveSession(t.Context(), &SessionState{
			Key:         SessionKey{ChatID: 123, UserID: 456},
			FlowID:      "flow",
			CurrentStep: "start",
			Data:        json.RawMessage(`{}`),
			ExpiresAt:   &expiresAt,
		}))

		err = handleUpdate(t, newTestGroup(manager), messageUpdate("hello"))

		require.ErrorIs(t, err, errTest)
	})
}

func TestContext_GoMoreEdgeCases(t *testing.T) {
	t.Run("target_missing_after_build", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		flow, err := New[testData]("flow").
			Steps(NewStep[testData]("start").CanGo("next"), NewStep[testData]("next")).
			Build()
		require.NoError(t, err)
		delete(flow.steps, "next")
		ctx := &Context[testData]{
			manager: manager,
			flow:    flow,
			Context: &th.Context{},
			session: &SessionState{CurrentStep: "start"},
		}

		err = ctx.Go("next")

		var stepErr StepNotFoundError
		require.ErrorAs(t, err, &stepErr)
		assert.Equal(t, "next", stepErr.StepID)
	})

	t.Run("load_after_go_enter_error", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		flow, err := New[testData]("flow").
			Steps(
				NewStep[testData]("start").CanGo("next"),
				NewStep[testData]("next").Enter(func(_ *Context[testData]) error { return errTest }),
			).
			Build()
		require.NoError(t, err)
		require.NoError(t, manager.Register(flow))

		ctx := &Context[testData]{
			manager: manager,
			flow:    flow,
			Context: &th.Context{},
			session: &SessionState{
				Key:         SessionKey{ChatID: 123, UserID: 456},
				FlowID:      "flow",
				CurrentStep: "start",
				Data:        json.RawMessage(`{}`),
			},
			data: &testData{},
		}

		err = ctx.Go("next")

		require.ErrorIs(t, err, errTest)
	})
}

func TestManager_OptionAndRegisterEdges(t *testing.T) {
	t.Run("nil_storage_and_nil_key_func", func(t *testing.T) {
		manager := NewManager(nil, WithKeyFunc(nil))
		assert.NotNil(t, manager.storage)
		key, ok := manager.keyFunc(messageUpdate("hello"))
		assert.True(t, ok)
		assert.Equal(t, SessionKey{ChatID: 123, UserID: 456}, key)
	})

	t.Run("register_error", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		err := manager.Register(fakeFlowRunner{id: "fake", err: errTest})

		require.ErrorIs(t, err, errTest)
	})
}

func TestStep_CanGoInitializesNilMap(t *testing.T) {
	step := &Step[testData]{id: "step"}
	step.CanGo("next")
	assert.True(t, step.allows("next"))
}
