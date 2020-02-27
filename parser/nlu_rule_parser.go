//parser/nlu_rule_parser.go
package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"time"
	"encoding/json"
)


//export run_parser
func RunParser(filename string) string {
	start := time.Now()
	input, _ := antlr.NewFileStream(filename)
	lexer := NewAnnotationGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := NewAnnotationGrammarParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Annotation()
	listener := NewNluRuleListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	elapsed := time.Since(start)
	fmt.Println("TIME", elapsed.Seconds())
	parsed, _ := json.MarshalIndent(listener.ParsedRules, "", " ")
	return string(parsed)
}
