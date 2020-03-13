package grammar

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ParseError struct {
	buf   strings.Builder
	count uint64
}

func (p *ParseError) Append(e ParseError) {
	var s string
	if p.count > 0 {
		s = "\n" + e.Error()
	} else {
		s = e.Error()
	}

	p.buf.WriteString(s)
	p.count++
}

func (p ParseError) Len() uint64 {
	return p.count
}

func (p ParseError) Error() string {
	return p.buf.String()
}

func NewParseError(s string) ParseError {
	var b strings.Builder
	b.WriteString(s)

	return ParseError{
		buf: b,
	}
}

type NLUExampleErrorListener struct {
	errs    chan ParseError
	verbose bool
}

func (l *NLUExampleErrorListener) SyntaxError(r antlr.Recognizer, sym interface{}, line, col int, msg string, e antlr.RecognitionException) {
	l.errs <- NewParseError(fmt.Sprintf("SyntaxError: %s %d:%d", msg, line, col))
}

func (l *NLUExampleErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	if !l.verbose {
		return
	}

	l.errs <- NewParseError(fmt.Sprintf("AmbiguityError: %d:%d", startIndex, stopIndex))
}

func (l *NLUExampleErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	if !l.verbose {
		return
	}

	l.errs <- NewParseError(fmt.Sprintf("ContextError: %d:%d", startIndex, stopIndex))
}

func (l *NLUExampleErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	if !l.verbose {
		return
	}

	l.errs <- NewParseError(fmt.Sprintf("ContextError: %d:%d", startIndex, stopIndex))
}

func (l *NLUExampleErrorListener) Errors() <-chan ParseError {
	return l.errs
}

func (l *NLUExampleErrorListener) Close() {
	close(l.errs)
}

func NewNLUExampleErrorListener(bufSize uint16, verbose bool) *NLUExampleErrorListener {
	return &NLUExampleErrorListener{
		errs: make(chan ParseError, bufSize),
	}
}
