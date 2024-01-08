package json

import (
	"encoding/json"

	"github.com/valyala/fastjson"
)

type RawMessage = json.RawMessage

var ParserPoll = &fastjson.ParserPool{}

func init() {
	for i := 0; i < 8; i++ {
		ParserPoll.Put(&fastjson.Parser{})
	}
}
