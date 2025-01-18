package telego

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

func TestWebhookFastHTTP(t *testing.T) {
	srv := &fasthttp.Server{}
	handler := WebhookFastHTTP(srv, "/", "secret")

	err := handler(func(ctx context.Context, data []byte) error {
		require.NotNil(t, ctx)
		if len(data) == 0 {
			return errors.New("empty data")
		}
		return nil
	})
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		ctx := &fasthttp.RequestCtx{}
		ctx.Request.SetRequestURI("/")
		ctx.Request.Header.SetMethod(fasthttp.MethodPost)
		ctx.Request.Header.Set(WebhookSecretTokenHeader, "secret")
		ctx.Request.SetBody([]byte("{}"))
		srv.Handler(ctx)

		assert.Equal(t, fasthttp.StatusOK, ctx.Response.StatusCode())
	})

	t.Run("error_method", func(t *testing.T) {
		ctx := &fasthttp.RequestCtx{}
		ctx.Request.SetRequestURI("/")
		ctx.Request.Header.SetMethod(fasthttp.MethodGet)
		srv.Handler(ctx)

		assert.Equal(t, fasthttp.StatusMethodNotAllowed, ctx.Response.StatusCode())
	})

	t.Run("error_handler", func(t *testing.T) {
		ctx := &fasthttp.RequestCtx{}
		ctx.Request.SetRequestURI("/")
		ctx.Request.Header.SetMethod(fasthttp.MethodPost)
		ctx.Request.Header.Set(WebhookSecretTokenHeader, "secret")
		srv.Handler(ctx)

		assert.Equal(t, fasthttp.StatusInternalServerError, ctx.Response.StatusCode())
	})

	t.Run("secret_token_invalid", func(t *testing.T) {
		ctx := &fasthttp.RequestCtx{}
		ctx.Request.SetRequestURI("/")
		ctx.Request.Header.SetMethod(fasthttp.MethodPost)
		srv.Handler(ctx)

		assert.Equal(t, fasthttp.StatusUnauthorized, ctx.Response.StatusCode())
	})
}

func TestWebhookHTTPServer(t *testing.T) {
	srv := &http.Server{} //nolint:gosec
	handler := WebhookHTTPServer(srv, "/", "secret")

	err := handler(func(ctx context.Context, data []byte) error {
		require.NotNil(t, ctx)
		if len(data) == 0 {
			return errors.New("empty data")
		}
		return nil
	})
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(WebhookSecretTokenHeader, "secret")
		req.Body = io.NopCloser(bytes.NewReader([]byte("{}")))
		srv.Handler.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusOK, rc.Code)
	})

	t.Run("error_method", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		srv.Handler.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusMethodNotAllowed, rc.Code)
	})

	t.Run("error_handler", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		req.Header.Set(WebhookSecretTokenHeader, "secret")
		srv.Handler.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusInternalServerError, rc.Code)
	})

	t.Run("error_read", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", errReader{})
		req.Header.Set(WebhookSecretTokenHeader, "secret")
		srv.Handler.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusInternalServerError, rc.Code)
	})

	t.Run("error_close", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", errReaderCloser{reader: strings.NewReader("ok")})
		req.Header.Set(WebhookSecretTokenHeader, "secret")
		srv.Handler.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusInternalServerError, rc.Code)
	})

	t.Run("secret_token_invalid", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		srv.Handler.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusUnauthorized, rc.Code)
	})
}

func TestWebhookHTTPServeMux(t *testing.T) {
	mux := &http.ServeMux{}
	handler := WebhookHTTPServeMux(mux, "POST /", "secret")

	err := handler(func(ctx context.Context, data []byte) error {
		require.NotNil(t, ctx)
		if len(data) == 0 {
			return errors.New("empty data")
		}
		return nil
	})
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(WebhookSecretTokenHeader, "secret")
		req.Body = io.NopCloser(bytes.NewReader([]byte("{}")))
		mux.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusOK, rc.Code)
	})

	t.Run("error_method", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		mux.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusMethodNotAllowed, rc.Code)
	})

	t.Run("error_handler", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		req.Header.Set(WebhookSecretTokenHeader, "secret")
		mux.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusInternalServerError, rc.Code)
	})

	t.Run("error_read", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", errReader{})
		req.Header.Set(WebhookSecretTokenHeader, "secret")
		mux.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusInternalServerError, rc.Code)
	})

	t.Run("error_close", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", errReaderCloser{reader: strings.NewReader("ok")})
		req.Header.Set(WebhookSecretTokenHeader, "secret")
		mux.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusInternalServerError, rc.Code)
	})

	t.Run("secret_token_invalid", func(t *testing.T) {
		rc := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		mux.ServeHTTP(rc, req)

		assert.Equal(t, http.StatusUnauthorized, rc.Code)
	})
}

type errReader struct{}

func (e errReader) Read(_ []byte) (n int, err error) {
	return 0, errTest
}

type errReaderCloser struct {
	reader io.Reader
}

func (e errReaderCloser) Close() error {
	return errTest
}

func (e errReaderCloser) Read(b []byte) (n int, err error) {
	return e.reader.Read(b)
}
