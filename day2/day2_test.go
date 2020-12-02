package day2

import (
	"testing"
	"strings"
)

func TestCompute(t *testing.T) {
	rawInput := "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"

	testCases := []struct{
		arg string
		answer int
	}{
		{
			arg: "part1",
			answer: 2,
		},
		{
			arg: "part2",
			answer: 1,
		},
	}

	for _, testCase := range testCases {
		input := strings.NewReader(rawInput)
		day2, err := NewDay(input, []string{testCase.arg})
		if err !=nil {
			t.Error(err)
		}
		answer,err := day2.Compute()
		if err !=nil {
			t.Error(err)
		}
		if answer != testCase.answer {
			t.Errorf("Expected %v, got %v", testCase.answer, answer)
		}
	}
}
