package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/mymmrac/go-telegram-bot-api/logger"
)

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

func main() {
	log := logger.CreateLogrusLogger(logrus.ErrorLevel)

	typePatternReg := regexp.MustCompile(removeNewline(typePattern))
	fieldPatternReg := regexp.MustCompile(removeNewline(fieldPattern))

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
			return
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
			fieldDescription := removeTags(cleanDescription(fieldMatched[3]))
			isOptional := strings.HasPrefix(fieldDescription, "Optional.")
			omitemptyStr := ""
			if isOptional {
				omitemptyStr = ",omitempty"
			}
			fieldType := convertType(removeTags(fieldMatched[2]), isOptional)

			fmt.Fprintf(file, "\t// %s - %s\n\t%s %s `json:\"%s%s\"`\n\n",
				snakeToCamelCase(fieldName), fieldDescription, snakeToCamelCase(fieldName), fieldType, fieldName, omitemptyStr)
		}

		fmt.Fprintf(file, "}\n\n")
	}
}

func convertType(text string, isOptional bool) string {
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
			return "[]" + convertType(strings.ReplaceAll(text, "Array of ", ""), false)
		}

		if isOptional {
			return "*" + text
		}
		return text
	}
}
