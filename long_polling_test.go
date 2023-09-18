package telego

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil).MinTimes(1)

		assert.NotPanics(t, func() {
			_, err := m.Bot.UpdatesViaLongPolling(nil)
			assert.NoError(t, err)
			time.Sleep(time.Millisecond * 10)
			m.Bot.StopLongPolling()
			time.Sleep(time.Millisecond * 500)
		})
	})

	t.Run("error_get_update", func(t *testing.T) {
		m := newMockedBot(ctrl)

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest).MinTimes(1)

		assert.NotPanics(t, func() {
			_, err := m.Bot.UpdatesViaLongPolling(nil)
			assert.NoError(t, err)
			time.Sleep(time.Millisecond * 10)
			m.Bot.StopLongPolling()
		})
	})

	t.Run("error_already_running", func(t *testing.T) {
		m := newMockedBot(ctrl)

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest).AnyTimes()

		assert.NotPanics(t, func() {
			_, err := m.Bot.UpdatesViaLongPolling(nil)
			assert.NoError(t, err)

			_, err = m.Bot.UpdatesViaLongPolling(nil)
			assert.Error(t, err)

			m.Bot.StopLongPolling()
		})
	})

	t.Run("error_options", func(t *testing.T) {
		m := newMockedBot(ctrl)

		assert.NotPanics(t, func() {
			_, err := m.Bot.UpdatesViaLongPolling(nil, WithLongPollingUpdateInterval(-time.Second))
			assert.Error(t, err)
		})
	})

	t.Run("success_with_context", func(t *testing.T) {
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil).MinTimes(1)

		assert.NotPanics(t, func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			updates, err := m.Bot.UpdatesViaLongPolling(nil, WithLongPollingContext(ctx))
			assert.NoError(t, err)

			time.Sleep(time.Millisecond * 10)

			cancel()
			<-updates

			assert.True(t, m.Bot.IsRunningLongPolling())
			m.Bot.StopLongPolling()
			assert.False(t, m.Bot.IsRunningLongPolling())
		})
	})
}

func TestBot_IsRunningLongPolling(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("stopped", func(t *testing.T) {
		assert.False(t, m.Bot.IsRunningLongPolling())
	})

	t.Run("running", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil).AnyTimes()

		resp := telegoResponse(t, []Update{})
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil).AnyTimes()

		_, err := m.Bot.UpdatesViaLongPolling(nil)
		require.NoError(t, err)

		assert.True(t, m.Bot.IsRunningLongPolling())

		m.Bot.StopLongPolling()
		assert.False(t, m.Bot.IsRunningLongPolling())
	})
}

func TestBot_StopLongPolling(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bot := &Bot{}

		bot.longPollingContext = &longPollingContext{
			running: true,
			stop:    make(chan struct{}),
		}
		assert.NotPanics(t, func() {
			bot.StopLongPolling()
		})

		assert.Nil(t, bot.longPollingContext)
	})

	t.Run("success_no_context", func(t *testing.T) {
		bot := &Bot{}

		assert.NotPanics(t, func() {
			bot.StopLongPolling()
		})
	})
}

func TestWithLongPollingUpdateInterval(t *testing.T) {
	ctx := &longPollingContext{}
	interval := time.Second

	t.Run("success", func(t *testing.T) {
		err := WithLongPollingUpdateInterval(interval)(ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, interval, ctx.updateInterval)
	})

	t.Run("error", func(t *testing.T) {
		err := WithLongPollingUpdateInterval(-interval)(ctx)
		assert.Error(t, err)
	})
}

func TestWithLongPollingRetryTimeout(t *testing.T) {
	ctx := &longPollingContext{}
	timeout := time.Second

	t.Run("success", func(t *testing.T) {
		err := WithLongPollingRetryTimeout(timeout)(ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, timeout, ctx.retryTimeout)
	})

	t.Run("error", func(t *testing.T) {
		err := WithLongPollingRetryTimeout(-timeout)(ctx)
		assert.Error(t, err)
	})
}

func TestWithLongPollingBuffer(t *testing.T) {
	ctx := &longPollingContext{}
	buffer := uint(1)

	err := WithLongPollingBuffer(buffer)(ctx)
	assert.NoError(t, err)
	assert.EqualValues(t, buffer, ctx.updateChanBuffer)
}

func TestWithLongPollingContext(t *testing.T) {
	lCtx := &longPollingContext{}

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		err := WithLongPollingContext(ctx)(lCtx)
		assert.NoError(t, err)
		assert.EqualValues(t, ctx, lCtx.ctx)
	})

	t.Run("error", func(t *testing.T) {
		//nolint:staticcheck
		err := WithLongPollingContext(nil)(lCtx)
		assert.Error(t, err)
	})
}
