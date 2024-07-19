//go:build integration

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/chococola/telego"
)

const (
	img1Jpg   = "img1.jpg"
	kittenMp3 = "kitten.mp3"

	exampleMp3 = "https://file-examples.com/storage/fe0e4ffeec64469f8a2ba23/2017/11/file_example_MP3_700KB.mp3"
)

var (
	bot     *telego.Bot
	timeNow string

	chatID int64
)

func TestMain(m *testing.M) {
	var err error
	bot, err = telego.NewBot(env("TOKEN"), telego.WithDiscardLogger())
	expect(err == nil, "Create bot:", err)

	chatID, err = strconv.ParseInt(env("CHAT_ID"), 10, 64)
	expect(err == nil, "Parse chat ID:", err)

	timeNow = time.Now().Local().String()

	os.Exit(m.Run())
}

func expect(ok bool, args ...any) {
	if !ok {
		fmt.Println(args...)
		os.Exit(1)
	}
}

func env(key string) string {
	value, ok := os.LookupEnv(key)
	expect(ok, "Environment variable", key, "not set")

	return value
}

func open(filename string) *os.File {
	file, err := os.Open(filepath.Join("files", filename))
	expect(err == nil, "Open file:", err)

	return file
}
