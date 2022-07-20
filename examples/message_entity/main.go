package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Create Bot
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Send message with provided message entities
	_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(1234567),
		tu.Entity("Hi").Bold(), tu.Entity(" "), tu.Entity("There").Italic().Spoiler(), tu.Entity("\n"),
		tu.Entity("The Link").TextLink("https://example.com").Italic(), tu.Entity("\n"),
		tu.Entity("User: "), tu.Entity("???").TextMentionWithID(1234567),
	))

	// Create text with corresponding message entities
	text, entities := tu.MessageEntities(
		// No entity text (pain text)
		tu.Entity("telego"),

		// Formatting
		tu.Entity("telego").Italic(),        // “italic” (italic text)
		tu.Entity("telego").Bold(),          // “bold” (bold text)
		tu.Entity("telego").Strikethrough(), // “strikethrough” (strikethrough text)
		tu.Entity("telego").Underline(),     // “underline” (underlined text)
		tu.Entity("telego").Spoiler(),       // “spoiler” (spoiler message)
		tu.Entity("telego").Code(),          // “code” (monowidth string)
		tu.Entity("telego").Pre("go"),       // “pre” (monowidth block)
		tu.Entity("telego").Hashtag(),       // “hashtag” (#hashtag)
		tu.Entity("telego").Cashtag(),       // “cashtag” ($USD)
		tu.Entity("telego").URL(),           // “url” (https://telegram.org)
		tu.Entity("telego").BotCommand(),    // “bot_command” (/start@jobs_bot)
		tu.Entity("telego").Email(),         // “email” (do-not-reply@telegram.org)
		tu.Entity("telego").PhoneNumber(),   // “phone_number” (+1-212-555-0123)
		tu.Entity("telego").Mention(),       // “mention” (@username)

		// Links
		tu.Entity("telego").TextLink("https://example.com"), // “text_link” (for clickable text URLs)

		// Mentions
		// “text_mention” (for users without usernames (https://telegram.org/blog/edit#new-mentions))
		tu.Entity("telego").TextMention(&telego.User{}),
		tu.Entity("telego").TextMentionWithID(1234567),

		// Combination
		tu.Entity("telego").Italic().Bold().Spoiler(),
		tu.Entity("telego").URL().Bold(),
		tu.Entity("telego").Spoiler().Email(),
	)

	fmt.Println()
	fmt.Printf("Text: %q\nEntities: %+v\n", text, entities)
}
