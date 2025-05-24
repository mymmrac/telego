package telego

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const timeout = time.Second

func TestBot_UpdatesViaLongPolling(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("success", func(t *testing.T) {
		m := newMockedBot(ctrl)

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil).MinTimes(1)

		expectedUpdates := []Update{
			{UpdateID: 1},
			{UpdateID: 2},
		}
		resp := telegoResponse(t, expectedUpdates)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil).MinTimes(1)

		assert.NotPanics(t, func() {
			updates, err := m.Bot.UpdatesViaLongPolling(t.Context(), nil)
			require.NoError(t, err)

			time.Sleep(time.Millisecond * 10)
			select {
			case <-time.After(timeout):
				t.Fatal("Timeout")
			case update := <-updates:
				assert.NotZero(t, update.UpdateID)
			}
		})
	})

	t.Run("error_get_update", func(t *testing.T) {
		m := newMockedBot(ctrl)

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest).MinTimes(1)

		assert.NotPanics(t, func() {
			_, err := m.Bot.UpdatesViaLongPolling(t.Context(), nil)
			require.NoError(t, err)
			time.Sleep(time.Millisecond * 10)
		})
	})

	t.Run("error_already_running", func(t *testing.T) {
		m := newMockedBot(ctrl)

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest).AnyTimes()

		assert.NotPanics(t, func() {
			_, err := m.Bot.UpdatesViaLongPolling(t.Context(), nil)
			require.NoError(t, err)

			_, err = m.Bot.UpdatesViaLongPolling(t.Context(), nil)
			require.Error(t, err)
		})
	})

	t.Run("error_options", func(t *testing.T) {
		m := newMockedBot(ctrl)

		assert.NotPanics(t, func() {
			_, err := m.Bot.UpdatesViaLongPolling(t.Context(), nil, WithLongPollingUpdateInterval(-time.Second))
			require.Error(t, err)
		})
	})
}

func TestWithLongPollingUpdateInterval(t *testing.T) {
	ctx := &longPolling{}
	interval := time.Second

	t.Run("success", func(t *testing.T) {
		err := WithLongPollingUpdateInterval(interval)(ctx)
		require.NoError(t, err)
		assert.Equal(t, interval, ctx.updateInterval)
	})

	t.Run("error", func(t *testing.T) {
		err := WithLongPollingUpdateInterval(-interval)(ctx)
		require.Error(t, err)
	})
}

func TestWithLongPollingRetryTimeout(t *testing.T) {
	ctx := &longPolling{}

	t.Run("success", func(t *testing.T) {
		err := WithLongPollingRetryTimeout(timeout)(ctx)
		require.NoError(t, err)
		assert.Equal(t, timeout, ctx.retryTimeout)
	})

	t.Run("error", func(t *testing.T) {
		err := WithLongPollingRetryTimeout(-timeout)(ctx)
		require.Error(t, err)
	})
}

func TestWithLongPollingBuffer(t *testing.T) {
	ctx := &longPolling{}
	buffer := uint(1)

	err := WithLongPollingBuffer(buffer)(ctx)
	require.NoError(t, err)
	assert.Equal(t, buffer, ctx.updateChanBuffer)
}
