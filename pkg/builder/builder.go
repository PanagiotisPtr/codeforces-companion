package builder

import (
	"os"
	"strconv"
	"strings"

	"github.com/panagiotisptr/codeforces-parser/pkg/parser"
)

func BuildTestCases(problemLink string, rootDir string) []error {
	testCases, errors := parser.GetTestCases(problemLink)

	if len(errors) != 0 {
		return errors
	}

	for testCaseNum, testCase := range testCases {
		testCaseName := rootDir + "testcase_" + strconv.Itoa(testCaseNum+1)
		fin, err := os.Create(testCaseName + ".in")
		if err != nil {
			errors = append(errors, err)
		} else {
			fin.WriteString(testCase.Inputs)
		}

		fout, err := os.Create(testCaseName + ".out")
		if err != nil {
			errors = append(errors, err)
		} else {
			fout.WriteString(testCase.Outputs)
		}

		fin.Close()
		fout.Close()
	}

	return nil
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

	for _, problemLink := range problemLinks {
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
		if err != nil {
			err = parser.GetProblemPdf(problemLink, problemFolder)
		}
	}

	return errors
}
