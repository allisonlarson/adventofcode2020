package day5

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	seatIds := []int{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		seats := []int{}
		rows := []int{}
		for i := 0; i < 128; i++ {
			seats = append(seats, i)
		}
		for i := 0; i < 8; i++ {
			rows = append(rows, i)
		}

		pass := scanner.Text()
		for _, p := range pass {
			seatChange := len(seats) / 2
			rowChange := len(rows) / 2

			if p == 'F' { // Lower half
				seats = seats[:seatChange]
			}
			if p == 'B' { // Higher half
				seats = seats[seatChange:]
			}
			if p == 'L' { // Lower half
				rows = rows[:rowChange]
			}
			if p == 'R' { // Higher half
				rows = rows[rowChange:]
			}

		}

		seatId := seats[0]*8 + rows[0]
		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)
	if runPartTwo {
		for i := 0; i < len(seatIds)-1; i++ {
			if seatIds[i+1]-seatIds[i] != 1 {
				return seatIds[i] + 1, nil
			}
		}
		return 0, fmt.Errorf("Nothing found")
	}
	return seatIds[len(seatIds)-1], nil
}
