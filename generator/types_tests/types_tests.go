package main

import (
	"fmt"
	"github.com/mymmrac/telego/generator"
	"io/ioutil"
	"os"
	"regexp"
)

const funcPattern = `
func \(\w* \*(\w*)\) (\w*)\(\) string {
	return (\w*)
}
`

func main() {
	funcPatternReg := regexp.MustCompile(generator.RemoveNewline(funcPattern))

	typesBytes, err := ioutil.ReadFile("types.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	types := string(typesBytes)
	types = generator.RemoveNewline(types)

	file, err := os.Create("types_test.go.generated")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = fmt.Fprintf(file, `package telego

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

`)

	funcs := funcPatternReg.FindAllStringSubmatch(types, -1)

	fmt.Println("Count of func:", len(funcs))

	for _, f := range funcs {
		funcType := f[1]
		funcName := f[2]
		funcReturn := f[3]

		_, _ = fmt.Fprintf(file, `func Test%s_%s(t *testing.T) {
	assert.Equal(t, %s, (&%s{}).%s())
}

`, funcType, funcName, funcReturn, funcType, funcName)
	}
}
