//go:build sonic && !stdjson

package json

import "github.com/bytedance/sonic"

func init() {
	Marshal = sonic.ConfigStd.Marshal
	Unmarshal = sonic.ConfigStd.Unmarshal
}
