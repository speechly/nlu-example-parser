GOPKGS		:= $(shell go list ./... | grep -v "/test")

BUILDOS		:= darwin
BUILDARCH	:= amd64
BUILDNAME	:= ./bin/parser

ANTLR_JAR	:= ./tools/antlr-4.8-complete.jar

all: compile-grammar test build verify
.PHONY: all

test:
	CGO_ENABLED=1 go test -race -cover -covermode=atomic -v $(GOPKGS)
.PHONY: test

verify: static-build
	./bin/golden -input_file_path ./examples/test_multi_intent_data.md | diff - test/golden.json
	./bin/golden -input_file_path ./examples/test_incorrect_data.md 2>&1 | diff - test/golden_error
.PHONY: verify

static-build:
	CGO_ENABLED=0 go build -o ./bin/golden ./cmd/parser
.PHONY: static-build

build: $(BUILDNAME)

$(BUILDNAME): compile-grammar
	CGO_ENABLED=0 GOOS=$(BUILDOS) GOARCH=$(BUILDARCH) go build -o $(BUILDNAME) ./cmd/parser

compile-grammar: $(ANTLR_JAR)
	java -jar $(ANTLR_JAR) -Dlanguage=Go -package grammar internal/grammar/AnnotationGrammar.g4
.PHONY: compile-grammar

$(ANTLR_JAR):
	curl -o $(ANTLR_JAR) https://www.antlr.org/download/antlr-4.8-complete.jar
