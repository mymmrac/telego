package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type tgMethodParameter struct {
	name          string
	nameSnakeCase string
	typ           string
	optional      bool
	description   string
}

type tgMethodParameters []tgMethodParameter

type tgMethod struct {
	name        string
	nameTitle   string
	description string
	parameters  tgMethodParameters
	returnType  string
}

func (m tgMethod) hasReturnValue() bool {
	return m.returnType != returnTypeNotFound && m.returnType != ""
}

type tgMethods []tgMethod

const methodPattern = `
	<a class="anchor" name="\w+?" href="#\w+?">
		<i class="anchor-icon"></i>
	</a>
	([a-z]\w+?)
</h4>

(.+?)

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
		(.+?)
	</tbody>
</table>
|
<h4>)
`

const methodParameterPattern = `
<tr>
	<td>(\w+)</td>
	<td>(.+?)</td>
	<td>(\w+)</td>
	<td>(.+?)</td>
</tr>
`

const (
	returnTypePattern1 = `[Rr]eturns [a-z ]*?((?:Array of |)[A-Z]\w+)`
	returnTypePattern2 = `((?:Array of |)[A-Z]\w+)[a-z ]*?returned`
)

const returnTypeNotFound = "NOT_FOUND"

const curInternalFuncPattern = `^func \(.+?\) [a-z]\w+\(`

var (
	methodRegexp          = regexp.MustCompile(preparePattern(methodPattern))
	methodParameterRegexp = regexp.MustCompile(preparePattern(methodParameterPattern))

	returnTypeRegexp1 = regexp.MustCompile(returnTypePattern1)
	returnTypeRegexp2 = regexp.MustCompile(returnTypePattern2)

	curInternalFuncRegexp = regexp.MustCompile(curInternalFuncPattern)
)

func generateMethods(docs string) tgMethods {
	methodGroups := methodRegexp.FindAllStringSubmatch(docs, -1)
	methods := make(tgMethods, len(methodGroups))

	cs := cases.Title(language.English, cases.NoLower)

	for i, methodGroup := range methodGroups {
		method := tgMethod{
			name:        methodGroup[1],
			nameTitle:   cs.String(methodGroup[1]),
			description: replaceHTML(methodGroup[2]),
			parameters:  generateMethodParameters(methodGroup[3], methodGroup[1]),
			returnType:  parseReturnType(methodGroup[2]),
		}

		methodSpecialCases(&method)

		methods[i] = method
	}

	return methods
}

func generateMethodParameters(parameterDocs, methodName string) tgMethodParameters {
	parameterGroups := methodParameterRegexp.FindAllStringSubmatch(parameterDocs, -1)
	parameters := make(tgMethodParameters, len(parameterGroups))

	for i, parameterGroup := range parameterGroups {
		parameter := tgMethodParameter{
			name:          snakeToCamelCase(parameterGroup[1]),
			nameSnakeCase: parameterGroup[1],
			description:   replaceHTML(parameterGroup[4]),
		}

		parameter.optional = parameterGroup[3] == "Optional"
		parameter.typ = parseType(parameterGroup[2], parameter.optional)

		parameterSpecialCases(&parameter, methodName)

		parameters[i] = parameter
	}

	return parameters
}

func parseCurrentMethods(methods string) map[string][]string {
	additional := map[string][]string{}

	constCount := 0
	internalFuncCount := 0
	currentType := ""

	lines := strings.Split(methods, "\n")
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

		if currentType != "" && curInternalFuncRegexp.MatchString(line) {
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

			if end == -1 {
				continue
			}

			_, ok := additional[currentType]
			if !ok {
				additional[currentType] = []string{}
			}
			additional[currentType] = append(additional[currentType], strings.Join(lines[i:end+1], "\n"))
			internalFuncCount++
		}
	}

	logInfof("Const count: %d", constCount)
	logInfof("Internal func count: %d", internalFuncCount)

	return additional
}

func writeMethods(file *os.File, methods tgMethods, currentMethods string) {
	additional := parseCurrentMethods(currentMethods)

	data := strings.Builder{}

	logInfof("Methods: %d", len(methods))

	data.WriteString(`package telego

import (
	"fmt"

	ta "github.com/mymmrac/telego/telegoapi"
)
`)

	parametersCount := 0
	returnsCount := 0
	returnsNotFoundCount := 0
	successValuesCount := 0

	for _, m := range methods {
		parametersStruct := m.nameTitle + "Params"
		parametersArg := ""

		if len(m.parameters) > 0 {
			parametersArg = fmt.Sprintf("params *%s", parametersStruct)

			parametersStructDescription := fitTextToLine(fmt.Sprintf("%s - Represents parameters of %s method.",
				parametersStruct, m.name), "// ")
			data.WriteString(parametersStructDescription)

			data.WriteString(fmt.Sprintf("\ntype %s struct {\n", parametersStruct))

			parametersCount += len(m.parameters)
			for _, p := range m.parameters {
				optional := ""
				if p.optional {
					optional = optionalPrefix
				}

				parameterDescription := fitTextToLine(fmt.Sprintf("%s - %s%s", p.name, optional, p.description), "\t// ")
				data.WriteString(parameterDescription)

				omitempty := ""
				if p.optional {
					omitempty = omitemptySuffix
				}

				if strings.Contains(p.typ, " or ") || strings.Contains(p.typ, ",") {
					data.WriteString(fmt.Sprintf("\n\t// TYPES: %s", p.typ))
					p.typ = "INTERFACE"
				}

				data.WriteString(fmt.Sprintf("\n\t%s %s `json:\"%s%s\"`\n\n", p.name, p.typ, p.nameSnakeCase, omitempty))
			}

			data.WriteString("}\n\n")

			additions := additional[parametersStruct]
			for _, a := range additions {
				data.WriteString(a)
				data.WriteString("\n")
			}
		}

		methodDescription := fitTextToLine(fmt.Sprintf("%s - %s", m.nameTitle, m.description), "// ")
		data.WriteString(methodDescription)

		var returnType string
		hasReturnType := false
		if !m.hasReturnValue() {
			returnType = "error"
		} else {
			returnType = fmt.Sprintf("(%s, error)", m.returnType)
			hasReturnType = true

			returnsCount++
		}

		if m.returnType == returnTypeNotFound {
			returnsNotFoundCount++
		}

		data.WriteString(fmt.Sprintf("\nfunc (b *Bot) %s(%s) %s {\n", m.nameTitle, parametersArg, returnType))

		returnVar := returnTypeToVar(m.returnType)
		switch m.nameTitle {
		case "ExportChatInviteLink":
			returnVar = "inviteLink"
		case "GetChatMemberCount":
			returnVar = "chatMemberCount"
		case "CreateInvoiceLink":
			returnVar = "invoiceLink"
		}

		if hasReturnType {
			data.WriteString(fmt.Sprintf("\tvar %s %s\n", returnVar, m.returnType))

			successValue := ""
			if strings.Contains(m.description, "otherwise True is returned") {
				data.WriteString("\tvar success *bool\n")
				successValue = ", &success"
				successValuesCount++
			}

			if len(m.parameters) > 0 {
				data.WriteString(fmt.Sprintf("\terr := b.performRequest(\"%s\", params, &%s%s)\n", m.name, returnVar, successValue))
			} else {
				data.WriteString(fmt.Sprintf("\terr := b.performRequest(\"%s\", nil, &%s%s)\n", m.name, returnVar, successValue))
			}

			data.WriteString(fmt.Sprintf("\tif err != nil {\n\t\treturn nil, fmt.Errorf(\"telego: %s(): %%w\", err)\n\t}\n\n", m.name))
			data.WriteString(fmt.Sprintf("\treturn %s, nil\n}\n\n", returnVar))
		} else {
			if len(m.parameters) > 0 {
				data.WriteString(fmt.Sprintf("\terr := b.performRequest(\"%s\", params)\n", m.name))
			} else {
				data.WriteString(fmt.Sprintf("\terr := b.performRequest(\"%s\", nil)\n", m.name))
			}

			data.WriteString(fmt.Sprintf("\tif err != nil {\n\t\treturn fmt.Errorf(\"telego: %s(): %%w\", err)\n\t}\n\n", m.name))
			data.WriteString("\treturn nil\n}\n\n")
		}
	}

	logInfof("Method parameters: %d", parametersCount)
	logInfof("Method returns: %d", returnsCount)
	logInfof("Method returns not found: %d", returnsNotFoundCount)
	logInfof("Method returns success: %d", successValuesCount)

	_, err := file.WriteString(uppercaseWords(data.String()))
	exitOnErr(err)
}

func methodSpecialCases(method *tgMethod) {
	if method.returnType == "string" || method.returnType == "int" {
		method.returnType = "*" + method.returnType
	}
}

func parameterSpecialCases(parameter *tgMethodParameter, methodName string) {
	if parameter.typ == "*InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply" {
		parameter.typ = "ReplyMarkup"
	}

	if parameter.typ == "[]InputMediaAudio, InputMediaDocument, InputMediaPhoto and InputMediaVideo" {
		parameter.typ = "[]InputMedia"
	}

	if (parameter.name == "UserId" || parameter.name == "ChatId" || parameter.name == "SenderChatId" ||
		strings.Contains(parameter.description, "number of seconds")) &&
		parameter.typ == "int" {
		parameter.typ = "int64"
	}

	if parameter.name == "Scope" && parameter.typ == "*BotCommandScope" {
		parameter.typ = parameter.typ[1:]
	}

	if parameter.name == "MenuButton" && parameter.typ == "*MenuButton" {
		parameter.typ = parameter.typ[1:]
	}

	if strings.Contains(parameter.name, "Date") && parameter.typ == "int" {
		parameter.typ = "int64"
	}

	if methodName == "promoteChatMember" && parameter.typ == "bool" {
		parameter.typ = "*bool"
	}

	if methodName == "editForumTopic" && parameter.nameSnakeCase == "icon_custom_emoji_id" {
		parameter.typ = "*string"
	}

	if methodName == "sendPoll" {
		switch parameter.nameSnakeCase {
		case "is_anonymous":
			parameter.typ = "*bool"
		case "correct_option_id":
			parameter.typ = "*int"
		}
	}
}

func parseReturnType(methodDescription string) string {
	methodDescription = removeHTML(methodDescription)
	var returnType string

	returns1 := returnTypeRegexp1.FindStringSubmatch(methodDescription)
	if len(returns1) != 0 {
		returnType = returns1[1]
	}

	returns2 := returnTypeRegexp2.FindStringSubmatch(methodDescription)
	if len(returns2) != 0 {
		returnType = returns2[1]
	}

	switch returnType {
	case "":
		return returnTypeNotFound
	case "True", "error":
		return ""
	case "Messages":
		return "[]Message"
	default:
		return parseType(returnType, true)
	}
}

func returnTypeToVar(returnType string) string {
	returnVar := strings.TrimPrefix(returnType, "*")
	if strings.HasPrefix(returnVar, "[]") {
		returnVar = strings.TrimPrefix(returnVar, "[]") + "s"
	}
	return firstToLower(returnVar)
}
