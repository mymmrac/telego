package telegoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		cmd      string
		username string
		args     []string
	}{
		{
			name:     "no_command",
			text:     "test",
			cmd:      "",
			username: "",
			args:     []string{},
		},
		{
			name:     "without_args",
			text:     "/test",
			cmd:      "test",
			username: "",
			args:     []string{},
		},
		{
			name:     "with_args",
			text:     "/test abc   123",
			cmd:      "test",
			username: "",
			args:     []string{"abc", "123"},
		},
		{
			name:     "with_username",
			text:     "/test@user ok",
			cmd:      "test",
			username: "@user",
			args:     []string{"ok"},
		},
		{
			name:     "multiline",
			text:     "/test@user  ok\n   test ",
			cmd:      "test",
			username: "@user",
			args:     []string{"ok", "test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, username, args := ParseCommand(tt.text)
			assert.Equal(t, tt.cmd, name)
			assert.Equal(t, tt.username, username)
			assert.Equal(t, tt.args, args)
		})
	}
}

func TestParseCommandPayload(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		cmd      string
		username string
		payload  string
	}{
		{
			name:     "no_command",
			text:     "test",
			cmd:      "",
			username: "",
			payload:  "",
		},
		{
			name:     "without_payload",
			text:     "/test",
			cmd:      "test",
			username: "",
			payload:  "",
		},
		{
			name:     "with_payload",
			text:     "/test abc   123",
			cmd:      "test",
			username: "",
			payload:  "abc   123",
		},
		{
			name:     "with_username",
			text:     "/test@user ok",
			cmd:      "test",
			username: "@user",
			payload:  "ok",
		},
		{
			name:     "multiline",
			text:     "/test@user   ok\n   test ",
			cmd:      "test",
			username: "@user",
			payload:  "ok\n   test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, username, payload := ParseCommandPayload(tt.text)
			assert.Equal(t, tt.cmd, name)
			assert.Equal(t, tt.username, username)
			assert.Equal(t, tt.payload, payload)
		})
	}
}
