package day1

import (
	"strings"
	"testing"
	"os"
)

func TestCompute(t *testing.T) {
	input := "1721\n979\n366\n299\n675\n1456"

	testCases := []struct {
		arg    string
		answer int
	}{
		{
			arg:    "part1",
			answer: 514579,
		},
		{
			arg:    "part2",
			answer: 241861950,
		},
	}

	for _, testCase := range testCases {
		day1, err := NewDay1(strings.NewReader(input), []string{testCase.arg})
		if err != nil {
			t.Error(err)
		}

		answer, err := day1.Compute()
		if err != nil {
			t.Error(err)
		}

		if answer != testCase.answer {
			t.Errorf("Expected %v, got %v", testCase.answer, answer)
		}
	}
}

func BenchmarkCompute(b *testing.B) {
	input, err := os.Open("../inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	day1, err := NewDay1(input, []string{"part1"})
	if err != nil {
		b.Error(err)
	}
	for i := 0; i <b.N; i++ {

		_, err = day1.Compute()
		if err != nil {
			b.Error(err)
		}
	}
}
