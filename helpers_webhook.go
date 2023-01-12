package telego

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fasthttp/router"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/telegoapi"
)

const defaultWebhookUpdateChanBuffer = 100 // Limited by number of updates in single Bot.GetUpdates() call

const webhookHealthAPIPath = "/health"

// longPullingContext represents configuration of getting updates via webhook
type webhookContext struct {
	running     bool
	configured  bool
	runningLock sync.RWMutex
	stop        chan struct{}

	server *fasthttp.Server
	router *router.Router

	updateChanBuffer uint
}

// WebhookOption represents an option that can be applied to webhookContext
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
			return errors.New("webhook server is nil")
		}

		ctx.server = server
		return nil
	}
}

// WithWebhookRouter sets HTTP router to use for webhook. Default is router.New()
// Note: For webhook to work properly POST route with a path specified in Bot.UpdatesViaWebhook() must be unset.
func WithWebhookRouter(rtr *router.Router) WebhookOption {
	return func(ctx *webhookContext) error {
		if rtr == nil {
			return errors.New("webhook router is nil")
		}

		ctx.router = rtr
		return nil
	}
}

// WithWebhookHealthAPI sets basic health API on `/health` path of the router. Keep in mind that should be
// specified only after WithWebhookRouter() option if any.
func WithWebhookHealthAPI() WebhookOption {
	return func(ctx *webhookContext) error {
		ctx.router.GET(webhookHealthAPIPath, func(ctx *fasthttp.RequestCtx) {
			httpHealthResponse(ctx)
		})
		return nil
	}
}

// StartListeningForWebhook start server for listening for webhook. Any error that occurs will stop the webhook.
// Calling before UpdatesViaWebhook() will return an error.
// Calling if already running (before StopWebhook() method) will return an error.
// Note: After you done with getting updates you should call StopWebhook() method to stop the server
func (b *Bot) StartListeningForWebhook(address string) error {
	return b.StartListeningForWebhookCustom(func(server *fasthttp.Server) error {
		return server.ListenAndServe(address)
	})
}

// StartListeningForWebhookTLS start server with TLS for listening for webhook. Any error that occurs will stop the
// webhook.
// Calling before UpdatesViaWebhook() will return an error.
// Calling if already running (before StopWebhook() method) will return an error.
// Note: After you done with getting updates you should call StopWebhook() method to stop the server
func (b *Bot) StartListeningForWebhookTLS(address, certificateFile, keyFile string) error {
	return b.StartListeningForWebhookCustom(func(server *fasthttp.Server) error {
		return server.ListenAndServeTLS(address, certificateFile, keyFile)
	})
}

// StartListeningForWebhookTLSEmbed start server with TLS (embed) for listening for webhook. Any error that occurs
// will stop the webhook.
// Calling before UpdatesViaWebhook() will return an error.
// Calling if already running (before StopWebhook() method) will return an error.
// Note: After you done with getting updates you should call StopWebhook() method to stop the server
func (b *Bot) StartListeningForWebhookTLSEmbed(address string, certificateData []byte, keyData []byte) error {
	return b.StartListeningForWebhookCustom(func(server *fasthttp.Server) error {
		return server.ListenAndServeTLSEmbed(address, certificateData, keyData)
	})
}

// StartListeningForWebhookUNIX start server with UNIX address for listening for webhook. Any error that occurs will
// stop the webhook.
// Calling before UpdatesViaWebhook() will return an error.
// Calling if already running (before StopWebhook() method) will return an error.
// Note: After you done with getting updates you should call StopWebhook() method to stop the server
func (b *Bot) StartListeningForWebhookUNIX(address string, mode os.FileMode) error {
	return b.StartListeningForWebhookCustom(func(server *fasthttp.Server) error {
		return server.ListenAndServeUNIX(address, mode)
	})
}

// StartListeningForWebhookCustom checks the configuration and starts webhook server using provided listen func.
// Note: Listening func can be nil (useful for running multiple bots on the same server).
func (b *Bot) StartListeningForWebhookCustom(listenAndServe func(server *fasthttp.Server) error) error {
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

	if listenAndServe == nil {
		return nil
	}

	go func() {
		if err := listenAndServe(ctx.server); err != nil {
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

// StopWebhook shutdown webhook server used in the UpdatesViaWebhook() method.
// Stopping will stop new updates from coming, but processing updates should be handled by the caller.
// Stop will only ensure that no more updates will come in update chan.
// Calling StopLongPulling() multiple times does nothing.
func (b *Bot) StopWebhook() error {
	ctx := b.webhookContext
	if ctx == nil {
		return nil
	}

	return b.StopWebhookCustom(ctx.server.Shutdown)
}

// StopWebhookCustom shutdown webhook server used in the UpdatesViaWebhook() method using provided shutdown func.
// Calling StopLongPulling() multiple times does nothing.
// Note: Shutdown func can be nil (useful for running multiple bots on the same server).
//
// Warning: If after shutdown func any updates will be passed into updates chan, the program will panic.
// Ensure that any active connections to webhook stopped before returning from shutdown func.
func (b *Bot) StopWebhookCustom(shutdown func() error) error {
	ctx := b.webhookContext
	if ctx == nil {
		return nil
	}

	ctx.runningLock.Lock()
	defer ctx.runningLock.Unlock()

	if ctx.running {
		ctx.running = false

		var err error
		if shutdown != nil {
			err = shutdown()
		}

		close(ctx.stop)
		b.webhookContext = nil
		return err
	}

	b.webhookContext = nil
	return nil
}

// UpdatesViaWebhook receive updates in chan from webhook.
// New POST route with a provided path will be added to the router.
// Calling if already configured (before StopWebhook() method) will return an error.
// Note: UpdatesViaWebhook() will redefine webhook's server handler.
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

	ctx.router.POST(path, func(ctx *fasthttp.RequestCtx) {
		var update Update
		err = json.Unmarshal(ctx.PostBody(), &update)
		if err != nil {
			httpRespondWithError(ctx, fmt.Errorf("decoding update: %w", err))

			b.log.Errorf("Webhook decoding error: %v", err)
			return
		}

		updatesChan <- update
		ctx.SetStatusCode(fasthttp.StatusOK)
	})

	ctx.server.Handler = ctx.router.Handler

	go func() {
		<-ctx.stop
		close(updatesChan)
	}()

	return updatesChan, nil
}

func (b *Bot) createWebhookContext(options []WebhookOption) (*webhookContext, error) {
	ctx := &webhookContext{
		server: &fasthttp.Server{},
		router: router.New(),

		updateChanBuffer: defaultWebhookUpdateChanBuffer,
	}

	for _, option := range options {
		if err := option(ctx); err != nil {
			return nil, fmt.Errorf("telego: options: %w", err)
		}
	}

	return ctx, nil
}

func httpRespondWithError(ctx *fasthttp.RequestCtx, err error) {
	//nolint:errcheck
	errMsg, _ := json.Marshal(map[string]string{
		"error": err.Error(),
	})

	ctx.SetStatusCode(fasthttp.StatusBadRequest)
	ctx.SetContentType(telegoapi.ContentTypeJSON)

	//nolint:errcheck
	_, _ = ctx.Write(errMsg)
}

func httpHealthResponse(ctx *fasthttp.RequestCtx) {
	//nolint:errcheck
	healthData, _ := json.Marshal(map[string]interface{}{
		"running": true,
		"time":    time.Now().Local(),
	})

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType(telegoapi.ContentTypeJSON)

	//nolint:errcheck
	_, _ = ctx.Write(healthData)
}
