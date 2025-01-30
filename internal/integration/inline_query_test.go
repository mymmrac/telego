//go:build integration && interactive

package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func TestInlineQuery(t *testing.T) {
	ctx := context.Background()

	updates, err := bot.UpdatesViaLongPolling(ctx, &telego.GetUpdatesParams{
		AllowedUpdates: []string{
			telego.InlineQueryUpdates,
			telego.ChosenInlineResultUpdates,
		},
	})
	require.NoError(t, err)

	bh, err := th.NewBotHandler(bot, updates)
	require.NoError(t, err)

	bh.HandleInlineQuery(func(ctx *th.Context, query telego.InlineQuery) error {
		t.Log(query.Query)

		err = ctx.Bot().AnswerInlineQuery(ctx, &telego.AnswerInlineQueryParams{
			InlineQueryID: query.ID,
			Results: []telego.InlineQueryResult{
				tu.ResultArticle("1", "Echo", tu.TextMessage("["+query.Query+"]")).WithDescription(query.Query),
			},
			CacheTime:  1,
			IsPersonal: true,
		})
		if err != nil {
			return fmt.Errorf("answer inline query: %w", err)
		}

		return nil
	})

	bh.HandleChosenInlineResult(func(ctx *th.Context, result telego.ChosenInlineResult) error {
		t.Log(result.Query, result.ResultID)
		return nil
	})

	err = bh.Start()
	require.NoError(t, err)
}
