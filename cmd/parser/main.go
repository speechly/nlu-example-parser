package main

import (
	"bufio"
	"fmt"
	"os"

	"speechly/nlu-rules-parser/pkg/nlurules"
)

func main() {
	fname := os.Args[1]

	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(f)
	inputCh := make(chan string)

	go func() {
		defer close(inputCh)

		for scan.Scan() {
			line := scan.Text()
			fmt.Printf("Read rule from file: %s\n", line)
			inputCh <- line
		}

		if err := scan.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			return
		}
	}()

	outCh := nlurules.ParseAsync(inputCh)

	for u := range outCh {
		fmt.Printf("Received utterance %+v\n", u)
	}

	os.Exit(0)
}
