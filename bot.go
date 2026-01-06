package telego

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego/internal/json"
	ta "github.com/mymmrac/telego/telegoapi"
)

const (
	defaultBotAPIServer = "https://api.telegram.org"

	tokenRegexp = `^\d+:[\w-]{35}$` //nolint:gosec

	attachFile = `attach://`

	botPathPrefix = "/bot"
)

// ErrEmptyToken bot token is empty
var ErrEmptyToken = errors.New("telego: empty token")

// ErrInvalidToken bot token is invalid according to token regexp
var ErrInvalidToken = errors.New("telego: invalid token format")

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
	api         ta.Caller
	constructor ta.RequestConstructor

	useTestServerPath     bool
	reportWarningAsErrors bool

	running atomic.Int32

	myOnce     sync.Once
	myID       int64
	myUsername string
}

// Bot actions
const (
	runningNone        = 0
	runningLongPolling = 1
	runningWebhook     = 2
)

// run checks if bot is already running some action
func (b *Bot) run(actionToRun int32) error {
	if b.running.CompareAndSwap(runningNone, actionToRun) {
		return nil
	}
	switch b.running.Load() {
	case runningLongPolling:
		return errors.New("telego: long polling already running")
	case runningWebhook:
		return errors.New("telego: webhook already running")
	default:
		return errors.New("telego: unknown running state")
	}
}

// BotOption represents an option that can be applied to [Bot]
type BotOption func(bot *Bot) error

// NewBot creates new bots with given options (order is important).
// If no options are specified, default values are used.
// Note: Default logger (that logs only errors if not configured) will hide your bot token, but it still may log
// sensitive information, it's only safe to use default logger in testing environment.
func NewBot(token string, options ...BotOption) (*Bot, error) {
	if token == "" {
		return nil, ErrEmptyToken
	}
	if !validateToken(token) {
		return nil, ErrInvalidToken
	}

	b := &Bot{
		token:       token,
		apiURL:      defaultBotAPIServer,
		log:         newDefaultLogger(token),
		api:         ta.FastHTTPCaller{Client: &fasthttp.Client{}},
		constructor: ta.DefaultConstructor{},
	}

	for _, option := range options {
		if err := option(b); err != nil {
			return nil, fmt.Errorf("telego: bot options: %w", err)
		}
	}

	return b, nil
}

// Token returns bot token
func (b *Bot) Token() string {
	return b.token
}

// SecretToken returns a secret token that is HEX encoded SHA-256 hash of your bot's token (this is useful for
// webhooks as using bot's token directly is a security risk and will not work as it contains forbidden symbols)
func (b *Bot) SecretToken() string {
	hash := sha256.Sum256([]byte(b.token))
	return hex.EncodeToString(hash[:])
}

// Logger returns bot logger
func (b *Bot) Logger() Logger {
	return b.log
}

// updateMe updates bot ID and username
func (b *Bot) updateMe() {
	me, err := b.GetMe(context.Background())
	if err != nil {
		b.log.Errorf("Error on get me: %s", err)
	} else {
		b.myID = me.ID
		b.myUsername = me.Username
	}
}

// ID returns bot ID by calling [Bot.GetMe] method once, if error occurs ID will be 0
func (b *Bot) ID() int64 {
	b.myOnce.Do(b.updateMe)
	return b.myID
}

// Username returns bot username by calling [Bot.GetMe] method once, if error occurs username will be empty
func (b *Bot) Username() string {
	b.myOnce.Do(b.updateMe)
	return b.myUsername
}

// FileDownloadURL returns URL that can be used to download a file by its file path retrieved from [Bot.GetFile] method
func (b *Bot) FileDownloadURL(filepath string) string {
	if b.useTestServerPath {
		return b.apiURL + "/file/bot" + b.token + "/test/" + filepath
	}
	return b.apiURL + "/file/bot" + b.token + "/" + filepath
}

// performRequest executes and parses response of method
func (b *Bot) performRequest(ctx context.Context, methodName string, parameters any, vs ...any) error {
	response, err := b.constructAndCallRequest(ctx, methodName, parameters)
	if err != nil {
		b.log.Errorf("Execution error %s: %s", methodName, err)
		return fmt.Errorf("internal execution: %w", err)
	}
	b.log.Debugf("API response %s: %s", methodName, response.String())

	if !response.Ok {
		return fmt.Errorf("api: %w", response.Error)
	}

	if response.Result != nil {
		var unmarshalErr error
		for i := range vs {
			unmarshalErr = json.Unmarshal(response.Result, &vs[i])
			if unmarshalErr == nil {
				break
			}
		}

		if unmarshalErr != nil {
			return fmt.Errorf("unmarshal to %s: %w", reflect.TypeOf(vs[len(vs)-1]), unmarshalErr)
		}
	}

	if b.reportWarningAsErrors && response.Error != nil {
		return response.Error
	}

	return nil
}

// constructAndCallRequest creates and executes request with parsing of parameters
func (b *Bot) constructAndCallRequest(ctx context.Context, methodName string, parameters any) (*ta.Response, error) {
	filesParams, hasFiles := filesParameters(parameters)
	var data *ta.RequestData

	debug := &strings.Builder{}

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

	var url string
	if b.useTestServerPath {
		url = b.apiURL + botPathPrefix + b.token + "/test/" + methodName
	} else {
		url = b.apiURL + botPathPrefix + b.token + "/" + methodName
	}

	debugData := strings.TrimSuffix(debug.String(), "\n")
	b.log.Debugf("API call to: %q, with data: %s", url, debugData)

	response, err := b.api.Call(ctx, url, data)
	if err != nil {
		return nil, fmt.Errorf("request call: %w", err)
	}

	return response, nil
}

// filesParameters gets all files from parameters
func filesParameters(parameters any) (files map[string]ta.NamedReader, hasFiles bool) {
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
func parseParameters(v any) (map[string]string, error) {
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
		key, _, _ = strings.Cut(key, ",") // Trim omitempty and omitzero
		if key == "" {
			return nil, fmt.Errorf("%s does not have `json` tag, or it's empty", paramsStructType.String())
		}

		fieldValue := paramsStruct.Field(i)
		value, err := parseField(fieldValue)
		if err != nil {
			return nil, fmt.Errorf("parse of %s: %w", paramsStructType.String(), err)
		}
		if value == "" {
			continue
		}

		params[key] = value
	}

	return params, nil
}

// parseField parses struct field to string value
func parseField(field reflect.Value) (string, error) {
	if field.IsZero() || !field.CanInterface() {
		return "", nil
	}

	var value string
	var rawString bool

	fieldInterface := field.Interface()
	if value, rawString = fieldInterface.(string); !rawString {
		data, err := json.Marshal(fieldInterface)
		if err != nil {
			return "", err
		}

		value = string(data)
	}

	// Trim double quotes for values marshaled into string (like file names)
	if !rawString && len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"' {
		value = value[1 : len(value)-1]
	}

	return value, nil
}

// isNil checks if the value, or it's underlying interface is nil
func isNil(v any) bool {
	if v == nil {
		return true
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Interface, reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return reflect.ValueOf(v).IsNil()
	default:
		return false
	}
}

// logRequestWithFiles logs request with files
func logRequestWithFiles(debug *strings.Builder, parameters map[string]string, files map[string]ta.NamedReader) {
	debugFiles := make([]string, 0, len(files))
	for k, v := range files {
		if isNil(v) {
			continue
		}

		if k == v.Name() {
			debugFiles = append(debugFiles, fmt.Sprintf("%q", k))
		} else {
			debugFiles = append(debugFiles, fmt.Sprintf("%q: %q", k, v.Name()))
		}
	}
	//nolint:errcheck
	debugJSON, _ := json.Marshal(parameters)
	_, _ = fmt.Fprintf(debug, "parameters: %s, files: {%s}", debugJSON, strings.Join(debugFiles, ", "))
}

// ToPtr converts value into a pointer to value
func ToPtr[T any](value T) *T {
	return &value
}
