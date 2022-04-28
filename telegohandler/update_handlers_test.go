package telegohandler

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

func testHandler(t *testing.T, bh *BotHandler, wg *sync.WaitGroup) {
	t.Helper()

	wg.Add(1)

	timeoutSignal := time.After(timeout)
	done := make(chan struct{})

	go bh.Start()

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	select {
	case <-timeoutSignal:
		t.Fatal("Timeout")
	case <-done:
	}
	bh.Stop()
}

func testHandlerSetup(t *testing.T, bh *BotHandler) {
	t.Helper()

	require.Equal(t, 1, len(bh.handlers))
	require.NotNil(t, bh.handlers[0].Handler)
	require.NotNil(t, bh.handlers[0].Predicates)
	require.Equal(t, 1, len(bh.handlers[0].Predicates))
}

func TestBotHandler_HandleMessage(t *testing.T) {
	bh := newBotHandler(t)

	wg := &sync.WaitGroup{}
	handler := MessageHandler(func(bot *telego.Bot, message telego.Message) { wg.Done() })

	bh.HandleMessage(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{Message: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleEditedMessage(t *testing.T) {
	bh := newBotHandler(t)

	wg := &sync.WaitGroup{}
	handler := MessageHandler(func(bot *telego.Bot, message telego.Message) { wg.Done() })

	bh.HandleEditedMessage(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{EditedMessage: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleChannelPost(t *testing.T) {
	bh := newBotHandler(t)

	wg := &sync.WaitGroup{}
	handler := MessageHandler(func(bot *telego.Bot, message telego.Message) { wg.Done() })

	bh.HandleChannelPost(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChannelPost: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleCallbackQuery(t *testing.T) {
	bh := newBotHandler(t)

	wg := &sync.WaitGroup{}
	handler := CallbackQueryHandler(func(bot *telego.Bot, query telego.CallbackQuery) { wg.Done() })

	bh.HandleCallbackQuery(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{CallbackQuery: &telego.CallbackQuery{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}
