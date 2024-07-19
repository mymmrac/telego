package telego

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/chococola/telego/internal/json"
	ta "github.com/chococola/telego/telegoapi"
	mockapi "github.com/chococola/telego/telegoapi/mock"
)

var (
	data      = &ta.RequestData{}
	emptyResp = &ta.Response{
		Ok: true,
	}

	expectedMessage = &Message{
		MessageID: 1,
	}

	testPortStart = 3100
	testPortLock  = sync.Mutex{}
)

func testAddress(t *testing.T) string {
	t.Helper()

	testPortLock.Lock()
	defer testPortLock.Unlock()

	testPortStart++
	return fmt.Sprintf("127.0.0.1:%d", testPortStart)
}

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
	bot, _ := NewBot(token,
		WithAPICaller(mb.MockAPICaller),
		WithRequestConstructor(mb.MockRequestConstructor),
		WithDiscardLogger(),
		WithWarnings())

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
