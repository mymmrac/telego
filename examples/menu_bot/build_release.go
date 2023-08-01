//go:build release

package main

import (
	"context"

	"github.com/fasthttp/router"
	"github.com/mymmrac/telego"
	"github.com/valyala/fasthttp"
)

func Webhook(_ context.Context, bot *telego.Bot, secret string) (telego.WebhookServer, string) {
	return telego.FastHTTPWebhookServer{
		Logger:      bot.Logger(),
		Server:      &fasthttp.Server{},
		Router:      router.New(),
		SecretToken: secret,
	}, "https://example.org"
}
