package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/panagiotisptr/codeforces-companion/pkg/tester"
)

func readMultilineString(query string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	rv := ""
	fmt.Println(query)
	for true {
		text, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		} else {
			if len(text) == 0 {
				break
			} else if text[0] == '\n' {
				break
			} else {
				rv += text
			}
		}
	}

	return rv, nil
}

func SaveTestCase(inputs string, outputs string, name string) error {
	err := ioutil.WriteFile(name+".in", []byte(inputs), 0644)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(name+".out", []byte(outputs), 0644)
	if err != nil {
		return err
	}

	return nil
}

func existsInSlice(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}

func main() {
	testIn, err := readMultilineString("Testcase Inputs")
	for err != nil {
		fmt.Printf("An error occured while reading testcase inputs: %v", err)
		testIn, err = readMultilineString("Testcase Inputs")
	}

	testOut, err := readMultilineString("Testcase Outputs")
	for err != nil {
		fmt.Printf("An error occured while reading testcase inputs: %v", err)
		testOut, err = readMultilineString("Testcase Inputs")
	}

	fmt.Println("Input:\n" + testIn)
	fmt.Println("Output:\n" + testOut)

	testcases, _, _ := tester.GetTestcases()
	numTestCases := len(testcases)
	fmt.Println("testcases:", testcases)
	testName := "testcase_" + strconv.Itoa(numTestCases)
	for existsInSlice(testcases, testName+".in") {
		numTestCases++
		testName = "testcase_" + strconv.Itoa(numTestCases)
	}

	fmt.Println(testName)
	err = SaveTestCase(testIn, testOut, testName)
	if err != nil {
		fmt.Printf("Error occured when saving test case: %v", err)
	}
}
