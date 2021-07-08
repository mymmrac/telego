package generator

import (
	"html"
	"regexp"
	"strings"
)

const PackageName = "telego"
const DocsURL = "https://core.telegram.org/bots/api"
const MaxLineLen = 110

const (
	URLPattern = `<a.*?href="(.+?)".*?>(.*?)</a>`
	ImgPattern = `<img.*?alt="(.+?)".*?>`

	TagPattern         = `<.+?>(.+?)</.+?>`
	UnclosedTagPattern = `<.+?>`
)

var (
	URLPatternReg = regexp.MustCompile(URLPattern)
	ImgPatternReg = regexp.MustCompile(ImgPattern)

	TagPatternReg         = regexp.MustCompile(TagPattern)
	UnclosedTagPatternReg = regexp.MustCompile(UnclosedTagPattern)
)

func RemoveNewline(text string) string {
	return strings.ReplaceAll(text, "\n", "")
}

func CleanDescription(text string) string {
	return ImgPatternReg.ReplaceAllString(
		URLPatternReg.ReplaceAllString(text, "$2 ($1)"), "$1")
}

func RemoveTags(text string) string {
	return html.UnescapeString(
		UnclosedTagPatternReg.ReplaceAllString(
			TagPatternReg.ReplaceAllString(text, "$1"), ""))
}

func SnakeToCamelCase(text string, firstUpper bool) string {
	nextUpper := firstUpper
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

func ConvertType(text string, isOptional bool) string {
	switch text {
	case "String":
		return "string"
	case "Integer":
		return "int"
	case "Float number", "Float":
		return "float64"
	case "Boolean", "True":
		return "bool"
	default:
		if strings.HasPrefix(text, "Array of ") {
			return "[]" + ConvertType(strings.ReplaceAll(text, "Array of ", ""), false)
		}

		if isOptional {
			return "*" + text
		}
		return text
	}
}

func FitLine(text string, lineLen int) []string {
	words := strings.Split(text, " ")
	result := make([]string, 0)
	line := strings.Builder{}
	for _, word := range words {
		if line.Len()+len(word)+1 > lineLen {
			result = append(result, line.String())
			line.Reset()
		}
		line.WriteString(word + " ")
	}

	if line.Len() != 0 {
		result = append(result, line.String())
	}
	return result
}
