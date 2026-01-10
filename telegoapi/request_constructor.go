//go:generate mockgen -typed -package mock -destination=mock/request_constructor.go github.com/mymmrac/telego/telegoapi RequestConstructor

package telegoapi

import (
	"fmt"
	"io"
	"mime/multipart"
	"reflect"

	"github.com/mymmrac/telego/internal/json"
)

// DefaultConstructor default implementation of [RequestConstructor]
type DefaultConstructor struct{}

// JSONRequest is default implementation
func (d DefaultConstructor) JSONRequest(parameters any) (*RequestData, error) {
	data, err := json.Marshal(parameters)
	if err != nil {
		return nil, fmt.Errorf("encode json: %w", err)
	}

	return &RequestData{
		ContentType: ContentTypeJSON,
		BodyRaw:     data,
	}, nil
}

// MultipartRequest is default implementation
func (d DefaultConstructor) MultipartRequest(
	parameters map[string]string, filesParameters map[string]NamedReader,
) (*RequestData, error) {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	data := &RequestData{
		ContentType: writer.FormDataContentType(),
		BodyStream:  pr,
	}

	go func() {
		defer func() { _ = pw.CloseWithError(writer.Close()) }()

		for field, value := range parameters {
			if err := writer.WriteField(field, value); err != nil {
				_ = pw.CloseWithError(fmt.Errorf("write field: %w", err))
				return
			}
		}

		for field, file := range filesParameters {
			if isNil(file) {
				continue
			}

			wr, err := writer.CreateFormFile(field, file.Name())
			if err != nil {
				_ = pw.CloseWithError(fmt.Errorf("write file header: %w", err))
				return
			}

			if _, err = io.Copy(wr, file); err != nil {
				_ = pw.CloseWithError(fmt.Errorf("write file: %w", err))
				return
			}
		}
	}()

	return data, nil
}

// isNil checks if the value, or it's underlying interface is nil
func isNil(v any) bool {
	if v == nil {
		return true
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Interface, reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return reflect.ValueOf(v).IsNil()
	default:
		return false
	}
}
