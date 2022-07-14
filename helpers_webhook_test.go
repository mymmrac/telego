package telego

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
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
