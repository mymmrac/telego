package main

import (
	"fmt"
	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	c := telego.IntOrStringChatID{
		StringValue: "@test",
		//IntValue: 4234,
	}
	b, _ := c.MarshalJSON()
	fmt.Println(string(b))
}
