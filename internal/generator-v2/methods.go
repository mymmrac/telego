package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
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

var (
	methodRegexp          *regexp.Regexp
	methodParameterRegexp *regexp.Regexp

	returnTypeRegexp1 *regexp.Regexp
	returnTypeRegexp2 *regexp.Regexp
)

func init() {
	var err error
	methodRegexp, err = regexp.Compile(preparePattern(methodPattern))
	exitOnErr(err)

	methodParameterRegexp, err = regexp.Compile(preparePattern(methodParameterPattern))
	exitOnErr(err)

	returnTypeRegexp1, err = regexp.Compile(returnTypePattern1)
	exitOnErr(err)

	returnTypeRegexp2, err = regexp.Compile(returnTypePattern2)
	exitOnErr(err)
}

func generateMethods(docs string) tgMethods {
	var methods tgMethods

	methodGroups := methodRegexp.FindAllStringSubmatch(docs, -1)
	for _, methodGroup := range methodGroups {
		method := tgMethod{
			name:        methodGroup[1],
			nameTitle:   strings.Title(methodGroup[1]),
			description: replaceHTML(methodGroup[2]),
			parameters:  generateMethodParameters(methodGroup[3]),
			returnType:  parseReturnType(methodGroup[2]),
		}

		methods = append(methods, method)
	}

	return methods
}

func generateMethodParameters(parametersDocs string) tgMethodParameters {
	var parameters tgMethodParameters

	parameterGroups := methodParameterRegexp.FindAllStringSubmatch(parametersDocs, -1)
	for _, parameterGroup := range parameterGroups {
		parameter := tgMethodParameter{
			name:          parameterGroup[1],
			nameSnakeCase: snakeToCamelCase(parameterGroup[1]),
			description:   replaceHTML(parameterGroup[4]),
		}

		parameter.optional = parameterGroup[3] == "Optional"
		parameter.typ = parseType(parameterGroup[2], parameter.optional)

		parameters = append(parameters, parameter)
	}

	return parameters
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
	case "", "True", "error":
		return ""
	default:
		return parseType(returnType, true)
	}
}

func writeMethods(file *os.File, methods tgMethods) {
	for _, m := range methods {
		fmt.Println(m.name, m.nameTitle, m.returnType)
		fmt.Println(m.description)

		for _, p := range m.parameters {
			fmt.Println(p)
		}

		fmt.Println()
	}
}
