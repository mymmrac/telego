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
	updates, err := bot.UpdatesViaLongPolling(nil)
	require.NoError(t, err)

	bh, err := th.NewBotHandler(bot, updates)
	require.NoError(t, err)

	messages := bh.Group(th.AnyMessageWithFrom())

	const userIDKey = "user-id"
	messages.Use(func(bot *telego.Bot, update telego.Update, next th.Handler) {
		t.Log("User ID middleware")
		ctx := context.WithValue(update.Context(), userIDKey, update.Message.From.ID)
		next(bot, update.WithContext(ctx))
	})

	messages.Handle(
		func(bot *telego.Bot, update telego.Update) {
			userID := update.Context().Value(userIDKey).(int64)
			t.Log("User ID:", userID)
		},
		func(update telego.Update) bool {
			return update.Context().Value(userIDKey) != nil
		},
	)

	bh.Start()
}
