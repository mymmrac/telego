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
	"strconv"
	"strings"
)

const (
	defaultAPIURL = "https://api.telegram.org"

	jsonContentType = "application/json"

	tokenRegexp = `^\d{9}:[\w-]{35}$` //nolint:gosec
)

func validateToken(token string) bool {
	reg := regexp.MustCompile(tokenRegexp)
	return reg.MatchString(token)
}

var (
	// ErrInvalidToken - Bot token is invalid according to token regexp
	ErrInvalidToken = errors.New("invalid token")

	// ErrNilHTTPClient - Provided nil HTTP client
	ErrNilHTTPClient = errors.New("nil http client")
)

// Bot - Represents telegram bot
type Bot struct {
	token  string
	apiURL string
	client *http.Client
}

func botCreator(token, apiURL string, client *http.Client) (*Bot, error) {
	if !validateToken(token) {
		return nil, ErrInvalidToken
	}
	if client == nil {
		return nil, ErrNilHTTPClient
	}
	return &Bot{
		token:  token,
		apiURL: apiURL,
		client: client,
	}, nil
}

// NewBot - Creates new bot
func NewBot(token string) (*Bot, error) {
	return botCreator(token, defaultAPIURL, http.DefaultClient)
}

// NewBotWithAPI - Creates new bot with API URL
func NewBotWithAPI(token, apiURL string) (*Bot, error) {
	return botCreator(token, apiURL, http.DefaultClient)
}

// NewBotWithClient - Creates new bot with HTTP client
func NewBotWithClient(token string, client *http.Client) (*Bot, error) {
	return botCreator(token, defaultAPIURL, client)
}

// NewBotWithAPIAndClient - Creates new bot with API URL and HTTP client
func NewBotWithAPIAndClient(token, apiURL string, client *http.Client) (*Bot, error) {
	return botCreator(token, apiURL, client)
}

type apiResponse struct {
	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result,omitempty"`
	APIError
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
	return fmt.Sprintf("%d %s", a.ErrorCode, a.Description)
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

	return apiResp, nil
}

func (b Bot) apiRequestMultipartFormData(methodName string,
	parameters map[string]string, fileParameters map[string]*os.File) (*apiResponse, error) {
	url := b.apiURL + "/bot" + b.token + "/" + methodName

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	for filed, file := range fileParameters {
		if file == nil {
			continue
		}

		wr, err := writer.CreateFormFile(filed, file.Name())
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

	return apiResp, nil
}

func (b *Bot) performRequest(methodName string, parameters, v interface{}) error {
	var (
		resp *apiResponse
		err  error
	)

	switch p := parameters.(type) {
	case fileCompatible:
		if p.isDirectFile() {
			params := toParams(parameters)
			fileParams := p.fileParameters()

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

func toParams(v interface{}) map[string]string {
	buf := bytes.Buffer{}
	_ = json.NewEncoder(&buf).Encode(v)

	var m map[string]interface{}
	_ = json.NewDecoder(strings.NewReader(buf.String())).Decode(&m)

	params := make(map[string]string)
	extractParams(m, "", params)

	return params
}

func extractParams(v1 interface{}, prefix string, params map[string]string) {
	switch v2 := v1.(type) {
	case map[string]interface{}:
		for key, v3 := range v2 {
			if prefix == "" {
				extractParams(v3, key, params)
			} else {
				extractParams(v3, prefix+"."+key, params)
			}
		}
	case []interface{}:
		for i, v3 := range v2 {
			extractParams(v3, prefix+"."+strconv.Itoa(i), params)
		}
	default:
		params[prefix] = fmt.Sprintf("%v", v2)
	}
}
