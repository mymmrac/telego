//go:generate mockgen -typed -package mock -destination=mock/caller.go github.com/mymmrac/telego/telegoapi Caller

package telegoapi

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/internal/json"
)

// FastHTTPCaller fasthttp implementation of [Caller]
type FastHTTPCaller struct {
	Client *fasthttp.Client
}

// DefaultFastHTTPCaller is a default fasthttp caller
var DefaultFastHTTPCaller = &FastHTTPCaller{
	Client: &fasthttp.Client{},
}

// Call is a fasthttp implementation
func (a FastHTTPCaller) Call(ctx context.Context, url string, data *RequestData) (*Response, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		// Continue
	}

	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	request.SetRequestURI(url)
	request.Header.SetContentType(data.ContentType)
	request.Header.SetMethod(fasthttp.MethodPost)

	switch {
	case data.BodyRaw != nil:
		request.SetBodyRaw(data.BodyRaw)
	case data.BodyStream != nil:
		request.SetBodyStream(data.BodyStream, -1)
	default:
		return nil, errors.New("body is not provided")
	}

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	var err error
	deadline, ok := ctx.Deadline()
	if ok {
		err = a.Client.DoDeadline(request, response, deadline)
	} else {
		err = a.Client.Do(request, response)
	}
	if err != nil {
		return nil, fmt.Errorf("fasthttp do request: %w", err)
	}

	if statusCode := response.StatusCode(); statusCode >= fasthttp.StatusInternalServerError {
		return nil, fmt.Errorf("internal server error: %d", statusCode)
	}

	apiResp := &Response{}
	err = json.Unmarshal(response.Body(), apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	return apiResp, nil
}

// HTTPCaller http implementation of [Caller]
type HTTPCaller struct {
	Client *http.Client
}

// DefaultHTTPCaller is a default http caller
var DefaultHTTPCaller = &HTTPCaller{
	Client: http.DefaultClient,
}

// Call is an http implementation
func (h HTTPCaller) Call(ctx context.Context, url string, data *RequestData) (*Response, error) {
	var requestBody io.Reader
	switch {
	case data.BodyRaw != nil:
		requestBody = bytes.NewReader(data.BodyRaw)
	case data.BodyStream != nil:
		requestBody = data.BodyStream
	default:
		return nil, errors.New("body is not provided")
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("http create request: %w", err)
	}
	request.Header.Set(ContentTypeHeader, data.ContentType)

	response, err := h.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("http do request: %w", err)
	}
	defer func() { _ = response.Body.Close() }() //nolint:errcheck

	if response.StatusCode >= http.StatusInternalServerError {
		return nil, fmt.Errorf("internal server error: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	apiResp := &Response{}
	err = json.Unmarshal(body, apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	return apiResp, nil
}

// RetryCaller decorator over [Caller] that provides retries with exponential backoff
// Depending on [RetryRateLimit] will wait for rate limit timeout to reset or abort, defaults to do nothing
// Delay = min((ExponentBase ^ AttemptNumber) * StartDelay, MaxDelay)
type RetryCaller struct {
	// Underling caller
	Caller Caller
	// Max number of attempts to make a call
	MaxAttempts int
	// Exponent base for delay
	ExponentBase float64
	// Starting delay duration
	StartDelay time.Duration
	// Maximum delay duration
	MaxDelay time.Duration
	// Rate limit behavior
	RateLimit RetryRateLimit
	// Buffer request data, if set to true requests that usually stream body using io.Reader will be buffered
	// to support retrying such requests
	//
	// Warning: Enabling this may lead to excessive memory consumption and OOMKill
	BufferRequestData bool
}

// RetryRateLimit mode for handling rate limits
type RetryRateLimit uint

const (
	// RetryRateLimitSkip do not handle rate limits
	RetryRateLimitSkip RetryRateLimit = iota
	// RetryRateLimitAbort abort retry if hit rate limit
	RetryRateLimitAbort
	// RetryRateLimitWait wait for rate limit timeout to reset
	RetryRateLimitWait
	// RetryRateLimitWaitOrAbort wait for rate limit timeout to reset if it's less than max delay else abort
	RetryRateLimitWaitOrAbort
)

// ErrMaxRetryAttempts returned when max retry attempts reached
var ErrMaxRetryAttempts = errors.New("max retry attempts reached")

// Call makes calls using provided caller with retries
func (r *RetryCaller) Call(ctx context.Context, url string, data *RequestData) (response *Response, err error) {
	if data.BodyStream != nil && r.BufferRequestData {
		data.BodyRaw, err = io.ReadAll(data.BodyStream)
		if err != nil {
			return nil, fmt.Errorf("read body: %w", err)
		}
		data.BodyStream = nil
	}

	for i := 0; i < r.MaxAttempts; i++ {
		response, err = r.Caller.Call(ctx, url, data)
		if err == nil && (response.Error == nil || response.ErrorCode == 0) {
			return response, nil
		}
		if err == nil {
			err = response.Error
		}

		if i == r.MaxAttempts-1 {
			break
		}

		delay, ok := r.handleError(err)
		if !ok {
			return nil, err
		}
		if delay == 0 {
			delay = min(time.Duration(math.Pow(r.ExponentBase, float64(i)))*r.StartDelay, r.MaxDelay)
		}

		select {
		case <-ctx.Done():
			return nil, errors.Join(err, ctx.Err())
		case <-time.After(delay):
			// Continue
		}
	}

	return nil, errors.Join(err, ErrMaxRetryAttempts)
}

func (r *RetryCaller) handleError(err error) (time.Duration, bool) {
	var apiErr *Error
	if !errors.As(err, &apiErr) {
		return 0, true // Unknown error
	}

	switch apiErr.ErrorCode {
	case http.StatusTooManyRequests:
		switch r.RateLimit { //nolint:revive
		case RetryRateLimitSkip:
			return 0, true // Skip handling
		case RetryRateLimitAbort:
			return 0, false // Abort retry
		case RetryRateLimitWait:
			if apiErr.Parameters == nil {
				return 0, true // Unknown retry after time
			}
			return time.Duration(apiErr.Parameters.RetryAfter) * time.Second, true // Retry after delay
		case RetryRateLimitWaitOrAbort:
			if apiErr.Parameters == nil {
				return 0, true // Unknown retry after time
			}
			delay := time.Duration(apiErr.Parameters.RetryAfter) * time.Second
			if delay > r.MaxDelay {
				return 0, false // Abort, delay too long
			}
			return delay, true // Retry after delay
		default:
			return 0, true // Skip unknown rate limit behavior
		}
	case http.StatusBadRequest, http.StatusRequestEntityTooLarge, http.StatusForbidden, http.StatusUnauthorized:
		return 0, false // Non retryable errors
	default:
		return 0, true // Unknown status
	}
}
