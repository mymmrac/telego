package telegoflow

import "github.com/mymmrac/telego"

// KeyFunc extracts a session key from an update.
type KeyFunc func(update telego.Update) (SessionKey, bool)

// DefaultKeyFunc extracts a key from message-like updates and callback queries.
//
// The key is chat scoped and user scoped: chat_id:user_id.
func DefaultKeyFunc(update telego.Update) (SessionKey, bool) {
	if update.Message != nil && update.Message.From != nil {
		return SessionKey{ChatID: update.Message.Chat.ID, UserID: update.Message.From.ID}, true
	}
	if update.EditedMessage != nil && update.EditedMessage.From != nil {
		return SessionKey{ChatID: update.EditedMessage.Chat.ID, UserID: update.EditedMessage.From.ID}, true
	}
	if update.BusinessMessage != nil && update.BusinessMessage.From != nil {
		return SessionKey{ChatID: update.BusinessMessage.Chat.ID, UserID: update.BusinessMessage.From.ID}, true
	}
	if update.EditedBusinessMessage != nil && update.EditedBusinessMessage.From != nil {
		return SessionKey{
			ChatID: update.EditedBusinessMessage.Chat.ID,
			UserID: update.EditedBusinessMessage.From.ID,
		}, true
	}
	if update.GuestMessage != nil && update.GuestMessage.From != nil {
		return SessionKey{ChatID: update.GuestMessage.Chat.ID, UserID: update.GuestMessage.From.ID}, true
	}
	if update.CallbackQuery != nil && update.CallbackQuery.Message != nil {
		return SessionKey{
			ChatID: update.CallbackQuery.Message.GetChat().ID,
			UserID: update.CallbackQuery.From.ID,
		}, true
	}

	return SessionKey{}, false
}
