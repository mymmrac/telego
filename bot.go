package telego

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/telegoapi"
)

const (
	defaultBotAPIServer = "https://api.telegram.org"

	tokenRegexp = `^\d{9,10}:[\w-]{35}$` //nolint:gosec

	attachFile = `attach://`

	omitEmptySuffix = ",omitempty"
)

// ErrInvalidToken Bot token is invalid according to token regexp
var ErrInvalidToken = errors.New("invalid token")

// validateToken validates if token matches format
func validateToken(token string) bool {
	reg := regexp.MustCompile(tokenRegexp)
	return reg.MatchString(token)
}

// Bot represents Telegram bot
type Bot struct {
	token          string
	apiURL         string
	log            Logger
	api            telegoapi.Caller
	constructor    telegoapi.RequestConstructor
	updateInterval time.Duration

	stop               chan struct{}
	startedLongPulling bool
	startedWebhook     bool

	server *fasthttp.Server
}

// BotOption represents option that can be applied to Bot
type BotOption func(bot *Bot) error

// NewBot creates new bot with given options. If no options specified default values are used.
// Note: Default logger (that logs only errors if not configured) will hide your bot token, but it still may log
// sensitive information, it's only safe to use default logger in testing environment.
func NewBot(token string, options ...BotOption) (*Bot, error) {
	if !validateToken(token) {
		return nil, ErrInvalidToken
	}

	b := &Bot{
		token:          token,
		apiURL:         defaultBotAPIServer,
		log:            newDefaultLogger(token),
		api:            telegoapi.FasthttpAPICaller{Client: &fasthttp.Client{}},
		constructor:    telegoapi.DefaultConstructor{},
		updateInterval: defaultUpdateInterval,

		server: &fasthttp.Server{},
	}

	for _, option := range options {
		if err := option(b); err != nil {
			return nil, fmt.Errorf("options: %w", err)
		}
	}

	return b, nil
}

// Token returns bot token
func (b *Bot) Token() string {
	return b.token
}

// performRequest executes and parses response of method
func (b *Bot) performRequest(methodName string, parameters, v interface{}) error {
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
		err = json.Unmarshal(resp.Result, &v)
		if err != nil {
			return fmt.Errorf("unmarshal to %s: %w", reflect.TypeOf(v), err)
		}
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

		i := 0
		debugFiles := make([]string, len(filesParams))
		for k, v := range filesParams {
			debugFiles[i] = fmt.Sprintf("%q: %q", k, v.Name())
			i++
		}
		//nolint:errcheck
		debugJSON, _ := json.Marshal(parsedParameters)

		debug.WriteString(fmt.Sprintf("parameters: %s, files: {%s}", debugJSON, strings.Join(debugFiles, ", ")))
	} else {
		var err error
		data, err = b.constructor.JSONRequest(parameters)
		if err != nil {
			return nil, fmt.Errorf("json request: %w", err)
		}

		debug.WriteString(data.Buffer.String())
	}

	url := b.apiURL + "/bot" + b.token + "/" + methodName

	debugData := strings.TrimSuffix(debug.String(), "\n")
	b.log.Debugf("API call to: %q, with data: %s", url, debugData)

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

// parseParameters parses parameter struct to key value structure
func parseParameters(v interface{}) (map[string]string, error) {
	valueOfV := reflect.ValueOf(v)
	if valueOfV.Kind() != reflect.Ptr && valueOfV.Kind() != reflect.Interface {
		return nil, fmt.Errorf("%s not a pointer or interface", valueOfV.Kind())
	}

	paramsStruct := valueOfV.Elem()
	if paramsStruct.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s not a struct", paramsStruct.Kind())
	}
	paramsStructType := paramsStruct.Type()

	params := make(map[string]string)

	for i := 0; i < paramsStruct.NumField(); i++ {
		structField := paramsStructType.Field(i)
		field := paramsStruct.Field(i)

		stringValue, ok, err := parseField(field)
		if err != nil {
			return nil, fmt.Errorf("parse field: %w", err)
		}
		if !ok {
			continue
		}

		key := structField.Tag.Get("json")
		key = strings.TrimSuffix(key, omitEmptySuffix)
		params[key] = stringValue
	}

	return params, nil
}

// parseField parses struct field to string value
func parseField(field reflect.Value) (string, bool, error) {
	if field.IsZero() {
		return "", false, nil
	}

	value := field.Interface()
	var stringValue string

	switch field.Kind() {
	case reflect.Struct, reflect.Slice, reflect.Interface, reflect.Ptr:
		buf := bytes.Buffer{}

		if err := json.NewEncoder(&buf).Encode(value); err != nil {
			return "", false, fmt.Errorf("encoding json: %w", err)
		}

		stringValue = buf.String()
		stringValue = strings.TrimSuffix(stringValue, "\n")
		if len(stringValue) == 0 {
			return "", false, nil
		}

		if len(stringValue) >= 2 && stringValue[0] == '"' && stringValue[len(stringValue)-1] == '"' {
			stringValue = stringValue[1 : len(stringValue)-1]
		}
	default:
		stringValue = fmt.Sprintf("%v", value)
	}

	return stringValue, true, nil
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
