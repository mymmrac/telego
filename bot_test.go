package telego

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testToken    = "1234567890:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc"
	invalidToken = "abc"
)

var errTest = errors.New("error")

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
			isValid := validateToken(tt.token)
			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

func TestNewBot(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bot, err := NewBot(testToken)

		assert.NoError(t, err)
		assert.NotNil(t, bot)
	})

	t.Run("success with options", func(t *testing.T) {
		bot, err := NewBot(testToken, func(_ *Bot) error { return nil })

		assert.NoError(t, err)
		assert.NotNil(t, bot)
	})

	t.Run("error", func(t *testing.T) {
		bot, err := NewBot(invalidToken)

		assert.Error(t, err)
		assert.Nil(t, bot)
	})

	t.Run("error with options", func(t *testing.T) {
		bot, err := NewBot(testToken, func(_ *Bot) error { return errTest })

		assert.ErrorIs(t, err, errTest)
		assert.Nil(t, bot)
	})
}

func TestBot_Token(t *testing.T) {
	bot := getBot(t)

	assert.Equal(t, testToken, bot.Token())
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
			name: "error not pointer",
			parameters: struct {
				a int
			}{},
			parsedParameters: nil,
			isError:          true,
		},
		{
			name:             "error not struct",
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

var testFile = &os.File{}

type testStruct struct{}

func (ts *testStruct) fileParameters() map[string]*os.File {
	return map[string]*os.File{
		"test": testFile,
	}
}

func Test_filesParameters(t *testing.T) {
	tests := []struct {
		name       string
		parameters interface{}
		files      map[string]*os.File
		hasFiles   bool
	}{
		{
			name:       "with files",
			parameters: &testStruct{},
			files: map[string]*os.File{
				"test": testFile,
			},
			hasFiles: true,
		},
		{
			name:       "no files",
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
