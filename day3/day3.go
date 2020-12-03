package day3

import (
	"bufio"
	"io"
)

type route struct {
	xSlope   int
	ySlope   int
	x        int
	y        int
	treesHit int
}

func newRoute(xSlope, ySlope int) *route {
	return &route{
		xSlope:   xSlope,
		ySlope:   ySlope,
		x:        xSlope,
		y:        0,
		treesHit: 0,
	}
}

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	routes := []*route{}
	routes = append(routes, newRoute(3, 1))
	if runPartTwo {
		routes = append(routes, newRoute(1, 1))
		routes = append(routes, newRoute(5, 1))
		routes = append(routes, newRoute(7, 1))
		routes = append(routes, newRoute(1, 2))
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		for _, route := range routes {
			var treeHit bool
			treeHit, route.x, route.y = processTreeLine(scanner.Text(), route.x, route.y, route.xSlope, route.ySlope)
			if treeHit {
				route.treesHit++
			}
		}
	}
	if runPartTwo {
		total := 1
		for _, route := range routes {
			total = total * route.treesHit
		}
		return total, scanner.Err()
	}
	return routes[0].treesHit, scanner.Err()
}

func processTreeLine(line string, x, y, xSlope, ySlope int) (bool, int, int) {
	treeHit := false
	if y >= ySlope && y%ySlope == 0 {
		if string(line[x%len(line)]) == "#" {
			treeHit = true
		}
		x += xSlope
	}
	y++
	return treeHit, x, y
}
