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
	typePatternReg := regexp.MustCompile(generator.RemoveNewline(typePattern))
	fieldPatternReg := regexp.MustCompile(generator.RemoveNewline(fieldPattern))

	file, err := os.Create("types.go")
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

	fmt.Fprintf(file, "package %s\n\n", generator.PackageName)

	body := generator.RemoveNewline(string(bodyBytes))

	typeMatch := typePatternReg.FindAllStringSubmatch(body, -1)

	for _, typeMatched := range typeMatch {
		typeName := typeMatched[1]
		typeDescription := generator.RemoveTags(generator.CleanDescription(typeMatched[2]))

		fmt.Fprintf(file, "// %s - %s\ntype %s struct {\n", typeName, typeDescription, typeName)

		typeDefinitionTable := typeMatched[3]
		fieldMatch := fieldPatternReg.FindAllStringSubmatch(typeDefinitionTable, -1)

		for _, fieldMatched := range fieldMatch {
			fieldName := fieldMatched[1]
			fieldDescription := generator.RemoveTags(generator.CleanDescription(fieldMatched[3]))
			isOptional := strings.HasPrefix(fieldDescription, "Optional.")
			omitemptyStr := ""
			if isOptional {
				omitemptyStr = ",omitempty"
			}
			fieldType := generator.ConvertType(generator.RemoveTags(fieldMatched[2]), isOptional)

			fmt.Fprintf(file, "\t// %s - %s\n\t%s %s `json:\"%s%s\"`\n\n",
				generator.SnakeToCamelCase(fieldName, true), fieldDescription,
				generator.SnakeToCamelCase(fieldName, true), fieldType, fieldName, omitemptyStr)
		}

		fmt.Fprintf(file, "}\n\n")
	}
}
