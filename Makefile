GOPKGS := $(shell go list ./... | grep -v "/test")

all: test build
.PHONY: all

test:
	go test -race -cover -covermode=atomic -v $(GOPKGS)
.PHONY: test

build:
	CGO_ENABLED=0 go build -o ./bin/parser ./cmd/parser
.PHONY: build
