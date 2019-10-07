package testcase

import (
	"bufio"
	"io"
	"io/ioutil"
)

type Testcase struct {
	Name    string
	Inputs  string
	Outputs string
}

func NewTestcase(name string, inputs string, outputs string) *Testcase {
	t := new(Testcase)
	t.Name = name
	t.Inputs = inputs
	t.Outputs = outputs

	return t
}

func readMultilineString(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	rv := ""
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

func (t *Testcase) ReadInput(r io.Reader) error {
	inputs, err := readMultilineString(r)
	if err != nil {
		return err
	}
	t.Inputs = inputs

	return nil
}

func (t *Testcase) ReadOutput(r io.Reader) error {
	outputs, err := readMultilineString(r)
	if err != nil {
		return err
	}
	t.Outputs = outputs

	return nil
}

func (t *Testcase) SaveTestcase() error {
	err := ioutil.WriteFile(t.Name+".in", []byte(t.Inputs), 0644)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(t.Name+".out", []byte(t.Outputs), 0644)
	if err != nil {
		return err
	}

	return nil
}
