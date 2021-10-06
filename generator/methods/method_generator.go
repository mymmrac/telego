package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mymmrac/telego/generator"
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

//nolint:funlen,gocognit
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
	returnValuesCount := 0

	fmt.Println("Method count:", len(allMethods))

	data := strings.Builder{}

	_, _ = data.WriteString(fmt.Sprintf("package %s\n\n", generator.PackageName))

	for _, currentMethod := range allMethods {
		methodName := currentMethod[1]
		funcName := strings.Title(methodName)

		paramsStructName := funcName + "Params"

		methodDescription := generator.RemoveTags(generator.CleanDescription(currentMethod[2]))
		methodDescriptionWithoutTags := generator.RemoveTags(currentMethod[2])

		returns := ""

		returnsAfter := regexp.MustCompile(`[Rr]eturns [a-z ]*?((?:Array of |)[A-Z]\w+)`).
			FindStringSubmatch(methodDescriptionWithoutTags)
		if len(returnsAfter) != 0 {
			returns = returnsAfter[1]
			returnValuesCount++
		}

		returnsBefore := regexp.MustCompile(`((?:Array of |)[A-Z]\w+)[a-z ]*?returned`).
			FindStringSubmatch(methodDescriptionWithoutTags)
		if len(returnsBefore) != 0 {
			returns = returnsBefore[1]
			returnValuesCount++
		}

		returnType := ""
		switch returns {
		case "", "True", "error":
		//	Do noting
		default:
			returnType = generator.ConvertType(returns, true)
		}

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

		returnFunc := "error"
		returnVar := ""
		returnVarName := "nil"
		returnEnd := ""
		returnNil := ""
		if returnType != "" {
			returnFunc = fmt.Sprintf("(%s, error)", returnType)

			returnVarName = returnType[1:]
			if strings.HasPrefix(returnVarName, "]") {
				returnVarName = returnVarName[1:] + "s"
			}
			returnVarName = string(returnVarName[0]|('a'-'A')) + returnVarName[1:]

			returnVar = fmt.Sprintf("\n\tvar %s %s", returnVarName, returnType)
			returnEnd = returnVarName + ", "

			returnNil = "nil, "

			returnVarName = "&" + returnVarName
		}

		_, _ = data.WriteString(fmt.Sprintf(`
%s
func (b *Bot) %s(%s) %s {%s
	err := b.performRequest("%s", %s, %s)
	if err != nil {
		return %sfmt.Errorf("%s(): %%w", err)
	}

	return %snil
}

`,
			funcDescription, funcName, params, returnFunc, returnVar, methodName, paramsOrNil, returnVarName,
			returnNil, methodName, returnEnd))
	}

	dataString := data.String()
	dataString = generator.UppercaseWords(dataString)
	_, _ = file.WriteString(dataString)

	fmt.Println("Return values:", returnValuesCount)
}
