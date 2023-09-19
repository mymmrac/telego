package telego

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	// Limited by the number of updates in single [Bot.GetUpdates] method call
	defaultLongPollingUpdateChanBuffer = 100
	defaultLongPollingUpdateInterval   = time.Second * 0 // 0s
	defaultLongPollingRetryTimeout     = time.Second * 8 // 8s

	defaultLongPollingUpdateTimeoutInSeconds = 8 // 8s
)

// longPollingContext represents configuration of getting updates via long polling
type longPollingContext struct {
	running     bool
	runningLock sync.RWMutex
	stop        chan struct{}
	ctx         context.Context

	updateChanBuffer uint
	updateInterval   time.Duration
	retryTimeout     time.Duration
}

// LongPollingOption represents an option that can be applied to longPollingContext
type LongPollingOption func(ctx *longPollingContext) error

// WithLongPollingUpdateInterval sets an update interval for long polling. Ensure that between two calls of
// [Bot.GetUpdates] method will be at least specified time, but it could be longer.
// Default is 0s.
// Note: Telegram has built in a timeout mechanism, to properly use it, set GetUpdatesParams.Timeout to desired timeout
// and update interval to 0 (default, recommended way).
func WithLongPollingUpdateInterval(updateInterval time.Duration) LongPollingOption {
	return func(ctx *longPollingContext) error {
		if updateInterval < 0 {
			return fmt.Errorf("update interval is negative: %s", updateInterval)
		}

		ctx.updateInterval = updateInterval
		return nil
	}
}

// WithLongPollingRetryTimeout sets updates retry timeout for long polling.
// Ensure that between two calls of [Bot.GetUpdates] method will be at least specified time if an error occurred,
// but it could be longer.
// Default is 8s.
func WithLongPollingRetryTimeout(retryTimeout time.Duration) LongPollingOption {
	return func(ctx *longPollingContext) error {
		if retryTimeout < 0 {
			return fmt.Errorf("retry timeout is negative: %s", retryTimeout)
		}

		ctx.retryTimeout = retryTimeout
		return nil
	}
}

// WithLongPollingBuffer sets buffering for update chan.
// Default is 100.
func WithLongPollingBuffer(chanBuffer uint) LongPollingOption {
	return func(ctx *longPollingContext) error {
		ctx.updateChanBuffer = chanBuffer
		return nil
	}
}

// WithLongPollingContext sets context used in long polling, this context will be added to each update
//
// Warning: Canceling the context doesn't stop long polling, it only closes update chan,
// be sure to stop long polling by calling [Bot.StopLongPolling] method
func WithLongPollingContext(ctx context.Context) LongPollingOption {
	return func(lCtx *longPollingContext) error {
		if ctx == nil {
			return errors.New("context is nil")
		}

		lCtx.ctx = ctx
		return nil
	}
}

// UpdatesViaLongPolling receive updates in chan using the [Bot.GetUpdates] method.
// Calling if already running (before [Bot.StopLongPolling] method) will return an error.
// Note: After you done with getting updates, you should call [Bot.StopLongPolling] method which will close update chan.
//
// Warning: If nil is passed as get update parameters, then the default timeout of 8s will be applied,
// but if a non-nil parameter is passed, you should remember to explicitly specify timeout
func (b *Bot) UpdatesViaLongPolling(params *GetUpdatesParams, options ...LongPollingOption) (<-chan Update, error) {
	if b.longPollingContext != nil {
		return nil, errors.New("telego: long polling context already exists")
	}

	ctx, err := b.createLongPollingContext(options)
	if err != nil {
		return nil, err
	}

	ctx.runningLock.Lock()
	defer ctx.runningLock.Unlock()

	b.longPollingContext = ctx
	ctx.stop = make(chan struct{})
	ctx.running = true

	updatesChan := make(chan Update, ctx.updateChanBuffer)

	if params == nil {
		params = &GetUpdatesParams{
			Timeout: defaultLongPollingUpdateTimeoutInSeconds,
		}
	}

	go b.doLongPolling(ctx, params, updatesChan)

	return updatesChan, nil
}

func (b *Bot) doLongPolling(ctx *longPollingContext, params *GetUpdatesParams, updatesChan chan<- Update) {
	defer close(updatesChan)

	for {
		select {
		case <-ctx.stop:
			return
		case <-ctx.ctx.Done():
			return
		default:
			// Continue getting updates
		}

		var updates []Update
		updates, err := b.GetUpdates(params)
		if err != nil {
			b.log.Errorf("Getting updates: %s", err)
			b.log.Errorf("Retrying to get updates in %s", ctx.retryTimeout.String())

			time.Sleep(ctx.retryTimeout)
			continue
		}

		for _, update := range updates {
			if update.UpdateID >= params.Offset {
				params.Offset = update.UpdateID + 1

				select {
				case <-ctx.stop:
					return
				case <-ctx.ctx.Done():
					return
				default:
					if safeSend(updatesChan, update.WithContext(ctx.ctx)) {
						b.log.Debugf("Long polling update chan closed")
						return
					}
				}
			}
		}

		time.Sleep(ctx.updateInterval)
	}
}

func (b *Bot) createLongPollingContext(options []LongPollingOption) (*longPollingContext, error) {
	ctx := &longPollingContext{
		updateChanBuffer: defaultLongPollingUpdateChanBuffer,
		updateInterval:   defaultLongPollingUpdateInterval,
		retryTimeout:     defaultLongPollingRetryTimeout,
		ctx:              context.Background(),
	}

	for _, option := range options {
		if err := option(ctx); err != nil {
			return nil, fmt.Errorf("telego: options: %w", err)
		}
	}

	return ctx, nil
}

// IsRunningLongPolling tells if [Bot.UpdatesViaLongPolling] method is running
func (b *Bot) IsRunningLongPolling() bool {
	ctx := b.longPollingContext
	if ctx == nil {
		return false
	}

	ctx.runningLock.RLock()
	defer ctx.runningLock.RUnlock()

	return ctx.running
}

// StopLongPolling stop reviving updates from [Bot.UpdatesViaLongPolling] method, stopping is non-blocking, it closes
// update chan, so it's caller's responsibility to process all unhandled updates after calling stop.
// Stop will only ensure that no more updates will come in update chan.
// Calling [Bot.StopLongPolling] method multiple times will do nothing.
func (b *Bot) StopLongPolling() {
	ctx := b.longPollingContext
	if ctx == nil {
		return
	}

	ctx.runningLock.Lock()
	defer ctx.runningLock.Unlock()

	if ctx.running {
		close(ctx.stop)
		ctx.running = false
		b.longPollingContext = nil
	}
}
