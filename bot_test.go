package telego

import (
	"bytes"
	"errors"
	"testing"

	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego/telegoapi"
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
	t.Run("success", func(t *testing.T) {
		bot, err := NewBot(token)

		assert.NoError(t, err)
		assert.NotNil(t, bot)
	})

	t.Run("success_with_options", func(t *testing.T) {
		bot, err := NewBot(token, func(_ *Bot) error { return nil })

		assert.NoError(t, err)
		assert.NotNil(t, bot)
	})

	t.Run("error", func(t *testing.T) {
		bot, err := NewBot(invalidToken)

		assert.Error(t, err)
		assert.Nil(t, bot)
	})

	t.Run("error_with_options", func(t *testing.T) {
		bot, err := NewBot(token, func(_ *Bot) error { return errTest })

		assert.ErrorIs(t, err, errTest)
		assert.Nil(t, bot)
	})
}

func TestBot_Token(t *testing.T) {
	bot, err := NewBot(token)
	assert.NoError(t, err)

	assert.Equal(t, token, bot.Token())
}

func Test_parseParameters(t *testing.T) {
	n := 1

	tests := []struct {
		name             string
		parameters       interface{}
		parsedParameters map[string]string
		isError          bool
	}{
		{
			name: "success",
			parameters: &struct {
				Empty  string `json:"empty,omitempty"`
				Number int    `json:"number"`
				Array  []int  `json:"array"`
				Struct *struct {
					N int `json:"n"`
				} `json:"struct"`
			}{
				Number: 10,
				Array:  []int{1, 2, 3},
				Struct: &struct {
					N int `json:"n"`
				}{2},
			},
			parsedParameters: map[string]string{
				"number": "10",
				"array":  "[1,2,3]",
				"struct": "{\"n\":2}",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedParameters, err := parseParameters(tt.parameters)
			if tt.isError {
				assert.Error(t, err)
				assert.Nil(t, parsedParameters)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.parsedParameters, parsedParameters)
		})
	}
}

type testStruct struct{}

func (ts *testStruct) fileParameters() map[string]telegoapi.NamedReader {
	return map[string]telegoapi.NamedReader{
		"test": &testNamedReade{},
	}
}

func Test_filesParameters(t *testing.T) {
	tests := []struct {
		name       string
		parameters interface{}
		files      map[string]telegoapi.NamedReader
		hasFiles   bool
	}{
		{
			name:       "with_files",
			parameters: &testStruct{},
			files: map[string]telegoapi.NamedReader{
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

func (p *paramsWithFile) fileParameters() map[string]telegoapi.NamedReader {
	return map[string]telegoapi.NamedReader{
		"test": &testNamedReade{},
	}
}

type notStructParamsWithFile string

func (p *notStructParamsWithFile) fileParameters() map[string]telegoapi.NamedReader {
	return map[string]telegoapi.NamedReader{
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

	url := m.Bot.apiURL + "/bot" + m.Bot.token + "/" + methodName

	expectedResp := &telegoapi.Response{
		Ok: true,
	}

	paramsBytes, err := json.Marshal(params)
	assert.NoError(t, err)

	expectedData := &telegoapi.RequestData{
		ContentType: telegoapi.ContentTypeJSON,
		Buffer:      bytes.NewBuffer(paramsBytes),
	}

	t.Run("success_json", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(params).
			Return(expectedData, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(url, expectedData).
			Return(expectedResp, nil).
			Times(1)

		resp, err := m.Bot.constructAndCallRequest(methodName, params)
		assert.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})

	t.Run("error_json", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(params).
			Return(nil, errTest).
			Times(1)

		resp, err := m.Bot.constructAndCallRequest(methodName, params)
		assert.ErrorIs(t, err, errTest)
		assert.Nil(t, resp)
	})

	t.Run("success_multipart", func(t *testing.T) {
		paramsFile := &paramsWithFile{N: 1}
		paramsMap := map[string]string{
			"n": "1",
		}

		paramsBytesFile, err := json.Marshal(paramsFile)
		assert.NoError(t, err)

		expectedDataFile := &telegoapi.RequestData{
			ContentType: telegoapi.ContentTypeJSON,
			Buffer:      bytes.NewBuffer(paramsBytesFile),
		}

		m.MockRequestConstructor.EXPECT().
			MultipartRequest(paramsMap, gomock.Any()).
			Return(expectedDataFile, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(url, expectedDataFile).
			Return(expectedResp, nil).
			Times(1)

		resp, err := m.Bot.constructAndCallRequest(methodName, paramsFile)
		assert.NoError(t, err)
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

		resp, err := m.Bot.constructAndCallRequest(methodName, paramsFile)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_multipart_params", func(t *testing.T) {
		notStruct := notStructParamsWithFile("test")

		resp, err := m.Bot.constructAndCallRequest(methodName, &notStruct)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_call", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(params).
			Return(expectedData, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(url, expectedData).
			Return(nil, errTest).
			Times(1)

		resp, err := m.Bot.constructAndCallRequest(methodName, params)
		assert.Error(t, err)
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

	var result int

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(&telegoapi.RequestData{}, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any()).
			Return(&telegoapi.Response{
				Ok:     true,
				Result: bytes.NewBufferString("1").Bytes(),
				Error:  nil,
			}, nil)

		err := m.Bot.performRequest(methodName, params, &result)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("error_not_ok", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(&telegoapi.RequestData{}, nil).
			Times(1)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any()).
			Return(&telegoapi.Response{
				Ok:     false,
				Result: nil,
				Error:  &telegoapi.Error{},
			}, nil)

		err := m.Bot.performRequest(methodName, params, &result)
		assert.Error(t, err)
	})

	t.Run("error_construct_and_call", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest).
			Times(1)

		err := m.Bot.performRequest(methodName, params, &result)
		assert.Error(t, err)
	})
}

func Test_isNil(t *testing.T) {
	var n *int
	a := 1
	m := &a

	tests := []struct {
		name  string
		i     interface{}
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
