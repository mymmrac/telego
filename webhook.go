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

// WebhookServer represents generic webhook server
type WebhookServer interface {
	Start(address string) error
	Stop(ctx context.Context) error
	RegisterHandler(path string, handler func(data []byte) error) error
}

// webhookContext represents configuration of getting updates via webhook
type webhookContext struct {
	running     bool
	configured  bool
	runningLock sync.RWMutex
	stop        chan struct{}

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

// WithWebhookSet calls Bot.SetWebhook() before starting webhook
func WithWebhookSet(params *SetWebhookParams) WebhookOption {
	return func(bot *Bot, ctx *webhookContext) error {
		return bot.SetWebhook(params)
	}
}

// UpdatesViaWebhook receive updates in chan from webhook.
// A new handler with a provided path will be registered on server.
// Calling if already configured (before StopWebhook() method) will return an error.
// Note: Once stopped, update chan will be closed
func (b *Bot) UpdatesViaWebhook(path string, options ...WebhookOption) (<-chan Update, error) {
	if b.webhookContext != nil {
		return nil, errors.New("telego: webhook context already exist")
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
		var update Update
		err = json.Unmarshal(data, &update)
		if err != nil {
			b.log.Errorf("Webhook decoding error: %s", err)
			return fmt.Errorf("telego: webhook decoding update: %w", err)
		}

		select {
		case updatesChan <- update:
			return nil
		case <-ctx.stop:
			return fmt.Errorf("telego: webhook stopped")
		}
	})
	if err != nil {
		return nil, fmt.Errorf("telego: webhook register handler: %w", err)
	}

	go func() {
		<-ctx.stop
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
	}

	for _, option := range options {
		if err := option(b, ctx); err != nil {
			return nil, fmt.Errorf("telego: options: %w", err)
		}
	}

	return ctx, nil
}

// StartWebhook start server for listening for webhook.
// Any error that occurs will stop the webhook.
// Calling before UpdatesViaWebhook() will return an error.
// Calling if already running (before StopWebhook() method) will return an error.
// Note: After you done with getting updates you should call StopWebhook() method to stop the server
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
		ctx.running = false
		close(ctx.stop)
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

// StopWebhookWithContext shutdown webhook server used in the UpdatesViaWebhook() method.
// Stopping will stop new updates from coming, but processing updates should be handled by the caller.
// Stop will only ensure that no more updates will come in update chan.
// Calling StopWebhookWithContext() multiple times does nothing.
func (b *Bot) StopWebhookWithContext(ctx context.Context) error {
	webhookCtx := b.webhookContext
	if webhookCtx == nil {
		return nil
	}

	webhookCtx.runningLock.Lock()
	defer webhookCtx.runningLock.Unlock()

	if webhookCtx.running {
		webhookCtx.running = false

		err := webhookCtx.server.Stop(ctx)

		close(webhookCtx.stop)
		b.webhookContext = nil
		return err
	}

	b.webhookContext = nil
	return nil
}

// StopWebhook shutdown webhook server used in the UpdatesViaWebhook() method
// Note: For more info, see StopWebhookWithContext()
func (b *Bot) StopWebhook() error {
	return b.StopWebhookWithContext(context.Background())
}
