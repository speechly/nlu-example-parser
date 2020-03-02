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

	inputFileFlag := flag.String("input_file", "", "input file name")
	flag.Parse()

	f, err := os.Open(*inputFileFlag)
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(f)
	inputCh := make(chan string)

	go func() {
		defer close(inputCh)

		for scan.Scan() {
			line := scan.Text()
			inputCh <- line
		}

		if err := scan.Err(); err != nil {
			return
		}
	}()

	outCh := nlurules.ParseAsync(inputCh, false, false)

	for u := range outCh {

		out, err := json.Marshal(u)
			if err != nil {
				os.Exit(1)
			}
			fmt.Println(string(out))
	}

	os.Exit(0)
}
