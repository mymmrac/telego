package telego

import (
	"bytes"
	stdJson "encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	defaultBotAPIServer = "https://api.telegram.org"

	jsonContentType = "application/json"

	tokenRegexp = `^\d{9}:[\w-]{35}$` //nolint:gosec

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

type logMode string

const (
	debugMode logMode = "DEBUG"
	errorMode logMode = "ERROR"
)

func logStarting(mode logMode) string {
	timeNow := ansiBlue + time.Now().Local().Format(time.UnixDate) + ansiReset
	switch mode {
	case debugMode:
		return fmt.Sprintf("[%s] %sDEBUG%s", timeNow, ansiYellow, ansiReset)
	case errorMode:
		return fmt.Sprintf("[%s] %sERROR%s", timeNow, ansiRed, ansiReset)
	}
	return "LOG"
}

// Bot - Represents telegram bot
type Bot struct {
	token          string
	apiURL         string
	client         *http.Client
	stopChannel    chan struct{}
	updateInterval time.Duration
	debugMode      bool
	printErrors    bool
}

// NewBot - Creates new bot
func NewBot(token string) (*Bot, error) {
	if !validateToken(token) {
		return nil, ErrInvalidToken
	}

	return &Bot{
		token:          token,
		apiURL:         defaultBotAPIServer,
		client:         http.DefaultClient,
		updateInterval: defaultUpdateInterval,
		debugMode:      false,
		printErrors:    true,
	}, nil
}

func (b *Bot) DebugMode(enabled bool) {
	b.debugMode = enabled
}

func (b *Bot) PrintErrors(enabled bool) {
	b.printErrors = enabled
}

func (b *Bot) SetToken(token string) error {
	if !validateToken(token) {
		return ErrInvalidToken
	}
	b.token = token
	return nil
}

func (b *Bot) SetAPIServer(apiURL string) error {
	if apiURL == "" {
		return errors.New("empty bot api server url")
	}
	b.apiURL = apiURL
	return nil
}

func (b *Bot) SetClient(client *http.Client) error {
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

	resp, err := b.client.Post(url, jsonContentType, buffer)
	if err != nil {
		return nil, fmt.Errorf("post request: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	apiResp := &apiResponse{}

	err = json.NewDecoder(resp.Body).Decode(apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	if b.debugMode {
		fmt.Printf("%s API response %s: %s\n", logStarting(debugMode), methodName, apiResp.String())
	}

	return apiResp, nil
}

func (b Bot) apiRequestMultipartFormData(methodName string,
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
		if value == "" {
			continue
		}

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

	resp, err := b.client.Post(url, writer.FormDataContentType(), buffer)
	if err != nil {
		return nil, fmt.Errorf("post request: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	apiResp := &apiResponse{}

	err = json.NewDecoder(resp.Body).Decode(apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	if b.debugMode {
		fmt.Printf("%s API response %s: %s\n", logStarting(debugMode), methodName, apiResp.String())
	}

	return apiResp, nil
}

func (b *Bot) performRequest(methodName string, parameters, v interface{}) error {
	var (
		resp *apiResponse
		err  error
	)

	switch p := parameters.(type) {
	case fileCompatible:
		fileParams := p.fileParameters()
		isDirectFile := false
		for _, file := range fileParams {
			if file != nil {
				isDirectFile = true
				break
			}
		}

		if isDirectFile {
			params, err := toParams(parameters)
			if err != nil {
				return fmt.Errorf("get params: %w", err)
			}

			resp, err = b.apiRequestMultipartFormData(methodName, params, fileParams)
			if err != nil {
				return fmt.Errorf("request multipart form data: %w", err)
			}
		} else {
			resp, err = b.apiRequest(methodName, parameters)
			if err != nil {
				return fmt.Errorf("request: %w", err)
			}
		}
	default:
		resp, err = b.apiRequest(methodName, parameters)
		if err != nil {
			return fmt.Errorf("request: %w", err)
		}
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

func toParams(v interface{}) (map[string]string, error) {
	val := reflect.ValueOf(v).Elem()
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s not a struct", val.Kind())
	}
	typ := val.Type()

	params := make(map[string]string)

	for i := 0; i < val.NumField(); i++ {
		structField := typ.Field(i)
		field := val.Field(i)

		if field.IsZero() || (field.Kind() == reflect.Ptr && field.IsNil()) {
			continue
		}

		if field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface {
			field = field.Elem()
		}

		key := structField.Tag.Get("json")
		if strings.HasSuffix(key, omitEmptySuffix) {
			key = key[:len(key)-len(omitEmptySuffix)]
		}
		value := field.Interface()

		kind := field.Kind()
		if kind == reflect.Struct || kind == reflect.Slice || kind == reflect.Map {
			buf := bytes.Buffer{}

			err := json.NewEncoder(&buf).Encode(value)
			if err != nil {
				return nil, fmt.Errorf("encoding json: %w", err)
			}

			strVal := buf.String()
			strVal = strVal[:len(strVal)-1] // remove "\n" from end of the string
			if strVal == "" {
				continue
			}
			params[key] = strVal

			continue
		}

		strVal := fmt.Sprintf("%v", value)
		if strVal == "" {
			continue
		}
		params[key] = strVal
	}

	return params, nil
}
