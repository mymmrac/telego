package web_apps

import (
	"context"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	ctx := context.Background()
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set chat button to open Web App
	_ = bot.SetChatMenuButton(ctx, &telego.SetChatMenuButtonParams{
		ChatID: 123,
		MenuButton: &telego.MenuButtonWebApp{
			Type: telego.ButtonTypeWebApp,
			Text: "Example",
			WebApp: telego.WebAppInfo{
				URL: "https://www.example.com/main",
			},
		},
	})

	// Also from @BotFather you can set Main Mini App
	// (https://core.telegram.org/bots/webapps#launching-the-main-mini-app)

	// Use keyboard buttons to open Mini App
	_, _ = bot.SendMessage(ctx, tu.Message(tu.ID(123), "Hello, World!").
		WithReplyMarkup(tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("Example").WithWebApp(tu.WebAppInfo("https://www.example.com/inline")),
			),
		)),
	)

	// In similar way inline keyboard button can also be used
	_, _ = bot.SendMessage(ctx, tu.Message(tu.ID(123), "Hello, World!").
		WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Example").WithWebApp(tu.WebAppInfo("https://www.example.com/inline")),
			),
		)),
	)

	// Button in inline query results
	_ = bot.AnswerInlineQuery(ctx, &telego.AnswerInlineQueryParams{
		InlineQueryID: "some-id",
		Results: []telego.InlineQueryResult{
			&telego.InlineQueryResultArticle{
				// ...
			},
		},
		Button: &telego.InlineQueryResultsButton{
			Text:   "Example",
			WebApp: &telego.WebAppInfo{URL: "https://www.example.com/inline-query"},
		},
	})

	// When receiving data from Mini Apps you MUST validate data received from client, you can use helper function
	values, _ := tu.ValidateWebAppData(bot.Token(), "window.Telegram.WebApp.initData")
	fmt.Println(values)
}
