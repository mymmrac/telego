package telegoflow

import (
	"context"
	"sync"
)

// MemoryStorage stores sessions in memory.
//
// It is safe for concurrent use and useful for tests, examples, and bots that
// don't need to survive process restarts.
type MemoryStorage struct {
	mu       sync.RWMutex
	sessions map[SessionKey]*SessionState
}

// NewMemoryStorage creates an in-memory session storage.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		sessions: make(map[SessionKey]*SessionState),
	}
}

// LoadSession loads a session by key.
func (s *MemoryStorage) LoadSession(_ context.Context, key SessionKey) (*SessionState, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, ok := s.sessions[key]
	if !ok {
		return nil, false, nil
	}

	return cloneSession(session), true, nil
}

// SaveSession creates or replaces a session.
func (s *MemoryStorage) SaveSession(_ context.Context, session *SessionState) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[session.Key] = cloneSession(session)
	return nil
}

// DeleteSession removes a session by key.
func (s *MemoryStorage) DeleteSession(_ context.Context, key SessionKey) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.sessions, key)
	return nil
}

func cloneSession(session *SessionState) *SessionState {
	if session == nil {
		return nil
	}

	clone := *session
	if session.Data != nil {
		clone.Data = append([]byte(nil), session.Data...)
	}
	if session.ExpiresAt != nil {
		expiresAt := *session.ExpiresAt
		clone.ExpiresAt = &expiresAt
	}

	return &clone
}
