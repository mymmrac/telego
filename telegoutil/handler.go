package telegoutil

import (
	"strings"

	th "github.com/mymmrac/telego/telegohandler"
)

// ParseCommand returns command and its arguments if any
func ParseCommand(text string) (string, []string) {
	matches := th.CommandRegexp.FindStringSubmatch(text)
	if len(matches) != th.CommandMatchGroupsLen {
		return "", nil
	}

	if matches[2] == "" {
		return matches[1], []string{}
	}

	return matches[1], strings.Split(matches[2], " ")
}
