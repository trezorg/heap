HAS_GOLANGCI := $(shell command -v golangci-lint;)

test:
	go test -v ./...

golangci:
ifndef HAS_GOLANGCI
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.1
endif
	golangci-lint run

lint: golangci

.PHONY: test lint
