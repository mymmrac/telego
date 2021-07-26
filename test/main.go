package main

import (
	"fmt"
	"os"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	testToken := os.Getenv("TEST_TOKEN")

	//myID := telego.ChatID{ID: 331849104}
	//groupID := telego.ChatID{ID: -1001516926498}

	bot, err := telego.NewBot(testToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	bot.DebugMode(true)

	_, err = bot.GetMe()
	if err != nil {
		fmt.Println(err)
		return
	}

	//file, err := os.Open("doc.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//p := &telego.SendMediaGroupParams{
	//	ChatID: myID,
	//	Media: []telego.InputMedia{
	//		telego.InputMediaDocument{
	//			Type:  "document",
	//			Media: telego.InputFile{File: file},
	//		},
	//		//telego.InputMediaPhoto{
	//		//	Type:  "photo",
	//		//	Media: telego.InputFile{URL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRzJZk-efp0id1yxpUHPYwJ1t8vuAwMI_SXfh77dRFWsg1X1ancplws5_DH_WSJ52MHyH8&usqp=CAU"},
	//		//},
	//		//telego.InputMediaPhoto{
	//		//	Type:  "photo",
	//		//	Media: telego.InputFile{URL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTqSw1_1Ar_u3f2lVhYkhz-R0KaaZtDKwx6Y5H1HGceAmx0sqexKzXkSawLG5PRoRKcy6A&usqp=CAU"},
	//		//},
	//	},
	//}
	//msgs, err := bot.SendMediaGroup(p)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, m := range msgs {
	//	fmt.Println(m)
	//}

	//err = bot.SetMyCommands(&telego.SetMyCommandsParams{
	//	Commands: []telego.BotCommand{
	//		{
	//			Command:     "test",
	//			Description: "Test",
	//		},
	//	},
	//	Scope: telego.BotCommandScopeAllGroupChats{Type: "all_group_chats"},
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//commands, err := bot.GetMyCommands(nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//for _, c := range commands {
	//	fmt.Println(c.Command, c.Description)
	//}

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
	//	fmt.Println(u.Message.Chat)
	//}
	//
	//p := &telego.GetChatAdministratorsParams{ChatID: telego.ChatID{ID: -1001516926498}}
	//admins, err := bot.GetChatAdministrators(p)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//for _, u := range admins {
	//	switch cm := u.(type) {
	//	case telego.ChatMemberAdministrator:
	//		fmt.Println("admin", cm.User)
	//	case telego.ChatMemberOwner:
	//		fmt.Println("owner", cm.User)
	//	default:
	//		fmt.Println(cm.MemberStatus())
	//	}
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
