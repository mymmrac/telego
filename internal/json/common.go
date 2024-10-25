package json

import (
	"encoding/json"

	"github.com/valyala/fastjson"
)

type RawMessage = json.RawMessage

var ParserPoll = &fastjson.ParserPool{}

var (
	Marshal   func(v any) ([]byte, error)
	Unmarshal func(data []byte, v any) error
)
