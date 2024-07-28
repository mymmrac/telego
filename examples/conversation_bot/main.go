package main

import (
	"fmt"
	"net/mail"
	"os"
	"strconv"
	"sync"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// State of the user
type State uint

const (
	StateDefault State = iota
	StateAskName
	StateAskAge
	StateAskEmail
	StateConfirm
)

// User data and state
type User struct {
	State State
	Name  string
	Age   uint
	Email string
}

func main() {
	botToken := os.Getenv("TOKEN")

	// Create Bot with debug on
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler
	bh, _ := th.NewBotHandler(bot, updates)

	// Create user storage (for this simple exampled in-memory map is sufficient,
	// but in the real world you might want to use persistent storage like PostgreSQL or Redis)
	users := make(map[int64]User)
	// Since this is in-memory storage, we must use mutex
	lock := sync.RWMutex{}

	// Handle any message
	bh.HandleMessage(func(bot *telego.Bot, msg telego.Message) {
		userID := msg.From.ID

		lock.RLock()
		user := users[userID]
		lock.RUnlock()

		var text string
		switch user.State {
		case StateDefault:
			// Welcome message for new users
			text = "Hello stranger, what's your name?"
			user.State = StateAskName
		case StateAskName:
			// Specify name (no validation)
			user.Name = msg.Text
			user.State = StateAskAge
			text = "How old are you?"
		case StateAskAge:
			// Specify age (validate that its positive number)
			var age uint64
			age, err = strconv.ParseUint(msg.Text, 10, 32)
			if err != nil || age == 0 {
				text = "Invalid age, please try again"
				// No state change
			} else {
				user.Age = uint(age)
				user.State = StateAskEmail
				text = "What's your email?"
			}
		case StateAskEmail:
			// Specify email (validate that its valid email address)
			var addr *mail.Address
			addr, err = mail.ParseAddress(msg.Text)
			if err != nil {
				text = "Invalid email, please try again"
				// No state change
			} else {
				user.Email = addr.Address
				user.State = StateConfirm
				text = fmt.Sprintf(
					"Your name is %s, your age is %d, and your email is %s, all corrent? (Y/N)",
					user.Name, user.Age, user.Email,
				)
			}
		case StateConfirm:
			if msg.Text == "Y" {
				text = "Thanks for your data!"
				user.State = StateDefault
				// Do some logic
			} else {
				text = "Ok, let's start again"
				user.State = StateDefault
			}
		default:
			panic("unknown state")
		}

		lock.Lock()
		users[userID] = user
		lock.RUnlock()

		_, _ = bot.SendMessage(tu.Message(msg.Chat.ChatID(), text))
	})

	// Stop handling updates on exit
	defer bh.Stop()
	defer bot.StopLongPolling()

	// Start handling updates
	bh.Start()
}
