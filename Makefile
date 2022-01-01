export PATH := $(PATH):$(shell go env GOPATH)/bin

lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0

mock-install:
	go install github.com/golang/mock/mockgen@v1.6.0

generate:
	go generate ./...

lint:
	golangci-lint run

test:
	go test -coverprofile cover.out \
	$(shell go list ./... | grep -v /examples/ | grep -v /test | grep -v /generator | grep -v /mock)

cover: test
	go tool cover -func cover.out

race:
	go test -race ./...

pre-commit: test lint

# Usage: make generator RUN="types types-tests methods"
generator: ./internal/generator-v2
	go run ./internal/generator-v2 $$RUN

.PHONY: lint-install mock-install generate lint test cover race pre-commit generator
