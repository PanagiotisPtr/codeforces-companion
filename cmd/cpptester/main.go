package main

import (
	"os"

	"github.com/panagiotisptr/codeforces-companion/pkg/tester"
)

func main() {
	filename := os.Args[1]
	tester.TestCpp(filename)
}
