package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const (
	baseURL = "https://core.telegram.org"
	docsURL = baseURL + "/bots/api"

	maxLineLen      = 110
	omitemptySuffix = ",omitempty"
	optionalPrefix  = "Optional. "

	packageName     = "telego"
	typesFilename   = "./types.go.generated"
	methodsFilename = "./methods.go.generated"
)

func main() {
	info("Reading docs...")
	start := time.Now()
	docs, err := docsText()
	exitOnErr(err)

	docs = removeNl(docs)
	info("Download docs in: %s", time.Since(start))

	start = time.Now()
	typesFile := openFile(typesFilename)
	defer func() { _ = typesFile.Close() }()

	types := generateTypes(docs)
	writeTypes(typesFile, types)

	formatFile(typesFile.Name())
	info("Generated types in: %s", time.Since(start))

	//start = time.Now()
	//methodsFile := openFile(methodsFilename)
	//defer func() { _ = methodsFile.Close() }()
	//
	//methods := generateMethods(docs)
	//writeMethods(methodsFile, methods)
	//
	//formatFile(methodsFile.Name())
	//info("Generated methods in: %s", time.Since(start))

	info("Done")
}

func openFile(filename string) *os.File {
	file, err := os.Create(filename)
	exitOnErr(err)
	info("File %q created", file.Name())

	return file
}

func docsText() (string, error) {
	response, err := http.Get(docsURL)
	if err != nil {
		return "", err
	}
	defer func() { _ = response.Body.Close() }()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func formatFile(filename string) {
	err := exec.Command("gofmt", "-w", filename).Run()
	exitOnErr(err)
}
