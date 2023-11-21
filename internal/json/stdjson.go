//go:build stdjson && !(sonic && avx && (linux || windows || darwin) && amd64)

package json

import "encoding/json"

var Marshal = json.Marshal
var Unmarshal = json.Unmarshal
var NewEncoder = json.NewEncoder
var NewDecoder = json.NewDecoder
