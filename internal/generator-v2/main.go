package main

import (
	"bytes"
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

	packageName              = "telego"
	generatedTypesFilename   = "./types.go.generated"
	generatedMethodsFilename = "./methods.go.generated"
)

const (
	runTypesGeneration      = "types"
	runTypesTestsGeneration = "types-tests"
	runMethodsGeneration    = "methods"
	// TODO: Methods test generator
)

func main() {
	if len(os.Args) <= 1 {
		logError("Generation args not specified")
		os.Exit(1)
	}
	args := os.Args[1:]

	var docs string

	for _, arg := range args {
		logInfo("==== %s ====", arg)
		switch arg {
		case runTypesGeneration:
			if docs == "" {
				docs = readDocs()
			} else {
				logInfo("Reusing docs")
			}

			start := time.Now()
			typesFile := openFile(generatedTypesFilename)

			types := generateTypes(docs)
			writeTypes(typesFile, types)
			_ = typesFile.Close()

			formatFile(typesFile.Name())
			logInfo("Generated types in: %s", time.Since(start))
		case runMethodsGeneration:
			if docs == "" {
				docs = readDocs()
			} else {
				logInfo("Reusing docs")
			}

			start := time.Now()
			methodsFile := openFile(generatedMethodsFilename)

			methods := generateMethods(docs)
			writeMethods(methodsFile, methods)
			_ = methodsFile.Close()

			formatFile(methodsFile.Name())
			logInfo("Generated methods in: %s", time.Since(start))
		case runTypesTestsGeneration:
			start := time.Now()
			generateTypesTests()
			logInfo("Generated types tests in: %s", time.Since(start))
		default:
			logError("Unknown generation arg: %q", arg)
			os.Exit(1)
		}
	}
	logInfo("Done")
}

func openFile(filename string) *os.File {
	file, err := os.Create(filename)
	exitOnErr(err)
	logInfo("File %q created", file.Name())

	return file
}

func readDocs() string {
	logInfo("Reading docs...")
	start := time.Now()
	docs, err := docsText()
	exitOnErr(err)

	docs = removeNl(docs)
	logInfo("Downloaded docs in: %s", time.Since(start))
	return docs
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
	buf := bytes.Buffer{}

	cmd := exec.Command("gofmt", "-w", filename)
	cmd.Stderr = &buf

	if err := cmd.Run(); err != nil {
		logError("Gofmt: %v\n%s", err, buf.String())
		os.Exit(1)
	}
}
