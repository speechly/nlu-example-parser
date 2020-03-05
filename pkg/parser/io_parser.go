package parser

import (
	"bytes"
	"encoding/json"
	"io"
	"sync"
)

const (
	newline    = '\n'
	newlineLen = 1
)

type IOParser struct {
	inputBuffer  []byte
	inputLock    sync.Mutex
	outputBuffer []byte
	outputLock   sync.Mutex
	parse        *Parser
}

func NewIOParser(bufSize uint64, parserBufSize uint64, debug bool) *IOParser {
	return &IOParser{
		inputBuffer:  make([]byte, 0, bufSize),
		outputBuffer: make([]byte, 0, bufSize),
		parse:        NewBufferedParser(parserBufSize, debug),
	}
}

func (p *IOParser) Read(b []byte) (int, error) {
	p.outputLock.Lock()
	defer p.outputLock.Unlock()

	var (
		buflen    = len(b)
		outbuflen = len(p.outputBuffer)
	)

	// If somehow we have exactly cap(b) in output buffer, just copy and return.
	if outbuflen == buflen {
		copy(b, p.outputBuffer)
		p.outputBuffer = p.outputBuffer[:0]
		return buflen, nil
	}

	// If we have more than cap(b), copy, resize and return.
	if outbuflen > buflen {
		copy(b, p.outputBuffer[:buflen])

		// Proper way to resize the slice without losing it's length
		copy(p.outputBuffer, p.outputBuffer[buflen:])
		p.outputBuffer = p.outputBuffer[:outbuflen-buflen]

		return buflen, nil
	}

	// Not enough data left in output buffer, so drain it unconditionally.
	copy(b[:outbuflen], p.outputBuffer)
	p.outputBuffer = p.outputBuffer[:0]

	// Remember how much we can write more
	written := outbuflen

	// Loop until we have either exhausted the parser OR the buffer.
	for {
		res, more := <-p.parse.Results()
		if !more {
			return written, io.EOF
		}

		data, err := json.Marshal(res)
		if err != nil {
			return written, err
		}

		var (
			datalen = len(data)
			capleft = buflen - written
		)

		// We have exactly enough data to fill in the rest of the buffer, so copy and return.
		if datalen == capleft {
			copy(b[written:], data)
			return buflen, nil
		}

		// We have more data to write than we can fit in the output.
		// Copy however many bytes we can from serialized result and store the rest in output buffer.
		if datalen > capleft {
			copy(b[written:], data[:capleft])
			p.outputBuffer = append(p.outputBuffer, data[capleft:]...)
			return buflen, nil
		}

		// We can fit more than we have available for writing.
		// Write as much as we can and continue to the next iteration.
		copy(b[written:], data)
		written = written + datalen
	}
}

func (p *IOParser) Write(b []byte) (int, error) {
	p.inputLock.Lock()
	defer p.inputLock.Unlock()

	written := 0

	for {
		buflen := len(b)

		// Quick check for empty buffer.
		if buflen == 0 {
			return written, nil
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
		// Re-slice the buffer to make sure we are progressively moving forward on it.
		head := b[:linesep]
		b = b[linesep+newlineLen:]

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
	}
}

func (p *IOParser) Close() error {
	p.parse.Close()
	return nil
}
