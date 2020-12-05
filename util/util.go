package util

import (
	"bufio"
	"io"
	"strconv"
	"strings"
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

func ReadInput(rawInput io.Reader) (string, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, rawInput)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
