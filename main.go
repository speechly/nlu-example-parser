//main.go
package main

import (
	"./parser"
	"os"
	"fmt"
)

func main() {
	filename := os.Args[1]
	fmt.Println(parser.RunParser(filename))
}
