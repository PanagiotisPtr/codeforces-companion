package parser

import (
	"sort"
	"strings"
	"testing"

	"github.com/panagiotisptr/codeforces-companion/pkg/testcase"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func TestGetProblemLinks(t *testing.T) {
	got, errors := GetProblemLinks("https://codeforces.com/contest/1228")

	if len(errors) != 0 {
		t.Errorf("GetProblemLinks(): Unexpected errors %v", errors)
	}

	want := []string{
		"https://codeforces.com//contest/1228/problem/A",
		"https://codeforces.com//contest/1228/problem/B",
		"https://codeforces.com//contest/1228/problem/C",
		"https://codeforces.com//contest/1228/problem/D",
		"https://codeforces.com//contest/1228/problem/E",
		"https://codeforces.com//contest/1228/problem/F",
	}
	sort.Strings(got)

	if len(got) != len(want) {
		t.Errorf("GetProblemLinks(): Expected number of links %d, got %d", len(want), len(got))
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("GetProblemLinks(): Unexpected link %s, wanted %s", got[i], want[i])
		}
	}
}

func TestGetTestCases(t *testing.T) {
	got, errors := GetTestCases("https://codeforces.com//contest/1228/problem/B")

	if len(errors) != 0 {
		t.Errorf("GetTestCases(): Returned errors %v", errors)
	}

	want := []testcase.Testcase{
		testcase.Testcase{
			Inputs:  `3 4 0 3 1 0 2 3 0`,
			Outputs: `2`,
		},
		testcase.Testcase{
			Inputs:  `1 1 0 1`,
			Outputs: `0`,
		},
		testcase.Testcase{
			Inputs:  `19 16 16 16 16 16 15 15 0 5 0 4 9 9 1 4 4 0 8 16 12 6 12 19 15 8 6 19 19 14 6 9 16 10 11 15 4`,
			Outputs: `797922655`,
		},
	}

	if len(got) != len(want) {
		t.Errorf("GetProblemLinks(): Expected number of test cases %d, got %d", len(got), len(want))
	}

	for i := 0; i < len(got); i++ {
		gotInputs := standardizeSpaces(got[i].Inputs)
		wantInputs := standardizeSpaces(want[i].Inputs)
		if gotInputs != wantInputs {
			t.Errorf("GetProblemLinks(): Expected inputs %s, got %s", wantInputs, gotInputs)
		}

		gotOutputs := standardizeSpaces(got[i].Outputs)
		wantOutputs := standardizeSpaces(want[i].Outputs)
		if gotOutputs != wantOutputs {
			t.Errorf("GetProblemLinks(): Expected outputs %s, got\n%s", wantOutputs, gotOutputs)
		}
	}
}

func TestGetProblemPdf(t *testing.T) {
	err := GetProblemPdf("https://codeforces.com//contest/1228/problem/B", "")

	if err != nil {
		t.Errorf("GetProblemPdf(): Unexpected error %v", err)
	}
}
