package telego

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/telegoapi"
)

const defaultWebhookUpdateChanBuffer = 100 // Limited by number of updates in single Bot.GetUpdates() call

// longPullingContext represents configuration of getting updates via webhook
type webhookContext struct {
	running     bool
	configured  bool
	runningLock sync.RWMutex
	stop        chan struct{}

	server           *fasthttp.Server
	updateChanBuffer uint
}

// WebhookOption represents option that can be applied to webhookContext
type WebhookOption func(ctx *webhookContext) error

// WithWebhookBuffer sets buffering for update chan. Default is 100.
func WithWebhookBuffer(chanBuffer uint) WebhookOption {
	return func(ctx *webhookContext) error {
		ctx.updateChanBuffer = chanBuffer
		return nil
	}
}

// WithWebhookServer sets HTTP server to use for webhook. Default is &fasthttp.Server{}
func WithWebhookServer(server *fasthttp.Server) WebhookOption {
	return func(ctx *webhookContext) error {
		if server == nil {
			return errors.New("telego: webhook server is nil")
		}

		ctx.server = server
		return nil
	}
}

// StartListeningForWebhook start server for listening for webhook. Any error that occurs will stop the webhook.
// Calling before UpdatesViaWebhook() will return an error.
// Calling if already running (before StopWebhook() method) will return an error.
// Note: After you done with getting updates you should call StopWebhook() method to stop the server
func (b *Bot) StartListeningForWebhook(address string) error {
	return b.startListeningForWebhook(func(server *fasthttp.Server) error {
		return server.ListenAndServe(address)
	})
}

// StartListeningForWebhookTLS start server with TLS for listening for webhook. Any error that occurs will stop the
// webhook.
// Calling before UpdatesViaWebhook() will return an error.
// Calling if already running (before StopWebhook() method) will return an error.
// Note: After you done with getting updates you should call StopWebhook() method to stop the server
func (b *Bot) StartListeningForWebhookTLS(address, certificateFile, keyFile string) error {
	return b.startListeningForWebhook(func(server *fasthttp.Server) error {
		return server.ListenAndServeTLS(address, certificateFile, keyFile)
	})
}

// StartListeningForWebhookTLSEmbed start server with TLS (embed) for listening for webhook. Any error that occurs
// will stop the webhook.
// Calling before UpdatesViaWebhook() will return an error.
// Calling if already running (before StopWebhook() method) will return an error.
// Note: After you done with getting updates you should call StopWebhook() method to stop the server
func (b *Bot) StartListeningForWebhookTLSEmbed(address string, certificateData []byte, keyData []byte) error {
	return b.startListeningForWebhook(func(server *fasthttp.Server) error {
		return server.ListenAndServeTLSEmbed(address, certificateData, keyData)
	})
}

// StartListeningForWebhookUNIX start server with UNIX address for listening for webhook. Any error that occurs will
// stop the webhook.
// Calling before UpdatesViaWebhook() will return an error.
// Calling if already running (before StopWebhook() method) will return an error.
// Note: After you done with getting updates you should call StopWebhook() method to stop the server
func (b *Bot) StartListeningForWebhookUNIX(address string, mode os.FileMode) error {
	return b.startListeningForWebhook(func(server *fasthttp.Server) error {
		return server.ListenAndServeUNIX(address, mode)
	})
}

// startListeningForWebhook checks configuration and starts webhook server
func (b *Bot) startListeningForWebhook(listenAndServe func(server *fasthttp.Server) error) error {
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

	go func() {
		err := listenAndServe(ctx.server)
		if err != nil {
			b.log.Errorf("Listening for webhook: %v", err)

			ctx.runningLock.Lock()
			ctx.running = false
			close(ctx.stop)
			b.webhookContext = nil
			ctx.runningLock.Unlock()
		}
	}()

	return nil
}

// IsRunningWebhook tells if StartListeningForWebhook[TLS|TLSEmbed|UNIX]() is running
func (b *Bot) IsRunningWebhook() bool {
	ctx := b.webhookContext
	if ctx == nil {
		return false
	}

	ctx.runningLock.RLock()
	defer ctx.runningLock.RUnlock()

	return ctx.running
}

// StopWebhook shutdown webhook server used in UpdatesViaWebhook() method. Stopping will stop new updates from coming,
// but not processes updates should be handled by the caller. Stop will only ensure that no more updates will come
// in update chan.
// Calling StopLongPulling() multiple times does nothing.
func (b *Bot) StopWebhook() error {
	ctx := b.webhookContext
	if ctx == nil {
		return nil
	}

	ctx.runningLock.Lock()
	defer ctx.runningLock.Unlock()

	if ctx.running {
		ctx.running = false
		err := ctx.server.Shutdown()
		close(ctx.stop)
		b.webhookContext = nil
		return err
	}

	b.webhookContext = nil
	return nil
}

// UpdatesViaWebhook receive updates in chan from webhook.
// Calling if already configured (before StopWebhook() method) will return an error.
// TODO: Add end to end test for this
func (b *Bot) UpdatesViaWebhook(path string, options ...WebhookOption) (<-chan Update, error) {
	if b.webhookContext != nil {
		return nil, errors.New("telego: webhook context already exist")
	}

	ctx, err := b.createWebhookContext(options)
	if err != nil {
		return nil, err
	}

	ctx.runningLock.Lock()
	b.webhookContext = ctx
	ctx.stop = make(chan struct{})
	ctx.configured = true
	ctx.runningLock.Unlock()

	updatesChan := make(chan Update, ctx.updateChanBuffer)

	ctx.server.Handler = func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) != path {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			b.log.Errorf("Unknown path was used in webhook: %q", ctx.Path())
			return
		}

		if method := string(ctx.Method()); method != fasthttp.MethodPost {
			err = fmt.Errorf("unexpected HTTP method, expected: %q, got: %q", fasthttp.MethodPost, method)
			b.respondWithError(ctx, err)

			b.log.Errorf("Webhook unexpected HTTP method, expected: %q, got: %q", fasthttp.MethodPost, method)
			return
		}

		var update Update
		err = json.Unmarshal(ctx.PostBody(), &update)
		if err != nil {
			b.respondWithError(ctx, fmt.Errorf("decoding update: %w", err))

			b.log.Errorf("Webhook decoding error: %v", err)
			return
		}

		updatesChan <- update
		ctx.SetStatusCode(fasthttp.StatusOK)
	}

	go func() {
		<-ctx.stop
		close(updatesChan)
	}()

	return updatesChan, nil
}

func (b *Bot) createWebhookContext(options []WebhookOption) (*webhookContext, error) {
	ctx := &webhookContext{
		server:           &fasthttp.Server{},
		updateChanBuffer: defaultWebhookUpdateChanBuffer,
	}

	for _, option := range options {
		if err := option(ctx); err != nil {
			return nil, fmt.Errorf("options: %w", err)
		}
	}

	return ctx, nil
}

func (b *Bot) respondWithError(ctx *fasthttp.RequestCtx, err error) {
	//nolint:errcheck
	// Marshal will never return an error in such case
	errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})

	ctx.SetStatusCode(fasthttp.StatusBadRequest)
	ctx.SetContentType(telegoapi.ContentTypeJSON)

	//nolint:errcheck
	// Write never returns an error
	_, _ = ctx.Write(errMsg)
}
