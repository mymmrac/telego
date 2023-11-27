package telegoutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net/url"
	"strings"

	"github.com/mymmrac/telego"
	ta "github.com/mymmrac/telego/telegoapi"
)

// namedReaderImpl represents simplest implementation of telegoapi.NamedReader
type namedReaderImpl struct {
	reader io.Reader
	name   string
}

func (r namedReaderImpl) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func (r namedReaderImpl) Name() string {
	return r.name
}

// NameReader "names" io.Reader and returns valid telegoapi.NamedReader
func NameReader(reader io.Reader, name string) ta.NamedReader {
	return namedReaderImpl{
		reader: reader,
		name:   name,
	}
}

// UpdateProcessor allows you to process updates and still use updates chan.
// New updates chan will be closed when the original chan is closed.
// Warning: Deep copy of update is passed, [telego.Update.Clone] method can panic, please read its comment.
func UpdateProcessor(updates <-chan telego.Update, buffer uint, processor func(update telego.Update) telego.Update,
) <-chan telego.Update {
	processedUpdates := make(chan telego.Update, buffer)

	go func() {
		defer close(processedUpdates)
		for update := range updates {
			processedUpdates <- processor(update.Clone())
		}
	}()

	return processedUpdates
}

// WebAppSecret represents secret used to hash web app data
const WebAppSecret = "WebAppData"

// Web app data query names
const (
	WebAppQueryID      = "query_id"
	WebAppUser         = "user"
	WebAppReceiver     = "receiver"
	WebAppChat         = "chat"
	WebAppStartParam   = "start_param"
	WebAppCanSendAfter = "can_send_after"
	WebAppAuthDate     = "auth_date"
	WebAppHash         = "hash"
)

// ValidateWebAppData validates the integrity of value provided by `window.Telegram.WebApp.initData` from web app and
// returns url.Values containing all fields that were provided
// More info: https://core.telegram.org/bots/webapps#validating-data-received-via-the-mini-app
func ValidateWebAppData(token string, data string) (url.Values, error) {
	appData, err := url.ParseQuery(data)
	if err != nil {
		return nil, errors.New("telego: parse query: bad data")
	}

	hash := appData.Get(WebAppHash)
	if hash == "" {
		return nil, errors.New("telego: no hash found")
	}

	appData.Del(WebAppHash)

	// Can't return error because [url.Values.Encode] method always inescapable
	//nolint:errcheck
	appDataToCheck, _ := url.QueryUnescape(strings.ReplaceAll(appData.Encode(), "&", "\n"))

	secretKey := hmacHash([]byte(token), []byte(WebAppSecret))
	if hex.EncodeToString(hmacHash([]byte(appDataToCheck), secretKey)) != hash {
		return nil, errors.New("telego: invalid hash")
	}

	appData.Add(WebAppHash, hash)
	return appData, nil
}

// Login widget data query names
const (
	LoginWidgetID        = "id"
	LoginWidgetFirstName = "first_name"
	LoginWidgetLastName  = "last_name"
	LoginWidgetUsername  = "username"
	LoginWidgetPhotoURL  = "photo_url"
	LoginWidgetAuthDate  = "auth_date"
	LoginWidgetHash      = "hash"
)

// ValidateLoginWidgetData validates the integrity of value provided by Telegram Login Widget and
// returns url.Values containing all fields that were provided
// More info: https://core.telegram.org/widgets/login#checking-authorization
func ValidateLoginWidgetData(token string, data string) (url.Values, error) {
	appData, err := url.ParseQuery(data)
	if err != nil {
		return nil, errors.New("telego: parse query: bad data")
	}

	hash := appData.Get(LoginWidgetHash)
	if hash == "" {
		return nil, errors.New("telego: no hash found")
	}

	appData.Del(LoginWidgetHash)

	// Can't return error because [url.Values.Encode] method always inescapable
	//nolint:errcheck
	appDataToCheck, _ := url.QueryUnescape(strings.ReplaceAll(appData.Encode(), "&", "\n"))

	secretKey := sha256.Sum256([]byte(token))
	if hex.EncodeToString(hmacHash([]byte(appDataToCheck), secretKey[:])) != hash {
		return nil, errors.New("telego: invalid hash")
	}

	appData.Add(LoginWidgetHash, hash)
	return appData, nil
}

// hmacHash hashes data with a provided key using HMAC and SHA256
func hmacHash(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	_, _ = h.Write(data)
	return h.Sum(nil)
}
