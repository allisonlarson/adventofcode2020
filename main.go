package main

import (
	"aoc/day1"
	"aoc/day2"
	"errors"
	"fmt"
	"io"
	"os"
)

type DayRunner interface {
	Compute() (int, error)
}

func getDayRunner(dayOption string, file io.Reader, args []string) (DayRunner, error) {
	switch dayOption {
	case "day1":
		return day1.NewDay1(file, os.Args[3:])
	case "day2":
		return day2.NewDay(file, os.Args[3:])
	}
	return nil, errors.New("Invalid option specified")
}

func main() {
	day := os.Args[1]
	input := os.Args[2]
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	dayRunner, err := getDayRunner(day, file, os.Args[3:])
	if err != nil {
		panic(err)
	}

	answer, err := dayRunner.Compute()
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}
