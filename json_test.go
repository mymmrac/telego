//go:build !race

package telego

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego/internal/json"
)

func TestSetMarshal(t *testing.T) {
	tests := []struct {
		name    string
		marshal func(v any) ([]byte, error)
		isValid bool
	}{
		{
			name:    "nil",
			marshal: nil,
			isValid: false,
		},
		{
			name:    "ok",
			marshal: json.Marshal,
			isValid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isValid {
				assert.NotPanics(t, func() {
					SetJSONMarshal(tt.marshal)
				})
			} else {
				assert.Panics(t, func() {
					SetJSONMarshal(tt.marshal)
				})
			}
		})
	}
}

func TestSetUnmarshal(t *testing.T) {
	tests := []struct {
		name      string
		unmarshal func(data []byte, v any) error
		isValid   bool
	}{
		{
			name:      "nil",
			unmarshal: nil,
			isValid:   false,
		},
		{
			name:      "ok",
			unmarshal: json.Unmarshal,
			isValid:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isValid {
				assert.NotPanics(t, func() {
					SetJSONUnmarshal(tt.unmarshal)
				})
			} else {
				assert.Panics(t, func() {
					SetJSONUnmarshal(tt.unmarshal)
				})
			}
		})
	}
}
