package telegohandler

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego"
)

const (
	text        = `Test text`
	textLower   = `test text`
	textPart    = `t te`
	textPrefix  = `Test`
	textSuffix  = `text`
	command1    = `/test abc 123`
	command2    = `/hmm bcd`
	commandName = `test`
)

//nolint:funlen
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
			predicate: TextEqual(text),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   true,
		},
		{
			name:      "text_equal_not_matches",
			predicate: TextEqual(text),
			update:    telego.Update{Message: &telego.Message{Text: textLower}},
			matches:   false,
		},
		{
			name:      "text_equal_fold_matches",
			predicate: TextEqualFold(text),
			update:    telego.Update{Message: &telego.Message{Text: textLower}},
			matches:   true,
		},
		{
			name:      "text_equal_fold_not_matches",
			predicate: TextEqualFold(text),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   false,
		},
		{
			name:      "text_contains_matches",
			predicate: TextContains(textPart),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   true,
		},
		{
			name:      "text_contains_not_matches",
			predicate: TextContains(textPart),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   false,
		},
		{
			name:      "text_prefix_matches",
			predicate: TextPrefix(textPrefix),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   true,
		},
		{
			name:      "text_prefix_not_matches",
			predicate: TextPrefix(textPrefix),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   false,
		},
		{
			name:      "text_suffix_matches",
			predicate: TextSuffix(textSuffix),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   true,
		},
		{
			name:      "text_suffix_not_matches",
			predicate: TextSuffix(textSuffix),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   false,
		},
		{
			name:      "text_matches_matches",
			predicate: TextMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   true,
		},
		{
			name:      "text_matches_not_matches",
			predicate: TextMatches(regexp.MustCompile(`^\w+ \w+$`)),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   false,
		},
		{
			name:      "any_command_matches",
			predicate: AnyCommand(),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   true,
		},
		{
			name:      "any_command_not_matches",
			predicate: AnyCommand(),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   false,
		},
		{
			name:      "command_equal_matches",
			predicate: CommandEqual(commandName),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   true,
		},
		{
			name:      "command_equal_not_matches",
			predicate: CommandEqual(commandName),
			update:    telego.Update{Message: &telego.Message{Text: command2}},
			matches:   false,
		},
		{
			name:      "command_equal_no_message",
			predicate: CommandEqual(commandName),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "command_equal_no_command",
			predicate: CommandEqual(commandName),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   false,
		},
		{
			name:      "command_equal_argc_matches",
			predicate: CommandEqualArgc(commandName, 2),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   true,
		},
		{
			name:      "command_equal_argc_not_matches",
			predicate: CommandEqualArgc(commandName, 3),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   false,
		},
		{
			name:      "command_equal_argc_no_message",
			predicate: CommandEqualArgc(commandName, 0),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "command_equal_argc_no_command",
			predicate: CommandEqualArgc(commandName, 0),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   false,
		},
		{
			name:      "command_equal_argv_matches",
			predicate: CommandEqualArgv(commandName, "abc", "123"),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   true,
		},
		{
			name:      "command_equal_argv_not_matches",
			predicate: CommandEqualArgv(commandName, "abc", "abc"),
			update:    telego.Update{Message: &telego.Message{Text: command1}},
			matches:   false,
		},
		{
			name:      "command_equal_argv_no_message",
			predicate: CommandEqualArgv(commandName),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "command_equal_argv_no_command",
			predicate: CommandEqualArgv(commandName),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.matches, tt.predicate(tt.update))
		})
	}
}
