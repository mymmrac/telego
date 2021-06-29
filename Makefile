lint-install:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0

lint:
	$(shell go env GOPATH)/bin/golangci-lint run