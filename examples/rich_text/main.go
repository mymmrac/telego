package main

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

	// Sending rich message
	msg, err := bot.SendRichMessage(ctx, &telego.SendRichMessageParams{
		ChatID: tu.ID(1234567),
		RichMessage: tu.RichMessage(
			tu.RichBlockSectionHeading(tu.RichTextPlain("Hello"), 1),
			tu.RichBlockParagraph(tu.RichTextPlain("World")),
			tu.RichBlockPhoto(*tu.MediaPhoto(tu.FileFromURL("https://telegram.org/example/photo.jpg"))).
				WithCaption(
					tu.RichBlockCaption(tu.RichTextPlain("Image")).
						WithCredit(tu.RichTextPlain("Telegram")),
				),
			tu.RichBlockCollage(
				tu.RichBlockPhoto(*tu.MediaPhoto(tu.FileFromURL("https://telegram.org/example/photo.jpg"))).
					WithCaption(tu.RichBlockCaption(tu.RichTextBold(tu.RichTextPlain("Image 1")))),
				tu.RichBlockAnimation(*tu.MediaAnimation(tu.FileFromURL("https://telegram.org/example/animation.gif"))).
					WithCaption(tu.RichBlockCaption(tu.RichTextItalic(tu.RichTextPlain("Animation 2")))),
			).WithCaption(tu.RichBlockCaption(tu.RichTextPlain("Collage"))),
			tu.RichBlockTableGrid(tu.RichBlockTableCols(3,
				tu.RichBlockTableCell(tu.RichTextPlain("Cell 1")),
				tu.RichBlockTableCell(tu.RichTextStrikethrough(tu.RichTextBold(tu.RichTextPlain("Cell 2")))),
				tu.RichBlockTableCell(tu.RichTextPlain("Cell 3")),
				tu.RichBlockTableCell(tu.RichTextSpoiler(tu.RichTextPlain("Cell 4"))),
			)),
		),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg.RichMessage)
}
