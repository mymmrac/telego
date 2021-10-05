package api

import (
	"bytes"
	stdJson "encoding/json"
	"net"
	"net/http"
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
			name: "error with parameters",
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
