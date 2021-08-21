package generator

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const (
	baseURL = "https://core.telegram.org"
	docsURL = baseURL + "/bots/api"

	PackageName = "telego"

	MaxLineLen = 110

	OmitemptySuffix = ",omitempty"
)

const (
	urlPattern = `<a.*?href="(.+?)".*?>(.*?)</a>`
	imgPattern = `<img.*?alt="(.+?)".*?>`

	externalLinkPattern  = `--(http[s]:\/\/.+?)--`
	linkOnPagePattern    = `--(#.+?)--`
	linkNotOnPagePattern = `--(\/.+?)--`

	tagPattern         = `<.+?>(.+?)</.+?>`
	unclosedTagPattern = `<.+?>`
)

var (
	urlReg = regexp.MustCompile(urlPattern)
	imgReg = regexp.MustCompile(imgPattern)

	externalLinkReg  = regexp.MustCompile(externalLinkPattern)
	linkOnPageReg    = regexp.MustCompile(linkOnPagePattern)
	linkNotOnPageReg = regexp.MustCompile(linkNotOnPagePattern)

	tagReg         = regexp.MustCompile(tagPattern)
	unclosedTagReg = regexp.MustCompile(unclosedTagPattern)
)

func RemoveNewline(text string) string {
	return strings.ReplaceAll(text, "\n", "")
}

func CleanDescription(text string) string {
	return linkNotOnPageReg.ReplaceAllString(
		linkOnPageReg.ReplaceAllString(
			externalLinkReg.ReplaceAllString(
				imgReg.ReplaceAllString(
					urlReg.ReplaceAllString(
						text, "$2 --$1--"),
					"$1"),
				"($1)"),
			fmt.Sprintf("(%s$1)", docsURL)),
		fmt.Sprintf("(%s$1)", baseURL))
}

func RemoveTags(text string) string {
	return html.UnescapeString(
		unclosedTagReg.ReplaceAllString(
			tagReg.ReplaceAllString(text, "$1"), ""))
}

func SnakeToCamelCase(text string, firstUpper bool) string {
	nextUpper := firstUpper
	result := strings.Builder{}
	result.Grow(len(text))
	for _, currentChar := range []byte(text) {
		if currentChar == '_' {
			nextUpper = true
			continue
		}

		if nextUpper {
			nextUpper = false
			currentChar += 'A'
			currentChar -= 'a'
			result.WriteByte(currentChar)
		} else {
			result.WriteByte(currentChar)
		}
	}

	return result.String()
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
	case "Integer or String":
		return "ChatID"
	default:
		if strings.HasPrefix(text, "Array of ") {
			return "[]" + ConvertType(strings.Replace(text, "Array of ", "", 1), false)
		}

		if isOptional {
			return "*" + text
		}
		return text
	}
}

func FitLine(text string, maxLineLength int) []string {
	words := strings.Split(text, " ")
	result := make([]string, 0)
	line := strings.Builder{}
	for _, word := range words {
		if line.Len()+len(word)+1 > maxLineLength {
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

func GetDocsText() (string, error) {
	response, err := http.Get(docsURL)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	body := RemoveNewline(string(bodyBytes))
	return body, nil
}
