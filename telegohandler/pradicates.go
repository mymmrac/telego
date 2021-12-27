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

// HasMassage is true if message isn't nil
func HasMassage() Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil
	}
}

// CommandRegexp matches to command and has match groups on command and arguments
var CommandRegexp = regexp.MustCompile(`/(\w+) ?(.*)`)

// HasCommand is true if message isn't nil, and it matches to command regexp
func HasCommand() Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && CommandRegexp.MatchString(update.Message.Text)
	}
}

// TextEqual is true if message isn't nil, and it's equal to specified text
func TextEqual(text string) Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && update.Message.Text == text
	}
}

// TextEqualFold is true if message isn't nil, and it's equal fold (more general form of case-insensitivity) to
// specified text
func TextEqualFold(text string) Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && strings.EqualFold(update.Message.Text, text)
	}
}

// HasText is true if message isn't nil, and it contains specified text
func HasText(text string) Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && strings.Contains(update.Message.Text, text)
	}
}

// HasPrefix is true if message isn't nil, and it has specified prefix
func HasPrefix(prefix string) Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && strings.HasPrefix(update.Message.Text, prefix)
	}
}

// HasSuffix is true if message isn't nil, and it has specified suffix
func HasSuffix(suffix string) Predicate {
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
