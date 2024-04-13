package telegoutil

import (
	"strings"

	th "github.com/mymmrac/telego/telegohandler"
)

// ParseCommand returns command, bot username and its arguments if any
func ParseCommand(text string) (cmd string, username string, args []string) {
	var payload string
	cmd, username, payload = ParseCommandPayload(text)
	return cmd, username, strings.Fields(payload)
}

// ParseCommandPayload returns command, bot username and its payload if any
func ParseCommandPayload(text string) (cmd string, username string, payload string) {
	matches := th.CommandRegexp.FindStringSubmatch(text)
	if len(matches) != th.CommandMatchGroupsLen {
		return "", "", ""
	}
	return matches[th.CommandMatchCmdGroup], matches[th.CommandMatchBotUsernameGroup], matches[th.CommandMatchArgsGroup]
}
