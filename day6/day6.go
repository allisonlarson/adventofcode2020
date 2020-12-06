package day6

import (
	"aoc/util"
	"io"
	"strings"
)

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	input, err := util.ReadInput(r)
	if err != nil {
		return 0, err
	}
	totalCount := 0
	for _, s := range strings.Split(input, "\n\n") {
		yesQuestions := map[rune]int{}
		people := strings.Fields(s)
		questions := strings.ReplaceAll(s, "\n", "")
		for _, s := range questions {
			_, ok := yesQuestions[s]
			if ok {
				yesQuestions[s]++
			} else {
				yesQuestions[s] = 1
			}
		}

		for _, v := range yesQuestions {
			if runPartTwo {
				if v == len(people) {
					totalCount++
				}
			} else {
				totalCount++
			}
		}
	}
	return totalCount, nil
}
