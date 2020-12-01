package day1

import (
	"aoc/util"
	"fmt"
	"io"
)

type Day1 struct {
	inputs []int
	args   []string
}

func NewDay1(file io.Reader, args []string) (*Day1, error) {
	input, err := util.ParseInts(file)
	if err != nil {
		return nil, err
	}
	return &Day1{
		inputs: input,
		args:   args,
	}, nil
}

func (d *Day1) Compute() (int, error) {
	for i, num := range d.inputs {
		for _, num2 := range d.inputs[i+1:] {
			if len(d.args) != 0 && d.args[0] == "part2" {
				for _, num3 := range d.inputs[i+1:] {
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
