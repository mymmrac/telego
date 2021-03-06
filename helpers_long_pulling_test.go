package telego

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBot_UpdatesViaLongPulling(t *testing.T) {
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
			_, err := m.Bot.UpdatesViaLongPulling(nil)
			assert.NoError(t, err)
			time.Sleep(time.Millisecond * 10)
			m.Bot.StopLongPulling()
			time.Sleep(time.Millisecond * 500)
		})
	})

	t.Run("error_get_update", func(t *testing.T) {
		m := newMockedBot(ctrl)

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest).MinTimes(1)

		assert.NotPanics(t, func() {
			_, err := m.Bot.UpdatesViaLongPulling(nil)
			assert.NoError(t, err)
			time.Sleep(time.Millisecond * 10)
			m.Bot.StopLongPulling()
		})
	})

	t.Run("error_already_running", func(t *testing.T) {
		m := newMockedBot(ctrl)

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest).AnyTimes()

		assert.NotPanics(t, func() {
			_, err := m.Bot.UpdatesViaLongPulling(nil)
			assert.NoError(t, err)

			_, err = m.Bot.UpdatesViaLongPulling(nil)
			assert.Error(t, err)

			m.Bot.StopLongPulling()
		})
	})

	t.Run("error_options", func(t *testing.T) {
		m := newMockedBot(ctrl)

		assert.NotPanics(t, func() {
			_, err := m.Bot.UpdatesViaLongPulling(nil, WithLongPullingUpdateInterval(-time.Second))
			assert.Error(t, err)
		})
	})
}

func TestBot_IsRunningLongPulling(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("stopped", func(t *testing.T) {
		assert.False(t, m.Bot.IsRunningLongPulling())
	})

	t.Run("running", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil).AnyTimes()

		resp := telegoResponse(t, []Update{})
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil).AnyTimes()

		_, err := m.Bot.UpdatesViaLongPulling(nil)
		require.NoError(t, err)

		assert.True(t, m.Bot.IsRunningLongPulling())

		m.Bot.StopLongPulling()
		assert.False(t, m.Bot.IsRunningLongPulling())
	})
}

func TestBot_StopLongPulling(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bot := &Bot{}

		bot.longPullingContext = &longPullingContext{
			running: true,
			stop:    make(chan struct{}),
		}
		assert.NotPanics(t, func() {
			bot.StopLongPulling()
		})

		assert.Nil(t, bot.longPullingContext)
	})

	t.Run("success_no_context", func(t *testing.T) {
		bot := &Bot{}

		assert.NotPanics(t, func() {
			bot.StopLongPulling()
		})
	})
}

func TestWithLongPullingUpdateInterval(t *testing.T) {
	ctx := &longPullingContext{}
	interval := time.Second

	t.Run("success", func(t *testing.T) {
		err := WithLongPullingUpdateInterval(interval)(ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, interval, ctx.updateInterval)
	})

	t.Run("error", func(t *testing.T) {
		err := WithLongPullingUpdateInterval(-interval)(ctx)
		assert.Error(t, err)
	})
}

func TestWithLongPullingRetryTimeout(t *testing.T) {
	ctx := &longPullingContext{}
	timeout := time.Second

	t.Run("success", func(t *testing.T) {
		err := WithLongPullingRetryTimeout(timeout)(ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, timeout, ctx.retryTimeout)
	})

	t.Run("error", func(t *testing.T) {
		err := WithLongPullingRetryTimeout(-timeout)(ctx)
		assert.Error(t, err)
	})
}

func TestWithLongPullingBuffer(t *testing.T) {
	ctx := &longPullingContext{}
	buffer := uint(1)

	err := WithLongPullingBuffer(buffer)(ctx)
	assert.NoError(t, err)
	assert.EqualValues(t, buffer, ctx.updateChanBuffer)
}
