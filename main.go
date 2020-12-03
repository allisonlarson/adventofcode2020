package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"errors"
	"fmt"
	"io"
	"os"
)

type ComputeFn func(io.Reader, bool) (int, error)

func getDayRunner(dayOption string) (ComputeFn, error) {
	switch dayOption {
	case "day1":
		return day1.Compute, nil
	case "day2":
		return day2.Compute, nil
	case "day3":
		return day3.Compute, nil
	default:
		return nil, errors.New("Invalid option specified")
	}
}

func main() {
	day := os.Args[1]
	input := os.Args[2]
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	runPartTwo := (len(os.Args[3:]) != 0) && (os.Args[3] == "part2")

	computeFn, err := getDayRunner(day)
	if err != nil {
		panic(err)
	}

	answer, err := computeFn(file, runPartTwo)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}
