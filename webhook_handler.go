package telego

import (
	"io"
	"net/http"

	"github.com/valyala/fasthttp"
)

// WebhookSecretTokenHeader represents secret token header name, see [SetWebhookParams.SecretToken] for more details
const WebhookSecretTokenHeader = "X-Telegram-Bot-Api-Secret-Token" //nolint:gosec

// WebhookFastHTTP registers new POST handler for the desired path with optional secret token, replacing
// original fasthttp handler for the server
func WebhookFastHTTP(server *fasthttp.Server, path string, secretToken ...string) func(handler WebhookHandler) error {
	if path == "" {
		path = "/"
	}
	return func(handler WebhookHandler) error {
		server.Handler = func(fCtx *fasthttp.RequestCtx) {
			if string(fCtx.Path()) != path {
				fCtx.SetStatusCode(fasthttp.StatusNotFound)
				return
			}

			if !fCtx.Request.Header.IsPost() {
				fCtx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
				return
			}

			if len(secretToken) > 0 && secretToken[0] != string(fCtx.Request.Header.Peek(WebhookSecretTokenHeader)) {
				fCtx.SetStatusCode(fasthttp.StatusUnauthorized)
				return
			}

			if err := handler(fCtx, fCtx.PostBody()); err != nil {
				fCtx.SetStatusCode(fasthttp.StatusInternalServerError)
				return
			}

			fCtx.SetStatusCode(fasthttp.StatusOK)
		}
		return nil
	}
}

// WebhookHTTPServer registers new POST handler for the desired path with optional secret token, replacing
// original http handler for the server
func WebhookHTTPServer(server *http.Server, path string, secretToken ...string) func(handler WebhookHandler) error {
	if path == "" {
		path = "/"
	}
	return func(handler WebhookHandler) error {
		server.Handler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Path != path {
				writer.WriteHeader(http.StatusNotFound)
				return
			}

			if request.Method != http.MethodPost {
				writer.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			if len(secretToken) > 0 && secretToken[0] != request.Header.Get(WebhookSecretTokenHeader) {
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}

			data, err := io.ReadAll(request.Body)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err = request.Body.Close(); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err = handler(request.Context(), data); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			writer.WriteHeader(http.StatusOK)
		})
		return nil
	}
}

// WebhookHTTPServeMux registers new handler for the desired pattern with optional secret token
func WebhookHTTPServeMux(mux *http.ServeMux, pattern string, secretToken ...string) func(handler WebhookHandler) error {
	if pattern == "" {
		pattern = "POST /"
	}
	return func(handler WebhookHandler) error {
		mux.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
			if len(secretToken) > 0 && secretToken[0] != request.Header.Get(WebhookSecretTokenHeader) {
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}

			data, err := io.ReadAll(request.Body)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err = request.Body.Close(); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err = handler(request.Context(), data); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			writer.WriteHeader(http.StatusOK)
		})
		return nil
	}
}
