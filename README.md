## Speechly NLU rules parser

This repository contains the grammar definitions for Speechly NLU rules and a parser for these rules, written in Go.

The parser can be used as a standalone command-line client or included in other Go projects as a dependency.

### Usage

In order to use the CLI you can build it using Make and run it against [some example rules](examples/test_multi_intent_data.md):

```sh
# Compile the binary
make build

# Run it against the examples
./bin/parser ./examples/test_multi_intent_data.md
```
