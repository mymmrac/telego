package telego

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/mymmrac/telego/internal/json"
	ta "github.com/mymmrac/telego/telegoapi"
)

func TestBot_UpdatesViaWebhook(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("success", func(t *testing.T) {
		b, err := NewBot(token, WithDiscardLogger())
		require.NoError(t, err)

		_, err = b.UpdatesViaWebhook(testCtx, func(handler WebhookHandler) error {
			return nil
		})
		require.NoError(t, err)
	})

	t.Run("error_webhook_exist", func(t *testing.T) {
		b := &Bot{}

		_, err := b.UpdatesViaWebhook(testCtx, func(handler WebhookHandler) error {
			return nil
		})
		require.NoError(t, err)

		_, err = b.UpdatesViaWebhook(testCtx, func(handler WebhookHandler) error {
			return nil
		})
		require.Error(t, err)
	})

	t.Run("error_long_polling_exist", func(t *testing.T) {
		m := newMockedBot(ctrl)

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil).AnyTimes()

		resp := telegoResponse(t, []Update{
			{UpdateID: 1},
			{UpdateID: 2},
		})
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil).AnyTimes()

		_, err := m.Bot.UpdatesViaLongPolling(testCtx, nil)
		require.NoError(t, err)

		_, err = m.Bot.UpdatesViaWebhook(testCtx, func(handler WebhookHandler) error {
			return nil
		})
		require.Error(t, err)
	})

	t.Run("end_to_end", func(t *testing.T) {
		b, err := NewBot(token, WithDiscardLogger())
		require.NoError(t, err)

		pushUpdate := make(chan struct{})

		expectedUpdate := Update{
			UpdateID: 1,
			Message:  &Message{Text: "ok"},
		}
		expectedUpdateBytes, err := json.Marshal(expectedUpdate)
		require.NoError(t, err)

		updates, err := b.UpdatesViaWebhook(testCtx, func(handler WebhookHandler) error {
			go func() {
				<-pushUpdate
				err = handler(testCtx, expectedUpdateBytes)
				assert.NoError(t, err)
			}()
			return nil
		})
		require.NoError(t, err)

		pushUpdate <- struct{}{}

		select {
		case update, ok := <-updates:
			require.True(t, ok)
			update.ctx = nil
			assert.Equal(t, expectedUpdate, update)
		case <-time.After(timeout):
			t.Fatalf("Timeout")
		}
	})
}

func TestWithWebhookBuffer(t *testing.T) {
	ctx := &webhook{}
	buffer := uint(1)

	err := WithWebhookBuffer(buffer)(nil, ctx)
	require.NoError(t, err)
	assert.EqualValues(t, buffer, ctx.updateChanBuffer)
}

func TestWithWebhookSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)
	ctx := &webhook{}

	m.MockRequestConstructor.EXPECT().JSONRequest(gomock.Any()).Return(&ta.RequestData{
		Buffer: bytes.NewBuffer(nil),
	}, nil)

	m.MockAPICaller.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ta.Response{Ok: true}, nil)

	err := WithWebhookSet(testCtx, &SetWebhookParams{})(m.Bot, ctx)
	require.NoError(t, err)
}
