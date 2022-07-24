package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

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

const testCase = 21

func main() {
	testToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(testToken, telego.WithDefaultDebugLogger())
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
			fmt.Println(bot.IsRunningLongPulling())

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

		bh, _ := th.NewBotHandler(bot, updates)
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

		bh, _ := th.NewBotHandler(bot, updates)

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

		bh, _ := th.NewBotHandler(bot, updates)

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			matches := th.CommandRegexp.FindStringSubmatch(msg.Text)
			_, _ = bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), fmt.Sprintf("%#v", matches)))
		}, th.AnyCommand())

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			_, _ = bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), fmt.Sprintf("Whaaat? %s", msg.Text)))
		}, th.AnyMessage(), th.Not(th.AnyCommand()))

		bh.Start()
		defer bh.Stop()
	case 18:
		updates, err := bot.UpdatesViaLongPulling(nil)
		assert(err == nil, err)

		defer bot.StopLongPulling()

		bh, err := th.NewBotHandler(bot, updates)
		assert(err == nil, err)

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			_, _ = bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), "Running test"))
		}, th.CommandEqualArgv("run", "test"))

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			_, _ = bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), "Running update"))
		}, th.CommandEqualArgv("run", "update"))

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			msg := update.Message
			m := tu.Message(tu.ID(msg.Chat.ID), "Run usage:\n```/run test```\n```/run update```")
			m.ParseMode = telego.ModeMarkdownV2
			_, _ = bot.SendMessage(m)
		}, th.Union(
			th.CommandEqualArgc("run", 0),
			th.CommandEqualArgv("help", "run"),
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
		}, th.AnyCommand())

		bh.Start()
		defer bh.Stop()
	case 19:
		updates, _ := bot.UpdatesViaLongPulling(nil)
		defer bot.StopLongPulling()

		bh, _ := th.NewBotHandler(bot, updates)

		bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
			_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Hmm?"))
		}, th.TextEqual("Hmm"))

		bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
			_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Hello"))
		})

		bh.Start()
		defer bh.Stop()
	case 20:
		img := tu.File(mustOpen("img1.jpg"))
		img2 := tu.File(mustOpen("img2.jpg"))
		audio := tu.File(mustOpen("kitten.mp3"))
		voice := tu.File(mustOpen("kitten.ogg"))
		doc := tu.File(mustOpen("doc.txt"))
		video := tu.File(mustOpen("sample.mp4"))
		note := tu.File(mustOpen("note.mp4"))
		gif := tu.File(mustOpen("cat.mp4"))

		_, err = bot.SendMessage(tu.Message(myID, "Test"))
		assert(err == nil, err)

		_, err = bot.SendPhoto(tu.Photo(myID, img))
		assert(err == nil, err)

		_, err = bot.SendAudio(tu.Audio(myID, audio))
		assert(err == nil, err)

		_, err = bot.SendDocument(tu.Document(myID, doc))
		assert(err == nil, err)

		time.Sleep(time.Second * 3)

		_, err = bot.SendVideo(tu.Video(myID, video))
		assert(err == nil, err)

		_, err = bot.SendAnimation(tu.Animation(myID, gif))
		assert(err == nil, err)

		_, err = bot.SendVoice(tu.Voice(myID, voice))
		assert(err == nil, err)

		_, err = bot.SendVideoNote(tu.VideoNote(myID, note))
		assert(err == nil, err)

		time.Sleep(time.Second * 3)

		img = tu.File(mustOpen("img1.jpg"))
		img2 = tu.File(mustOpen("img2.jpg"))

		_, err = bot.SendMediaGroup(tu.MediaGroup(myID, tu.MediaPhoto(img), tu.MediaPhoto(img2)))
		assert(err == nil, err)

		_, err = bot.SendLocation(tu.Location(myID, 42, 24))
		assert(err == nil, err)

		_, err = bot.SendVenue(tu.Venue(myID, 42, 24, "The Thing", "Things str."))
		assert(err == nil, err)

		_, err = bot.SendContact(tu.Contact(myID, "+424242", "The 42"))
		assert(err == nil, err)

		time.Sleep(time.Second * 3)

		_, err = bot.SendPoll(tu.Poll(myID, "42?", "42", "24"))
		assert(err == nil, err)

		_, err = bot.SendDice(tu.Dice(myID, telego.EmojiBasketball))
		assert(err == nil, err)

		err = bot.SendChatAction(tu.ChatAction(myID, telego.ChatActionTyping))
		assert(err == nil, err)
	case 21:
		updates, _ := bot.UpdatesViaLongPulling(nil, telego.WithLongPullingUpdateInterval(time.Second))
		defer bot.StopLongPulling()

		bh, _ := th.NewBotHandler(bot, updates)

		bh.HandleInlineQuery(func(bot *telego.Bot, query telego.InlineQuery) {
			err = bot.AnswerInlineQuery(&telego.AnswerInlineQueryParams{
				InlineQueryID: query.ID,
				Results: []telego.InlineQueryResult{
					&telego.InlineQueryResultArticle{
						Type:                telego.ResultTypeArticle,
						ID:                  "1",
						Title:               "Hmm",
						InputMessageContent: tu.TextMessage("Hmm"),
						ReplyMarkup: tu.InlineKeyboard(tu.InlineKeyboardRow(
							tu.InlineKeyboardButton("GG?").WithCallbackData("ok"),
						)),
					},
				},
			})
			assert(err == nil, err)
		})

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
			_, err = bot.EditMessageText(&telego.EditMessageTextParams{
				Text:            "GG?",
				InlineMessageID: query.InlineMessageID,
			})
			assert(err == nil, err)

			err = bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
				CallbackQueryID: query.ID,
				Text:            "OK",
			})
			assert(err == nil, err)
		})

		defer bh.Stop()
		bh.Start()
	case 22:
		updates, _ := bot.UpdatesViaLongPulling(nil)
		defer bot.StopLongPulling()

		bh, _ := th.NewBotHandler(bot, updates)

		auth := func(update telego.Update) bool {
			var userID int64

			if update.Message != nil && update.Message.From != nil {
				userID = update.Message.From.ID
			}

			if update.CallbackQuery != nil {
				userID = update.CallbackQuery.From.ID
			}

			if userID == 0 {
				return false
			}

			if userID == 1234 {
				return true
			}

			return false
		}

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			// DO AUTHORIZED STUFF...
		}, auth)

		bh.Handle(func(bot *telego.Bot, update telego.Update) {
			// DO NOT AUTHORIZED STUFF...
		}, th.Not(auth))

		defer bh.Stop()
		bh.Start()
	case 23:
		updates, _ := bot.UpdatesViaLongPulling(nil)
		defer bot.StopLongPulling()

		bh, _ := th.NewBotHandler(bot, updates)

		ok := false
		middleware := func(update telego.Update) bool {
			return ok
		}

		bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
			ok = true
			fmt.Println("SET OK")
		}, th.CommandEqual("ok"))

		bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
			fmt.Println("OK")
		}, middleware)

		defer bh.Stop()
		bh.Start()
	case 24:
		updates, err := bot.UpdatesViaLongPulling(nil)
		assert(err == nil, err)

		fmt.Println(bot.IsRunningLongPulling())
		time.Sleep(time.Second * 10)

		fmt.Println(bot.IsRunningLongPulling())
		bot.StopLongPulling()
		fmt.Println(bot.IsRunningLongPulling())

		for upd := range updates {
			fmt.Println(upd)
		}
	case 25:
		fmt.Println(bot.IsRunningWebhook())

		err = bot.StopWebhook()
		assert(err == nil, err)

		_, err = bot.UpdatesViaWebhook("/")
		assert(err == nil, err)

		fmt.Println(bot.IsRunningWebhook())

		err = bot.StartListeningForWebhook(":8080")
		assert(err == nil, err)

		fmt.Println(bot.IsRunningWebhook())

		err = bot.StopWebhook()
		assert(err == nil, err)

		err = bot.StopWebhook()
		assert(err == nil, err)

		fmt.Println(bot.IsRunningWebhook())

		fmt.Println("====")

		_, err = bot.UpdatesViaWebhook("/")
		assert(err == nil, err)

		fmt.Println(bot.IsRunningWebhook())

		err = bot.StartListeningForWebhook(":8080")
		assert(err == nil, err)

		fmt.Println(bot.IsRunningWebhook())

		err = bot.StopWebhook()
		assert(err == nil, err)

		fmt.Println(bot.IsRunningWebhook())
	case 26:
		r := router.New()
		r.GET("/", func(ctx *fasthttp.RequestCtx) {
			ctx.SetStatusCode(fasthttp.StatusAccepted)
		})

		_, err = bot.UpdatesViaWebhook("/", telego.WithWebhookRouter(r), telego.WithWebhookHealthAPI())
		assert(err == nil, err)

		err = bot.StartListeningForWebhook(":8080")
		assert(err == nil, err)

		defer func() {
			_ = bot.StopWebhook()
		}()
		select {}
	case 27:
		note := tu.File(mustOpen("note.mp4"))

		_, err = bot.SendVideoNote(tu.VideoNote(myID, note))
		assert(err == nil, err)
	}
}

func assert(ok bool, args ...interface{}) {
	if !ok {
		fmt.Println(args...)
		os.Exit(1)
	}
}

func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
