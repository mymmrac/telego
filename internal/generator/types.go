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

const curTypePattern = `^type (\w+) (?:struct|interface) {`
const curConstPattern = `^const \(`
const curFuncPattern = `^func \(`
const curInterfacePattern = `^type \w+ interface {`

var (
	typeRegexp        = regexp.MustCompile(preparePattern(typePattern))
	typeFieldRegexp   = regexp.MustCompile(preparePattern(typeFieldPattern))
	curTypeRegexp     = regexp.MustCompile(curTypePattern)
	curConstRegexp    = regexp.MustCompile(curConstPattern)
	curFuncRegexp     = regexp.MustCompile(curFuncPattern)
	curInterfaceRegex = regexp.MustCompile(curInterfacePattern)
)

func generateTypes(docs string) tgTypes {
	typeGroups := typeRegexp.FindAllStringSubmatch(docs, -1)
	types := make(tgTypes, len(typeGroups))

	for i, typeGroup := range typeGroups {
		types[i] = tgType{
			name:        typeGroup[1],
			description: replaceHTML(typeGroup[2]),
			fields:      generateTypeFields(typeGroup[3], typeGroup[1]),
		}
	}

	return types
}

func generateTypeFields(fieldDocs, typeName string) tgTypeFields {
	if fieldDocs == "" {
		return nil
	}

	fieldGroups := typeFieldRegexp.FindAllStringSubmatch(fieldDocs, -1)
	fields := make(tgTypeFields, len(fieldGroups))

	for i, fieldGroup := range fieldGroups {
		field := tgTypeField{
			name:          snakeToCamelCase(fieldGroup[1]),
			nameSnakeCase: fieldGroup[1],
			description:   replaceHTML(fieldGroup[3]),
		}

		if strings.HasPrefix(field.description, optionalPrefix) {
			field.optional = true
		}

		field.typ = parseType(fieldGroup[2], field.optional)
		fieldSpecialCases(&field, typeName)

		fields[i] = field
	}

	return fields
}

func parseCurrentTypes(types string) map[string][]string {
	additional := map[string][]string{}

	constCount := 0
	funcOrInterfaceCount := 0
	currentType := ""

	lines := strings.Split(types, "\n")
	for i, line := range lines {
		typeMatches := curTypeRegexp.FindStringSubmatch(line)
		if len(typeMatches) > 0 {
			currentType = typeMatches[1]
			continue
		}

		if currentType != "" && curConstRegexp.MatchString(line) {
			end := i + 1
			for ; ; end++ {
				if end >= len(lines) {
					end = -1
					break
				}

				if lines[end] == ")" {
					break
				}
			}

			if end == -1 {
				continue
			}

			_, ok := additional[currentType]
			if !ok {
				additional[currentType] = []string{}
			}
			additional[currentType] = append(additional[currentType], strings.Join(lines[i-1:end+1], "\n"))
			constCount++

			continue
		}

		if currentType != "" && (curFuncRegexp.MatchString(line) || curInterfaceRegex.MatchString(line)) {
			start := i - 1
			for ; ; start-- {
				if start < 0 {
					start = -1
					break
				}

				if !strings.HasPrefix(lines[start], "//") {
					break
				}
			}

			end := i + 1
			for ; ; end++ {
				if end >= len(lines) {
					end = -1
					break
				}

				if lines[end] == "}" {
					break
				}
			}

			if start == -1 || end == -1 {
				continue
			}

			_, ok := additional[currentType]
			if !ok {
				additional[currentType] = []string{}
			}
			additional[currentType] = append(additional[currentType], strings.Join(lines[start:end+1], "\n"))
			funcOrInterfaceCount++
		}
	}

	logInfo("Const count: %d", constCount)
	logInfo("Func & interface count: %d", funcOrInterfaceCount)

	return additional
}

func writeTypes(file *os.File, types tgTypes, currentTypes string) {
	additional := parseCurrentTypes(currentTypes)

	data := strings.Builder{}

	logInfo("Types: %d", len(types))

	data.WriteString(`package telego

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/chococola/telego/internal/json"
	"github.com/chococola/telego/telegoapi"
)
`)

	fieldsCount := 0
	for _, t := range types {
		typeDescription := fitTextToLine(fmt.Sprintf("%s - %s", t.name, t.description), "// ")
		data.WriteString(typeDescription)

		if len(t.fields) == 0 && !strings.Contains(t.description, "holds no information") {
			data.WriteString(fmt.Sprintf(
				"\ntype %s interface {\n\t// TODO: Add methods\n\t// Disallow external implementations\n\ti%s()\n",
				t.name, t.name,
			))
		} else {
			data.WriteString(fmt.Sprintf("\ntype %s struct {", t.name))
		}

		if len(t.fields) > 0 {
			data.WriteString("\n")
		}

		fieldsCount += len(t.fields)
		for _, f := range t.fields {
			fieldDescription := fitTextToLine(fmt.Sprintf("%s - %s", f.name, f.description), "\t// ")
			data.WriteString(fieldDescription)

			omitempty := ""
			if f.optional {
				omitempty = omitemptySuffix
			}

			data.WriteString(fmt.Sprintf("\n\t%s %s `json:\"%s%s\"`\n\n", f.name, f.typ, f.nameSnakeCase, omitempty))
		}

		data.WriteString("}\n\n")

		additions := additional[t.name]
		for _, a := range additions {
			data.WriteString(a)
			data.WriteString("\n")
		}
	}

	logInfo("Type fields: %d", fieldsCount)

	_, err := file.WriteString(uppercaseWords(data.String()))
	exitOnErr(err)
}

func fieldSpecialCases(field *tgTypeField, typeName string) {
	if strings.Contains(field.name, "Date") && field.typ == "int" {
		field.typ = "int64"
	}

	if (strings.Contains(field.description, "64-bit integer") ||
		strings.Contains(field.description, "64 bit integer")) && field.typ == "int" {
		field.typ = "int64"
	}

	if strings.Contains(field.description, "32-bit identifier") && field.typ == "int" {
		field.typ = "int32"
	}

	if field.name == "UserId" && field.typ == "int" {
		field.typ = "int64"
	}
	if field.name == "UserIds" && field.typ == "[]int" {
		field.typ = "[]int64"
	}

	if field.name == "Media" && field.typ == "string" {
		field.typ = "InputFile"
	}

	if field.name == "InputMessageContent" && field.typ == "*InputMessageContent" {
		field.typ = "InputMessageContent"
	}

	if typeName == "ChatPermissions" ||
		((typeName == "KeyboardButtonRequestUsers" || typeName == "KeyboardButtonRequestChat") &&
			field.typ == "bool" && field.optional) {
		field.typ = "*bool"
	}

	if typeName == "InlineKeyboardButton" && field.typ == "string" &&
		(field.name == "SwitchInlineQuery" || field.name == "SwitchInlineQueryCurrentChat") {
		field.typ = "*string"
	}

	if field.typ == "*MaybeInaccessibleMessage" {
		field.typ = "MaybeInaccessibleMessage"
	}

	if field.typ == "*MessageOrigin" {
		field.typ = "MessageOrigin"
	}

	if field.typ == "*ReactionType" {
		field.typ = "ReactionType"
	}

	if field.typ == "*ChatBoostSource" {
		field.typ = "ChatBoostSource"
	}

	if field.typ == "*RevenueWithdrawalState" {
		field.typ = "RevenueWithdrawalState"
	}

	if field.typ == "*TransactionPartner" {
		field.typ = "TransactionPartner"
	}
}
