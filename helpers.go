package telego

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/go-telegram-bot-api/api"
)

const (
	updateChanBuffer = 100

	defaultUpdateInterval = time.Second / 2 // 0.5s
	retryTimeout          = time.Second * 3 // 3s
)

const listeningForWebhookErrMsg = "Listening for webhook: %v"

// SetUpdateInterval sets interval of calling GetUpdates in GetUpdatesChan method. Ensures that between two calls
// of GetUpdates will be at least specified time, but it could be longer.
func (b *Bot) SetUpdateInterval(interval time.Duration) {
	b.updateInterval = interval
}

// StopGettingUpdates stop reviving updates from GetUpdatesChan method
func (b *Bot) StopGettingUpdates() {
	close(b.stopChannel)
}

// GetUpdatesChan receive updates in chan
func (b *Bot) GetUpdatesChan(params *GetUpdatesParams) (chan Update, error) {
	b.stopChannel = make(chan struct{})
	updatesChan := make(chan Update, updateChanBuffer)

	go func() {
		for {
			select {
			case <-b.stopChannel:
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

// StartListeningForWebhookTLS start server with TLS for listening for webhook
func (b *Bot) StartListeningForWebhookTLS(address, certificateFile, keyFile string) {
	go func() {
		err := fasthttp.ListenAndServeTLS(address, certificateFile, keyFile, b.webhookHandler)
		if err != nil {
			b.log.Errorf(listeningForWebhookErrMsg, err)
		}
	}()
}

// StartListeningForWebhookTLSEmbed start server with TLS (embed) for listening for webhook
func (b *Bot) StartListeningForWebhookTLSEmbed(address string, certificateData []byte, keyData []byte) {
	go func() {
		err := fasthttp.ListenAndServeTLSEmbed(address, certificateData, keyData, b.webhookHandler)
		if err != nil {
			b.log.Errorf(listeningForWebhookErrMsg, err)
		}
	}()
}

// StartListeningForWebhook start server for listening for webhook
func (b *Bot) StartListeningForWebhook(address string) {
	go func() {
		err := fasthttp.ListenAndServe(address, b.webhookHandler)
		if err != nil {
			b.log.Errorf(listeningForWebhookErrMsg, err)
		}
	}()
}

// ListenForWebhook receive updates in chan from webhook
func (b *Bot) ListenForWebhook(path string) (chan Update, error) {
	updatesChan := make(chan Update, updateChanBuffer)

	b.webhookHandler = func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) != path {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			b.log.Errorf("Unknown path was used in webhook: %q", ctx.Path())
			return
		}

		if method := string(ctx.Method()); method != fasthttp.MethodPost {
			err := fmt.Errorf("used invalid HTTP method: %q, required method: %q", method, fasthttp.MethodPost)
			respondWithError(ctx, err)

			b.log.Errorf("Webhook invalid HTTP method: %q", method)
			return
		}

		var update Update
		err := json.Unmarshal(ctx.PostBody(), &update)
		if err != nil {
			respondWithError(ctx, fmt.Errorf("decoding update: %w", err))

			b.log.Errorf("Webhook decoding error: %v", err)
			return
		}

		updatesChan <- update

		ctx.SetStatusCode(fasthttp.StatusOK)
	}

	return updatesChan, nil
}

func respondWithError(ctx *fasthttp.RequestCtx, err error) {
	errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})

	ctx.SetStatusCode(fasthttp.StatusBadRequest)
	ctx.SetContentType(api.ContentTypeJSON)

	_, _ = ctx.Write(errMsg)
}
