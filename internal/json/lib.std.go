//go:build stdjson && !sonic

package json

import "encoding/json"

func init() {
	Marshal = json.Marshal
	Unmarshal = json.Unmarshal
}
