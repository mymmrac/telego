package telego

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"
)

const (
	// Limited by the number of updates in single [Bot.GetUpdates] method call
	defaultLongPollingUpdateChanBuffer = 100
	defaultLongPollingUpdateInterval   = time.Second * 0 // 0s
	defaultLongPollingRetryTimeout     = time.Second * 8 // 8s

	defaultLongPollingUpdateTimeoutInSeconds = 8 // 8s
)

// longPolling represents configuration of getting updates via long polling
type longPolling struct {
	updateChanBuffer uint
	updateInterval   time.Duration
	retryTimeout     time.Duration
}

// LongPollingOption represents an option that can be applied to long polling
type LongPollingOption func(lp *longPolling) error

// WithLongPollingUpdateInterval sets an update interval for long polling. Ensure that between two calls of
// [Bot.GetUpdates] method will be at least specified time, but it could be longer.
// Default is 0s.
// Note: Telegram has built in a timeout mechanism, to properly use it, set GetUpdatesParams.Timeout to desired timeout
// and update interval to 0 (default, recommended way).
func WithLongPollingUpdateInterval(updateInterval time.Duration) LongPollingOption {
	return func(lp *longPolling) error {
		if updateInterval < 0 {
			return fmt.Errorf("update interval is negative: %s", updateInterval)
		}
		lp.updateInterval = updateInterval
		return nil
	}
}

// WithLongPollingRetryTimeout sets updates retry timeout for long polling.
// Ensure that between two calls of [Bot.GetUpdates] method will be at least specified time if an error occurred,
// but it could be longer. If zero is passed, reties will be disabled and on error update chan will be closed.
// Default is 8s.
func WithLongPollingRetryTimeout(retryTimeout time.Duration) LongPollingOption {
	return func(lp *longPolling) error {
		if retryTimeout < 0 {
			return fmt.Errorf("retry timeout is negative: %s", retryTimeout)
		}
		lp.retryTimeout = retryTimeout
		return nil
	}
}

// WithLongPollingBuffer sets buffering for update chan.
// Default is 100.
func WithLongPollingBuffer(chanBuffer uint) LongPollingOption {
	return func(lp *longPolling) error {
		lp.updateChanBuffer = chanBuffer
		return nil
	}
}

// UpdatesViaLongPolling receive updates in chan using the [Bot.GetUpdates] method.
// Calling if already running long polling or webhook will return an error.
//
// Warning: If nil is passed as get update parameters, then the default timeout of 8s will be applied,
// but if a non-nil parameter is passed, you should remember to explicitly specify timeout
//
// Note: After you done with getting updates, you should close context this will close the update chan
// Note: Value of params is reused to call [Bot.GetUpdates] method many times, because of this we copy params value
func (b *Bot) UpdatesViaLongPolling(
	ctx context.Context, params *GetUpdatesParams, options ...LongPollingOption,
) (<-chan Update, error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.longPollingRunning {
		return nil, errors.New("telego: long polling already running")
	}
	if b.webhookRunning {
		return nil, errors.New("telego: webhook already running")
	}

	lp, err := b.createLongPolling(options)
	if err != nil {
		return nil, err
	}

	updatesChan := make(chan Update, lp.updateChanBuffer)

	if params == nil {
		params = &GetUpdatesParams{
			Timeout: defaultLongPollingUpdateTimeoutInSeconds,
		}
	} else {
		params = &GetUpdatesParams{
			Offset:         params.Offset,
			Limit:          params.Limit,
			Timeout:        params.Timeout,
			AllowedUpdates: slices.Clone(params.AllowedUpdates),
		}
	}

	b.longPollingRunning = true
	go b.doLongPolling(ctx, lp, params, updatesChan)

	return updatesChan, nil
}

func (b *Bot) doLongPolling(ctx context.Context, lp *longPolling, params *GetUpdatesParams, updatesChan chan<- Update) {
	defer func() {
		b.lock.Lock()
		b.longPollingRunning = false
		b.lock.Unlock()

		close(updatesChan)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Continue getting updates
		}

		var updates []Update
		updates, err := b.GetUpdates(ctx, params)
		if err != nil {
			b.log.Errorf("Getting updates: %s", err)
			if lp.retryTimeout == 0 || errors.Is(err, context.Canceled) {
				return
			}

			b.log.Errorf("Retrying getting updates in %s...", lp.retryTimeout.String())
			time.Sleep(lp.retryTimeout)
			continue
		}

		for _, update := range updates {
			if update.UpdateID >= params.Offset {
				params.Offset = update.UpdateID + 1

				select {
				case <-ctx.Done():
					return
				case updatesChan <- update.WithContext(ctx):
					// Continue
				}
			}
		}

		if lp.updateInterval > 0 {
			time.Sleep(lp.updateInterval)
		}
	}
}

func (b *Bot) createLongPolling(options []LongPollingOption) (*longPolling, error) {
	lp := &longPolling{
		updateChanBuffer: defaultLongPollingUpdateChanBuffer,
		updateInterval:   defaultLongPollingUpdateInterval,
		retryTimeout:     defaultLongPollingRetryTimeout,
	}

	for _, option := range options {
		if err := option(lp); err != nil {
			return nil, fmt.Errorf("telego: long polling options: %w", err)
		}
	}

	return lp, nil
}
