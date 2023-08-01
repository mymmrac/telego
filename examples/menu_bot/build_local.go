//go:build !release

package main

import (
	"context"
	"log"

	"github.com/fasthttp/router"
	"github.com/mymmrac/telego"
	"github.com/valyala/fasthttp"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func Webhook(ctx context.Context, bot *telego.Bot, secret string) (telego.WebhookServer, string) {
	tun, err := ngrok.Listen(ctx, config.HTTPEndpoint(config.WithForwardsTo(":8080")))
	if err != nil {
		log.Fatalf("Ngrok listen: %s", err)
	}

	srv := &fasthttp.Server{}
	return telego.FuncWebhookServer{
		Server: telego.FastHTTPWebhookServer{
			Logger:      bot.Logger(),
			Server:      srv,
			Router:      router.New(),
			SecretToken: secret,
		},
		StartFunc: func(_ string) error {
			return srv.Serve(tun)
		},
	}, tun.URL()
}
