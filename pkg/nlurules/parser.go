package nlurules

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"speechly/nlu-rules-parser/pkg/parser"
)

func ParseAsync(in <-chan string) <-chan Utterance {
	listener := NewNluRuleListener()

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
			p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			p.BuildParseTrees = true

			antlr.ParseTreeWalkerDefault.Walk(listener, p.Annotation())
		}
	}()

	return listener.Utterances()
}
