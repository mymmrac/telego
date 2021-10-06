package api

import (
	"bytes"
	stdJson "encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func Test_Response_String_and_Error_Error(t *testing.T) {
	tests := []struct {
		name string
		resp Response
		text string
	}{
		{
			name: "success",
			resp: Response{
				Ok:     true,
				Error:  nil,
				Result: stdJson.RawMessage{},
			},
			text: "Ok: true, Err: {<nil>}, Result: ",
		},
		{
			name: "error",
			resp: Response{
				Ok: false,
				Error: &Error{
					Description: "bad request",
					ErrorCode:   400,
					Parameters:  nil,
				},
				Result: nil,
			},
			text: "Ok: false, Err: {400 \"bad request\"}, Result: ",
		},
		{
			name: "error_with_parameters",
			resp: Response{
				Ok: false,
				Error: &Error{
					Description: "bad request",
					ErrorCode:   400,
					Parameters: &ResponseParameters{
						MigrateToChatID: 1,
						RetryAfter:      2,
					},
				},
				Result: nil,
			},
			text: "Ok: false, Err: {400 \"bad request\" migrate to chat id: 1, retry after: 2}, Result: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			text := tt.resp.String()
			assert.Equal(t, tt.text, text)
		})
	}
}

type Server struct {
	t *testing.T
}

func (s *Server) Handle(ctx *fasthttp.RequestCtx) {
	assert.True(s.t, ctx.IsPost())
	assert.Equal(s.t, ContentTypeJSON, string(ctx.Request.Header.ContentType()))

	_, err := ctx.WriteString("{\"ok\": true}")
	assert.NoError(s.t, err)

	ctx.SetStatusCode(http.StatusOK)
}

func newTestClient(t *testing.T) (client *fasthttp.Client, teardown func()) {
	ln := fasthttputil.NewInmemoryListener()

	api := &Server{
		t: t,
	}
	srv := fasthttp.Server{
		Handler: api.Handle,
	}

	go func() {
		if err := srv.Serve(ln); err != nil {
			panic(err)
		}
	}()

	teardown = func() {
		assert.NoError(t, ln.Close())
	}

	client = &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}

	return client, teardown
}

func TestFasthttpAPICaller_Call(t *testing.T) {
	client, teardown := newTestClient(t)
	defer teardown()

	caller := FasthttpAPICaller{Client: client}

	data := &RequestData{
		ContentType: ContentTypeJSON,
		Buffer:      bytes.NewBufferString("test"),
	}

	resp, err := caller.Call("http://localhost", data)
	require.NoError(t, err)
	assert.True(t, resp.Ok)
}

func TestDefaultConstructor_JSONRequest(t *testing.T) {
	tests := []struct {
		name       string
		parameters interface{}
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
				assert.Error(t, err)
				assert.Nil(t, data)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.data, data,
				fmt.Sprintf("Expected: %q, actual: %q", tt.data.Buffer.String(), data.Buffer.String()))
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
				assert.Error(t, err)
				assert.Nil(t, data)
				return
			}
			assert.NoError(t, err)
			assert.Contains(t, data.ContentType, tt.contentType)

			for _, expectedData := range tt.data {
				assert.Contains(t, data.Buffer.String(), expectedData)
			}
		})
	}
}
