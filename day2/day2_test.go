package day2

import (
	"testing"
	"strings"
)

func TestCompute(t *testing.T) {
	rawInput := "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"

	testCases := []struct{
		runPartTwo bool
		answer int
	}{
		{
			answer: 2,
		},
		{
			runPartTwo: true,
			answer: 1,
		},
	}

	for _, testCase := range testCases {
		input := strings.NewReader(rawInput)
		answer,err := Compute(input, testCase.runPartTwo)
		if err !=nil {
			t.Error(err)
		}
		if answer != testCase.answer {
			t.Errorf("Expected %v, got %v", testCase.answer, answer)
		}
	}
}
