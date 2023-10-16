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
			args:     nil,
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
			text:     "/test abc 123",
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
