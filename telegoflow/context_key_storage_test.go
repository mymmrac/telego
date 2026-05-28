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

func TestContext_AccessorsAndCallback(t *testing.T) {
	storage := NewMemoryStorage()
	manager := NewManager(storage)
	var sawCallback bool
	var sawUpdate bool

	flow, err := New[testData]("flow").
		Steps(NewStep[testData]("callback").Handle(func(ctx *Context[testData]) error {
			sawUpdate = ctx.Update().CallbackQuery != nil
			callback := ctx.CallbackQuery()
			require.NotNil(t, callback)
			assert.Equal(t, "payload", callback.Data)
			assert.NotNil(t, ctx.Message())
			assert.Empty(t, ctx.Text())
			sawCallback = true
			return ctx.Complete()
		}, th.AnyCallbackQuery())).
		Build()
	require.NoError(t, err)
	require.NoError(t, manager.Register(flow))
	require.NoError(t, storage.SaveSession(t.Context(), &SessionState{
		Key:         SessionKey{ChatID: 123, UserID: 456},
		FlowID:      "flow",
		CurrentStep: "callback",
		Data:        json.RawMessage(`{}`),
	}))

	require.NoError(t, handleUpdate(t, newTestGroup(manager), callbackUpdate("payload")))
	assert.True(t, sawUpdate)
	assert.True(t, sawCallback)
}

func TestContext_MessageVariants(t *testing.T) {
	baseMessage := func(text string) *telego.Message {
		return &telego.Message{
			MessageID: 10,
			From:      &telego.User{ID: 456},
			Chat:      telego.Chat{ID: 123},
			Text:      text,
		}
	}

	tests := []struct {
		name   string
		update telego.Update
		text   string
	}{
		{name: "message", update: telego.Update{Message: baseMessage("message")}, text: "message"},
		{name: "edited_message", update: telego.Update{EditedMessage: baseMessage("edited")}, text: "edited"},
		{
			name:   "business_message",
			update: telego.Update{BusinessMessage: baseMessage("business")},
			text:   "business",
		},
		{
			name:   "edited_business_message",
			update: telego.Update{EditedBusinessMessage: baseMessage("edited_business")},
			text:   "edited_business",
		},
		{name: "guest_message", update: telego.Update{GuestMessage: baseMessage("guest")}, text: "guest"},
		{name: "callback_accessible_message", update: callbackUpdate("data"), text: ""},
		{name: "no_message", update: telego.Update{Poll: &telego.Poll{ID: "poll"}}, text: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context[testData]{update: tt.update}
			msg := ctx.Message()
			if tt.text == "" {
				assert.Empty(t, ctx.Text())
				return
			}
			require.NotNil(t, msg)
			assert.Equal(t, tt.text, ctx.Text())
		})
	}
}

func TestContext_GoErrors(t *testing.T) {
	t.Run("current_step_missing", func(t *testing.T) {
		flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
		require.NoError(t, err)
		ctx := &Context[testData]{flow: flow, session: &SessionState{CurrentStep: "missing"}}

		err = ctx.Go("start")

		var stepErr StepNotFoundError
		require.ErrorAs(t, err, &stepErr)
		assert.Equal(t, "missing", stepErr.StepID)
	})

	t.Run("transition_not_allowed", func(t *testing.T) {
		flow, err := New[testData]("flow").Steps(NewStep[testData]("start"), NewStep[testData]("next")).Build()
		require.NoError(t, err)
		ctx := &Context[testData]{flow: flow, session: &SessionState{CurrentStep: "start"}}

		err = ctx.Go("next")

		var transitionErr TransitionNotAllowedError
		require.ErrorAs(t, err, &transitionErr)
		assert.Equal(t, "start", transitionErr.From)
		assert.Equal(t, "next", transitionErr.To)
	})
}

func TestFlow_LifecycleErrorHooksAndMiddleware(t *testing.T) {
	storage := NewMemoryStorage()
	manager := NewManager(storage)
	var order []string

	flow, err := New[testData]("flow").
		Steps(NewStep[testData]("start").
			Use(func(next Handler[testData]) Handler[testData] {
				return func(ctx *Context[testData]) error {
					order = append(order, "step_before")
					err := next(ctx)
					order = append(order, "step_after")
					return err
				}
			}).
			Enter(func(ctx *Context[testData]) error {
				order = append(order, "enter")
				return errTest
			})).
		Use(func(next Handler[testData]) Handler[testData] {
			return func(ctx *Context[testData]) error {
				order = append(order, "flow_before")
				err := next(ctx)
				order = append(order, "flow_after")
				return err
			}
		}).
		OnError(func(ctx *Context[testData], err error) error {
			order = append(order, "on_error")
			require.ErrorIs(t, err, errTest)
			return nil
		}).
		Build()
	require.NoError(t, err)
	require.NoError(t, manager.Register(flow))

	require.NoError(t, flow.Start(&th.Context{}, messageUpdate("/start")))
	assert.Equal(t, []string{"flow_before", "step_before", "enter", "step_after", "flow_after", "on_error"}, order)
}

func TestFlow_StartWithData(t *testing.T) {
	storage := NewMemoryStorage()
	manager := NewManager(storage)
	flow, err := New[testData]("flow").Steps(NewStep[testData]("start")).Build()
	require.NoError(t, err)
	require.NoError(t, manager.Register(flow))

	handler := flow.StartWithData(testData{Name: "initial", Count: 7})
	require.NoError(t, handler(&th.Context{}, messageUpdate("/start")))

	_, data, ok := loadData[testData](t, storage, SessionKey{ChatID: 123, UserID: 456})
	require.True(t, ok)
	assert.Equal(t, testData{Name: "initial", Count: 7}, data)
}

func TestDefaultKeyFunc(t *testing.T) {
	msg := func(chatID, userID int64) *telego.Message {
		return &telego.Message{From: &telego.User{ID: userID}, Chat: telego.Chat{ID: chatID}}
	}

	tests := []struct {
		name   string
		update telego.Update
		key    SessionKey
		ok     bool
	}{
		{name: "message", update: telego.Update{Message: msg(1, 2)}, key: SessionKey{ChatID: 1, UserID: 2}, ok: true},
		{
			name:   "edited",
			update: telego.Update{EditedMessage: msg(3, 4)},
			key:    SessionKey{ChatID: 3, UserID: 4},
			ok:     true,
		},
		{
			name:   "business",
			update: telego.Update{BusinessMessage: msg(5, 6)},
			key:    SessionKey{ChatID: 5, UserID: 6},
			ok:     true,
		},
		{
			name:   "edited_business",
			update: telego.Update{EditedBusinessMessage: msg(7, 8)},
			key:    SessionKey{ChatID: 7, UserID: 8},
			ok:     true,
		},
		{
			name:   "guest",
			update: telego.Update{GuestMessage: msg(9, 10)},
			key:    SessionKey{ChatID: 9, UserID: 10},
			ok:     true,
		},
		{name: "callback", update: callbackUpdate("data"), key: SessionKey{ChatID: 123, UserID: 456}, ok: true},
		{name: "message_no_from", update: telego.Update{Message: &telego.Message{Chat: telego.Chat{ID: 1}}}},
		{
			name:   "callback_no_message",
			update: telego.Update{CallbackQuery: &telego.CallbackQuery{From: telego.User{ID: 1}}},
		},
		{name: "none", update: telego.Update{Poll: &telego.Poll{ID: "poll"}}, ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, ok := DefaultKeyFunc(tt.update)
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.key, key)
		})
	}
}

func TestMemoryStorageCloneIsolationAndHelpers(t *testing.T) {
	storage := NewMemoryStorage()
	key := SessionKey{ChatID: 1, UserID: 2}
	expiresAt := fixedNow().Add(time.Minute)
	state := &SessionState{
		Key:         key,
		FlowID:      "flow",
		CurrentStep: "step",
		Data:        json.RawMessage(`{"name":"original"}`),
		CreatedAt:   fixedNow(),
		UpdatedAt:   fixedNow(),
		ExpiresAt:   &expiresAt,
	}
	require.NoError(t, storage.SaveSession(t.Context(), state))

	state.Data[9] = 'X'
	*state.ExpiresAt = fixedNow().Add(-time.Hour)
	loaded, ok, err := storage.LoadSession(t.Context(), key)
	require.NoError(t, err)
	require.True(t, ok)
	assert.JSONEq(t, `{"name":"original"}`, string(loaded.Data))
	assert.True(t, loaded.ExpiresAt.After(fixedNow()))

	loaded.Data[9] = 'Y'
	loadedAgain, ok, err := storage.LoadSession(t.Context(), key)
	require.NoError(t, err)
	require.True(t, ok)
	assert.JSONEq(t, `{"name":"original"}`, string(loadedAgain.Data))

	assert.Equal(t, "1:2", key.String())
	assert.False(t, (*SessionState)(nil).Expired(fixedNow()))
	assert.Nil(t, newSessionInfo(nil))
	assert.Nil(t, cloneSession(nil))

	require.NoError(t, storage.DeleteSession(t.Context(), key))
	_, ok, err = storage.LoadSession(t.Context(), key)
	require.NoError(t, err)
	assert.False(t, ok)
}
