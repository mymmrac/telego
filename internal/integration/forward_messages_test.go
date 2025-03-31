//go:build integration && interactive

package integration

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func TestForwardMessages(t *testing.T) {
	ctx := context.Background()

	updates, err := bot.UpdatesViaLongPolling(ctx, &telego.GetUpdatesParams{
		AllowedUpdates: []string{
			telego.MessageUpdates,
		},
	})
	require.NoError(t, err)

	bh, err := th.NewBotHandler(bot, updates)
	require.NoError(t, err)

	lock := sync.Mutex{}
	groupedMessages := map[string]chan int{}

	bh.HandleMessage(func(ctx *th.Context, message telego.Message) error {
		if message.MediaGroupID == "" {
			_, err = ctx.Bot().ForwardMessage(ctx, &telego.ForwardMessageParams{
				ChatID:          message.Chat.ChatID(),
				MessageThreadID: message.MessageThreadID,
				FromChatID:      message.Chat.ChatID(),
				MessageID:       message.MessageID,
			})
			if err != nil {
				return err
			}
		} else {
			lock.Lock()
			defer lock.Unlock()

			var ok bool
			var ch chan int

			if ch, ok = groupedMessages[message.MediaGroupID]; ok {
				ch <- message.MessageID
				return nil
			}

			ch = make(chan int, 1)
			groupedMessages[message.MediaGroupID] = ch

			go func() {
				messageIDs := make([]int, 0, 10)
				// Add the initial message
				messageIDs = append(messageIDs, message.MessageID)

			loop:
				for {
					// Limit the number of messages to 10 (Telegram limit)
					if len(messageIDs) == 10 {
						break
					}

					select {
					case messageID := <-ch:
						messageIDs = append(messageIDs, messageID)
					// Wait for 1 second for other messages in a group
					case <-time.After(time.Second):
						break loop
					}
				}

				lock.Lock()
				delete(groupedMessages, message.MediaGroupID)
				lock.Unlock()

				close(ch)
				// Drain message IDs if any left
				for messageID := range ch {
					messageIDs = append(messageIDs, messageID)
				}

				// Sort message IDs to preserve the order
				slices.Sort(messageIDs)

				_, err = ctx.Bot().ForwardMessages(ctx.WithoutCancel(), &telego.ForwardMessagesParams{
					ChatID:          message.Chat.ChatID(),
					MessageThreadID: message.MessageThreadID,
					FromChatID:      message.Chat.ChatID(),
					MessageIDs:      messageIDs,
				})
				if err != nil {
					fmt.Println(err)
				}
			}()
		}
		return nil
	})

	err = bh.Start()
	require.NoError(t, err)
}
