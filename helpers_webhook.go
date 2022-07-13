package telego

import (
	"errors"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/telegoapi"
)

const (
	defaultWebhookUpdateChanBuffer = 100 // Limited by number of updates in single Bot.GetUpdates() call

	listeningForWebhookErrMsg = "Listening for webhook: %v"
)

// StartListeningForWebhook start server for listening for webhook
// Note: After you done with getting updates you should call StopWebhook method
func (b *Bot) StartListeningForWebhook(address string) {
	b.startedWebhook = true // TODO: Add mutex
	go func() {
		err := b.server.ListenAndServe(address)
		if err != nil {
			b.log.Errorf(listeningForWebhookErrMsg, err)
		}
	}()
}

// StartListeningForWebhookTLS start server with TLS for listening for webhook
// Note: After you done with getting updates you should call StopWebhook method
func (b *Bot) StartListeningForWebhookTLS(address, certificateFile, keyFile string) {
	b.startedWebhook = true
	go func() {
		err := b.server.ListenAndServeTLS(address, certificateFile, keyFile)
		if err != nil {
			b.log.Errorf(listeningForWebhookErrMsg, err)
		}
	}()
}

// StartListeningForWebhookTLSEmbed start server with TLS (embed) for listening for webhook
// Note: After you done with getting updates you should call StopWebhook method
func (b *Bot) StartListeningForWebhookTLSEmbed(address string, certificateData []byte, keyData []byte) {
	b.startedWebhook = true
	go func() {
		err := b.server.ListenAndServeTLSEmbed(address, certificateData, keyData)
		if err != nil {
			b.log.Errorf(listeningForWebhookErrMsg, err)
		}
	}()
}

// IsRunningWebhook tells if StartListeningForWebhook... is running
func (b *Bot) IsRunningWebhook() bool {
	return b.startedWebhook
}

// StopWebhook shutdown webhook server used in UpdatesViaWebhook method
// Note: Should be called only after both UpdatesViaWebhook and StartListeningForWebhook...
func (b *Bot) StopWebhook() error { // TODO: [?] Graceful shutdown
	if b.startedWebhook {
		b.startedWebhook = false
		close(b.stop)
		return b.server.Shutdown()
	}
	return nil
}

// UpdatesViaWebhook receive updates in chan from webhook
func (b *Bot) UpdatesViaWebhook(path string) (chan Update, error) {
	if b.startedWebhook {
		return nil, errors.New("calling UpdatesViaWebhook after starting webhook is not allowed")
	}

	updatesChan := make(chan Update, defaultWebhookUpdateChanBuffer)
	b.stop = make(chan struct{})

	b.server.Handler = func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) != path {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			b.log.Errorf("Unknown path was used in webhook: %q", ctx.Path())
			return
		}

		if method := string(ctx.Method()); method != fasthttp.MethodPost {
			err := fmt.Errorf("unexpected HTTP method, expected: %q, got: %q", fasthttp.MethodPost, method)
			b.respondWithError(ctx, err)

			b.log.Errorf("Webhook unexpected HTTP method, expected: %q, got: %q", fasthttp.MethodPost, method)
			return
		}

		var update Update
		err := json.Unmarshal(ctx.PostBody(), &update)
		if err != nil {
			b.respondWithError(ctx, fmt.Errorf("decoding update: %w", err))

			b.log.Errorf("Webhook decoding error: %v", err)
			return
		}

		updatesChan <- update

		ctx.SetStatusCode(fasthttp.StatusOK)
	}

	go func() {
		<-b.stop
		close(updatesChan)
	}()

	return updatesChan, nil
}

func (b *Bot) respondWithError(ctx *fasthttp.RequestCtx, err error) {
	//nolint:errcheck
	errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})

	ctx.SetStatusCode(fasthttp.StatusBadRequest)
	ctx.SetContentType(telegoapi.ContentTypeJSON)

	_, err = ctx.Write(errMsg)
	if err != nil {
		b.log.Error("Writing HTTP:", err)
	}
}
