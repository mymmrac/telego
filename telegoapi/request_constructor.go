//go:generate mockgen -typed -package mock -destination=mock/request_constructor.go github.com/mymmrac/telego/telegoapi RequestConstructor

package telegoapi

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"reflect"

	"github.com/goccy/go-json"
)

// DefaultConstructor default implementation of RequestConstructor
type DefaultConstructor struct{}

// JSONRequest is default implementation
func (d DefaultConstructor) JSONRequest(parameters any) (*RequestData, error) {
	data := &RequestData{
		ContentType: ContentTypeJSON,
		Buffer:      &bytes.Buffer{},
	}

	err := json.NewEncoder(data.Buffer).Encode(parameters)
	if err != nil {
		return nil, fmt.Errorf("encode json: %w", err)
	}

	return data, nil
}

// MultipartRequest is default implementation
func (d DefaultConstructor) MultipartRequest(parameters map[string]string, filesParameters map[string]NamedReader) (
	*RequestData, error,
) {
	data := &RequestData{
		Buffer: &bytes.Buffer{},
	}
	writer := multipart.NewWriter(data.Buffer)

	for field, file := range filesParameters {
		if isNil(file) {
			continue
		}

		wr, err := writer.CreateFormFile(field, file.Name())
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(wr, file)
		if err != nil {
			return nil, err
		}
	}

	for field, value := range parameters {
		if err := writer.WriteField(field, value); err != nil {
			return nil, err
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("closing writer: %w", err)
	}

	data.ContentType = writer.FormDataContentType()
	return data, nil
}

func isNil(i any) bool {
	if i == nil {
		return true
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	default:
		return false
	}
}
