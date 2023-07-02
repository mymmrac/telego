package telegohandler

import (
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
	handler := MessageHandler(func(bot *telego.Bot, message telego.Message) { wg.Done() })

	bh.HandleMessage(handler)
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
	handler := MessageHandler(func(bot *telego.Bot, message telego.Message) { wg.Done() })

	bh.HandleEditedMessage(handler)
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
	handler := MessageHandler(func(bot *telego.Bot, message telego.Message) { wg.Done() })

	bh.HandleChannelPost(handler)
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
	handler := MessageHandler(func(bot *telego.Bot, message telego.Message) { wg.Done() })

	bh.HandleEditedChannelPost(handler)
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
	handler := InlineQueryHandler(func(bot *telego.Bot, query telego.InlineQuery) { wg.Done() })

	bh.HandleInlineQuery(handler)
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
	handler := ChosenInlineResultHandler(func(bot *telego.Bot, query telego.ChosenInlineResult) { wg.Done() })

	bh.HandleChosenInlineResult(handler)
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
	handler := CallbackQueryHandler(func(bot *telego.Bot, query telego.CallbackQuery) { wg.Done() })

	bh.HandleCallbackQuery(handler)
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
	handler := ShippingQueryHandler(func(bot *telego.Bot, query telego.ShippingQuery) { wg.Done() })

	bh.HandleShippingQuery(handler)
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
	handler := PreCheckoutQueryHandler(func(bot *telego.Bot, query telego.PreCheckoutQuery) { wg.Done() })

	bh.HandlePreCheckoutQuery(handler)
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
	handler := PollHandler(func(bot *telego.Bot, poll telego.Poll) { wg.Done() })

	bh.HandlePoll(handler)
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
	handler := PollAnswerHandler(func(bot *telego.Bot, pollAnswer telego.PollAnswer) { wg.Done() })

	bh.HandlePollAnswer(handler)
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
	handler := ChatMemberUpdatedHandler(func(bot *telego.Bot, chatMember telego.ChatMemberUpdated) { wg.Done() })

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

func TestBotHandler_HandleChatMemberUpdated(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChatMemberUpdated(nil) })

	wg := &sync.WaitGroup{}
	handler := ChatMemberUpdatedHandler(func(bot *telego.Bot, chatMember telego.ChatMemberUpdated) { wg.Done() })

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

func TestBotHandler_HandleChatJoinRequest(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChatJoinRequest(nil) })

	wg := &sync.WaitGroup{}
	handler := ChatJoinRequestHandler(func(bot *telego.Bot, chatJoinRequest telego.ChatJoinRequest) { wg.Done() })

	bh.HandleChatJoinRequest(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChatJoinRequest: &telego.ChatJoinRequest{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}
