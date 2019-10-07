# Codeforces Companion

This is a tool that allows you to get the testcases, problems (as pdfs) and initial setup for your language of choice for a codeforces competition. More features will be added soon like support for different languages, better testing, concurrent downloads etc. This project is still in development

## To build the code
Building the code is dead simple as it comes with a Makefile. Just do
```
make build
```

## Tests
Although some tests exist for now they run with `go test` but will be added on the Makefile soon so you should be able to do `make tests` soon enough...

## Using the binaries
For now the binaries will be in /bin at the root of the project. For the parser just pass the link to the contest as the first argument and for the cpptester pass the name of your filename (need to run the cpptester inside the directory that you have your testcase files). This is still work in progress. Command line arguments and help will be added soon.
