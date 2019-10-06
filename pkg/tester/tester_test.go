package tester

import (
	"fmt"
	"testing"
)

func TestTestCpp(t *testing.T) {
	err := TestCpp("../solution.cpp")
	if err != nil {
		fmt.Println(err)
	}
}
