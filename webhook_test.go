package telego

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/fasthttp/router"
	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/telegoapi"
)

func testWebhookBot(t *testing.T) *Bot {
	t.Helper()

	b, err := NewBot(token, WithDiscardLogger())
	require.NoError(t, err)

	b.webhookContext = &webhookContext{
		running:    false,
		configured: true,
		server: FastHTTPWebhookServer{
			Server: &fasthttp.Server{},
			Router: router.New(),
		},
		stop: make(chan struct{}),
	}

	return b
}

func TestBot_StartWebhook(t *testing.T) {
	t.Run("success_and_error_already_running_and_start_fail", func(t *testing.T) {
		b := testWebhookBot(t)

		testAddr := testAddress(t)

		assert.NotPanics(t, func() {
			go func() {
				err := b.StartWebhook(testAddr)
				assert.NoError(t, err)
			}()
			time.Sleep(time.Millisecond * 10)
		})

		assert.NotPanics(t, func() {
			err := b.StartWebhook("test")
			assert.Error(t, err)
		})

		b.webhookContext.running = false
		assert.NotPanics(t, func() {
			err := b.StartWebhook(testAddr)
			assert.Error(t, err)
		})
	})

	t.Run("error_not_configured", func(t *testing.T) {
		b := testWebhookBot(t)
		b.webhookContext.configured = false

		assert.NotPanics(t, func() {
			err := b.StartWebhook("test")
			assert.Error(t, err)
		})
	})
}

func TestBot_StopWebhook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		b := testWebhookBot(t)

		assert.NotPanics(t, func() {
			err := b.StopWebhook()
			assert.NoError(t, err)
		})
	})

	t.Run("success_no_context", func(t *testing.T) {
		b := &Bot{}

		assert.NotPanics(t, func() {
			err := b.StopWebhook()
			assert.NoError(t, err)
		})
	})
}

func TestBot_UpdatesViaWebhook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		b, err := NewBot(token, WithDiscardLogger())
		require.NoError(t, err)

		srv := &fasthttp.Server{}
		_, err = b.UpdatesViaWebhook("/bot", WithWebhookServer(FastHTTPWebhookServer{
			Server: srv,
			Router: router.New(),
		}))
		require.NoError(t, err)

		assert.NotPanics(t, func() {
			t.Run("invalid_path_error", func(t *testing.T) {
				ctx := &fasthttp.RequestCtx{}
				srv.Handler(ctx)
			})

			t.Run("invalid_method_error", func(t *testing.T) {
				ctx := &fasthttp.RequestCtx{}
				ctx.Request.SetRequestURI("/bot")
				srv.Handler(ctx)
			})
		})
	})

	t.Run("error_context_exist", func(t *testing.T) {
		b := &Bot{}

		b.webhookContext = &webhookContext{}
		_, err := b.UpdatesViaWebhook("/bot")
		require.Error(t, err)
	})

	t.Run("error_create_context", func(t *testing.T) {
		b := &Bot{}

		_, err := b.UpdatesViaWebhook("/bot", WithWebhookServer(nil))
		require.Error(t, err)
	})

	t.Run("end_to_end", func(t *testing.T) {
		b, err := NewBot(token, WithDiscardLogger())
		require.NoError(t, err)

		require.False(t, b.IsRunningWebhook())

		updates, err := b.UpdatesViaWebhook("/")
		require.NoError(t, err)

		require.False(t, b.IsRunningWebhook())

		addr := testAddress(t)
		go func() {
			err = b.StartWebhook(addr)
			require.NoError(t, err)
		}()
		time.Sleep(time.Millisecond * 10)

		require.True(t, b.IsRunningWebhook())

		expectedUpdate := Update{
			UpdateID: 1,
			Message:  &Message{Text: "ok"},
		}
		expectedUpdateBytes, err := json.Marshal(expectedUpdate)
		require.NoError(t, err)

		go func() {
			resp, errHTTP := http.Post(fmt.Sprintf("http://%s", addr), telegoapi.ContentTypeJSON,
				bytes.NewBuffer([]byte{}))
			assert.NoError(t, errHTTP)
			assert.NoError(t, resp.Body.Close())

			require.NotNil(t, resp)
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

			resp, errHTTP = http.Post(fmt.Sprintf("http://%s", addr), telegoapi.ContentTypeJSON,
				bytes.NewBuffer(expectedUpdateBytes))
			assert.NoError(t, errHTTP)
			assert.NoError(t, resp.Body.Close())

			require.NotNil(t, resp)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		}()

		update, ok := <-updates
		require.True(t, ok)

		assert.Equal(t, expectedUpdate, update)

		err = b.StopWebhook()
		assert.NoError(t, err)

		assert.False(t, b.IsRunningWebhook())

		_, ok = <-updates
		assert.False(t, ok)
	})
}

func TestBot_IsRunningWebhook(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("stopped", func(t *testing.T) {
		assert.False(t, m.Bot.IsRunningWebhook())
	})

	t.Run("running", func(t *testing.T) {
		_, err := m.Bot.UpdatesViaWebhook("/bot")
		require.NoError(t, err)

		go func() {
			err = m.Bot.StartWebhook(testAddress(t))
			assert.NoError(t, err)
		}()
		time.Sleep(time.Millisecond * 10)

		assert.True(t, m.Bot.IsRunningWebhook())

		err = m.Bot.StopWebhook()
		assert.NoError(t, err)

		assert.False(t, m.Bot.IsRunningWebhook())
	})

	t.Run("running_order_error", func(t *testing.T) {
		err := m.Bot.StartWebhook(testAddress(t))
		assert.Error(t, err)

		_, err = m.Bot.UpdatesViaWebhook("/bot")
		assert.NoError(t, err)
	})
}

func TestWithWebhookBuffer(t *testing.T) {
	ctx := &webhookContext{}
	buffer := uint(1)

	err := WithWebhookBuffer(buffer)(nil, ctx)
	assert.NoError(t, err)
	assert.EqualValues(t, buffer, ctx.updateChanBuffer)
}

func TestWithWebhookServer(t *testing.T) {
	ctx := &webhookContext{}
	server := FastHTTPWebhookServer{}

	t.Run("success", func(t *testing.T) {
		err := WithWebhookServer(server)(nil, ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, server, ctx.server)
	})

	t.Run("error", func(t *testing.T) {
		err := WithWebhookServer(nil)(nil, ctx)
		assert.Error(t, err)
	})
}

func TestWithWebhookSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)
	ctx := &webhookContext{}

	m.MockRequestConstructor.EXPECT().JSONRequest(gomock.Any()).Return(&telegoapi.RequestData{
		Buffer: bytes.NewBuffer(nil),
	}, nil)

	m.MockAPICaller.EXPECT().Call(gomock.Any(), gomock.Any()).Return(&telegoapi.Response{Ok: true}, nil)

	err := WithWebhookSet(&SetWebhookParams{})(m.Bot, ctx)
	assert.NoError(t, err)
}
