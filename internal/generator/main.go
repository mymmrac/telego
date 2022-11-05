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

	typesFilename   = "./types.go"
	methodsFilename = "./methods.go"

	generatedTypesFilename               = "./types.go.generated"
	generatedTypesTestsFilename          = "./types_test.go.generated"
	generatedTypesSettersFilename        = "./types_setters.go.generated"
	generatedTypesSettersTestsFilename   = "./types_setters_test.go.generated"
	generatedMethodsFilename             = "./methods.go.generated"
	generatedMethodsTestsFilename        = "./methods_test.go.generated"
	generatedMethodsSettersFilename      = "./methods_setters.go.generated"
	generatedMethodsSettersTestsFilename = "./methods_setters_test.go.generated"
)

const (
	runTypesGeneration               = "types"
	runTypesTestsGeneration          = "types-tests"
	runTypesSettersGeneration        = "types-setters"
	runTypesSettersTestsGeneration   = "types-setters-tests"
	runMethodsGeneration             = "methods"
	runMethodsTestsGeneration        = "methods-tests"
	runMethodsSettersGeneration      = "methods-setters"
	runMethodsSettersTestsGeneration = "methods-setters-tests"
)

var typeStructsSetters = []string{
	"ReplyKeyboardMarkup",
	"ForceReply",
	"ReplyKeyboardRemove",
	"InlineKeyboardMarkup",
	"KeyboardButton",
	"InlineKeyboardButton",
	"MenuButtonWebApp",

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

	"InputMediaAnimation",
	"InputMediaDocument",
	"InputMediaAudio",
	"InputMediaPhoto",
	"InputMediaVideo",

	"InputTextMessageContent",
	"InputLocationMessageContent",
	"InputVenueMessageContent",
	"InputContactMessageContent",
	"InputInvoiceMessageContent",
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
		start := time.Now()
		switch arg {
		case runTypesGeneration:
			docs := sr.Docs()
			currentTypes := sr.TypesData()

			typesFile := openFile(generatedTypesFilename)

			types := generateTypes(docs)
			writeTypes(typesFile, types, currentTypes)
			_ = typesFile.Close()

			formatFile(typesFile.Name())
		case runMethodsGeneration:
			docs := sr.Docs()
			currentMethods := sr.MethodsData()

			methodsFile := openFile(generatedMethodsFilename)

			methods := sr.Methods(docs)
			writeMethods(methodsFile, methods, currentMethods)
			_ = methodsFile.Close()

			formatFile(methodsFile.Name())
		case runTypesTestsGeneration:
			types := sr.TypesData()

			generateTypesTests(types)
		case runMethodsTestsGeneration:
			docs := sr.Docs()
			methods := sr.Methods(docs)

			generateMethodsTests(methods)
		case runMethodsSettersGeneration:
			methodsSetters := sr.MethodsSetters()

			methodsSettersFile := openFile(generatedMethodsSettersFilename)
			writeSetters(methodsSettersFile, methodsSetters, true, nil)
			_ = methodsSettersFile.Close()

			formatFile(methodsSettersFile.Name())
		case runMethodsSettersTestsGeneration:
			methodsSetters := sr.MethodsSetters()

			methodsSettersTestsFile := openFile(generatedMethodsSettersTestsFilename)
			writeSettersTests(methodsSettersTestsFile, methodsSetters, nil)
			_ = methodsSettersTestsFile.Close()

			formatFile(methodsSettersTestsFile.Name())
		case runTypesSettersGeneration:
			types := removeNl(sr.TypesData())
			typesSetters := sr.TypesSetters(types)

			typesSettersFile := openFile(generatedTypesSettersFilename)
			writeSetters(typesSettersFile, typesSetters, false, typeStructsNoPointerSetters)
			_ = typesSettersFile.Close()

			formatFile(typesSettersFile.Name())
		case runTypesSettersTestsGeneration:
			types := removeNl(sr.TypesData())
			typesSetters := sr.TypesSetters(types)

			typesSettersTestsFile := openFile(generatedTypesSettersTestsFilename)
			writeSettersTests(typesSettersTestsFile, typesSetters, typeStructsNoPointerSetters)
			_ = typesSettersTestsFile.Close()

			formatFile(typesSettersTestsFile.Name())
		default:
			logError("Unknown generation arg: %q", arg)
			os.Exit(1)
		}

		logInfo("Generated %s in: %s\n", arg, time.Since(start))
	}

	logInfo("==== end ====")
	logInfo("Generation successful")
}

type sharedResources struct {
	docs    string
	methods tgMethods

	typesData   string
	methodsData string

	methodsSetters tgSetters
	typesSetters   tgSetters
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

		typesBytes, err := os.ReadFile(typesFilename)
		exitOnErr(err)

		logInfo("Types length: %d", len(typesBytes))

		r.typesData = string(typesBytes)
	} else {
		logInfo("Reusing types data")
	}

	return r.typesData
}

func (r *sharedResources) MethodsData() string {
	if r.methodsData == "" {
		logInfo("Reading methods from: %q", methodsFilename)

		mehtodsBytes, err := os.ReadFile(methodsFilename)
		exitOnErr(err)

		logInfo("Methods length: %d", len(mehtodsBytes))

		r.methodsData = string(mehtodsBytes)
	} else {
		logInfo("Reusing methods data")
	}

	return r.methodsData
}

func (r *sharedResources) MethodsSetters() tgSetters {
	if r.methodsSetters == nil {
		methods := removeNl(r.MethodsData())

		r.methodsSetters = generateSetters(methods, nil)
	} else {
		logInfo("Reusing methods setters")
	}

	return r.methodsSetters
}

func (r *sharedResources) TypesSetters(types string) tgSetters {
	if r.typesSetters == nil {
		r.typesSetters = generateSetters(types, typeStructsSetters)
	} else {
		logInfo("Reusing types setters")
	}

	return r.typesSetters
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

	cmd := exec.Command("goimports", "-w", filename)
	cmd.Stderr = &buf

	if err := cmd.Run(); err != nil {
		logError("Gofmt: %v\n%s", err, buf.String())
		os.Exit(1)
	}
}
