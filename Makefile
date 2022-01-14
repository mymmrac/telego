export PATH := $(PATH):$(shell go env GOPATH)/bin
# TODO: Document all targets, create `help` target
# TODO: Add "how to run" to contribution guidelines

lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0

mock-install:
	go install github.com/golang/mock/mockgen@v1.6.0

generate:
	go generate ./...

lint:
	golangci-lint run

# TODO: [?] Test examples
test:
	go test -coverprofile cover.out \
	$(shell go list ./... | grep -v /examples/ | grep -v /test | grep -v /internal/ | grep -v /mock)

cover: test
	go tool cover -func cover.out

race:
	go test -race ./...

pre-commit: test lint

# Usage: make generator RUN="types types-tests methods methods-tests"
generator: ./internal/generator
	go run ./internal/generator $$RUN

generator-clean-up:
	rm *.generated

.PHONY: lint-install mock-install generate lint test cover race pre-commit generator generator-clean-up
