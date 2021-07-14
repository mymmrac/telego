package telego

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
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
	Ok     bool `json:"ok"`
	Error  *apiError
	Result json.RawMessage `json:"result,omitempty"`
}

type apiError struct {
	Description string              `json:"description,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

func (e apiError) Error() string {
	if e.Parameters != nil {
		return fmt.Sprintf("%d %s migrate to chat id: %d, retry after: %d",
			e.ErrorCode, e.Description, e.Parameters.MigrateToChatID, e.Parameters.RetryAfter)
	}
	return fmt.Sprintf("%d %s", e.ErrorCode, e.Description)
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

func (b Bot) performRequest(methodName string, parameters, v interface{}) error {
	resp, err := b.apiRequest(methodName, parameters)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}

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
