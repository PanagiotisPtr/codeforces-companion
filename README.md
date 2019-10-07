# Codeforces Companion

This is a tool that allows you to get the testcases, problems (as pdfs) and initial setup for your language of choice for a codeforces competition. I might add more features over time like support for different languages.

## Highlights
Problems for competitions are downloaded concurrently using go-routines. The code is separated into modules each of which is responsible for 1 thing. There are a some unit tests to ensure that the code works. Also uses templates to set up your initial code file.

## Overview of the tools
When you build the code you have 3 tools. The competition parser (`cfc-parser`) that allows you to parse a competition. The code tester, only works with C++ currently (`cfc-tester`). For those two programs you can also use `--help` to find out how they work. The last thing is `cfc-mktest` which allows you to quickly handcraft a testcase. It doesn't have `--help` but instead when run, it asks for user input with a prompt. 

## To build the code
Building the code is dead simple as it comes with a Makefile. Just do
```
make build
```

## Tests
I have written a few unit tests for this code but not that many. To run the tests for a package simply go to that package's folder and run `go test`
