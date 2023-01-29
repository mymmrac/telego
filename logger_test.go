package telego

import (
	"bytes"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	data1  = "test"
	data2  = "ok"
	data3  = "hello"
	format = "%s %s"
)

func testLogger() (*logger, *bytes.Buffer) {
	buffer := &bytes.Buffer{}

	return &logger{
		Out:         buffer,
		DebugMode:   false,
		PrintErrors: false,
		mutex:       sync.Mutex{},
	}, buffer
}

func Test_logger_Debugf(t *testing.T) {
	l, b := testLogger()

	t.Run("disabled", func(t *testing.T) {
		l.Debugf(format, data1, data2)
		assert.Equal(t, "", b.String())
	})

	t.Run("enabled", func(t *testing.T) {
		l.DebugMode = true
		l.Debugf(format, data1, data2)
		assert.Contains(t, b.String(), data1)
		assert.Contains(t, b.String(), data2)
	})
}

func Test_logger_Errorf(t *testing.T) {
	l, b := testLogger()

	t.Run("disabled", func(t *testing.T) {
		l.Errorf(format, data1, data2)
		assert.Equal(t, "", b.String())
	})

	t.Run("enabled", func(t *testing.T) {
		l.PrintErrors = true
		l.Errorf(format, data1, data2)
		assert.Contains(t, b.String(), data1)
		assert.Contains(t, b.String(), data2)
	})
}

func Test_logger_replacer(t *testing.T) {
	l, b := testLogger()

	l.DebugMode = true
	l.Replacer = strings.NewReplacer(data1, data3)
	l.Debugf(data1, data2)
	assert.NotContains(t, b.String(), data1)
	assert.Contains(t, b.String(), data2)
	assert.Contains(t, b.String(), data3)
}

func Test_newDefaultLogger(t *testing.T) {
	l := newDefaultLogger("token")

	assert.Equal(t, os.Stderr, l.Out)
	assert.False(t, l.DebugMode)
	assert.True(t, l.PrintErrors)
	assert.NotNil(t, l.Replacer)
}
