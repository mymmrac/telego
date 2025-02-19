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

// Start starts handling of updates, blocks execution, caller is responsible for handling all unhandled updates in the
// update channel after bot handler stop (start will return an error in this case)
// Note: Calling if already running will return an error
func (h *BotHandler) Start() error {
	h.lock.Lock()

	if h.running {
		h.lock.Unlock()
		return errors.New("telego: bot handler already running")
	}

	h.running = true
	h.stop = make(chan struct{})

	// Prevents calling Wait before single Add call
	h.handlers.Add(1)
	defer h.handlers.Done()

	h.lock.Unlock()

	depth := h.baseGroup.depth(1)

	for {
		select {
		case <-h.stop:
			if unhandled := len(h.updates); unhandled > 0 {
				return fmt.Errorf("telego: bot handler stopped, %d update(s) left unhandled", unhandled)
			}
			return nil
		case update, ok := <-h.updates:
			if !ok {
				return nil
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
					ctx: ctx,
					ctxBase: &ctxBase{
						bot:      h.bot,
						updateID: update.UpdateID,
						group:    h.baseGroup,
						stack:    append(make([]int, 0, depth), -1),
					},
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

// StopWithContext stops handling of updates, blocks until all updates have been processes (only if update was received
// from the update channel) or when context is canceled, if not all updates were received from the update channel
// [BotHandler.Start] will return an error, if context is canceled context error will be returned
// Note: Calling [BotHandler.StopWithContext] method multiple times or before [BotHandler.Start] does nothing
func (h *BotHandler) StopWithContext(ctx context.Context) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	if !h.running {
		return nil
	}

	close(h.stop)
	h.running = false

	select {
	case <-ctx.Done():
		return ctx.Err()
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
		return ctx.Err()
	case <-wait:
		// Wait for handlers to complete
		return nil
	}
}

// Stop stops handling of updates, will block until all updates have been processes.
// It's recommended to use [BotHandler.StopWithContext] if you want to force stop after some timeout.
func (h *BotHandler) Stop() error {
	return h.StopWithContext(context.Background())
}

// Handle registers new handler in the base group, update will be processed only by first-matched route,
// order of registration determines the order of matching routes.
// Important to notice handler's context will be automatically canceled once the handler will finish processing or
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

// BaseGroup returns a base group used by default in [BotHandler] methods
func (h *BotHandler) BaseGroup() *HandlerGroup {
	return h.baseGroup
}
