package telego

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/telegoapi"
)

const (
	defaultBotAPIServer  = "https://api.telegram.org"
	defaultBotEmptyValue = "TELEGO_EMPTY_VALUE"

	tokenRegexp = `^\d{9,10}:[\w-]{35}$` //nolint:gosec

	attachFile = `attach://`

	omitEmptySuffix = ",omitempty"
)

// ErrInvalidToken Bot token is invalid according to token regexp
var ErrInvalidToken = errors.New("telego: invalid token")

// validateToken validates if token matches format
func validateToken(token string) bool {
	reg := regexp.MustCompile(tokenRegexp)
	return reg.MatchString(token)
}

// Bot represents Telegram bot
type Bot struct {
	token       string
	apiURL      string
	log         Logger
	api         telegoapi.Caller
	constructor telegoapi.RequestConstructor

	healthCheckRequested bool
	warningAsErrors      bool
	replaceToEmpty       string

	longPollingContext *longPollingContext
	webhookContext     *webhookContext
}

// BotOption represents an option that can be applied to Bot
type BotOption func(bot *Bot) error

// NewBot creates new bots with given options.
// If no options specified default values are used.
// Note: Default logger (that logs only errors if not configured) will hide your bot token, but it still may log
// sensitive information, it's only safe to use default logger in testing environment.
func NewBot(token string, options ...BotOption) (*Bot, error) {
	if !validateToken(token) {
		return nil, ErrInvalidToken
	}

	b := &Bot{
		token:       token,
		apiURL:      defaultBotAPIServer,
		log:         newDefaultLogger(token),
		api:         telegoapi.FastHTTPCaller{Client: &fasthttp.Client{}},
		constructor: telegoapi.DefaultConstructor{},
	}

	for _, option := range options {
		if err := option(b); err != nil {
			return nil, fmt.Errorf("telego: options: %w", err)
		}
	}

	if b.healthCheckRequested {
		if _, err := b.GetMe(); err != nil {
			return nil, fmt.Errorf("telego: health check: %w", err)
		}
	}

	return b, nil
}

// Token returns bot token
func (b *Bot) Token() string {
	return b.token
}

// EmptyValue returns value that will be erased from all requests useful for things like SwitchInlineQuery in
// telego.InlineKeyboardButton that have empty string as valid parameter value
// Warning: Only works if at least one of the bot options, WithEmptyValues or WithCustomEmptyValues are used
func (b *Bot) EmptyValue() string {
	return b.replaceToEmpty
}

// performRequest executes and parses response of method
func (b *Bot) performRequest(methodName string, parameters interface{}, vs ...interface{}) error {
	resp, err := b.constructAndCallRequest(methodName, parameters)
	if err != nil {
		b.log.Errorf("Execution error %s: %s", methodName, err)
		return fmt.Errorf("internal execution: %w", err)
	}
	b.log.Debugf("API response %s: %s", methodName, resp.String())

	if !resp.Ok {
		return fmt.Errorf("api: %w", resp.Error)
	}

	if resp.Result != nil {
		var unmarshalErr error
		for i := range vs {
			unmarshalErr = json.Unmarshal(resp.Result, &vs[i])
			if unmarshalErr == nil {
				break
			}
		}

		if unmarshalErr != nil {
			return fmt.Errorf("unmarshal to %s: %w", reflect.TypeOf(vs[len(vs)-1]), unmarshalErr)
		}
	}

	if b.warningAsErrors && resp.Error != nil {
		return resp.Error
	}

	return nil
}

// constructAndCallRequest creates and executes request with parsing of parameters
func (b *Bot) constructAndCallRequest(methodName string, parameters interface{}) (*telegoapi.Response, error) {
	filesParams, hasFiles := filesParameters(parameters)
	var data *telegoapi.RequestData

	debug := strings.Builder{}

	if hasFiles {
		parsedParameters, err := parseParameters(parameters)
		if err != nil {
			return nil, fmt.Errorf("parsing parameters: %w", err)
		}

		data, err = b.constructor.MultipartRequest(parsedParameters, filesParams)
		if err != nil {
			return nil, fmt.Errorf("multipart request: %w", err)
		}

		logRequestWithFiles(debug, parsedParameters, filesParams)
	} else {
		var err error
		data, err = b.constructor.JSONRequest(parameters)
		if err != nil {
			return nil, fmt.Errorf("json request: %w", err)
		}

		_, _ = debug.WriteString(data.Buffer.String())
	}

	url := b.apiURL + "/bot" + b.token + "/" + methodName

	debugData := strings.TrimSuffix(debug.String(), "\n")
	b.log.Debugf("API call to: %q, with data: %s", url, debugData)

	if b.replaceToEmpty != "" {
		data.Buffer = bytes.NewBuffer(bytes.ReplaceAll(data.Buffer.Bytes(), []byte(b.replaceToEmpty), []byte{}))
	}

	resp, err := b.api.Call(url, data)
	if err != nil {
		return nil, fmt.Errorf("request call: %w", err)
	}

	return resp, nil
}

// filesParameters gets all files from parameters
func filesParameters(parameters interface{}) (files map[string]telegoapi.NamedReader, hasFiles bool) {
	if parametersWithFiles, ok := parameters.(fileCompatible); ok && !isNil(parameters) {
		files = parametersWithFiles.fileParameters()
		for _, file := range files {
			if !isNil(file) {
				hasFiles = true
				break
			}
		}
	}
	return files, hasFiles
}

// parseParameters parses parameter struct to key value structure, v should be a pointer to struct
func parseParameters(v interface{}) (map[string]string, error) {
	valueOfV := reflect.ValueOf(v)
	if valueOfV.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("%q not a pointer", valueOfV.Kind())
	}

	paramsStruct := valueOfV.Elem()
	if paramsStruct.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%q not a struct", paramsStruct.Kind())
	}
	paramsStructType := paramsStruct.Type()

	params := make(map[string]string)

	for i := 0; i < paramsStructType.NumField(); i++ {
		fieldType := paramsStructType.Field(i)
		key := fieldType.Tag.Get("json")
		key = strings.TrimSuffix(key, omitEmptySuffix)
		if key == "" {
			return nil, fmt.Errorf("%s does not have `json` tag", paramsStructType.String())
		}

		fieldValue := paramsStruct.Field(i)
		value, ok, err := parseField(fieldValue)
		if err != nil {
			return nil, fmt.Errorf("parse of %s: %w", paramsStructType.String(), err)
		}
		if !ok {
			continue
		}

		params[key] = value
	}

	return params, nil
}

// parseField parses struct field to string value
func parseField(field reflect.Value) (string, bool, error) {
	if field.IsZero() || !field.CanInterface() {
		return "", false, nil
	}

	data, err := json.Marshal(field.Interface())
	if err != nil {
		return "", false, err
	}

	value := string(data)

	// Trim double quotes in strings
	if len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"' {
		value = value[1 : len(value)-1]
	}

	if len(value) == 0 {
		return "", false, nil
	}

	return value, true, nil
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	default:
		return false
	}
}

func logRequestWithFiles(debug strings.Builder, parameters map[string]string, files map[string]telegoapi.NamedReader) {
	i := 0
	debugFiles := make([]string, len(files))
	for k, v := range files {
		if k == v.Name() {
			debugFiles[i] = fmt.Sprintf("%q", k)
		} else {
			debugFiles[i] = fmt.Sprintf("%q: %q", k, v.Name())
		}
		i++
	}
	//nolint:errcheck
	debugJSON, _ := json.Marshal(parameters)

	_, _ = debug.WriteString(fmt.Sprintf("parameters: %s, files: {%s}", debugJSON, strings.Join(debugFiles, ", ")))
}
