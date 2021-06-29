package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/mymmrac/go-telegram-bot-api/logger"
)

const packageName = "telego"
const docsURL = "https://core.telegram.org/bots/api"

const typePattern = `
<h4><a class="anchor" name="\w+?" href="#\w+?"><i class="anchor-icon"></i></a>([A-Z]\w+?)</h4>
<p>(.+?)</p>
.*?
(?:
<table class="table">
<thead>
<tr>
<th>Field</th>
<th>Type</th>
<th>Description</th>
</tr>
</thead>
<tbody>
(.*?)
</tbody>
</table>
|)
`

const fieldPattern = `
<tr>
<td>(\w+)</td>
<td>(.+?)</td>
<td>(.+?)</td>
</tr>
`

const urlPattern = `<a.*?href="(.+?)".*?>(.*?)</a>`

const tagPattern = `<.+?>(.+?)</.+?>`

var urlPatternReg *regexp.Regexp
var tagPatternReg *regexp.Regexp

func main() {
	log := logger.CreateLogrusLogger(logrus.ErrorLevel)

	typePatternReg := regexp.MustCompile(removeNewline(typePattern))
	fieldPatternReg := regexp.MustCompile(removeNewline(fieldPattern))
	urlPatternReg = regexp.MustCompile(urlPattern)
	tagPatternReg = regexp.MustCompile(tagPattern)

	file, err := os.Create("types.go")
	if err != nil {
		log.Error(err)
		return
	}

	response, err := http.Get(docsURL)
	if err != nil {
		log.Error(err)
		return
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Fprintf(file, "package %s\n\n", packageName)

	body := removeNewline(string(bodyBytes))

	typeMatch := typePatternReg.FindAllStringSubmatch(body, -1)

	for _, typeMatched := range typeMatch {
		typeName := typeMatched[1]
		typeDescription := removeTags(cleanDescription(typeMatched[2]))

		fmt.Fprintf(file, "// %s - %s\ntype %s struct {\n", typeName, typeDescription, typeName)

		typeDefinitionTable := typeMatched[3]
		fieldMatch := fieldPatternReg.FindAllStringSubmatch(typeDefinitionTable, -1)

		for _, fieldMatched := range fieldMatch {
			fieldName := fieldMatched[1]
			fieldType := convertType(removeTags(fieldMatched[2]))
			fieldDescription := removeTags(cleanDescription(fieldMatched[3]))
			omitemptyStr := ""
			if strings.HasPrefix(fieldDescription, "Optional.") {
				omitemptyStr = ",omitempty"
			}

			fmt.Fprintf(file, "\t// %s - %s\n\t%s %s `json:\"%s%s\"`\n\n",
				snakeToCamelCase(fieldName), fieldDescription, snakeToCamelCase(fieldName), fieldType, fieldName, omitemptyStr)
		}

		fmt.Fprintf(file, "}\n\n")
	}
}

func removeNewline(text string) string {
	return strings.ReplaceAll(text, "\n", "")
}

func cleanDescription(text string) string {
	return html.UnescapeString(urlPatternReg.ReplaceAllString(text, "$2 ($1)"))
}

func removeTags(text string) string {
	return tagPatternReg.ReplaceAllString(text, "$1")
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

func convertType(text string) string {
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
			return "[]" + convertType(strings.ReplaceAll(text, "Array of ", ""))
		}

		return "*" + text
	}
}
