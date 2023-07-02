package telegohandler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithDone(t *testing.T) {
	bh := &BotHandler{}
	err := WithStopTimeout(-time.Second)(bh)
	assert.Error(t, err)
}

func TestWithStopTimeout(t *testing.T) {
	bh := &BotHandler{}
	err := WithDone(nil)(bh)
	assert.Error(t, err)
}
