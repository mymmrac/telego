package json

import (
	"encoding/json"

	"github.com/valyala/fastjson"
)

type RawMessage = json.RawMessage

var ParserPoll = &fastjson.ParserPool{}
