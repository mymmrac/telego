package telego

import (
	stdJson "encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testToken    = "1234567890:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc"
	invalidToken = "abc"
)

func getBot(t *testing.T) *Bot {
	bot, err := NewBot(testToken)
	assert.NoError(t, err)

	return bot
}

func Test_validateToken(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		isValid bool
	}{
		{
			name:    "empty",
			token:   "",
			isValid: false,
		},
		{
			name:    "not valid",
			token:   invalidToken,
			isValid: false,
		},
		{
			name:    "valid 1",
			token:   testToken,
			isValid: true,
		},
		{
			name:    "valid 2",
			token:   "123456789:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc",
			isValid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := validateToken(tt.token)
			assert.Equal(t, tt.isValid, actual)
		})
	}
}

func TestNewBot(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bot, err := NewBot(testToken)

		assert.NoError(t, err)
		assert.NotNil(t, bot)
	})

	t.Run("error", func(t *testing.T) {
		bot, err := NewBot(invalidToken)

		assert.Error(t, err)
		assert.Nil(t, bot)
	})
}

// func TestBot_Logger(t *testing.T) {
// 	bot := getBot(t)
//
// 	t.Run("default-logger", func(t *testing.T) {
// 		assert.NotPanics(t, func() {
// 			bot.DefaultLogger(true, true)
// 		})
// 	})
//
// 	t.Run("set-logger", func(t *testing.T) {
// 		assert.NotPanics(t, func() {
// 			var l Logger
// 			bot.SetLogger(l)
// 		})
// 	})
// }

// func TestBot_SetToken(t *testing.T) {
// 	bot := getBot(t)
//
// 	tests := []struct {
// 		name  string
// 		token string
// 		err   error
// 	}{
// 		{
// 			name:  "success",
// 			token: testToken,
// 			err:   nil,
// 		},
// 		{
// 			name:  "error",
// 			token: invalidToken,
//			err:   ErrInvalidToken,
//		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			actual := bot.SetToken(tt.token)
// 			assert.Equal(t, tt.err, actual)
// 		})
// 	}
// }

func TestBot_Token(t *testing.T) {
	bot := getBot(t)

	assert.Equal(t, testToken, bot.Token())
}

// func TestBot_SetAPIServer(t *testing.T) {
// 	bot := getBot(t)
//
// 	tests := []struct {
// 		name  string
// 		url   string
// 		isErr bool
// 	}{
// 		{
// 			name:  "success",
// 			url:   defaultBotAPIServer,
// 			isErr: false,
// 		},
// 		{
// 			name:  "empty",
// 			url:   "",
// 			isErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			actual := bot.SetAPIServer(tt.url)
// 			if tt.isErr {
// 				assert.Error(t, actual)
// 				return
// 			}
// 			assert.NoError(t, actual)
// 		})
// 	}
// }

// func TestBot_SetClient(t *testing.T) {
// 	bot := getBot(t)
//
// 	tests := []struct {
// 		name   string
// 		client *fasthttp.Client
// 		isErr  bool
// 	}{
// 		{
// 			name:   "success",
// 			client: &fasthttp.Client{},
// 	 		isErr:  false,
// 		},
// 		{
// 			name:   "error",
// 			client: nil,
// 			isErr:  true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			actual := bot.SetClient(tt.client)
// 			if tt.isErr {
// 				assert.Error(t, actual)
// 				return
// 			}
// 			assert.NoError(t, actual)
// 		})
// 	}
// }

func Test_apiResponse_String_and_APIError_Error(t *testing.T) {
	tests := []struct {
		name string
		resp apiResponse
		text string
	}{
		{
			name: "success",
			resp: apiResponse{
				Ok:       true,
				APIError: nil,
				Result:   stdJson.RawMessage{},
			},
			text: "Ok: true, Err: {<nil>}, Result: ",
		},
		{
			name: "error",
			resp: apiResponse{
				Ok: false,
				APIError: &APIError{
					Description: "bad request",
					ErrorCode:   400,
					Parameters:  nil,
				},
				Result: nil,
			},
			text: "Ok: false, Err: {400 \"bad request\"}, Result: ",
		},
		{
			name: "error with parameters",
			resp: apiResponse{
				Ok: false,
				APIError: &APIError{
					Description: "bad request",
					ErrorCode:   400,
					Parameters: &ResponseParameters{
						MigrateToChatID: 1,
						RetryAfter:      2,
					},
				},
				Result: nil,
			},
			text: "Ok: false, Err: {400 \"bad request\" migrate to chat id: 1, retry after: 2}, Result: ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.resp.String()
			assert.Equal(t, tt.text, actual)
		})
	}
}
