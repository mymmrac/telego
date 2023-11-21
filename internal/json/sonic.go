//go:build sonic && avx && (linux || windows || darwin) && amd64

package json

import "github.com/bytedance/sonic"

var Marshal = sonic.ConfigStd.Marshal
var Unmarshal = sonic.ConfigStd.Unmarshal
var NewEncoder = sonic.ConfigStd.NewEncoder
var NewDecoder = sonic.ConfigStd.NewDecoder
