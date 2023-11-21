//go:build sonic && amd64 && avx512 && !stdjson

package json

import "github.com/bytedance/sonic"

var Marshal = sonic.ConfigStd.Marshal
var MarshalIndent = sonic.ConfigStd.MarshalIndent
var Unmarshal = sonic.ConfigStd.Unmarshal
var NewEncoder = sonic.ConfigStd.NewEncoder
var NewDecoder = sonic.ConfigStd.NewDecoder
var Valid = sonic.ConfigStd.Valid
