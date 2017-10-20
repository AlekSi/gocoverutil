package main

import (
	"flag"

	"github.com/AlekSi/gocoverutil/internal/test/package2"
)

var vF = flag.Bool("v", false, "Verbose output: print numbers before words")

func main() {
	package2.FizzBuzz(*vF)
}
