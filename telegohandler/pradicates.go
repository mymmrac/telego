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

// HasMassage is true if message isn't nil
func HasMassage() Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil
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

// ContainsText is true if message isn't nil, and it contains specified text
func ContainsText(text string) Predicate {
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

// CommandRegexp matches to command and has match groups on command and arguments
var CommandRegexp = regexp.MustCompile(`^/(\w+) ?(.*)$`)

// CommandMatchGroupsLen represents length of match groups of CommandRegexp
const CommandMatchGroupsLen = 3

// HasCommand is true if message isn't nil, and it matches to command regexp
func HasCommand() Predicate {
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

// CommandEqualWithArgc is true if message isn't nil, and it contains specified command with number of args
func CommandEqualWithArgc(command string, argc int) Predicate {
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

// CommandEqualWithArgv is true if message isn't nil, and it contains specified command and args
func CommandEqualWithArgv(command string, argv ...string) Predicate {
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
