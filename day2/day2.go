package day2

import (
	"io"
	"bufio"
	"strings"
	"strconv"
)

func Compute(r io.Reader, isPart2 bool) (int, error) {
	result := 0
	scanner := bufio.NewScanner(r)
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

		if isPart2 {
			chars := strings.Split(pw, "")
			if (chars[first-1] == val) != (chars[second-1] == val) {
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
