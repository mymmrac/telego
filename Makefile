# Adds $GOPATH/bit to $PATH
export PATH := $(PATH):$(shell go env GOPATH)/bin

help: ## Display this help message
	@echo "Usage:"
	@grep -E "^[a-zA-Z_-]+:.*? ## .+$$" $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}'

lint: ## Run golangci-lint
	golangci-lint run

lint-install: ## Install golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint-list: ## Run golangci-lint linters (print enabled & disabled linters)
	golangci-lint linters

test: ## Run tests
	go test -coverprofile cover.out \
	$(shell go list ./... | grep -v /examples/ | grep -v /test | grep -v /internal/ | grep -v /mock)

cover: test ## Run tests & show coverage
	go tool cover -func cover.out

race: ## Run tests with race flag
	go test -race -count=1 ./...

bench: ## Run benchmarks
	go test -bench=. -benchtime=10s -benchmem

build-examples: ## Build examples into bin folder
	go build -o bin/ ./examples/*

pre-commit: test lint build-examples ## Run tests, linter and build examples

generate: ## Generate (used for mock generation)
	go generate ./...

mock-install: ## Install mockgen
	go install github.com/golang/mock/mockgen@latest

generator: ./internal/generator ## Run specific generation
	go run ./internal/generator $$RUN

generator-all: ./internal/generator ## Run all generation
	go run ./internal/generator types types-tests types-setters types-setters-tests methods methods-tests \
		methods-setters methods-setters-tests

generator-clean-up: ## Remove generated files
	rm *.generated

.PHONY: help lint lint-install test cover race build-examples pre-commit generate mock-install generator \
generator-all generator-clean-up
