package telego

import (
	"context"
	"errors"
	"io"
	"net/http"
	"sync"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// WebhookSecretTokenHeader represents secret token header name, see SetWebhookParams.SecretToken for more details
const WebhookSecretTokenHeader = "X-Telegram-Bot-Api-Secret-Token" //nolint:gosec

// FastHTTPWebhookServer represents fasthttp implementation of WebhookServer.
// The Server and Router are required fields, optional Logger and SecretToken can be provided.
type FastHTTPWebhookServer struct {
	Logger      Logger
	Server      *fasthttp.Server
	Router      *router.Router
	SecretToken string
}

// Start starts server
func (f FastHTTPWebhookServer) Start(address string) error {
	return f.Server.ListenAndServe(address)
}

// Stop stops server
func (f FastHTTPWebhookServer) Stop(ctx context.Context) error {
	return f.Server.ShutdownWithContext(ctx)
}

// RegisterHandler registers new POST handler for the desired path
// Note: If server's handler is not set, it will be set to router's handler
func (f FastHTTPWebhookServer) RegisterHandler(path string, handler func(data []byte) error) error {
	f.Router.POST(path, func(ctx *fasthttp.RequestCtx) {
		if f.SecretToken != "" {
			secretToken := ctx.Request.Header.Peek(WebhookSecretTokenHeader)
			if f.SecretToken != string(secretToken) {
				if f.Logger != nil {
					f.Logger.Errorf("Webhook handler: unauthorized: secret token does not match")
				}

				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				return
			}
		}

		if err := handler(ctx.PostBody()); err != nil {
			if f.Logger != nil {
				f.Logger.Errorf("Webhook handler: %s", err)
			}

			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		ctx.SetStatusCode(fasthttp.StatusOK)
	})

	if f.Server.Handler == nil {
		f.Server.Handler = f.Router.Handler
	}

	return nil
}

// HTTPWebhookServer represents http implementation of WebhookServer.
// The Server and ServeMux are required fields, optional Logger and SecretToken can be provided.
type HTTPWebhookServer struct {
	Logger      Logger
	Server      *http.Server
	ServeMux    *http.ServeMux
	SecretToken string
}

// Start starts server
func (h HTTPWebhookServer) Start(address string) error {
	if h.Server.Addr == "" {
		h.Server.Addr = address
	}
	if err := h.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

// Stop stops server
func (h HTTPWebhookServer) Stop(ctx context.Context) error {
	return h.Server.Shutdown(ctx)
}

// RegisterHandler registers new POST handler for the desired path
// Note: If server's handler is not set, it will be set to serve mux handler
func (h HTTPWebhookServer) RegisterHandler(path string, handler func(data []byte) error) error {
	h.ServeMux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		if !h.validateRequest(writer, request) {
			return
		}

		data, err := io.ReadAll(request.Body)
		if err != nil {
			if h.Logger != nil {
				h.Logger.Errorf("Webhook handler: read body: %s", err)
			}

			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = request.Body.Close(); err != nil {
			if h.Logger != nil {
				h.Logger.Errorf("Webhook handler: close body: %s", err)
			}
		}

		if err = handler(data); err != nil {
			if h.Logger != nil {
				h.Logger.Errorf("Webhook handler: %s", err)
			}

			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusOK)
	})

	if h.Server.Handler == nil {
		h.Server.Handler = h.ServeMux
	}

	return nil
}

func (h HTTPWebhookServer) validateRequest(writer http.ResponseWriter, request *http.Request) bool {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}

	if h.SecretToken != "" {
		secretToken := request.Header.Get(WebhookSecretTokenHeader)
		if h.SecretToken != secretToken {
			if h.Logger != nil {
				h.Logger.Errorf("Webhook handler: unauthorized: secret token does not match")
			}

			writer.WriteHeader(http.StatusUnauthorized)
			return false
		}
	}

	return true
}

// MultiBotWebhookServer represents multi bot implementation of WebhookServer,
// stable for running multiple bots from single server
type MultiBotWebhookServer struct {
	Server WebhookServer

	startOnce sync.Once
	stopOnce  sync.Once
}

// Start starts server only once
func (m *MultiBotWebhookServer) Start(address string) error {
	var err error
	m.startOnce.Do(func() {
		err = m.Server.Start(address)
	})
	return err
}

// Stop stops server only once
func (m *MultiBotWebhookServer) Stop(ctx context.Context) error {
	var err error
	m.stopOnce.Do(func() {
		err = m.Server.Stop(ctx)
	})
	return err
}

// RegisterHandler registers new handler for the desired path
func (m *MultiBotWebhookServer) RegisterHandler(path string, handler func(data []byte) error) error {
	return m.Server.RegisterHandler(path, handler)
}
