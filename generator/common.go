package main

import (
	"html"
	"regexp"
	"strings"
)

const packageName = "telego"
const docsURL = "https://core.telegram.org/bots/api"

const (
	urlPattern = `<a.*?href="(.+?)".*?>(.*?)</a>`
	imgPattern = `<img.*?alt="(.+?)".*?>`

	tagPattern         = `<.+?>(.+?)</.+?>`
	unclosedTagPattern = `<.+?>`
)

var (
	urlPatternReg = regexp.MustCompile(urlPattern)
	imgPatternReg = regexp.MustCompile(imgPattern)

	tagPatternReg         = regexp.MustCompile(tagPattern)
	unclosedTagPatternReg = regexp.MustCompile(unclosedTagPattern)
)

func removeNewline(text string) string {
	return strings.ReplaceAll(text, "\n", "")
}

func cleanDescription(text string) string {
	return imgPatternReg.ReplaceAllString(
		urlPatternReg.ReplaceAllString(text, "$2 ($1)"), "$1")
}

func removeTags(text string) string {
	return html.UnescapeString(
		unclosedTagPatternReg.ReplaceAllString(
			tagPatternReg.ReplaceAllString(text, "$1"), ""))
}

func snakeToCamelCase(text string) string {
	nextUpper := true
	sb := strings.Builder{}
	sb.Grow(len(text))
	for _, v := range []byte(text) {
		if v == '_' {
			nextUpper = true
			continue
		}

		if nextUpper {
			nextUpper = false
			v += 'A'
			v -= 'a'
			sb.WriteByte(v)
		} else {
			sb.WriteByte(v)
		}
	}

	return sb.String()
}
