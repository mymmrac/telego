package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/mymmrac/go-telegram-bot-api/generator"
)

const methodPattern = `
<h4><a class="anchor" name="\w+?" href="#\w+?"><i class="anchor-icon"></i></a>([a-z]\w+?)</h4>
<p>(.+?)</p>
(?:<blockquote>.+?</blockquote>|)
(?:<p>.+?</p>|)
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

//nolint:funlen
func main() {
	methodPatternReg := regexp.MustCompile(generator.RemoveNewline(methodPattern))
	paramsPatternReg := regexp.MustCompile(generator.RemoveNewline(paramsPattern))

	file, err := os.Create("methods_generated.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := http.Get(generator.DocsURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := generator.RemoveNewline(string(bodyBytes))
	methodMatch := methodPatternReg.FindAllStringSubmatch(body, -1)

	fmt.Fprintf(file, "package %s\n\n", generator.PackageName)

	for _, methodMatched := range methodMatch {
		methodName := methodMatched[1]
		funcName := strings.Title(methodName)
		paramsStructName := funcName + "Params"

		methodDescription := generator.RemoveTags(generator.CleanDescription(methodMatched[2]))
		funcDescriptionLines := generator.FitLine(fmt.Sprintf("// %s - %s",
			funcName, methodDescription), generator.MaxLineLen)
		funcDescription := strings.Join(funcDescriptionLines, "\n// ")

		paramsDefinitionTable := methodMatched[3]
		paramsMatch := paramsPatternReg.FindAllStringSubmatch(paramsDefinitionTable, -1)

		params := ""
		paramsOrNil := "nil"

		if len(paramsMatch) != 0 {
			fmt.Fprintf(file, "// %s - Represents parameters of %s method.\ntype %s struct {\n",
				paramsStructName, methodName, paramsStructName)

			for _, paramMatched := range paramsMatch {
				paramName := paramMatched[1]
				fieldName := generator.SnakeToCamelCase(paramName, true)

				isOptional := paramMatched[3] == "Optional"
				omitemptyStr := ""
				optionalStr := ""
				if isOptional {
					omitemptyStr = ",omitempty"
					optionalStr = "Optional. "
				}

				paramDescription := generator.RemoveTags(generator.CleanDescription(paramMatched[4]))
				filedDescriptionLines := generator.FitLine(fmt.Sprintf("// %s - %s%s",
					fieldName, optionalStr, paramDescription), generator.MaxLineLen)
				fieldDescription := strings.Join(filedDescriptionLines, "\n\t// ")

				fieldType := generator.ConvertType(generator.RemoveTags(paramMatched[2]), isOptional)

				fmt.Fprintf(file, "\t%s\n\t%s %s `json:\"%s%s\"`\n\n",
					fieldDescription, fieldName, fieldType, paramName, omitemptyStr)
			}

			fmt.Fprintf(file, "}\n")

			params = fmt.Sprintf("params *%s", paramsStructName)
			paramsOrNil = "params"
		}

		fmt.Fprintf(file, `
%s
func (b *Bot) %s(%s) error {
	err := b.performRequest("%s", %s, nil)
	if err != nil {
		return fmt.Errorf("%s(): %%w", err)
	}

	return nil
}

`,
			funcDescription, funcName, params, methodName, paramsOrNil, methodName)
	}
}
