//parser/nlu_rule_parser.go
package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)



func RunAsync(in <-chan string) <-chan Utterance {
	listener := NewNluRuleListener()

	go func() {

		for str := range in {
			input := antlr.NewInputStream(str)
			lexer := NewAnnotationGrammarLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := NewAnnotationGrammarParser(stream)
			p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			p.BuildParseTrees = true
			tree := p.Annotation()
			antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		}
		defer listener.Close()
	}()

	return listener.Utterances()
}