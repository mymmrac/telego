package telego

import (
	"strings"
	"testing"

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

type testConstructorType struct{}

func (testConstructorType) JSONRequest(_ interface{}) (*telegoapi.RequestData, error) {
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

func (testLoggerType) Debug(_ ...interface{}) {
	panic("implement me")
}

func (testLoggerType) Debugf(_ string, _ ...interface{}) {
	panic("implement me")
}

func (testLoggerType) Error(_ ...interface{}) {
	panic("implement me")
}

func (testLoggerType) Errorf(_ string, _ ...interface{}) {
	panic("implement me")
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
