package telego

import (
	"bytes"
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"go.uber.org/mock/gomock"

	"github.com/mymmrac/telego/internal/json"
	ta "github.com/mymmrac/telego/telegoapi"
	mockapi "github.com/mymmrac/telego/telegoapi/mock"
)

type testCallerType struct{}

func (c testCallerType) Call(_ context.Context, _ string, _ *ta.RequestData) (*ta.Response, error) {
	panic("implement me")
}

func TestWithAPICaller(t *testing.T) {
	bot := &Bot{}
	caller := testCallerType{}

	err := WithAPICaller(caller)(bot)
	require.NoError(t, err)
	assert.EqualValues(t, caller, bot.api)
}

func TestWithFastHTTPClient(t *testing.T) {
	bot := &Bot{}
	client := &fasthttp.Client{}

	err := WithFastHTTPClient(client)(bot)
	require.NoError(t, err)
}

func TestWithHTTPClient(t *testing.T) {
	bot := &Bot{}
	client := &http.Client{}

	err := WithHTTPClient(client)(bot)
	require.NoError(t, err)
}

type testConstructorType struct{}

func (testConstructorType) JSONRequest(_ any) (*ta.RequestData, error) {
	panic("implement me")
}

func (testConstructorType) MultipartRequest(_ map[string]string, _ map[string]ta.NamedReader,
) (*ta.RequestData, error) {
	panic("implement me")
}

func TestWithRequestConstructor(t *testing.T) {
	bot := &Bot{}
	constructor := &testConstructorType{}

	err := WithRequestConstructor(constructor)(bot)
	require.NoError(t, err)
	assert.EqualValues(t, constructor, bot.constructor)
}

func TestWithDefaultLogger(t *testing.T) {
	bot := &Bot{}

	err := WithDefaultLogger(true, false)(bot)
	require.NoError(t, err)

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
		require.NoError(t, err)

		log, ok := bot.log.(*logger)
		require.True(t, ok)

		assert.True(t, log.DebugMode)
		assert.True(t, log.PrintErrors)
		assert.Nil(t, log.Replacer)
	})

	t.Run("not_nil_replacer", func(t *testing.T) {
		err := WithExtendedDefaultLogger(true, true, strings.NewReplacer("old", "new"))(bot)
		require.NoError(t, err)

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
	require.NoError(t, err)

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
	require.NoError(t, err)
	assert.EqualValues(t, log, bot.log)
}

func TestWithAPIServer(t *testing.T) {
	bot := &Bot{}

	t.Run("success", func(t *testing.T) {
		err := WithAPIServer("test")(bot)
		require.NoError(t, err)
		assert.Equal(t, "test", bot.apiURL)
	})

	t.Run("error", func(t *testing.T) {
		err := WithAPIServer("")(bot)
		require.Error(t, err)
	})
}

func TestWithDefaultDebugLogger(t *testing.T) {
	bot := &Bot{}

	err := WithDefaultDebugLogger()(bot)
	require.NoError(t, err)

	log, ok := bot.log.(*logger)
	require.True(t, ok)

	assert.True(t, log.DebugMode)
	assert.True(t, log.PrintErrors)
	assert.NotNil(t, log.Replacer)
}

func TestWithTestServerPath(t *testing.T) {
	bot := &Bot{}

	err := WithTestServerPath()(bot)
	require.NoError(t, err)

	assert.True(t, bot.useTestServerPath)
}

func TestWithHealthCheck(t *testing.T) {
	ctrl := gomock.NewController(t)

	caller := mockapi.NewMockCaller(ctrl)
	constructor := mockapi.NewMockRequestConstructor(ctrl)

	expectedResp := &ta.Response{
		Ok:     true,
		Result: json.RawMessage(`{}`),
	}

	expectedData := &ta.RequestData{
		ContentType: ta.ContentTypeJSON,
		Buffer:      bytes.NewBuffer([]byte{}),
	}

	constructor.EXPECT().
		JSONRequest(nil).
		Return(expectedData, nil).
		Times(1)

	caller.EXPECT().
		Call(t.Context(), defaultBotAPIServer+botPathPrefix+token+"/getMe", expectedData).
		Return(expectedResp, nil).
		Times(1)

	bot, err := NewBot(token,
		WithAPICaller(caller),
		WithRequestConstructor(constructor),
		WithHealthCheck(t.Context()),
	)
	require.NoError(t, err)
	require.NotNil(t, bot)
}

func TestWithWarnings(t *testing.T) {
	bot := &Bot{}

	err := WithWarnings()(bot)
	require.NoError(t, err)

	assert.True(t, bot.reportWarningAsErrors)
}
