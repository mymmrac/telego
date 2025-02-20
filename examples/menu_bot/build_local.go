//go:build !release

package main

import (
	"context"
	"log"
	"net"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func WebhookListener(ctx context.Context) (net.Listener, string) {
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(config.WithForwardsTo(":8080")),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		log.Fatalf("Ngrok listen: %s", err)
	}
	return tun, tun.URL()
}
