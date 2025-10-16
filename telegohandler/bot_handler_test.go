package telegohandler

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

const (
	token = "1234567890:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc"

	timeout      = time.Second
	smallTimeout = time.Millisecond * 10
	hugeTimeout  = time.Hour
)

var errTest = errors.New("error")

func newTestBotHandler(t *testing.T) *BotHandler {
	t.Helper()

	bot, err := telego.NewBot(token)
	require.NoError(t, err)

	updates := make(chan telego.Update)

	bh, err := NewBotHandler(bot, updates)
	require.NoError(t, err)
	return bh
}

func TestNewBotHandler(t *testing.T) {
	bot, err := telego.NewBot(token)
	require.NoError(t, err)

	updates := make(chan telego.Update)

	var bh *BotHandler

	t.Run("success", func(t *testing.T) {
		bh, err = NewBotHandler(bot, updates)
		require.NoError(t, err)

		assert.Equal(t, bot, bh.bot)
		assert.EqualValues(t, updates, bh.updates)
		assert.Equal(t, &HandlerGroup{}, bh.baseGroup)
		assert.Nil(t, bh.stop)
	})

	t.Run("success_with_options", func(t *testing.T) {
		bh, err = NewBotHandler(bot, updates, func(_ *BotHandler) error { return nil })
		require.NoError(t, err)

		assert.Equal(t, bot, bh.bot)
		assert.EqualValues(t, updates, bh.updates)
		assert.Equal(t, &HandlerGroup{}, bh.baseGroup)
		assert.Nil(t, bh.stop)
	})

	t.Run("error_with_options", func(t *testing.T) {
		bh, err = NewBotHandler(bot, updates, func(_ *BotHandler) error { return errTest })

		require.ErrorIs(t, err, errTest)
		assert.Nil(t, bh)
	})
}

func TestBotHandler_Start(t *testing.T) {
	bot, err := telego.NewBot(token)
	require.NoError(t, err)

	updates := make(chan telego.Update)

	bh, err := NewBotHandler(bot, updates)
	require.NoError(t, err)

	wg := sync.WaitGroup{}
	h1 := 0
	h2 := 0

	bh.Handle(func(_ *Context, _ telego.Update) error {
		defer wg.Done()
		h1++
		return nil
	}, func(ctx context.Context, update telego.Update) bool {
		return update.UpdateID == 1
	})

	bh.Handle(func(_ *Context, _ telego.Update) error {
		defer wg.Done()
		h2++
		return nil
	})

	timeoutSignal := time.After(timeout * 10)
	done := make(chan struct{})

	assert.NotPanics(t, func() {
		wg.Add(2)

		go func() {
			errStart := bh.Start()
			assert.NoError(t, errStart)
		}()

		time.Sleep(smallTimeout)
		err = bh.Start()
		assert.Error(t, err)

		defer func() {
			err = bh.Stop()
			assert.NoError(t, err)
		}()

		updates <- telego.Update{}
		updates <- telego.Update{UpdateID: 1}

		go func() {
			wg.Wait()
			done <- struct{}{}
		}()

		select {
		case <-timeoutSignal:
			t.Fatal("Timeout")
		case <-done:
			assert.Equal(t, 1, h1)
			assert.Equal(t, 1, h2)
		}
	})
}

func TestBotHandler_HandleError(t *testing.T) {
	bot, err := telego.NewBot(token)
	require.NoError(t, err)

	updates := make(chan telego.Update)

	wg := sync.WaitGroup{}

	var receivedError error
	bh, err := NewBotHandler(bot, updates, WithErrorHandler(func(_ *Context, _ telego.Update, err error) {
		defer wg.Done()
		receivedError = err
	}))
	require.NoError(t, err)

	bh.Handle(func(_ *Context, _ telego.Update) error {
		return errTest
	})

	timeoutSignal := time.After(timeout * 10)
	done := make(chan struct{})

	assert.NotPanics(t, func() {
		wg.Add(1)

		go func() {
			errStart := bh.Start()
			assert.NoError(t, errStart)
		}()

		defer func() {
			err = bh.Stop()
			assert.NoError(t, err)
		}()

		updates <- telego.Update{}

		go func() {
			wg.Wait()
			done <- struct{}{}
		}()

		select {
		case <-timeoutSignal:
			t.Fatal("Timeout")
		case <-done:
			assert.Equal(t, errTest, receivedError)
		}
	})
}

func TestBotHandler_Stop(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		bh := newTestBotHandler(t)
		bh.stop = make(chan struct{})
		assert.NotPanics(t, func() {
			err := bh.Stop()
			assert.NoError(t, err)
		})
	})

	t.Run("with_timeout", func(t *testing.T) {
		bot, err := telego.NewBot(token)
		require.NoError(t, err)

		updates := make(chan telego.Update)

		bh, err := NewBotHandler(bot, updates)
		require.NoError(t, err)

		bh.Handle(func(_ *Context, _ telego.Update) error {
			time.Sleep(hugeTimeout)
			t.Fatal("timeout didn't work")
			return nil
		})

		timeoutSignal := time.After(timeout)
		done := make(chan struct{})

		assert.NotPanics(t, func() {
			go func() {
				errStart := bh.Start()
				assert.NoError(t, errStart)
			}()
			for !bh.IsRunning() {
				// Wait for handler to start
			}

			updates <- telego.Update{}
			time.Sleep(smallTimeout)

			ctx, cancel := context.WithTimeout(t.Context(), smallTimeout)
			go func() {
				errStop := bh.StopWithContext(ctx)
				assert.ErrorIs(t, errStop, context.DeadlineExceeded)
				done <- struct{}{}
				cancel()
			}()

			select {
			case <-timeoutSignal:
				t.Fatal("Timeout")
			case <-done:
			}
		})
	})

	t.Run("without_timeout", func(t *testing.T) {
		bot, err := telego.NewBot(token)
		require.NoError(t, err)

		updates := make(chan telego.Update, 2)

		bh, err := NewBotHandler(bot, updates)
		require.NoError(t, err)

		called1 := atomic.Int32{}
		bh.Handle(
			func(_ *Context, _ telego.Update) error {
				called1.Add(1)
				return nil
			},
			func(_ context.Context, update telego.Update) bool { return update.UpdateID == 0 },
		)

		called2 := atomic.Int32{}
		bh.Handle(
			func(_ *Context, _ telego.Update) error {
				called2.Add(1)
				return errTest
			},
		)

		timeoutSignal := time.After(timeout)
		done := make(chan struct{})

		assert.NotPanics(t, func() {
			go func() {
				errStart := bh.Start()
				assert.NoError(t, errStart)
			}()
			for !bh.IsRunning() {
				// Wait for handler to start
			}

			updates <- telego.Update{}
			updates <- telego.Update{UpdateID: 1}
			time.Sleep(smallTimeout)

			ctx, cancel := context.WithTimeout(t.Context(), hugeTimeout)
			go func() {
				errStop := bh.StopWithContext(ctx)
				assert.NoError(t, errStop)
				done <- struct{}{}
				cancel()
			}()

			select {
			case <-timeoutSignal:
				t.Fatal("Timeout")
			case <-done:
			}

			assert.Equal(t, int32(1), called1.Load())
			assert.Equal(t, int32(1), called2.Load())
		})
	})

	t.Run("with_unhandled_updates_error", func(t *testing.T) {
		bot, err := telego.NewBot(token)
		require.NoError(t, err)

		updates := make(chan telego.Update, 1000)
		for range cap(updates) - 1 {
			updates <- telego.Update{}
		}

		bh, err := NewBotHandler(bot, updates)
		require.NoError(t, err)

		done := make(chan struct{})
		assert.NotPanics(t, func() {
			go func() {
				startErr := bh.Start()
				if len(updates) == 0 {
					assert.NoError(t, startErr)
				} else {
					assert.Error(t, startErr)
				}
				done <- struct{}{}
			}()
			for !bh.IsRunning() {
				// Wait for handler to start
			}

			err = bh.StopWithContext(t.Context())
			assert.NoError(t, err)
		})

		select {
		case <-done:
		case <-time.After(timeout):
			t.Fatal("Timeout")
		}
	})

	t.Run("with_canceled", func(t *testing.T) {
		bot, err := telego.NewBot(token)
		require.NoError(t, err)

		updates := make(chan telego.Update)

		bh, err := NewBotHandler(bot, updates)
		require.NoError(t, err)

		bh.Handle(func(_ *Context, _ telego.Update) error {
			time.Sleep(hugeTimeout)
			t.Fatal("timeout didn't work")
			return nil
		})

		timeoutSignal := time.After(timeout)
		done := make(chan struct{})

		assert.NotPanics(t, func() {
			go func() {
				errStart := bh.Start()
				assert.NoError(t, errStart)
			}()
			for !bh.IsRunning() {
				// Wait for handler to start
			}

			updates <- telego.Update{}

			ctx, cancel := context.WithCancel(t.Context())
			cancel()
			go func() {
				errStop := bh.StopWithContext(ctx)
				assert.ErrorIs(t, errStop, context.Canceled)
				done <- struct{}{}
			}()

			select {
			case <-timeoutSignal:
				t.Fatal("Timeout")
			case <-done:
			}
		})
	})

	t.Run("stop_checked", func(t *testing.T) {
		bot, err := telego.NewBot(token)
		require.NoError(t, err)

		updates := make(chan telego.Update, 1)

		bh, err := NewBotHandler(bot, updates)
		require.NoError(t, err)

		bh.Handle(func(_ *Context, _ telego.Update) error {
			t.Fatal("handled after stop")
			return nil
		})

		assert.NotPanics(t, func() {
			go func() {
				errStart := bh.Start()
				assert.NoError(t, errStart)
			}()
			for !bh.IsRunning() {
				// Wait for handler to start
			}

			err = bh.Stop()
			assert.NoError(t, err)

			updates <- telego.Update{}
		})
	})

	t.Run("updates_close", func(t *testing.T) {
		bot, err := telego.NewBot(token)
		require.NoError(t, err)

		updates := make(chan telego.Update, 1)

		bh, err := NewBotHandler(bot, updates)
		require.NoError(t, err)

		bh.Handle(func(_ *Context, _ telego.Update) error {
			t.Fatal("handled after stop")
			return nil
		})

		assert.NotPanics(t, func() {
			go func() {
				errStart := bh.Start()
				assert.NoError(t, errStart)
			}()
			for !bh.IsRunning() {
				// Wait for handler to start
			}
			close(updates)
		})

		time.Sleep(smallTimeout)
		assert.True(t, bh.IsRunning())
	})
}

func TestBotHandler_Handle(t *testing.T) {
	bh := newTestBotHandler(t)

	handler := Handler(func(_ *Context, _ telego.Update) error { return nil })
	predicate := Predicate(func(_ context.Context, _ telego.Update) bool { return false })

	bh.Handle(handler, predicate)

	require.Len(t, bh.baseGroup.routes, 1)
	assert.NotNil(t, bh.baseGroup.routes[0].handler)
	assert.NotEmpty(t, bh.baseGroup.routes[0].predicates)
}

func TestBotHandler_Group(t *testing.T) {
	bh := newTestBotHandler(t)

	predicate := Predicate(func(_ context.Context, _ telego.Update) bool { return false })

	newGr := bh.Group(predicate)

	require.Len(t, bh.baseGroup.routes, 1)
	assert.Equal(t, newGr, bh.baseGroup.routes[0].group)
	assert.NotEmpty(t, bh.baseGroup.routes[0].predicates)
}

func TestBotHandler_Use(t *testing.T) {
	bh := newTestBotHandler(t)

	middleware := Handler(func(ctx *Context, update telego.Update) error {
		return ctx.Next(update)
	})

	bh.Use(middleware)

	require.Len(t, bh.baseGroup.routes, 1)
	assert.NotNil(t, bh.baseGroup.routes[0].handler)
}

func TestBotHandler_IsRunning(t *testing.T) {
	bh := newTestBotHandler(t)

	t.Run("stopped", func(t *testing.T) {
		assert.False(t, bh.IsRunning())
	})

	t.Run("running", func(t *testing.T) {
		go func() {
			err := bh.Start()
			assert.NoError(t, err)
		}()
		time.Sleep(smallTimeout)
		assert.True(t, bh.IsRunning())

		err := bh.Stop()
		require.NoError(t, err)
		assert.False(t, bh.IsRunning())
	})
}

func TestBotHandler_BaseGroup(t *testing.T) {
	bh := newTestBotHandler(t)

	assert.Equal(t, bh.baseGroup, bh.BaseGroup())
}
