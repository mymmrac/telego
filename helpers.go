package telego

import (
	"errors"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/telegoapi"
)

const (
	updateChanBuffer = 100

	defaultUpdateInterval = time.Second / 2 // 0.5s
	retryTimeout          = time.Second * 3 // 3s
)

const listeningForWebhookErrMsg = "Listening for webhook: %v"

// SetUpdateInterval sets interval of calling GetUpdates in UpdatesViaLongPulling method. Ensures that between two
// calls of GetUpdates will be at least specified time, but it could be longer.
func (b *Bot) SetUpdateInterval(interval time.Duration) {
	b.updateInterval = interval
}

// UpdatesViaLongPulling receive updates in chan using GetUpdates method
// Note: After you done with getting updates you should call StopLongPulling method
func (b *Bot) UpdatesViaLongPulling(params *GetUpdatesParams) (chan Update, error) {
	b.stop = make(chan struct{})
	b.startedLongPulling = true
	updatesChan := make(chan Update, updateChanBuffer)

	if params == nil {
		params = &GetUpdatesParams{}
	}

	go func() {
		for {
			select {
			case <-b.stop:
				close(updatesChan)
				return
			default:
				// Continue getting updates
			}

			updates, err := b.GetUpdates(params)
			if err != nil {
				b.log.Errorf("Getting updates: %v", err)
				b.log.Errorf("Retrying to get updates in %s", retryTimeout.String())

				time.Sleep(retryTimeout)
				continue
			}

			for _, update := range updates {
				if update.UpdateID >= params.Offset {
					params.Offset = update.UpdateID + 1
					updatesChan <- update
				}
			}

			time.Sleep(b.updateInterval)
		}
	}()

	return updatesChan, nil
}

// IsRunningLongPulling tells if UpdatesViaLongPulling is running
func (b *Bot) IsRunningLongPulling() bool {
	return b.startedLongPulling
}

// StopLongPulling stop reviving updates from UpdatesViaLongPulling method
func (b *Bot) StopLongPulling() {
	if b.startedLongPulling {
		b.startedLongPulling = false
		close(b.stop)
	}
}

// StartListeningForWebhook start server for listening for webhook
// Note: After you done with getting updates you should call StopWebhook method
func (b *Bot) StartListeningForWebhook(address string) {
	b.startedWebhook = true
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
func (b *Bot) StopWebhook() error {
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

	updatesChan := make(chan Update, updateChanBuffer)
	b.stop = make(chan struct{})

	b.server.Handler = func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) != path {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			b.log.Errorf("Unknown path was used in webhook: %q", ctx.Path())
			return
		}

		if method := string(ctx.Method()); method != fasthttp.MethodPost {
			err := fmt.Errorf("used invalid HTTP method: %q, required method: %q", method, fasthttp.MethodPost)
			b.respondWithError(ctx, err)

			b.log.Errorf("Webhook invalid HTTP method: %q", method)
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
	errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})

	ctx.SetStatusCode(fasthttp.StatusBadRequest)
	ctx.SetContentType(telegoapi.ContentTypeJSON)

	_, err = ctx.Write(errMsg)
	if err != nil {
		b.log.Error("Writing HTTP:", err)
	}
}
