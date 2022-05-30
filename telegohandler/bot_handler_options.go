package telegohandler

import (
	"time"
)

// WithStopTimeout sets wait timeout to use. If 0 (default value), calling Stop will immediately stop processing
// updates, else will wait for specified time before force stopping.
func WithStopTimeout(timeout time.Duration) BotHandlerOption {
	return func(bh *BotHandler) error {
		bh.stopTimeout = timeout
		return nil
	}
}
