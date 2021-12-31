package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mymmrac/telego/internal/generator"
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

	data := strings.Builder{}

	_, _ = data.WriteString(fmt.Sprintf("package %s\n\n", generator.PackageName))

	allTypes := typePatternReg.FindAllStringSubmatch(body, -1)

	fmt.Println("Types count:", len(allTypes))

	fieldsCount := 0

	for _, currentType := range allTypes {
		typeName := currentType[1]

		typeDescription := generator.RemoveTags(generator.CleanDescription(currentType[2]))
		typeDescriptionLines := generator.FitLine(fmt.Sprintf("// %s - %s",
			typeName, typeDescription), generator.MaxLineLen)
		typeDescriptionFitted := strings.Join(typeDescriptionLines, "\n// ")

		_, _ = data.WriteString(fmt.Sprintf("%s\ntype %s struct {\n", typeDescriptionFitted, typeName))

		typeDefinitionTable := currentType[3]
		allFields := fieldPatternReg.FindAllStringSubmatch(typeDefinitionTable, -1)

		fieldsCount += len(allFields)

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

			_, _ = data.WriteString(fmt.Sprintf("%s\n\t%s %s `json:\"%s%s\"`\n\n",
				fieldDescriptionFitted, fieldNameCamelCase, fieldType, fieldName, omitempty))
		}

		_, _ = data.WriteString("}\n\n")
	}

	fmt.Println("Fields count:", fieldsCount)

	dataString := data.String()
	dataString = generator.UppercaseWords(dataString)
	_, _ = file.WriteString(dataString)
}
