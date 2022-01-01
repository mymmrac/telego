package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type tgTypeField struct {
	name          string
	nameSnakeCase string
	typ           string
	description   string
	optional      bool
}

type tgTypeFields []tgTypeField

type tgType struct {
	name        string
	description string
	fields      tgTypeFields
}

type tgTypes []tgType

const typePattern = `
	<a class="anchor" name="\w+?" href="#\w+?">
		<i class="anchor-icon"></i>
	</a>
	([A-Z]\w+?)
</h4>

(.+?)

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
		(.+?)
	</tbody>
</table>
|
<h4>)
`

const typeFieldPattern = `
<tr>
	<td>(.+?)<\/td>
	<td>(.+?)<\/td>
	<td>(.+?)<\/td>
</tr>
`

var (
	typeRegexp      *regexp.Regexp
	typeFieldRegexp *regexp.Regexp
)

func init() {
	var err error
	typeRegexp, err = regexp.Compile(preparePattern(typePattern))
	exitOnErr(err)

	typeFieldRegexp, err = regexp.Compile(preparePattern(typeFieldPattern))
	exitOnErr(err)
}

func generateTypes(docs string) tgTypes {
	var types tgTypes

	typeGroups := typeRegexp.FindAllStringSubmatch(docs, -1)
	for _, typeGroup := range typeGroups {
		typ := tgType{
			name:        typeGroup[1],
			description: replaceHTML(typeGroup[2]),
			fields:      generateTypeFields(typeGroup[3]),
		}

		types = append(types, typ)
	}

	return types
}

func generateTypeFields(fieldDocs string) tgTypeFields {
	var fields tgTypeFields

	fieldGroups := typeFieldRegexp.FindAllStringSubmatch(fieldDocs, -1)
	for _, fieldGroup := range fieldGroups {
		field := tgTypeField{
			name:          snakeToCamelCase(fieldGroup[1]),
			nameSnakeCase: fieldGroup[1],
			description:   replaceHTML(fieldGroup[3]),
		}

		if strings.HasPrefix(field.description, optionalPrefix) {
			field.optional = true
		}

		field.typ = parseType(fieldGroup[2], field.optional)

		fieldSpecialCases(&field)

		fields = append(fields, field)
	}

	return fields
}

func writeTypes(file *os.File, types tgTypes) {
	data := strings.Builder{}

	logInfo("Types: %d", len(types))

	data.WriteString(fmt.Sprintf("package %s\n", packageName))
	data.WriteString(`
import (
	"errors"
	"fmt"

	"github.com/mymmrac/telego/telegoapi"
)
`)

	fieldsCount := 0
	for _, t := range types {
		typeDescriptionLines := fitLine(fmt.Sprintf("// %s - %s", t.name, t.description))
		data.WriteString(strings.Join(typeDescriptionLines, "\n// "))

		data.WriteString(fmt.Sprintf("\ntype %s struct {\n", t.name))

		fieldsCount += len(t.fields)
		for _, f := range t.fields {
			fieldDescriptionLines := fitLine(fmt.Sprintf("\t// %s - %s", f.name, f.description))
			data.WriteString(strings.Join(fieldDescriptionLines, "\n\t// "))

			omitempty := ""
			if f.optional {
				omitempty = omitemptySuffix
			}

			data.WriteString(fmt.Sprintf("\n\t%s %s `json:\"%s%s\"`\n\n", f.name, f.typ, f.nameSnakeCase, omitempty))
		}

		data.WriteString("}\n\n")
	}

	logInfo("Type fields: %d", fieldsCount)

	_, err := file.WriteString(uppercaseWords(data.String()))
	exitOnErr(err)
}

func fieldSpecialCases(field *tgTypeField) {
	if strings.Contains(field.name, "Date") && field.typ == "int" {
		field.typ = "int64"
	}

	if (strings.Contains(field.description, "64-bit integer") ||
		strings.Contains(field.description, "64 bit integer")) && field.typ == "int" {
		field.typ = "int64"
	}

	if field.name == "UserId" && field.typ == "int" {
		field.typ = "int64"
	}
}
