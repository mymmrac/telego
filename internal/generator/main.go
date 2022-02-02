package main

import (
	"bytes"
	"io"
	"io/ioutil"
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

	typesFilename   = "./types.go"
	methodsFilename = "./methods.go"

	generatedTypesFilename          = "./types.go.generated"
	generatedTypesTestsFilename     = "./types_test.go.generated"
	generatedTypesSettersFilename   = "./types_setters.go.generated"
	generatedMethodsFilename        = "./methods.go.generated"
	generatedMethodsTestsFilename   = "./methods_test.go.generated"
	generatedMethodsSettersFilename = "./methods_setters.go.generated"
)

const (
	runTypesGeneration          = "types"
	runTypesTestsGeneration     = "types-tests"
	runTypesSettersGeneration   = "types-setters"
	runMethodsGeneration        = "methods"
	runMethodsTestsGeneration   = "methods-tests"
	runMethodsSettersGeneration = "methods-setters"
)

var typeStructsSetters = []string{
	"ReplyKeyboardMarkup",
	"KeyboardButton",
	"InlineKeyboardButton",

	"InlineQueryResultCachedAudio",
	"InlineQueryResultCachedDocument",
	"InlineQueryResultCachedGif",
	"InlineQueryResultCachedMpeg4Gif",
	"InlineQueryResultCachedPhoto",
	"InlineQueryResultCachedSticker",
	"InlineQueryResultCachedVideo",
	"InlineQueryResultCachedVoice",
	"InlineQueryResultArticle",
	"InlineQueryResultAudio",
	"InlineQueryResultContact",
	"InlineQueryResultGame",
	"InlineQueryResultDocument",
	"InlineQueryResultGif",
	"InlineQueryResultLocation",
	"InlineQueryResultMpeg4Gif",
	"InlineQueryResultPhoto",
	"InlineQueryResultVenue",
	"InlineQueryResultVideo",
	"InlineQueryResultVoice",
}

var typeStructsNoPointerSetters = []string{
	"KeyboardButton",
	"InlineKeyboardButton",
}

func main() {
	if len(os.Args) <= 1 {
		logError("Generation args not specified")
		os.Exit(1)
	}
	args := os.Args[1:]

	sr := sharedResources{}

	for _, arg := range args {
		logInfo("==== %s ====", arg)
		switch arg {
		case runTypesGeneration:
			docs := sr.Docs()

			start := time.Now()
			typesFile := openFile(generatedTypesFilename)

			types := generateTypes(docs)
			writeTypes(typesFile, types)
			_ = typesFile.Close()

			formatFile(typesFile.Name())
			logInfo("Generated types in: %s", time.Since(start))
		case runMethodsGeneration:
			docs := sr.Docs()

			start := time.Now()
			methodsFile := openFile(generatedMethodsFilename)

			methods := sr.Methods(docs)
			writeMethods(methodsFile, methods)
			_ = methodsFile.Close()

			formatFile(methodsFile.Name())
			logInfo("Generated methods in: %s", time.Since(start))
		case runTypesTestsGeneration:
			start := time.Now()
			types := sr.TypesData()

			generateTypesTests(types)
			logInfo("Generated types tests in: %s", time.Since(start))
		case runMethodsTestsGeneration:
			docs := sr.Docs()
			methods := sr.Methods(docs)

			start := time.Now()
			generateMethodsTests(methods)
			logInfo("Generated methods tests in: %s", time.Since(start))
		case runMethodsSettersGeneration:
			start := time.Now()

			logInfo("Reading methods from: %q", methodsFilename)
			methodsBytes, err := ioutil.ReadFile(methodsFilename)
			exitOnErr(err)

			logInfo("Methods length: %d", len(methodsBytes))
			methods := removeNl(string(methodsBytes))

			methodsSettersFile := openFile(generatedMethodsSettersFilename)

			methodsSetters := generateSetters(methods, nil)
			writeSetters(methodsSettersFile, methodsSetters, true, nil)
			_ = methodsSettersFile.Close()

			formatFile(methodsSettersFile.Name())
			logInfo("Generated methods setters in: %s", time.Since(start))
		case runTypesSettersGeneration:
			start := time.Now()

			types := removeNl(sr.TypesData())

			typesSettersFile := openFile(generatedTypesSettersFilename)

			typesSetters := generateSetters(types, typeStructsSetters)
			writeSetters(typesSettersFile, typesSetters, false, typeStructsNoPointerSetters)
			_ = typesSettersFile.Close()

			formatFile(typesSettersFile.Name())
			logInfo("Generated types setters in: %s", time.Since(start))
		default:
			logError("Unknown generation arg: %q", arg)
			os.Exit(1)
		}
	}

	logInfo("==== end ====")
	logInfo("Generation successful")
}

type sharedResources struct {
	docs    string
	methods tgMethods

	typesData string
}

func (r *sharedResources) Docs() string {
	if r.docs == "" {
		r.docs = readDocs()
	} else {
		logInfo("Reusing docs")
	}
	return r.docs
}

func (r *sharedResources) Methods(docs string) tgMethods {
	if r.methods == nil {
		r.methods = generateMethods(docs)
	} else {
		logInfo("Reusing methods")
	}
	return r.methods
}

func (r *sharedResources) TypesData() string {
	if r.typesData == "" {
		logInfo("Reading types from: %q", typesFilename)

		typesBytes, err := ioutil.ReadFile(typesFilename)
		exitOnErr(err)

		logInfo("Types length: %d", len(typesBytes))

		r.typesData = string(typesBytes)
	} else {
		logInfo("Reusing types data")
	}

	return r.typesData
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
