package parser

import (
	"context"
	"errors"

	"github.com/speechly/nlu-example-parser/internal/grammar"
)

type StreamParser struct {
	parse          *Parser
	ch             chan string
	done           chan struct{}
	debug, verbose bool
}

func NewStreamParser(bufSize uint64, debug bool, verbose bool) *StreamParser {
	return &StreamParser{
		parse:   NewParser(),
		ch:      make(chan string, bufSize),
		done:    make(chan struct{}),
		debug:   debug,
		verbose: verbose,
	}
}

func (s *StreamParser) Start(ctx context.Context) error {
	go func() {
		defer close(s.done)
		defer s.parse.Close()

		for {
			select {
			case <-ctx.Done():
				return
			case line, more := <-s.ch:
				if !more {
					return
				}

				s.parse.Parse(line)
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

func (s *StreamParser) Results() <-chan grammar.Utterance {
	return s.parse.Results()
}
