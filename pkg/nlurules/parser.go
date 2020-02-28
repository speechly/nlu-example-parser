package nlurules

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"speechly/nlu-rules-parser/pkg/parser"
)

// ParseAsync processes input strings in a streaming fashion.
// It assumes that a single string passed through the input channel contains a parseable utterance.
//
// TODO: combine logDiagnostics and verbose into a ternary log level.
func ParseAsync(in <-chan string, logDiagnostics bool, verbose bool) <-chan Utterance {
	listener := NewNluRuleListener(1, logDiagnostics)

	go func() {
		defer listener.Close()

		for str := range in {
			stream := antlr.NewCommonTokenStream(
				parser.NewAnnotationGrammarLexer(
					antlr.NewInputStream(str),
				),
				0,
			)

			p := parser.NewAnnotationGrammarParser(stream)
			p.BuildParseTrees = true

			if logDiagnostics {
				p.AddErrorListener(antlr.NewDiagnosticErrorListener(verbose))
			}

			antlr.ParseTreeWalkerDefault.Walk(listener, p.Annotation())
		}
	}()

	return listener.Utterances()
}
