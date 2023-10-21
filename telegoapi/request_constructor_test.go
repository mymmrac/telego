package telegoapi

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultConstructor_JSONRequest(t *testing.T) {
	tests := []struct {
		name       string
		parameters any
		data       *RequestData
		isError    bool
	}{
		{
			name: "success",
			parameters: struct {
				N int    `json:"n"`
				S string `json:"s"`
				E int    `json:"e,omitempty"`
			}{
				N: 1,
				S: "test",
			},
			data: &RequestData{
				ContentType: ContentTypeJSON,
				Buffer:      bytes.NewBufferString(`{"n":1,"s":"test"}` + "\n"),
			},
			isError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DefaultConstructor{}
			data, err := d.JSONRequest(tt.parameters)
			if tt.isError {
				require.Error(t, err)
				assert.Nil(t, data)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.data, data,
				fmt.Sprintf("Expected: %q, actual: %q", tt.data.Buffer.String(), data.Buffer.String()))
		})
	}
}

func TestDefaultConstructor_MultipartRequest(t *testing.T) {
	tests := []struct {
		name            string
		parameters      map[string]string
		filesParameters map[string]NamedReader
		contentType     string
		data            []string
		isError         bool
	}{
		{
			name: "success",
			parameters: map[string]string{
				"testParam": "1",
			},
			filesParameters: map[string]NamedReader{
				"testFile": newTestFile("Hello World", "testF"),
				"nilFile":  nil,
			},
			contentType: "multipart/form-data; boundary=",
			data: []string{
				"Content-Disposition: form-data; name=\"testFile\"; filename=\"testF\"\r\n" +
					"Content-Type: application/octet-stream\r\n\r\nHello World\r\n",
				"Content-Disposition: form-data; name=\"testParam\"\r\n\r\n1\r\n",
			},
			isError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DefaultConstructor{}
			data, err := d.MultipartRequest(tt.parameters, tt.filesParameters)
			if tt.isError {
				require.Error(t, err)
				assert.Nil(t, data)
				return
			}
			require.NoError(t, err)
			assert.Contains(t, data.ContentType, tt.contentType)

			for _, expectedData := range tt.data {
				assert.Contains(t, data.Buffer.String(), expectedData)
			}
		})
	}
}

type testFile struct {
	data     io.Reader
	fileName string
}

func newTestFile(data, name string) testFile {
	return testFile{
		data:     strings.NewReader(data),
		fileName: name,
	}
}

func (t testFile) Read(p []byte) (n int, err error) {
	return t.data.Read(p)
}

func (t testFile) Name() string {
	return t.fileName
}

func Test_isNil(t *testing.T) {
	var nr NamedReader
	var ns any
	var s []int
	ns = s

	tests := []struct {
		name  string
		value any
		isNil bool
	}{
		{
			name:  "nil",
			value: nil,
			isNil: true,
		},
		{
			name:  "not_nil",
			value: true,
			isNil: false,
		},
		{
			name:  "nil_interface",
			value: nr,
			isNil: true,
		},
		{
			name:  "nil_value_interface",
			value: ns,
			isNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isNil, isNil(tt.value))
		})
	}
}
