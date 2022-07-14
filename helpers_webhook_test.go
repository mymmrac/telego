package telego

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

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
		server:     &fasthttp.Server{},
		stop:       make(chan struct{}),
	}

	return b
}

func TestBot_StartListeningForWebhook(t *testing.T) {
	t.Run("success_and_error_already_running", func(t *testing.T) {
		b := testWebhookBot(t)

		assert.NotPanics(t, func() {
			err := b.StartListeningForWebhook(testAddress(t))
			assert.NoError(t, err)

			time.Sleep(time.Millisecond * 10)
		})

		assert.NotPanics(t, func() {
			err := b.StartListeningForWebhook("test")
			assert.Error(t, err)
		})
	})

	t.Run("error_not_configured", func(t *testing.T) {
		b := testWebhookBot(t)
		b.webhookContext.configured = false

		assert.NotPanics(t, func() {
			err := b.StartListeningForWebhook("test")
			assert.Error(t, err)
		})
	})
}

func TestBot_StartListeningForWebhookTLS(t *testing.T) {
	b := testWebhookBot(t)

	assert.NotPanics(t, func() {
		err := b.StartListeningForWebhookTLS(testAddress(t), "", "")
		assert.NoError(t, err)

		time.Sleep(time.Millisecond * 10)
	})
}

func TestBot_StartListeningForWebhookTLSEmbed(t *testing.T) {
	b := testWebhookBot(t)

	c, k, err := fasthttp.GenerateTestCertificate(testAddress(t))
	require.NoError(t, err)

	assert.NotPanics(t, func() {
		err = b.StartListeningForWebhookTLSEmbed(testAddress(t), c, k)
		assert.NoError(t, err)

		time.Sleep(time.Millisecond * 10)
	})
}

func TestBot_StartListeningForWebhookUNIX(t *testing.T) {
	b := testWebhookBot(t)

	assert.NotPanics(t, func() {
		err := b.StartListeningForWebhookUNIX(filepath.Join(t.TempDir(), testAddress(t)), os.ModeDir)
		assert.NoError(t, err)

		time.Sleep(time.Millisecond * 10)
	})
}

func TestBot_respondWithError(t *testing.T) {
	b, err := NewBot(token, WithDiscardLogger())
	require.NoError(t, err)

	ctx := &fasthttp.RequestCtx{}

	b.respondWithError(ctx, errTest)
	assert.Equal(t, fasthttp.StatusBadRequest, ctx.Response.StatusCode())
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

		_, err = b.UpdatesViaWebhook("/bot")
		require.NoError(t, err)

		assert.NotPanics(t, func() {
			t.Run("invalid_path_error", func(t *testing.T) {
				ctx := &fasthttp.RequestCtx{}
				b.webhookContext.server.Handler(ctx)
			})

			t.Run("invalid_method_error", func(t *testing.T) {
				ctx := &fasthttp.RequestCtx{}
				ctx.Request.SetRequestURI("/bot")
				b.webhookContext.server.Handler(ctx)
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
		err = b.StartListeningForWebhook(addr)
		require.NoError(t, err)

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
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

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

		err = m.Bot.StartListeningForWebhook(testAddress(t))
		assert.NoError(t, err)

		assert.True(t, m.Bot.IsRunningWebhook())

		err = m.Bot.StopWebhook()
		assert.NoError(t, err)

		assert.False(t, m.Bot.IsRunningWebhook())
	})

	t.Run("running_order_error", func(t *testing.T) {
		err := m.Bot.StartListeningForWebhook(testAddress(t))
		assert.Error(t, err)

		_, err = m.Bot.UpdatesViaWebhook("/bot")
		assert.NoError(t, err)
	})
}

func TestWebhookBuffer(t *testing.T) {
	config := &webhookContext{}
	buffer := uint(1)

	err := WithWebhookBuffer(buffer)(config)
	assert.NoError(t, err)
	assert.EqualValues(t, buffer, config.updateChanBuffer)
}

func TestWithWebhookServer(t *testing.T) {
	config := &webhookContext{}
	server := &fasthttp.Server{}

	t.Run("success", func(t *testing.T) {
		err := WithWebhookServer(server)(config)
		assert.NoError(t, err)
		assert.EqualValues(t, server, config.server)
	})

	t.Run("error", func(t *testing.T) {
		err := WithWebhookServer(nil)(config)
		assert.Error(t, err)
	})
}
