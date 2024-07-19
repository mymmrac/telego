package main

import (
	"fmt"
	"strings"
)

func generateMethodsTests(methods tgMethods) {
	methodsTestsFile := openFile(generatedMethodsTestsFilename)
	defer func() { _ = methodsTestsFile.Close() }()

	data := strings.Builder{}

	data.WriteString(`package telego

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	ta "github.com/chococola/telego/telegoapi"
)
`)

	for _, m := range methods {
		data.WriteString(fmt.Sprintf("func TestBot_%s(t *testing.T) {\n", m.nameTitle))
		data.WriteString(`	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)
`)

		respVar := "emptyResp"
		actualVar := returnTypeToVar(m.returnType)
		expectedVar := fmt.Sprintf("expected%s", firstToUpper(actualVar))
		if m.hasReturnValue() {
			expectedData := strings.Replace(m.returnType, "*", "&", 1) + "{}"

			// Special case
			if expectedVar != "expectedMessage" {
				data.WriteString(fmt.Sprintf("\n\t\t%s := %s", expectedVar, expectedData))
			}

			respVar = "resp"
			data.WriteString(fmt.Sprintf("\n\t\tresp := telegoResponse(t, %s)", expectedVar))
		}

		parameters := ""
		if len(m.parameters) > 0 {
			parameters = "nil"
		}

		data.WriteString(`
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any()).
			Return(` + respVar + `, nil)`)
		data.WriteString("\n\n")

		if m.hasReturnValue() {
			data.WriteString(fmt.Sprintf(`		%s, err := m.Bot.%s(%s)
		require.NoError(t, err)
		assert.Equal(t, %s, %s)`, actualVar, m.nameTitle, parameters, expectedVar, actualVar))
		} else {
			data.WriteString(fmt.Sprintf(`		err := m.Bot.%s(%s)
		require.NoError(t, err)`, m.nameTitle, parameters))
		}

		data.WriteString("\n\t})\n\n")

		data.WriteString(`	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)`)
		data.WriteString("\n\n")

		if m.hasReturnValue() {
			data.WriteString(fmt.Sprintf(`		%s, err := m.Bot.%s(%s)
		require.Error(t, err)
		assert.Nil(t, %s)`, actualVar, m.nameTitle, parameters, actualVar))
		} else {
			data.WriteString(fmt.Sprintf(`		err := m.Bot.%s(%s)
		require.Error(t, err)`, m.nameTitle, parameters))
		}

		data.WriteString("\n\t})\n}\n\n")
	}

	_, err := methodsTestsFile.WriteString(uppercaseWords(data.String()))
	exitOnErr(err)

	formatFile(methodsTestsFile.Name())
}
