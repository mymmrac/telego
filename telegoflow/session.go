package telegoflow

import (
	"encoding/json"
	"fmt"
	"time"
)

// SessionKey identifies a single conversation session.
//
// The default key uses both chat ID and user ID, so the same user can have
// independent conversations in private chats and groups.
type SessionKey struct {
	ChatID int64 `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

// String returns a stable textual representation of the session key.
func (k SessionKey) String() string {
	return fmt.Sprintf("%d:%d", k.ChatID, k.UserID)
}

// SessionState contains persistent session metadata and encoded typed data.
type SessionState struct {
	Key         SessionKey      `json:"key"`
	FlowID      string          `json:"flow_id"`
	CurrentStep string          `json:"current_step"`
	Data        json.RawMessage `json:"data"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	ExpiresAt   *time.Time      `json:"expires_at,omitempty"`
}

// Expired reports whether the session has passed its expiration time.
func (s *SessionState) Expired(now time.Time) bool {
	return s != nil && s.ExpiresAt != nil && now.After(*s.ExpiresAt)
}

// SessionInfo contains non-generic information about an active session.
type SessionInfo struct {
	Key         SessionKey
	FlowID      string
	CurrentStep string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ExpiresAt   *time.Time
}

func newSessionInfo(state *SessionState) *SessionInfo {
	if state == nil {
		return nil
	}

	return &SessionInfo{
		Key:         state.Key,
		FlowID:      state.FlowID,
		CurrentStep: state.CurrentStep,
		CreatedAt:   state.CreatedAt,
		UpdatedAt:   state.UpdatedAt,
		ExpiresAt:   state.ExpiresAt,
	}
}
