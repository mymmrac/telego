package main

import (
	"fmt"
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

	file, err := os.Create("methods.go.generated")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := generator.GetDocsText()
	if err != nil {
		fmt.Println(err)
		return
	}

	allMethods := methodPatternReg.FindAllStringSubmatch(body, -1)

	data := strings.Builder{}

	_, _ = data.WriteString(fmt.Sprintf("package %s\n\n", generator.PackageName))

	for _, currentMethod := range allMethods {
		methodName := currentMethod[1]
		funcName := strings.Title(methodName)

		paramsStructName := funcName + "Params"

		methodDescription := generator.RemoveTags(generator.CleanDescription(currentMethod[2]))
		funcDescriptionLines := generator.FitLine(fmt.Sprintf("// %s - %s",
			funcName, methodDescription), generator.MaxLineLen)
		funcDescription := strings.Join(funcDescriptionLines, "\n// ")

		paramsDefinitionTable := currentMethod[3]
		allParams := paramsPatternReg.FindAllStringSubmatch(paramsDefinitionTable, -1)

		params := ""
		paramsOrNil := "nil"

		if len(allParams) != 0 {
			_, _ = data.WriteString(fmt.Sprintf("// %s - Represents parameters of %s method.\ntype %s struct {\n",
				paramsStructName, methodName, paramsStructName))

			for _, currentParam := range allParams {
				paramName := currentParam[1]
				fieldName := generator.SnakeToCamelCase(paramName, true)

				isOptional := currentParam[3] == "Optional"
				omitempty := ""
				optional := ""
				if isOptional {
					omitempty = generator.OmitemptySuffix
					optional = "Optional. "
				}

				paramDescription := generator.RemoveTags(generator.CleanDescription(currentParam[4]))
				filedDescriptionLines := generator.FitLine(fmt.Sprintf("// %s - %s%s",
					fieldName, optional, paramDescription), generator.MaxLineLen)
				fieldDescription := strings.Join(filedDescriptionLines, "\n\t// ")

				fieldType := generator.ConvertType(generator.RemoveTags(currentParam[2]), isOptional)

				_, _ = data.WriteString(fmt.Sprintf("\t%s\n\t%s %s `json:\"%s%s\"`\n\n",
					fieldDescription, fieldName, fieldType, paramName, omitempty))
			}

			_, _ = data.WriteString("}\n")

			params = fmt.Sprintf("params *%s", paramsStructName)
			paramsOrNil = "params"
		}

		_, _ = data.WriteString(fmt.Sprintf(`
%s
func (b *Bot) %s(%s) error {
	err := b.performRequest("%s", %s, nil)
	if err != nil {
		return fmt.Errorf("%s(): %%w", err)
	}

	return nil
}

`,
			funcDescription, funcName, params, methodName, paramsOrNil, methodName))
	}

	dataString := data.String()
	dataString = generator.UppercaseWords(dataString)
	_, _ = file.WriteString(dataString)
}
