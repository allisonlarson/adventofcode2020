package day1

import (
	"aoc/util"
	"fmt"
	"io"
)

func Compute(r io.Reader, isPartTwo bool) (int, error) {
	input, err := util.ParseInts(r)
	if err != nil {
		return 0, err
	}

	for i, num := range input {
		for _, num2 := range input[i+1:] {
			if isPartTwo {
				for _, num3 := range input[i+1:] {
					if num+num2+num3 == 2020 {
						return num * num2 * num3, nil
					}
				}
			} else {
				if num+num2 == 2020 {
					return num * num2, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("no answer found")
}
