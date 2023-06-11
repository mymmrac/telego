package telego

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/fasthttp/router"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

const defaultWebhookUpdateChanBuffer = 128

// WebhookHandler user handler for incoming updates
type WebhookHandler func(data []byte) error

// WebhookServer represents generic webhook server
type WebhookServer interface {
	Start(address string) error
	Stop(ctx context.Context) error
	RegisterHandler(path string, handler WebhookHandler) error
}

// webhookContext represents configuration of getting updates via webhook
type webhookContext struct {
	running     bool
	configured  bool
	runningLock sync.RWMutex
	stop        chan struct{}
	ctx         context.Context

	server WebhookServer

	updateChanBuffer uint
}

// WebhookOption represents an option that can be applied to webhookContext
type WebhookOption func(bot *Bot, ctx *webhookContext) error

// WithWebhookBuffer sets buffering for update chan. Default is 128.
func WithWebhookBuffer(chanBuffer uint) WebhookOption {
	return func(_ *Bot, ctx *webhookContext) error {
		ctx.updateChanBuffer = chanBuffer
		return nil
	}
}

// WithWebhookServer sets webhook server to use for webhook. Default is FastHTTPWebhookServer
func WithWebhookServer(server WebhookServer) WebhookOption {
	return func(_ *Bot, ctx *webhookContext) error {
		if server == nil {
			return errors.New("webhook server is nil")
		}

		ctx.server = server
		return nil
	}
}

// WithWebhookSet calls [Bot.SetWebhook] method before starting webhook
// Note: Calling [Bot.SetWebhook] method multiple times in a row may give "too many requests" errors
func WithWebhookSet(params *SetWebhookParams) WebhookOption {
	return func(bot *Bot, ctx *webhookContext) error {
		return bot.SetWebhook(params)
	}
}

// WithWebhookContext sets context used in webhook server, this context will be added to each update
//
// Warning: Canceling the context doesn't stop webhook server, it only closes update chan,
// be sure to stop server by calling [Bot.StopWebhook] or [Bot.StopWebhookWithContext] methods
func WithWebhookContext(ctx context.Context) WebhookOption {
	return func(_ *Bot, wCtx *webhookContext) error {
		if ctx == nil {
			return errors.New("context is nil")
		}

		wCtx.ctx = ctx
		return nil
	}
}

// UpdatesViaWebhook receive updates in chan from webhook.
// A new handler with a provided path will be registered on server.
// Calling if already configured (before [Bot.StopWebhook] method) will return an error.
// Note: Once stopped, update chan will be closed
func (b *Bot) UpdatesViaWebhook(path string, options ...WebhookOption) (<-chan Update, error) {
	if b.webhookContext != nil {
		return nil, errors.New("telego: webhook context already exists")
	}

	ctx, err := b.createWebhookContext(options)
	if err != nil {
		return nil, err
	}

	ctx.runningLock.Lock()
	defer ctx.runningLock.Unlock()

	b.webhookContext = ctx
	ctx.stop = make(chan struct{})
	ctx.configured = true

	updatesChan := make(chan Update, ctx.updateChanBuffer)

	err = ctx.server.RegisterHandler(path, func(data []byte) error {
		b.log.Debugf("Webhook request with data: %s", string(data))

		var update Update
		err = json.Unmarshal(data, &update)
		if err != nil {
			b.log.Errorf("Webhook decoding error: %s", err)
			return fmt.Errorf("telego: webhook decoding update: %w", err)
		}

		select {
		case <-ctx.stop:
			return fmt.Errorf("telego: webhook stopped")
		case <-ctx.ctx.Done():
			return fmt.Errorf("telego: %w", ctx.ctx.Err())
		default:
			if safeSend(updatesChan, update.WithContext(ctx.ctx)) {
				return fmt.Errorf("telego: webhook stopped")
			}
			return nil
		}
	})
	if err != nil {
		return nil, fmt.Errorf("telego: webhook register handler: %w", err)
	}

	go func() {
		select {
		case <-ctx.stop:
		case <-ctx.ctx.Done():
		}
		close(updatesChan)
	}()

	return updatesChan, nil
}

func (b *Bot) createWebhookContext(options []WebhookOption) (*webhookContext, error) {
	ctx := &webhookContext{
		server: FastHTTPWebhookServer{
			Logger: b.Logger(),
			Server: &fasthttp.Server{},
			Router: router.New(),
		},
		updateChanBuffer: defaultWebhookUpdateChanBuffer,
		ctx:              context.Background(),
	}

	for _, option := range options {
		if err := option(b, ctx); err != nil {
			return nil, fmt.Errorf("telego: options: %w", err)
		}
	}

	return ctx, nil
}

// StartWebhook start server for listening for webhook, blocking operation.
// Any error that occurs will stop the webhook.
// Calling before [Bot.UpdatesViaWebhook] method will return an error.
// Calling if already running (before [Bot.StopWebhook] method) will return an error.
// Note: After you done with getting updates, you should call [Bot.StopWebhook] method to stop the server
func (b *Bot) StartWebhook(address string) error {
	ctx := b.webhookContext
	if ctx == nil {
		return errors.New("telego: webhook context does not exist")
	}

	ctx.runningLock.RLock()
	if !ctx.configured {
		ctx.runningLock.RUnlock()
		return errors.New("telego: webhook context not configured")
	}

	if ctx.running {
		ctx.runningLock.RUnlock()
		return errors.New("telego: webhook already running")
	}
	ctx.runningLock.RUnlock()

	ctx.runningLock.Lock()
	ctx.running = true
	ctx.runningLock.Unlock()

	if err := ctx.server.Start(address); err != nil {
		ctx.runningLock.Lock()
		if ctx.running {
			close(ctx.stop)
			ctx.running = false
		}
		b.webhookContext = nil
		ctx.runningLock.Unlock()

		return err
	}

	return nil
}

// IsRunningWebhook tells if webhook server is running
func (b *Bot) IsRunningWebhook() bool {
	ctx := b.webhookContext
	if ctx == nil {
		return false
	}

	ctx.runningLock.RLock()
	defer ctx.runningLock.RUnlock()

	return ctx.running
}

// StopWebhookWithContext shutdown webhook server used in the [Bot.UpdatesViaWebhook] method.
// Stopping will stop new updates from coming, but processing updates should be handled by the caller.
// Stop will only ensure that no more updates will come in update chan.
// Calling [Bot.StopWebhookWithContext] method multiple times does nothing.
func (b *Bot) StopWebhookWithContext(ctx context.Context) error {
	webhookCtx := b.webhookContext
	if webhookCtx == nil {
		return nil
	}

	webhookCtx.runningLock.Lock()
	defer webhookCtx.runningLock.Unlock()

	if webhookCtx.running {
		err := webhookCtx.server.Stop(ctx)

		close(webhookCtx.stop)
		webhookCtx.running = false

		b.webhookContext = nil
		return err
	}

	b.webhookContext = nil
	return nil
}

// StopWebhook shutdown webhook server used in the [Bot.UpdatesViaWebhook] method
// Note: For more info, see [Bot.StopWebhookWithContext] method
func (b *Bot) StopWebhook() error {
	return b.StopWebhookWithContext(context.Background())
}
