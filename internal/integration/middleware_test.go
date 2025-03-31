//go:build integration && interactive

package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func TestMiddleware(t *testing.T) {
	ctx := context.Background()

	updates, err := bot.UpdatesViaLongPolling(ctx, nil)
	require.NoError(t, err)

	bh, err := th.NewBotHandler(bot, updates)
	require.NoError(t, err)

	messages := bh.Group(th.AnyMessageWithFrom())

	const userIDKey = "user-id"
	messages.Use(func(ctx *th.Context, update telego.Update) error {
		t.Log("User ID middleware")
		ctx.WithValue(userIDKey, update.Message.From.ID)
		return ctx.Next(update)
	})

	messages.Handle(
		func(ctx *th.Context, update telego.Update) error {
			userID := ctx.Value(userIDKey).(int64)
			t.Log("User ID:", userID)
			return nil
		},
		func(ctx context.Context, update telego.Update) bool {
			return ctx.Value(userIDKey) != nil
		},
	)

	err = bh.Start()
	require.NoError(t, err)
}
