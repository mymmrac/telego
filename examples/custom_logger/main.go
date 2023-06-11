package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime/debug"
	"strings"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoutil"
)

const (
	ansiReset = "\u001B[0m"
	ansiRedBG = "\u001B[41m"
	BUG       = ansiRedBG + "Ж" + ansiReset
)

var (
	// Log error with time and source of error
	letf = log.New(os.Stdout, BUG, log.Ltime|log.Lshortfile)
	// Log debug with time and source of debug
	ltf = log.New(os.Stdout, " ", log.Ltime|log.Lshortfile)

	// Log error with time for custom loger
	let = log.New(os.Stdout, BUG, log.Ltime)
	// Log debug with time for custom loger
	lt = log.New(os.Stdout, " ", log.Ltime)
)

// Custom loger type
type customLogger struct{}

// Custom logger method for debug
func (customLogger) Debugf(format string, args ...any) {
	lt.Print(woToken(format, args...))
}

// Custom logger method for error
func (customLogger) Errorf(format string, args ...any) {
	let.Print(woToken(format, args...))
}

func main() {

	// Try create bot with empty botToken
	botToken := ""

	ltf.Println("Try create bot with empty botToken")
	_, err := telego.NewBot(botToken)
	if err != nil {
		letf.Println(err)
	}

	// Create bot with custom logger
	botToken = os.Getenv("TOKEN")

	ltf.Println("Create bot with custom logger")
	bot, err := telego.NewBot(botToken,
		// Create you custom logger that implements telego.Logger (default: telego has build in default logger)
		// Note: Please keep in mind that logger may expose sensitive information, use in development only or configure
		// it not to leak unwanted content
		telego.WithLogger(telego.Logger(customLogger{})),
	)
	if err != nil {
		letf.Println(err)
		os.Exit(1)
	}

	// Call method DeleteMessage
	bot.Logger().Debugf("Call method DeleteMessage")
	err = bot.DeleteMessage(telegoutil.Delete(telegoutil.ID(1), 1))
	if err != nil {
		bot.Logger().Errorf("%+v\n", err)
	}
}

// Get source of code
func src(deep int) (s string) {
	s = string(debug.Stack())
	s = strings.Split(s, "\n")[deep]
	s = strings.Split(s, " +0x")[0]
	_, s = path.Split(s)
	s += ":"
	return
}

// Hide bot token
func woToken(format string, args ...any) (s string) {
	s = src(10) + " " + fmt.Sprintf(format, args...)
	btStart := strings.Index(s, "/bot") + 4
	if btStart > 4-1 {
		btLen := strings.Index(s[btStart:], "/")
		if btLen > 0 {
			s = s[:btStart] + s[btStart+btLen:]
		}
	}
	return
}
