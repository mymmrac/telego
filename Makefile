export PATH := $(PATH):$(shell go env GOPATH)/bin

lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1

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

pre-commit: test lint

.PHONY: lint-install lint test cover pre-commit mock-install
