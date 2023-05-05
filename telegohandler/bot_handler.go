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
// Note: Predicate can't change the update, because it uses a copy, not original value
type Predicate func(update telego.Update) bool

// Middleware applies any function on update before calling the handler
type Middleware func(next Handler) Handler

// BotHandler represents a bot handler that can handle updated matching by predicates
type BotHandler struct {
	bot       *telego.Bot
	updates   <-chan telego.Update
	baseGroup *HandlerGroup

	running        bool
	runningLock    sync.RWMutex
	stop           chan struct{}
	handledUpdates *sync.WaitGroup
	stopTimeout    time.Duration
}

// BotHandlerOption represents an option that can be applied to bot handler
type BotHandlerOption func(bh *BotHandler) error

// NewBotHandler creates new bot handler
func NewBotHandler(bot *telego.Bot, updates <-chan telego.Update, options ...BotHandlerOption) (*BotHandler, error) {
	bh := &BotHandler{
		bot:            bot,
		updates:        updates,
		baseGroup:      &HandlerGroup{},
		handledUpdates: &sync.WaitGroup{},
	}

	for _, option := range options {
		if err := option(bh); err != nil {
			return nil, fmt.Errorf("options: %w", err)
		}
	}

	return bh, nil
}

// Start starts handling of updates, blocks execution
// Calling [BotHandler.Start] method multiple times after the first one does nothing.
// Note: After you done with handling updates, you should call [BotHandler.Stop] method,
// because stopping updates chan will do nothing.
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
	defer h.handledUpdates.Done()
	h.runningLock.Unlock()

	for {
		select {
		case <-h.stop:
			return
		case update, ok := <-h.updates:
			if !ok {
				return
			}

			h.processUpdate(update)
		}
	}
}

// processUpdate checks all groups and handlers, tries to process update in first matched handler
func (h *BotHandler) processUpdate(update telego.Update) {
	_ = h.baseGroup.useGroups(h.bot, update, h.handledUpdates)
}

// IsRunning tells if Start is running
func (h *BotHandler) IsRunning() bool {
	h.runningLock.RLock()
	defer h.runningLock.RUnlock()

	return h.running
}

// Stop stops handling of updates, will block until all updates has been processes or on timeout. If timeout set to 0,
// bot handler will not wait for all handlers to complete processing.
// Note: Calling [BotHandler.Stop] method multiple times does nothing. Calling before [BotHandler.Start] method does
// nothing.
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

// Handle registers new handler in the base group, update will be processed only by first-matched handler,
// order of registration determines the order of matching handlers.
// Important to notice, update's context will be automatically canceled once the handler will finish processing.
// Note: All handlers will process updates in parallel, there is no guaranty on order of processed updates, also keep
// in mind that predicates are checked sequentially.
//
// Warning: Panics if nil handler or predicates passed
func (h *BotHandler) Handle(handler Handler, predicates ...Predicate) {
	h.baseGroup.Handle(handler, predicates...)
}

// Group creates a new group of handlers and middlewares from the base group
// Note: Updates first checked by group and only after that by handler
//
// Warning: Panics if nil predicates passed
func (h *BotHandler) Group(predicates ...Predicate) *HandlerGroup {
	return h.baseGroup.Group(predicates...)
}

// Use applies middleware to the base group
// Note: The Handler chain will be stopped if middleware doesn't call the next func
//
// Warning: Panics if nil middlewares passed
func (h *BotHandler) Use(middlewares ...Middleware) {
	h.baseGroup.Use(middlewares...)
}
