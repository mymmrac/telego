//go:build integration && interactive

package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func TestGetMessage(t *testing.T) {
	ctx := t.Context()

	updates, err := bot.UpdatesViaLongPolling(ctx, &telego.GetUpdatesParams{
		AllowedUpdates: []string{
			telego.MessageUpdates,
		},
	})
	require.NoError(t, err)

	bh, err := th.NewBotHandler(bot, updates)
	require.NoError(t, err)

	bh.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		var data []byte
		data, err = json.Marshal(message)
		require.NoError(t, err)

		t.Log(string(data))

		return nil
	})

	err = bh.Start()
	require.NoError(t, err)
}
