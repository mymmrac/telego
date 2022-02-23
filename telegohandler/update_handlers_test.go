package telegohandler

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

// TODO: Refactor to be more general
func TestBotHandler_HandleMessage(t *testing.T) {
	bh := newBotHandler(t)

	wg := sync.WaitGroup{}
	handler := MessageHandler(func(bot *telego.Bot, message telego.Message) { wg.Done() })

	bh.HandleMessage(handler)

	require.Equal(t, 1, len(bh.handlers))
	assert.NotNil(t, bh.handlers[0].Handler)
	assert.NotNil(t, bh.handlers[0].Predicates)
	assert.Equal(t, 1, len(bh.handlers[0].Predicates))

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{Message: &telego.Message{}}

	bh.updates = updates
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
