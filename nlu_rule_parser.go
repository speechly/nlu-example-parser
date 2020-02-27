//nlu_rule_parser.go
package main

import (
	"C"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"./nlu_rule_parser"
	"fmt"
	"time"
	"encoding/json"
	"os"
)


//export run_parser
func run_parser(filename string) *C.char {
	start := time.Now()
	input, _ := antlr.NewFileStream(filename)
	lexer := parser.NewAnnotationGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := parser.NewAnnotationGrammarParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Annotation()
	listener := parser.NewNluRuleListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	elapsed := time.Since(start)
	fmt.Println("TIME", elapsed.Seconds())
	parsed, _ := json.MarshalIndent(listener.ParsedRules, "", " ")
	return C.CString(string(parsed))
}

func main() {
	filename := os.Args[1]
	fmt.Println(run_parser(filename))
}
