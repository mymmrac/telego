package telegoapi

import (
	"bytes"
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

const (
	errJSONPath = "/json_err"
	err500Path  = "/500"
)

var _ Caller = DefaultFastHTTPCaller

func TestFastHTTPCaller_Call(t *testing.T) {
	ln := fasthttputil.NewInmemoryListener()

	api := &fasthttpServer{
		t: t,
	}
	srv := fasthttp.Server{
		Handler: api.Handle,
	}

	go func() {
		if err := srv.Serve(ln); err != nil {
			panic(err)
		}
	}()

	teardown := func() {
		require.NoError(t, ln.Close())
	}

	client := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}

	defer teardown()

	caller := FastHTTPCaller{Client: client}

	data := &RequestData{
		ContentType: ContentTypeJSON,
		Buffer:      bytes.NewBufferString("test"),
	}

	ctx := t.Context()

	t.Run("success", func(t *testing.T) {
		resp, err := caller.Call(ctx, "http://localhost", data)
		require.NoError(t, err)
		assert.True(t, resp.Ok)
	})

	t.Run("error_fasthttp_do_request", func(t *testing.T) {
		resp, err := caller.Call(ctx, "abc", data)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_500", func(t *testing.T) {
		resp, err := caller.Call(ctx, "http://localhost/500", data)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_json", func(t *testing.T) {
		resp, err := caller.Call(ctx, "http://localhost/json_err", data)
		require.Error(t, err)
		assert.Nil(t, resp)
	})
}

type fasthttpServer struct {
	t *testing.T
}

func (s *fasthttpServer) Handle(ctx *fasthttp.RequestCtx) {
	assert.True(s.t, ctx.IsPost())
	assert.Equal(s.t, ContentTypeJSON, string(ctx.Request.Header.ContentType())) //nolint:testifylint

	switch string(ctx.Path()) {
	case err500Path:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	case errJSONPath:
		ctx.SetStatusCode(fasthttp.StatusOK)
		_, err := ctx.WriteString("abc")
		require.NoError(s.t, err)
	default:
		ctx.SetStatusCode(fasthttp.StatusOK)
		_, err := ctx.WriteString("{\"ok\": true}")
		require.NoError(s.t, err)
	}
}

var _ Caller = DefaultHTTPCaller

func TestHTTPCaller_Call(t *testing.T) {
	api := &httpServer{
		t: t,
	}

	srv := httptest.NewServer(api)
	defer srv.Close()

	caller := HTTPCaller{Client: srv.Client()}

	data := &RequestData{
		ContentType: ContentTypeJSON,
		Buffer:      bytes.NewBufferString("test"),
	}

	ctx := t.Context()

	t.Run("success", func(t *testing.T) {
		resp, err := caller.Call(ctx, srv.URL, data)
		require.NoError(t, err)
		assert.True(t, resp.Ok)
	})

	t.Run("error_http_create_request", func(t *testing.T) {
		resp, err := caller.Call(ctx, "\x00", data)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_http_do_request", func(t *testing.T) {
		resp, err := caller.Call(ctx, "abc", data)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_500", func(t *testing.T) {
		resp, err := caller.Call(ctx, srv.URL+err500Path, data)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_json", func(t *testing.T) {
		resp, err := caller.Call(ctx, srv.URL+errJSONPath, data)
		require.Error(t, err)
		assert.Nil(t, resp)
	})
}

type httpServer struct {
	t *testing.T
}

func (h *httpServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	assert.Equal(h.t, http.MethodPost, req.Method)
	assert.Equal(h.t, ContentTypeJSON, req.Header.Get(ContentTypeHeader)) //nolint:testifylint

	switch req.RequestURI {
	case err500Path:
		resp.WriteHeader(http.StatusInternalServerError)
	case errJSONPath:
		resp.WriteHeader(http.StatusOK)
		_, err := resp.Write([]byte("abc"))
		assert.NoError(h.t, err)
	default:
		resp.WriteHeader(http.StatusOK)
		_, err := resp.Write([]byte("{\"ok\": true}"))
		assert.NoError(h.t, err)
	}
}

type testRetryCaller struct {
	resp     *Response
	err      error
	attempts int
	okAfter  int
}

func (t *testRetryCaller) Call(_ context.Context, _ string, _ *RequestData) (*Response, error) {
	t.attempts++
	if t.okAfter != 0 && t.attempts > t.okAfter {
		return t.resp, nil
	}
	return t.resp, t.err
}

func TestRetryCaller_Call(t *testing.T) {
	ctx := t.Context()
	expectedResp := &Response{Ok: true}

	t.Run("success", func(t *testing.T) {
		retryCaller := &RetryCaller{
			Caller: &testRetryCaller{
				resp: expectedResp,
				err:  nil,
			},
			MaxAttempts: 1,
		}
		resp, err := retryCaller.Call(ctx, "", nil)
		require.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})

	t.Run("success_retry", func(t *testing.T) {
		retryCaller := &RetryCaller{
			Caller: &testRetryCaller{
				resp:    expectedResp,
				err:     errors.New("test"),
				okAfter: 2,
			},
			MaxAttempts: 3,
		}
		resp, err := retryCaller.Call(ctx, "", nil)
		require.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})

	t.Run("error_retry", func(t *testing.T) {
		retryCaller := &RetryCaller{
			Caller: &testRetryCaller{
				resp: nil,
				err:  errors.New("test"),
			},
			MaxAttempts: 2,
		}
		resp, err := retryCaller.Call(ctx, "", nil)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("max_delay", func(t *testing.T) {
		retryCaller := &RetryCaller{
			Caller: &testRetryCaller{
				resp: nil,
				err:  errors.New("test"),
			},
			MaxAttempts:  2,
			ExponentBase: 2,
			StartDelay:   10,
			MaxDelay:     1,
		}
		resp, err := retryCaller.Call(ctx, "", nil)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_retry_rate_limit_skip", func(t *testing.T) {
		retryCaller := &RetryCaller{
			Caller: &testRetryCaller{
				resp: nil,
				err: &Error{
					ErrorCode:  429,
					Parameters: &ResponseParameters{RetryAfter: 1},
				},
			},
			MaxAttempts: 2,
			RateLimit:   RetryRateLimitSkip,
		}
		resp, err := retryCaller.Call(ctx, "", nil)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_retry_rate_limit_wait", func(t *testing.T) {
		retryCaller := &RetryCaller{
			Caller: &testRetryCaller{
				resp: nil,
				err: &Error{
					ErrorCode:  429,
					Parameters: &ResponseParameters{RetryAfter: 1},
				},
			},
			MaxAttempts: 2,
			RateLimit:   RetryRateLimitWait,
		}
		resp, err := retryCaller.Call(ctx, "", nil)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_retry_rate_limit_abort", func(t *testing.T) {
		retryCaller := &RetryCaller{
			Caller: &testRetryCaller{
				resp: nil,
				err: &Error{
					ErrorCode:  429,
					Parameters: &ResponseParameters{RetryAfter: 1},
				},
			},
			MaxAttempts: 2,
			RateLimit:   RetryRateLimitAbort,
		}
		resp, err := retryCaller.Call(ctx, "", nil)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_retry_rate_limit_wait_or_abort", func(t *testing.T) {
		retryCaller := &RetryCaller{
			Caller: &testRetryCaller{
				resp: nil,
				err: &Error{
					ErrorCode:  429,
					Parameters: &ResponseParameters{RetryAfter: 1},
				},
			},
			MaxAttempts: 2,
			RateLimit:   RetryRateLimitWaitOrAbort,
		}
		resp, err := retryCaller.Call(ctx, "", nil)
		require.Error(t, err)
		assert.Nil(t, resp)
	})
}
