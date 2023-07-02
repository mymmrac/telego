package telegohandler

import (
	"errors"
	"fmt"
	"time"
)

// WithStopTimeout sets wait timeout to use.
// If 0 (default value), calling [BotHandler.Stop] method will immediately stop processing updates,
// else will wait for a specified time before force stopping.
func WithStopTimeout(timeout time.Duration) BotHandlerOption {
	return func(bh *BotHandler) error {
		if timeout < 0 {
			return fmt.Errorf("timeout is negative: %s", timeout)
		}

		bh.stopTimeout = timeout
		return nil
	}
}

// WithDone sets done chan, if received value or closed, then bot handler will be stopped by calling [BotHandler.Stop]
// method automatically
func WithDone(done <-chan struct{}) BotHandlerOption {
	return func(bh *BotHandler) error {
		if done == nil {
			return errors.New("done chan is nil")
		}

		bh.done = done
		return nil
	}
}
