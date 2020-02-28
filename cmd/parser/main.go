package main

import (
	"bufio"
	"fmt"
	"os"
	"flag"
	"encoding/json"

	"speechly/nlu-rules-parser/pkg/nlurules"
)

func main() {

	fname := os.Args[1]

	jsonFlag := flag.Bool("-json", true, "return results as json")
	flag.Parse()

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
			if ! *jsonFlag {
				fmt.Printf("Read rule from file: %s\n", line)
			}

			inputCh <- line
		}

		if err := scan.Err(); err != nil {
			if ! *jsonFlag {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
			return
		}
	}()

	outCh := nlurules.ParseAsync(inputCh, false, false)

	for u := range outCh {
		if *jsonFlag {
		out, err := json.MarshalIndent(u, "", "  ")
			if err != nil {
				os.Exit(1)
			}
			fmt.Printf(string(out))
		} else {
			fmt.Printf("Received utterance %+v\n", u)
		}

	}

	os.Exit(0)
}
