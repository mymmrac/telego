//go:build !sonic && !stdjson

package json

import (
	"github.com/goccy/go-json"
)

var Marshal = json.Marshal
var Unmarshal = json.Unmarshal
var NewEncoder = json.NewEncoder
var NewDecoder = json.NewDecoder
