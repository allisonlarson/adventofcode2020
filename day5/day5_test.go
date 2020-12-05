package day5

import (
	"strings"
	"testing"
)

func TestComputePart1(t *testing.T) {
	rawInput := `
	FBFBBFFRLR
	BFFFBBFRRR
	FFFBBBFRRR
	BBFFBBFRLL`
	answer, err := Compute(strings.NewReader(rawInput), false)
	if err != nil {
		t.Error(err)
	}
	if answer != 821 {
		t.Errorf("Expected 820, got %v", answer)
	}
}
