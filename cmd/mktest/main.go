package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/panagiotisptr/codeforces-companion/pkg/testcase"
	"github.com/panagiotisptr/codeforces-companion/pkg/tester"
)

func existsInSlice(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}

func main() {
	var ts testcase.Testcase

	fmt.Println("Testcase Inputs")
	err := ts.ReadInput(os.Stdin)
	for err != nil {
		fmt.Printf("An error occured while reading testcase inputs: %v", err)
		err = ts.ReadInput(os.Stdin)
	}

	fmt.Println("Testcase Outputs")
	err = ts.ReadOutput(os.Stdin)
	for err != nil {
		fmt.Printf("An error occured while reading testcase inputs: %v", err)
		err = ts.ReadOutput(os.Stdin)
	}

	testcases, _, _ := tester.GetTestcases()
	numTestCases := len(testcases)
	fmt.Println("testcases:", testcases)
	testName := "testcase_" + strconv.Itoa(numTestCases)
	for existsInSlice(testcases, testName+".in") {
		numTestCases++
		testName = "testcase_" + strconv.Itoa(numTestCases)
	}

	ts.Name = testName
	err = ts.SaveTestcase()
	if err != nil {
		fmt.Printf("Error occured when saving test case: %v", err)
	}
}
