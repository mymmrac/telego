package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type tgSetter struct {
	structType         string
	fieldName          string
	fieldSnakeCaseName string
	fieldType          string
}

type tgSetters []tgSetter

const typeStructPattern = `
type (\w+) struct {
	(.+?)
}
`

const fieldPattern = "(\\w+) ([\\*A-Za-z\\[\\]]+) `json:\"(\\w+)"

var (
	typeStructRegexp = regexp.MustCompile(preparePattern(typeStructPattern))
	fieldRegexp      = regexp.MustCompile(fieldPattern)
)

func generateSetters(typesData string, desiredStructs []string) tgSetters {
	structsGroups := typeStructRegexp.FindAllStringSubmatch(typesData, -1)
	logInfo("Structs count: %d", len(structsGroups))

	var setters tgSetters

	for _, structsGroup := range structsGroups {
		structType := structsGroup[1]
		if len(desiredStructs) > 0 {
			found := contains(desiredStructs, structType)

			if !found {
				continue
			}
		}

		fields := structsGroup[2]

		fieldsGroups := fieldRegexp.FindAllStringSubmatch(fields, -1)

		for _, fieldsGroup := range fieldsGroups {
			setter := tgSetter{
				structType:         structType,
				fieldName:          fieldsGroup[1],
				fieldSnakeCaseName: fieldsGroup[3],
				fieldType:          fieldsGroup[2],
			}

			if (strings.HasPrefix(setter.structType, "InlineQueryResult") ||
				strings.HasPrefix(setter.structType, "InputMedia")) &&
				setter.fieldName == "Type" {
				continue
			}

			setters = append(setters, setter)
		}
	}

	logInfo("Setters count: %d", len(setters))

	return setters
}

func contains(slice []string, a string) bool {
	found := false
	for _, s := range slice {
		if s == a {
			found = true
			break
		}
	}

	return found
}

func writeSetters(file *os.File, setters tgSetters, receiverDefault bool, noPointerStructs []string) {
	data := strings.Builder{}

	data.WriteString(`package telego` + "\n\n")

	for _, setter := range setters {
		data.WriteString(fmt.Sprintf("// With%s adds %s parameter\n", setter.fieldName,
			strings.ReplaceAll(setter.fieldSnakeCaseName, "_", " ")))

		setterSpecialCase(&setter)

		if strings.HasPrefix(setter.fieldType, "[]") {
			setter.fieldType = strings.Replace(setter.fieldType, "[]", "...", 1)
		}

		r := "p"
		if !receiverDefault {
			r = firstToLower(string(setter.structType[0]))
		}

		noPointer := contains(noPointerStructs, setter.structType)

		var s string
		if setter.fieldType != "bool" {
			s = fmt.Sprintf("func (%s *%s) With%s(%s %s) *%s {\n", r, setter.structType,
				setter.fieldName, firstToLower(setter.fieldName), setter.fieldType, setter.structType)
		} else {
			s = fmt.Sprintf("func (%s *%s) With%s() *%s {\n", r, setter.structType,
				setter.fieldName, setter.structType)
		}

		if noPointer {
			s = strings.ReplaceAll(s, "*"+setter.structType, setter.structType)
		}

		if len(s) > maxLineLen+11 {
			s = strings.Replace(s, ") *", ",\n) *", 1)
		}

		data.WriteString(s)

		if setter.fieldType != "bool" {
			data.WriteString(fmt.Sprintf("\t%s.%s = %s\n", r, setter.fieldName, firstToLower(setter.fieldName)))
		} else {
			data.WriteString(fmt.Sprintf("\t%s.%s = true\n", r, setter.fieldName))
		}

		data.WriteString(fmt.Sprintf("\treturn %s\n}\n\n", r))
	}

	_, err := file.WriteString(uppercaseWords(data.String()))
	exitOnErr(err)
}

func setterSpecialCase(setter *tgSetter) {
	if setter.fieldName == "Type" {
		setter.fieldName = "XXXType"
	}
}
