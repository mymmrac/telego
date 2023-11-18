//go:build stdjson && !(sonic && amd64 && avx512)

package json

import "encoding/json"

var Marshal = json.Marshal
var MarshalIndent = json.MarshalIndent
var Unmarshal = json.Unmarshal
var NewEncoder = json.NewEncoder
var NewDecoder = json.NewDecoder
var Valid = json.Valid
