//go:generate mockgen -package mock -destination=mock/caller.go github.com/mymmrac/telego/api Caller
//go:generate mockgen -package mock -destination=mock/request-constructor.go github.com/mymmrac/telego/api RequestConstructor

package api

import (
	"bytes"
	stdJson "encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"reflect"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

const (
	// ContentTypeJSON http content type
	ContentTypeJSON = "application/json"
)

// json jsoniter replacement for json package
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Response represents response returned by Telegram API
type Response struct {
	Ok     bool               `json:"ok"`
	Result stdJson.RawMessage `json:"result,omitempty"`
	*Error
}

func (r Response) String() string {
	if r.Result == nil || len(r.Result) == 0 {
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
func (a Error) Error() string {
	if a.Parameters != nil {
		return fmt.Sprintf("%d %q migrate to chat id: %d, retry after: %d",
			a.ErrorCode, a.Description, a.Parameters.MigrateToChatID, a.Parameters.RetryAfter)
	}
	return fmt.Sprintf("%d %q", a.ErrorCode, a.Description)
}

// ResponseParameters - Contains information about why a request was unsuccessful.
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

// RequestData represents data needed to execute request
type RequestData struct {
	ContentType string
	Buffer      *bytes.Buffer
}

// Caller represents way to call API with request
type Caller interface {
	Call(url string, data *RequestData) (*Response, error)
}

// NamedReader represents a way to send files (or other data).
// Implemented by os.File.
type NamedReader interface {
	io.Reader
	Name() string
}

// RequestConstructor represents way to construct API request
type RequestConstructor interface {
	JSONRequest(parameters interface{}) (*RequestData, error)
	MultipartRequest(parameters map[string]string, filesParameters map[string]NamedReader) (*RequestData, error)
}

// FasthttpAPICaller fasthttp implementation of Caller
type FasthttpAPICaller struct {
	Client *fasthttp.Client
}

// Call is fasthttp implementation
func (a FasthttpAPICaller) Call(url string, data *RequestData) (*Response, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetContentType(data.ContentType)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBodyRaw(data.Buffer.Bytes())

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := a.Client.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("fasthttp do request: %w", err)
	}

	if statusCode := resp.StatusCode(); statusCode >= fasthttp.StatusInternalServerError {
		return nil, fmt.Errorf("internal server error: %d", statusCode)
	}

	apiResp := &Response{}
	err = json.Unmarshal(resp.Body(), apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	return apiResp, nil
}

// DefaultConstructor default implementation of RequestConstructor
type DefaultConstructor struct{}

// JSONRequest is default implementation
func (d DefaultConstructor) JSONRequest(parameters interface{}) (*RequestData, error) {
	data := &RequestData{
		ContentType: ContentTypeJSON,
		Buffer:      &bytes.Buffer{},
	}

	err := json.NewEncoder(data.Buffer).Encode(parameters)
	if err != nil {
		return nil, fmt.Errorf("encode json: %w", err)
	}

	return data, nil
}

// MultipartRequest is default implementation
func (d DefaultConstructor) MultipartRequest(
	parameters map[string]string, filesParameters map[string]NamedReader) (*RequestData, error) {
	data := &RequestData{
		Buffer: &bytes.Buffer{},
	}
	writer := multipart.NewWriter(data.Buffer)

	for field, file := range filesParameters {
		if isNil(file) {
			continue
		}

		wr, err := writer.CreateFormFile(field, file.Name())
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(wr, file)
		if err != nil {
			return nil, err
		}
	}

	for field, value := range parameters {
		wr, err := writer.CreateFormField(field)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(wr, strings.NewReader(value))
		if err != nil {
			return nil, err
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, fmt.Errorf("closing writer: %w", err)
	}

	data.ContentType = writer.FormDataContentType()
	return data, nil
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}

	return false
}
