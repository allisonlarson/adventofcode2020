package day5

import (
	"aoc/util"
	"io"
	"sort"
	"strconv"
	"strings"
)

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	seatIds := []int{}
	input, err := util.ReadInput(r)
	if err != nil {
		return 0, err
	}
	for _, s := range strings.Fields(input) {
		bin := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1").Replace(s)
		seatId, err := strconv.ParseUint(bin, 2, 10)
		if err != nil {
			return 0, err
		}

		seatIds = append(seatIds, int(seatId))
	}

	sort.Ints(seatIds)
	if runPartTwo {
		return seatIds[sort.Search(len(seatIds), func(i int) bool { return seatIds[i] != seatIds[0]+i })] - 1, nil
	}
	return seatIds[len(seatIds)-1], nil
}
