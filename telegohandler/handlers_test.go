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
		errStart := bh.Start()
		assert.NoError(t, errStart)
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
	err := bh.Stop()
	assert.NoError(t, err)
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

func TestBotHandler_HandleBusinessConnection(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleBusinessConnection(nil) })

	wg := &sync.WaitGroup{}
	handler := BusinessConnectionHandler(func(_ *Context, _ telego.BusinessConnection) error { wg.Done(); return nil })

	bh.HandleBusinessConnection(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{BusinessConnection: &telego.BusinessConnection{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleBusinessMessage(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleBusinessMessage(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandler(func(_ *Context, _ telego.Message) error { wg.Done(); return nil })

	bh.HandleBusinessMessage(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{BusinessMessage: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleEditedBusinessMessage(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleEditedBusinessMessage(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageHandler(func(_ *Context, _ telego.Message) error { wg.Done(); return nil })

	bh.HandleEditedBusinessMessage(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{EditedBusinessMessage: &telego.Message{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleDeletedBusinessMessages(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleDeletedBusinessMessages(nil) })

	wg := &sync.WaitGroup{}
	handler := DeletedBusinessMessagesHandler(func(_ *Context, _ telego.BusinessMessagesDeleted) error {
		wg.Done()
		return nil
	})

	bh.HandleDeletedBusinessMessages(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{DeletedBusinessMessages: &telego.BusinessMessagesDeleted{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleMessageReaction(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleMessageReaction(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageReactionHandler(func(_ *Context, _ telego.MessageReactionUpdated) error { wg.Done(); return nil })

	bh.HandleMessageReaction(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{MessageReaction: &telego.MessageReactionUpdated{}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleMessageReactionCount(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleMessageReactionCount(nil) })

	wg := &sync.WaitGroup{}
	handler := MessageReactionCountHandler(func(_ *Context, _ telego.MessageReactionCountUpdated) error {
		wg.Done()
		return nil
	})

	bh.HandleMessageReactionCount(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{MessageReactionCount: &telego.MessageReactionCountUpdated{}}

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

func TestBotHandler_HandlePurchasedPaidMedia(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandlePurchasedPaidMedia(nil) })

	wg := &sync.WaitGroup{}
	handler := PurchasedPaidMediaHandler(func(_ *Context, _ telego.PaidMediaPurchased) error {
		wg.Done()
		return nil
	})

	bh.HandlePurchasedPaidMedia(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{PurchasedPaidMedia: &telego.PaidMediaPurchased{}}

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

func TestBotHandler_HandleChatBoost(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleChatBoost(nil) })

	wg := &sync.WaitGroup{}
	handler := ChatBoostHandler(func(_ *Context, _ telego.ChatBoostUpdated) error { wg.Done(); return nil })

	bh.HandleChatBoost(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{ChatBoost: &telego.ChatBoostUpdated{
		Boost: telego.ChatBoost{
			Source: &telego.ChatBoostSourcePremium{
				Source: telego.BoostSourcePremium,
			},
		},
	}}

	bh.updates = updates
	testHandler(t, bh, wg)
}

func TestBotHandler_HandleRemovedChatBoost(t *testing.T) {
	bh := newTestBotHandler(t)

	require.Panics(t, func() { bh.HandleRemovedChatBoost(nil) })

	wg := &sync.WaitGroup{}
	handler := RemovedChatBoostHandler(func(_ *Context, _ telego.ChatBoostRemoved) error { wg.Done(); return nil })

	bh.HandleRemovedChatBoost(handler)
	testHandlerSetup(t, bh)

	updates := make(chan telego.Update, 1)
	updates <- telego.Update{RemovedChatBoost: &telego.ChatBoostRemoved{
		Source: &telego.ChatBoostSourcePremium{
			Source: telego.BoostSourcePremium,
		},
	}}

	bh.updates = updates
	testHandler(t, bh, wg)
}
