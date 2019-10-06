package tester

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"

	"github.com/fatih/color"
)

func CompileCpp(filename string) (string, error) {
	name := getFilename(removeExtension(filename))
	executableName := name + ".tmp.out"
	cmd := exec.Command("g++", "-std=c++14", filename, "-o", executableName)
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return executableName, nil
}

func RunProgram(programName string, programInput string) (string, error) {
	cmd := exec.Command("./" + programName)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	inputString, err := ioutil.ReadFile(programInput)
	if err != nil {
		return "", err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, string(inputString))
	}()

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	res, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func GetTestcases() ([]string, []string, error) {
	var inputs []string
	var outputs []string
	dirname := "." + string(filepath.Separator)
	d, err := os.Open(dirname)
	if err != nil {
		return inputs, outputs, err
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		return inputs, outputs, err
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".in" {
				outputName := removeExtension(file.Name()) + ".out"
				if _, err := os.Stat(outputName); err == nil {
					outputs = append(outputs, outputName)
					inputs = append(inputs, file.Name())
				}
			}
		}
	}

	return inputs, outputs, nil
}

func CheckOutput(programName string, input string, output string) error {
	programOutput, err := RunProgram(programName, input)
	if err != nil {
		return err
	}

	answerBytes, err := ioutil.ReadFile(output)
	if err != nil {
		return err
	}
	answer := string(answerBytes)

	answer = removeUnnecessaryEndlines(answer)
	programOutput = removeUnnecessaryEndlines(programOutput)

	if programOutput != answer {
		color.Red("[FAIL] %s", input)
		color.Red("Expected:\n%s\nGot:\n%s", answer, programOutput)
	} else {
		color.Green("[PASS] %s", input)
	}

	return nil
}

func TestCpp(filename string) error {
	programName, err := CompileCpp(filename)
	if err != nil {
		return err
	}

	testcasesIn, testcasesOut, err := GetTestcases()
	if err != nil {
		return err // probably need to defer cleanup in this case
	}

	sort.Strings(testcasesIn)
	sort.Strings(testcasesOut)

	for i, _ := range testcasesIn {
		err := CheckOutput(programName, testcasesIn[i], testcasesOut[i])
		if err != nil {
			return err
		}
	}

	err = os.Remove(programName)
	if err != nil {
		return err
	}

	return nil
}

func removeExtension(s string) string {
	stop := 0
	for index, char := range s {
		if char == '.' {
			stop = index
		}
	}

	return s[:stop]
}

// Helpers

func getFilename(s string) string {
	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] == '/' {
			break
		}
	}

	if i < 0 {
		return s
	}

	return s[i+1 : len(s)]
}

func removeUnnecessaryEndlines(s string) string {
	for len(s) > 0 && s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}

	return s
}
