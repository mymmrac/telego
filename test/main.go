package main

import (
	"fmt"
	"os"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

var myID = telego.ChatID{ID: 331849104}
var groupID = telego.ChatID{ID: -1001516926498}

func main() {
	testToken := os.Getenv("TEST_TOKEN")

	bot, err := telego.NewBot(testToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.DefaultLogger(true, true)

	_, err = bot.GetMe()
	if err != nil {
		fmt.Println(err)
		return
	}

	//message := &telego.SendMessageParams{
	//	ChatID: myID,
	//	Text:   "Test",
	//	ReplyMarkup: &telego.ReplyKeyboardMarkup{
	//		Keyboard: [][]telego.KeyboardButton{
	//			{
	//				{
	//					Text: "1",
	//				},
	//				{
	//					Text: "2",
	//				},
	//			},
	//			{
	//				{
	//					Text: "3",
	//				},
	//			},
	//		},
	//		ResizeKeyboard: true,
	//		//OneTimeKeyboard:       true,
	//		InputFieldPlaceholder: "Number?",
	//	},
	//}
	//
	//msg, _ := bot.SendMessage(message)
	//fmt.Println(msg)

	//updChan, err := bot.GetUpdatesChan(&telego.GetUpdatesParams{})
	//defer bot.StopGettingUpdates()
	//
	//for upd := range updChan {
	//	fmt.Println(upd)
	//
	//	if upd.Message != nil {
	//		_, err := bot.CopyMessage(&telego.CopyMessageParams{
	//			ChatID:     telego.ChatID{ID: upd.Message.Chat.ID},
	//			FromChatID: telego.ChatID{ID: upd.Message.Chat.ID},
	//			MessageID:  upd.Message.MessageID,
	//		})
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//	}
	//}

	//p := &telego.ExportChatInviteLinkParams{ChatID: groupID}
	//link, err := bot.ExportChatInviteLink(p)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(*link)

	//p := &telego.SendMediaGroupParams{
	//	ChatID: myID,
	//	Media: []telego.InputMedia{
	//		//&telego.InputMediaDocument{
	//		//	Type:  "document",
	//		//	Media: telego.InputFile{File: mustOpen("doc.txt")},
	//		//},
	//		&telego.InputMediaPhoto{
	//			Type:  "photo",
	//			Media: telego.InputFile{URL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRzJZk-efp0id1yxpUHPYwJ1t8vuAwMI_SXfh77dRFWsg1X1ancplws5_DH_WSJ52MHyH8&usqp=CAU"},
	//		},
	//		//telego.InputMediaPhoto{
	//		//	Type:  "photo",
	//		//	Media: telego.InputFile{URL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTqSw1_1Ar_u3f2lVhYkhz-R0KaaZtDKwx6Y5H1HGceAmx0sqexKzXkSawLG5PRoRKcy6A&usqp=CAU"},
	//		//},
	//		&telego.InputMediaPhoto{
	//			Type:  "photo",
	//			Media: telego.InputFile{File: mustOpen("img1.jpg")},
	//		},
	//		&telego.InputMediaPhoto{
	//			Type:  "photo",
	//			Media: telego.InputFile{File: mustOpen("img2.jpg")},
	//		},
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
	//			Description: "Test OK",
	//		},
	//	},
	//	Scope: &telego.BotCommandScopeAllGroupChats{Type: "all_group_chats"},
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

	dp := &telego.SendDocumentParams{
		ChatID: myID,
		//Document: telego.InputFile{File: mustOpen("doc.txt")},
		Document: telego.InputFile{FileID: "BQACAgIAAxkDAAMmYP_FFDZSpqgMsWpK0GCB3hQaI8MAApUPAALeHgABSHe5TRKuQ2NGIAQ"},
		//Caption:  "Hello world",
		//ReplyMarkup: &telego.InlineKeyboardMarkup{InlineKeyboard: [][]telego.InlineKeyboardButton{
		//	{
		//		{
		//			Text:         "Test",
		//			CallbackData: "1",
		//		},
		//	},
		//}},
	}
	msg, err := bot.SendDocument(dp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg.Document)
}

func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
