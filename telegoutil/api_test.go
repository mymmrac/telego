package telegoutil

import (
	"io"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

func TestNameReader(t *testing.T) {
	buf := strings.NewReader(text1)

	nameReader := NameReader(buf, text2)

	data, err := io.ReadAll(nameReader)
	require.NoError(t, err)

	assert.Equal(t, text1, string(data))
	assert.Equal(t, text2, nameReader.Name())
}

func TestNameBytes(t *testing.T) {
	nameBytes := NameBytes([]byte(text1), text2)

	data, err := io.ReadAll(nameBytes)
	require.NoError(t, err)

	assert.Equal(t, text1, string(data))
	assert.Equal(t, text2, nameBytes.Name())
}

func TestUpdateProcessor(t *testing.T) {
	updates := make(chan telego.Update)

	wg := sync.WaitGroup{}

	processedUpdates := UpdateProcessor(updates, 10, func(update telego.Update) telego.Update {
		wg.Done()
		update.UpdateID *= 10
		return update
	})

	const updateCount = 2
	wg.Add(updateCount)

	updates <- telego.Update{UpdateID: 1}
	updates <- telego.Update{UpdateID: 2}

	wg.Wait()

	count := 0
	for update := range processedUpdates {
		count++
		assert.Equal(t, 0, update.UpdateID%10)

		if count == updateCount {
			close(updates)
		}
	}
}

func TestTypesConstants(t *testing.T) {
	tests := [][]string{
		{
			WebAppSecret,
		},
		{
			WebAppQueryID, WebAppUser, WebAppReceiver, WebAppChat, WebAppStartParam, WebAppCanSendAfter, WebAppAuthDate,
			WebAppHash,
		},
		{
			LoginWidgetID, LoginWidgetFirstName, LoginWidgetLastName, LoginWidgetUsername, LoginWidgetPhotoURL,
			LoginWidgetAuthDate, LoginWidgetHash,
		},
	}

	for _, tt := range tests {
		assert.NotEmpty(t, tt)
		for _, ct := range tt {
			assert.NotEmpty(t, ct)
		}
	}
}

func TestValidateWebAppData(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		data    string
		values  url.Values
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "success",
			token: "token",
			data:  "hash=2960d6823fe22f39b2b0b547baf7fa95a053411d4685a57a16b38cc57346a4e6&ok=true",
			values: map[string][]string{
				"hash": {"2960d6823fe22f39b2b0b547baf7fa95a053411d4685a57a16b38cc57346a4e6"},
				"ok":   {"true"},
			},
			wantErr: assert.NoError,
		},
		{
			name:    "error_invalid_query",
			token:   "",
			data:    "%_",
			values:  nil,
			wantErr: assert.Error,
		},
		{
			name:    "error_empty_hash",
			token:   "",
			data:    "hash=",
			values:  nil,
			wantErr: assert.Error,
		},
		{
			name:    "error_no_hash",
			token:   "",
			data:    "abc=%a2",
			values:  nil,
			wantErr: assert.Error,
		},
		{
			name:    "error_invalid_hash",
			token:   "",
			data:    "hash=test",
			values:  nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, err := ValidateWebAppData(tt.token, tt.data)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.values, values)
		})
	}
}

func TestValidateLoginWidgetData(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		data    string
		values  url.Values
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "success",
			token: "token",
			data:  "hash=22414c1fdbad44e8156276e1c2a99b332daf5218a65731c745108e34d82c2348&ok=true",
			values: map[string][]string{
				"hash": {"22414c1fdbad44e8156276e1c2a99b332daf5218a65731c745108e34d82c2348"},
				"ok":   {"true"},
			},
			wantErr: assert.NoError,
		},
		{
			name:    "error_invalid_query",
			token:   "",
			data:    "%_",
			values:  nil,
			wantErr: assert.Error,
		},
		{
			name:    "error_empty_hash",
			token:   "",
			data:    "hash=",
			values:  nil,
			wantErr: assert.Error,
		},
		{
			name:    "error_no_hash",
			token:   "",
			data:    "abc=%a2",
			values:  nil,
			wantErr: assert.Error,
		},
		{
			name:    "error_invalid_hash",
			token:   "",
			data:    "hash=test",
			values:  nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, err := ValidateLoginWidgetData(tt.token, tt.data)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.values, values)
		})
	}
}
