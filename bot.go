package telego

import (
	"bytes"
	"encoding/json"
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

	"github.com/sirupsen/logrus"
)

// TODO: Make use of https://github.com/json-iterator/go

const (
	defaultBotAPIServer = "https://api.telegram.org"

	jsonContentType = "application/json"

	tokenRegexp = `^\d{9}:[\w-]{35}$` //nolint:gosec

	attachFile = `attach://`
)

func validateToken(token string) bool {
	reg := regexp.MustCompile(tokenRegexp)
	return reg.MatchString(token)
}

var (
	// ErrInvalidToken - Bot token is invalid according to token regexp
	ErrInvalidToken = errors.New("invalid token")
)

// Bot - Represents telegram bot
type Bot struct {
	token  string
	apiURL string
	client *http.Client
	log    *logrus.Logger
}

// NewBot - Creates new bot
func NewBot(token string) (*Bot, error) {
	if !validateToken(token) {
		return nil, ErrInvalidToken
	}

	log := logrus.StandardLogger()
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = time.RFC1123
	formatter.FullTimestamp = true
	log.SetFormatter(formatter)
	log.SetLevel(logrus.ErrorLevel)

	return &Bot{
		token:  token,
		apiURL: defaultBotAPIServer,
		client: http.DefaultClient,
		log:    log,
	}, nil
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

func (b *Bot) DebugMode(is bool) {
	if is {
		b.log.SetLevel(logrus.DebugLevel)
		return
	}
	b.log.SetLevel(logrus.ErrorLevel)
}

type apiResponse struct {
	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result,omitempty"`
	APIError
}

func (a apiResponse) String() string {
	return fmt.Sprintf("[OK: %t, ERR: %s]: %s", a.Ok, a.APIError.Error(), string(a.Result))
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
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		return nil, fmt.Errorf("post request: %w", err)
	}

	apiResp := &apiResponse{}

	err = json.NewDecoder(resp.Body).Decode(apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	b.log.Debugf("API response %s: %s", methodName, apiResp.String())

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
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		return nil, fmt.Errorf("post request: %w", err)
	}

	apiResp := &apiResponse{}

	err = json.NewDecoder(resp.Body).Decode(apiResp)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	b.log.Debug(apiResp.String())

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

// TODO: Refactor (without json encoding)

func toParams(v interface{}) (map[string]string, error) {
	tmpBuf := bytes.Buffer{}
	err := json.NewEncoder(&tmpBuf).Encode(v)
	if err != nil {
		return nil, fmt.Errorf("encoding json: %w", err)
	}

	var m map[string]interface{}
	err = json.NewDecoder(&tmpBuf).Decode(&m)
	if err != nil {
		return nil, fmt.Errorf("decoding json: %w", err)
	}

	params := make(map[string]string)

	for key, value := range m {
		kind := reflect.ValueOf(value).Kind()
		if kind == reflect.Struct || kind == reflect.Slice || kind == reflect.Map {
			buf := bytes.Buffer{}

			err = json.NewEncoder(&buf).Encode(value)
			if err != nil {
				return nil, fmt.Errorf("encoding json: %w", err)
			}

			strVal := buf.String()
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
