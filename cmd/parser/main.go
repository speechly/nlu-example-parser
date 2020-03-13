package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/speechly/nlu-example-parser/pkg/parser"

	"golang.org/x/sync/errgroup"
)

const (
	bytesInMB = 1024 * 1024
)

var (
	inputFileFlag     = flag.String("input_file_path", "", "path to input file")
	bufSizeFlag       = flag.Uint64("buffer_size_mb", 1, "size of the input and output buffers for input and output files")
	parserBufSizeFlag = flag.Uint("parser_buffer_size_lines", 100, "size of the buffer with parsed lines")
)

func main() {
	flag.Parse()

	filename := *inputFileFlag
	if filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	bufSize := *bufSizeFlag * bytesInMB
	if bufSize < 1 {
		flag.Usage()
		os.Exit(1)
	}

	parserBufSize := *parserBufSizeFlag
	if bufSize < 1 {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		os.Exit(printErr(err))
	}

	parser := parser.NewIOParser(bufSize, uint16(parserBufSize))

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		defer parser.Close() // nolint: errcheck
		defer f.Close()      // nolint: errcheck

		var (
			buf  = make([]byte, bufSize)
			done = false
		)

		for {
			n, err := f.Read(buf)
			if err != nil {
				if err != io.EOF {
					return err
				}

				if n == 0 {
					return nil
				}

				done = true
			}

			if _, err := parser.Write(buf[:n]); err != nil {
				return err
			}

			if done {
				return nil
			}
		}
	})

	g.Go(func() error {
		var (
			buf  = make([]byte, bufSize)
			done = false
		)

		for {
			n, err := parser.Read(buf)
			if err != nil {
				if err != io.EOF {
					return err
				}

				if n == 0 {
					return nil
				}

				done = true
			}

			if _, err := os.Stdout.Write(buf[:n]); err != nil {
				return err
			}

			if done {
				return nil
			}
		}
	})

	if err := g.Wait(); err != nil {
		os.Exit(printErr(err))
	}

	os.Exit(0)
}

func printErr(err error) int {
	for _, l := range strings.Split(err.Error(), "\n") {
		fmt.Fprintf(os.Stderr, "{\"error\":\"%s\"}\n", l)
	}

	return 1
}
