## Speechly NLU example parser

This repository contains the grammar definitions for Speechly NLU rules and a parser for these rules, written in Go.

The parser can be used as a standalone command-line client or included in other Go projects as a dependency.

### Usage

#### Command-line interface

In order to use the CLI you can build it using Make and run it against [some example rules](examples/test_multi_intent_data.md):

```sh
# Compile the binary
make build

# Run it against the examples
./bin/parser ./examples/test_multi_intent_data.md
```

#### As a library

You can use the streaming version of a parser in your Go program the following way:

```golang
package main

import (
	"context"
	"fmt"

	"speechly/nlu-rules-parser/pkg/nlurules"
)

const (
	bufSize = 10
	debug   = false
	verbose = false

	testUtterance = "*change change the [kitchen door    knobs](object)! *_ and *remove remove  the [table](object)"
)

func main() {
	ctx := context.Background()

	// Create a new parser.
	p := nlurules.NewStreamParser(bufSize, debug, verbose)

	// Start the parser.
	if err := p.Start(ctx); err != nil {
		panic(err)
	}

	go func() {
		// Stop the parser after finishing writing to it.
		defer p.Stop(ctx)

		for i := 0; i < 10; i++ {
			if err := p.Write(ctx, testUtterance); err != nil {
				panic(err)
			}
		}
	}()

	// Consume the results.
	for u := range p.Results() {
		fmt.Printf("Received utterance %+v\n", u)
	}
}
```
