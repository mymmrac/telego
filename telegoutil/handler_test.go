package telegoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name        string
		text        string
		commandName string
		args        []string
	}{
		{
			name:        "no_command",
			text:        "test",
			commandName: "",
			args:        nil,
		},
		{
			name:        "without_args",
			text:        "/test",
			commandName: "test",
			args:        []string{},
		},
		{
			name:        "with_args",
			text:        "/test abc 123",
			commandName: "test",
			args:        []string{"abc", "123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, args := ParseCommand(tt.text)
			assert.Equal(t, tt.commandName, name)
			assert.Equal(t, tt.args, args)
		})
	}
}
