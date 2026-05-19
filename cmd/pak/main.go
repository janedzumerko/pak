package main

import (
	"fmt"
	"os"

	"github.com/janedzumerko/pak/internal/gitx"
)

func main() {
	args := os.Args[1:]

	code, err := gitx.Passthrough(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pak: %v\n", err)
		os.Exit(1)
	}

	os.Exit(code)
}
