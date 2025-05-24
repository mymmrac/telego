package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	baseURL = "https://core.telegram.org"
	docsURL = baseURL + "/bots/api"

	maxLineLen      = 110
	omitemptySuffix = ",omitempty"
	optionalPrefix  = "Optional. "

	typesFilename   = "../../types.go"
	methodsFilename = "../../methods.go"

	generatedTypesFilename               = "types.go.generated"
	generatedTypesTestsFilename          = "types_test.go.generated"
	generatedTypesSettersFilename        = "types_setters.go.generated"
	generatedTypesSettersTestsFilename   = "types_setters_test.go.generated"
	generatedMethodsFilename             = "methods.go.generated"
	generatedMethodsTestsFilename        = "methods_test.go.generated"
	generatedMethodsSettersFilename      = "methods_setters.go.generated"
	generatedMethodsSettersTestsFilename = "methods_setters_test.go.generated"
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
	"KeyboardButtonRequestUsers",
	"KeyboardButtonRequestChat",
	"ReplyParameters",
	"InputPollOption",

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

	"InputPaidMediaPhoto",
	"InputPaidMediaVideo",

	"InputTextMessageContent",
	"InputLocationMessageContent",
	"InputVenueMessageContent",
	"InputContactMessageContent",
	"InputInvoiceMessageContent",

	"InputSticker",

	"InputProfilePhotoStatic",
	"InputProfilePhotoAnimated",

	"InputStoryContentPhoto",
	"InputStoryContentVideo",
}

var typeStructsNoPointerSetters = []string{
	"KeyboardButton",
	"InlineKeyboardButton",
}

func main() {
	if len(os.Args) <= 1 {
		logErrorf("Generation args not specified")
		os.Exit(1)
	}
	args := os.Args[1:]

	sr := sharedResources{}

	for _, arg := range args {
		logInfof("==== %s ====", arg)
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
			logErrorf("Unknown generation arg: %q", arg)
			os.Exit(1)
		}

		logInfof("Generated %s in: %s\n", arg, time.Since(start))
	}

	logInfof("==== end ====")
	logInfof("Generation successful")
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
		logInfof("Reusing docs")
	}
	return r.docs
}

func (r *sharedResources) Methods(docs string) tgMethods {
	if r.methods == nil {
		r.methods = generateMethods(docs)
	} else {
		logInfof("Reusing methods")
	}
	return r.methods
}

func (r *sharedResources) TypesData() string {
	if r.typesData == "" {
		logInfof("Reading types from: %q", typesFilename)

		typesBytes, err := os.ReadFile(typesFilename)
		exitOnErr(err)

		logInfof("Types length: %d", len(typesBytes))

		r.typesData = string(typesBytes)
	} else {
		logInfof("Reusing types data")
	}

	return r.typesData
}

func (r *sharedResources) MethodsData() string {
	if r.methodsData == "" {
		logInfof("Reading methods from: %q", methodsFilename)

		methodsBytes, err := os.ReadFile(methodsFilename)
		exitOnErr(err)

		logInfof("Methods length: %d", len(methodsBytes))

		r.methodsData = string(methodsBytes)
	} else {
		logInfof("Reusing methods data")
	}

	return r.methodsData
}

func (r *sharedResources) MethodsSetters() tgSetters {
	if r.methodsSetters == nil {
		methods := removeNl(r.MethodsData())

		r.methodsSetters = generateSetters(methods, nil)
	} else {
		logInfof("Reusing methods setters")
	}

	return r.methodsSetters
}

func (r *sharedResources) TypesSetters(types string) tgSetters {
	if r.typesSetters == nil {
		r.typesSetters = generateSetters(types, typeStructsSetters)
	} else {
		logInfof("Reusing types setters")
	}

	return r.typesSetters
}

func openFile(filename string) *os.File {
	file, err := os.Create(filepath.Join("../../", filename))
	exitOnErr(err)
	logInfof("File %q created", file.Name())

	return file
}

func readDocs() string {
	logInfof("Reading docs...")
	start := time.Now()
	docs, err := docsText()
	exitOnErr(err)

	docs = removeNl(docs)
	logInfof("Downloaded docs in: %s", time.Since(start))
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
	cmd := exec.Command("goimports", "-w", filename)

	buf := bytes.Buffer{}
	cmd.Stderr = &buf

	if err := cmd.Run(); err != nil {
		logErrorf("Gofmt: %v\n%s", err, buf.String())
		os.Exit(1)
	}

	cmd = exec.Command("gofumpt", "-w", filename)

	buf = bytes.Buffer{}
	cmd.Stderr = &buf

	if err := cmd.Run(); err != nil {
		logErrorf("Gofumpt: %v\n%s", err, buf.String())
		os.Exit(1)
	}
}
