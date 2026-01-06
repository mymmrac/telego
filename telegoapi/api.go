package telegoapi

import (
	"context"
	"fmt"
	"io"

	"github.com/mymmrac/telego/internal/json"
)

const (
	// ContentTypeHeader http content type header
	ContentTypeHeader = "Content-Type"

	// ContentTypeJSON http JSON content type
	ContentTypeJSON = "application/json"
)

// Response represents response returned by Telegram API
type Response struct {
	*Error

	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result,omitempty"`
}

func (r Response) String() string {
	if len(r.Result) == 0 {
		return fmt.Sprintf("Ok: %t, Err: [%v]", r.Ok, r.Error)
	}
	return fmt.Sprintf("Ok: %t, Err: [%v], Result: %s", r.Ok, r.Error, r.Result)
}

// Error represents error from telegram API
type Error struct {
	Description string              `json:"description,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

// Error converts Error to human-readable string
func (a *Error) Error() string {
	if a.Parameters != nil {
		return fmt.Sprintf("%d %q, migrate to chat ID: %d, retry after: %d",
			a.ErrorCode, a.Description, a.Parameters.MigrateToChatID, a.Parameters.RetryAfter)
	}
	return fmt.Sprintf("%d %q", a.ErrorCode, a.Description)
}

// ResponseParameters - Describes why a request was unsuccessful.
type ResponseParameters struct {
	// MigrateToChatID - Optional. The group has been migrated to a supergroup with the specified identifier.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// RetryAfter - Optional. In case of exceeding flood control, the number of seconds left to wait before the
	// request can be repeated
	RetryAfter int `json:"retry_after,omitempty"`
}

// RequestData represents data needed to execute the request; at least one BodyRaw or BodyStream must be provided
type RequestData struct {
	// ContentType request content type
	ContentType string
	// BodyRaw raw body, slice must not be modified until the request is completed
	BodyRaw []byte
	// BodyStream body stream that will be read, ignored if BodyRaw is provided
	BodyStream io.Reader
}

// Caller represents way to call API with a request
type Caller interface {
	// Call accepts the URL and request data (as raw bytes or stream) and calls API
	Call(ctx context.Context, url string, data *RequestData) (*Response, error)
}

// NamedReader represents a way to send files (or other data).
// Implemented by [os.File].
// Note: Name method may be called multiple times and should return unique names for all files sent in one request.
//
// Warning: Since, for sending data (files) reader data will be copied, using the same reader multiple times as is
// will not work.
// For [os.File] you can use os.File.Seek(0, io.SeekStart) to prepare for a new request.
type NamedReader interface {
	// Reader of file data
	io.Reader
	// Name returns file name
	Name() string
}

// RequestConstructor represents a way to construct an API request
type RequestConstructor interface {
	// JSONRequest marshals parameters as JSON, usually returns body as raw bytes
	JSONRequest(parameters any) (*RequestData, error)
	// MultipartRequest marshals parameters and files as multipart form data, usually returns body as stream
	MultipartRequest(parameters map[string]string, filesParameters map[string]NamedReader) (*RequestData, error)
}
