package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/panagiotisptr/codeforces-companion/pkg/builder"
)

func main() {
	var competitionUrl string
	fs := flag.NewFlagSet("", flag.ExitOnError)
	fs.StringVar(&competitionUrl, "competitionUrl", "", "Url to the codeforces competition")
	fs.Parse(os.Args[1:])
	fmt.Println("Parsing competition")
	builder.BuildCompetition(competitionUrl)
	fmt.Println("Done")
}
