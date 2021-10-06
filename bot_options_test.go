package telego

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/api"
)

type testCallerType struct{}

func (c testCallerType) Call(_ string, _ *api.RequestData) (*api.Response, error) {
	panic("implement me")
}

func TestCustomAPICaller(t *testing.T) {
	bot := &Bot{}
	caller := testCallerType{}

	err := CustomAPICaller(caller)(bot)
	assert.NoError(t, err)
	assert.EqualValues(t, caller, bot.api)
}

func TestFastHTTPClient(t *testing.T) {
	bot := &Bot{}
	client := &fasthttp.Client{}

	err := FastHTTPClient(client)(bot)
	assert.NoError(t, err)
}

type testConstructorType struct{}

func (testConstructorType) JSONRequest(_ interface{}) (*api.RequestData, error) {
	panic("implement me")
}

func (testConstructorType) MultipartRequest(_ map[string]string, _ map[string]api.NamedReader,
) (*api.RequestData, error) {
	panic("implement me")
}

func TestCustomRequestConstructor(t *testing.T) {
	bot := &Bot{}
	constructor := &testConstructorType{}

	err := CustomRequestConstructor(constructor)(bot)
	assert.NoError(t, err)
	assert.EqualValues(t, constructor, bot.constructor)
}

func TestDefaultLogger(t *testing.T) {
	bot := &Bot{}

	err := DefaultLogger(true, true)(bot)
	assert.NoError(t, err)
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

func TestSetLogger(t *testing.T) {
	bot := &Bot{}
	log := &testLoggerType{}

	err := SetLogger(log)(bot)
	assert.NoError(t, err)
	assert.EqualValues(t, log, bot.log)
}

func TestSetAPIServer(t *testing.T) {
	bot := &Bot{}

	t.Run("success", func(t *testing.T) {
		err := SetAPIServer("test")(bot)
		assert.NoError(t, err)
		assert.Equal(t, "test", bot.apiURL)
	})

	t.Run("error", func(t *testing.T) {
		err := SetAPIServer("")(bot)
		assert.Error(t, err)
	})
}
