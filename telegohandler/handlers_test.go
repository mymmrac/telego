package telegohandler

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

func testHandler(t *testing.T, bh *BotHandler, wg *sync.WaitGroup) {
	t.Helper()

	wg.Add(1)

	timeoutSignal := time.After(timeout)
	done := make(chan struct{})

	go func() {
		err := bh.Start()
		assert.NoError(t, err)
	}()

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

	require.Len(t, bh.baseGroup.routes, 1)
	require.NotNil(t, bh.baseGroup.routes[0].handler)
	require.NotNil(t, bh.baseGroup.routes[0].predicates)
	require.Len(t, bh.baseGroup.routes[0].predicates, 1)
}

func TestBotHandler_HandleMessage(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleMessage(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandler(func(_ *Context, _ telego.Message) error { wg.Done(); return nil })

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
	handler := MessageHandler(func(_ *Context, _ telego.Message) error { wg.Done(); return nil })

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
	handler := MessageHandler(func(_ *Context, _ telego.Message) error { wg.Done(); return nil })

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
	handler := MessageHandler(func(_ *Context, _ telego.Message) error { wg.Done(); return nil })

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
	handler := InlineQueryHandler(func(_ *Context, _ telego.InlineQuery) error { wg.Done(); return nil })

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
	handler := ChosenInlineResultHandler(func(_ *Context, _ telego.ChosenInlineResult) error {
		wg.Done()
		return nil
	})

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
	handler := CallbackQueryHandler(func(_ *Context, _ telego.CallbackQuery) error { wg.Done(); return nil })

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
	handler := ShippingQueryHandler(func(_ *Context, _ telego.ShippingQuery) error { wg.Done(); return nil })

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
	handler := PreCheckoutQueryHandler(func(_ *Context, _ telego.PreCheckoutQuery) error {
		wg.Done()
		return nil
	})

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
	handler := PollHandler(func(_ *Context, _ telego.Poll) error { wg.Done(); return nil })

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
	handler := PollAnswerHandler(func(_ *Context, _ telego.PollAnswer) error { wg.Done(); return nil })

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
	handler := ChatMemberUpdatedHandler(func(_ *Context, _ telego.ChatMemberUpdated) error {
		wg.Done()
		return nil
	})

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
	handler := ChatMemberUpdatedHandler(func(_ *Context, _ telego.ChatMemberUpdated) error {
		wg.Done()
		return nil
	})

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
	handler := ChatJoinRequestHandler(func(_ *Context, _ telego.ChatJoinRequest) error { wg.Done(); return nil })

	bh.HandleChatJoinRequest(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChatJoinRequest: &telego.ChatJoinRequest{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}
