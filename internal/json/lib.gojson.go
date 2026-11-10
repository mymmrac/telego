//go:build !sonic && !stdjson

package json

import "github.com/goccy/go-json"

func init() {
	Marshal = json.Marshal
	Unmarshal = json.Unmarshal
}
