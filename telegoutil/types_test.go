package telegoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testNamedReade struct{}

func (t testNamedReade) Read(_ []byte) (n int, err error) {
	panic("implement me")
}

func (t testNamedReade) Name() string {
	return "test"
}

func TestFile(t *testing.T) {
	n := testNamedReade{}
	f := File(n)
	assert.Equal(t, n, f.File)
}

func TestFileByID(t *testing.T) {
	fileID := "file"
	f := FileByID(fileID)
	assert.Equal(t, fileID, f.FileID)
}

func TestFileByURL(t *testing.T) {
	url := "url"
	f := FileByURL(url)
	assert.Equal(t, url, f.URL)
}

func TestID(t *testing.T) {
	var intID int64 = 123
	chatID := ID(intID)
	assert.Equal(t, intID, chatID.ID)
}

func TestUsername(t *testing.T) {
	var username = "username"
	chatID := Username(username)
	assert.Equal(t, username, chatID.Username)
}
