package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var (
	myID            = tu.ID(331849104)
	groupID         = tu.ID(-1001516926498)
	channelUsername = tu.Username("@mymmrTest")
	groupUsername   = tu.Username("@botesup")
	userUsername    = tu.Username("@mymmrac")
)

const testCase = 18

func main() {
	testToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(testToken,
		telego.WithDefaultLogger(true, true))
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
		updChan, err := bot.UpdatesViaLongPulling(&telego.GetUpdatesParams{})
		if err != nil {
			fmt.Println(err)
			return
		}
		defer bot.StopLongPulling()

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
	case 13:
		msg := &telego.SendMessageParams{
			ChatID: myID,
			Text: `	case 12:
		msg := &telego.SendMessageParams{
			ChatID: channelUsername,
			Text:   "Test msg",
		}
		_, err = bot.SendMessage(msg)
		if err != nil {
			fmt.Println(err)
			return
		}`,
		}

		msg.Entities = []telego.MessageEntity{
			{
				Type:     telego.EntityTypePre,
				Offset:   0,
				Length:   len(msg.Text),
				Language: "go",
			},
		}

		_, err = bot.SendMessage(msg)
		if err != nil {
			fmt.Println(err)
			return
		}
	case 14:
		_, err := bot.SendMessage(tu.Message(groupUsername, "Test 1"))
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = bot.SendMessage(tu.Message(userUsername, "Test 2"))
		if err != nil {
			fmt.Println(err)
			return
		}
	case 15:
		updates, _ := bot.UpdatesViaLongPulling(nil)
		defer bot.StopLongPulling()

		bh := th.NewBotHandler(bot, updates)
		defer bh.Stop()

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			fmt.Println(update.Message.Text)
		}, func(update telego.Update) bool {
			return update.Message != nil
		})

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			fmt.Println("====")
			fmt.Println(update.Message.Text)
			fmt.Println("====")
		}, func(update telego.Update) bool {
			return update.Message != nil && update.Message.Text == "OK"
		})

		bh.Start()
	case 16:
		updates, _ := bot.UpdatesViaLongPulling(nil)
		defer bot.StopLongPulling()

		bh := th.NewBotHandler(bot, updates)

		count := 0

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			fmt.Println("ZERO")
			_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), fmt.Sprintf("Count is zero")))
			count = 1
		}, func(update telego.Update) bool {
			return update.Message != nil && count == 0
		})

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			fmt.Println("ONE")
			_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), fmt.Sprintf("Count is one")))
			count = 2
		}, func(update telego.Update) bool {
			return update.Message != nil && count == 1
		})

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			fmt.Println("BIG")
			_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), fmt.Sprintf("Count is big: %d", count)))
			count++
		}, func(update telego.Update) bool {
			return update.Message != nil && count > 1
		})

		bh.Start()
		defer bh.Stop()
	case 17:
		updates, _ := bot.UpdatesViaLongPulling(nil)
		defer bot.StopLongPulling()

		bh := th.NewBotHandler(bot, updates)

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			matches := th.CommandRegexp.FindStringSubmatch(msg.Text)
			_, _ = bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), fmt.Sprintf("%#v", matches)))
		}, th.HasCommand())

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			_, _ = bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), fmt.Sprintf("Whaaat? %s", msg.Text)))
		}, th.HasMassage(), th.Not(th.HasCommand()))

		bh.Start()
		defer bh.Stop()
	case 18:
		updates, _ := bot.UpdatesViaLongPulling(nil)
		defer bot.StopLongPulling()

		bh := th.NewBotHandler(bot, updates)

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			_, _ = bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), "Running test"))
		}, th.CommandEqualWithArgv("run", "test"))

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			_, _ = bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), "Running update"))
		}, th.CommandEqualWithArgv("run", "update"))

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			m := tu.Message(tu.ID(msg.Chat.ID), "Run usage:\n```/run test```\n```/run update```")
			m.ParseMode = telego.ModeMarkdownV2
			_, _ = bot.SendMessage(m)
		}, th.Union(
			th.CommandEqualWithArgc("run", 0),
			th.CommandEqualWithArgv("help", "run"),
		))

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			m := tu.Message(tu.ID(msg.Chat.ID), "Unknown subcommand\nRun usage:\n```/run test```\n```/run update```")
			m.ParseMode = telego.ModeMarkdownV2
			_, _ = bot.SendMessage(m)
		}, th.CommandEqual("run"))

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			_, _ = bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), "Help: /run"))
		}, th.CommandEqual("help"))

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			_, _ = bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), "Unknown command, use: /run"))
		}, th.HasCommand())

		bh.Start()
		defer bh.Stop()
	}
}

func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
