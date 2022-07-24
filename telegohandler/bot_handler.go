package telegohandler

import (
	"fmt"
	"sync"
	"time"

	"github.com/mymmrac/telego"
)

// Handler handles update that came from bot
type Handler func(bot *telego.Bot, update telego.Update)

// Predicate allows filtering updates for handlers
type Predicate func(update telego.Update) bool

// BotHandler represents bot handler that can handle updated matching by predicates
type BotHandler struct {
	bot      *telego.Bot
	updates  <-chan telego.Update
	handlers []*conditionalHandler

	running        bool
	runningLock    sync.RWMutex
	stop           chan struct{}
	handledUpdates sync.WaitGroup
	stopTimeout    time.Duration
}

// BotHandlerOption represents option that can be applied to bot handler
type BotHandlerOption func(bh *BotHandler) error

// NewBotHandler creates new bot handler
func NewBotHandler(bot *telego.Bot, updates <-chan telego.Update, options ...BotHandlerOption) (*BotHandler, error) {
	bh := &BotHandler{
		bot:      bot,
		updates:  updates,
		handlers: make([]*conditionalHandler, 0),
	}

	for _, option := range options {
		if err := option(bh); err != nil {
			return nil, fmt.Errorf("options: %w", err)
		}
	}

	return bh, nil
}

// Start starts handling of updates, blocks execution
// Calling Start() multiple times after the first one does nothing.
// Note: After you done with handling updates you should call Stop() method, because stopping updates chan will do
// nothing.
func (h *BotHandler) Start() {
	h.runningLock.RLock()
	if h.running {
		h.runningLock.RUnlock()
		return
	}
	h.runningLock.RUnlock()

	h.runningLock.Lock()
	h.stop = make(chan struct{})
	h.running = true
	// Prevents calling Wait before single Add call
	h.handledUpdates.Add(1)
	h.runningLock.Unlock()

	for {
		select {
		case <-h.stop:
			h.handledUpdates.Done()
			return
		case update := <-h.updates:
			h.processUpdate(update)
		}
	}
}

// processUpdate checks all handlers and tries to process update in first matched handler
func (h *BotHandler) processUpdate(update telego.Update) {
	for _, ch := range h.handlers {
		if !ch.match(update) {
			continue
		}

		h.handledUpdates.Add(1)
		go func() {
			ch.Handler(h.bot, update)
			h.handledUpdates.Done()
		}()

		return
	}
}

// IsRunning tells if Start is running
func (h *BotHandler) IsRunning() bool {
	h.runningLock.RLock()
	defer h.runningLock.RUnlock()

	return h.running
}

// Stop stops handling of updates, will block until all updates has been processes or on timeout. If timeout set to 0,
// bot handler will not wait for all handlers to done processing.
// Note: Calling Stop() multiple times does nothing. Calling before Start() does nothing.
func (h *BotHandler) Stop() {
	h.runningLock.Lock()
	defer h.runningLock.Unlock()

	if h.running {
		close(h.stop)

		wait := make(chan struct{})
		go func() {
			h.handledUpdates.Wait()
			wait <- struct{}{}
		}()

		select {
		case <-time.After(h.stopTimeout):
		case <-wait:
		}

		h.running = false
	}
}

// Handle registers new handler, update will be processed only by first matched handler, order of registration
// determines order of matching handlers
// Note: All handlers will process updates in parallel, there is no guaranty on order of processed updates also, keep
// in mind that predicates checked sequentially
//
// Warning: Panics if nil handler or predicate passed
func (h *BotHandler) Handle(handler Handler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil handlers not allowed")
	}

	for _, p := range predicates {
		if p == nil {
			panic("Telego: nil predicates not allowed")
		}
	}

	h.handlers = append(h.handlers, &conditionalHandler{
		Handler:    handler,
		Predicates: predicates,
	})
}
