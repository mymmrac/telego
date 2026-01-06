//go:build integration && webhook

package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	tu "github.com/mymmrac/telego/telegoutil"
)

func TestWebhookInfo(t *testing.T) {
	ctx := t.Context()

	info, err := bot.GetWebhookInfo(ctx)
	require.NoError(t, err)

	t.Logf("WebhookInfo: %+v", info)
}

func TestWebhook(t *testing.T) {
	ctx := t.Context()

	err := bot.SetWebhook(ctx, tu.Webhook("https://example.org"))
	require.NoError(t, err)
}

func TestDeleteWebhook(t *testing.T) {
	ctx := t.Context()

	err := bot.DeleteWebhook(ctx, nil)
	require.NoError(t, err)
}
