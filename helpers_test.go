package telego

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBot_SetUpdateInterval(t *testing.T) {
	bot := &Bot{}
	ui := time.Second

	bot.SetUpdateInterval(ui)
	assert.Equal(t, ui, bot.updateInterval)
}

func TestBot_StopGettingUpdates(t *testing.T) {
	bot := &Bot{}

	bot.stopChannel = make(chan struct{})
	assert.NotPanics(t, func() {
		bot.StopGettingUpdates()
	})
}
