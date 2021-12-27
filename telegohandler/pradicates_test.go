package telegohandler

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego"
)

const (
	text       = `Test text`
	textLower  = `test text`
	textPart   = `t te`
	textPrefix = `Test`
	textSuffix = `text`
	command    = `/test abc 123`
)

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
			name:      "has_massage_matches",
			predicate: HasMassage(),
			update:    telego.Update{Message: &telego.Message{}},
			matches:   true,
		},
		{
			name:      "has_massage_not_matches",
			predicate: HasMassage(),
			update:    telego.Update{},
			matches:   false,
		},
		{
			name:      "has_command_matches",
			predicate: HasCommand(),
			update:    telego.Update{Message: &telego.Message{Text: command}},
			matches:   true,
		},
		{
			name:      "has_command_not_matches",
			predicate: HasCommand(),
			update:    telego.Update{Message: &telego.Message{Text: text}},
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
			update:    telego.Update{Message: &telego.Message{Text: command}},
			matches:   false,
		},
		{
			name:      "has_text_matches",
			predicate: HasText(textPart),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   true,
		},
		{
			name:      "has_text_not_matches",
			predicate: HasText(textPart),
			update:    telego.Update{Message: &telego.Message{Text: command}},
			matches:   false,
		},
		{
			name:      "has_prefix_matches",
			predicate: HasPrefix(textPrefix),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   true,
		},
		{
			name:      "has_prefix_not_matches",
			predicate: HasPrefix(textPrefix),
			update:    telego.Update{Message: &telego.Message{Text: command}},
			matches:   false,
		},
		{
			name:      "has_suffix_matches",
			predicate: HasSuffix(textSuffix),
			update:    telego.Update{Message: &telego.Message{Text: text}},
			matches:   true,
		},
		{
			name:      "has_suffix_not_matches",
			predicate: HasSuffix(textSuffix),
			update:    telego.Update{Message: &telego.Message{Text: command}},
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
			update:    telego.Update{Message: &telego.Message{Text: command}},
			matches:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.matches, tt.predicate(tt.update))
		})
	}
}
