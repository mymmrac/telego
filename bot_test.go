package telego

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/mymmrac/telego/internal/json"
	ta "github.com/mymmrac/telego/telegoapi"
	mockapi "github.com/mymmrac/telego/telegoapi/mock"
)

const (
	token        = "1234567890:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc"
	invalidToken = "invalid-token"

	methodName = "testMethod"
)

var errTest = errors.New("error")

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
			name:    "not_valid",
			token:   invalidToken,
			isValid: false,
		},
		{
			name:    "valid_1",
			token:   token,
			isValid: true,
		},
		{
			name:    "valid_2",
			token:   "123456789:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc",
			isValid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := validateToken(tt.token)
			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

func TestNewBot(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("success", func(t *testing.T) {
		bot, err := NewBot(token)

		require.NoError(t, err)
		assert.NotNil(t, bot)
	})

	t.Run("success_with_options", func(t *testing.T) {
		bot, err := NewBot(token, func(_ *Bot) error { return nil })

		require.NoError(t, err)
		assert.NotNil(t, bot)
	})

	t.Run("error", func(t *testing.T) {
		bot, err := NewBot(invalidToken)

		require.Error(t, err)
		assert.Nil(t, bot)
	})

	t.Run("error_with_options", func(t *testing.T) {
		bot, err := NewBot(token, func(_ *Bot) error { return errTest })

		require.ErrorIs(t, err, errTest)
		assert.Nil(t, bot)
	})

	t.Run("with_health_check", func(t *testing.T) {
		caller := mockapi.NewMockCaller(ctrl)
		constructor := mockapi.NewMockRequestConstructor(ctrl)

		expectedData := &ta.RequestData{
			ContentType: ta.ContentTypeJSON,
			Buffer:      bytes.NewBuffer([]byte{}),
		}

		t.Run("success", func(t *testing.T) {
			expectedResp := &ta.Response{
				Ok:     true,
				Result: json.RawMessage(`{}`),
			}

			constructor.EXPECT().
				JSONRequest(nil).
				Return(expectedData, nil).
				Times(1)

			caller.EXPECT().
				Call(testCtx, defaultBotAPIServer+botPathPrefix+token+"/getMe", expectedData).
				Return(expectedResp, nil).
				Times(1)

			bot, err := NewBot(token,
				WithAPICaller(caller),
				WithRequestConstructor(constructor),
				WithHealthCheck(testCtx),
			)

			require.NoError(t, err)
			assert.NotNil(t, bot)
		})

		t.Run("error", func(t *testing.T) {
			expectedResp := &ta.Response{
				Ok:    false,
				Error: &ta.Error{},
			}

			constructor.EXPECT().
				JSONRequest(nil).
				Return(expectedData, nil).
				Times(1)

			caller.EXPECT().
				Call(testCtx, defaultBotAPIServer+botPathPrefix+token+"/getMe", expectedData).
				Return(expectedResp, nil).
				Times(1)

			bot, err := NewBot(token,
				WithAPICaller(caller),
				WithRequestConstructor(constructor),
				WithHealthCheck(testCtx),
			)

			require.Error(t, err)
			assert.Nil(t, bot)
		})
	})
}

func TestBot_Token(t *testing.T) {
	bot, err := NewBot(token)
	require.NoError(t, err)

	assert.Equal(t, token, bot.Token())
}

func TestBot_Logger(t *testing.T) {
	bot, err := NewBot(token)
	require.NoError(t, err)

	assert.Equal(t, bot.log, bot.Logger())
}

func TestBot_FileDownloadURL(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		bot, err := NewBot(token)
		require.NoError(t, err)

		filepath := "file.txt"
		url := bot.FileDownloadURL(filepath)
		assert.Equal(t, bot.apiURL+"/file"+botPathPrefix+bot.token+"/"+filepath, url)
	})

	t.Run("test", func(t *testing.T) {
		bot, err := NewBot(token, WithTestServerPath())
		require.NoError(t, err)

		filepath := "file.txt"
		url := bot.FileDownloadURL(filepath)
		assert.Equal(t, bot.apiURL+"/file"+botPathPrefix+bot.token+"/test/"+filepath, url)
	})
}

type testErrorMarshal struct {
	Number int `json:"number"`
}

func (t testErrorMarshal) MarshalJSON() ([]byte, error) {
	return nil, errTest
}

type testEmptyMarshal struct {
	Number int `json:"number"`
}

func (t testEmptyMarshal) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}

func Test_parseParameters(t *testing.T) {
	n := 1

	tests := []struct {
		name             string
		parameters       any
		parsedParameters map[string]string
		isError          bool
	}{
		{
			name: "success",
			parameters: &struct {
				Empty       string    `json:"empty,omitempty"`
				EmptyNoOmit string    `json:"empty_no_omit"`
				Number      int       `json:"number"`
				Array       []int     `json:"array"`
				Text        string    `json:"text"`
				Struct      *struct { //revive:disable:nested-structs
					N int `json:"n"`
				} `json:"struct"`
			}{
				Number: 10,
				Array:  []int{1, 2, 3},
				Text:   "ok",
				Struct: &struct {
					N int `json:"n"`
				}{2},
			},
			parsedParameters: map[string]string{
				"number": "10",
				"array":  "[1,2,3]",
				"struct": "{\"n\":2}",
				"text":   "ok",
			},
			isError: false,
		},
		{
			name: "error_not_pointer",
			parameters: struct {
				a int
			}{},
			parsedParameters: nil,
			isError:          true,
		},
		{
			name:             "error_not_struct",
			parameters:       &n,
			parsedParameters: nil,
			isError:          true,
		},
		{
			name: "error_no_tag",
			parameters: &struct {
				Number int
			}{
				Number: 1,
			},
			parsedParameters: nil,
			isError:          true,
		},
		{
			name: "error_marshal",
			parameters: &struct {
				NonMarshaled testErrorMarshal `json:"non_marshaled"`
			}{
				NonMarshaled: testErrorMarshal{1},
			},
			parsedParameters: nil,
			isError:          true,
		},
		{
			name: "success_get_update",
			parameters: &GetUpdatesParams{
				Offset:         1,
				Limit:          2,
				Timeout:        3,
				AllowedUpdates: []string{"ok1", "ok2"},
			},
			parsedParameters: map[string]string{
				"offset":          "1",
				"limit":           "2",
				"timeout":         "3",
				"allowed_updates": "[\"ok1\",\"ok2\"]",
			},
			isError: false,
		},
		{
			name: "success_send_photo",
			parameters: &SendPhotoParams{
				ChatID:              ChatID{ID: 1},
				Photo:               InputFile{URL: "ok1"},
				Caption:             "ok2",
				DisableNotification: true,
				ReplyMarkup: &InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						{
							{
								Text: "ok3",
							},
						},
					},
				},
			},
			parsedParameters: map[string]string{
				"caption":              "ok2",
				"chat_id":              "1",
				"disable_notification": "true",
				"photo":                "ok1",
				"reply_markup":         "{\"inline_keyboard\":[[{\"text\":\"ok3\"}]]}",
			},
			isError: false,
		},
		{
			name: "success_empty_marshal",
			parameters: &struct {
				EmptyMarshaled testEmptyMarshal `json:"empty_marshaled"`
			}{
				EmptyMarshaled: testEmptyMarshal{1},
			},
			parsedParameters: map[string]string{},
			isError:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedParameters, err := parseParameters(tt.parameters)
			if tt.isError {
				require.Error(t, err)
				assert.Nil(t, parsedParameters)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.parsedParameters, parsedParameters)
		})
	}
}

type testStruct struct{}

func (ts *testStruct) fileParameters() map[string]ta.NamedReader {
	return map[string]ta.NamedReader{
		"test": &testNamedReade{},
	}
}

func Test_filesParameters(t *testing.T) {
	tests := []struct {
		name       string
		parameters any
		files      map[string]ta.NamedReader
		hasFiles   bool
	}{
		{
			name:       "with_files",
			parameters: &testStruct{},
			files: map[string]ta.NamedReader{
				"test": &testNamedReade{},
			},
			hasFiles: true,
		},
		{
			name:       "no_files",
			parameters: 1,
			files:      nil,
			hasFiles:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, hasFiles := filesParameters(tt.parameters)
			assert.Equal(t, tt.hasFiles, hasFiles)
			assert.Equal(t, tt.files, files)
		})
	}
}

type paramsWithFile struct {
	N int `json:"n"`
}

func (p *paramsWithFile) fileParameters() map[string]ta.NamedReader {
	return map[string]ta.NamedReader{
		"test": &testNamedReade{},
	}
}

type notStructParamsWithFile string

func (p *notStructParamsWithFile) fileParameters() map[string]ta.NamedReader {
	return map[string]ta.NamedReader{
		"test": &testNamedReade{},
	}
}

func TestBot_constructAndCallRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	params := struct {
		N int `json:"n"`
	}{
		N: 1,
	}

	url := m.Bot.apiURL + botPathPrefix + m.Bot.token + "/" + methodName

	expectedResp := &ta.Response{
		Ok: true,
	}

	paramsBytes, err := json.Marshal(params)
	require.NoError(t, err)

	expectedData := &ta.RequestData{
		ContentType: ta.ContentTypeJSON,
		Buffer:      bytes.NewBuffer(paramsBytes),
	}

	t.Run("success_json", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(params).
			Return(expectedData, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(testCtx, url, expectedData).
			Return(expectedResp, nil).
			Times(1)

		resp, err := m.Bot.constructAndCallRequest(testCtx, methodName, params)
		require.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})

	t.Run("error_json", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(params).
			Return(nil, errTest).
			Times(1)

		resp, err := m.Bot.constructAndCallRequest(testCtx, methodName, params)
		require.ErrorIs(t, err, errTest)
		assert.Nil(t, resp)
	})

	t.Run("success_multipart", func(t *testing.T) {
		paramsFile := &paramsWithFile{N: 1}
		paramsMap := map[string]string{
			"n": "1",
		}

		paramsBytesFile, err := json.Marshal(paramsFile)
		require.NoError(t, err)

		expectedDataFile := &ta.RequestData{
			ContentType: ta.ContentTypeJSON,
			Buffer:      bytes.NewBuffer(paramsBytesFile),
		}

		m.MockRequestConstructor.EXPECT().
			MultipartRequest(paramsMap, gomock.Any()).
			Return(expectedDataFile, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(testCtx, url, expectedDataFile).
			Return(expectedResp, nil).
			Times(1)

		resp, err := m.Bot.constructAndCallRequest(testCtx, methodName, paramsFile)
		require.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})

	t.Run("error_multipart", func(t *testing.T) {
		paramsFile := &paramsWithFile{N: 1}
		paramsMap := map[string]string{
			"n": "1",
		}

		m.MockRequestConstructor.EXPECT().
			MultipartRequest(paramsMap, gomock.Any()).
			Return(nil, errTest).
			Times(1)

		resp, err := m.Bot.constructAndCallRequest(testCtx, methodName, paramsFile)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_multipart_params", func(t *testing.T) {
		notStruct := notStructParamsWithFile("test")

		resp, err := m.Bot.constructAndCallRequest(testCtx, methodName, &notStruct)
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_call", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(params).
			Return(expectedData, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(testCtx, url, expectedData).
			Return(nil, errTest).
			Times(1)

		resp, err := m.Bot.constructAndCallRequest(testCtx, methodName, params)
		require.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestBot_performRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	params := struct {
		N int `json:"n"`
	}{
		N: 1,
	}

	t.Run("success", func(t *testing.T) {
		var result int

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(&ta.RequestData{}, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&ta.Response{
				Ok:     true,
				Result: bytes.NewBufferString("1").Bytes(),
				Error:  nil,
			}, nil)

		err := m.Bot.performRequest(testCtx, methodName, params, &result)
		require.NoError(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("success_unmarshal_second", func(t *testing.T) {
		var result1 int
		var result2 bool

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(&ta.RequestData{}, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&ta.Response{
				Ok:     true,
				Result: bytes.NewBufferString("true").Bytes(),
				Error:  nil,
			}, nil)

		err := m.Bot.performRequest(testCtx, methodName, params, &result1, &result2)
		require.NoError(t, err)
		assert.Equal(t, 0, result1)
		assert.True(t, result2)
	})

	t.Run("error_not_ok", func(t *testing.T) {
		var result int

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(&ta.RequestData{}, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&ta.Response{
				Ok:     false,
				Result: nil,
				Error:  &ta.Error{},
			}, nil)

		err := m.Bot.performRequest(testCtx, methodName, params, &result)
		require.Error(t, err)
	})

	t.Run("error_construct_and_call", func(t *testing.T) {
		var result int

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest).
			Times(1)

		err := m.Bot.performRequest(testCtx, methodName, params, &result)
		require.Error(t, err)
	})

	t.Run("error_unmarshal", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(&ta.RequestData{}, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&ta.Response{
				Ok:     true,
				Result: bytes.NewBufferString("1").Bytes(),
				Error:  nil,
			}, nil)

		var stringResult string
		err := m.Bot.performRequest(testCtx, methodName, params, &stringResult)
		require.Error(t, err)
		assert.Equal(t, "", stringResult)
	})

	t.Run("error_warning", func(t *testing.T) {
		var result int

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(&ta.RequestData{}, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&ta.Response{
				Ok:     true,
				Result: bytes.NewBufferString("1").Bytes(),
				Error:  &ta.Error{ErrorCode: 1},
			}, nil)

		err := m.Bot.performRequest(testCtx, methodName, params, &result)
		assert.Equal(t, &ta.Error{ErrorCode: 1}, err)
		assert.Equal(t, 1, result)
	})
}

func Test_isNil(t *testing.T) {
	var n *int
	a := 1
	m := &a

	tests := []struct {
		name  string
		i     any
		isNil bool
	}{
		{
			name:  "nil",
			i:     nil,
			isNil: true,
		},
		{
			name:  "nil_ptr",
			i:     n,
			isNil: true,
		},
		{
			name:  "value",
			i:     m,
			isNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isNil, isNil(tt.i))
		})
	}
}

func TestToPtr(t *testing.T) {
	assert.True(t, *ToPtr(true))
	assert.False(t, *ToPtr(false))

	assert.Equal(t, "", *ToPtr(""))
	assert.Equal(t, "a", *ToPtr("a"))
}
