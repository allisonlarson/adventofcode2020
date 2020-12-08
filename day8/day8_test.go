package day8

import (
	"strings"
	"testing"
)

func TestComputePart1(t *testing.T) {
	rawInput := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

	answer, err := Compute(strings.NewReader(rawInput), false)
	if err != nil {
		t.Error(err)
	}
	if answer != 5 {
		t.Errorf("Expected 5, got %v", answer)
	}
}

func TestComputePart2(t *testing.T) {
	rawInput := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

	answer, err := Compute(strings.NewReader(rawInput), true)
	if err != nil {
		t.Error(err)
	}
	if answer != 8 {
		t.Errorf("Expected 8, got %v", answer)
	}
}
