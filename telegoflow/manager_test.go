package telegoflow

import (
	"context"
	"encoding/json"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func TestManager_Register(t *testing.T) {
	manager := NewManager(NewMemoryStorage())
	flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
	require.NoError(t, err)

	t.Run("success_and_nil_ignored", func(t *testing.T) {
		err = manager.Register(nil, flow)
		require.NoError(t, err)
		assert.Equal(t, manager, flow.manager)
		assert.Equal(t, flow, manager.flows["flow"])
	})

	t.Run("duplicate", func(t *testing.T) {
		err = manager.Register(flow)

		var duplicateErr DuplicateFlowError
		require.ErrorAs(t, err, &duplicateErr)
		assert.Equal(t, "flow", duplicateErr.FlowID)
	})
}

func TestManager_Middleware_StartAndRouteFlow(t *testing.T) {
	storage := NewMemoryStorage()
	manager := NewManager(storage)
	key := SessionKey{ChatID: 123, UserID: 456}

	enteredName := 0
	enteredAge := 0
	completed := false

	flow, err := New[testData]("registration").
		Steps(
			NewStep[testData]("name").
				Enter(func(ctx *Context[testData]) error {
					enteredName++
					assert.Equal(t, "registration", ctx.FlowID())
					assert.Equal(t, "name", ctx.StepID())
					assert.Equal(t, key, ctx.SessionKey())
					assert.Equal(t, int64(456), ctx.UserID())
					assert.Equal(t, telego.ChatID{ID: 123}, ctx.ChatID())
					assert.Equal(t, "/register", ctx.Text())
					assert.NotNil(t, ctx.Message())
					return nil
				}).
				Handle(func(ctx *Context[testData]) error {
					ctx.Data().Name = ctx.Text()
					return ctx.Go("age")
				}, th.AnyMessageWithText()).
				CanGo("age"),
			NewStep[testData]("age").
				Enter(func(ctx *Context[testData]) error {
					enteredAge++
					assert.Equal(t, "Alice", ctx.Data().Name)
					return nil
				}).
				Handle(func(ctx *Context[testData]) error {
					ctx.Data().Count++
					return ctx.Complete()
				}, th.TextEqual("done")),
		).
		OnComplete(func(ctx *Context[testData]) error {
			completed = true
			assert.Equal(t, "Alice", ctx.Data().Name)
			assert.Equal(t, 1, ctx.Data().Count)
			return nil
		}).
		Build()
	require.NoError(t, err)
	require.NoError(t, manager.Register(flow))

	group := newTestGroup(manager, func(group *th.HandlerGroup) {
		group.Handle(flow.Start, th.CommandEqual("register"))
	})

	require.NoError(t, handleUpdate(t, group, messageUpdate("/register")))
	state, data, ok := loadData[testData](t, storage, key)
	require.True(t, ok)
	assert.Equal(t, "name", state.CurrentStep)
	assert.Empty(t, data.Name)
	assert.Equal(t, 1, enteredName)
	assert.Equal(t, 0, enteredAge)

	require.NoError(t, handleUpdate(t, group, messageUpdate("Alice")))
	state, data, ok = loadData[testData](t, storage, key)
	require.True(t, ok)
	assert.Equal(t, "age", state.CurrentStep)
	assert.Equal(t, "Alice", data.Name)
	assert.Equal(t, 1, enteredAge)

	require.NoError(t, handleUpdate(t, group, messageUpdate("ignored")))
	state, data, ok = loadData[testData](t, storage, key)
	require.True(t, ok)
	assert.Equal(t, "age", state.CurrentStep)
	assert.Equal(t, 0, data.Count)
	assert.False(t, completed)

	require.NoError(t, handleUpdate(t, group, messageUpdate("done")))
	_, _, ok = loadData[testData](t, storage, key)
	assert.False(t, ok)
	assert.True(t, completed)
}

func TestManager_Middleware_NoSessionCallsNext(t *testing.T) {
	manager := NewManager(NewMemoryStorage())
	called := false
	group := newTestGroup(manager, func(group *th.HandlerGroup) {
		group.Handle(func(_ *th.Context, update telego.Update) error {
			called = true
			assert.Equal(t, "hello", update.Message.Text)
			return nil
		}, th.AnyMessageWithText())
	})

	require.NoError(t, handleUpdate(t, group, messageUpdate("hello")))
	assert.True(t, called)
}

func TestManager_Middleware_NoKeyCallsNext(t *testing.T) {
	manager := NewManager(NewMemoryStorage())
	called := false
	group := newTestGroup(manager, func(group *th.HandlerGroup) {
		group.Handle(func(_ *th.Context, update telego.Update) error {
			called = true
			assert.NotNil(t, update.Poll)
			return nil
		})
	})

	require.NoError(t, handleUpdate(t, group, telego.Update{Poll: &telego.Poll{ID: "poll"}}))
	assert.True(t, called)
}

func TestManager_ActiveSession(t *testing.T) {
	storage := NewMemoryStorage()
	manager := NewManager(storage)
	manager.nowFunc = fixedNow
	key := SessionKey{ChatID: 123, UserID: 456}

	info, ok, err := manager.ActiveSession(&th.Context{}, messageUpdate("hello"))
	require.NoError(t, err)
	assert.False(t, ok)
	assert.Nil(t, info)

	state := &SessionState{
		Key:         key,
		FlowID:      "flow",
		CurrentStep: "step",
		Data:        json.RawMessage(`{}`),
		CreatedAt:   fixedNow(),
		UpdatedAt:   fixedNow(),
	}
	require.NoError(t, storage.SaveSession(t.Context(), state))

	info, ok, err = manager.ActiveSession(&th.Context{}, messageUpdate("hello"))
	require.NoError(t, err)
	require.True(t, ok)
	require.NotNil(t, info)
	assert.Equal(t, key, info.Key)
	assert.Equal(t, "flow", info.FlowID)
	assert.Equal(t, "step", info.CurrentStep)

	expiresAt := fixedNow().Add(-time.Second)
	state.ExpiresAt = &expiresAt
	require.NoError(t, storage.SaveSession(t.Context(), state))

	info, ok, err = manager.ActiveSession(&th.Context{}, messageUpdate("hello"))
	require.NoError(t, err)
	assert.False(t, ok)
	assert.Nil(t, info)
}

func TestManager_Cancel(t *testing.T) {
	storage := NewMemoryStorage()
	manager := NewManager(storage)
	cancelCalled := false
	key := SessionKey{ChatID: 123, UserID: 456}

	flow, err := New[testData]("flow").
		Steps(NewStep[testData]("start")).
		OnCancel(func(ctx *Context[testData]) error {
			cancelCalled = true
			assert.Equal(t, "saved", ctx.Data().Name)
			return nil
		}).
		Build()
	require.NoError(t, err)
	require.NoError(t, manager.Register(flow))

	data, err := json.Marshal(testData{Name: "saved"})
	require.NoError(t, err)
	require.NoError(t, storage.SaveSession(t.Context(), &SessionState{
		Key:         key,
		FlowID:      "flow",
		CurrentStep: "start",
		Data:        data,
	}))

	require.NoError(t, manager.Cancel(&th.Context{}, messageUpdate("/cancel")))
	_, _, ok := loadData[testData](t, storage, key)
	assert.False(t, ok)
	assert.True(t, cancelCalled)
}

func TestManager_CancelErrors(t *testing.T) {
	t.Run("no_key", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		err := manager.Cancel(&th.Context{}, telego.Update{Poll: &telego.Poll{ID: "poll"}})

		var keyErr NoSessionKeyError
		require.ErrorAs(t, err, &keyErr)
	})

	t.Run("no_session", func(t *testing.T) {
		manager := NewManager(NewMemoryStorage())
		require.NoError(t, manager.Cancel(&th.Context{}, messageUpdate("/cancel")))
	})

	t.Run("unknown_flow", func(t *testing.T) {
		storage := NewMemoryStorage()
		manager := NewManager(storage)
		require.NoError(t, storage.SaveSession(t.Context(), &SessionState{
			Key:    SessionKey{ChatID: 123, UserID: 456},
			FlowID: "missing",
			Data:   json.RawMessage(`{}`),
		}))

		err := manager.Cancel(&th.Context{}, messageUpdate("/cancel"))

		var flowErr FlowNotFoundError
		require.ErrorAs(t, err, &flowErr)
		assert.Equal(t, "missing", flowErr.FlowID)
	})
}

func TestManager_Timeout(t *testing.T) {
	storage := NewMemoryStorage()
	manager := NewManager(storage)
	manager.nowFunc = fixedNow
	timeoutCalled := false
	key := SessionKey{ChatID: 123, UserID: 456}

	flow, err := New[testData]("flow").
		Steps(NewStep[testData]("start")).
		WithTimeout(time.Minute).
		OnTimeout(func(ctx *Context[testData]) error {
			timeoutCalled = true
			assert.Equal(t, "old", ctx.Data().Name)
			return nil
		}).
		Build()
	require.NoError(t, err)
	require.NoError(t, manager.Register(flow))

	data, err := json.Marshal(testData{Name: "old"})
	require.NoError(t, err)
	expiresAt := fixedNow().Add(-time.Second)
	require.NoError(t, storage.SaveSession(t.Context(), &SessionState{
		Key:         key,
		FlowID:      "flow",
		CurrentStep: "start",
		Data:        data,
		ExpiresAt:   &expiresAt,
	}))

	group := newTestGroup(manager)
	require.NoError(t, handleUpdate(t, group, messageUpdate("hello")))

	_, _, ok := loadData[testData](t, storage, key)
	assert.False(t, ok)
	assert.True(t, timeoutCalled)
}

func TestManager_StorageErrors(t *testing.T) {
	t.Run("middleware_load_error", func(t *testing.T) {
		storage := newPersistentTestStorage()
		storage.loadErr = errTest
		manager := NewManager(storage)
		group := newTestGroup(manager)

		err := handleUpdate(t, group, messageUpdate("hello"))

		require.ErrorIs(t, err, errTest)
	})

	t.Run("start_save_error", func(t *testing.T) {
		storage := newPersistentTestStorage()
		storage.saveErr = errTest
		manager := NewManager(storage)
		flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
		require.NoError(t, err)
		require.NoError(t, manager.Register(flow))

		err = flow.Start(&th.Context{}, messageUpdate("/start"))

		require.ErrorIs(t, err, errTest)
	})

	t.Run("complete_delete_error", func(t *testing.T) {
		storage := newPersistentTestStorage()
		storage.deleteErr = errTest
		manager := NewManager(storage)
		flow, err := New[testData]("flow").
			Steps(NewStep[testData]("start").Enter(func(ctx *Context[testData]) error {
				return ctx.Complete()
			})).
			Build()
		require.NoError(t, err)
		require.NoError(t, manager.Register(flow))

		err = flow.Start(&th.Context{}, messageUpdate("/start"))

		require.ErrorIs(t, err, errTest)
	})
}

func TestManager_PersistentStorageRestoresAcrossManagers(t *testing.T) {
	storage := newPersistentTestStorage()
	key := SessionKey{ChatID: 123, UserID: 456}

	firstManager := NewManager(storage)
	firstFlow, err := New[testData]("flow").
		Steps(
			NewStep[testData]("name").
				Handle(func(ctx *Context[testData]) error {
					ctx.Data().Name = ctx.Text()
					return ctx.Go("count")
				}, th.AnyMessageWithText()).
				CanGo("count"),
			NewStep[testData]("count").
				Enter(func(ctx *Context[testData]) error {
					ctx.Data().Count++
					return nil
				}),
		).
		Build()
	require.NoError(t, err)
	require.NoError(t, firstManager.Register(firstFlow))

	group := newTestGroup(firstManager, func(group *th.HandlerGroup) {
		group.Handle(firstFlow.Start, th.CommandEqual("start"))
	})
	require.NoError(t, handleUpdate(t, group, messageUpdate("/start")))
	require.NoError(t, handleUpdate(t, group, messageUpdate("Alice")))

	state, data, ok := loadData[testData](t, storage, key)
	require.True(t, ok)
	assert.Equal(t, "count", state.CurrentStep)
	assert.Equal(t, "Alice", data.Name)
	assert.Equal(t, 1, data.Count)

	secondManager := NewManager(storage)
	secondEntered := false
	secondFlow, err := New[testData]("flow").
		Steps(
			NewStep[testData]("name").CanGo("count"),
			NewStep[testData]("count").
				Handle(func(ctx *Context[testData]) error {
					secondEntered = true
					assert.Equal(t, "Alice", ctx.Data().Name)
					assert.Equal(t, 1, ctx.Data().Count)
					ctx.Data().Count++
					return ctx.Complete()
				}, th.TextEqual("finish")),
		).
		Build()
	require.NoError(t, err)
	require.NoError(t, secondManager.Register(secondFlow))

	secondGroup := newTestGroup(secondManager)
	require.NoError(t, handleUpdate(t, secondGroup, messageUpdate("finish")))

	_, _, ok = loadData[testData](t, storage, key)
	assert.False(t, ok)
	assert.True(t, secondEntered)
}

func TestManager_ConcurrentUpdatesSameSessionAreSerialized(t *testing.T) {
	storage := NewMemoryStorage()
	manager := NewManager(storage)
	started := make(chan struct{})
	release := make(chan struct{})
	var running atomic.Int32
	var maxRunning atomic.Int32
	var handled atomic.Int32

	flow, err := New[testData]("flow").
		Steps(NewStep[testData]("step").Handle(func(ctx *Context[testData]) error {
			current := running.Add(1)
			for {
				currentMax := maxRunning.Load()
				if current <= currentMax || maxRunning.CompareAndSwap(currentMax, current) {
					break
				}
			}
			if handled.Add(1) == 1 {
				close(started)
				<-release
			}
			ctx.Data().Count++
			running.Add(-1)
			return nil
		})).
		Build()
	require.NoError(t, err)
	require.NoError(t, manager.Register(flow))
	require.NoError(t, flow.Start(&th.Context{}, messageUpdate("/start")))

	group := newTestGroup(manager)
	bot := newTestBot(t)
	firstDone := make(chan error, 1)
	secondDone := make(chan error, 1)

	go func() { firstDone <- group.HandleUpdate(context.Background(), bot, messageUpdate("first")) }()
	<-started
	go func() { secondDone <- group.HandleUpdate(context.Background(), bot, messageUpdate("second")) }()

	select {
	case <-time.After(20 * time.Millisecond):
		assert.Equal(t, int32(1), handled.Load())
		assert.Equal(t, int32(1), maxRunning.Load())
	case err := <-secondDone:
		t.Fatalf("second update finished before first was released: %v", err)
	}

	close(release)
	require.NoError(t, <-firstDone)
	require.NoError(t, <-secondDone)
	assert.Equal(t, int32(2), handled.Load())
	assert.Equal(t, int32(1), maxRunning.Load())

	_, data, ok := loadData[testData](t, storage, SessionKey{ChatID: 123, UserID: 456})
	require.True(t, ok)
	assert.Equal(t, 2, data.Count)
}

func TestManager_ConcurrentUpdatesDifferentSessionsRunInParallel(t *testing.T) {
	storage := NewMemoryStorage()
	manager := NewManager(storage)
	started := make(chan struct{}, 2)
	release := make(chan struct{})
	var running atomic.Int32
	var maxRunning atomic.Int32

	flow, err := New[testData]("flow").
		Steps(NewStep[testData]("step").Handle(func(ctx *Context[testData]) error {
			current := running.Add(1)
			for {
				currentMax := maxRunning.Load()
				if current <= currentMax || maxRunning.CompareAndSwap(currentMax, current) {
					break
				}
			}
			started <- struct{}{}
			<-release
			ctx.Data().Count++
			running.Add(-1)
			return nil
		})).
		Build()
	require.NoError(t, err)
	require.NoError(t, manager.Register(flow))
	require.NoError(t, flow.Start(&th.Context{}, messageUpdateFor(123, 456, "/start")))
	require.NoError(t, flow.Start(&th.Context{}, messageUpdateFor(999, 789, "/start")))

	group := newTestGroup(manager)
	bot := newTestBot(t)
	firstDone := make(chan error, 1)
	secondDone := make(chan error, 1)
	go func() {
		firstDone <- group.HandleUpdate(context.Background(), bot, messageUpdateFor(123, 456, "first"))
	}()
	go func() {
		secondDone <- group.HandleUpdate(context.Background(), bot, messageUpdateFor(999, 789, "second"))
	}()

	<-started
	<-started
	assert.Equal(t, int32(2), maxRunning.Load())
	close(release)
	require.NoError(t, <-firstDone)
	require.NoError(t, <-secondDone)
}

func TestManager_CustomKeyFunc(t *testing.T) {
	storage := NewMemoryStorage()
	manager := NewManager(storage, WithKeyFunc(func(update telego.Update) (SessionKey, bool) {
		if update.Message == nil {
			return SessionKey{}, false
		}
		return SessionKey{ChatID: 1, UserID: 2}, true
	}))
	flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
	require.NoError(t, err)
	require.NoError(t, manager.Register(flow))

	require.NoError(t, flow.Start(&th.Context{}, messageUpdate("/start")))
	_, _, ok := loadData[testData](t, storage, SessionKey{ChatID: 1, UserID: 2})
	assert.True(t, ok)
}

func TestManager_HandleUnknownFlowAndBadData(t *testing.T) {
	t.Run("unknown_flow", func(t *testing.T) {
		storage := NewMemoryStorage()
		manager := NewManager(storage)
		require.NoError(t, storage.SaveSession(t.Context(), &SessionState{
			Key:    SessionKey{ChatID: 123, UserID: 456},
			FlowID: "missing",
			Data:   json.RawMessage(`{}`),
		}))

		err := handleUpdate(t, newTestGroup(manager), messageUpdate("hello"))

		var flowErr FlowNotFoundError
		require.ErrorAs(t, err, &flowErr)
		assert.Equal(t, "missing", flowErr.FlowID)
	})

	t.Run("bad_json", func(t *testing.T) {
		storage := NewMemoryStorage()
		manager := NewManager(storage)
		flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
		require.NoError(t, err)
		require.NoError(t, manager.Register(flow))
		require.NoError(t, storage.SaveSession(t.Context(), &SessionState{
			Key:         SessionKey{ChatID: 123, UserID: 456},
			FlowID:      "flow",
			CurrentStep: "start",
			Data:        json.RawMessage(`{bad`),
		}))

		err = handleUpdate(t, newTestGroup(manager), messageUpdate("hello"))

		var dataErr SessionDataError
		require.ErrorAs(t, err, &dataErr)
		assert.Equal(t, "flow", dataErr.FlowID)
	})
}

func TestManager_ContextCancellationPassedToStorage(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	storage := NewMemoryStorage()
	manager := NewManager(storage)
	flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
	require.NoError(t, err)
	require.NoError(t, manager.Register(flow))

	base := (&th.Context{}).WithContext(ctx)
	err = flow.Start(base, messageUpdate("/start"))

	// MemoryStorage doesn't reject canceled contexts, but this verifies the public
	// path accepts a telegohandler context carrying arbitrary cancellation state.
	require.NoError(t, err)
}

func TestManager_ErrorTypeFormatting(t *testing.T) {
	assert.Equal(t, "telego: flow session key not found in update", (NoSessionKeyError{}).Error())
	assert.Contains(t, (FlowNotFoundError{FlowID: "f"}).Error(), "f")
	assert.Contains(t, (DuplicateFlowError{FlowID: "f"}).Error(), "f")
	assert.Contains(t, (FlowNotRegisteredError{FlowID: "f"}).Error(), "f")
	assert.Contains(t, (StepNotFoundError{FlowID: "f", StepID: "s"}).Error(), "s")
	assert.Contains(t, (DuplicateStepError{StepID: "s"}).Error(), "s")
	assert.Contains(t, (InvalidStartStepError{FlowID: "f", StepID: "s"}).Error(), "s")
	assert.Contains(t, (InvalidTransitionError{FlowID: "f", From: "a", To: "b"}).Error(), "b")
	assert.Contains(t, (TransitionNotAllowedError{FlowID: "f", From: "a", To: "b"}).Error(), "not allowed")
	assert.Contains(t, (SessionDataError{FlowID: "f", Err: errTest}).Error(), "test error")
	assert.ErrorIs(t, (SessionDataError{FlowID: "f", Err: errTest}).Unwrap(), errTest)
}
