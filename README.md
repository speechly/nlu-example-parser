# Speechly NLU example parser

This repository contains the grammar definitions for Speechly NLU examples and an AST parser for these rules, written in Go.

The parser can be used as a standalone command-line client or included in other Go projects as a library.

## About Speechly

Speechly is a developer tool for building real-time multimodal voice user interfaces. It enables developers and designers to enhance their current touch user interface with voice functionalities for better user experience. Speechly key features:

#### Speechly key features

- Fully streaming API
- Multi modal from the ground up
- Easy to configure for any use case
- Fast to integrate to any touch screen application
- Supports natural corrections such as "Show me red – i mean blue t-shirts"
- Real time visual feedback encourages users to go on with their voice

| Example application | Description |
| :---: | --- |
| <img src="https://i.imgur.com/v9o1JHf.gif" width=50%> | Instead of using buttons, input fields and dropdowns, Speechly enables users to interact with the application by using voice. <br />User gets real-time visual feedback on the form as they speak and are encouraged to go on. If there's an error, the user can either correct it by using traditional touch user interface or by voice. |

## Usage

### CLI

Get the parser, a file with examples and run it:

```sh
nlu-example-parser -input_file_path nlu-examples.md > output.json
```

More CLI options available, you can tune the performance / memory usage or enable debug output:

```
Usage of nlu-example-parser:
  -input_file_path string
    path to input file (required)
  -buffer_size_mb uint
    size of the input and output buffers for input and output files (default 1)
  -parser_buffer_size_lines uint
    size of the buffer with parsed lines (default 100)
```

You can check out [some example rules](examples/test_multi_intent_data.md) and [sample output](test/golden.json).

#### Installation

Pre-built binaries are automatically compiled on every release - https://github.com/speechly/nlu-example-parser/releases, you can download them using e.g. `curl`:

```sh
# amd64 or 386
$ export ARCH="amd64"

# darwin, linux, freebsd
$ export PLATFORM="darwin"
$ export FILENAME="speechly-nluexamplesparser-${PLATFORM}-${ARCH}.tar.gz"

# If you want specific version
$ export VERSION="v0.2.0"
$ curl -LJO https://github.com/speechly/nlu-example-parser/releases/download/${VERSION}/${FILENAME}

# Or if you want just the latest
$ curl -LJO https://github.com/speechly/nlu-example-parser/releases/latest/download/${PARSER}

$ tar -xzf ${FILENAME}
```

Alternatively, you can build it yourself, make sure you have `make` and `go` installed. Minimum support Go version is `1.13`.

```sh
# Clone the repo
$ git clone git@github.com:speechly/nlu-example-parser.git
$ cd nlu-example-parser

# Compile the parser
$ make build

# Run it
$ ./bin/parser -input_file_path ./examples/test_multi_intent_data.md
```

### Go library

Get it using `go get`:

```sh
$ go get github.com/speechly/nlu-example-parser
```

The library exposes a couple of different parser API, you can choose which one you want to use.

#### Parser

The simplest version, you can use it the following way:

```golang
package main

import (
	"fmt"

	"github.com/speechly/nlu-example-parser/pkg/parser"
)

const (
	testUtterance = "*change change the [kitchen door    knobs](object)! *_ and *remove remove  the [table](object)"
)

func main() {
	// Create a new parser.
	p := parser.NewParser()

	// Consume the results.
	done := make(chan struct{})
	go func() {
		defer close(done)

		for u := range p.Results() {
			fmt.Printf("Received utterance %+v\n", u)
		}
	}()

	// Send some utterances.
	for i := 0; i < 10; i++ {
		p.Parse(testUtterance)
	}

	p.Close()
	<-done
}
```

#### StreamingParser

You can use the streaming version of a parser in your Go program the following way:

```golang
package main

import (
	"context"
	"fmt"

	"github.com/speechly/nlu-example-parser/pkg/parser"
)

const (
	bufSize = 10
	verbose = false

	testUtterance = "*change change the [kitchen door    knobs](object)! *_ and *remove remove  the [table](object)"
)

func main() {
	ctx := context.Background()

	// Create a new parser.
	p := parser.NewStreamParser(bufSize, verbose)

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

#### IOParser

You can use the `io.ReadWriteCloser` version of a parser in your Go program the following way:

```golang
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/speechly/nlu-example-parser/pkg/parser"
)

const (
	bufSize       = 10
	parserBufSize = 10

	testUtterance = "*change change the [kitchen door    knobs](object)! *_ and *remove remove  the [table](object)\n"
)

func main() {
	var wg sync.WaitGroup

	r := strings.NewReader(testUtterance)
	parser := parser.NewIOParser(bufSize, parserBufSize)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer parser.Close()

		if _, err := io.Copy(parser, r); err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "parser encountered an error: %s\n", err.Error())
			os.Exit(1)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if _, err := io.Copy(os.Stdout, parser); err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "parser encountered an error: %s\n", err.Error())
		}
	}()

	wg.Wait()
	os.Exit(0)
}
```
