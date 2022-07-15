package telegoutil

import (
	"io/ioutil"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego"
)

func TestNameReader(t *testing.T) {
	buf := strings.NewReader(text1)

	nr := NameReader(buf, text2)

	data, err := ioutil.ReadAll(nr)
	assert.NoError(t, err)

	assert.Equal(t, text1, string(data))
	assert.Equal(t, text2, nr.Name())
}

func TestUpdateProcessor(t *testing.T) {
	updates := make(chan telego.Update)

	wg := sync.WaitGroup{}

	processedUpdates := UpdateProcessor(updates, 10, func(update telego.Update) telego.Update {
		wg.Done()
		update.UpdateID *= 10
		return update
	})

	const updateCount = 2
	wg.Add(updateCount)

	updates <- telego.Update{UpdateID: 1}
	updates <- telego.Update{UpdateID: 2}

	wg.Wait()

	count := 0
	for update := range processedUpdates {
		count++
		assert.True(t, update.UpdateID%10 == 0)

		if count == updateCount {
			close(updates)
		}
	}
}
