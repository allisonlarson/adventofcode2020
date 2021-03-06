package main

import (
	"aoc/day1"
	"aoc/day10"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"aoc/day5"
	"aoc/day6"
	"aoc/day7"
	"aoc/day8"
	"aoc/day9"
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
	case "day4":
		return day4.Compute, nil
	case "day5":
		return day5.Compute, nil
	case "day6":
		return day6.Compute, nil
	case "day7":
		return day7.Compute, nil
	case "day8":
		return day8.Compute, nil
	case "day9":
		return day9.Compute, nil
	case "day10":
		return day10.Compute, nil
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
