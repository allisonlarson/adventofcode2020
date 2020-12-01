package util

import (
	"bufio"
	"io"
	"strconv"
)

func ParseInts(rawInput io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(rawInput)
	var result []int
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, input)
	}
	return result, scanner.Err()
}
