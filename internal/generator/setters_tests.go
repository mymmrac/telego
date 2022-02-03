package main

import (
	"fmt"
	"os"
	"strings"
)

type setterTest struct {
	name  string
	value string
}

func writeSettersTests(file *os.File, setters tgSetters, noPointerStructs []string) {
	data := strings.Builder{}

	data.WriteString(`package telego

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)`)
	data.WriteString("\n\n")

	currentStruct := setters[0].structType
	var currentSetters []setterTest
	counter := 0
	for i, setter := range setters {
		if currentStruct == setter.structType {
			currentSetters = append(currentSetters, setterTest{
				name:  setter.fieldName,
				value: parseSetterType(setter, &counter),
			})

			if i != len(setters)-1 {
				continue
			}
		}

		data.WriteString(fmt.Sprintf("func Test%s_Setters(t *testing.T) {\n", currentStruct))

		pointer := "&"
		if contains(noPointerStructs, currentStruct) {
			pointer = ""
		}
		valueName := firstToLower(string(currentStruct[0]))
		data.WriteString(fmt.Sprintf("\t%s := (%s%s{}).\n", valueName, pointer, currentStruct))

		for j, s := range currentSetters {
			if s.value != "true" {
				v := s.value
				if strings.HasPrefix(s.value, "[]") {
					v += "..."
				}

				data.WriteString(fmt.Sprintf("\t\tWith%s(%s)", s.name, v))
			} else {
				data.WriteString(fmt.Sprintf("\t\tWith%s()", s.name))
			}

			if j != len(currentSetters)-1 {
				data.WriteString(".\n")
			}
		}
		data.WriteString("\n\n")

		data.WriteString(fmt.Sprintf("\tassert.Equal(t, %s%s{\n", pointer, currentStruct))

		for _, s := range currentSetters {
			data.WriteString(fmt.Sprintf("\t\t%s: %s,\n", s.name, s.value))
		}

		data.WriteString(fmt.Sprintf("\t}, %s)\n", valueName))
		data.WriteString("}\n\n")

		currentStruct = setter.structType
		currentSetters = []setterTest{
			{
				name:  setter.fieldName,
				value: parseSetterType(setter, &counter),
			},
		}
		counter = 0
	}

	_, err := file.WriteString(uppercaseWords(data.String()))
	exitOnErr(err)
}

func parseSetterType(setter tgSetter, counter *int) string {
	switch setter.fieldType {
	case "bool":
		return "true"
	case "string":
		return "\"" + setter.fieldName + "\""
	case "int":
		*counter++
		return fmt.Sprintf("%d", *counter)
	case "ChatID":
		*counter++
		return fmt.Sprintf("ChatID{ID: %d}", *counter)
	case "[]string":
		return fmt.Sprintf("[]string{\"%s\"}", setter.fieldName)
	case "InputFile":
		return "testInputFile"
	case "*InputFile":
		return "&testInputFile"
	case "[]MessageEntity":
		return fmt.Sprintf("[]MessageEntity{{Type: \"%s\"}}", setter.fieldName)
	case "ReplyMarkup":
		return "&ReplyKeyboardRemove{}"
	case "[]InputMedia":
		return "[]InputMedia{&InputMediaAnimation{}}"
	case "[]BotCommand":
		return fmt.Sprintf("[]BotCommand{{Command: \"%s\"}}", setter.fieldName)
	case "BotCommandScope":
		return "&BotCommandScopeDefault{}"
	case "*InlineKeyboardMarkup":
		return "&InlineKeyboardMarkup{}"
	case "[]InlineQueryResult":
		return "[]InlineQueryResult{&InlineQueryResultArticle{}}"
	case "[]LabeledPrice":
		return "[]LabeledPrice{}"
	case "[]int":
		return "[]int{}"
	case "[]ShippingOption":
		return "[]ShippingOption{}"
	case "[]PassportElementError":
		return "[]PassportElementError{}"
	case "*MaskPosition":
		return "&MaskPosition{}"
	case "InputMedia":
		return "&InputMediaAnimation{}"
	case "ChatPermissions":
		return "ChatPermissions{}"
	default:
		return "UNKNOWN: " + setter.fieldType
	}
}

/*
func TestGetUpdatesParams_Setters(t *testing.T) {
	g := (&GetUpdatesParams{}).
		WithOffset(1).
		WithLimit(2).
		WithTimeout(3).
		WithAllowedUpdates("AllowedUpdates")

	assert.Equal(t, &GetUpdatesParams{
		Offset:         1,
		Limit:          2,
		Timeout:        3,
		AllowedUpdates: []string{"AllowedUpdates"},
	}, g)
}
*/
