package builder

import (
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/panagiotisptr/codeforces-companion/pkg/parser"
)

func BuildTestCases(problemLink string, rootDir string) []error {
	testCases, errors := parser.GetTestCases(problemLink)

	if len(errors) != 0 {
		return errors
	}

	for testCaseNum, testCase := range testCases {
		testCaseName := rootDir + "testcase_" + strconv.Itoa(testCaseNum+1)
		testCase.Name = testCaseName
		err := testCase.SaveTestcase()
		if err != nil {
			errors = append(errors, err)
		}
	}

	return nil
}

func parseProblem(problemLink string, competitionName string) []error {
	var errors []error
	linkPaths := strings.Split(problemLink, "/")
	problemName := linkPaths[len(linkPaths)-1]
	problemFolder := competitionName + "/" + problemName + "/"

	err := os.MkdirAll(problemFolder, os.ModePerm)
	if err != nil {
		err = parser.GetProblemPdf(problemLink, problemFolder)
	}

	testCaseErrors := BuildTestCases(problemLink, problemFolder)
	err = parser.GetProblemPdf(problemLink, problemFolder)

	if len(testCaseErrors) > 0 {
		errors = append(errors, testCaseErrors...)
	}

	return errors
}

func BuildCompetition(competitionUrl string) []error {
	var errors []error

	linkPaths := strings.Split(competitionUrl, "/")
	competitionName := "codeforces_" + linkPaths[len(linkPaths)-2] + "_" + linkPaths[len(linkPaths)-1]

	err := os.MkdirAll(competitionName, os.ModePerm)

	if err != nil {
		return []error{err}
	}

	problemLinks, errors := parser.GetProblemLinks(competitionUrl)

	if len(errors) != 0 {
		return errors
	}

	var wg sync.WaitGroup
	for _, problemLink := range problemLinks {
		wg.Add(1)
		go func(pl string, cn string) {
			defer wg.Done()
			parseProblem(pl, cn)
		}(problemLink, competitionName)
	}

	wg.Wait()

	return errors
}
