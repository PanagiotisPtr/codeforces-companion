package builder

import "testing"

func TestBuildCompetition(t *testing.T) {
	errors := BuildCompetition("https://codeforces.com/contest/1228")

	if len(errors) != 0 {
		t.Errorf("BuildCompetition(): Unexpected errors %v", errors)
	}
}
