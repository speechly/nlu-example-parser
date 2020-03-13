package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/speechly/nlu-example-parser/internal/grammar"
)

type Parser struct {
	lis    *grammar.NluRuleListener
	errLis *grammar.NLUExampleErrorListener
}

func NewParser() *Parser {
	return NewBufferedParser(1)
}

func NewBufferedParser(bufSize uint16) *Parser {
	return &Parser{
		lis:    grammar.NewNluRuleListener(bufSize),
		errLis: grammar.NewNLUExampleErrorListener(bufSize, false),
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
	parse.RemoveErrorListeners()
	parse.AddErrorListener(p.errLis)
	parse.BuildParseTrees = true

	antlr.ParseTreeWalkerDefault.Walk(p.lis, parse.Annotation())
}

func (p *Parser) Results() <-chan grammar.Utterance {
	return p.lis.Utterances()
}

func (p *Parser) Close() {
	p.lis.Close()
	p.errLis.Close()
}

func (p *Parser) Errors() <-chan grammar.ParseError {
	return p.errLis.Errors()
}
