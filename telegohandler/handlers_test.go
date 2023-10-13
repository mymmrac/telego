package telegohandler

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

func testHandler(t *testing.T, bh *BotHandler, wg *sync.WaitGroup) {
	t.Helper()

	wg.Add(1)

	timeoutSignal := time.After(timeout)
	done := make(chan struct{})

	go bh.Start()

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	select {
	case <-timeoutSignal:
		t.Fatal("Timeout")
	case <-done:
	}
	bh.Stop()
}

func testHandlerSetup(t *testing.T, bh *BotHandler) {
	t.Helper()

	require.Equal(t, 1, len(bh.baseGroup.handlers))
	require.NotNil(t, bh.baseGroup.handlers[0].handler)
	require.NotNil(t, bh.baseGroup.handlers[0].predicates)
	require.Equal(t, 1, len(bh.baseGroup.handlers[0].predicates))
}

func TestBotHandler_HandleMessage(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleMessage(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandler(func(_ *telego.Bot, _ telego.Message) { wg.Done() })

	bh.HandleMessage(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{Message: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleMessageCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleMessageCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.Message) { wg.Done() })

	bh.HandleMessageCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{Message: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleEditedMessage(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleEditedMessage(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandler(func(_ *telego.Bot, _ telego.Message) { wg.Done() })

	bh.HandleEditedMessage(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{EditedMessage: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleEditedMessageCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleEditedMessageCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.Message) { wg.Done() })

	bh.HandleEditedMessageCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{EditedMessage: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleChannelPost(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChannelPost(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandler(func(_ *telego.Bot, _ telego.Message) { wg.Done() })

	bh.HandleChannelPost(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChannelPost: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleChannelPostCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChannelPostCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.Message) { wg.Done() })

	bh.HandleChannelPostCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChannelPost: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleEditedChannelPost(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleEditedChannelPost(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandler(func(_ *telego.Bot, _ telego.Message) { wg.Done() })

	bh.HandleEditedChannelPost(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{EditedChannelPost: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleEditedChannelPostCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleEditedChannelPostCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.Message) { wg.Done() })

	bh.HandleEditedChannelPostCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{EditedChannelPost: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleInlineQuery(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleInlineQuery(nil) })

	wg := &sync.WaitGroup{}
	handler := InlineQueryHandler(func(_ *telego.Bot, _ telego.InlineQuery) { wg.Done() })

	bh.HandleInlineQuery(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{InlineQuery: &telego.InlineQuery{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleInlineQueryCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleInlineQueryCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := InlineQueryHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.InlineQuery) { wg.Done() })

	bh.HandleInlineQueryCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{InlineQuery: &telego.InlineQuery{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleChosenInlineResult(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChosenInlineResult(nil) })

	wg := &sync.WaitGroup{}
	handler := ChosenInlineResultHandler(func(_ *telego.Bot, _ telego.ChosenInlineResult) { wg.Done() })

	bh.HandleChosenInlineResult(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChosenInlineResult: &telego.ChosenInlineResult{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleChosenInlineResultCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChosenInlineResultCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := ChosenInlineResultHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.ChosenInlineResult) {
		wg.Done()
	})

	bh.HandleChosenInlineResultCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChosenInlineResult: &telego.ChosenInlineResult{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleCallbackQuery(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleCallbackQuery(nil) })

	wg := &sync.WaitGroup{}
	handler := CallbackQueryHandler(func(_ *telego.Bot, _ telego.CallbackQuery) { wg.Done() })

	bh.HandleCallbackQuery(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{CallbackQuery: &telego.CallbackQuery{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleCallbackQueryCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleCallbackQueryCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := CallbackQueryHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.CallbackQuery) { wg.Done() })

	bh.HandleCallbackQueryCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{CallbackQuery: &telego.CallbackQuery{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleShippingQuery(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleShippingQuery(nil) })

	wg := &sync.WaitGroup{}
	handler := ShippingQueryHandler(func(_ *telego.Bot, _ telego.ShippingQuery) { wg.Done() })

	bh.HandleShippingQuery(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ShippingQuery: &telego.ShippingQuery{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleShippingQueryCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleShippingQueryCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := ShippingQueryHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.ShippingQuery) { wg.Done() })

	bh.HandleShippingQueryCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ShippingQuery: &telego.ShippingQuery{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandlePreCheckoutQuery(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandlePreCheckoutQuery(nil) })

	wg := &sync.WaitGroup{}
	handler := PreCheckoutQueryHandler(func(_ *telego.Bot, _ telego.PreCheckoutQuery) { wg.Done() })

	bh.HandlePreCheckoutQuery(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{PreCheckoutQuery: &telego.PreCheckoutQuery{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandlePreCheckoutQueryCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandlePreCheckoutQueryCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := PreCheckoutQueryHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.PreCheckoutQuery) {
		wg.Done()
	})

	bh.HandlePreCheckoutQueryCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{PreCheckoutQuery: &telego.PreCheckoutQuery{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandlePoll(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandlePoll(nil) })

	wg := &sync.WaitGroup{}
	handler := PollHandler(func(_ *telego.Bot, _ telego.Poll) { wg.Done() })

	bh.HandlePoll(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{Poll: &telego.Poll{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandlePollCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandlePollCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := PollHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.Poll) { wg.Done() })

	bh.HandlePollCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{Poll: &telego.Poll{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandlePollAnswer(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandlePollAnswer(nil) })

	wg := &sync.WaitGroup{}
	handler := PollAnswerHandler(func(_ *telego.Bot, _ telego.PollAnswer) { wg.Done() })

	bh.HandlePollAnswer(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{PollAnswer: &telego.PollAnswer{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandlePollAnswerCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandlePollAnswerCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := PollAnswerHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.PollAnswer) { wg.Done() })

	bh.HandlePollAnswerCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{PollAnswer: &telego.PollAnswer{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleMyChatMemberUpdated(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleMyChatMemberUpdated(nil) })

	wg := &sync.WaitGroup{}
	handler := ChatMemberUpdatedHandler(func(_ *telego.Bot, _ telego.ChatMemberUpdated) { wg.Done() })

	bh.HandleMyChatMemberUpdated(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{MyChatMember: &telego.ChatMemberUpdated{
		OldChatMember: &telego.ChatMemberMember{Status: telego.MemberStatusMember},
		NewChatMember: &telego.ChatMemberMember{Status: telego.MemberStatusMember},
	}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleMyChatMemberUpdatedCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleMyChatMemberUpdatedCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := ChatMemberUpdatedHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.ChatMemberUpdated) {
		wg.Done()
	})

	bh.HandleMyChatMemberUpdatedCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{MyChatMember: &telego.ChatMemberUpdated{
		OldChatMember: &telego.ChatMemberMember{Status: telego.MemberStatusMember},
		NewChatMember: &telego.ChatMemberMember{Status: telego.MemberStatusMember},
	}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleChatMemberUpdated(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChatMemberUpdated(nil) })

	wg := &sync.WaitGroup{}
	handler := ChatMemberUpdatedHandler(func(_ *telego.Bot, _ telego.ChatMemberUpdated) { wg.Done() })

	bh.HandleChatMemberUpdated(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChatMember: &telego.ChatMemberUpdated{
		OldChatMember: &telego.ChatMemberMember{Status: telego.MemberStatusMember},
		NewChatMember: &telego.ChatMemberMember{Status: telego.MemberStatusMember},
	}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleChatMemberUpdatedCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChatMemberUpdatedCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := ChatMemberUpdatedHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.ChatMemberUpdated) {
		wg.Done()
	})

	bh.HandleChatMemberUpdatedCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChatMember: &telego.ChatMemberUpdated{
		OldChatMember: &telego.ChatMemberMember{Status: telego.MemberStatusMember},
		NewChatMember: &telego.ChatMemberMember{Status: telego.MemberStatusMember},
	}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleChatJoinRequest(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChatJoinRequest(nil) })

	wg := &sync.WaitGroup{}
	handler := ChatJoinRequestHandler(func(_ *telego.Bot, _ telego.ChatJoinRequest) { wg.Done() })

	bh.HandleChatJoinRequest(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChatJoinRequest: &telego.ChatJoinRequest{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleChatJoinRequestCtx(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChatJoinRequestCtx(nil) })

	wg := &sync.WaitGroup{}
	handler := ChatJoinRequestHandlerCtx(func(_ context.Context, _ *telego.Bot, _ telego.ChatJoinRequest) { wg.Done() })

	bh.HandleChatJoinRequestCtx(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChatJoinRequest: &telego.ChatJoinRequest{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}
