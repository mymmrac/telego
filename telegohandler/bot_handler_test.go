package telegohandler

import (
	"errors"
	"sync"
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

		assert.ErrorIs(t, err, errTest)
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

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		defer wg.Done()
		h1++
	}, func(update telego.Update) bool {
		return update.UpdateID == 1
	})

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		defer wg.Done()
		h2++
	})

	timeoutSignal := time.After(timeout)
	done := make(chan struct{})

	assert.NotPanics(t, func() {
		wg.Add(2)

		go bh.Start()

		// Check if multiple Start calls do nothing
		time.Sleep(smallTimeout)
		bh.Start()

		defer bh.Stop()

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

//revive:disable:cognitive-complexity
//nolint:gocognit
func TestBotHandler_Stop(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		bh := newTestBotHandler(t)
		bh.stop = make(chan struct{})
		assert.NotPanics(t, func() {
			bh.Stop()
		})
	})

	t.Run("with_timeout", func(t *testing.T) {
		bot, err := telego.NewBot(token)
		require.NoError(t, err)

		updates := make(chan telego.Update)

		bh, err := NewBotHandler(bot, updates, WithStopTimeout(smallTimeout))
		require.NoError(t, err)

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			time.Sleep(hugeTimeout)
			t.Fatal("timeout didn't work")
		})

		timeoutSignal := time.After(timeout)
		done := make(chan struct{})

		assert.NotPanics(t, func() {
			go bh.Start()
			for !bh.IsRunning() { //nolint:revive
				// Wait for handler to start
			}

			updates <- telego.Update{}

			go func() {
				bh.Stop()
				done <- struct{}{}
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

		updates := make(chan telego.Update)

		bh, err := NewBotHandler(bot, updates, WithStopTimeout(hugeTimeout))
		require.NoError(t, err)

		bh.Handle(func(bot *telego.Bot, update telego.Update) {})

		timeoutSignal := time.After(timeout)
		done := make(chan struct{})

		assert.NotPanics(t, func() {
			go bh.Start()
			for !bh.IsRunning() { //nolint:revive
				// Wait for handler to start
			}

			updates <- telego.Update{}

			go func() {
				bh.Stop()
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

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			t.Fatal("handled after stop")
		})

		assert.NotPanics(t, func() {
			go bh.Start()
			for !bh.IsRunning() { //nolint:revive
				// Wait for handler to start
			}

			bh.Stop()

			updates <- telego.Update{}
		})
	})

	t.Run("with_done", func(t *testing.T) {
		bot, err := telego.NewBot(token)
		require.NoError(t, err)

		updates := make(chan telego.Update, 1)

		done := make(chan struct{})
		bh, err := NewBotHandler(bot, updates, WithDone(done))
		require.NoError(t, err)

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			t.Fatal("handled after stop")
		})

		assert.NotPanics(t, func() {
			go bh.Start()
			for !bh.IsRunning() { //nolint:revive
				// Wait for handler to start
			}

			close(done)
			time.Sleep(smallTimeout)

			updates <- telego.Update{}
		})
		assert.False(t, bh.IsRunning())
	})

	t.Run("updates_close", func(t *testing.T) {
		bot, err := telego.NewBot(token)
		require.NoError(t, err)

		updates := make(chan telego.Update, 1)

		bh, err := NewBotHandler(bot, updates)
		require.NoError(t, err)

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			t.Fatal("handled after stop")
		})

		assert.NotPanics(t, func() {
			go bh.Start()
			for !bh.IsRunning() { //nolint:revive
				// Wait for handler to start
			}

			close(updates)
		})

		time.Sleep(smallTimeout)
		assert.False(t, bh.IsRunning())
	})
}

func TestBotHandler_Handle(t *testing.T) {
	bh := newTestBotHandler(t)

	handler := Handler(func(bot *telego.Bot, update telego.Update) {})
	predicate := Predicate(func(update telego.Update) bool { return false })

	bh.Handle(handler, predicate)

	require.Equal(t, 1, len(bh.baseGroup.handlers))
	assert.NotNil(t, bh.baseGroup.handlers[0].handler)
	assert.NotNil(t, bh.baseGroup.handlers[0].predicates)

	bh.baseGroup.handlers = make([]conditionalHandler, 0)
}

func TestBotHandler_Group(t *testing.T) {
	bh := newTestBotHandler(t)

	predicate := Predicate(func(update telego.Update) bool { return false })

	newGr := bh.Group(predicate)

	require.Equal(t, 1, len(bh.baseGroup.groups))
	assert.Equal(t, newGr, bh.baseGroup.groups[0])
	assert.NotEmpty(t, bh.baseGroup.groups[0].predicates)
}

func TestBotHandler_Use(t *testing.T) {
	bh := newTestBotHandler(t)

	middleware := Middleware(func(bot *telego.Bot, update telego.Update, next Handler) {
		next(bot, update)
	})

	bh.Use(middleware)

	require.Equal(t, 1, len(bh.baseGroup.middlewares))
	assert.NotNil(t, bh.baseGroup.middlewares[0])
}

func TestBotHandler_IsRunning(t *testing.T) {
	bh := newTestBotHandler(t)

	t.Run("stopped", func(t *testing.T) {
		assert.False(t, bh.IsRunning())
	})

	t.Run("running", func(t *testing.T) {
		go bh.Start()
		time.Sleep(smallTimeout)
		assert.True(t, bh.IsRunning())

		bh.Stop()
		assert.False(t, bh.IsRunning())
	})
}

func TestBotHandler_BaseGroup(t *testing.T) {
	bh := newTestBotHandler(t)

	assert.Equal(t, bh.baseGroup, bh.BaseGroup())
}
