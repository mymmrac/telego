package telego

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// json - Jsoniter replacement for json package
var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	defaultBotAPIServer = "https://api.telegram.org"

	contentTypeJSON = "application/json"

	tokenRegexp = `^\d{9,10}:[\w-]{35}$` //nolint:gosec

	attachFile = `attach://`

	omitEmptySuffix = ",omitempty"
)

var (
	// ErrInvalidToken - Bot token is invalid according to token regexp
	ErrInvalidToken = errors.New("invalid token")
)

// validateToken - Validates if token matches format
func validateToken(token string) bool {
	reg := regexp.MustCompile(tokenRegexp)
	return reg.MatchString(token)
}

// Bot - Represents Telegram bot
type Bot struct {
	token          string
	apiURL         string
	log            Logger
	api            apiCaller
	constructor    requestConstructor
	updateInterval time.Duration

	stopChannel    chan struct{}
	webhookHandler fasthttp.RequestHandler
}

// BotOption - Represents option that can be applied to Bot
type BotOption func(bot *Bot) error

// NewBot - Creates new bot with given options. If no options specified default values are used
func NewBot(token string, options ...BotOption) (*Bot, error) {
	if !validateToken(token) {
		return nil, ErrInvalidToken
	}

	b := &Bot{
		token:          token,
		apiURL:         defaultBotAPIServer,
		log:            newLogger(),
		api:            fasthttpAPICaller{Client: &fasthttp.Client{}},
		constructor:    defaultConstructor{},
		updateInterval: defaultUpdateInterval,
	}

	for _, option := range options {
		if err := option(b); err != nil {
			return nil, fmt.Errorf("optins: %w", err)
		}
	}

	return b, nil
}

// Token - Returns bot token
func (b *Bot) Token() string {
	return b.token
}

// performRequest - Executes and parses response of method
func (b *Bot) performRequest(methodName string, parameters, v interface{}) error {
	resp, err := b.constructAndCallRequest(methodName, parameters)
	if err != nil {
		b.log.Errorf("Execution error %s: %s", methodName, err)
		return fmt.Errorf("internal execution: %w", err)
	}
	b.log.Debugf("API response %s: %s", methodName, resp.String())

	if !resp.Ok {
		return fmt.Errorf("api: %w", resp.APIError)
	}

	if resp.Result != nil {
		err = json.Unmarshal(resp.Result, &v)
		if err != nil {
			return fmt.Errorf("unmarshal to %s: %w", reflect.TypeOf(v), err)
		}
	}

	return nil
}

// constructAndCallRequest - Creates and executes request with parsing of parameters
func (b *Bot) constructAndCallRequest(methodName string, parameters interface{}) (*apiResponse, error) {
	filesParams, hasFiles := filesParameters(parameters)
	var data *requestData

	if hasFiles {
		parsedParameters, err := parseParameters(parameters)
		if err != nil {
			return nil, fmt.Errorf("parsing parameters: %w", err)
		}

		data, err = b.constructor.MultipartRequest(parsedParameters, filesParams)
		if err != nil {
			return nil, fmt.Errorf("multipart request: %w", err)
		}
	} else {
		var err error
		data, err = b.constructor.JSONRequest(parameters)
		if err != nil {
			return nil, fmt.Errorf("json request: %w", err)
		}
	}

	url := b.apiURL + "/bot" + b.token + "/" + methodName
	resp, err := b.api.Call(url, data)
	if err != nil {
		return nil, fmt.Errorf("request call: %w", err)
	}

	return resp, nil
}

// filesParameters - Gets all files from parameters
func filesParameters(parameters interface{}) (files map[string]*os.File, hasFiles bool) {
	if parametersWithFiles, ok := parameters.(fileCompatible); ok {
		files = parametersWithFiles.fileParameters()
		for _, file := range files {
			if file != nil {
				hasFiles = true
				break
			}
		}
	}
	return files, hasFiles
}

// parseParameters - Parses parameter struct to key value structure
func parseParameters(v interface{}) (map[string]string, error) {
	paramsStruct := reflect.ValueOf(v).Elem()
	if paramsStruct.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s not a struct", paramsStruct.Kind())
	}
	paramsStructType := paramsStruct.Type()

	params := make(map[string]string)

	for i := 0; i < paramsStruct.NumField(); i++ {
		structField := paramsStructType.Field(i)
		field := paramsStruct.Field(i)

		if field.IsZero() {
			continue
		}

		if field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface {
			field = field.Elem()
		}

		value := field.Interface()
		var stringValue string

		switch field.Kind() {
		case reflect.Struct, reflect.Slice:
			buf := bytes.Buffer{}

			if err := json.NewEncoder(&buf).Encode(value); err != nil {
				return nil, fmt.Errorf("encoding json: %w", err)
			}

			stringValue = buf.String()
			stringValue = strings.TrimSuffix(stringValue, "\n")
		default:
			stringValue = fmt.Sprintf("%v", value)
		}

		if stringValue == "" {
			continue
		}

		key := structField.Tag.Get("json")
		key = strings.TrimSuffix(key, omitEmptySuffix)
		params[key] = stringValue
	}

	return params, nil
}
