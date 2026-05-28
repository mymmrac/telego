package telegoflow

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

const testToken = "1234567890:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc"

var errTest = errors.New("test error")

type testData struct {
	Name  string `json:"name,omitempty"`
	Count int    `json:"count,omitempty"`
}

type persistentTestStorage struct {
	mu        sync.RWMutex
	sessions  map[SessionKey][]byte
	loadErr   error
	saveErr   error
	deleteErr error
}

func newPersistentTestStorage() *persistentTestStorage {
	return &persistentTestStorage{sessions: make(map[SessionKey][]byte)}
}

func (s *persistentTestStorage) LoadSession(_ context.Context, key SessionKey) (*SessionState, bool, error) {
	if s.loadErr != nil {
		return nil, false, s.loadErr
	}

	s.mu.RLock()
	data, ok := s.sessions[key]
	s.mu.RUnlock()
	if !ok {
		return nil, false, nil
	}

	var session SessionState
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, false, err
	}
	return &session, true, nil
}

func (s *persistentTestStorage) SaveSession(_ context.Context, session *SessionState) error {
	if s.saveErr != nil {
		return s.saveErr
	}

	data, err := json.Marshal(session)
	if err != nil {
		return err
	}

	s.mu.Lock()
	s.sessions[session.Key] = data
	s.mu.Unlock()
	return nil
}

func (s *persistentTestStorage) DeleteSession(_ context.Context, key SessionKey) error {
	if s.deleteErr != nil {
		return s.deleteErr
	}

	s.mu.Lock()
	delete(s.sessions, key)
	s.mu.Unlock()
	return nil
}

func newTestBot(t *testing.T) *telego.Bot {
	t.Helper()

	bot, err := telego.NewBot(testToken)
	require.NoError(t, err)
	return bot
}

func newTestGroup(manager *Manager, handlers ...func(group *th.HandlerGroup)) *th.HandlerGroup {
	group := &th.HandlerGroup{}
	group.Use(manager.Middleware())
	for _, handler := range handlers {
		handler(group)
	}
	return group
}

func handleUpdate(t *testing.T, group *th.HandlerGroup, update telego.Update) error {
	t.Helper()
	return group.HandleUpdate(t.Context(), newTestBot(t), update)
}

func messageUpdate(text string) telego.Update {
	return messageUpdateFor(123, 456, text)
}

func messageUpdateFor(chatID int64, userID int64, text string) telego.Update {
	return telego.Update{
		UpdateID: int(userID),
		Message: &telego.Message{
			MessageID: 10,
			From:      &telego.User{ID: userID, FirstName: "Test"},
			Chat:      telego.Chat{ID: chatID, Type: telego.ChatTypePrivate},
			Text:      text,
		},
	}
}

func callbackUpdate(data string) telego.Update {
	return telego.Update{
		UpdateID: 2,
		CallbackQuery: &telego.CallbackQuery{
			ID:   "callback-id",
			From: telego.User{ID: 456, FirstName: "Test"},
			Message: &telego.Message{
				MessageID: 10,
				Chat:      telego.Chat{ID: 123, Type: telego.ChatTypePrivate},
			},
			Data: data,
		},
	}
}

func loadData[T any](t *testing.T, storage Storage, key SessionKey) (*SessionState, T, bool) {
	t.Helper()

	state, ok, err := storage.LoadSession(t.Context(), key)
	require.NoError(t, err)
	var data T
	if !ok {
		return nil, data, false
	}
	require.NoError(t, json.Unmarshal(state.Data, &data))
	return state, data, true
}

func fixedNow() time.Time {
	return time.Date(2026, 5, 28, 12, 0, 0, 0, time.UTC)
}
