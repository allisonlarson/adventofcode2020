package day9

import (
	"strings"
	"testing"
)

func TestComputePart1(t *testing.T) {
	rawInput := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

	answer, err := Compute(strings.NewReader(rawInput), false)
	if err != nil {
		t.Error(err)
	}
	if answer != 127 {
		t.Errorf("Expected 127, got %v", answer)
	}
}

func TestComputePart2(t *testing.T) {
	rawInput := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

	answer, err := Compute(strings.NewReader(rawInput), true)
	if err != nil {
		t.Error(err)
	}
	if answer != 62 {
		t.Errorf("Expected 62, got %v", answer)
	}
}
