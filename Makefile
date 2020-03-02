GOPKGS := $(shell go list ./... | grep -v "/test")
ANTLR4 := java -jar antlr-4.8-complete.jar

all: test build
.PHONY: all

test:
	go test -race -cover -covermode=atomic -v $(GOPKGS)
.PHONY: test

build:
	CGO_ENABLED=0 go build -o ./bin/parser ./cmd/parser

antlr-4.8-complete.jar:
	curl -o antlr-4.8-complete.jar https://www.antlr.org/download/antlr-4.8-complete.jar

compile-grammar: antlr-4.8-complete.jar
	$(ANTLR4) -Dlanguage=Go -package grammar internal/grammar/AnnotationGrammar.g4

.PHONY: build
