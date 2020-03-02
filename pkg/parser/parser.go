package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"speechly/nlu-example-parser/internal/grammar"
)

// ParseAsync processes input strings in a streaming fashion.
// It assumes that a single string passed through the input channel contains a parseable utterance.
//
// TODO: combine logDiagnostics and verbose into a ternary log level.
func ParseAsync(in <-chan string, logDiagnostics bool, verbose bool) <-chan grammar.Utterance {
	listener := grammar.NewNluRuleListener(1, logDiagnostics)

	go func() {
		defer listener.Close()

		for str := range in {
			stream := antlr.NewCommonTokenStream(
				grammar.NewAnnotationGrammarLexer(
					antlr.NewInputStream(str),
				),
				0,
			)

			p := grammar.NewAnnotationGrammarParser(stream)
			p.BuildParseTrees = true

			if logDiagnostics {
				p.AddErrorListener(antlr.NewDiagnosticErrorListener(verbose))
			}

			antlr.ParseTreeWalkerDefault.Walk(listener, p.Annotation())
		}
	}()

	return listener.Utterances()
}
