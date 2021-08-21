package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mymmrac/go-telegram-bot-api/generator"
)

const typePattern = `
<h4><a class="anchor" name="\w+?" href="#\w+?"><i class="anchor-icon"></i></a>([A-Z]\w+?)</h4>
<p>(.+?)</p>
(?:<div.*?>.+?</div>|)
(?:<p>.+?</p>|)
(?:<blockquote>.+?</blockquote>|)
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

	file, err := os.Create("types.go.generated")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := generator.GetDocsText()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = fmt.Fprintf(file, "package %s\n\n", generator.PackageName)

	allTypes := typePatternReg.FindAllStringSubmatch(body, -1)

	for _, currentType := range allTypes {
		typeName := currentType[1]

		typeDescription := generator.RemoveTags(generator.CleanDescription(currentType[2]))
		typeDescriptionLines := generator.FitLine(fmt.Sprintf("// %s - %s",
			typeName, typeDescription), generator.MaxLineLen)
		typeDescriptionFitted := strings.Join(typeDescriptionLines, "\n// ")

		_, _ = fmt.Fprintf(file, "%s\ntype %s struct {\n", typeDescriptionFitted, typeName)

		typeDefinitionTable := currentType[3]
		allFields := fieldPatternReg.FindAllStringSubmatch(typeDefinitionTable, -1)

		for _, currentFiled := range allFields {
			fieldName := currentFiled[1]
			fieldNameCamelCase := generator.SnakeToCamelCase(fieldName, true)

			fieldDescription := generator.RemoveTags(generator.CleanDescription(currentFiled[3]))
			fieldDescriptionLines := generator.FitLine(fmt.Sprintf("\t// %s - %s",
				fieldNameCamelCase, fieldDescription), generator.MaxLineLen)
			fieldDescriptionFitted := strings.Join(fieldDescriptionLines, "\n\t// ")

			isOptional := strings.HasPrefix(fieldDescription, "Optional.")
			omitempty := ""
			if isOptional {
				omitempty = generator.OmitemptySuffix
			}

			fieldType := generator.ConvertType(generator.RemoveTags(currentFiled[2]), isOptional)

			_, _ = fmt.Fprintf(file, "%s\n\t%s %s `json:\"%s%s\"`\n\n",
				fieldDescriptionFitted, fieldNameCamelCase, fieldType, fieldName, omitempty)
		}

		_, _ = fmt.Fprintf(file, "}\n\n")
	}
}
