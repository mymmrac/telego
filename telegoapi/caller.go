//go:generate mockgen -typed -package mock -destination=mock/caller.go github.com/mymmrac/telego/telegoapi Caller

package telegoapi

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

// FastHTTPCaller fasthttp implementation of Caller
type FastHTTPCaller struct {
	Client *fasthttp.Client
}

// DefaultFastHTTPCaller is a default fast http caller
var DefaultFastHTTPCaller = &FastHTTPCaller{
	Client: &fasthttp.Client{},
}

// Call is a fasthttp implementation
func (a FastHTTPCaller) Call(url string, data *RequestData) (*Response, error) {
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

// HTTPCaller http implementation of Caller
type HTTPCaller struct {
	Client *http.Client
}

// DefaultHTTPCaller is a default http caller
var DefaultHTTPCaller = &HTTPCaller{
	Client: http.DefaultClient,
}

// Call is a http implementation
func (h HTTPCaller) Call(url string, data *RequestData) (*Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, data.Buffer)
	if err != nil {
		return nil, fmt.Errorf("http create request: %w", err)
	}
	req.Header.Set(ContentTypeHeader, data.ContentType)

	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http do request: %w", err)
	}

	if resp.StatusCode >= http.StatusInternalServerError {
		return nil, fmt.Errorf("internal server error: %d", resp.StatusCode)
	}

	apiResp := &Response{}
	err = json.NewDecoder(resp.Body).Decode(apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	if err = resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("close body: %w", err)
	}

	return apiResp, nil
}

// RetryCaller decorator over Caller that provides reties with exponential backoff
// Delay = (ExponentBase ^ AttemptNumber) * StartDelay or MaxDelay
type RetryCaller struct {
	Caller       Caller
	MaxAttempts  int
	ExponentBase float64
	StartDelay   time.Duration
	MaxDelay     time.Duration
}

// ErrMaxRetryAttempts returned when max retry attempts reached
var ErrMaxRetryAttempts = errors.New("max retry attempts reached")

// Call makes calls using provided caller with retries
func (r *RetryCaller) Call(url string, data *RequestData) (resp *Response, err error) {
	for i := 0; i < r.MaxAttempts; i++ {
		resp, err = r.Caller.Call(url, data)
		if err == nil {
			return resp, nil
		}

		if i == r.MaxAttempts-1 {
			break
		}

		delay := time.Duration(math.Pow(r.ExponentBase, float64(i))) * r.StartDelay
		if delay > r.MaxDelay {
			delay = r.MaxDelay
		}
		time.Sleep(delay)
	}
	return nil, errors.Join(err, ErrMaxRetryAttempts)
}
