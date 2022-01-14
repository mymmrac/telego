# Adds $GOPATH/bit to $PATH
export PATH := $(PATH):$(shell go env GOPATH)/bin
# TODO: Add "how to run" to contribution guidelines

help: ## Display this help message
	@echo "Usage:\n"
	@grep -E "^[a-zA-Z_-]+:.*? ## .+$$" $(MAKEFILE_LIST) | sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint-install: ## Install golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0

mock-install: ## Install mockgen
	go install github.com/golang/mock/mockgen@v1.6.0

generate: ## Generate (used for mock generation)
	go generate ./...

lint: ## Run golangci-lint
	golangci-lint run

# TODO: [?] Test examples
test: ## Run tests
	go test -coverprofile cover.out \
	$(shell go list ./... | grep -v /examples/ | grep -v /test | grep -v /internal/ | grep -v /mock)

cover: test ## Run tests & show coverage
	go tool cover -func cover.out

# TODO: Fix data race
race: ## Run tests with race flag
	go test -race ./...

pre-commit: test lint ## Run tests & linter

generator: ./internal/generator ## Run generation, example: make generator RUN="types types-tests methods methods-tests"
	go run ./internal/generator $$RUN

generator-clean-up: ## Remove generated files
	rm *.generated

.PHONY: lint-install mock-install generate lint test cover race pre-commit generator generator-clean-up help
