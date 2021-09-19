package telego

import (
	"bytes"
	stdJson "encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

// apiResponse - Represents response returned by Telegram API
type apiResponse struct {
	Ok     bool               `json:"ok"`
	Result stdJson.RawMessage `json:"result,omitempty"`
	*APIError
}

func (a apiResponse) String() string {
	return fmt.Sprintf("Ok: %t, Err: {%v}, Result: %s", a.Ok, a.APIError, a.Result)
}

// APIError - Represents error from telegram API
type APIError struct {
	Description string              `json:"description,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

// Error - Converts APIError to human-readable string
func (a APIError) Error() string {
	if a.Parameters != nil {
		return fmt.Sprintf("%d %q migrate to chat id: %d, retry after: %d",
			a.ErrorCode, a.Description, a.Parameters.MigrateToChatID, a.Parameters.RetryAfter)
	}
	return fmt.Sprintf("%d %q", a.ErrorCode, a.Description)
}

// requestData - Represents data needed to execute request
type requestData struct {
	ContentType string
	Buffer      *bytes.Buffer
}

// apiCaller - Represents way to call API with request
type apiCaller interface {
	Call(url string, data *requestData) (*apiResponse, error)
}

// requestConstructor - Represents way to construct API request
type requestConstructor interface {
	JSONRequest(parameters interface{}) (*requestData, error)
	MultipartRequest(parameters map[string]string, filesParameters map[string]*os.File) (*requestData, error)
}

// fasthttpAPICaller - Fasthttp implementation of apiCaller
type fasthttpAPICaller struct {
	Client *fasthttp.Client
}

func (a fasthttpAPICaller) Call(url string, data *requestData) (*apiResponse, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetContentType(data.ContentType)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBodyRaw(data.Buffer.Bytes())

	resp := fasthttp.AcquireResponse()
	err := a.Client.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("fasthttp do request: %w", err)
	}

	if statusCode := resp.StatusCode(); statusCode >= fasthttp.StatusInternalServerError {
		return nil, fmt.Errorf("internal server error: %d", statusCode)
	}

	apiResp := &apiResponse{}
	err = json.Unmarshal(resp.Body(), apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	return apiResp, nil
}

// defaultConstructor - Default implementation of requestConstructor
type defaultConstructor struct{}

func (d defaultConstructor) JSONRequest(parameters interface{}) (*requestData, error) {
	data := &requestData{
		ContentType: contentTypeJSON,
		Buffer:      &bytes.Buffer{},
	}

	err := json.NewEncoder(data.Buffer).Encode(parameters)
	if err != nil {
		return nil, fmt.Errorf("encode json: %w", err)
	}

	return data, nil
}

func (d defaultConstructor) MultipartRequest(
	parameters map[string]string, filesParameters map[string]*os.File) (*requestData, error) {
	data := &requestData{
		Buffer: &bytes.Buffer{},
	}
	writer := multipart.NewWriter(data.Buffer)

	for field, file := range filesParameters {
		if file == nil {
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
