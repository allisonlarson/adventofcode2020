package day10

import (
	"aoc/util"
	"io"
	"sort"
)

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	ones := 0
	threes := 0

	inputs, err := util.ParseInts(r)
	if err != nil {
		return 0, err
	}
	sort.Ints(inputs)
	inputs = append([]int{0}, inputs...)
	inputs = append(inputs, inputs[len(inputs)-1]+3)
	joltCount := 0
	for _, i := range inputs {
		count := i - joltCount
		if count == 1 {
			ones++
		}
		if count == 3 {
			threes++
		}
		joltCount = i
	}

	if runPartTwo {
		return FindAdapterCombo(inputs, 0, make(map[int]int)), nil
	}

	return ones * threes, nil
}

func FindAdapterCombo(adapters []int, index int, testedAdapters map[int]int) int {
	if index >= len(adapters)-3 {
		return 1
	}

	adapter := adapters[index]
	if tested, ok := testedAdapters[adapter]; ok {
		return tested
	}

	var count int
	for i := index + 1; i <= index+3; i++ {
		a := adapters[i]
		diff := a - adapter
		if diff > 0 && diff < 4 {
			count += FindAdapterCombo(adapters, i, testedAdapters)
		}
	}

	testedAdapters[adapter] = count
	return count
}
