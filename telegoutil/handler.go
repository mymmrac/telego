package telegoutil

import (
	"strings"

	th "github.com/mymmrac/telego/telegohandler"
)

// ParseCommand returns command, bot username and its arguments if any
func ParseCommand(text string) (cmd string, username string, args []string) {
	matches := th.CommandRegexp.FindStringSubmatch(text)
	if len(matches) != th.CommandMatchGroupsLen {
		return "", "", nil
	}

	if matches[th.CommandMatchArgsGroup] == "" {
		return matches[th.CommandMatchCmdGroup], matches[th.CommandMatchBotUsernameGroup], []string{}
	}

	return matches[th.CommandMatchCmdGroup], matches[th.CommandMatchBotUsernameGroup],
		strings.Split(matches[th.CommandMatchArgsGroup], " ")
}
