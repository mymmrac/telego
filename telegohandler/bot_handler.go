package telegohandler

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/mymmrac/telego"
)

// Handler handles update that came from bot
type Handler func(ctx *Context, update telego.Update) error

// Predicate allows filtering updates for handlers
// Note: Predicate can't change the update, because it uses a copy, not original value
type Predicate func(ctx context.Context, update telego.Update) bool

// BotHandler represents a bot handler that can handle updated matching by predicates
type BotHandler struct {
	bot       *telego.Bot
	updates   <-chan telego.Update
	baseGroup *HandlerGroup

	running  bool
	lock     sync.RWMutex
	stop     chan struct{}
	handlers sync.WaitGroup
}

// BotHandlerOption represents an option that can be applied to bot handler
type BotHandlerOption func(bh *BotHandler) error

// NewBotHandler creates new bot handler
// Note: Currently no options available, they may be added in future
func NewBotHandler(bot *telego.Bot, updates <-chan telego.Update, options ...BotHandlerOption) (*BotHandler, error) {
	bh := &BotHandler{
		bot:       bot,
		updates:   updates,
		baseGroup: &HandlerGroup{},
	}

	for _, option := range options {
		if err := option(bh); err != nil {
			return nil, fmt.Errorf("telego: bot handler options: %w", err)
		}
	}

	return bh, nil
}

// Start starts handling of updates, blocks execution
// Note: Calling if already running will return an error
func (h *BotHandler) Start() error {
	h.lock.RLock()
	if h.running {
		h.lock.RUnlock()
		return errors.New("telego: bot handler already running")
	}
	h.lock.RUnlock()

	h.lock.Lock()
	h.stop = make(chan struct{})
	h.running = true
	// Prevents calling Wait before single Add call
	h.handlers.Add(1)
	defer h.handlers.Done()
	h.lock.Unlock()

	for {
		select {
		case <-h.stop:
			return nil
		case update, ok := <-h.updates:
			if !ok {
				return errors.New("telego: updates channel closed")
			}

			// Process update
			h.handlers.Add(1)
			go func() {
				defer h.handlers.Done()

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				go func() {
					select {
					case <-ctx.Done():
						// Done processing
					case <-h.stop:
						cancel()
					}
				}()

				bCtx := &Context{
					bot:      h.bot,
					ctx:      ctx,
					updateID: update.UpdateID,
					group:    h.baseGroup,
					stack:    []int{-1}, // TODO: Pre allocate
				}

				if err := bCtx.Next(update); err != nil {
					h.bot.Logger().Errorf("Error processing update %d, err: %s", update.UpdateID, err)
				}
			}()
		}
	}
}

// IsRunning tells if Start is running
func (h *BotHandler) IsRunning() bool {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.running
}

// StopWithContext stops handling of updates, blocks until all updates have been processes or when context is canceled
// Note: Calling [BotHandler.StopWithContext] method multiple times or before [BotHandler.Start] does nothing
func (h *BotHandler) StopWithContext(ctx context.Context) {
	h.lock.Lock()
	defer h.lock.Unlock()

	if !h.running {
		return
	}

	close(h.stop)

	select {
	case <-ctx.Done():
		h.running = false
		return
	default:
		// Continue
	}

	wait := make(chan struct{})
	go func() {
		h.handlers.Wait()
		close(wait)
	}()

	select {
	case <-ctx.Done():
		// Wait for context to be done
	case <-wait:
		// Wait for handlers to complete
	}

	h.running = false
}

// Stop stops handling of updates, will block until all updates have been processes.
// It's recommended to use [BotHandler.StopWithContext] if you want to force stop after some timeout.
func (h *BotHandler) Stop() {
	h.StopWithContext(context.Background())
}

// Handle registers new handler in the base group, update will be processed only by first-matched route,
// order of registration determines the order of matching routes.
// Important to notice, handler's context will be automatically canceled once the handler will finish processing or
// the bot handler stopped.
// Note: All handlers will process updates in parallel, there is no guaranty on order of processed updates, also keep
// in mind that middlewares and predicates are run sequentially.
//
// Warning: Panics if nil handler or predicates passed
func (h *BotHandler) Handle(handler Handler, predicates ...Predicate) {
	h.baseGroup.Handle(handler, predicates...)
}

// Group creates a new group of handlers and middlewares from the base group, update will be processed only by
// first-matched route, order of registration determines the order of matching routes
//
// Warning: Panics if nil predicates passed
func (h *BotHandler) Group(predicates ...Predicate) *HandlerGroup {
	return h.baseGroup.Group(predicates...)
}

// Use applies middleware to the base group, update will be processed only by first-matched route,
// order of registration determines the order of matching routes.
// Note: The chain will be stopped if middleware doesn't call the [Context.Next]
//
// Warning: Panics if nil middlewares passed
func (h *BotHandler) Use(middlewares ...Handler) {
	h.baseGroup.Use(middlewares...)
}

// BaseGroup returns a base group that is used by default in [BotHandler] methods
func (h *BotHandler) BaseGroup() *HandlerGroup {
	return h.baseGroup
}
