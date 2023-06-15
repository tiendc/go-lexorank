prepare:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.49.0

test:
	@go test -cover -v ./...

lint:
	golangci-lint run -v ./...
