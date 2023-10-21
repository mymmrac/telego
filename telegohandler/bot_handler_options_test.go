package telegohandler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWithDone(t *testing.T) {
	bh := &BotHandler{}
	err := WithStopTimeout(-time.Second)(bh)
	require.Error(t, err)
}

func TestWithStopTimeout(t *testing.T) {
	bh := &BotHandler{}
	err := WithDone(nil)(bh)
	require.Error(t, err)
}
