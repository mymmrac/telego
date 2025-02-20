//go:build release

package main

import (
	"context"
	"log"
	"net"
)

func WebhookListener(_ context.Context) (net.Listener, string) {
	ln, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatalf("Listen: %s", err)
	}
	return ln, "https://example.org"
}
