package main

import (
	"os"

	"github.com/panagiotisptr/codeforces-parser/pkg/tester"
)

func main() {
	filename := os.Args[1]
	tester.TestCpp(filename)
}
