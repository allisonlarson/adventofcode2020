package day3

import (
	"strings"
	"testing"
)

func TestCompute(t *testing.T) {
	rawInput := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	answer, err := Compute(strings.NewReader(rawInput), false)
	if err != nil {
		t.Error(err)
	}

	if answer != 7 {
		t.Errorf("Expected 7, got %v", answer)
	}

	answer, err = Compute(strings.NewReader(rawInput), true)
	if err != nil {
		t.Error(err)
	}

	if answer != 336 {
		t.Errorf("Expected 336, got %v", answer)
	}
}
