package day1

import (
	"strings"
	"testing"
)

func TestCompute(t *testing.T) {
	input := "1721\n979\n366\n299\n675\n1456"

	testCases := []struct {
		runPartTwo bool
		answer int
	}{
		{
			answer: 514579,
		},
		{
			runPartTwo: true,
			answer: 241861950,
		},
	}

	for _, testCase := range testCases {
		answer, err := Compute(strings.NewReader(input), testCase.runPartTwo)
		if err != nil {
			t.Error(err)
		}

		if answer != testCase.answer {
			t.Errorf("Expected %v, got %v", testCase.answer, answer)
		}
	}
}
