package telegohandler

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

const (
	token = "1234567890:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc"

	timeout = time.Second
)

func newBotHandler(t *testing.T) *BotHandler {
	t.Helper()

	bot, err := telego.NewBot(token)
	require.NoError(t, err)

	updates := make(chan telego.Update)

	return NewBotHandler(bot, updates)
}

func TestNewBotHandler(t *testing.T) {
	bot, err := telego.NewBot(token)
	require.NoError(t, err)

	updates := make(chan telego.Update)

	bh := NewBotHandler(bot, updates)
	assert.Equal(t, bot, bh.bot)
	assert.Equal(t, updates, bh.updates)
	assert.Equal(t, []conditionalHandler{}, bh.handlers)
	assert.Nil(t, bh.stop)
}

func TestBotHandler_Start(t *testing.T) {
	bot, err := telego.NewBot(token)
	require.NoError(t, err)

	updates := make(chan telego.Update)

	bh := NewBotHandler(bot, updates)

	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	h1 := 0
	h2 := 0

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		h1++
	}, func(update telego.Update) bool {
		return update.UpdateID == 1
	})

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		h2++
	})

	timeoutSignal := time.After(timeout)
	done := make(chan struct{})

	assert.NotPanics(t, func() {
		wg.Add(2)

		go bh.Start()
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

func TestBotHandler_Stop(t *testing.T) {
	bh := newBotHandler(t)
	bh.stop = make(chan struct{})
	assert.NotPanics(t, func() {
		bh.Stop()
	})
}

func TestBotHandler_Handle(t *testing.T) {
	bh := newBotHandler(t)

	t.Run("panic_nil_handler", func(t *testing.T) {
		assert.Panics(t, func() {
			bh.Handle(nil)
		})
	})

	handler := Handler(func(bot *telego.Bot, update telego.Update) {})

	t.Run("panic_nil_predicate", func(t *testing.T) {
		assert.Panics(t, func() {
			bh.Handle(handler, nil)
		})
	})

	t.Run("without_predicates", func(t *testing.T) {
		bh.Handle(handler)

		require.Equal(t, 1, len(bh.handlers))
		assert.NotNil(t, bh.handlers[0].Handler)
		assert.Nil(t, bh.handlers[0].Predicates)

		bh.handlers = make([]conditionalHandler, 0)
	})

	predicate := Predicate(func(update telego.Update) bool { return false })

	t.Run("with_predicates", func(t *testing.T) {
		bh.Handle(handler, predicate)

		require.Equal(t, 1, len(bh.handlers))
		assert.NotNil(t, bh.handlers[0].Handler)
		assert.NotNil(t, bh.handlers[0].Predicates)

		bh.handlers = make([]conditionalHandler, 0)
	})
}

func TestBotHandler_IsRunning(t *testing.T) {
	bh := newBotHandler(t)

	t.Run("stopped", func(t *testing.T) {
		assert.False(t, bh.IsRunning())
	})

	t.Run("running", func(t *testing.T) {
		go bh.Start()
		time.Sleep(time.Millisecond * 10)
		assert.True(t, bh.IsRunning())

		bh.Stop()
		assert.False(t, bh.IsRunning())
	})
}
