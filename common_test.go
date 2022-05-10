package telego

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego/telegoapi"
	mockAPI "github.com/mymmrac/telego/telegoapi/mock"
)

var (
	data = &telegoapi.RequestData{}
	resp = &telegoapi.Response{
		Ok: true,
	}

	expectedMessage = &Message{
		MessageID: 1,
	}
)

func setResult(t *testing.T, v interface{}) {
	t.Helper()

	bytesData, err := json.Marshal(v)
	assert.NoError(t, err)
	resp.Result = bytesData
}

type mockedBot struct {
	MockAPICaller          *mockAPI.MockCaller
	MockRequestConstructor *mockAPI.MockRequestConstructor
	Bot                    *Bot
}

func newMockedBot(ctrl *gomock.Controller) mockedBot {
	mb := mockedBot{
		MockAPICaller:          mockAPI.NewMockCaller(ctrl),
		MockRequestConstructor: mockAPI.NewMockRequestConstructor(ctrl),
	}

	//nolint:errcheck
	bot, _ := NewBot(token,
		WithAPICaller(mb.MockAPICaller),
		WithRequestConstructor(mb.MockRequestConstructor),
		WithDiscardLogger())

	mb.Bot = bot

	return mb
}

type testNamedReade struct{}

func (t testNamedReade) Read(_ []byte) (n int, err error) {
	panic("implement me")
}

func (t testNamedReade) Name() string {
	return "test"
}

var testInputFile = InputFile{
	File: testNamedReade{},
}
