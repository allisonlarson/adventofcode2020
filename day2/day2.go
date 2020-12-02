package day2

import (
	"io"
	"bufio"
	"strings"
	"strconv"
)
type Day struct {
	inputs io.Reader
	args []string
}

func NewDay(file io.Reader, args []string) (*Day,error) {
	return &Day{inputs: file, args: args},nil
}

func (d *Day) Compute() (int, error) {
	result := 0
	part2 := (len(d.args) != 0) && (d.args[0] == "part2")
	scanner := bufio.NewScanner(d.inputs)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		nums := strings.Split(input[0], "-")

		first, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		second, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, err
		}
		val := input[1][:len(input[1])-1]
		pw := input[2]

		if part2 {
			chars := strings.Split(pw, "")
			if (chars[first-1] == val) != (chars[second-1] == val) { // sad XOR
				result += 1
			}
		} else {
			occurances := strings.Count(pw, val)
			if occurances >= first && occurances <= second {
				result += 1
			}
		}
	}
	return result, scanner.Err()
}
