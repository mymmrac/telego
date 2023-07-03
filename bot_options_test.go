package telego

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/telegoapi"
)

type testCallerType struct{}

func (c testCallerType) Call(_ string, _ *telegoapi.RequestData) (*telegoapi.Response, error) {
	panic("implement me")
}

func TestWithAPICaller(t *testing.T) {
	bot := &Bot{}
	caller := testCallerType{}

	err := WithAPICaller(caller)(bot)
	assert.NoError(t, err)
	assert.EqualValues(t, caller, bot.api)
}

func TestWithFastHTTPClient(t *testing.T) {
	bot := &Bot{}
	client := &fasthttp.Client{}

	err := WithFastHTTPClient(client)(bot)
	assert.NoError(t, err)
}

func TestWithHTTPClient(t *testing.T) {
	bot := &Bot{}
	client := &http.Client{}

	err := WithHTTPClient(client)(bot)
	assert.NoError(t, err)
}

type testConstructorType struct{}

func (testConstructorType) JSONRequest(_ any) (*telegoapi.RequestData, error) {
	panic("implement me")
}

func (testConstructorType) MultipartRequest(_ map[string]string, _ map[string]telegoapi.NamedReader,
) (*telegoapi.RequestData, error) {
	panic("implement me")
}

func TestWithRequestConstructor(t *testing.T) {
	bot := &Bot{}
	constructor := &testConstructorType{}

	err := WithRequestConstructor(constructor)(bot)
	assert.NoError(t, err)
	assert.EqualValues(t, constructor, bot.constructor)
}

func TestWithDefaultLogger(t *testing.T) {
	bot := &Bot{}

	err := WithDefaultLogger(true, false)(bot)
	assert.NoError(t, err)

	log, ok := bot.log.(*logger)
	require.True(t, ok)

	assert.True(t, log.DebugMode)
	assert.False(t, log.PrintErrors)
	assert.NotNil(t, log.Replacer)
}

func TestWithExtendedDefaultLogger(t *testing.T) {
	bot := &Bot{}

	t.Run("nil_replacer", func(t *testing.T) {
		err := WithExtendedDefaultLogger(true, true, nil)(bot)
		assert.NoError(t, err)

		log, ok := bot.log.(*logger)
		require.True(t, ok)

		assert.True(t, log.DebugMode)
		assert.True(t, log.PrintErrors)
		assert.Nil(t, log.Replacer)
	})

	t.Run("not_nil_replacer", func(t *testing.T) {
		err := WithExtendedDefaultLogger(true, true, strings.NewReplacer("old", "new"))(bot)
		assert.NoError(t, err)

		log, ok := bot.log.(*logger)
		require.True(t, ok)

		assert.True(t, log.DebugMode)
		assert.True(t, log.PrintErrors)
		assert.NotNil(t, log.Replacer)
	})
}

func TestWithDiscardLogger(t *testing.T) {
	bot := &Bot{}

	err := WithDiscardLogger()(bot)
	assert.NoError(t, err)

	log, ok := bot.log.(*logger)
	require.True(t, ok)

	assert.False(t, log.DebugMode)
	assert.False(t, log.PrintErrors)
	assert.NotNil(t, log.Replacer)
}

type testLoggerType struct{}

func (testLoggerType) Debugf(_ string, _ ...any) {
	// NoOp
}

func (testLoggerType) Errorf(_ string, _ ...any) {
	// NoOp
}

func TestWithLogger(t *testing.T) {
	bot := &Bot{}
	log := &testLoggerType{}

	err := WithLogger(log)(bot)
	assert.NoError(t, err)
	assert.EqualValues(t, log, bot.log)
}

func TestWithAPIServer(t *testing.T) {
	bot := &Bot{}

	t.Run("success", func(t *testing.T) {
		err := WithAPIServer("test")(bot)
		assert.NoError(t, err)
		assert.Equal(t, "test", bot.apiURL)
	})

	t.Run("error", func(t *testing.T) {
		err := WithAPIServer("")(bot)
		assert.Error(t, err)
	})
}

func TestWithDefaultDebugLogger(t *testing.T) {
	bot := &Bot{}

	err := WithDefaultDebugLogger()(bot)
	assert.NoError(t, err)

	log, ok := bot.log.(*logger)
	require.True(t, ok)

	assert.True(t, log.DebugMode)
	assert.True(t, log.PrintErrors)
	assert.NotNil(t, log.Replacer)
}

func TestWithHealthCheck(t *testing.T) {
	bot := &Bot{}

	err := WithHealthCheck()(bot)
	assert.NoError(t, err)

	assert.True(t, bot.healthCheckRequested)
}

func TestWithWarnings(t *testing.T) {
	bot := &Bot{}

	err := WithWarnings()(bot)
	assert.NoError(t, err)

	assert.True(t, bot.warningAsErrors)
}

func TestWithRetry(t *testing.T) {
	bot := &Bot{}

	t.Run("success", func(t *testing.T) {
		err := WithRetry(1, 1, 2, 3)(bot)
		assert.NoError(t, err)
		assert.Equal(t, 1, bot.retryOptions.maxAttempts)
		assert.Equal(t, 1, bot.retryOptions.delayFactor)
		assert.Equal(t, time.Duration(2), bot.retryOptions.startDelay)
		assert.Equal(t, time.Duration(3), bot.retryOptions.maxDelay)
	})

	t.Run("error", func(t *testing.T) {
		err := WithRetry(0, 1, 2, 3)(bot)
		assert.Error(t, err)

		err = WithRetry(1, 0, 2, 3)(bot)
		assert.Error(t, err)
	})
}
