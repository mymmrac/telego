package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const funcPattern = `
func \(\w* \*(\w*)\) (\w*)\(\) string {
	return (\w*)
}
`

const (
	typesFilename = "./types.go"

	generatedTypesTestsFilename = "types_test.go.generated"
)

var funcRegexp *regexp.Regexp

func init() {
	var err error
	funcRegexp, err = regexp.Compile(funcPattern)
	exitOnErr(err)
}

func generateTypesTests() {
	info("Reading types from: %s", typesFilename)

	typesBytes, err := ioutil.ReadFile(typesFilename)
	exitOnErr(err)

	info("Types length: %d", len(typesBytes))

	types := string(typesBytes)

	typesTestsFile := openFile(generatedTypesTestsFilename)
	defer func() { _ = typesTestsFile.Close() }()

	data := strings.Builder{}

	data.WriteString(fmt.Sprintf(`package %s

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestTypesInterfaces(t *testing.T) {
`, packageName))

	funcs := funcRegexp.FindAllStringSubmatch(types, -1)

	info("Func count: %d", len(funcs))

	for _, f := range funcs {
		funcType := f[1]
		funcName := f[2]
		funcReturn := f[3]

		data.WriteString(fmt.Sprintf("\tassert.Implements(t, (*INTERFACE)(nil), &%s{})\n", funcType))
		data.WriteString(fmt.Sprintf("\tassert.Equal(t, %s, (&%s{}).%s())\n\n", funcReturn, funcType, funcName))
	}

	data.WriteString("}\n")

	_, err = typesTestsFile.WriteString(data.String())
	exitOnErr(err)

	formatFile(typesTestsFile.Name())
	info("Done")
}
