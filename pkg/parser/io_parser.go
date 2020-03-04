package parser

import (
	"bytes"
	"encoding/json"
	"io"
	"sync"
)

const newlineByte = byte('\n')

type IOParser struct {
	inputBuffer  *bytes.Buffer
	inputLock    sync.Mutex
	outputBuffer *bytes.Buffer
	outputLock   sync.Mutex
	parse        *Parser
}

func NewIOParser(bufSize uint64, parserBufSize uint64, debug bool) *IOParser {
	return &IOParser{
		outputBuffer: bytes.NewBuffer(make([]byte, 0, bufSize)),
		inputBuffer:  bytes.NewBuffer(make([]byte, 0, bufSize)),
		parse:        NewBufferedParser(parserBufSize, debug),
	}
}

func (p *IOParser) Read(b []byte) (int, error) {
	p.outputLock.Lock()
	defer p.outputLock.Unlock()

	// Loop until we have either
	// A) exhausted the parser AND the output buffer
	// B) exhausted the space in b
	var written int
	for {
		// Try to fill b from the buffer.
		n, err := p.outputBuffer.Read(b[written:])
		written += n

		// If we filled b from the buffer
		// AND it was an EOF,
		// OR if we encountered an unexpected error, return both n and err.
		if (err != nil && err != io.EOF) || (err == io.EOF && written == len(b)) {
			return written, err
		}

		// If we filled b, simply return.
		if written == len(b) {
			return written, nil
		}

		res, more := <-p.parse.Results()
		if !more {
			return written, io.EOF
		}

		data, err := json.Marshal(res)
		if err != nil {
			return written, err
		}

		// Write JSON data into output buffer.
		// Write always return a nil error and n = len(b),
		// but we don't need either, so just discard them.
		_, _ = p.outputBuffer.Write(data)
	}
}

func (p *IOParser) Write(b []byte) (int, error) {
	p.inputLock.Lock()
	defer p.inputLock.Unlock()

	// Write always return a nil error and n = len(b).
	n, _ := p.inputBuffer.Write(b)

	// Try to read as many lines as we can and send them to the parser.
	for {
		line, err := p.inputBuffer.ReadString(newlineByte)
		if err != nil {
			// No more lines, return.
			if err == io.EOF {
				return n, nil
			}

			// Something didn't go well, return an error.
			// Since the data is in the buffer, return `n` as the amount of bytes written.
			return n, err
		}

		p.parse.Parse(line)
	}
}

func (p *IOParser) Close() error {
	p.parse.Close()
	return nil
}
