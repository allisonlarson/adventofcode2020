package day3

import (
	"os"
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

func BenchmarkComputePart1(b *testing.B) {
	file, err := os.Open("../inputs/day3.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_, err := Compute(file, false)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComputePart2(b *testing.B) {
	file, err := os.Open("../inputs/day3.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_, err = Compute(file, true)
		if err != nil {
			b.Error(err)
		}
	}
}
