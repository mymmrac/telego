package main

import (
	"fmt"
	"html"
	"os"
	"regexp"
	"strings"
)

const (
	linkPattern = `<a.+?href="(.+?)".*?>(.+?)<\/a>`

	externalURLPattern = `--(http[s]:\/\/.+?)--`
	internalURLPattern = `--(\/.+?)--`
	anchorURLPattern   = `--(#.+?)--`

	imagePattern = `<img.+?alt="(.+?)".*?>`

	tagNlPattern = `<(?:p|div|li|blockquote|br).*?>`
	tagPattern   = `<.+?>`

	tagElemPattern = `<.+?>(.+?)<\/.+?>`

	multiSpacePattern = `(\s)\s+`
)

var (
	linkRegexp *regexp.Regexp

	externalURLRegexp *regexp.Regexp
	internalURLRegexp *regexp.Regexp
	anchorURLRegexp   *regexp.Regexp

	imageRegexp *regexp.Regexp

	tagRegexp   *regexp.Regexp
	tagNlRegexp *regexp.Regexp

	tagElemRegexp *regexp.Regexp

	multiSpaceRegexp *regexp.Regexp
)

func init() {
	var err error
	linkRegexp, err = regexp.Compile(linkPattern)
	exitOnErr(err)

	externalURLRegexp, err = regexp.Compile(externalURLPattern)
	exitOnErr(err)

	internalURLRegexp, err = regexp.Compile(internalURLPattern)
	exitOnErr(err)

	anchorURLRegexp, err = regexp.Compile(anchorURLPattern)
	exitOnErr(err)

	imageRegexp, err = regexp.Compile(imagePattern)
	exitOnErr(err)

	tagNlRegexp, err = regexp.Compile(tagNlPattern)
	exitOnErr(err)

	tagRegexp, err = regexp.Compile(tagPattern)
	exitOnErr(err)

	tagElemRegexp, err = regexp.Compile(tagElemPattern)
	exitOnErr(err)

	multiSpaceRegexp, err = regexp.Compile(multiSpacePattern)
	exitOnErr(err)
}

func logInfo(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

func logError(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}

func exitOnErr(err error) {
	if err != nil {
		logError("%v", err)
		os.Exit(1)
	}
}

func join(lines []string) string {
	return strings.Join(lines, "")
}

func forEach(lines []string, op func(text string) string) []string {
	var newLines []string
	for _, line := range lines {
		newLines = append(newLines, op(line))
	}
	return newLines
}

func splitNl(text string) []string {
	return strings.Split(text, "\n")
}

func removeNl(text string) string {
	return strings.ReplaceAll(text, "\n", "")
}

func replaceHTML(text string) string {
	text = imageRegexp.ReplaceAllString(text, "$1")

	text = linkRegexp.ReplaceAllString(text, "$2 --$1--")

	text = externalURLRegexp.ReplaceAllString(text, "($1)")
	text = internalURLRegexp.ReplaceAllString(text, fmt.Sprintf("(%s$1)", baseURL))
	text = anchorURLRegexp.ReplaceAllString(text, fmt.Sprintf("(%s$1)", docsURL))

	text = tagNlRegexp.ReplaceAllString(text, "\n")
	text = tagRegexp.ReplaceAllString(text, "")

	text = html.UnescapeString(text)
	text = trimSpaces(text)

	return text
}

func removeHTML(text string) string {
	text = tagElemRegexp.ReplaceAllString(text, "$1")
	text = tagRegexp.ReplaceAllString(text, "")

	text = html.UnescapeString(text)

	return text
}

func splitTextToFitLine(text string) []string {
	words := strings.Split(text, " ")
	result := make([]string, 0)
	line := strings.Builder{}
	for _, word := range words {
		if strings.Contains(word, "\n") {
			ws := strings.Split(word, "\n")
			if len(ws) != 2 {
				os.Exit(2)
			}

			if line.Len()+len(ws[0])+1 > maxLineLen {
				result = append(result, line.String())
				line.Reset()
			}

			line.WriteString(ws[0] + " ")
			result = append(result, line.String())
			line.Reset()

			word = ws[1]
		}

		if line.Len()+len(word)+1 > maxLineLen {
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

func fitTextToLine(text, delimiter string) string {
	lines := splitTextToFitLine(delimiter + text)
	return strings.Join(lines, "\n"+delimiter)
}

func trimSpaces(text string) string {
	text = strings.TrimSpace(text)
	text = multiSpaceRegexp.ReplaceAllString(text, "$1")

	return text
}

func preparePattern(pattern string) string {
	return join(forEach(splitNl(pattern), strings.TrimSpace))
}

func snakeToCamelCase(text string) string {
	nextUpper := true
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

func parseType(text string, optional bool) string {
	text = removeHTML(text)

	switch text {
	case "String":
		return "string"
	case "Integer", "Int":
		return "int"
	case "Float number", "Float":
		return "float64"
	case "Boolean", "True":
		return "bool"
	case "Integer or String":
		return "ChatID"
	case "InputFile or String":
		if optional {
			return "*InputFile"
		}
		return "InputFile"
	default:
		if strings.HasPrefix(text, "Array of ") {
			text = strings.Replace(text, "Array of ", "", 1)
			return "[]" + parseType(text, false)
		}

		if optional {
			return "*" + text
		}
		return text
	}
}

func uppercaseWords(text string) string {
	text = strings.ReplaceAll(text, "Id ", "ID ")
	text = strings.ReplaceAll(text, "Id\n", "ID\n")
	text = strings.ReplaceAll(text, "Id)", "ID)")
	text = strings.ReplaceAll(text, "Ids", "IDs")
	text = strings.ReplaceAll(text, "Id,", "ID,")
	text = strings.ReplaceAll(text, "Id{", "ID{")
	text = strings.ReplaceAll(text, " id ", " ID ")

	text = strings.ReplaceAll(text, "Url ", "URL ")
	text = strings.ReplaceAll(text, " url ", " URL ")
	text = strings.ReplaceAll(text, " url's ", " URL's ")

	text = strings.ReplaceAll(text, "IpAddress", "IPAddress")

	return text
}

func firstToLower(text string) string {
	switch len(text) {
	case 0:
		return text
	case 1:
		return string(text[0] | ('a' - 'A'))
	default:
		return string(text[0]|('a'-'A')) + text[1:]
	}
}
