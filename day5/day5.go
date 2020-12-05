package day5

import (
	"bufio"
	"io"
	"sort"
)

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	seatIds := []int{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		rows := make([]int, 128)
		seat := make([]int, 8)
		for i := range rows {
			rows[i] = i
		}
		for i := range seat {
			seat[i] = i
		}

		pass := scanner.Text()
		for _, p := range pass {
			seatChange := len(seat) / 2
			rowChange := len(rows) / 2
			if p == 'F' { // Lower half
				rows = rows[:rowChange]
			}
			if p == 'B' { // Higher half
				rows = rows[rowChange:]
			}
			if p == 'L' { // Lower half
				seat = seat[:seatChange]
			}
			if p == 'R' { // Higher half
				seat = seat[seatChange:]
			}
		}

		seatId := rows[0]*8 + seat[0]
		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)
	if runPartTwo {
		return seatIds[sort.Search(len(seatIds), func(i int) bool { return seatIds[i] != seatIds[0]+i })] - 1, nil
	}
	return seatIds[len(seatIds)-1], nil
}
