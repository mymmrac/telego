package telegohandler

import (
	"context"
	"regexp"
	"strings"

	"github.com/mymmrac/telego"
)

// Any is always true
func Any() Predicate {
	return func(_ context.Context, _ telego.Update) bool {
		return true
	}
}

// None is always false
func None() Predicate {
	return func(_ context.Context, _ telego.Update) bool {
		return false
	}
}

// And is true if all the predicates are true
func And(predicates ...Predicate) Predicate {
	return func(ctx context.Context, update telego.Update) bool {
		for _, p := range predicates {
			if !p(ctx, update) {
				return false
			}
		}
		return true
	}
}

// Or is true if at least one of the predicates is true
func Or(predicates ...Predicate) Predicate {
	return func(ctx context.Context, update telego.Update) bool {
		for _, p := range predicates {
			if p(ctx, update) {
				return true
			}
		}
		return false
	}
}

// Not is true if predicate is false
func Not(predicate Predicate) Predicate {
	return func(ctx context.Context, update telego.Update) bool {
		return !predicate(ctx, update)
	}
}

func anyMessage(message *telego.Message) bool {
	return message != nil
}

// AnyMessage is true if the message isn't nil
func AnyMessage() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMessage(update.Message)
	}
}

func anyMessageWithPhoto(message *telego.Message) bool {
	return message.Photo != nil
}

func anyMessageWithAudio(message *telego.Message) bool {
	return message.Audio != nil
}

func anyMessageWithVideo(message *telego.Message) bool {
	return message.Video != nil
}

func anyMessageWithVideoNote(message *telego.Message) bool {
	return message.VideoNote != nil
}

func anyMessageWithVoice(message *telego.Message) bool {
	return message.Voice != nil
}

func anyMessageWithPaidMedia(message *telego.Message) bool {
	return message.PaidMedia != nil
}
func anyMessageWithDocument(message *telego.Message) bool {
	return message.Document != nil
}

func anyMessageWithContact(message *telego.Message) bool {
	return message.Contact != nil
}

func anyMessageWithGame(message *telego.Message) bool {
	return message.Game != nil
}

func anyMessageWithStory(message *telego.Message) bool {
	return message.Story != nil
}

func anyMessageWithLocation(message *telego.Message) bool {
	return message.Location != nil
}

func anyMessageWithPoll(message *telego.Message) bool {
	return message.Poll != nil
}

func anyMessageWithDice(message *telego.Message) bool {
	return message.Dice != nil
}

func anyMessageWithSticker(message *telego.Message) bool {
	return message.Sticker != nil
}

// AnyMessageWithMedia is true when message contain any media content
func AnyMessageWithMedia() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return (anyMessageWithPhoto(update.Message) || anyMessageWithAudio(update.Message) || anyMessageWithVideo(update.Message) || anyMessageWithVideoNote(update.Message) || anyMessageWithVoice(update.Message) || anyMessageWithPaidMedia(update.Message) || anyMessageWithDocument(update.Message) || anyMessageWithContact(update.Message) || anyMessageWithGame(update.Message) || anyMessageWithStory(update.Message) || anyMessageWithLocation(update.Message) || anyMessageWithPoll(update.Message) || anyMessageWithDice(update.Message) || anyMessageWithSticker(update.Message))
	}

}

func anyMassageWithText(message *telego.Message) bool {
	return message != nil && message.Text != ""
}

// AnyMessageWithText is true if the message isn't nil and its text is not empty
func AnyMessageWithText() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassageWithText(update.Message)
	}
}

func anyMassageWithFrom(message *telego.Message) bool {
	return message != nil && message.From != nil
}

// AnyMessageWithFrom is true if the message isn't nil and its from (sender) is not nil
func AnyMessageWithFrom() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassageWithFrom(update.Message)
	}
}

func baseTextEqual(message *telego.Message, text string) bool {
	return message != nil && message.Text == text
}

// TextEqual is true if the message isn't nil, and its text is equal to the specified text
func TextEqual(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextEqual(update.Message, text)
	}
}

func baseTextEqualFold(message *telego.Message, text string) bool {
	return message != nil && strings.EqualFold(message.Text, text)
}

// TextEqualFold is true if the message isn't nil, and its text equal fold (more general form of case-insensitivity
// equal) to the specified text
func TextEqualFold(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextEqualFold(update.Message, text)
	}
}

func baseTextContains(message *telego.Message, text string) bool {
	return message != nil && strings.Contains(message.Text, text)
}

// TextContains is true if the message isn't nil, and its text contains specified text
func TextContains(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextContains(update.Message, text)
	}
}

func baseTextPrefix(message *telego.Message, prefix string) bool {
	return message != nil && strings.HasPrefix(message.Text, prefix)
}

// TextPrefix is true if the message isn't nil, and its text has specified prefix
func TextPrefix(prefix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextPrefix(update.Message, prefix)
	}
}

func baseTextSuffix(message *telego.Message, suffix string) bool {
	return message != nil && strings.HasSuffix(message.Text, suffix)
}

// TextSuffix is true if the message isn't nil, and its text has specified suffix
func TextSuffix(suffix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextSuffix(update.Message, suffix)
	}
}

func baseTextMatches(message *telego.Message, pattern *regexp.Regexp) bool {
	return message != nil && pattern.MatchString(message.Text)
}

// TextMatches is true if the message isn't nil, and its text matches specified regexp
func TextMatches(pattern *regexp.Regexp) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextMatches(update.Message, pattern)
	}
}

// CommandRegexp matches to command and has match groups on command, bot username and arguments
var CommandRegexp = regexp.MustCompile(`(?s)^/(\w+)(?:@(\w+))?(?:\s+(.+?)\s*)?$`)

// Command match group indexes in the [CommandRegexp]
const (
	CommandMatchCmdGroup         = 1
	CommandMatchBotUsernameGroup = 2
	CommandMatchArgsGroup        = 3
)

// CommandMatchGroupsLen represents the length of match groups in the [CommandRegexp]
const CommandMatchGroupsLen = 4

// AnyCommand is true if the message isn't nil, and it matches to command regexp
func AnyCommand() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.Message != nil && CommandRegexp.MatchString(update.Message.Text)
	}
}

// AnyCommandToMe is true if the message isn't nil, and it matches to command regexp, and the message either in
// private chat, or it contains bot's username or reply to bot's message
// Note: It's better to use Group Privacy Mode (https://core.telegram.org/bots/features#privacy-mode) instead
func AnyCommandToMe(botUsername string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Text)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return update.Message.Chat.Type == telego.ChatTypePrivate ||
			strings.EqualFold(matches[CommandMatchBotUsernameGroup], botUsername) ||
			(update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.From != nil &&
				strings.EqualFold(update.Message.ReplyToMessage.From.Username, botUsername))
	}
}

// CommandEqual is true if the message isn't nil, and it contains specified command
func CommandEqual(command string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Text)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[CommandMatchCmdGroup], command)
	}
}

// CommandEqualArgc is true if the message isn't nil, and it contains specified command with a number of args
func CommandEqualArgc(command string, argc int) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Text)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[CommandMatchCmdGroup], command) &&
			(argc == 0 && matches[CommandMatchArgsGroup] == "" ||
				len(strings.Fields(matches[CommandMatchArgsGroup])) == argc)
	}
}

// CommandEqualArgv is true if the message isn't nil, and it contains specified command and args
func CommandEqualArgv(command string, argv ...string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Text)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[CommandMatchCmdGroup], command) &&
			(len(argv) == 0 && matches[CommandMatchArgsGroup] == "" ||
				matches[CommandMatchArgsGroup] == strings.Join(argv, " "))
	}
}

// SuccessPayment is true if the message isn't nil, and contains success payment
func SuccessPayment() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.Message != nil && update.Message.SuccessfulPayment != nil
	}
}

// AnyEditedMessage is true if the edited message isn't nil
func AnyEditedMessage() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassage(update.EditedMessage)
	}
}

// AnyEditedMessageWithText is true if the edited message isn't nil and its text is not empty
func AnyEditedMessageWithText() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassageWithText(update.EditedMessage)
	}
}

// AnyEditedMessageWithFrom is true if the edited message isn't nil and its from (sender) is not nil
func AnyEditedMessageWithFrom() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassageWithFrom(update.EditedMessage)
	}
}

// EditedTextEqual is true if the edited message isn't nil, and its text equals to the specified text
func EditedTextEqual(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextEqual(update.EditedMessage, text)
	}
}

// EditedTextEqualFold is true if the edited message isn't nil, and its text equal fold (more general form of
// case-insensitivity equal) to the specified text
func EditedTextEqualFold(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextEqualFold(update.EditedMessage, text)
	}
}

// EditedTextContains is true if the edited message isn't nil, and its text contains specified text
func EditedTextContains(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextContains(update.EditedMessage, text)
	}
}

// EditedTextPrefix is true if the edited message isn't nil, and its text has specified prefix
func EditedTextPrefix(prefix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextPrefix(update.EditedMessage, prefix)
	}
}

// EditedTextSuffix is true if the edited message isn't nil, and its text has specified suffix
func EditedTextSuffix(suffix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextSuffix(update.EditedMessage, suffix)
	}
}

// EditedTextMatches is true if the edited message isn't nil, and its text matches specified regexp
func EditedTextMatches(pattern *regexp.Regexp) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextMatches(update.EditedMessage, pattern)
	}
}

// AnyChannelPost is true if channel post isn't nil
func AnyChannelPost() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassage(update.ChannelPost)
	}
}

// AnyChannelPostWithText is true if channel post isn't nil and its text is not empty
func AnyChannelPostWithText() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassageWithText(update.ChannelPost)
	}
}

// PostTextEqual is true if channel post isn't nil, and its text equals to the specified text
func PostTextEqual(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextEqual(update.ChannelPost, text)
	}
}

// PostTextEqualFold is true if channel post isn't nil, and its text equal fold (more general form of case-insensitivity
// equal) to the specified text
func PostTextEqualFold(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextEqualFold(update.ChannelPost, text)
	}
}

// PostTextContains is true if channel post isn't nil, and its text contains specified text
func PostTextContains(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextContains(update.ChannelPost, text)
	}
}

// PostTextPrefix is true if channel post isn't nil, and its text has specified prefix
func PostTextPrefix(prefix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextPrefix(update.ChannelPost, prefix)
	}
}

// PostTextSuffix is true if channel post isn't nil, and its text has specified suffix
func PostTextSuffix(suffix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextSuffix(update.ChannelPost, suffix)
	}
}

// PostTextMatches is true if channel post isn't nil, and its text matches specified regexp
func PostTextMatches(pattern *regexp.Regexp) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextMatches(update.ChannelPost, pattern)
	}
}

// AnyEditedChannelPost is true if the edited channel post isn't nil
func AnyEditedChannelPost() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassage(update.EditedChannelPost)
	}
}

// AnyEditedChannelPostWithText is true if edited channel post isn't nil and its text is not empty
func AnyEditedChannelPostWithText() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassageWithText(update.EditedChannelPost)
	}
}

// EditedPostTextEqual is true if edited channel post isn't nil, and its text equals to the specified text
func EditedPostTextEqual(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextEqual(update.EditedChannelPost, text)
	}
}

// EditedPostTextEqualFold is true if edited channel post isn't nil, and its text equal fold (more general form of
// case-insensitivity equal) to the specified text
func EditedPostTextEqualFold(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextEqualFold(update.EditedChannelPost, text)
	}
}

// EditedPostTextContains is true if edited channel post isn't nil, and its text contains specified text
func EditedPostTextContains(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextContains(update.EditedChannelPost, text)
	}
}

// EditedPostTextPrefix is true if edited channel post isn't nil, and its text has specified prefix
func EditedPostTextPrefix(prefix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextPrefix(update.EditedChannelPost, prefix)
	}
}

// EditedPostTextSuffix is true if edited channel post isn't nil, and its text has specified suffix
func EditedPostTextSuffix(suffix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextSuffix(update.EditedChannelPost, suffix)
	}
}

// EditedPostTextMatches is true if edited channel post isn't nil, and its text matches specified regexp
func EditedPostTextMatches(pattern *regexp.Regexp) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseTextMatches(update.EditedChannelPost, pattern)
	}
}

// AnyBusinessConnection is true if business connection isn't nil
func AnyBusinessConnection() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.BusinessConnection != nil
	}
}

// AnyBusinessMessage is true if the business message isn't nil
func AnyBusinessMessage() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassage(update.BusinessMessage)
	}
}

// AnyEditedBusinessMessage is true if edited business message isn't nil
func AnyEditedBusinessMessage() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassage(update.EditedBusinessMessage)
	}
}

// AnyDeletedBusinessMessages is true if deleted business messages isn't nil
func AnyDeletedBusinessMessages() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.DeletedBusinessMessages != nil
	}
}

// AnyMessageReaction is true if message reaction isn't nil
func AnyMessageReaction() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.MessageReaction != nil
	}
}

// AnyMessageReactionCount is true if message reaction count isn't nil
func AnyMessageReactionCount() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.MessageReactionCount != nil
	}
}

// AnyInlineQuery is true if inline query isn't nil
func AnyInlineQuery() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.InlineQuery != nil
	}
}

// InlineQueryEqual is true if inline query isn't nil, and its query equal to specified text
func InlineQueryEqual(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.InlineQuery != nil && update.InlineQuery.Query == text
	}
}

// InlineQueryEqualFold is true if inline query isn't nil, and its query equal fold (more general form of
// case-insensitivity equal) to the specified text
func InlineQueryEqualFold(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.InlineQuery != nil && strings.EqualFold(update.InlineQuery.Query, text)
	}
}

// InlineQueryContains is true if inline query isn't nil, and its query contains specified text
func InlineQueryContains(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.InlineQuery != nil && strings.Contains(update.InlineQuery.Query, text)
	}
}

// InlineQueryPrefix is true if inline query isn't nil, and its query has specified prefix
func InlineQueryPrefix(prefix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.InlineQuery != nil && strings.HasPrefix(update.InlineQuery.Query, prefix)
	}
}

// InlineQuerySuffix is true if inline query isn't nil, and its query has specified suffix
func InlineQuerySuffix(suffix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.InlineQuery != nil && strings.HasSuffix(update.InlineQuery.Query, suffix)
	}
}

// InlineQueryMatches is true if inline query isn't nil, and its query matches specified regexp
func InlineQueryMatches(pattern *regexp.Regexp) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.InlineQuery != nil && pattern.MatchString(update.InlineQuery.Query)
	}
}

// AnyChosenInlineResult is true if the chosen inline result isn't nil
func AnyChosenInlineResult() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.ChosenInlineResult != nil
	}
}

// AnyCallbackQuery is true if the callback query isn't nil
func AnyCallbackQuery() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.CallbackQuery != nil
	}
}

// AnyCallbackQueryWithMessage is true if callback query and its message isn't nil
func AnyCallbackQueryWithMessage() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.CallbackQuery != nil && update.CallbackQuery.Message != nil
	}
}

// CallbackDataEqual is true if callback query isn't nil, and its data equal to specified text
func CallbackDataEqual(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.CallbackQuery != nil && update.CallbackQuery.Data == text
	}
}

// CallbackDataEqualFold is true if the callback query isn't nil, and its data equal fold (more general form of
// case-insensitivity equal) to the specified text
func CallbackDataEqualFold(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.CallbackQuery != nil && strings.EqualFold(update.CallbackQuery.Data, text)
	}
}

// CallbackDataContains is true if the callback query isn't nil, and its data contains specified text
func CallbackDataContains(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.CallbackQuery != nil && strings.Contains(update.CallbackQuery.Data, text)
	}
}

// CallbackDataPrefix is true if the callback query isn't nil, and its data has specified prefix
func CallbackDataPrefix(prefix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.CallbackQuery != nil && strings.HasPrefix(update.CallbackQuery.Data, prefix)
	}
}

// CallbackDataSuffix is true if the callback query isn't nil, and its data has specified suffix
func CallbackDataSuffix(suffix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.CallbackQuery != nil && strings.HasSuffix(update.CallbackQuery.Data, suffix)
	}
}

// CallbackDataMatches is true if the callback query isn't nil, and its data matches specified regexp
func CallbackDataMatches(pattern *regexp.Regexp) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.CallbackQuery != nil && pattern.MatchString(update.CallbackQuery.Data)
	}
}

// AnyShippingQuery is true if shipping query isn't nil
func AnyShippingQuery() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.ShippingQuery != nil
	}
}

// AnyPreCheckoutQuery is true if the pre checkout query isn't nil
func AnyPreCheckoutQuery() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.PreCheckoutQuery != nil
	}
}

// AnyPurchasedPaidMedia is true if the purchased paid media isn't nil
func AnyPurchasedPaidMedia() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.PurchasedPaidMedia != nil
	}
}

// AnyPoll is true if the poll isn't nil
func AnyPoll() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.Poll != nil
	}
}

// AnyPollAnswer is true if the poll answer isn't nil
func AnyPollAnswer() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.PollAnswer != nil
	}
}

// AnyMyChatMember is true if my chat member isn't nil
func AnyMyChatMember() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.MyChatMember != nil
	}
}

// AnyChatMember is true if chat member isn't nil
func AnyChatMember() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.ChatMember != nil
	}
}

// AnyChatJoinRequest is true if chat join request isn't nil
func AnyChatJoinRequest() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.ChatJoinRequest != nil
	}
}

// AnyChatBoost is true if chat boost isn't nil
func AnyChatBoost() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.ChatBoost != nil
	}
}

// AnyRemovedChatBoost is true if removed chat boost isn't nil
func AnyRemovedChatBoost() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.RemovedChatBoost != nil
	}
}

func anyMassageWithCaption(message *telego.Message) bool {
	return message != nil && message.Caption != ""
}

// AnyMessageWithCaption is true if the message isn't nil and its caption is not empty
func AnyMessageWithCaption() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassageWithCaption(update.Message)
	}
}

func baseCaptionEqual(message *telego.Message, text string) bool {
	return message != nil && message.Caption == text
}

// CaptionEqual is true if the message isn't nil, and its caption is equal to the specified text
func CaptionEqual(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionEqual(update.Message, text)
	}
}

func baseCaptionEqualFold(message *telego.Message, text string) bool {
	return message != nil && strings.EqualFold(message.Caption, text)
}

// CaptionEqualFold is true if the message isn't nil, and its caption equal fold (more general form of
// case-insensitivity equal) to the specified text
func CaptionEqualFold(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionEqualFold(update.Message, text)
	}
}

func baseCaptionContains(message *telego.Message, text string) bool {
	return message != nil && strings.Contains(message.Caption, text)
}

// CaptionContains is true if the message isn't nil, and its caption contains specified text
func CaptionContains(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionContains(update.Message, text)
	}
}

func baseCaptionPrefix(message *telego.Message, prefix string) bool {
	return message != nil && strings.HasPrefix(message.Caption, prefix)
}

// CaptionPrefix is true if the message isn't nil, and its caption has specified prefix
func CaptionPrefix(prefix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionPrefix(update.Message, prefix)
	}
}

func baseCaptionSuffix(message *telego.Message, suffix string) bool {
	return message != nil && strings.HasSuffix(message.Caption, suffix)
}

// CaptionSuffix is true if the message isn't nil, and its caption has specified suffix
func CaptionSuffix(suffix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionSuffix(update.Message, suffix)
	}
}

func baseCaptionMatches(message *telego.Message, pattern *regexp.Regexp) bool {
	return message != nil && pattern.MatchString(message.Caption)
}

// CaptionMatches is true if the message isn't nil, and its caption matches specified regexp
func CaptionMatches(pattern *regexp.Regexp) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionMatches(update.Message, pattern)
	}
}

// AnyCaptionCommand is true if the message isn't nil, and its caption matches to command regexp
func AnyCaptionCommand() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return update.Message != nil && CommandRegexp.MatchString(update.Message.Caption)
	}
}

// CaptionCommandEqual is true if the message isn't nil, and its caption contains specified command
func CaptionCommandEqual(command string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Caption)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[CommandMatchCmdGroup], command)
	}
}

// CaptionCommandEqualArgc is true if the message isn't nil, and its caption contains specified
// command with a number of args
func CaptionCommandEqualArgc(command string, argc int) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Caption)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[CommandMatchCmdGroup], command) &&
			(argc == 0 && matches[CommandMatchArgsGroup] == "" ||
				len(strings.Fields(matches[CommandMatchArgsGroup])) == argc)
	}
}

// CaptionCommandEqualArgv is true if the message isn't nil, and its caption contains specified command and args
func CaptionCommandEqualArgv(command string, argv ...string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Caption)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[CommandMatchCmdGroup], command) &&
			(len(argv) == 0 && matches[CommandMatchArgsGroup] == "" ||
				matches[CommandMatchArgsGroup] == strings.Join(argv, " "))
	}
}

// AnyEditedMessageWithCaption is true if the edited message isn't nil and its caption is not empty
func AnyEditedMessageWithCaption() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassageWithCaption(update.EditedMessage)
	}
}

// EditedCaptionEqual is true if the edited message isn't nil, and its caption equals to the specified text
func EditedCaptionEqual(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionEqual(update.EditedMessage, text)
	}
}

// EditedCaptionEqualFold is true if the edited message isn't nil, and its caption equal fold (more general form of
// case-insensitivity equal) to the specified text
func EditedCaptionEqualFold(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionEqualFold(update.EditedMessage, text)
	}
}

// EditedCaptionContains is true if the edited message isn't nil, and its caption contains specified text
func EditedCaptionContains(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionContains(update.EditedMessage, text)
	}
}

// EditedCaptionPrefix is true if the edited message isn't nil, and its caption has specified prefix
func EditedCaptionPrefix(prefix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionPrefix(update.EditedMessage, prefix)
	}
}

// EditedCaptionSuffix is true if the edited message isn't nil, and its caption has specified suffix
func EditedCaptionSuffix(suffix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionSuffix(update.EditedMessage, suffix)
	}
}

// EditedCaptionMatches is true if the edited message isn't nil, and its caption matches specified regexp
func EditedCaptionMatches(pattern *regexp.Regexp) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionMatches(update.EditedMessage, pattern)
	}
}

// AnyChannelPostWithCaption is true if channel post isn't nil and its caption is not empty
func AnyChannelPostWithCaption() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassageWithCaption(update.ChannelPost)
	}
}

// PostCaptionEqual is true if channel post isn't nil, and its caption equals to the specified text
func PostCaptionEqual(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionEqual(update.ChannelPost, text)
	}
}

// PostCaptionEqualFold is true if channel post isn't nil, and its caption equal fold (more general form of
// case-insensitivity equal) to the specified text
func PostCaptionEqualFold(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionEqualFold(update.ChannelPost, text)
	}
}

// PostCaptionContains is true if channel post isn't nil, and its caption contains specified text
func PostCaptionContains(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionContains(update.ChannelPost, text)
	}
}

// PostCaptionPrefix is true if channel post isn't nil, and its caption has specified prefix
func PostCaptionPrefix(prefix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionPrefix(update.ChannelPost, prefix)
	}
}

// PostCaptionSuffix is true if channel post isn't nil, and its caption has specified suffix
func PostCaptionSuffix(suffix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionSuffix(update.ChannelPost, suffix)
	}
}

// PostCaptionMatches is true if channel post isn't nil, and its caption matches specified regexp
func PostCaptionMatches(pattern *regexp.Regexp) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionMatches(update.ChannelPost, pattern)
	}
}

// AnyEditedChannelPostWithCaption is true if edited channel post isn't nil and its caption is not empty
func AnyEditedChannelPostWithCaption() Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return anyMassageWithCaption(update.EditedChannelPost)
	}
}

// EditedPostCaptionEqual is true if edited channel post isn't nil, and its caption equals to the specified text
func EditedPostCaptionEqual(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionEqual(update.EditedChannelPost, text)
	}
}

// EditedPostCaptionEqualFold is true if edited channel post isn't nil, and its caption equal fold (more general form of
// case-insensitivity equal) to the specified text
func EditedPostCaptionEqualFold(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionEqualFold(update.EditedChannelPost, text)
	}
}

// EditedPostCaptionContains is true if edited channel post isn't nil, and its caption contains specified text
func EditedPostCaptionContains(text string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionContains(update.EditedChannelPost, text)
	}
}

// EditedPostCaptionPrefix is true if edited channel post isn't nil, and its caption has specified prefix
func EditedPostCaptionPrefix(prefix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionPrefix(update.EditedChannelPost, prefix)
	}
}

// EditedPostCaptionSuffix is true if edited channel post isn't nil, and its caption has specified suffix
func EditedPostCaptionSuffix(suffix string) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionSuffix(update.EditedChannelPost, suffix)
	}
}

// EditedPostCaptionMatches is true if edited channel post isn't nil, and its caption matches specified regexp
func EditedPostCaptionMatches(pattern *regexp.Regexp) Predicate {
	return func(_ context.Context, update telego.Update) bool {
		return baseCaptionMatches(update.EditedChannelPost, pattern)
	}
}
