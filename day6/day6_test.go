package day6

import (
	"strings"
	"testing"
)

func TestComputePart1(t *testing.T) {
	rawInput := `abc

a
b
c

ab
ac

a
a
a
a

b`
	answer, err := Compute(strings.NewReader(rawInput), false)
	if err != nil {
		t.Error(err)
	}
	if answer != 11 {
		t.Errorf("Expected 11, got %v", answer)
	}
}

func TestComputePart2(t *testing.T) {
	rawInput := `abc

a
b
c

ab
ac

a
a
a
a

b`
	answer, err := Compute(strings.NewReader(rawInput), true)
	if err != nil {
		t.Error(err)
	}
	if answer != 6 {
		t.Errorf("Expected 6, got %v", answer)
	}
}
