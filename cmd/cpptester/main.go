package main

import (
	"flag"
	"os"

	"github.com/panagiotisptr/codeforces-companion/pkg/tester"
)

func main() {
	var filename string
	fs := flag.NewFlagSet("solution.cpp", flag.ExitOnError)
	fs.StringVar(&filename, "filename", "solution.cpp", "C++ file to test")
	fs.Parse(os.Args[1:])
	tester.TestCpp(filename)
}
