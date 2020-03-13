package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/speechly/nlu-example-parser/internal/grammar"
)


type NluParsingErrorListener struct {
	errors int
}

func (l *NluParsingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors += 1
}

func (l *NluParsingErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.errors += 1
}

func (l *NluParsingErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.errors += 1
}

func (l *NluParsingErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	l.errors += 1
}

func NewNluParsingErrorListener() *NluParsingErrorListener {
	return &NluParsingErrorListener{
		errors: 0,
	}
}

type Parser struct {
	lis   *grammar.NluRuleListener
	debug bool
	errLis NluParsingErrorListener
}

func NewParser(debug bool) *Parser {
	return NewBufferedParser(1, debug)
}

func NewBufferedParser(bufSize uint64, debug bool) *Parser {
	return &Parser{
		lis:   grammar.NewNluRuleListener(bufSize, debug),
		debug: debug,
		errLis: NluParsingErrorListener(),
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

	parse.AddErrorListener(errLis)

	antlr.ParseTreeWalkerDefault.Walk(p.lis, parse.Annotation())
}

func (p *Parser) Results() <-chan grammar.Utterance {
	return p.lis.Utterances()
}

func (p *Parser) Close() {
	p.lis.Close()
	p.errLis.Close()
}

func (p *Parser) Errors() <-chan error {
	return p.errLis.Errs()
}
