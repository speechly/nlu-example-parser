GOPKGS := $(shell go list ./... | grep -v "/test")

all: lint test build
.PHONY: all

lint:
	golangci-lint run --exclude-use-default=false
.PHONY: lint

test:
	go test -race -cover -covermode=atomic -v $(GOPKGS)
.PHONY: test

build:
	CGO_ENABLED=0 go build -o ./bin/parser ./cmd/parser
.PHONY: build
