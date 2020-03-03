ANTLR_JAR	:= ./tools/antlr-4.8-complete.jar
GOPKGS		:= $(shell go list ./... | grep -v "/test")
GOOS			:= darwin
GOARCH		:= amd64
GONAME		:= ./bin/parser

all: compile-grammar test build
.PHONY: all

test:
	go test -race -cover -covermode=atomic -v $(GOPKGS)
.PHONY: test

build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(GONAME) ./cmd/parser
.PHONY: build

compile-grammar: $(ANTLR_JAR)
	java -jar $(ANTLR_JAR) -Dlanguage=Go -package grammar internal/grammar/AnnotationGrammar.g4
.PHONY: compile-grammar

$(ANTLR_JAR):
	curl -o $(ANTLR_JAR) https://www.antlr.org/download/antlr-4.8-complete.jar
