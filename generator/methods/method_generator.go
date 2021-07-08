package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/mymmrac/go-telegram-bot-api/generator"
	"github.com/mymmrac/go-telegram-bot-api/logger"
)

const methodPattern = `
<h4><a class="anchor" name="\w+?" href="#\w+?"><i class="anchor-icon"></i></a>([a-z]\w+?)</h4>
<p>(.+?)</p>
.*?
(?:
<table class="table">
<thead>
<tr>
<th>Parameter</th>
<th>Type</th>
<th>Required</th>
<th>Description</th>
</tr>
</thead>
<tbody>
(.*?)
</tbody>
</table>
|)
`

const paramsPattern = `
<tr>
<td>(\w+)</td>
<td>(.+?)</td>
<td>(\w+)</td>
<td>(.+?)</td>
</tr>
`

func main() {
	log := logger.CreateLogrusLogger(logrus.ErrorLevel)

	methodPatternReg := regexp.MustCompile(generator.RemoveNewline(methodPattern))
	paramsPatternReg := regexp.MustCompile(generator.RemoveNewline(paramsPattern))

	file, err := os.Create("methods.go")
	if err != nil {
		log.Error(err)
		return
	}

	response, err := http.Get(generator.DocsURL)
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

	fmt.Fprintf(file, "package %s\n\n", generator.PackageName)

	body := generator.RemoveNewline(string(bodyBytes))

	methodMatch := methodPatternReg.FindAllStringSubmatch(body, -1)

	for _, methodMatched := range methodMatch {
		methodName := methodMatched[1]
		methodDescription := generator.RemoveTags(generator.CleanDescription(methodMatched[2]))

		allMethodDescription := strings.Builder{}
		allMethodDescription.WriteString(fmt.Sprintf("%s - %s{|}{|}Parameters:{|}", methodName, methodDescription))
		allMethodParams := strings.Builder{}

		paramsDefinitionTable := methodMatched[3]
		paramsMatch := paramsPatternReg.FindAllStringSubmatch(paramsDefinitionTable, -1)

		for i, paramMatched := range paramsMatch {
			paramName := generator.SnakeToCamelCase(paramMatched[1], false)
			paramOptional := paramMatched[3]
			isOptional := paramOptional == "Optional"
			paramDescription := generator.RemoveTags(generator.CleanDescription(paramMatched[4]))
			paramType := generator.ConvertType(generator.RemoveTags(paramMatched[2]), isOptional)

			optional := ""
			if isOptional {
				optional = " (optional)"
			}

			allMethodDescription.WriteString(fmt.Sprintf("%s%s - %s{|}{|}", paramName, optional, paramDescription))

			allMethodParams.WriteString(fmt.Sprintf("%s %s", paramName, paramType))
			if i != len(paramsMatch)-1 {
				allMethodParams.WriteString(", ")
			}
		}

		splitedDescription := strings.Builder{}
		lines := generator.FitLine(allMethodDescription.String(), generator.MaxLineLen)
		for _, line := range lines {
			splitedDescription.WriteString(fmt.Sprintf("// %s\n", line))
		}
		desc := strings.ReplaceAll(splitedDescription.String(), "{|}", "\n// ")

		fmt.Fprintf(file, "%sfunc %s(%s){\n\t\n}\n\n", desc, methodName, allMethodParams.String())
	}
}
