//go:build stdjson && !sonic

package json

import "encoding/json"

var Marshal = json.Marshal
var Unmarshal = json.Unmarshal
var NewEncoder = json.NewEncoder
var NewDecoder = json.NewDecoder
