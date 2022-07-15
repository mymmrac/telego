package telego

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	defaultLongPullingUpdateChanBuffer = 100             // Limited by number of updates in single Bot.GetUpdates() call
	defaultLongPullingUpdateInterval   = time.Second / 2 // 0.5s
	defaultLongPullingRetryTimeout     = time.Second * 3 // 3s
)

// longPullingContext represents configuration of getting updates via long pulling
type longPullingContext struct {
	running     bool
	runningLock sync.RWMutex
	stop        chan struct{}

	updateChanBuffer uint
	updateInterval   time.Duration
	retryTimeout     time.Duration
}

// LongPullingOption represents option that can be applied to longPullingContext
type LongPullingOption func(ctx *longPullingContext) error

// WithLongPullingUpdateInterval sets updates interval for long pulling. Ensures that between two calls of
// Bot.GetUpdates() will be at least specified time, but it could be longer. Default is 0.5s.
func WithLongPullingUpdateInterval(updateInterval time.Duration) LongPullingOption {
	return func(ctx *longPullingContext) error {
		if updateInterval < 0 {
			return errors.New("telego: update interval can't be negative")
		}

		ctx.updateInterval = updateInterval
		return nil
	}
}

// WithLongPullingRetryTimeout sets updates retry timeout for long pulling. Ensures that between two calls of
// Bot.GetUpdates() will be at least specified time if an error occurred, but it could be longer. Default is 3s.
func WithLongPullingRetryTimeout(retryTimeout time.Duration) LongPullingOption {
	return func(ctx *longPullingContext) error {
		if retryTimeout < 0 {
			return errors.New("telego: retry timeout can't be negative")
		}

		ctx.retryTimeout = retryTimeout
		return nil
	}
}

// WithLongPullingBuffer sets buffering for update chan. Default is 100.
func WithLongPullingBuffer(chanBuffer uint) LongPullingOption {
	return func(ctx *longPullingContext) error {
		ctx.updateChanBuffer = chanBuffer
		return nil
	}
}

// UpdatesViaLongPulling receive updates in chan using GetUpdates() method.
// Calling if already running (before StopLongPulling() method) will return an error.
// Note: After you done with getting updates you should call StopLongPulling() method which will close update chan.
func (b *Bot) UpdatesViaLongPulling(params *GetUpdatesParams, options ...LongPullingOption) (<-chan Update, error) {
	if b.longPullingContext != nil {
		return nil, errors.New("telego: long pulling context already exist")
	}

	ctx, err := b.createLongPullingContext(options)
	if err != nil {
		return nil, err
	}

	ctx.runningLock.Lock()
	defer ctx.runningLock.Unlock()

	b.longPullingContext = ctx
	ctx.stop = make(chan struct{})
	ctx.running = true

	updatesChan := make(chan Update, ctx.updateChanBuffer)

	if params == nil {
		params = &GetUpdatesParams{}
	}

	go func() {
		defer close(updatesChan)

		for {
			select {
			case <-ctx.stop:
				return
			default:
				// Continue getting updates
			}

			updates, err := b.GetUpdates(params)
			if err != nil {
				b.log.Errorf("Getting updates: %v", err)
				b.log.Errorf("Retrying to get updates in %s", ctx.retryTimeout.String())

				time.Sleep(ctx.retryTimeout)
				continue
			}

			for _, update := range updates {
				if update.UpdateID >= params.Offset {
					params.Offset = update.UpdateID + 1
					updatesChan <- update
				}
			}

			time.Sleep(ctx.updateInterval)
		}
	}()

	return updatesChan, nil
}

func (b *Bot) createLongPullingContext(options []LongPullingOption) (*longPullingContext, error) {
	ctx := &longPullingContext{
		updateChanBuffer: defaultLongPullingUpdateChanBuffer,
		updateInterval:   defaultLongPullingUpdateInterval,
		retryTimeout:     defaultLongPullingRetryTimeout,
	}

	for _, option := range options {
		if err := option(ctx); err != nil {
			return nil, fmt.Errorf("options: %w", err)
		}
	}

	return ctx, nil
}

// IsRunningLongPulling tells if UpdatesViaLongPulling() is running
func (b *Bot) IsRunningLongPulling() bool {
	ctx := b.longPullingContext
	if ctx == nil {
		return false
	}

	ctx.runningLock.RLock()
	defer ctx.runningLock.RUnlock()

	return ctx.running
}

// StopLongPulling stop reviving updates from UpdatesViaLongPulling() method, stopping is non-blocking, it closes update
// chan, so it's caller's responsibility to process all unhandled updates after calling stop. Stop will only ensure
// that no more updates will come in update chan.
// Calling StopLongPulling() multiple times does nothing.
func (b *Bot) StopLongPulling() {
	ctx := b.longPullingContext
	if ctx == nil {
		return
	}

	ctx.runningLock.Lock()
	defer ctx.runningLock.Unlock()

	if ctx.running {
		ctx.running = false
		close(ctx.stop)
		b.longPullingContext = nil
	}
}
