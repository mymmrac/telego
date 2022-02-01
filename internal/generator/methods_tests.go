package main

import (
	"fmt"
	"strings"
)

const (
	generatedMethodsTestsFilename = "methods_test.go.generated"
)

func generateMethodsTests(methods tgMethods) {
	methodsTestsFile := openFile(generatedMethodsTestsFilename)
	defer func() { _ = methodsTestsFile.Close() }()

	data := strings.Builder{}

	data.WriteString(`//nolint:dupl
package telego

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego/telegoapi"
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

		actualVar := returnTypeToVar(m.returnType)
		expectedVar := fmt.Sprintf("expected%s", firstToUpper(actualVar))
		if m.hasReturnValue() {
			expectedData := strings.Replace(m.returnType, "*", "&", 1) + "{}"

			// Special case
			if expectedVar != "expectedMessage" {
				data.WriteString(fmt.Sprintf("\n\t\t%s := %s", expectedVar, expectedData))
			}

			data.WriteString(fmt.Sprintf("\n\t\tsetResult(t, %s)", expectedVar))
		}

		parameters := ""
		if len(m.parameters) > 0 {
			parameters = "nil"
		}

		data.WriteString(`
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)`)
		data.WriteString("\n\n")

		if m.hasReturnValue() {
			data.WriteString(fmt.Sprintf(`		%s, err := m.Bot.%s(%s)
		assert.NoError(t, err)
		assert.Equal(t, %s, %s)`, actualVar, m.nameTitle, parameters, expectedVar, actualVar))
		} else {
			data.WriteString(fmt.Sprintf(`		err := m.Bot.%s(%s)
		assert.NoError(t, err)`, m.nameTitle, parameters))
		}

		data.WriteString("\n\t})\n\n")

		data.WriteString(`	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)`)
		data.WriteString("\n\n")

		if m.hasReturnValue() {
			data.WriteString(fmt.Sprintf(`		%s, err := m.Bot.%s(%s)
		assert.Error(t, err)
		assert.Nil(t, %s)`, actualVar, m.nameTitle, parameters, actualVar))
		} else {
			data.WriteString(fmt.Sprintf(`		err := m.Bot.%s(%s)
		assert.Error(t, err)`, m.nameTitle, parameters))
		}

		data.WriteString("\n\t})\n}\n\n")
	}

	_, err := methodsTestsFile.WriteString(uppercaseWords(data.String()))
	exitOnErr(err)

	formatFile(methodsTestsFile.Name())
}
