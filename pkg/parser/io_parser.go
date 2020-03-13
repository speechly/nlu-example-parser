package parser

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"sync"

	"github.com/speechly/nlu-example-parser/internal/grammar"
)

const (
	newline    = '\n'
	newlineLen = 1
)

var errEmptyBuffer = errors.New("cannot handle an empty buffer")

type IOParser struct {
	inputBuffer  []byte
	inputLock    sync.Mutex
	outputBuffer []byte
	outputLock   sync.Mutex
	parse        *Parser
	bufSize      uint64
	errLock      sync.Mutex
	parseErr     grammar.ParseError
	ctx          context.Context
	cancel       context.CancelFunc
}

func NewIOParser(bufSize uint64, parserBufSize uint16) *IOParser {
	ctx, cancel := context.WithCancel(context.Background())
	p := &IOParser{
		inputBuffer:  make([]byte, 0, bufSize),
		outputBuffer: make([]byte, 0, bufSize),
		parse:        NewBufferedParser(parserBufSize),
		bufSize:      bufSize,
		ctx:          ctx,
		cancel:       cancel,
	}

	p.handleErrs()

	return p
}

func (p *IOParser) handleErrs() {
	go func() {
		for {
			select {
			case e, more := <-p.parse.Errors():
				if !more {
					return
				}

				p.saveErr(e)
			case <-p.ctx.Done():
				return
			}
		}
	}()
}

func (p *IOParser) Read(b []byte) (int, error) {
	if err := p.checkErr(); err != nil {
		return 0, err
	}

	p.outputLock.Lock()
	defer p.outputLock.Unlock()

	buflen := len(b)

	// Quick escape for empty buffers.
	if buflen == 0 {
		return 0, errEmptyBuffer
	}

	// Copy as much as we can from output buffer to b.
	// `copy` guarantees that `min(len(src), len(dst))` elements will be copied.
	written := copy(b, p.outputBuffer)

	// Proper way to re-slice output buffer, without losing its capacity.
	c := copy(p.outputBuffer, p.outputBuffer[written:])
	p.outputBuffer = p.outputBuffer[:c]

	if written == buflen {
		return written, nil
	}

	// A separate buffer for json serialization.
	// TODO: switch outputBuffer to bytes.Buffer and get rid of this temporary solution altogether.
	jsonbuf := bytes.NewBuffer(make([]byte, 0, p.bufSize))

	// Loop until we have either exhausted the parser OR the buffer.
	for {
		if err := p.checkErr(); err != nil {
			return written, err
		}

		res, more := <-p.parse.Results()
		if !more {
			return written, io.EOF
		}

		// Unlike `json.Marshal` `json.Encoder.Encode` appends a newline to encoded JSON, which is a requirement.
		if err := json.NewEncoder(jsonbuf).Encode(res); err != nil {
			return written, err
		}

		// Copy as much as we can from JSON buffer to b.
		// `copy` guarantees that `min(len(src), len(dst))` elements will be copied.
		data := jsonbuf.Bytes()

		c := copy(b[written:], data)

		// Store any non-copied data in output buffer.
		if c < len(data) {
			p.outputBuffer = append(p.outputBuffer, data[c:]...)
		}

		// Update how much we written.
		jsonbuf.Reset()
		written += c
		if written == buflen {
			return written, nil
		}
	}
}

func (p *IOParser) Write(b []byte) (int, error) {
	if err := p.checkErr(); err != nil {
		return 0, err
	}

	p.inputLock.Lock()
	defer p.inputLock.Unlock()

	var (
		buflen  = len(b)
		written = 0
	)

	// Quick check for empty buffer.
	if buflen == 0 {
		return written, errEmptyBuffer
	}

	for {
		if err := p.checkErr(); err != nil {
			return written, err
		}

		// Check if incoming buffer contains a newline,
		// since it indicates that we can now send the utterance to AST parser.
		linesep := bytes.IndexRune(b, newline)

		// If incoming bytestring doesn't contain a newline,
		// then just append it to our "current line" buffer and return.
		if linesep == -1 {
			p.inputBuffer = append(p.inputBuffer, b...)
			return written + buflen, nil
		}

		// Split the buffer by the newline character.
		// Head contains the bytes [0 : newlineIndex]
		head := b[:linesep]

		// Check if we need to append input buffer to the head to get the full line.
		var line string
		if len(p.inputBuffer) == 0 {
			line = string(head)
		} else {
			// Update input buffer, get the line to be parsed and zero out the input buffer.
			p.inputBuffer = append(p.inputBuffer, head...)
			line = string(p.inputBuffer)
			p.inputBuffer = p.inputBuffer[:0]
		}

		p.parse.Parse(line)
		written += len(head) + newlineLen

		// Re-slice the buffer to make sure we are progressively moving forward on it.
		b = b[linesep+newlineLen:]
		buflen = len(b)

		if buflen == 0 {
			return written, nil
		}
	}
}

func (p *IOParser) Close() error {
	p.inputLock.Lock()
	defer p.inputLock.Unlock()

	// Calling `Close` signifies the EOF,
	// so flush whatever is left in the input buffer to the parser,
	// and then zero out the input buffer.
	if len(p.inputBuffer) != 0 {
		p.parse.Parse(string(p.inputBuffer))
		p.inputBuffer = p.inputBuffer[:0]
	}

	p.parse.Close()
	p.cancel()

	return nil
}

func (p *IOParser) checkErr() error {
	p.errLock.Lock()
	defer p.errLock.Unlock()

	if p.parseErr.Len() > 0 {
		return p.parseErr
	}

	return nil
}

func (p *IOParser) saveErr(e grammar.ParseError) {
	p.errLock.Lock()
	defer p.errLock.Unlock()

	p.parseErr.Append(e)
}
