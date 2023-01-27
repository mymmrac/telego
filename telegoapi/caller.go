//go:generate mockgen -package mock -destination=mock/caller.go github.com/mymmrac/telego/telegoapi Caller

package telegoapi

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

// FasthttpAPICaller fasthttp implementation of Caller
type FasthttpAPICaller struct {
	Client *fasthttp.Client
}

// Call is a fasthttp implementation
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
