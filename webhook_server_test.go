package telego

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/fasthttp/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

func TestFastHTTPWebhookServer_RegisterHandler(t *testing.T) {
	addr := testAddress(t)

	s := FastHTTPWebhookServer{
		Logger:      testLoggerType{},
		Server:      &fasthttp.Server{},
		Router:      router.New(),
		SecretToken: "secret",
	}

	go func() {
		err := s.Start(addr)
		require.NoError(t, err)
	}()

	err := s.RegisterHandler("/", func(data []byte) error {
		if len(data) == 0 {
			return nil
		}

		return errTest
	})
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		ctx := &fasthttp.RequestCtx{}
		ctx.Request.SetRequestURI("/")
		ctx.Request.Header.SetMethod(fasthttp.MethodPost)
		ctx.Request.Header.Set(WebhookSecretTokenHeader, s.SecretToken)
		s.Server.Handler(ctx)

		assert.Equal(t, fasthttp.StatusOK, ctx.Response.StatusCode())
	})

	t.Run("error_method", func(t *testing.T) {
		ctx := &fasthttp.RequestCtx{}
		ctx.Request.SetRequestURI("/")
		ctx.Request.Header.SetMethod(fasthttp.MethodGet)
		s.Server.Handler(ctx)

		assert.Equal(t, fasthttp.StatusMethodNotAllowed, ctx.Response.StatusCode())
	})

	t.Run("error_handler", func(t *testing.T) {
		ctx := &fasthttp.RequestCtx{}
		ctx.Request.SetRequestURI("/")
		ctx.Request.Header.SetMethod(fasthttp.MethodPost)
		ctx.Request.Header.Set(WebhookSecretTokenHeader, s.SecretToken)
		ctx.Request.SetBody([]byte("err"))
		s.Server.Handler(ctx)

		assert.Equal(t, fasthttp.StatusInternalServerError, ctx.Response.StatusCode())
	})

	t.Run("secret_token_invalid", func(t *testing.T) {
		ctx := &fasthttp.RequestCtx{}
		ctx.Request.SetRequestURI("/")
		ctx.Request.Header.SetMethod(fasthttp.MethodPost)
		s.Server.Handler(ctx)

		assert.Equal(t, fasthttp.StatusUnauthorized, ctx.Response.StatusCode())
	})

	err = s.Stop(context.Background())
	assert.NoError(t, err)
}

func TestHTTPWebhookServer_RegisterHandler(t *testing.T) {
	t.Run("error_start_fail", func(t *testing.T) {
		s := HTTPWebhookServer{
			Logger:   testLoggerType{},
			Server:   &http.Server{}, //nolint:gosec
			ServeMux: http.NewServeMux(),
		}

		testAddr := testAddress(t)
		go func() {
			err := http.ListenAndServe(testAddr, nil) //nolint:gosec
			require.NoError(t, err)
		}()

		time.Sleep(time.Millisecond * 10)

		err := s.Start(testAddr)
		require.Error(t, err)
	})

	t.Run("end_to_end", func(t *testing.T) {
		s := HTTPWebhookServer{
			Logger:      testLoggerType{},
			Server:      &http.Server{}, //nolint:gosec
			ServeMux:    http.NewServeMux(),
			SecretToken: "secret",
		}

		go func() {
			err := s.Start(testAddress(t))
			require.NoError(t, err)
		}()

		err := s.RegisterHandler("/", func(data []byte) error {
			if len(data) == 0 {
				return nil
			}

			return errTest
		})
		assert.NoError(t, err)

		t.Run("success", func(t *testing.T) {
			rc := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(WebhookSecretTokenHeader, s.SecretToken)

			s.Server.Handler.ServeHTTP(rc, req)

			assert.Equal(t, http.StatusOK, rc.Code)
		})

		t.Run("error_method", func(t *testing.T) {
			rc := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			s.Server.Handler.ServeHTTP(rc, req)

			assert.Equal(t, http.StatusMethodNotAllowed, rc.Code)
		})

		t.Run("error_handler", func(t *testing.T) {
			rc := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("err"))
			req.Header.Set(WebhookSecretTokenHeader, s.SecretToken)

			s.Server.Handler.ServeHTTP(rc, req)

			assert.Equal(t, http.StatusInternalServerError, rc.Code)
		})

		t.Run("error_read", func(t *testing.T) {
			rc := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", errReader{})
			req.Header.Set(WebhookSecretTokenHeader, s.SecretToken)

			s.Server.Handler.ServeHTTP(rc, req)

			assert.Equal(t, http.StatusInternalServerError, rc.Code)
		})

		t.Run("error_close", func(t *testing.T) {
			rc := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", errReaderCloser{reader: strings.NewReader("ok")})
			req.Header.Set(WebhookSecretTokenHeader, s.SecretToken)

			s.Server.Handler.ServeHTTP(rc, req)

			assert.Equal(t, http.StatusInternalServerError, rc.Code)
		})

		t.Run("secret_token_invalid", func(t *testing.T) {
			rc := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", nil)

			s.Server.Handler.ServeHTTP(rc, req)

			assert.Equal(t, http.StatusUnauthorized, rc.Code)
		})

		err = s.Stop(context.Background())
		assert.NoError(t, err)
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

func TestMultiBotWebhookServer_RegisterHandler(t *testing.T) {
	ts := &testServer{}
	s := &MultiBotWebhookServer{
		Server: ts,
	}

	assert.Equal(t, 0, ts.started)
	assert.Equal(t, 0, ts.stopped)
	assert.Equal(t, 0, ts.registered)

	err := s.Start("")
	assert.NoError(t, err)
	assert.Equal(t, 1, ts.started)

	err = s.Start("")
	assert.NoError(t, err)
	assert.Equal(t, 1, ts.started)

	err = s.RegisterHandler("", nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, ts.registered)

	err = s.RegisterHandler("", nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, ts.registered)

	err = s.Stop(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, ts.stopped)

	err = s.Stop(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, ts.stopped)
}

type testServer struct {
	started    int
	stopped    int
	registered int
}

func (t *testServer) Start(_ string) error {
	t.started++
	return nil
}

func (t *testServer) Stop(_ context.Context) error {
	t.stopped++
	return nil
}

func (t *testServer) RegisterHandler(_ string, _ func(data []byte) error) error {
	t.registered++
	return nil
}
