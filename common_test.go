package telego

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/mymmrac/telego/internal/json"
	ta "github.com/mymmrac/telego/telegoapi"
	mockapi "github.com/mymmrac/telego/telegoapi/mock"
)

var (
	data      = &ta.RequestData{}
	emptyResp = &ta.Response{
		Ok: true,
	}

	expectedMessage = &Message{
		MessageID: 1,
	}
)

func telegoResponse(t *testing.T, v any) *ta.Response {
	t.Helper()

	byteData, err := json.Marshal(v)
	require.NoError(t, err)
	return &ta.Response{
		Ok:     true,
		Result: byteData,
	}
}

type mockedBot struct {
	MockAPICaller          *mockapi.MockCaller
	MockRequestConstructor *mockapi.MockRequestConstructor
	Bot                    *Bot
}

func newMockedBot(ctrl *gomock.Controller) mockedBot {
	mb := mockedBot{
		MockAPICaller:          mockapi.NewMockCaller(ctrl),
		MockRequestConstructor: mockapi.NewMockRequestConstructor(ctrl),
	}

	//nolint:errcheck
	bot, _ := NewBot(validToken,
		WithAPICaller(mb.MockAPICaller),
		WithRequestConstructor(mb.MockRequestConstructor),
		WithDiscardLogger(),
		WithWarnings())

	mb.Bot = bot

	return mb
}

type testNamedReader struct{}

func (t testNamedReader) Read(_ []byte) (n int, err error) {
	panic("unreachable: testNamedReader.Read")
}

func (t testNamedReader) Name() string {
	return "test"
}

var testInputFile = InputFile{
	File: testNamedReader{},
}

var testInputNilFile = InputFile{
	File: nil,
}
