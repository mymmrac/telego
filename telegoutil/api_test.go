package telegoutil

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameReader(t *testing.T) {
	buf := strings.NewReader(text1)

	nr := NameReader(buf, text2)

	data, err := ioutil.ReadAll(nr)
	assert.NoError(t, err)

	assert.Equal(t, text1, string(data))
	assert.Equal(t, text2, nr.Name())
}
