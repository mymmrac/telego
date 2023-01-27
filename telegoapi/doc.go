/*
Package telegoapi provides API for calling Telegram for Telego.

This API package describes the main part of communication with Telegram Bot API.

The response represents the API response from Telegram with respectful result and error values.

Caller interface represents the general logic of sending requests to API and receiving responses from it. Currently,
Telego provides valyala/fasthttp and net/http implementation, but your own can be defined and specified
via telego.BotOption's.

RequestConstructor interface represents a general way of constructing RequestData used in Caller. Currently, Telego
provides an only default implementation that uses goccy/go-json instead of encoding/json and std mime/multipart
package.

NamedReader interface represents a general way of sending files that are provided to RequestConstructor. As io.Reader
can be provided, any valid reader and name method should return a unique name for every file in one request, otherwise
not all files will be sent properly.
*/
package telegoapi
