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

	methodsFile, err := os.Create("methods.go.generated")
	if err != nil {
		fmt.Println(err)
		return
	}

	testsFile, err := os.Create("methods_test.go.generated")
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

	methodsData := strings.Builder{}
	testsData := strings.Builder{}

	_, _ = methodsData.WriteString(fmt.Sprintf("package %s\n\n", generator.PackageName))
	_, _ = testsData.WriteString(fmt.Sprintf("package %s\n\n", generator.PackageName))

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
			_, _ = methodsData.WriteString(fmt.Sprintf("// %s - Represents parameters of %s method.\ntype %s struct {\n",
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

				_, _ = methodsData.WriteString(fmt.Sprintf("\t%s\n\t%s %s `json:\"%s%s\"`\n\n",
					fieldDescription, fieldName, fieldType, paramName, omitempty))
			}

			_, _ = methodsData.WriteString("}\n")

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

		expectedData := ""
		expectedAssert := ""
		expectedAssertNil := ""
		actualReturnVar := ""
		if returnVarName != "nil" {
			actualReturnVar = returnVarName[1:] + ", "
			expectedVarName := "expected" + strings.ToUpper(returnVarName[1:2]) + returnVarName[2:]
			expectedData = "\n\t\t" + expectedVarName + " := " + strings.Replace(returnType, "*", "&", 1) + "{}\n\t\tsetResult(t, " + expectedVarName + ")"
			expectedAssert = fmt.Sprintf("\n\t\tassert.Equal(t, %s, %s)", expectedVarName, returnVarName[1:])
			expectedAssertNil = fmt.Sprintf("\n\t\tassert.Nil(t, %s)", returnVarName[1:])
		}

		args := "nil"
		if paramsOrNil == "nil" {
			args = ""
		}

		_, _ = testsData.WriteString(fmt.Sprintf(`
func TestBot_%s(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)
		%s
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		%serr := m.Bot.%s(%s)
		assert.NoError(t, err)%s
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		%serr := m.Bot.%s(%s)
		assert.Error(t, err)%s
	})
}
`, funcName, expectedData, actualReturnVar, funcName, args, expectedAssert, actualReturnVar, funcName, args, expectedAssertNil))

		_, _ = methodsData.WriteString(fmt.Sprintf(`
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

	dataString := methodsData.String()
	dataString = generator.UppercaseWords(dataString)
	_, _ = methodsFile.WriteString(dataString)

	testDataString := testsData.String()
	testDataString = generator.UppercaseWords(testDataString)
	_, _ = testsFile.WriteString(testDataString)

	fmt.Println("Return values:", returnValuesCount)
}
