package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/speechly/nlu-example-parser/internal/grammar"
)

type Parser struct {
	lis   *grammar.NluRuleListener
	debug bool
}

func NewParser(debug bool) *Parser {
	return NewBufferedParser(1, debug)
}

func NewBufferedParser(bufSize uint64, debug bool) *Parser {
	return &Parser{
		lis:   grammar.NewNluRuleListener(bufSize, debug),
		debug: debug,
	}
}

// TODO: we should work on implementing a custom `antrl.CharStream`,
// instead of re-creating ANTLR machinery for every line of text.
func (p *Parser) Parse(line string) {
	stream := antlr.NewCommonTokenStream(
		grammar.NewAnnotationGrammarLexer(
			antlr.NewInputStream(line),
		),
		0,
	)

	parse := grammar.NewAnnotationGrammarParser(stream)
	parse.BuildParseTrees = true

	if p.debug {
		parse.AddErrorListener(antlr.NewDiagnosticErrorListener(p.debug))
	}

	antlr.ParseTreeWalkerDefault.Walk(p.lis, parse.Annotation())
}

func (p *Parser) Results() <-chan grammar.Utterance {
	return p.lis.Utterances()
}

func (p *Parser) Close() {
	p.lis.Close()
}
