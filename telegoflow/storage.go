package telegoflow

import "context"

// Storage persists flow sessions.
//
// Storage implementations don't need to know the concrete generic data type of
// a flow. Flow data is stored as JSON in [SessionState.Data].
type Storage interface {
	// LoadSession loads a session by key. The bool result is false when no session exists.
	LoadSession(ctx context.Context, key SessionKey) (*SessionState, bool, error)

	// SaveSession creates or replaces a session.
	SaveSession(ctx context.Context, session *SessionState) error

	// DeleteSession removes a session by key.
	DeleteSession(ctx context.Context, key SessionKey) error
}
