package telegohandler

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego"
)

const (
	testText        = `Test text`
	testTextLower   = `test text`
	testTextPart    = `t te`
	testTextPrefix  = `Test`
	testTextSuffix  = `text`
	testCommand1    = `/test abc 123`
	testCommand2    = `/hmm bcd`
	testCommandName = `test`
)

//nolint:funlen,maintidx
func TestPredicates(t *testing.T) {
	tests := []struct {
		name      string
		predicate Predicate
		update    telego.Update
		matches   bool
	}{
		{
			name: "union_matches",
			predicate: Union(
				func(update telego.Update) bool { return true },
				func(update telego.Update) bool { return false },
			),
			update:  telego.Update{},
			matches: true,
		},
		{
			name: "union_not_matches",
			predicate: Union(
				func(update telego.Update) bool { return false },
				func(update telego.Update) bool { return false },
			),
			update:  telego.Update{},
			matches: false,
		},
		{
			name:      "not_matches",
			predicate: Not(func(update telego.Update) bool { return false }),
			update:    telego.Update{},
			matches:   true,
		},
		{
			name:      "not_not_matches",
			predicate: Not(func(update telego.Update) bool { return true }),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "any_message_matches",
			predicate: AnyMessage(),
			update:    telego.Update{Message: &telego.Message{}},
			matches:   true,
		},
		{
			name:      "any_message_not_matches",
			predicate: AnyMessage(),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "text_equal_matches",
			predicate: TextEqual(testText),
			update:    telego.Update{Message: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "text_equal_not_matches",
			predicate: TextEqual(testText),
			update:    telego.Update{Message: &telego.Message{Text: testTextLower}},
			matches:   false,
		},
		{
			name:      "text_equal_fold_matches",
			predicate: TextEqualFold(testText),
			update:    telego.Update{Message: &telego.Message{Text: testTextLower}},
			matches:   true,
		},
		{
			name:      "text_equal_fold_not_matches",
			predicate: TextEqualFold(testText),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "text_contains_matches",
			predicate: TextContains(testTextPart),
			update:    telego.Update{Message: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "text_contains_not_matches",
			predicate: TextContains(testTextPart),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "text_prefix_matches",
			predicate: TextPrefix(testTextPrefix),
			update:    telego.Update{Message: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "text_prefix_not_matches",
			predicate: TextPrefix(testTextPrefix),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "text_suffix_matches",
			predicate: TextSuffix(testTextSuffix),
			update:    telego.Update{Message: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "text_suffix_not_matches",
			predicate: TextSuffix(testTextSuffix),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "text_matches_matches",
			predicate: TextMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{Message: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "text_matches_not_matches",
			predicate: TextMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "any_command_matches",
			predicate: AnyCommand(),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   true,
		},
		{
			name:      "any_command_not_matches",
			predicate: AnyCommand(),
			update:    telego.Update{Message: &telego.Message{Text: testText}},
			matches:   false,
		},
		{
			name:      "command_equal_matches",
			predicate: CommandEqual(testCommandName),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   true,
		},
		{
			name:      "command_equal_not_matches",
			predicate: CommandEqual(testCommandName),
			update:    telego.Update{Message: &telego.Message{Text: testCommand2}},
			matches:   false,
		},
		{
			name:      "command_equal_no_message",
			predicate: CommandEqual(testCommandName),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "command_equal_no_command",
			predicate: CommandEqual(testCommandName),
			update:    telego.Update{Message: &telego.Message{Text: testText}},
			matches:   false,
		},
		{
			name:      "command_equal_argc_matches",
			predicate: CommandEqualArgc(testCommandName, 2),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   true,
		},
		{
			name:      "command_equal_argc_not_matches",
			predicate: CommandEqualArgc(testCommandName, 3),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "command_equal_argc_no_message",
			predicate: CommandEqualArgc(testCommandName, 0),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "command_equal_argc_no_command",
			predicate: CommandEqualArgc(testCommandName, 0),
			update:    telego.Update{Message: &telego.Message{Text: testText}},
			matches:   false,
		},
		{
			name:      "command_equal_argv_matches",
			predicate: CommandEqualArgv(testCommandName, "abc", "123"),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   true,
		},
		{
			name:      "command_equal_argv_not_matches",
			predicate: CommandEqualArgv(testCommandName, "abc", "abc"),
			update:    telego.Update{Message: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "command_equal_argv_no_message",
			predicate: CommandEqualArgv(testCommandName),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "command_equal_argv_no_command",
			predicate: CommandEqualArgv(testCommandName),
			update:    telego.Update{Message: &telego.Message{Text: testText}},
			matches:   false,
		},
		{
			name:      "any_edited_message_matches",
			predicate: AnyEditedMessage(),
			update:    telego.Update{EditedMessage: &telego.Message{}},
			matches:   true,
		},
		{
			name:      "any_edited_message_not_matches",
			predicate: AnyEditedMessage(),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "edited_text_equal_matches",
			predicate: EditedTextEqual(testText),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "edited_text_equal_not_matches",
			predicate: EditedTextEqual(testText),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testTextLower}},
			matches:   false,
		},
		{
			name:      "edited_text_equal_fold_matches",
			predicate: EditedTextEqualFold(testText),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testTextLower}},
			matches:   true,
		},
		{
			name:      "edited_text_equal_fold_not_matches",
			predicate: EditedTextEqualFold(testText),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "edited_text_contains_matches",
			predicate: EditedTextContains(testTextPart),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "edited_text_contains_not_matches",
			predicate: EditedTextContains(testTextPart),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "edited_text_prefix_matches",
			predicate: EditedTextPrefix(testTextPrefix),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "edited_text_prefix_not_matches",
			predicate: EditedTextPrefix(testTextPrefix),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "edited_text_suffix_matches",
			predicate: EditedTextSuffix(testTextSuffix),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "edited_text_suffix_not_matches",
			predicate: EditedTextSuffix(testTextSuffix),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "edited_text_matches_matches",
			predicate: EditedTextMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "edited_text_matches_not_matches",
			predicate: EditedTextMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{EditedMessage: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "any_channel_post_matches",
			predicate: AnyChannelPost(),
			update:    telego.Update{ChannelPost: &telego.Message{}},
			matches:   true,
		},
		{
			name:      "any_channel_post_not_matches",
			predicate: AnyChannelPost(),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "post_text_equal_matches",
			predicate: PostTextEqual(testText),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "post_text_equal_not_matches",
			predicate: PostTextEqual(testText),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testTextLower}},
			matches:   false,
		},
		{
			name:      "post_text_equal_fold_matches",
			predicate: PostTextEqualFold(testText),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testTextLower}},
			matches:   true,
		},
		{
			name:      "post_text_equal_fold_not_matches",
			predicate: PostTextEqualFold(testText),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "post_text_contains_matches",
			predicate: PostTextContains(testTextPart),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "post_text_contains_not_matches",
			predicate: PostTextContains(testTextPart),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "post_text_prefix_matches",
			predicate: PostTextPrefix(testTextPrefix),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "post_text_prefix_not_matches",
			predicate: PostTextPrefix(testTextPrefix),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "post_text_suffix_matches",
			predicate: PostTextSuffix(testTextSuffix),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "post_text_suffix_not_matches",
			predicate: PostTextSuffix(testTextSuffix),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "post_text_matches_matches",
			predicate: PostTextMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "post_text_matches_not_matches",
			predicate: PostTextMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{ChannelPost: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "any_edited_channel_post_matches",
			predicate: AnyEditedChannelPost(),
			update:    telego.Update{EditedChannelPost: &telego.Message{}},
			matches:   true,
		},
		{
			name:      "any_edited_channel_post_not_matches",
			predicate: AnyEditedChannelPost(),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "edited_post_text_equal_matches",
			predicate: EditedPostTextEqual(testText),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "edited_post_text_equal_not_matches",
			predicate: EditedPostTextEqual(testText),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testTextLower}},
			matches:   false,
		},
		{
			name:      "edited_post_text_equal_fold_matches",
			predicate: EditedPostTextEqualFold(testText),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testTextLower}},
			matches:   true,
		},
		{
			name:      "edited_post_text_equal_fold_not_matches",
			predicate: EditedPostTextEqualFold(testText),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "edited_post_text_contains_matches",
			predicate: EditedPostTextContains(testTextPart),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "edited_post_text_contains_not_matches",
			predicate: EditedPostTextContains(testTextPart),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "edited_post_text_prefix_matches",
			predicate: EditedPostTextPrefix(testTextPrefix),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "edited_post_text_prefix_not_matches",
			predicate: EditedPostTextPrefix(testTextPrefix),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "edited_post_text_suffix_matches",
			predicate: EditedPostTextSuffix(testTextSuffix),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "edited_post_text_suffix_not_matches",
			predicate: EditedPostTextSuffix(testTextSuffix),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "edited_post_text_matches_matches",
			predicate: EditedPostTextMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testText}},
			matches:   true,
		},
		{
			name:      "edited_post_text_matches_not_matches",
			predicate: EditedPostTextMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{EditedChannelPost: &telego.Message{Text: testCommand1}},
			matches:   false,
		},
		{
			name:      "any_inline_query_matches",
			predicate: AnyInlineQuery(),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{}},
			matches:   true,
		},
		{
			name:      "any_inline_query_not_matches",
			predicate: AnyInlineQuery(),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "inline_query_equal_matches",
			predicate: InlineQueryEqual(testText),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testText}},
			matches:   true,
		},
		{
			name:      "inline_query_equal_not_matches",
			predicate: InlineQueryEqual(testText),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testTextLower}},
			matches:   false,
		},
		{
			name:      "inline_query_equal_fold_matches",
			predicate: InlineQueryEqualFold(testText),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testTextLower}},
			matches:   true,
		},
		{
			name:      "inline_query_equal_fold_not_matches",
			predicate: InlineQueryEqualFold(testText),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testCommand1}},
			matches:   false,
		},
		{
			name:      "inline_query_contains_matches",
			predicate: InlineQueryContains(testTextPart),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testText}},
			matches:   true,
		},
		{
			name:      "inline_query_contains_not_matches",
			predicate: InlineQueryContains(testTextPart),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testCommand1}},
			matches:   false,
		},
		{
			name:      "inline_query_prefix_matches",
			predicate: InlineQueryPrefix(testTextPrefix),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testText}},
			matches:   true,
		},
		{
			name:      "inline_query_prefix_not_matches",
			predicate: InlineQueryPrefix(testTextPrefix),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testCommand1}},
			matches:   false,
		},
		{
			name:      "inline_query_suffix_matches",
			predicate: InlineQuerySuffix(testTextSuffix),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testText}},
			matches:   true,
		},
		{
			name:      "inline_query_suffix_not_matches",
			predicate: InlineQuerySuffix(testTextSuffix),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testCommand1}},
			matches:   false,
		},
		{
			name:      "inline_query_matches_matches",
			predicate: InlineQueryMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testText}},
			matches:   true,
		},
		{
			name:      "inline_query_matches_not_matches",
			predicate: InlineQueryMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{InlineQuery: &telego.InlineQuery{Query: testCommand1}},
			matches:   false,
		},
		{
			name:      "any_callback_query_matches",
			predicate: AnyCallbackQuery(),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{}},
			matches:   true,
		},
		{
			name:      "any_callback_query_not_matches",
			predicate: AnyCallbackQuery(),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "any_callback_query_with_message_matches",
			predicate: AnyCallbackQueryWithMessage(),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Message: &telego.Message{}}},
			matches:   true,
		},
		{
			name:      "any_callback_query_with_message_not_matches",
			predicate: AnyCallbackQueryWithMessage(),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{}},
			matches:   false,
		},
		{
			name:      "callback_data_equal_matches",
			predicate: CallbackDataEqual(testText),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testText}},
			matches:   true,
		},
		{
			name:      "callback_data_equal_not_matches",
			predicate: CallbackDataEqual(testText),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testTextLower}},
			matches:   false,
		},
		{
			name:      "callback_data_equal_fold_matches",
			predicate: CallbackDataEqualFold(testText),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testTextLower}},
			matches:   true,
		},
		{
			name:      "callback_data_equal_fold_not_matches",
			predicate: CallbackDataEqualFold(testText),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testCommand1}},
			matches:   false,
		},
		{
			name:      "callback_data_contains_matches",
			predicate: CallbackDataContains(testTextPart),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testText}},
			matches:   true,
		},
		{
			name:      "callback_data_contains_not_matches",
			predicate: CallbackDataContains(testTextPart),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testCommand1}},
			matches:   false,
		},
		{
			name:      "callback_data_prefix_matches",
			predicate: CallbackDataPrefix(testTextPrefix),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testText}},
			matches:   true,
		},
		{
			name:      "callback_data_prefix_not_matches",
			predicate: CallbackDataPrefix(testTextPrefix),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testCommand1}},
			matches:   false,
		},
		{
			name:      "callback_data_suffix_matches",
			predicate: CallbackDataSuffix(testTextSuffix),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testText}},
			matches:   true,
		},
		{
			name:      "callback_data_suffix_not_matches",
			predicate: CallbackDataSuffix(testTextSuffix),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testCommand1}},
			matches:   false,
		},
		{
			name:      "callback_data_matches_matches",
			predicate: CallbackDataMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testText}},
			matches:   true,
		},
		{
			name:      "callback_data_matches_not_matches",
			predicate: CallbackDataMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{CallbackQuery: &telego.CallbackQuery{Data: testCommand1}},
			matches:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.matches, tt.predicate(tt.update))
		})
	}
}
