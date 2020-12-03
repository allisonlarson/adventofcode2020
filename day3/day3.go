package day3

import (
	"bufio"
	"io"
)

type route struct {
	xSlope   int
	ySlope   int
	x        int
	treesHit int
}

func newRoute(xSlope, ySlope int) *route {
	return &route{
		xSlope: xSlope,
		ySlope: ySlope,
		x:      xSlope,
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
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, route := range routes {
			if i >= route.ySlope && i%route.ySlope == 0 {
				if line[route.x%len(line)] == '#' {
					route.treesHit += 1
				}
				route.x += route.xSlope
			}
		}
		i++
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
