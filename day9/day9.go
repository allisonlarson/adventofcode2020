package day9

import (
	"aoc/util"
	"errors"
	"io"
	"sort"
)

const Preamble = 25

//const Preamble = 5

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	ints, err := util.ParseInts(r)
	if err != nil {
		return 0, err
	}
	workingInts := []int{}
	for _, i := range ints {
		if len(workingInts) != Preamble {
			workingInts = append(workingInts, i)
			continue
		}

		var valid bool
		for j, w1 := range workingInts {
			for _, w2 := range workingInts[j+1:] {
				if w1+w2 == i {
					valid = true
				}
			}
		}
		if !valid {
			if runPartTwo {
				return ComputePartTwo(i, ints)
			}
			return i, nil
		}

		workingInts = workingInts[1:]
		workingInts = append(workingInts, i)
	}
	return 0, errors.New("No solution found")
}

func ComputePartTwo(target int, ints []int) (int, error) {
	for i := 2; i < len(ints); i++ {
		answer, err := RunPartTwo(ints, i, target)
		if err == nil {
			return answer, nil
		}
	}

	return 0, errors.New("Really no answer")
}

func RunPartTwo(ints []int, comb, target int) (int, error) {
	for i := 0; i < len(ints); i++ {
		list := []int{}
		total := 0
		for j := 0; j < comb; j++ {
			index := i + j
			if index > len(ints)-1 {
				break
			}
			list = append(list, ints[i+j])
			total += ints[i+j]
		}

		if total == target {
			sort.Ints(list)
			return list[0] + list[len(list)-1], nil
		}
	}
	return 0, errors.New("No solution")
}
