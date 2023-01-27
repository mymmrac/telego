package telegoapi

import (
	"bytes"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestFasthttpAPICaller_Call(t *testing.T) {
	ln := fasthttputil.NewInmemoryListener()

	api := &server{
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

	teardown := func() {
		assert.NoError(t, ln.Close())
	}

	client := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}

	defer teardown()

	caller := FasthttpAPICaller{Client: client}

	data := &RequestData{
		ContentType: ContentTypeJSON,
		Buffer:      bytes.NewBufferString("test"),
	}

	t.Run("success", func(t *testing.T) {
		resp, err := caller.Call("http://localhost", data)
		require.NoError(t, err)
		assert.True(t, resp.Ok)
	})

	t.Run("error_fasthttp_do_request", func(t *testing.T) {
		resp, err := caller.Call("abc", data)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_500", func(t *testing.T) {
		resp, err := caller.Call("http://localhost/500", data)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error_json", func(t *testing.T) {
		resp, err := caller.Call("http://localhost/json_err", data)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
