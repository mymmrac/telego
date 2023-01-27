//go:generate mockgen -package mock -destination=mock/caller.go github.com/mymmrac/telego/telegoapi Caller

package telegoapi

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

// FastHTTPCaller fasthttp implementation of Caller
type FastHTTPCaller struct {
	Client *fasthttp.Client
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

// Call is a http implementation
func (h HTTPCaller) Call(url string, data *RequestData) (*Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, data.Buffer)
	if err != nil {
		return nil, fmt.Errorf("http create request: %w", err)
	}
	req.Header.Set(ContentTypeHeader, ContentTypeJSON)

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
