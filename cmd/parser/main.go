package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"speechly/nlu-example-parser/pkg/parser"
)

var (
	inputFileFlag = flag.String("input_file", "", "input file name")
	bufSizeFlag   = flag.Uint64("bufsize", 100, "size of the buffer with input file lines")
	debugFlag     = flag.Bool("debug", false, "enable debug output")
)

func main() {
	flag.Parse()

	filename := *inputFileFlag
	if filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	bufSize := *bufSizeFlag
	if bufSize < 1 {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer f.Close() // nolint: errcheck

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	scan := bufio.NewScanner(f)
	parser := parser.NewStreamParser(bufSize, *debugFlag, *debugFlag)

	go func() {
		if err := parser.Start(ctx); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		defer parser.Stop(ctx) // nolint: errcheck

		for scan.Scan() {
			line := scan.Text()

			if err := parser.Write(ctx, line); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		}

		if err := scan.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}()

	for u := range parser.Results() {
		out, err := json.Marshal(u)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Fprintln(os.Stdout, string(out))
	}

	os.Exit(0)
}
