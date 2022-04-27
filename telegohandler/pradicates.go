package telegohandler

import (
	"regexp"
	"strings"

	"github.com/mymmrac/telego"
)

// Union is true if at least one of predicates is true
func Union(predicates ...Predicate) Predicate {
	return func(update telego.Update) bool {
		for _, p := range predicates {
			if p(update) {
				return true
			}
		}
		return false
	}
}

// Not is true if predicate is false
func Not(predicate Predicate) Predicate {
	return func(update telego.Update) bool {
		return !predicate(update)
	}
}

// AnyMessage is true if message isn't nil
func AnyMessage() Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil
	}
}

// TextEqual is true if message isn't nil, and its equal to specified text
func TextEqual(text string) Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && update.Message.Text == text
	}
}

// TextEqualFold is true if message isn't nil, and its equal fold (more general form of case-insensitivity equal) to
// specified text
func TextEqualFold(text string) Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && strings.EqualFold(update.Message.Text, text)
	}
}

// TextContains is true if message isn't nil, and it contains specified text
func TextContains(text string) Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && strings.Contains(update.Message.Text, text)
	}
}

// TextPrefix is true if message isn't nil, and it has specified prefix
func TextPrefix(prefix string) Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && strings.HasPrefix(update.Message.Text, prefix)
	}
}

// TextSuffix is true if message isn't nil, and it has specified suffix
func TextSuffix(suffix string) Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && strings.HasSuffix(update.Message.Text, suffix)
	}
}

// TextMatches is true if message isn't nil, and it matches specified regexp
func TextMatches(pattern *regexp.Regexp) Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && pattern.MatchString(update.Message.Text)
	}
}

// CommandRegexp matches to command and has match groups on command and arguments
var CommandRegexp = regexp.MustCompile(`^/(\w+) ?(.*)$`)

// CommandMatchGroupsLen represents length of match groups of CommandRegexp
const CommandMatchGroupsLen = 3

// AnyCommand is true if message isn't nil, and it matches to command regexp
func AnyCommand() Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && CommandRegexp.MatchString(update.Message.Text)
	}
}

// CommandEqual is true if message isn't nil, and it contains specified command
func CommandEqual(command string) Predicate {
	return func(update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Text)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[1], command)
	}
}

// CommandEqualArgc is true if message isn't nil, and it contains specified command with number of args
func CommandEqualArgc(command string, argc int) Predicate {
	return func(update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Text)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[1], command) &&
			(argc == 0 && matches[2] == "" || len(strings.Split(matches[2], " ")) == argc)
	}
}

// CommandEqualArgv is true if message isn't nil, and it contains specified command and args
func CommandEqualArgv(command string, argv ...string) Predicate {
	return func(update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Text)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[1], command) &&
			(len(argv) == 0 && matches[2] == "" || matches[2] == strings.Join(argv, " "))
	}
}

// AnyEditedMessage is true if edited message isn't nil
func AnyEditedMessage() Predicate {
	return func(update telego.Update) bool {
		return update.EditedMessage != nil
	}
}

// EditedTextEqual is true if edited message isn't nil, and its equal to specified text
func EditedTextEqual(text string) Predicate {
	return func(update telego.Update) bool {
		return update.EditedMessage != nil && update.EditedMessage.Text == text
	}
}

// EditedTextEqualFold is true if edited message isn't nil, and its equal fold (more general form of case-insensitivity
// equal) to specified text
func EditedTextEqualFold(text string) Predicate {
	return func(update telego.Update) bool {
		return update.EditedMessage != nil && strings.EqualFold(update.EditedMessage.Text, text)
	}
}

// EditedTextContains is true if edited message isn't nil, and it contains specified text
func EditedTextContains(text string) Predicate {
	return func(update telego.Update) bool {
		return update.EditedMessage != nil && strings.Contains(update.EditedMessage.Text, text)
	}
}

// EditedTextPrefix is true if edited message isn't nil, and it has specified prefix
func EditedTextPrefix(prefix string) Predicate {
	return func(update telego.Update) bool {
		return update.EditedMessage != nil && strings.HasPrefix(update.EditedMessage.Text, prefix)
	}
}

// EditedTextSuffix is true if edited message isn't nil, and it has specified suffix
func EditedTextSuffix(suffix string) Predicate {
	return func(update telego.Update) bool {
		return update.EditedMessage != nil && strings.HasSuffix(update.EditedMessage.Text, suffix)
	}
}

// EditedTextMatches is true if edited message isn't nil, and it matches specified regexp
func EditedTextMatches(pattern *regexp.Regexp) Predicate {
	return func(update telego.Update) bool {
		return update.EditedMessage != nil && pattern.MatchString(update.EditedMessage.Text)
	}
}

// AnyCallbackQuery is true if callback query isn't nil
func AnyCallbackQuery() Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil
	}
}

// AnyCallbackQueryWithMessage is true if callback query and its message isn't nil
func AnyCallbackQueryWithMessage() Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && update.CallbackQuery.Message != nil
	}
}

// CallbackDataEqual is true if callback query isn't nil, and its data equal to specified text
func CallbackDataEqual(text string) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && update.CallbackQuery.Data == text
	}
}

// CallbackDataEqualFold is true if callback query isn't nil, and its data equal fold (more general form of
// case-insensitivity equal) to specified text
func CallbackDataEqualFold(text string) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && strings.EqualFold(update.CallbackQuery.Data, text)
	}
}

// CallbackDataContains is true if callback query isn't nil, and its data contains specified text
func CallbackDataContains(text string) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && strings.Contains(update.CallbackQuery.Data, text)
	}
}

// CallbackDataPrefix is true if callback query isn't nil, and its data has specified prefix
func CallbackDataPrefix(prefix string) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && strings.HasPrefix(update.CallbackQuery.Data, prefix)
	}
}

// CallbackDataSuffix is true if callback query isn't nil, and its data has specified suffix
func CallbackDataSuffix(suffix string) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && strings.HasSuffix(update.CallbackQuery.Data, suffix)
	}
}

// CallbackDataMatches is true if callback query isn't nil, and its data matches specified regexp
func CallbackDataMatches(pattern *regexp.Regexp) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && pattern.MatchString(update.CallbackQuery.Data)
	}
}
