package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
)

var (
	myID            = telego.ChatID{ID: 331849104}
	groupID         = telego.ChatID{ID: -1001516926498}
	channelUsername = telego.ChatID{Username: "@mymmrTest"}
)

const testCase = 2

func main() {
	testToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(testToken,
		telego.DefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = bot.GetMe()
	if err != nil {
		fmt.Println(err)
		return
	}

	switch testCase {
	case 1:
		message := &telego.SendMessageParams{
			ChatID: myID,
			Text:   "Test",
			ReplyMarkup: &telego.ReplyKeyboardMarkup{
				Keyboard: [][]telego.KeyboardButton{
					{
						{
							Text: "1",
						},
						{
							Text: "2",
						},
					},
					{
						{
							Text: "3",
						},
					},
				},
				ResizeKeyboard:        true,
				OneTimeKeyboard:       true,
				InputFieldPlaceholder: "Number?",
			},
		}

		msg, err := bot.SendMessage(message)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)
	case 2:
		updChan, err := bot.GetUpdatesViaLongPulling(&telego.GetUpdatesParams{})
		if err != nil {
			fmt.Println(err)
			return
		}
		defer bot.StopGettingUpdates()

		for upd := range updChan {
			fmt.Println(upd)

			if upd.Message != nil {
				_, err := bot.CopyMessage(&telego.CopyMessageParams{
					ChatID:     telego.ChatID{ID: upd.Message.Chat.ID},
					FromChatID: telego.ChatID{ID: upd.Message.Chat.ID},
					MessageID:  upd.Message.MessageID,
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	case 3:
		p := &telego.ExportChatInviteLinkParams{ChatID: groupID}
		link, err := bot.ExportChatInviteLink(p)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*link)
	case 4:
		p := &telego.SendMediaGroupParams{
			ChatID: myID,
			Media: []telego.InputMedia{
				//&telego.InputMediaDocument{
				//	Type:  "document",
				//	Media: telego.InputFile{File: mustOpen("doc.txt")},
				//},
				&telego.InputMediaPhoto{
					Type:  "photo",
					Media: telego.InputFile{URL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRzJZk-efp0id1yxpUHPYwJ1t8vuAwMI_SXfh77dRFWsg1X1ancplws5_DH_WSJ52MHyH8&usqp=CAU"},
				},
				//telego.InputMediaPhoto{
				//	Type:  "photo",
				//	Media: telego.InputFile{URL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTqSw1_1Ar_u3f2lVhYkhz-R0KaaZtDKwx6Y5H1HGceAmx0sqexKzXkSawLG5PRoRKcy6A&usqp=CAU"},
				//},
				&telego.InputMediaPhoto{
					Type:  "photo",
					Media: telego.InputFile{File: mustOpen("img1.jpg")},
				},
				&telego.InputMediaPhoto{
					Type:  "photo",
					Media: telego.InputFile{File: mustOpen("img2.jpg")},
				},
			},
		}
		msgs, err := bot.SendMediaGroup(p)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, m := range msgs {
			fmt.Println(m)
		}
	case 5:
		err = bot.SetMyCommands(&telego.SetMyCommandsParams{
			Commands: []telego.BotCommand{
				{
					Command:     "test",
					Description: "Test OK",
				},
			},
			Scope: &telego.BotCommandScopeAllGroupChats{Type: "all_group_chats"},
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	case 6:
		commands, err := bot.GetMyCommands(nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, c := range commands {
			fmt.Println(c.Command, c.Description)
		}
	case 7:
		updParams := &telego.GetUpdatesParams{
			AllowedUpdates: []string{"chat_member"},
		}
		upd, err := bot.GetUpdates(updParams)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, u := range upd {
			fmt.Println(u.Message.Chat)
		}
	case 8:
		p := &telego.GetChatAdministratorsParams{ChatID: telego.ChatID{ID: -1001516926498}}
		admins, err := bot.GetChatAdministrators(p)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, u := range admins {
			switch cm := u.(type) {
			case *telego.ChatMemberAdministrator:
				fmt.Println("admin", cm.User)
			case *telego.ChatMemberOwner:
				fmt.Println("owner", cm.User)
			default:
				fmt.Println(cm.MemberStatus())
			}
		}
	case 9:
		dp := &telego.SendDocumentParams{
			ChatID:   myID,
			Document: telego.InputFile{FileID: "BQACAgIAAxkDAAMmYP_FFDZSpqgMsWpK0GCB3hQaI8MAApUPAALeHgABSHe5TRKuQ2NGIAQ"},
			ReplyMarkup: &telego.InlineKeyboardMarkup{InlineKeyboard: [][]telego.InlineKeyboardButton{
				{
					{
						Text:         "Test",
						CallbackData: "1",
					},
				},
			}},
		}
		msg, err := bot.SendDocument(dp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg.Document)
	case 10:
		dp := &telego.SendDocumentParams{
			ChatID:   myID,
			Document: telego.InputFile{File: mustOpen("doc.txt")},
			Caption:  "Hello world",
		}
		msg, err := bot.SendDocument(dp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg.Document)
	case 11:
		photo := &telego.SendPhotoParams{
			ChatID:  channelUsername,
			Photo:   telego.InputFile{File: mustOpen("img1.jpg")},
			Caption: "https://test.ua/test_url",
		}

		msg, err := bot.SendPhoto(photo)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)
	case 12:
		msg := &telego.SendMessageParams{
			ChatID: channelUsername,
			Text:   "Test msg",
		}
		_, err = bot.SendMessage(msg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
