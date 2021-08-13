package telego

import (
	"bytes"
	stdJson "encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	defaultBotAPIServer = "https://api.telegram.org"

	jsonContentType = "application/json"

	tokenRegexp = `^\d{9,10}:[\w-]{35}$` //nolint:gosec

	attachFile = `attach://`

	omitEmptySuffix = ",omitempty"

	ansiReset  = "\u001B[0m"
	ansiRed    = "\u001B[31m"
	ansiYellow = "\u001B[33m"
	ansiBlue   = "\u001B[34m"
)

var (
	// ErrInvalidToken - Bot token is invalid according to token regexp
	ErrInvalidToken = errors.New("invalid token")
)

func validateToken(token string) bool {
	reg := regexp.MustCompile(tokenRegexp)
	return reg.MatchString(token)
}

// Bot - Represents telegram bot
type Bot struct {
	token          string
	apiURL         string
	client         *fasthttp.Client
	stopChannel    chan struct{}
	updateInterval time.Duration
	webhookHandler fasthttp.RequestHandler
	log            Logger
}

// NewBot - Creates new bot
func NewBot(token string) (*Bot, error) {
	if !validateToken(token) {
		return nil, ErrInvalidToken
	}

	return &Bot{
		token:          token,
		apiURL:         defaultBotAPIServer,
		client:         &fasthttp.Client{},
		updateInterval: defaultUpdateInterval,
		log:            newLogger(),
	}, nil
}

// DefaultLogger - Setup default logger. Redefines existing logger
func (b *Bot) DefaultLogger(debugMode, printErrors bool) {
	log := &logger{
		Out:         os.Stderr,
		DebugMode:   debugMode,
		PrintErrors: printErrors,
	}
	b.log = log
}

// SetLogger - Set logger
func (b *Bot) SetLogger(log Logger) {
	b.log = log
}

// SetToken - Sets bot token
func (b *Bot) SetToken(token string) error {
	if !validateToken(token) {
		return ErrInvalidToken
	}
	b.token = token
	return nil
}

// Token - Returns bot token
func (b *Bot) Token() string {
	return b.token
}

// SetAPIServer - Sets bot API server
func (b *Bot) SetAPIServer(apiURL string) error {
	if apiURL == "" {
		return errors.New("empty bot api server url")
	}
	b.apiURL = apiURL
	return nil
}

// SetClient - Sets fasthttp client to use
func (b *Bot) SetClient(client *fasthttp.Client) error {
	if client == nil {
		return errors.New("nil http client")
	}
	b.client = client
	return nil
}

type apiResponse struct {
	Ok     bool               `json:"ok"`
	Result stdJson.RawMessage `json:"result,omitempty"`
	APIError
}

func (a apiResponse) String() string {
	return fmt.Sprintf("Ok: %t, Err: {%v}, Result: %s", a.Ok, a.APIError, a.Result)
}

// APIError - Represents error from telegram API
type APIError struct {
	Description string              `json:"description,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

func (a APIError) Error() string {
	if a.Parameters != nil {
		return fmt.Sprintf("%d %s migrate to chat id: %d, retry after: %d",
			a.ErrorCode, a.Description, a.Parameters.MigrateToChatID, a.Parameters.RetryAfter)
	}
	return fmt.Sprintf("%d %q", a.ErrorCode, a.Description)
}

func (b *Bot) apiRequest(methodName string, parameters interface{}) (*apiResponse, error) {
	url := b.apiURL + "/bot" + b.token + "/" + methodName

	buffer := &bytes.Buffer{}
	err := json.NewEncoder(buffer).Encode(parameters)
	if err != nil {
		return nil, fmt.Errorf("encode json: %w", err)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetContentType(jsonContentType)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBodyRaw(buffer.Bytes())

	resp := fasthttp.AcquireResponse()
	err = b.client.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	if statusCode := resp.StatusCode(); statusCode >= fasthttp.StatusInternalServerError {
		b.log.Errorf("Internal server error, status code: %d", statusCode)
		return nil, fmt.Errorf("internal server error: %d", statusCode)
	}

	apiResp := &apiResponse{}
	err = json.Unmarshal(resp.Body(), apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	b.log.Debugf("API response %s: %s", methodName, apiResp.String())

	return apiResp, nil
}

func (b *Bot) apiRequestMultipartFormData(methodName string,
	parameters map[string]string, fileParameters map[string]*os.File) (*apiResponse, error) {
	url := b.apiURL + "/bot" + b.token + "/" + methodName

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	for field, file := range fileParameters {
		if file == nil {
			continue
		}

		wr, err := writer.CreateFormFile(field, file.Name())
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(wr, file)
		if err != nil {
			return nil, err
		}
	}

	for field, value := range parameters {
		wr, err := writer.CreateFormField(field)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(wr, strings.NewReader(value))
		if err != nil {
			return nil, err
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetContentType(writer.FormDataContentType())
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBodyRaw(buffer.Bytes())

	resp := fasthttp.AcquireResponse()
	err = b.client.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("multipart request: %w", err)
	}

	if statusCode := resp.StatusCode(); statusCode >= fasthttp.StatusInternalServerError {
		b.log.Errorf("Internal server error, status code: %d", statusCode)
		return nil, fmt.Errorf("internal server error: %d", statusCode)
	}

	apiResp := &apiResponse{}
	err = json.Unmarshal(resp.Body(), apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	b.log.Debugf("API response %s: %s", methodName, apiResp.String())

	return apiResp, nil
}

func (b *Bot) performRequest(methodName string, parameters, v interface{}) error {
	resp, err := b.executeRequest(methodName, parameters)
	if err != nil {
		return fmt.Errorf("execute: %w", err)
	}

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

func (b *Bot) executeRequest(methodName string, parameters interface{}) (*apiResponse, error) {
	isDirectFile := false
	var fileParams map[string]*os.File

	p, ok := parameters.(fileCompatible)
	if ok {
		fileParams = p.fileParameters()
		for _, file := range fileParams {
			if file != nil {
				isDirectFile = true
				break
			}
		}
	}

	if isDirectFile {
		params, err := toParams(parameters)
		if err != nil {
			return nil, fmt.Errorf("get params: %w", err)
		}

		resp, err := b.apiRequestMultipartFormData(methodName, params, fileParams)
		if err != nil {
			return nil, fmt.Errorf("request multipart form data: %w", err)
		}
		return resp, nil
	}

	resp, err := b.apiRequest(methodName, parameters)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	return resp, nil
}

func toParams(v interface{}) (map[string]string, error) {
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
