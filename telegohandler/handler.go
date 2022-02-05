package telegohandler

// TODO: Properly tests package

import (
	"sync"

	"github.com/mymmrac/telego"
)

// Handler handles update that came from bot
type Handler func(bot *telego.Bot, update telego.Update)

// Predicate allows filtering updates for handlers
type Predicate func(update telego.Update) bool

type conditionalHandler struct {
	Handler    Handler
	Predicates []Predicate
}

// BotHandler represents bot handler that can handle updated matching by predicates
type BotHandler struct {
	bot      *telego.Bot
	updates  chan telego.Update
	handlers []conditionalHandler

	stop    chan struct{}
	running bool
	mutex   sync.Mutex
}

// NewBotHandler creates new bot handler
func NewBotHandler(bot *telego.Bot, updates chan telego.Update) *BotHandler {
	return &BotHandler{
		bot:      bot,
		updates:  updates,
		handlers: make([]conditionalHandler, 0),
	}
}

// Start starts handling of updates
// Note: After you done with handling updates you should call Stop method
func (h *BotHandler) Start() {
	h.stop = make(chan struct{})

	h.mutex.Lock()
	h.running = true
	h.mutex.Unlock()

	for {
		select {
		case <-h.stop:
			return
		case update := <-h.updates:
			h.processUpdate(update)
		}
	}
}

// processUpdate checks all handlers and tries to process update in first matched handler
func (h *BotHandler) processUpdate(update telego.Update) {
	for _, ch := range h.handlers {
		ok := true
		for _, p := range ch.Predicates {
			if !p(update) {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}

		go ch.Handler(h.bot, update)
		return
	}
}

// IsRunning tells if Start is running
func (h *BotHandler) IsRunning() bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	return h.running
}

// Stop stops handling of updates
// Note: Should be called only after Start method
func (h *BotHandler) Stop() {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if h.running {
		h.running = false
		close(h.stop)
	}
}

// Handle registers new handler, update will be processed only by first matched handler, order of registration
// determines order of matching handlers
// Note: All handlers will process updates in parallel, there is no guaranty on order of processed updates also, keep
// in mind that predicates checked sequentially
func (h *BotHandler) Handle(handler Handler, predicates ...Predicate) {
	if handler == nil {
		panic("Telego: nil handlers not allowed")
	}

	for _, p := range predicates {
		if p == nil {
			panic("Telego: nil predicates not allowed")
		}
	}

	h.handlers = append(h.handlers, conditionalHandler{
		Handler:    handler,
		Predicates: predicates,
	})
}
