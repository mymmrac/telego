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
			if strings.HasPrefix(s.value, " ") {
				s.value = "ToPtr(" + s.value + ")"
			}

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
	case "*bool":
		return " true"
	case "string":
		return "\"" + setter.fieldName + "\""
	case "*string":
		return " \"" + setter.fieldName + "\""
	case "int", "int32", "int64":
		*counter++
		return fmt.Sprintf("%d", *counter)
	case "*int":
		*counter++
		return fmt.Sprintf(" %d", *counter)
	case "float64":
		*counter++
		return fmt.Sprintf("%d.0", *counter)
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
		return "&ReplyKeyboardRemove{RemoveKeyboard: true}"
	case "[]InputMedia":
		return fmt.Sprintf("[]InputMedia{&InputMediaAnimation{Type: \"%s\"}}", setter.fieldName)
	case "[]BotCommand":
		return fmt.Sprintf("[]BotCommand{{Command: \"%s\"}}", setter.fieldName)
	case "BotCommandScope":
		return fmt.Sprintf("&BotCommandScopeDefault{Type: \"%s\"}", setter.fieldName)
	case "*InlineKeyboardMarkup":
		return "&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}"
	case "[]InlineQueryResult":
		return fmt.Sprintf("[]InlineQueryResult{&InlineQueryResultArticle{Type: \"%s\"}}", setter.fieldName)
	case "[]LabeledPrice":
		return fmt.Sprintf("[]LabeledPrice{{Label: \"%s\"}}", setter.fieldName)
	case "[]int":
		*counter++
		return fmt.Sprintf("[]int{%d}", *counter)
	case "[]ShippingOption":
		return fmt.Sprintf("[]ShippingOption{{ID: \"%s\"}}", setter.fieldName)
	case "[]PassportElementError":
		return "[]PassportElementError{&PassportElementErrorDataField{}}"
	case "*MaskPosition":
		return fmt.Sprintf("&MaskPosition{Point: \"%s\"}", setter.fieldName)
	case "InputMedia":
		return fmt.Sprintf("&InputMediaAnimation{Type: \"%s\"}", setter.fieldName)
	case "ChatPermissions":
		return "ChatPermissions{CanSendMessages: ToPtr(true)}"
	case "InputMessageContent":
		return "&InputTextMessageContent{}"
	case "*CallbackGame":
		return "&CallbackGame{}"
	case "*LoginURL":
		return fmt.Sprintf("&LoginURL{URL: \"%s\"}", setter.fieldName)
	case "[][]InlineKeyboardButton":
		return "[][]InlineKeyboardButton{{}}"
	case "*KeyboardButtonPollType":
		return fmt.Sprintf("&KeyboardButtonPollType{Type: \"%s\"}", setter.fieldName)
	case "[][]KeyboardButton":
		return "[][]KeyboardButton{{}}"
	case "WebAppInfo":
		return "WebAppInfo{}"
	case "*WebAppInfo":
		return "&WebAppInfo{}"
	case "MenuButton":
		return fmt.Sprintf("&MenuButtonCommands{Type: \"%s\"}", setter.fieldName)
	case "*ChatAdministratorRights":
		return "&ChatAdministratorRights{IsAnonymous: true}"
	case "InlineQueryResult":
		return fmt.Sprintf("&InlineQueryResultArticle{Type: \"%s\"}", setter.fieldName)
	case "*KeyboardButtonRequestUsers":
		*counter++
		return fmt.Sprintf("&KeyboardButtonRequestUsers{RequestID: %d}", *counter)
	case "*KeyboardButtonRequestChat":
		*counter++
		return fmt.Sprintf("&KeyboardButtonRequestChat{RequestID: %d}", *counter)
	case "[]InputSticker":
		return "[]InputSticker{{}}"
	case "InputSticker":
		return "InputSticker{Sticker: testInputFile}"
	case "*SwitchInlineQueryChosenChat":
		return "&SwitchInlineQueryChosenChat{AllowUserChats: true}"
	case "*InlineQueryResultsButton":
		return "&InlineQueryResultsButton{}"
	case "*LinkPreviewOptions":
		return "&LinkPreviewOptions{IsDisabled: true}"
	case "*ReplyParameters":
		*counter++
		return fmt.Sprintf("&ReplyParameters{MessageID: %d}", *counter)
	case "[]ReactionType":
		return "[]ReactionType{&ReactionTypeEmoji{Type: ReactionEmoji}}"
	case "[]InputPollOption":
		return "[]InputPollOption{{}}"
	case "[]InputPaidMedia":
		return "[]InputPaidMedia{&InputPaidMediaPhoto{}}"
	case "*CopyTextButton":
		return "&CopyTextButton{}"
	default:
		return "UNKNOWN: " + setter.fieldType
	}
}
