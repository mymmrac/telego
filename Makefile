lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0

lint:
	$(shell go env GOPATH)/bin/golangci-lint run

test:
	go test -coverprofile cover.out

cover: test
	go tool cover -func cover.out

pre-commit: test lint

.PHONY: lint-install lint test cover pre-commit
