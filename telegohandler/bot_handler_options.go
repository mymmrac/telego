package telegohandler

// WithErrorHandler sets custom error handler to use, handler can be nil (this is the default and results in simply
// logging the error using Bot's logger)
// Note: Because of how handler routing works error handler can only receive original unmodified context, instead of
// the last context defined in user handlers/middlewares, please keep this in mind
func WithErrorHandler(handler ErrorHandler) BotHandlerOption {
	return func(bh *BotHandler) error {
		bh.errorHandler = handler
		return nil
	}
}
