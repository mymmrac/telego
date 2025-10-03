package telegohandler

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

func TestWithErrorHandler(t *testing.T) {
	bh := &BotHandler{}
	handler := func(ctx *Context, update telego.Update, err error) {}

	err := WithErrorHandler(handler)(bh)
	require.NoError(t, err)
}
