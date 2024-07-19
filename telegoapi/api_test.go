package telegoapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chococola/telego/internal/json"
)

func Test_Response_String_and_Error_Error(t *testing.T) {
	tests := []struct {
		name string
		resp Response
		text string
	}{
		{
			name: "success",
			resp: Response{
				Ok:     true,
				Error:  nil,
				Result: json.RawMessage(`{"test":true}`),
			},
			text: `Ok: true, Err: [<nil>], Result: {"test":true}`,
		},
		{
			name: "error",
			resp: Response{
				Ok: false,
				Error: &Error{
					Description: "bad request",
					ErrorCode:   400,
					Parameters:  nil,
				},
				Result: nil,
			},
			text: "Ok: false, Err: [400 \"bad request\"]",
		},
		{
			name: "error_with_parameters",
			resp: Response{
				Ok: false,
				Error: &Error{
					Description: "bad request",
					ErrorCode:   400,
					Parameters: &ResponseParameters{
						MigrateToChatID: 1,
						RetryAfter:      2,
					},
				},
				Result: nil,
			},
			text: "Ok: false, Err: [400 \"bad request\", migrate to chat ID: 1, retry after: 2]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			text := tt.resp.String()
			assert.Equal(t, tt.text, text)
		})
	}
}

func Test_Error(t *testing.T) {
	var err error = &Error{
		ErrorCode: 1,
	}

	var apiErr *Error
	if errors.As(err, &apiErr) {
		assert.Equal(t, 1, apiErr.ErrorCode)
	} else {
		assert.Fail(t, "not an API error")
	}
}
