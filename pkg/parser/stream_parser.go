package parser

import (
	"context"
	"errors"

	"speechly/nlu-example-parser/internal/grammar"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type StreamParser struct {
	lis            *NluRuleListener
	ch             chan string
	done           chan struct{}
	debug, verbose bool
}

func NewStreamParser(bufSize uint64, debug bool, verbose bool) *StreamParser {
	return &StreamParser{
		lis:     NewNluRuleListener(bufSize, debug),
		ch:      make(chan string, bufSize),
		done:    make(chan struct{}),
		debug:   debug,
		verbose: verbose,
	}
}

func (s *StreamParser) Start(ctx context.Context) error {
	go func() {
		defer close(s.done)
		defer s.lis.Close()

		for {
			select {
			case <-ctx.Done():
				return
			case str, more := <-s.ch:
				if !more {
					return
				}

				stream := antlr.NewCommonTokenStream(
					parser.NewAnnotationGrammarLexer(
						antlr.NewInputStream(str),
					),
					0,
				)

				p := parser.NewAnnotationGrammarParser(stream)
				p.BuildParseTrees = true

				if s.debug {
					p.AddErrorListener(antlr.NewDiagnosticErrorListener(s.verbose))
				}

				antlr.ParseTreeWalkerDefault.Walk(s.lis, p.Annotation())
			}
		}
	}()

	return nil
}

func (s *StreamParser) Stop(ctx context.Context) error {
	close(s.ch)

	select {
	case <-s.done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *StreamParser) Write(ctx context.Context, line string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-s.done:
		return errors.New("cannot write to a stopped parser")
	case s.ch <- line:
		return nil
	}
}

func (s *StreamParser) Results() <-chan Utterance {
	return s.lis.Utterances()
}
