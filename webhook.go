package telego

import (
	"context"
	"fmt"

	"github.com/mymmrac/telego/internal/json"
)

const defaultWebhookUpdateChanBuffer = 128

// WebhookHandler user handler for incoming updates, context will be passed into update, user is responsible go pass
// JSON data of update from his own server, and it will be decoded and returned to update chan, also user is
// responsible for validating request (and secret token from headers)
//
// Warning: Common approach of HTTP servers is to cancel context once request connection is closed,
// but in webhook handler update is sent to the channel and not processed in request lifetime,
// so remember to wrap context in [context.WithoutCancel] as webhook helper will not do that automatically
type WebhookHandler func(ctx context.Context, data []byte) error

// webhook represents configuration of getting updates via webhook
type webhook struct {
	updateChanBuffer uint
}

// WebhookOption represents an option that can be applied to webhook
type WebhookOption func(bot *Bot, ctx *webhook) error

// WithWebhookBuffer sets buffering for update chan. Default is 128.
func WithWebhookBuffer(chanBuffer uint) WebhookOption {
	return func(_ *Bot, wh *webhook) error {
		wh.updateChanBuffer = chanBuffer
		return nil
	}
}

// WithWebhookSet calls [Bot.SetWebhook] method before starting webhook
// Note: Calling [Bot.SetWebhook] method multiple times in a row may give "too many requests" errors
func WithWebhookSet(ctx context.Context, params *SetWebhookParams) WebhookOption {
	return func(bot *Bot, _ *webhook) error {
		return bot.SetWebhook(ctx, params)
	}
}

// UpdatesViaWebhook receive updates in chan from webhook. A new handler will be registered on server.
// Calling if already running webhook or long polling will return an error.
func (b *Bot) UpdatesViaWebhook(
	ctx context.Context, registerHandler func(handler WebhookHandler) error, options ...WebhookOption,
) (<-chan Update, error) {
	if err := b.run(runningWebhook); err != nil {
		return nil, err
	}

	wh, err := b.createWebhook(options)
	if err != nil {
		b.running.Store(runningNone)
		return nil, err
	}

	updatesChan := make(chan Update, wh.updateChanBuffer)

	err = registerHandler(func(ctx context.Context, data []byte) error {
		b.log.Debugf("Webhook request with data: %s", string(data))

		var update Update
		err = json.Unmarshal(data, &update)
		if err != nil {
			b.log.Errorf("Webhook decoding error: %s", err)
			return fmt.Errorf("telego: webhook decoding update: %w", err)
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("telego: webhook handler context: %w", ctx.Err())
		case updatesChan <- update.WithContext(ctx):
			return nil
		}
	})
	if err != nil {
		b.running.Store(runningNone)
		return nil, fmt.Errorf("telego: webhook register handler: %w", err)
	}

	go func() {
		<-ctx.Done()
		b.running.Store(runningNone)
		close(updatesChan)
	}()

	return updatesChan, nil
}

func (b *Bot) createWebhook(options []WebhookOption) (*webhook, error) {
	wh := &webhook{
		updateChanBuffer: defaultWebhookUpdateChanBuffer,
	}

	for _, option := range options {
		if err := option(b, wh); err != nil {
			return nil, fmt.Errorf("telego: webhook options: %w", err)
		}
	}

	return wh, nil
}
