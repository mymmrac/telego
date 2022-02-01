package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type tgMethodSetter struct {
	paramStructType    string
	paramName          string
	paramSnakeCaseName string
	paramType          string
}

type tgMethodSetters []tgMethodSetter

const methodParamPattern = `
type (\w+Params) struct {
	(.+?)
}
`

const paramPattern = "(\\w+) ([*A-Za-z\\[\\]]+) `json:\"(\\w+)"

const (
	methodsFilename = "./methods.go"

	generatedMethodsSettersFilename = "methods_setters.go.generated"
)

var (
	methodParamRegexp = regexp.MustCompile(preparePattern(methodParamPattern))
	paramRegexp       = regexp.MustCompile(paramPattern)
)

func generateMethodsSetters() tgMethodSetters {
	logInfo("Reading methods from: %q", methodsFilename)

	methodsBytes, err := ioutil.ReadFile(methodsFilename)
	exitOnErr(err)

	logInfo("Methods length: %d", len(methodsBytes))

	methods := removeNl(string(methodsBytes))

	methodsParamsGroups := methodParamRegexp.FindAllStringSubmatch(methods, -1)
	logInfo("Methods params count: %d", len(methodsParamsGroups))

	var setters tgMethodSetters

	count := 0
	for _, methodsParamsGroup := range methodsParamsGroups {
		paramStructType := methodsParamsGroup[1]
		params := methodsParamsGroup[2]

		paramsGroups := paramRegexp.FindAllStringSubmatch(params, -1)
		count += len(paramsGroups)

		for _, paramsGroup := range paramsGroups {
			setter := tgMethodSetter{
				paramStructType:    paramStructType,
				paramName:          paramsGroup[1],
				paramSnakeCaseName: paramsGroup[3],
				paramType:          paramsGroup[2],
			}

			setters = append(setters, setter)
		}
	}

	logInfo("Setters count: %d", count)

	return setters
}

func writeMethodsSetters(file *os.File, setters tgMethodSetters) {
	data := strings.Builder{}

	data.WriteString(`package telego` + "\n\n")

	for _, setter := range setters {
		data.WriteString(fmt.Sprintf("// With%s adds %s parameter\n", setter.paramName,
			strings.ReplaceAll(setter.paramSnakeCaseName, "_", " ")))

		setterSpecialCase(&setter)

		if strings.HasPrefix(setter.paramType, "[]") {
			setter.paramType = strings.Replace(setter.paramType, "[]", "...", 1)
		}

		if setter.paramType != "bool" {
			data.WriteString(fmt.Sprintf("func (p *%s) With%s(%s %s) *%s {\n", setter.paramStructType,
				setter.paramName, firstToLower(setter.paramName), setter.paramType, setter.paramStructType))

			data.WriteString(fmt.Sprintf("\tp.%s = %s\n", setter.paramName, firstToLower(setter.paramName)))
		} else {
			data.WriteString(fmt.Sprintf("func (p *%s) With%s() *%s {\n", setter.paramStructType,
				setter.paramName, setter.paramStructType))

			data.WriteString(fmt.Sprintf("\tp.%s = true\n", setter.paramName))
		}

		data.WriteString("\treturn p\n}\n\n")
	}

	_, err := file.WriteString(uppercaseWords(data.String()))
	exitOnErr(err)
}

func setterSpecialCase(setter *tgMethodSetter) {
	if setter.paramStructType == "SendPollParams" && setter.paramName == "Type" {
		setter.paramName = "PollType"
	}
}
