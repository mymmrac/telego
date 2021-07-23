package main

import (
	"fmt"
	"os"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	testToken := os.Getenv("TEST_TOKEN")

	bot, err := telego.NewBot(testToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(botUser.Username)

	//updParams := &telego.GetUpdatesParams{
	//	Offset:  0,
	//	Limit:   0,
	//	Timeout: 0,
	//	//AllowedUpdates: []string{"chat_member"},
	//}
	//upd, err := bot.GetUpdates(updParams)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, u := range upd {
	//	fmt.Println(u)
	//}

	//file, err := os.Open("doc.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//dp := &telego.SendDocumentParams{
	//	ChatID:   telego.ChatID{ID: 331849104},
	//	Document: telego.InputFile{File: file},
	//	Caption:  "Hello world",
	//}
	//msg, err := bot.SendDocument(dp)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(msg.Document)
}
